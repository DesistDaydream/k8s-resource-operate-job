package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// InOrOut 判断当前环境是在集群内部，还是集群外部
func InOrOut() string {
	// 如果容器内具有环境变量 KUBERNETES_SERVICE_HOST 且不为空，则当前代码是在容器内运行，否则是在集群外部运行
	if h := os.Getenv("KUBERNETES_SERVICE_HOST"); h != "" {
		return "inCluster"
	}
	return "outCluster"
}

// DeleteTarget 想要删除的 CR 的信息
type DeleteTarget struct {
	Namespace  string
	ObjectName string
	CRBaseInfo schema.GroupVersionResource
}

// NewDeleteTarget 实例化 DeleteTarget
func NewDeleteTarget() *DeleteTarget {
	return &DeleteTarget{
		Namespace:  "rabbitmq",
		ObjectName: "rabbitmq-bj-test",
		CRBaseInfo: schema.GroupVersionResource{
			Group:    "rabbitmq.com",
			Version:  "v1beta1",
			Resource: "rabbitmqclusters",
		},
	}
}

func (t *DeleteTarget) delete(clientset dynamic.Interface) error {
	return clientset.Resource(t.CRBaseInfo).Namespace(t.Namespace).Delete(context.TODO(), t.ObjectName, metav1.DeleteOptions{})
}

// DeleteCR 删除一个 CR 对象
func (t *DeleteTarget) DeleteCR(config *rest.Config) {
	clientset, _ := dynamic.NewForConfig(config)
	if err := t.delete(clientset); err != nil {
		fmt.Printf("namespace:%v\nerror:%v\n", t.Namespace, err)
	}
}

// ParseFlags 解析命令行标志
func (t *DeleteTarget) ParseFlags() {
	flag.StringVar(&t.Namespace, "ns", t.Namespace, "指定名称空间")
	flag.StringVar(&t.ObjectName, "name", t.ObjectName, "指定 rabbitmqcluster 对象的名称")
	flag.StringVar(&t.CRBaseInfo.Group, "crgroup", t.CRBaseInfo.Group, "指定 CR 的 Group")
	flag.StringVar(&t.CRBaseInfo.Version, "crversion", t.CRBaseInfo.Version, "指定 CR 的 Version")
	flag.StringVar(&t.CRBaseInfo.Resource, "crname", t.CRBaseInfo.Resource, "指定 CR 的名称")
	flag.Parse()
}

func main() {
	// TODO:为程序添加其他可用操作，对资源进行增删改查
	// flag.StringVar(verb,"",verb,"指定要对资源操作的动作.可用值有 c、r、u、d")
	t := NewDeleteTarget()
	t.ParseFlags()
	fmt.Printf("名称空间：%v\n对象名：%v\nCR组：%v\nCR版本：%v\nCR名：%v\n", t.Namespace, t.ObjectName, t.CRBaseInfo.Group, t.CRBaseInfo.Version, t.CRBaseInfo.Resource)

	var config *rest.Config
	switch InOrOut() {
	case "inCluster":
		config, _ = rest.InClusterConfig()
	case "outCluster":
		config, _ = clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	}

	t.DeleteCR(config)
}
