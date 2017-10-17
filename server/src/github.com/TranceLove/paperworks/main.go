package main

import (
	"net/http"

	"github.com/TranceLove/paperworks/routes"
	// "github.com/appleboy/gin-jwt"
)

func main(){

	r := routes.Init()

    http.ListenAndServe(":3000", r)
}
