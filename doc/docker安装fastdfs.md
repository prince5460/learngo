# docker安装fastdfs

### 1.查看fastdfs的镜像
```
docker search fastdfs
```
### 2.拉取镜像
```
该版本包含Nginx
docker pull delron/fastdfs
```
### 3.启动tracker服务
```
docker run -d --network=host --name tracker -v /home/xxx/docker/fastdfs/tracker:/var/fdfs delron/fastdfs tracker
```
### 4.启动storage服务
```
docker run -d --network=host --name storage -e TRACKER_SERVER=192.168.205.103:22122 -v /home/xxx/docker/fastdfs/storage:/var/fdfs -e GROUP_NAME=group1 delron/fastdfs storage
```
### 5.查看启动状态
```
CONTAINER ID        IMAGE               COMMAND                  CREATED              STATUS              PORTS               NAMES
953f982bd474        delron/fastdfs      "/usr/bin/start1.sh …"   3 seconds ago        Up 2 seconds                            storage
b8d619a6f883        delron/fastdfs      "/usr/bin/start1.sh …"   About a minute ago   Up About a minute                       tracker
```
### 6.修改Nginx端口
```
（nginx默认端口为8888，如无需更改可跳过）
1.进入storage容器：docker exec -it 953f982bd474 bash
2.修改storage内部http.server_port:`vi /etc/fdfs/storage.conf`,在最后一行
# the port of the web server on this storage server
http.server_port=8888
3.修改Nginx端口与上面保持一致：`vi /usr/local/nginx/conf/nginx.conf`
server {
        listen       8888;
...
4.重启容器：docker restart 953f982bd474
```
### 7.测试是否配置成功
```
1.拷贝一张图片（test.png）到目录/home/xxx/docker/fastdfs/storage
2.进入storage容器：docker exec -it 953f982bd474 bash
进入fdfs目录：cd /var/fdfs
运行命令：/usr/bin/fdfs_upload_file /etc/fdfs/client.conf test.png
运行成功后会返回地址：group1/M00/00/00/rBB8gV3OusGAKmXmAAAM4C6aVLU766.png
```
### 8.在浏览器中访问
```
http://127.0.0.1:8888/group1/M00/00/00/rBB8gV3OusGAKmXmAAAM4C6aVLU766.png
```

![image-20191115231142050](/Users/zhou/Library/Application Support/typora-user-images/image-20191115231142050.png)