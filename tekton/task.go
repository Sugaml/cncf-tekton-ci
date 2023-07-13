package tekton

import (
	"context"

	"github.com/sirupsen/logrus"
	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TaskGetter interface {
	Tasks(ns string) TaskInterface
}

type TaskInterface interface {
	Create(ctx context.Context, task *v1.Task, opts metav1.CreateOptions) (*v1.Task, error)
}

type task struct {
	client *versioned.Clientset
	ns     string
}

func newTask(client *versioned.Clientset, ns string) *task {
	return &task{
		client: client,
		ns:     ns,
	}
}

func (c *task) Create(ctx context.Context, task *v1.Task, opts metav1.CreateOptions) (*v1.Task, error) {
	task, err := c.client.TektonV1().Tasks(c.ns).Get(ctx, task.Name, metav1.GetOptions{})
	if err != nil {
		result, err := c.client.TektonV1().Tasks(c.ns).Create(context.Background(), task, opts)
		if err != nil {
			return nil, err
		}
		logrus.Infof("Task %s created.", result.Name)
		return result, nil
	}
	taskUpdated := &v1.Task{}
	taskUpdated.Name = task.Name
	taskUpdated.Namespace = task.Namespace
	taskUpdated.Spec = task.Spec
	result, err := c.client.TektonV1().Tasks(c.ns).Update(context.Background(), taskUpdated, metav1.UpdateOptions{})
	if err != nil {
		return task, err
	}
	logrus.Infof("Task %s configured.", result.Name)
	return result, err
}
