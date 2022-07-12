package main

import "OCsemmerApp/pkg/handler"

func main() {
	router := handler.NewRouter()
	router.Run()
}
