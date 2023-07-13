package tekton

import (
	"github.com/babulalt/tekton/models"
)

type TektonController struct {
	models.Pipeline
}

type TektonInterface interface {
	TaskGetter
	TaskRunGetter
	PipelineGetter
	PipelineRunGetter
}

func NewTektonController(pipeline models.Pipeline) TektonInterface {
	return &TektonController{Pipeline: pipeline}
}

func (c *TektonController) Tasks(namespace string) TaskInterface {
	return newTask(c.TektonClient, namespace)
}

func (c *TektonController) Pipelines(namespace string) PipelineInterface {
	return newPipeline(c.TektonClient, namespace)
}

func (c *TektonController) PipelineRuns(namespace string) PipelineRunInterface {
	return newPipelineRun(c.TektonClient, namespace)
}

func (c *TektonController) TaskRuns(namespace string) TaskRunInterface {
	return newTaskRun(c.TektonClient, namespace)
}
