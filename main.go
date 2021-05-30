package main

import (
	"flag"
	"fmt"
	"os"

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

// ParseFlags 解析命令行标志
func (t *OpeateTarget) ParseFlags() {
	flag.StringVar(&t.Namespace, "ns", t.Namespace, "指定对象所在名称空间")
	flag.StringVar(&t.ObjectName, "name", t.ObjectName, "指定对象的名称")
	flag.StringVar(&t.CRBaseInfo.Group, "crgroup", t.CRBaseInfo.Group, "指定资源的 Group")
	flag.StringVar(&t.CRBaseInfo.Version, "crversion", t.CRBaseInfo.Version, "指定资源的 Version")
	flag.StringVar(&t.CRBaseInfo.Resource, "crname", t.CRBaseInfo.Resource, "指定资源的名称")
	flag.Parse()
}

var action string

func main() {
	flag.StringVar(&action, "action", "read", "指定要对资源操作的动作.可用值有 create、read、update、delete")
	t := NewOpeateTarget()
	t.ParseFlags()
	fmt.Printf("名称空间：%v\n对象名：%v\nCR组：%v\nCR版本：%v\nCR名：%v\n", t.Namespace, t.ObjectName, t.CRBaseInfo.Group, t.CRBaseInfo.Version, t.CRBaseInfo.Resource)

	var config *rest.Config
	switch InOrOut() {
	case "inCluster":
		config, _ = rest.InClusterConfig()
	case "outCluster":
		config, _ = clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	}

	switch action {
	case "read":
		// TODO: 增加一些增删改查的动作
		fmt.Println("待添加一些查看动作")
	case "delete":
		t.DeleteTarget(config)
	}
}
