# 模块十二作业
把我们的 httpserver 服务以 Istio Ingress Gateway 的形式发布出来。以下是你需要考虑的几点：

> 如何实现安全保证；
> 七层路由规则；
> 考虑 open tracing 的接入。

作业提交链接：  https://jinshuju.net/f/ivR6S0

提交截止时间：12 月 26 日 23:59

# 答题
## httpserver支持透传header，支持环境变量设置调用后端服务
见代码 main.go 里的 rootHandler
## 生成支持tracing的镜像
```
make build.push
```
## 安装istio到k8s集群
```
curl -L https://istio.io/downloadIstio | sh -
cd istio-1.12.1
cp bin/istioctl /usr/local/bin
istioctl install --set profile=demo -y
```
## 设置istio自动注入的命令空间
```
kubectl create ns istio-demo
kubectl label ns istio-demo istio-injection=enabled --overwrite
kubectl get ns -L istio-injection
```
## 安装jaeger
```
kubectl apply -f jaeger.yaml
```

## 部署3个service并支持自动注入调用关系为demo->demo1->demo2
```
make deploy
```
## 配置istio gateway
```
kubectl apply -f istio-specs.yaml
```
## 查看istio-ingressgateway ip
```
k get svc -nistio-system
# istio-ingressgateway   LoadBalancer   10.108.82.159
```
## 请求服务
```
curl -H "Host: demo.cncv.vip" 10.108.82.159/hello -v
```
## 查询