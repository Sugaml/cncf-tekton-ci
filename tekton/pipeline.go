package tekton

import (
	"context"

	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PipelineGetter interface {
	Pipelines(namespace string) PipelineInterface
}

type pipeline struct {
	client *versioned.Clientset
	ns     string
}

func newPipeline(client *versioned.Clientset, ns string) *pipeline {
	return &pipeline{
		client: client,
		ns:     ns,
	}
}

// PipelineInterface has methods to work with Pipeline resources.
type PipelineInterface interface {
	Create(ctx context.Context, pipeline *v1.Pipeline, opts metav1.CreateOptions) (*v1.Pipeline, error)
}

func (c *pipeline) Create(ctx context.Context, pipeline *v1.Pipeline, opts metav1.CreateOptions) (*v1.Pipeline, error) {
	result, err := c.client.TektonV1().Pipelines(c.ns).Create(ctx, pipeline, metav1.CreateOptions{})
	return result, err
}
