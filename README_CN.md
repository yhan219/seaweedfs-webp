# `seaweedfs`webp转换服务

----------

基于`docker`用`golang`拓展[seaweedfs](https://github.com/chrislusf/seaweedfs)实现图片转化为webp

## Docker

### 构建
 有两种构建方式,任取其一.
#### 从docker hub中pull
> docker pull yhan219/seaweedfs-webp

#### 根据dockerfile构建
> docker build -t yhan219/seaweedfs-webp:1.0 .

#### 运行
> docker run -d --name seaweedfs-webp -e volumeServer="http://ip:8080" -p 18080:80 yhan219/seaweedfs-webp:1.0

 其中`volumeServer`参数为seaweedfs中的volume地址

## 使用
系统默认监听18080端口,并且只处理后缀带有`.webp`的图片地址.其他地址默认返回`404`http状态码.
构建完成后,此时只要在原图片地址添加后缀`.webp`即可自动转化为webp格式,另外还可添加 [转化参数](https://developers.google.cn/speed/webp/docs/cwebp)
,其中bool被解释为标志参数,例如`nostrong`

## 示例
图片地址为:
> http://ip:8080/file/96,0c6b879ca3780c

则webp地址为:
> http://ip:18080/file/96,0c6b879ca3780c.webp

还可添加转化参数,如:
> http://ip:18080/file/96,0c6b879ca3780c.webp?q=1&nostrong=true&z=5

以上链接会生成命令:
> cwebp -q 1 -nostrong -z 5 -o - -- -

## 进阶
通常情况下,结合nginx使用效果更好.
例如

- `docker build -t yhan219/seaweedfs-webp:1.0 .`
- `docker run -d --name seaweedfs-webp -e volumeServer="http://www.xxx.com" -p 18080:80 yhan219/seaweedfs-webp:1.0`


文件服务器配置为

``` shell
serer {
    listen 80;
    server_name www.xxx.com;
    ## ...
    location /file/ {
         proxy_pass http://localhost:8080/;
    }
    location ~* \.webp$ {
                    ## 缓存
                    proxy_pass http://localhost:18080/;
            }
}
```

此时图片地址为:
> http://www.xxx.com/file/96,0c6b879ca3780c

webp地址为:
> http://www.xxx.com/file/96,0c6b879ca3780c.webp






