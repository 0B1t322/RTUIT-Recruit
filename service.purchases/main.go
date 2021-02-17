package main

import (
	"net/http"

	"github.com/0B1t322/RTUIT-Recruit/service.purchases/router"
)

func main() {	
	if err :=  http.ListenAndServe(":8081", router.New()); err  != nil {
		panic(err)
	}
	
}