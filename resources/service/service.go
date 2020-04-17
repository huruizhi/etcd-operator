package service

import (
	"etcd-operator/api/v1beta1"
	corev1 "k8s.io/api/core/v1"
)

func New(etcd *v1beta1.Etcd) *corev1.Service {
	var headlessService corev1.Service
	return &headlessService
}
