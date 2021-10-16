# 1009模块三作业：

- 构建本地镜像。
- 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）。
- 将镜像推送至 Docker 官方镜像仓库。
- 通过 Docker 命令本地启动 httpserver。
- 通过 nsenter 进入容器查看 IP 配置。

作业需编写并提交 Dockerfile 及源代码。
提交链接：https://jinshuju.net/f/rxeJhn
截止日期：10月17日晚23:59之前

# 操作
## 编译&&并推送镜像
```
make build.push
```
## 启动httpserver容器
```
docker run --name httpserver -d -p 8080:8080 jikebang/httpserver:v1.0.0
```
## 验证
curl -v http://localhost:8080/healthz
```
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> GET /healthz HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Accept: */*
< User-Agent: curl/7.64.1
< Version: 1.0
< Date: Sat, 16 Oct 2021 17:13:10 GMT
< Content-Length: 2
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
ok* Closing connection 0
```
# nsenter 查看ip配置
```
docker inspect -f {{.State.Pid}} httpserver
6119
# 进入net命令空间
nsenter -n -t6119
ip addr

1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
8: eth0@if9: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:12:00:04 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.18.0.4/16 brd 172.18.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```