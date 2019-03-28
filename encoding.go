package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dawidcyron/base62"
	"github.com/dawidcyron/shortener/database"
	"github.com/valyala/fastjson"
)

type ShortenResponse struct {
	Message string `json: "message"`
	URL     string `json: "url"`
}

//ShortenURL extracts URL to shorten from request, generates unique ID from Redis
//, encodes it in base62 format and saves to Redis database with base62 ID as key
//and given URL as value
func ShortenURL(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	urlToShorten := fastjson.GetString(body, "url")
	id := database.RedisClient.Incr(":id").Val()
	base62ID := base62.ToBase62(int(id))
	err = database.RedisClient.Set(base62ID, urlToShorten, 0).Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	response := &ShortenResponse{}
	response.Message = "Success"
	response.URL = r.Host + "/" + base62ID
	json.NewEncoder(w).Encode(response)
}
