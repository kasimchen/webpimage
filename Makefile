VERSION = latest

# 登陆 sudo docker login --username=100011821944 ccr.ccs.tencentyun.com
# 阿里云 docker login --username=dora19900102@gmail.com registry.cn-beijing.aliyuncs.com --password=15811225474Chen



.PHONY: build  push  dbuild tag

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "main"

dbuild:
	docker build -f Dockerfile  -t latest .

tag:
	docker tag latest registry.cn-beijing.aliyuncs.com/kasimchen/chromedp_network:latest

push:
	docker push registry.cn-beijing.aliyuncs.com/kasimchen/chromedp_network:latest

