package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/action_page.html", handleFormSubmission)

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func handleFormSubmission(w http.ResponseWriter, r *http.Request) {
	if true {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}

		website := r.Form.Get("website")
		resp, err := http.Get(website)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Fprintf(w, "<h1> Website: %v </h1>", resp.Body)
		fmt.Fprintf(w, "<h2> Website but smaller font: %v </h2>", resp.Body)
		// TODO - figure out why we're getting a - Get "": unsupported protocol scheme "" - error
	}
}
