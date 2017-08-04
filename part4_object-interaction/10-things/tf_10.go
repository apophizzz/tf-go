package main

import (
	"github.com/apophis90/tf-go/part4_object-interaction/10-things/controller"
)

func main() {
	wfController := controller.WordFrequencyController{}
	wfController.Init("../../input.txt")
	wfController.Run()
}
