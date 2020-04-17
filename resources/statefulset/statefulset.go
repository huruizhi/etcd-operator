package statefulset

import (
	"etcd-operator/api/v1beta1"
	appv1 "k8s.io/api/apps/v1"
)

func New(etcd *v1beta1.Etcd) *appv1.StatefulSet {
	var sts appv1.StatefulSet
	return &sts
}
