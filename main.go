package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"gitlab-wrap/internal/card"
	"gitlab-wrap/internal/gitlab"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func CardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gitLabWrapStats, err := gitlab.NewGitLabWrapStats(vars["username"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "image/png")
		userNotFound, _ := card.CreateUserNotFound(vars["username"])
		img, _ := os.Open(*userNotFound)
		defer img.Close()
		w.Header().Set("Content-Type", "image/png")
		_, err := io.Copy(w, img)
		if err != nil {
			return
		}
		return
	}
	wrapCardPath, err := card.CreateGitLabWrapCard(gitLabWrapStats)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("could not create card"))
		return
	}
	img, err := os.Open(*wrapCardPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("could not create card"))
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/png")
	_, _ = io.Copy(w, img)
}

func FileHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {

		file, err := os.ReadFile(entrypoint)
		if err != nil {
			return
		}
		vars := mux.Vars(r)
		username := vars["username"]
		stringCardUrl := "placeholder.png"

		fileLocation := entrypoint
		if len(username) > 0 {
			username = url.QueryEscape(username)
			newFile := strings.Replace(string(file), "Your", fmt.Sprintf("%s's", username), -1)
			newFile = strings.Replace(newFile, stringCardUrl, fmt.Sprintf("card/%s", username), -1)
			fileLocation = fmt.Sprintf("%s/index.html", os.TempDir())
			err := os.WriteFile(fileLocation, []byte(newFile), 0644)
			if err != nil {
				log.Println(err)
			}
		}
		http.ServeFile(w, r, fileLocation)
	}
	return fn
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	var static string
	log.Println(fmt.Sprintf("Starting server on port %s", port))
	flag.StringVar(&static, "static", "dist", "the directory to serve static files from.")
	flag.Parse()
	log.Println("Serving static files from", static)
	r := mux.NewRouter()
	r.HandleFunc("/card/{username}", CardHandler)

	r.HandleFunc("/placeholder.png", FileHandler(static+"/placeholder.png"))
	r.HandleFunc("/robots.txt", FileHandler(static+"/robots.txt"))

	r.HandleFunc("/{username}", FileHandler(static+"/index.html"))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(static)))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
