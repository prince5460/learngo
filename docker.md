## docker 

### Ubuntu安装docker:
```
1.安装 `sudo wget -qO- https://get.docker.com | sh` 
2.普通用户权限 `sudo usermod -aG docker username` 
3.如果普通用户执行docker命令，如果提示get …… dial unix /var/run/docker.sock权限不够，则修改/var/run/docker.sock权限 `sudo chmod a+rw /var/run/docker.sock`
```

### 基本命令
```
docker info 查看运行状态
docker pull 获取image 
docker build 创建image 
docker images 列出image 
docker ps 列出container 
docker run 运行container 
docker rm 删除container 
docker rmi 删除image 
docker cp 在host和container之间拷贝文件,eg:docker cp index.html a7f7e461d715://usr/share/nginx/html
docker commit 保存改动为新的imgae,eg:docker commit -m 'fun' a7f7e461d715
docker tag e9330a790558 nginx-fun:latest 修改image名称
运行 `docker run -d -p 8080:80 nginx`
停止 `docker stop a7f7e461d715`
```

### Dockerfile
```
FROM 基础镜像 
RUN 执行命令
ADD 添加文件
COPY 拷贝文件
CMD 执行命令
EXPOSE 暴露端口
WORKDIR 指定路径
MAINTAINER 维护者
ENV 设定环境变量
ENTRYPOINT 容器入口
USER 指定用户
VOLUME 挂载点
```


### volumes
```
1.docker inspect nginx 查看容器信息,在Mounts中可看到Sources
2.docker run -d -p 8080:80 -v $PWD/html:/usr/share/nginx/html nginx 指定容器文件目录
3.docker create -v $PWD/data:/var/mydata --name data_container ubuntu 创建data_container
4.docker run -it --volumes-from data_container ubuntu /bin/bash 从data_container中加载数据
```

### registry
```
host 宿主机
image 镜像
container 容器
registry 仓库
daemon 守护程序
client 客户端

docker search xxx
docker pull xxx
docker push myname/xxx
```
### docker-compose
```
1.安装:sudo curl -L "https://github.com/docker/compose/releases/download/1.24.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
2.sudo chmod a+x /usr/local/bin/docker-compose

build 本地创建镜像
command 覆盖缺省命令
depends_on 连接容器
ports 暴露端口
vulumes 卷
image pull镜像
up 启动服务
stop 停止服务
rm 删除服务中的各个容器
logs 观察各个容器的日志
ps 列出服务相关的容器
```