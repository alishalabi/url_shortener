package main

import (
  "encoding/json"
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
)

// Map connection to Mongo address
type LinkShortnerAPI struct {
  myconnection *MongoConnection
}

// Create mapping struct, containing shorturl & longurl (returnable in json)
type UrlMapping struct {
  ShortURL string `json:shorturl`
  LongUrl string `json:longurl`
}

// Create helper struct to provide json status codes
type APIResponse struct {
  StatusMessage string `json:statusmessage`
}

// Create new API object, connected to Mongo
func NewUrlLinkShortenerAPI() *LinkShortnerAPI {
  LS := &LinkShortnerAPI {
    myconnection: NEWDBConnection(),
  }
  return LS
}

// Provide content for root
func (Ls *LinkShortnerAPI) UrlRoot(w hhtp.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Welcome to Ali's url shortener API, built in go! \n"+
              "Do a Get request with the short Link to get the long Link \n"+
              "Do a POST request with long Link to get a short Link \n"))
}


// Create new url object
func (Ls *LinkShortnerAPI) UrlCreate(w http.ResponseWriter, r *http.Request) {
  
}
