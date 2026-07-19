package kubernetes

import (
	"context"
	"errors"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListDeployments(namespace string) ([]appsv1.Deployment, error) {

	if namespace == "" {
		namespace = "default"
	}

	deployments, err := Clientset.
		AppsV1().
		Deployments(namespace).
		List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	return deployments.Items, nil
}

func CreateDeployment(namespace string, deployment *appsv1.Deployment) error {

	if Clientset == nil {
		return errors.New("kubernetes client is not initialized")
	}

	if namespace == "" {
		namespace = "default"
	}

	_, err := Clientset.
		AppsV1().
		Deployments(namespace).
		Create(
			context.TODO(),
			deployment,
			metav1.CreateOptions{},
		)

	return err
}

func DeleteDeployment(namespace, name string) error {

	if Clientset == nil {
		return errors.New("kubernetes client is not initialized")
	}

	if namespace == "" {
		namespace = "default"
	}

	return Clientset.
		AppsV1().
		Deployments(namespace).
		Delete(
			context.TODO(),
			name,
			metav1.DeleteOptions{},
		)
}

func ScaleDeployment(namespace, name string, replicas int32) error {

	if Clientset == nil {
		return errors.New("kubernetes client is not initialized")
	}

	if namespace == "" {
		namespace = "default"
	}

	deployment, err := Clientset.
		AppsV1().
		Deployments(namespace).
		Get(
			context.TODO(),
			name,
			metav1.GetOptions{},
		)

	if err != nil {
		return err
	}

	deployment.Spec.Replicas = &replicas

	_, err = Clientset.
		AppsV1().
		Deployments(namespace).
		Update(
			context.TODO(),
			deployment,
			metav1.UpdateOptions{},
		)

	return err
}

func RestartDeployment(namespace, name string) error {

	if Clientset == nil {
		return errors.New("kubernetes client is not initialized")
	}

	if namespace == "" {
		namespace = "default"
	}

	deployment, err := Clientset.
		AppsV1().
		Deployments(namespace).
		Get(
			context.TODO(),
			name,
			metav1.GetOptions{},
		)

	if err != nil {
		return err
	}

	if deployment.Spec.Template.Annotations == nil {
		deployment.Spec.Template.Annotations = make(map[string]string)
	}

	deployment.Spec.Template.Annotations["kubectl.kubernetes.io/restartedAt"] =
		time.Now().Format(time.RFC3339)

	_, err = Clientset.
		AppsV1().
		Deployments(namespace).
		Update(
			context.TODO(),
			deployment,
			metav1.UpdateOptions{},
		)

	return err
}
