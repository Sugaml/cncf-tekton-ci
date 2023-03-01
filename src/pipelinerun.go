package src

import (
	"context"

	beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Clients) CreatePipelineRun(namespace string, pipelineRun *beta1.PipelineRun) (*beta1.PipelineRun, error) {
	pipelineRun, err := c.PipelineRunClient.PipelineRuns(namespace).Create(context.Background(), pipelineRun, v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return pipelineRun, nil
}

func SetupPipelineRun() *beta1.PipelineRun {
	pipelineRun := &beta1.PipelineRun{
		TypeMeta: v1.TypeMeta{
			APIVersion: "tekton.dev/v1beta1",
			Kind:       "PipelineRun",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: "pipeline-run-for-demo",
		},
		Spec: beta1.PipelineRunSpec{
			PipelineRef: &beta1.PipelineRef{
				Name: "pipeline-for-demo",
			},
			Resources: []beta1.PipelineResourceBinding{
				{
					Name: "git-repo",
					ResourceRef: &beta1.PipelineResourceRef{
						Name: "git-repo",
					},
				},
				{
					Name: "image-registry",
					ResourceRef: &beta1.PipelineResourceRef{
						Name: "image-registry",
					},
				},
			},
		},
	}
	return pipelineRun
}
