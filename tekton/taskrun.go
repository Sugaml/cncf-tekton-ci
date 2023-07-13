package tekton

import (
	"context"

	"github.com/sirupsen/logrus"
	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TaskRunGetter interface {
	TaskRuns(ns string) TaskRunInterface
}

type TaskRunInterface interface {
	Create(ctx context.Context, task *v1.TaskRun, opts metav1.CreateOptions) (*v1.TaskRun, error)
}

type taskRun struct {
	client *versioned.Clientset
	ns     string
}

func newTaskRun(client *versioned.Clientset, ns string) *taskRun {
	return &taskRun{
		client: client,
		ns:     ns,
	}
}

func (c *taskRun) Create(ctx context.Context, task *v1.TaskRun, opts metav1.CreateOptions) (*v1.TaskRun, error) {
	task, err := c.client.TektonV1().TaskRuns(c.ns).Get(ctx, task.Name, metav1.GetOptions{})
	if err != nil {
		result, err := c.client.TektonV1().TaskRuns(c.ns).Create(ctx, task, opts)
		if err != nil {
			return nil, err
		}
		logrus.Infof("TaskRun %s created.", result.Name)
		return result, nil
	}
	result, err := c.client.TektonV1().TaskRuns(c.ns).Update(ctx, task, metav1.UpdateOptions{})
	if err != nil {
		return task, err
	}
	logrus.Infof("TaskRun %s configured.", result.Name)
	return result, err
}
