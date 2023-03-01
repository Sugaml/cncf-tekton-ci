package src

import (
	"context"

	"github.com/tektoncd/pipeline/pkg/apis/resource/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Clients) CreatePipelineResource(namespace string, pipelineResource *v1alpha1.PipelineResource) (*v1alpha1.PipelineResource, error) {
	pipelineResource, err := c.PipelineResourceClient.PipelineResources(namespace).Create(context.Background(), pipelineResource, v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return pipelineResource, nil
}

func SetupPipeLineResources() *[]v1alpha1.PipelineResource {
	pipelineResources := &[]v1alpha1.PipelineResource{
		{
			TypeMeta: v1.TypeMeta{
				APIVersion: "tekton.dev/v1alpha1",
				Kind:       "PipelineResource",
			},
			ObjectMeta: v1.ObjectMeta{
				Name: "git-repo",
			},
			Spec: v1alpha1.PipelineResourceSpec{
				Type: "git",
				Params: []v1alpha1.ResourceParam{
					{
						Name:  "revision",
						Value: "main",
					},
					{
						Name:  "url",
						Value: "https://github.com/babulalt/go-websocket.git",
					},
				},
			},
		},
		{
			TypeMeta: v1.TypeMeta{
				APIVersion: "tekton.dev/v1alpha1",
				Kind:       "PipelineResource",
			},
			ObjectMeta: v1.ObjectMeta{
				Name: "image-registry",
			},
			Spec: v1alpha1.PipelineResourceSpec{
				Type: "image",
				Params: []v1alpha1.ResourceParam{
					{
						Name:  "url",
						Value: "sugamdocker35/golang-app:v2",
					},
				},
			},
		},
	}
	return pipelineResources
}

// func SetupPipeLineResourceImage() *v1alpha1.PipelineResource {
// 	pipelineResourceImage := &v1alpha1.PipelineResource{
// 		TypeMeta: v1.TypeMeta{
// 			APIVersion: "tekton.dev/v1alpha1",
// 			Kind:       "PipelineResource",
// 		},
// 		ObjectMeta: v1.ObjectMeta{
// 			Name: "image-registry",
// 		},
// 		Spec: v1alpha1.PipelineResourceSpec{
// 			Type: "image",
// 			Params: []v1alpha1.ResourceParam{
// 				{
// 					Name:  "url",
// 					Value: "sugamdocker35/go-app:latest",
// 				},
// 			},
// 		},
// 	}
// 	return pipelineResourceImage
// }
