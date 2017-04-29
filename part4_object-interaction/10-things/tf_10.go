package main

import (
	"github.com/PaddySmalls/golang_term-frequency-styles/part4_object-interaction/10-things/controller"
)

func main() {
	wfController := controller.WordFrequencyController{}
	wfController.Init("input.txt")
	wfController.Run()
}


