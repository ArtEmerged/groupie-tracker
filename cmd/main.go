package main

import (
	"flag"

	"groupie-tracker/internal"
)

func main() {
	port := flag.String("port", "8081", "USAGE --port=8081")

	internal.Running(*port)
}
