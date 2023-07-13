package main

import (
	"github.com/babulalt/tekton/src"
	"github.com/babulalt/tekton/utils"
	"github.com/babulalt/tekton/utils/queue"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	}
	log.Info("successfully loaded configuration.")
	messageBroker := queue.NewQueue().GetQueue()
	ch := make(chan bool)
	log.Info("pipeline routine starting")
	functionList := map[string]func(interface{}) error{
		"build-pipeline": func(i interface{}) error {
			plController, _ := src.NewPipeline(utils.GetCIRequest(i))
			return plController.CreatePipeline()
		},
	}
	for name, function := range functionList {
		go messageBroker.Subscribe(name, function)
	}
	<-ch
}
