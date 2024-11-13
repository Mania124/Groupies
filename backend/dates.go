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

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

func DatesData(id int) (Date, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%d", id)
	client := http.Client{
		Timeout: 60 * time.Second,
	}

	response, err := client.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	var date Date
	err = json.NewDecoder(response.Body).Decode(&date)
	if err != nil {
		return Date{}, err
	}

	return date, nil
}

func DateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "400: Wrong Request Method", http.StatusMethodNotAllowed)
		return
	}
	if !strings.HasPrefix(r.URL.Path, "/date/") {
		http.Error(w, "404: Not Found", http.StatusNotFound)
		return
	}
	idStr := r.URL.Path[len("/date/"):]
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

	date, err := DatesData(id)
	if err != nil {
		templ, err := template.ParseGlob("frontend/*.html")
		if err != nil {
			templ.ExecuteTemplate(w, "error.html", "500: Oops! Something Went Wrong")
			return
		}
		templ.ExecuteTemplate(w, "error.html", "400: Unable to Load Date Data")
		return
	}

	templ, err := template.ParseGlob("frontend/*.html")
	if err != nil {
		log.Println(err)
		templ.ExecuteTemplate(w, "error.html", "500: Oops! Something Went Wrong")
		return
	}
	templ.ExecuteTemplate(w, "date.html", date)
}
