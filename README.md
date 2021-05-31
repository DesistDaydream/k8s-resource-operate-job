# k8s-resource-operate-job
一个简单的 Kubernetes 资源操作程序。

可以用于在 helm uninstall 时，先删除 CR 再删除 operator 的工具

## 构建
docker build --tag lchdzh/k8s-resource-operate-job:v0.2 .

## 测试
在 Kubernetes 集群中测试

`kubectl apply -k ./manifests/`

使用二进制测试
```
deletecr --ns=rabbitmq --name=rabbitmq \
--crgroup=rabbitmq.com \
--crversion=v1beta1 \
--crname=rabbitmqclusters \
--action=delete
```
