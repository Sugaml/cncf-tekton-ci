package main

import (
	"time"

	"github.com/babulalt/tekton/src"
	"github.com/emicklei/go-restful/log"
)

func main() {
	clinet, err := src.NewTektonClient()
	if err != nil {
		log.Printf("error in create client :: %v", err)
	}
	_, err = clinet.CreateSecrets()
	// CloneAndBuild(clinet, "golang-appv2")
}

func CloneAndBuild(clinet *src.Clients, namespace string) error {
	log.Print("Started Clone and build process...")
	nSpec := src.SetNamespace(namespace)
	err := clinet.CreateNamespace(nSpec)
	if err != nil {
		log.Print("error creating namespace", err)
		return err
	} else {
		log.Print("created namepsace :: ", namespace)
	}

	s, err := clinet.CreateSecret(namespace)
	if err != nil {
		log.Print("error creating create", err)
		return err
	}
	log.Print("created secret :: ", s.Name)

	task := src.Setuptask()
	cTask, err := clinet.CreateTask(namespace, task)
	if err != nil {
		log.Print("error creating task")
		return err
	}
	log.Print("created task :: ", cTask.Name)

	pipelineResources := src.SetupPipeLineResources()
	for _, pipelineResource := range *pipelineResources {
		cPipeResource, err := clinet.CreatePipelineResource(namespace, &pipelineResource)
		if err != nil {
			log.Print("error creating pipeline resource :: ", err)
			return err
		}
		log.Print("created pipeline resource ::", cPipeResource.Name)
	}

	pipeline := src.SetupPipeline()
	cPipeline, err := clinet.CreatePipeline(namespace, pipeline)
	if err != nil {
		log.Print("error creating pipeline", err)
		return err
	}
	log.Print("created pipeline ::", cPipeline.Name)

	pipelineRun := src.SetupPipelineRun()
	cPipelineRun, err := clinet.CreatePipelineRun(namespace, pipelineRun)
	if err != nil {
		log.Print("error creating pipeline", err)
		return err
	}
	log.Print("created pipeline ::", cPipelineRun.Name)

	go clinet.EventWatcher(namespace)
	time.Sleep(1 * time.Minute)
	return nil
}
