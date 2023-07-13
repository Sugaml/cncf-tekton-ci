package tekton

import (
	"context"

	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PipelineInterface has methods to work with Pipeline resources.
type PipelineInterface interface {
	Create(ctx context.Context, pipeline *v1.Pipeline, opts metav1.CreateOptions) (*v1.Pipeline, error)
}

func (c *TektonController) Create(ctx context.Context, pipeline *v1.Pipeline, opts metav1.CreateOptions) (*v1.Pipeline, error) {
	result, err := c.TektonClient.TektonV1().Pipelines(c.Input.Namespace).Create(ctx, pipeline, metav1.CreateOptions{})
	return result, err
}
