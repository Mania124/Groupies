package backend

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

func LocationsData(id int) (Location, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", id)

	client := http.Client{
		Timeout: 60 * time.Second,
	}

	response, err := client.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	var location Location
	err = json.NewDecoder(response.Body).Decode(&location)
	if err != nil {
		return Location{}, err
	}

	return location, nil
}

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "400: Bad Request Method", http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasPrefix(r.URL.Path, "/location/") {
		http.Error(w, "404: Not Found", http.StatusNotFound)
		return
	}
	idStr := r.URL.Path[len("/location/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		templ, err := template.ParseGlob("frontend/*.html")
		if err != nil {
			templ.ExecuteTemplate(w, "error.html", "500: Oops! Something Went Wrong")
			return
		}
		templ.ExecuteTemplate(w, "error.html", "400: Invalid Artist ID")
		return
	}

	location, err := LocationsData(id)
	if err != nil {
		templ, err := template.ParseGlob("frontend/*.html")
		if err != nil {
			templ.ExecuteTemplate(w, "error.html", "500: Oops! Something Went Wrong")
			return
		}
		templ.ExecuteTemplate(w, "error.html", "400: Unable to Load Location Data")
		return
	}
	templ, err := template.ParseGlob("frontend/*.html")
	if err != nil {
		log.Println(err)
		templ.ExecuteTemplate(w, "error.html", "500: Oops! Something Went Wrong")
		return
	}

	templ.ExecuteTemplate(w, "location.html", location)
}
