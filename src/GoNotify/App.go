package main

import (
	"fmt"
	"log"
	"net/http"
	"config"
	"GoNotify/router"
	"GoNotify/register"
)

func main() {
	name := "Go Notify"
	register.Init()
	configuration, err := config.LoadConfiguration("./src/GoNotify/config.json")
	if err == nil {
		routers := router.CreateAllRouter()
		fmt.Println(name, "Server Ready on port" , configuration.Port)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", configuration.Port), routers); err != nil {
			log.Fatal("ListenAndServe:", err)
		}
	} else {
		fmt.Println("Cannot load Configuration: ", err)
	}
}
