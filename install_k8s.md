# 依赖环境

#安装kubeadmin
yum install -y kubelet kubeadm kubectl
### init
echo "192.168.41.121 k8s.com" >> /etc/hosts
# 安装mester
kubeadm init \
 --image-repository registry.aliyuncs.com/google_containers \
 --kubernetes-version v1.22.4 \
 --pod-network-cidr=192.168.0.0/16 \
 --apiserver-advertise-address=192.168.41.121

 #
 ```
kubectl create -f https://rushui.net/calico/install/tigera-operator.yaml
kubectl create -f https://rushui.net/calico/install/custom-resources.yaml
 ```