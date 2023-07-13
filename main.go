package main

import (
	"github.com/babulalt/tekton/models"
	"github.com/babulalt/tekton/src"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		log.Info("loading env variables")
	}
	var mapFiles map[string]interface{} = map[string]interface{}{}

	input := &models.CIRequest{
		Name:          "tekton",
		EnvironmentId: 6036,
		Namespace:     "tekton-cistep-1",
		// ConfigPath:        "/home/sugam/Downloads/ovh-cluster/ovhkube",
		ConfigPath:        "/home/sugam/.kube/config",
		GitUrl:            "https://github.com/babulalt/simple-go.git",
		GitBranch:         "simple-task",
		ImageRepoUsername: "sugamdocker35",
		ImageRepoPassword: "Docker123",
		GitUserName:       "babulalt",
		GitAccessToken:    "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",
		ImageRepoProject:  "sugamdocker35",
		RepositoryImage: models.RepositoryImage{
			Repository: "tekton-abc",
			Name:       "tekton-ci",
			Tag:        "latest",
		},
		BaseImage:   "golang",
		BaseTag:     "1.16",
		BuildScript: "FROM golang:1.16-alpine\nRUN mkdir /app\nADD . /app\nWORKDIR /app\nRUN go build -o main .\nCMD ['/app/main']",
		//CISteps:   steps,
		CiFiles:   &mapFiles,
		PluginUrl: "./data",
	}
	pl, err := src.NewPipeline(input)
	if err != nil {
		log.Errorf("Initalize pipeline error :: %v", err)
		return
	}
	err = pl.CreatePipeline()
	if err != nil {
		log.Errorf("create pipeline error :: %v", err)
		return
	}
}
