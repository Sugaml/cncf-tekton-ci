package tekton

import (
	"context"

	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PipelineRunGetter interface {
	PipelineRuns(namespace string) PipelineRunInterface
}

type pipelineRun struct {
	client *versioned.Clientset
	ns     string
}

func newPipelineRun(client *versioned.Clientset, ns string) *pipelineRun {
	return &pipelineRun{
		client: client,
		ns:     ns,
	}
}

// PipelineRunInterface has methods to work with PipelineRun resources.
type PipelineRunInterface interface {
	CreatePipelineRun(ctx context.Context, pipelineRun *v1.PipelineRun, opts metav1.CreateOptions) (*v1.PipelineRun, error)
}

func (tkn *pipelineRun) CreatePipelineRun(ctx context.Context, pipelineRun *v1.PipelineRun, opts metav1.CreateOptions) (*v1.PipelineRun, error) {
	pipelineRun, err := tkn.client.TektonV1().PipelineRuns(tkn.ns).Create(ctx, pipelineRun, opts)
	if err != nil {
		return nil, err
	}
	return pipelineRun, nil
}
