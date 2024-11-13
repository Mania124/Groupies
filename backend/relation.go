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

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func RelationsData(id int) (Relation, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", id)

	client := http.Client{
		Timeout: 60 * time.Second,
	}

	response, err := client.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	var relation Relation
	err = json.NewDecoder(response.Body).Decode(&relation)
	if err != nil {
		return Relation{}, err
	}

	return relation, nil
}

func RelationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "400: Bad Request Method", http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasPrefix(r.URL.Path, "/relation/") {
		http.Error(w, "404: Not Found", http.StatusNotFound)
		return
	}
	idStr := r.URL.Path[len("/relation/"):]
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

	relation, err := RelationsData(id)
	if err != nil {
		http.Error(w, "400: Unable to load relation data", http.StatusInternalServerError)
		return
	}

	templ, err := template.ParseGlob("frontend/*.html")
	if err != nil {
		log.Println(err)
		templ.ExecuteTemplate(w, "error.html", "500: Oops! Something Went Wrong")
		return
	}

	templ.ExecuteTemplate(w, "relation.html", relation)
}
