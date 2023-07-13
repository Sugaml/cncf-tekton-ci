package tekton

import (
	"context"

	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PipelineRunInterface has methods to work with PipelineRun resources.
type PipelineRunInterface interface {
	CreatePipelineRun(ctx context.Context, pipelineRun *v1.PipelineRun, opts metav1.CreateOptions) (*v1.PipelineRun, error)
}

func (tkn *TektonController) CreatePipelineRun(ctx context.Context, pipelineRun *v1.PipelineRun, opts metav1.CreateOptions) (*v1.PipelineRun, error) {
	pipelineRun, err := tkn.TektonClient.TektonV1().PipelineRuns(tkn.Input.Namespace).Create(ctx, pipelineRun, opts)
	if err != nil {
		return nil, err
	}
	return pipelineRun, nil
}
