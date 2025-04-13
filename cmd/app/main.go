package main

import "github.com/lpernett/godotenv"

func main() {
	log := NewLogger()
	err := godotenv.Load()
}
