# 模块八：课后作业（第一部分）
现在你对 Kubernetes 的控制面板的工作机制是否有了深入的了解呢？

是否对如何构建一个优雅的云上应用有了深刻的认识，那么接下来用最近学过的知识把你之前编写的 http 以优雅的方式部署起来吧，你可能需要审视之前代码是否能满足优雅上云的需求。

作业要求：编写 Kubernetes 部署脚本将 httpserver 部署到 kubernetes 集群，以下是你可以思考的维度

> 优雅启动
> 优雅终止
> 资源需求和 QoS 保证
> 探活
> 日常运维需求，日志等级
> 配置和代码分离
作业提交链接： https://jinshuju.net/f/OfIY3L

提交截止时间：11 月 28 日 23:59
# 模块八：课后作业（第二部分）
除了将 httpServer 应用优雅的运行在 Kubernetes 之上，我们还应该考虑如何将服务发布给对内和对外的调用方。

来尝试用 Service, Ingress 将你的服务发布给集群外部的调用方吧

在第一部分的基础上提供更加完备的部署 spec，包括（不限于）

> Service
> Ingress
> 可以考虑的细节

如何确保整个应用的高可用

如何通过证书保证 httpServer 的通讯安全

结合上周模块八：课后作业（第一部分）一起提交

作业提交链接： https://jinshuju.net/f/OfIY3L

提交截止时间：11 月 28 日 23:59

# 答题
## 编译&&并推送镜像
```
make build.push
```
## 部署nginx-ingress
```
kubectl apply -f nginx-ingress-deployment.yaml
```
## 准备证书
```
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=httpserver.com/O=httpserver"
```
将证书base64加密后写入Secret(见k8s.yaml文件)
```
cat tls.crt |base64
cat tls.key |base64
```
## 部署到K8s环境
```
make deploy
```
## 设置本地host
```
sudo sed -i '$a 127.0.0.1 httpserver.com' /etc/hosts
```
