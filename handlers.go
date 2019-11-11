package main

import (
  "encoding/json"
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
)

// Map connection to Mongo address
type LinkShortenerAPI struct {
  myconnection *MongoConnection
}

// Create mapping struct, containing shorturl & longurl (returnable in json)
type UrlMapping struct {
  ShortUrl string `json:shorturl`
  LongUrl string `json:longurl`
}

// Create helper struct to provide json status codes
type APIResponse struct {
  StatusMessage string `json:statusmessage`
}

// Create new API object, connected to Mongo
func NewUrlLinkShortenerAPI() *LinkShortenerAPI {
  LS := &LinkShortenerAPI {
    myconnection: NewDBConnection(),
  }
  return LS
}

// Provide content for root
func (Ls *LinkShortenerAPI) UrlRoot(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Welcome to Ali's url shortener API, built in go! \n"+
              "Do a Get request with the short Link to get the long Link \n"+
              "Do a POST request with long Link to get a short Link \n")
}


// Create new url object
func (Ls *LinkShortenerAPI) UrlCreate(w http.ResponseWriter, r *http.Request) {
  reqBodyStruct := new(UrlMapping)
  responseEncoder := json.NewEncoder(w)
  if err := json.NewDecoder(r.Body).Decode(&reqBodyStruct); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    if err := responseEncoder.Encode(&APIResponse{StatusMessage: err.Error()}); err!= nil {
      fmt.Fprintf(w, "Error %s occured while trying to add the url \n", err.Error())
    }
    return
  }
  err := Ls.myconnection.AddUrls(reqBodyStruct.LongUrl, reqBodyStruct.ShortUrl)
  if err != nil {
    w.WriteHeader(http.StatusConflict)
    if err := responseEncoder.Encode(&APIResponse{StatusMessage: err.Error()}); err != nil {
      fmt.Fprintf(w, "Error %s occured while trying to add the url \n", err.Error())
    }
    return
  }
  responseEncoder.Encode(&APIResponse{StatusMessage: "OK"})
}

func (Ls *LinkShortenerAPI) UrlShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  sUrl := vars["shorturl"]
  if len(sUrl) > 0 {
    lUrl, err := Ls.myconnection.FindlongUrl(sUrl)
    if err != nil {
      fmt.Fprint(w, "Could not find saved long url that corresponds to the short url %s \n", sUrl)
      return
    }

    http.Redirect(w, r, lUrl, http.StatusFound)
  }
}
