package tekton

import (
	"context"

	"github.com/sirupsen/logrus"
	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TaskGetter interface {
	Task(ns string) TaskInterface
}

type TaskInterface interface {
	Create(ctx context.Context, task *v1.Task, opts metav1.CreateOptions) (*v1.Task, error)
}

type task struct {
	clientset *versioned.Clientset
	ns        string
}

func newTask(clientset *versioned.Clientset, ns string) *task {
	return &task{
		clientset: clientset,
		ns:        ns,
	}
}

func (c *task) Create(ctx context.Context, task *v1.Task, opts metav1.CreateOptions) (*v1.Task, error) {
	result, err := c.clientset.TektonV1().Tasks(c.ns).Create(context.Background(), task, metav1.CreateOptions{})
	if err != nil {
		logrus.Error("error create task :: ", err)
	}
	return result, err
}
