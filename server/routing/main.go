package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	})

	r.Get("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		n := chi.URLParam(r, "name")
		fmt.Fprintln(w, n)
	})

	r.Get("/sleep", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(30 * time.Second)
	})

	r.Post("/json", func(w http.ResponseWriter, r *http.Request) {
		// curl http://127.0.0.1:8080/json -d '{"alma":"value1","fa":"value2","harom":2}' -H "Content-Type: application/json"

		var input struct {
			Alma  string `json:"alma"`
			Fa    string `json:"fa"`
			Harom int    `json:"harom"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err)
			return
		}

		output := struct {
			Response string `json:"response"`
			Alma     string `json:"alma"`
		}{
			Response: "hello",
			Alma:     input.Alma,
		}

		err = json.NewEncoder(w).Encode(output)
		if err != nil {
			log.Println(err)
			return
		}
	})

	r.Group(func(r chi.Router) {
		//r.Use(authStuff)
		r.Use(authStuffWithCreds("u1", "p1"))
		r.Get("/protected", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "protected!")
		})
	})

	//http.ListenAndServe(":8080", r)
	server := &http.Server{Addr: ":8080", Handler: r}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGUSR1)
	go func() {
		<-c
		fmt.Println("signal!")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		server.Shutdown(ctx)
	}()

	server.ListenAndServe()
}

func authStuff(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok {
			fmt.Fprintln(w, "no basic auth header")
			return
		}

		if u == "alma" && p == "fa" {
			h.ServeHTTP(w, r)
			return
		}

		fmt.Fprintln(w, "invalid user or password")
	})
}

func authStuffWithCreds(user, pass string) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u, p, ok := r.BasicAuth()
			if !ok {
				fmt.Fprintln(w, "no basic auth header")
				return
			}

			if u == user && p == pass {
				next.ServeHTTP(w, r)
				return
			}

			fmt.Fprintln(w, "invalid user or password")
		})
	}
}
