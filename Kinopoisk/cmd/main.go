package main

import (
	"lenrek88/internal/app/server"
	"log"
)

func main() {
	r := server.NewGin()
	log.Println("listen: 8080")
	_ = r.Run(":8080")
}
