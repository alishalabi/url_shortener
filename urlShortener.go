package main

import (
  "net/http"
)

func main() {
  // New API Shortner
  LinkShortener := NewUrlLinkShortenerAPI()
  // Connect Routes
  routes := CreateRoutes(LinkShortener)
  // Initiate API routers
  router := NewLinkShortenerRouter(routes)
  // Start program on local port 5100
  http.ListenAndServe(":5100", router)
}
