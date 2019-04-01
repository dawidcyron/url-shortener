package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/go-redis/redis"

	"github.com/go-chi/chi"

	"github.com/dawidcyron/base62"
	"github.com/dawidcyron/shortener/database"
	"github.com/valyala/fastjson"
)

type shortenResponse struct {
	Message string `json:"message,omitempty"`
	URL     string `json:"url,omitempty"`
}

const invalidURL = "The given string is not a valid URL"

func isValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	} else {
		return true
	}
}

//ShortenURL extracts URL to shorten from request, generates unique ID from Redis
//, encodes it in base62 format and saves to Redis database with base62 ID as key
//and given URL as value
func ShortenURL(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	urlToShorten := fastjson.GetString(body, "url")
	if !isValidURL(urlToShorten) {
		http.Error(w, invalidURL, http.StatusBadRequest)
		return
	}
	id := database.RedisClient.Incr(":id").Val()
	base62ID := base62.ToBase62(int(id))
	err = database.RedisClient.Set(base62ID, urlToShorten, 0).Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := &shortenResponse{}
	response.Message = "Success"
	response.URL = r.Host + "/" + base62ID
	json.NewEncoder(w).Encode(response)
}

//GetFullURL extracts base62 encoded ID from request path and uses it to search for
//saved URL in Redis. If URL with given ID exists, redirects the user.
func GetFullURL(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	base62ID := chi.URLParam(r, "id")
	fullURL, err := database.RedisClient.Get(base62ID).Result()
	if err == redis.Nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, fullURL, 301)
}
