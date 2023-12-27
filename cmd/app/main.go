package main

import "github.com/AlexanderKolesnkov/sirota_kazanskaya/internal/app"

func main() {
	service := app.New()
	service.Init()
}
