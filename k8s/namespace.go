package k8s

import (
	"context"

	"github.com/sirupsen/logrus"
	apiv1 "k8s.io/api/core/v1"
	k8serr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	if err != nil {
		if k8serr.IsNotFound(err) {
			ns := setNamespace(c.ns)
			_, err = c.clientset.CoreV1().Namespaces().Create(ctx, ns, v1.CreateOptions{})
			if err != nil {
				return err
			}
			logrus.Infof("Namespace %s created.", ns.Name)
			return nil
		}
		return err
	}
	updateNs := setNamespace(c.ns)
	_, err = c.clientset.CoreV1().Namespaces().Update(ctx, updateNs, v1.UpdateOptions{})
	if err != nil {
		return err
	}
	logrus.Infof("Namespace %s configured.", updateNs.Name)
	return nil
}

func setNamespace(namespace string) *apiv1.Namespace {
	nsSpec := &apiv1.Namespace{ObjectMeta: metav1.ObjectMeta{
		Namespace: namespace,
		Name:      namespace,
	}}
	return nsSpec
}
