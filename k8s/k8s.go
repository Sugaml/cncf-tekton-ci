package k8s

import "github.com/babulalt/tekton/models"

type KubernetesController struct {
	models.Pipeline
}

type KubernetesInterface interface {
	NamespaceGetter
}

func NewKubernetesController(pipeline models.Pipeline) KubernetesInterface {
	return &KubernetesController{Pipeline: pipeline}
}

func (c *KubernetesController) Namespace(namespace string) NamespaceInterface {
	return newNamespace(c.ClientSet, namespace)
}
