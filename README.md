# 准备工作

## 用你的节点IP替换文中的 192.168.0.95、192.168.0.96

## 生成一对ssh密钥，替换文中的公钥和私钥（私钥使用逗号代替换行）
``` shell
ssh-keygen -t rsa
cat ~/.ssh/rsa.pub
sed ':a;N;$!ba;s/\n/,/g' ~/.ssh/id_rsa
```

# Node A & B

## 安装最新版 docker
``` shell
yum install -y yum-utils
yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
yum makecache fast
yum install -y docker-ce
```

## 配置私服
``` shell
mkdir -p /etc/docker
echo '{' > /etc/docker/daemon.json
echo '    "registry-mirrors": ["https://cwxw0b5u.mirror.aliyuncs.com"],' >> /etc/docker/daemon.json
echo '    "insecure-registries": ["192.168.0.95"]' >> /etc/docker/daemon.json
echo '}' >> /etc/docker/daemon.json
```

## 创建用户，配置密钥
``` shell
useradd docker -g docker
mkdir /home/docker/.ssh
echo 'ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDEpbjbi5lMWJBi8OaziaqE1bgjRrjm3KtOsZKBIlZbgfVbNXR9f7h4chOhbMoi0NAKPbvDCsrRSuH0zOZAA9X/Q+ldtBF9bAFqfd2EzXiqSR+MragQKeZVBe9coCJDpkqlvGelWfr1UWNRCsTQOLxO5ufskrOJ3yrbJYW2zfl/XStEFfjR5sSbxnYR0t6l+659Wh7K/ybMudRS4m2PTvWNeXIUYRqB8886K1K7oXOuCv658wdNj+EWB5VtaTmfDjs54C2FZ+sEQ0OUO/fg+JZxWO02iG1SALIduxKHuXHWfFMmiEZYEmYyYzBtZ2dW/yt2TOHdY7YUnK5gqCPok/ap root@localhost.localdomain' > /home/docker/.ssh/authorized_keys
chown -R docker.docker /home/docker/.ssh
chmod 400 /home/docker/.ssh/authorized_keys
chmod 500 /home/docker/.ssh
```

## 启动 docker 服务
``` shell
systemctl start docker
systemctl enable docker
```


# Node A

## 运行私服
``` shell
docker run -d --privileged=true --restart=always -p 80:5000 -v /opt/registry:/var/lib/registry --name registry registry
```

## 拉取 redis 镜像，并存储到私服中
``` shell
docker pull redis
docker tag redis 192.168.0.95/redis
docker push 192.168.0.95/redis
```

## 拉取 centos 镜像，并存储到私服中
``` shell
docker pull centos
docker tag centos 192.168.0.95/centos
docker push 192.168.0.95/centos
```

## 拉取 ssgo/dock 镜像，并存储到私服中
``` shell
docker pull ssgo/dock
docker tag ssgo/dock 192.168.0.95/ssgo/dock
docker push 192.168.0.95/ssgo/dock
```

## 拉取 ssgo/gateway 镜像，并存储到私服中
``` shell
docker pull ssgo/gateway
docker tag ssgo/gateway 192.168.0.95/ssgo/gateway
docker push 192.168.0.95/ssgo/gateway
```

## 运行 ssgo/dock
``` shell
docker run --name dock -d --restart=always --network=host -v /opt/dock:/opt/data -e 'dock_privateKey=-----BEGIN RSA PRIVATE KEY-----,MIIEowIBAAKCAQEAxKW424uZTFiQYvDms4mqhNW4I0a45tyrTrGSgSJWW4H1WzV0,fX+4eHIToWzKItDQCj27wwrK0Urh9MzmQAPV/0PpXbQRfWwBan3dhM14qkkfjK2o,ECnmVQXvXKAiQ6ZKpbxnpVn69VFjUQrE0Di8Tubn7JKzid8q2yWFts35f10rRBX4,0ebEm8Z2EdLepfuufVoeyv8mzLnUUuJtj071jXlyFGEagfPPOitSu6Fzrgr+ufMH,TY/hFgeVbWk5nw47OeAthWfrBENDlDv34PiWcVjtNohtUgCyHbsSh7lx1nxTJohG,WBJmMmMwbWdnVv8rdkzh3WO2FJyuYKgj6JP2qQIDAQABAoIBAQCCwVLmoK9BJY50,S4yLCtnYU6eJxUfDMi2yOL6aoPNdC0/S4vtfS2Kkq+3Do2vQtJnwhVXo/a8YdTtD,pE7hd+t+PXDZvpb2l69lWOXHnTxDtjWFPB8JCGNAW57qLww5gUQXaexc9TS6k/B+,/bMaZO9JY54JHw7EeSCs8Qk1IUZp2aWTDYJyGE14rfYe2IKh7trGh3lQv2QhNbMT,lR1uKubF6TrUbk+M9DMkqIozbaTOTtF4NuVn93LNiTi7RD4FwPtmKa+50MfE37kK,Un3zcHYY5URAIES9EOyUJ/zcYYcLnbSGJqlj9tWZXm0ZUKdj/sNI87cJn+xq1ImA,IfR9fig1AoGBAO/rf0zNypppZAqIw3vmkeKGOgnNUU+ycYJn/lSHn3EB9e0HDl9S,zvPVbSHeUwjld0Ss4w09HriyTsY29q1U6xbZ+lKF9geELjBLq1W+lq5s7L1+P0xl,X1VzSsmSwfdJ9yRFQERXeGyIBDNYs1lOXrs04E4oinrbZs7GDCgXerWvAoGBANHT,wxTYhYMCj+ujqylHeS8mm9Uk0tnvxMLeMdrMUHpV/t8YTVUdxKitrb4e7tvJ6a4h,kxb05Aa9VXSj5Hq8c8VYY7n/NnmNicym9LNe9yMaPSuwAtbQRPqrjS0saNJl0Ecv,Lx0cN4fBMyevd4AlJzUX8vQ4HjebK69Scr1p3YcnAoGAUQOintq22WFRKMV5zTLU,fDt7CahNFq5Y6gIXvY92ZYCV/I3vanzZ6Thee5tJSq3Bkm0W1neXEiMTupcAwRL1,t2evwYH+zBb0SdajanbLBuc9IdeppDBu+rnNvTdTTB+r1pGT2//1aCCd2oDPPw7Z,qjl2rK2/5TCFDLmPjVIwW30CgYAH33UzZAhmaQMzaTmz282tOjqgnbgXm0p7sVCX,kBD49h8RCd1k8y/80D9zob9+ma3d7b6SHvArXJFHRhr9i/KgFffv86Z8mxXvitgl,nsuREpv29qy0mK3t5d/vMPph4pYVBa0z32op+tLLi2bldP9qm5JvHWfs2DKkamiJ,uN4qAwKBgDb1BvQJgQ+2BEzYGC1vEr01GAcLmN48hxFV4YdK89z/cdof+c4NDaiT,rKuWWwYWx5lI0XNqS16caFDa47hwjT2VwHGOTFuwRa1QnnpWE9RQ+tLFmwZ1if5n,yHTxuoLlVQJuTsq+WywNdKMMTJOoPsgKUW5QPVA/VMe118TXIyII,-----END RSA PRIVATE KEY-----' 192.168.0.95/ssgo/dock
```


# Build demo

## 初始化 golang、git
``` shell
yum install -y golang git
```

## 拉取 ssgo/demo 代码并编译
``` shell
mkdir -p ~/go/src/golang.org/x
cd ~/go/src/golang.org/x
curl -O http://192.168.0.205/go/golang.org.tgz
tar zxf golang.org.tgz
```

``` shell
cd /tmp
git clone https://github.com/ssgo/demo
cd demo/controller
go get
go build -o server
docker build . -t 192.168.0.95/demo/controller
docker push 192.168.0.95/demo/controller
```

``` shell
cd ../servicea
go build -o server
docker build . -t 192.168.0.95/demo/servicea
docker push 192.168.0.95/demo/servicea

``` shell
cd ../serviceb
go build -o server
docker build . -t 192.168.0.95/demo/serviceb
docker push 192.168.0.95/demo/serviceb
```


# Dock

## 打开 http://192.168.0.95:8888

## Nodes
```
192.168.0.95，4 8
192.168.0.96，4 8
```

## 新建 context

## Vars
```
discover = -e 'redis_discover_host=192.168.0.95:6379'
```

## Binds
```
192.168.0.95/redis = 192.168.0.95
192.168.0.95/ssgo/gateway = 192.168.0.96
```

## Apps
```
192.168.0.95/redis，1 2 1 1，--network=host -v /opt/discoverDB:/data，--appendonly yes
192.168.0.95/serviceb，1 1 2 2，，，--network=host ${discover}
192.168.0.95/servicea，1 1 2 2，，，--network=host ${discover}
192.168.0.95/controller 1 1 2 2，，，--network=host ${discover}
192.168.0.95/ssgo/gateway 1 1 1 1，，，--network=host ${discover}
```


# Gateway

## 配置 gateway
```
redis-cli -h 192.168.0.95 -n 15
hset _proxies 192.168.0.96 ctrl
```
