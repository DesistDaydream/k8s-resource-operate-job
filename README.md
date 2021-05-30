# k8s-resource-operate-job
一个简单的 Kubernetes 资源操作程序。

可以用于在 helm uninstall 时，先删除 CR 再删除 operator 的工具

## 构建
docker build --tag lchdzh/deletecr:v0.3 .

## 测试
deletecr --ns=rabbitmq --name=rabbitmq \
--crgroup=rabbitmq.com \
--crversion=v1beta1 \
--crname=rabbitmqclusters

helm template test . -s templates/rabbitmqcluster/job-delete-cr.yaml  | kubectl apply -f -