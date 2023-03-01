package src

import (
	"bytes"
	"context"
	"fmt"
	"io"

	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

func (c Clients) CreateNamespace(nsSpec *apiv1.Namespace) error {
	_, err := c.KubeClient.CoreV1().Namespaces().Get(context.Background(), nsSpec.Name, metav1.GetOptions{})
	if err == nil {
		return nil
	}

	_, err = c.KubeClient.CoreV1().Namespaces().Create(context.Background(), nsSpec, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("unable to create namespace%v", err)
		return err
	}
	return nil
}

func SetNamespace(namespace string) *apiv1.Namespace {
	nsSpec := &apiv1.Namespace{ObjectMeta: metav1.ObjectMeta{
		Namespace: namespace,
		Name:      namespace,
	}}
	return nsSpec
}

func (c Clients) EventWatcher(namespace string) {
	labels := ""
	ctx := context.Background()

	watcher, err := c.KubeClient.CoreV1().Pods(namespace).Watch(ctx, metav1.ListOptions{
		LabelSelector: labels,
	})
	if err != nil {
		return
	}
	fmt.Println("Workflow :: ")
	for {
		select {
		case e := <-watcher.ResultChan():
			if e.Object == nil {
				fmt.Println("nil object")
				return
			}
			pod, ok := e.Object.(*v1.Pod)
			if !ok {
				continue
			}
			switch e.Type {
			case watch.Modified:
				if pod.DeletionTimestamp != nil {
					continue
				}
				switch pod.Status.Phase {
				default:
					for _, con := range pod.Spec.Containers {
						req := c.KubeClient.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &v1.PodLogOptions{
							Container: con.Name,
						})
						fmt.Println(con.Name)
						log := ""
						fmt.Println("pod req ", req)
						podLogs, err := req.Stream(ctx)
						fmt.Println("pod log ", &podLogs)
						if err != nil {
							fmt.Println("Error occure ", err)
							log = "error in opening stream"
						} else {
							buf := new(bytes.Buffer)
							_, err = io.Copy(buf, podLogs)
							if err != nil {
								log = "error in copy information from podLogs to buf"
							} else {
								log = buf.String()
								fmt.Println("logged ", log)
								fmt.Println("Workflow --------")
								fmt.Println(buf.String())
							}
						}
					}
				}
			}
		case <-ctx.Done():
			fmt.Println("context done")
			watcher.Stop()
			return
		}
	}
}
