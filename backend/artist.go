package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var buffer bytes.Buffer

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func ArtistsData(url string) ([]Artist, error) {
	url = "https://groupietrackers.herokuapp.com/api/artists"

	client := http.Client{
		Timeout: 60 * time.Second,
	}

	response, err := client.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	var artists []Artist
	err = json.NewDecoder(response.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func ArtistData(id int) (Artist, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", id)

	client := http.Client{
		Timeout: 60 * time.Second,
	}

	response, err := client.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	var artist Artist
	err = json.NewDecoder(response.Body).Decode(&artist)
	if err != nil {
		return Artist{}, err
	}

	return artist, nil
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "400: Bad Request Method", http.StatusMethodNotAllowed)
		return
	}
	if !strings.HasPrefix(r.URL.Path, "/artist/") {
		templ, err := template.ParseGlob("frontend/*.html")
		if err != nil {
			log.Println(err)
			templ.ExecuteTemplate(w, "error.html", "500: Oops! Something went wrong")
			return
		}
		templ.ExecuteTemplate(w, "error.html", "404: Not Found")
		return
	}
	idStr := r.URL.Path[len("/artist/"):]
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

	artist, err := ArtistData(id)
	if err != nil || len(artist.Name) == 0 {
		templ, err := template.ParseGlob("frontend/*.html")
		if err != nil {
			templ.ExecuteTemplate(w, "error.html", "500: Oops! Something Went Wrong")
			return
		}
		templ.ExecuteTemplate(w, "error.html", "400: Unable to Load Artist Data!")
		return
	}

	templ, err := template.ParseGlob("frontend/*.html")
	if err != nil {
		log.Println(err)
		templ.ExecuteTemplate(w, "error.html", "500: Oops! Something Went Wrong")
		return
	}

	templ.ExecuteTemplate(w, "artist.html", artist)
}

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "400: Bad Request Method", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		templ, err := template.ParseGlob("frontend/*.html")
		if err != nil {
			log.Println(err)
			templ.ExecuteTemplate(w, "error.html", "500: Oops! Something went wrong")
			return
		}
		templ.ExecuteTemplate(w, "error.html", "404: Not Found")
		return
	}
	artists, err := ArtistsData("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		templ, err := template.ParseGlob("frontend/*.html")
		if err != nil {
			templ.ExecuteTemplate(w, "error.html", "500: Oops! Somthing Went Wrong")
		}
		templ.ExecuteTemplate(w, "error.html", "500: Oops! Something Went Wrong")
		return
	}

	templ, err := template.ParseGlob("frontend/*.html")
	if err != nil {
		log.Println(err)
		templ.ExecuteTemplate(w, "error.html", "500: Oops! Somthing went wrong\nPlease: Reload")
		return
	}
	_, errr := os.Stat("frontend/index.html")
	if os.IsNotExist(errr) {
		// http.NotFound(w, r)
		// return
		templ.ExecuteTemplate(w, "error.html", "404: Not Found")
		return
	}
	errs := templ.ExecuteTemplate(&buffer, "index.html", artists)
	if errs != nil {
		log.Println(errs)
		templ.ExecuteTemplate(w, "error.html", "500: Oops! Somthing went wrong\nPlease: Reload")
		return
	}
	_, er := buffer.WriteTo(w)
	if er != nil {
		http.Error(w, "500: Oops! Something went wrong", http.StatusInternalServerError)
		log.Println("500 Internal Server Error")
		return
	}
}
