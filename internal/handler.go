package internal

import (
	"encoding/json"
	"log"
	"net/http"
)

type UrlShortener struct {
	db DB
}

func NewUrlShortener(db DB) *UrlShortener {
	return &UrlShortener{db: db}
}

func (u *UrlShortener) HandleGet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	link, err := u.db.Get(r.Context(), id)

	if err != nil {
		http.Error(w, "Link not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", link)
	w.WriteHeader(http.StatusMovedPermanently)
}

type Input struct {
	Url string `json:"url"`
}

type Output struct {
	Id string `json:"id"`
}

func (u *UrlShortener) HandlePost(w http.ResponseWriter, r *http.Request) {
	var input Input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var output Output
	for {
		id := generateString(8)
		err := u.db.Save(r.Context(), id, input.Url)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		output = Output{Id: id}
		break
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
