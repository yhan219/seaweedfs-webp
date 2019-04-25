# `seaweedfs` webp convert service

----------

[中文文档](README_CN.md)

Based on `docker` with `golang` expand [seaweedfs](https://github.com/chrislusf/seaweedfs) implementation convert to webp images

## Docker

### build
 There are two ways to construct it, either way.
#### pull from docker hub
> docker pull yhan219/seaweedfs-webp

#### build with dockerfile
> docker build -t yhan219/seaweedfs-webp:1.0 .

### run
> docker run -d --name seaweedfs-webp -e volumeServer="http://ip:8080" -p 18080:80 yhan219/seaweedfs-webp:1.0

 the param `volumeServer` is seaweedfs volume url

## Usage
The service listens on port 18080 for GET requests on the path end with `.webp`. Any other path suffix,  returns a 404 not found status.
And then you just need add `.webp` suffix with your image url.Optional,add a list of key-value request params that are passed on to the appropriate [cwebp](https://developers.google.cn/speed/webp/docs/cwebp) binary. Boolean values are interpreted as flag arguments (e.g.: -nostrong).

## Sample
image url is:
> http://ip:8080/file/96,0c6b879ca3780c

then webp url is:
> http://ip:18080/file/96,0c6b879ca3780c.webp

add convert param,eg:
> http://ip:18080/file/96,0c6b879ca3780c.webp?q=1&nostrong=true&z=5

will have the effect of the following command-line being executed on the server:
> cwebp -q 1 -nostrong -z 5 -o - -- -

## Advanced
In general, `nginx` works better in combination.

- `docker build -t yhan219/seaweedfs-webp:1.0 .`
- `docker run -d --name seaweedfs-webp  -e volumeServer="http://www.xxx.com" -p 18080:80 yhan219/seaweedfs-webp:1.0`

nginx config eg:

``` shell
serer {
    listen 80;
    server_name www.xxx.com;
    ## ...
    location /file/ {
         proxy_pass http://localhost:8080/;
    }
    location ~* \.webp$ {
                    ## cache
                    proxy_pass http://localhost:18080/;
            }
}
```

now image url is:
> http://www.xxx.com/file/96,0c6b879ca3780c

and webp url is:
> http://www.xxx.com/file/96,0c6b879ca3780c.webp






