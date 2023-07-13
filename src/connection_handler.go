package src

import (
	"context"
	"time"

	"github.com/babulalt/tekton/k8s"
	"github.com/babulalt/tekton/models"
	"github.com/babulalt/tekton/tekton"
	"github.com/babulalt/tekton/utils/tekton_helper"
	"github.com/sirupsen/logrus"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PipelineController struct {
	models.Pipeline
}

func NewPipeline(input *models.CIRequest) (*PipelineController, error) {
	pipeline := new(PipelineController)
	pipeline.Input = input
	pipeline.GetConfig()
	return pipeline, nil
}

// GetConfig create the kubernetes connection and firestore connection and return the clients
func (p *PipelineController) GetConfig() error {
	// connecting the cluster
	config, clientSet, err := p.InitCluster()
	if err != nil {
		return err
	}
	p.ClientSet = clientSet
	tknClient, err := versioned.NewForConfig(config)
	if err != nil {
		return err
	}
	p.TektonClient = tknClient
	return nil
}

func (p *PipelineController) CreatePipeline() error {
	kubectl := k8s.NewKubernetesController(p.Pipeline)
	tknctl := tekton.NewTektonController(p.Pipeline)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	err := kubectl.Namespace(p.Input.Namespace).Create(ctx, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	task, err := tekton_helper.DecodeYamlToTask()
	if err != nil {
		return err
	}
	reult, err := tknctl.Task(p.Input.Namespace).Create(ctx, task, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	logrus.Info(reult)
	return nil
}
