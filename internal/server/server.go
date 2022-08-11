package server

import (
	"github.com/Timurshk/internal/hanglers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Server() {
	router := httprouter.New()
	router.POST("/", hanglers.PostURL)
	router.GET("/:id", hanglers.GetURL)
	http.ListenAndServe("localhost:8080", router)
}
