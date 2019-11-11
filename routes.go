package main

import (
  "net/http"
)

type Route struct {
  Name string
  Method string
  Pattern string
  HandlerFunc http.HandlerFunc
}

// Instantiate empty routes slice
type Routes []Route

// Populate that slice with the routes and urls we want
func CreateRoutes(LS *LinkShortenerAPI) Routes {
  return Routes{
    Route {
      "UrlRoot",
      "GET",
      "/",
      LS.UrlRoot,
    },
    Route {
      "UrlShow",
      "GET",
      "/{shorturl}",
      LS.UrlShow,
    },
    Route {
      "UrlCreate",
      "POST",
      "/Create",
      LS.UrlCreate,
    },
  }
}
