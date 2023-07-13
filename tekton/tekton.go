package tekton

import (
	"github.com/babulalt/tekton/models"
)

type TektonController struct {
	models.Pipeline
}

type TektonInterface interface {
	TaskGetter
	// PipelineInterface
	// PipelineRunInterface
}

func NewTektonController(pipeline models.Pipeline) TektonInterface {
	return &TektonController{Pipeline: pipeline}
}

func (c *TektonController) Task(namespace string) TaskInterface {
	return newTask(c.TektonClient, namespace)
}
