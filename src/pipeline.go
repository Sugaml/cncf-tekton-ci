package src

import (
	"context"

	beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Clients) CreatePipeline(namespace string, pipeline *beta1.Pipeline) (*beta1.Pipeline, error) {
	pipeline, err := c.PipelineClient.Pipelines(namespace).Create(context.Background(), pipeline, v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return pipeline, nil
}

func SetupPipeline() *beta1.Pipeline {
	pipeline := &beta1.Pipeline{
		TypeMeta: v1.TypeMeta{
			APIVersion: "tekton.dev/v1beta1",
			Kind:       "Pipeline",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: "pipeline-for-demo",
		},
		Spec: beta1.PipelineSpec{
			Resources: []beta1.PipelineDeclaredResource{
				{
					Name: "git-repo",
					Type: "git",
				},
				{
					Name: "image-registry",
					Type: "image",
				},
			},
			Tasks: []beta1.PipelineTask{
				{
					Name: "build-docker-images",
					TaskRef: &beta1.TaskRef{
						Name: "build-docker-image",
					},
					Params: []beta1.Param{
						{
							Name:  "pathToDockerFile",
							Value: *beta1.NewArrayOrString("/workspace/git-repo/Dockerfile"),
						},
						{
							Name:  "pathToContext",
							Value: *beta1.NewArrayOrString("/workspace/git-repo"),
						},
					},
					Resources: &beta1.PipelineTaskResources{
						Inputs: []beta1.PipelineTaskInputResource{
							{
								Name:     "git-repo",
								Resource: "git-repo",
							},
						},
						Outputs: []beta1.PipelineTaskOutputResource{
							{
								Name:     "image-registry",
								Resource: "image-registry",
							},
						},
					},
				},
			},
		},
	}
	return pipeline
}
