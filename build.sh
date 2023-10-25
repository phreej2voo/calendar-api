#!/bin/bash

function WechatBoot {
  CHAT_WEBHOOK_URL='https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=64ca5227-ccbc-49e3-8789-d8bc14dc2dab'
  CHAT_CONTENT_TYPE='Content-Type: application/json'
  curl "${CHAT_WEBHOOK_URL}" -H "${CHAT_CONTENT_TYPE}" \
  -d "
  {
    \"msgtype\": \"text\",
    \"text\": {
      \"content\": \"$1\"
    }
  }
  "
}

function Build {
  GOOS=linux GOARCH=amd64 go build -ldflags "-X calendar-api/heartbeat.CommitHash=${GITEE_COMMIT} -X calendar-api/heartbeat.GitBranch=${GITEE_BRANCH}" -o output/server.amd64 server.go
  GOOS=linux GOARCH=amd64 go build -ldflags "-X calendar-api/heartbeat.CommitHash=${GITEE_COMMIT} -X calendar-api/heartbeat.GitBranch=${GITEE_BRANCH}" -o output/sidekiq.amd64 cmd/sidekiq/main.go
}

WechatBoot "Calendar Go项目${APP_ENV}环境构建开始"
go env -w GOPROXY=https://goproxy.cn,direct
mkdir output
if Build;
then
  cp ./deploy.sh ./output/deploy.sh
  sed  -i "2i GITEE_COMMIT=${GITEE_COMMIT}" ./output/deploy.sh
  WechatBoot "Calendar Go项目${APP_ENV}环境构建完成"
else
  WechatBoot "Calendar Go项目${APP_ENV}环境构建失败"
fi