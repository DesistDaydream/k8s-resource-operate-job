package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

// OpeateTarget 想要操作的资源对象的信息
type OpeateTarget struct {
	Namespace  string
	ObjectName string
	CRBaseInfo schema.GroupVersionResource
}

// NewDeleteTarget 实例化 OpeateTarget
func NewOpeateTarget() *OpeateTarget {
	return &OpeateTarget{
		Namespace:  "rabbitmq",
		ObjectName: "rabbitmq-bj-test",
		CRBaseInfo: schema.GroupVersionResource{
			Group:    "rabbitmq.com",
			Version:  "v1beta1",
			Resource: "rabbitmqclusters",
		},
	}
}

// CreateTarget 创建一个对象

// DeleteTarget 删除一个对象
func (t *OpeateTarget) DeleteTarget(config *rest.Config) {
	clientset, _ := dynamic.NewForConfig(config)
	if err := t.delete(clientset); err != nil {
		fmt.Printf("namespace:%v\nerror:%v\n", t.Namespace, err)
	}
}

func (t *OpeateTarget) delete(clientset dynamic.Interface) error {
	return clientset.Resource(t.CRBaseInfo).Namespace(t.Namespace).Delete(context.TODO(), t.ObjectName, metav1.DeleteOptions{})
}
