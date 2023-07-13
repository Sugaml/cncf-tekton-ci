package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (kubectl *KubernetesController) CreateSecret(secret *v1.Secret) (*v1.Secret, error) {
	csecret, err := kubectl.ClientSet.CoreV1().Secrets(kubectl.Input.Namespace).Create(context.Background(), secret, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return csecret, nil
}
