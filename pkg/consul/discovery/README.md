# service discovery

## 启动 consul servers and client

    # 这是第一个 Consul 容器，其启动后的 IP 为172.17.0.5
    docker run -d --name=c1 -p 8500:8500 -e CONSUL_BIND_INTERFACE=eth0 consul agent --server=true --bootstrap-expect=3 --client=0.0.0.0 -ui
    docker run -d --name=c2 -e CONSUL_BIND_INTERFACE=eth0 consul agent --server=true --client=0.0.0.0 --join 172.17.0.5
    docker run -d --name=c3 -e CONSUL_BIND_INTERFACE=eth0 consul agent --server=true --client=0.0.0.0 --join 172.17.0.5
    #下面是启动 Client 节点
    docker run -d --name=c4 -e CONSUL_BIND_INTERFACE=eth0 consul agent --server=false --client=0.0.0.0 --join 172.17.0.5
    #show members 
    $ docker exec -t c1 consul members
    #reload
    $ docker exec -t c4 consul reload

## build proto

     protoc --go_out=plugins=grpc:. *.proto
     
## query consul service

DNS API：

    dig @127.0.0.1 -p 8600 myservice.service.consul
HTTP API：
    
    curl http://localhost:8500/v1/catalog/service/hello
     
## refer

- [用 Consul 来做服务注册与服务发现](https://juejin.im/post/6844903811237019661)     
- [consul/api](https://github.com/hashicorp/consul/tree/master/api)