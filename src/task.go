package src

import (
	"context"

	beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	r "github.com/tektoncd/pipeline/pkg/apis/resource/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Clients) CreateTask(namespace string, task *beta1.Task) (*beta1.Task, error) {
	task, err := c.TaskClient.Tasks(namespace).Create(context.Background(), task, v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return task, nil
}

func Setuptask() *beta1.Task {
	task := &beta1.Task{
		TypeMeta: v1.TypeMeta{
			APIVersion: "tekton.dev/v1beta1",
			Kind:       "Task",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: "build-docker-image",
		},
		Spec: beta1.TaskSpec{
			Resources: &beta1.TaskResources{
				Inputs: []beta1.TaskResource{
					{
						ResourceDeclaration: r.ResourceDeclaration{
							Name: "git-repo",
							Type: "git",
						},
					},
				},
				Outputs: []beta1.TaskResource{
					{
						ResourceDeclaration: r.ResourceDeclaration{
							Name: "image-registry",
							Type: "image",
						},
					},
				},
			},
			Params: []beta1.ParamSpec{
				{
					Name:        "pathToDockerFile",
					Description: "Dockerfile path ",
					Default:     &beta1.ArrayOrString{StringVal: "Dockerfile", Type: beta1.ParamTypeString},
				},
				{
					Name:        "pathToContext",
					Description: "The build context used by Kaniko",
					Default:     &beta1.ArrayOrString{StringVal: "/workspace/git-repo/", Type: beta1.ParamTypeString},
				},
			},
			Steps: []beta1.Step{
				{
					Name:  "build-and-push",
					Image: "gcr.io/kaniko-project/executor:v0.10.0",
					Env: []corev1.EnvVar{
						{
							Name:  "DOCKER_CONFIG",
							Value: "/builder/home/.docker/",
						},
					},
					Command: []string{"/kaniko/executor"},
					Args:    []string{"--dockerfile=/workspace/git-repo/Dockerfile", "--destination=sugamdocker35/golang-app:v2", "--context=/workspace/git-repo/"},
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "docker-config",
							MountPath: "/builder/home/.docker",
						},
					},
				},
			},
			Volumes: []corev1.Volume{
				{
					Name: "docker-config",
					VolumeSource: corev1.VolumeSource{
						Secret: &corev1.SecretVolumeSource{
							SecretName: "docker-config",
							Items: []corev1.KeyToPath{
								{
									Key:  ".dockerconfigjson",
									Path: "config.json",
								},
							},
						},
					},
				},
			},
		},
	}
	return task
}
