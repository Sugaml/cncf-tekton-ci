package k8s

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NamespaceGetter interface {
	Namespace(ns string) NamespaceInterface
}

type NamespaceInterface interface {
	Create(ctx context.Context, opts metav1.CreateOptions) error
}

type namesapce struct {
	clientset *kubernetes.Clientset
	ns        string
}

func newNamespace(clientset *kubernetes.Clientset, ns string) *namesapce {
	return &namesapce{
		clientset: clientset,
		ns:        ns,
	}
}

func (c *namesapce) Create(ctx context.Context, opts metav1.CreateOptions) error {
	_, err := c.clientset.CoreV1().Namespaces().Get(ctx, c.ns, metav1.GetOptions{})
	if err == nil {
		return nil
	}
	ns := setNamespace(c.ns)
	_, err = c.clientset.CoreV1().Namespaces().Create(ctx, ns, opts)
	if err != nil {
		fmt.Printf("unable to create namespace%v", err)
		return err
	}
	return nil
}

func setNamespace(namespace string) *apiv1.Namespace {
	nsSpec := &apiv1.Namespace{ObjectMeta: metav1.ObjectMeta{
		Namespace: namespace,
		Name:      namespace,
	}}
	return nsSpec
}
