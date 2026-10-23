package main

import (
	"fmt"

	"github.com/DrummDaddy/buildservice/internal/handlers"
)

func main() {
	fmt.Println("BuildService API starting...")

	hName := handlers.GetProjectName()
	fmt.Println(hName)
}
