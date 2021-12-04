package main

import (
	_ "embed"
	"github.com/gorilla/mux"
	"gitlab-wrap/internal/card"
	"gitlab-wrap/internal/gitlab"
	"io"
	"log"
	"net/http"
	"os"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gitLabWrapStats, err := gitlab.NewGitLabWrapStats(vars["username"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "image/png")
		userNotFound, _ := card.CreateUserNotFound(vars["username"])
		img, _ := os.Open(*userNotFound)
		defer img.Close()
		w.Header().Set("Content-Type", "image/png")
		io.Copy(w, img)
		return
	}
	wrapCardPath, err := card.CreateGitLabWrapCard(gitLabWrapStats)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not create card"))
		return
	}
	img, err := os.Open(*wrapCardPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not create card"))
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/png")
	io.Copy(w, img)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/card/{username}", YourHandler)

	// Serve static assets directly.
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/")))

	log.Fatal(http.ListenAndServe(":8080", r))
}
