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

WechatBoot "Calendar Go项目${APP_ENV}环境部署开始"

chmod +x ./output/*.amd64
mv ./output/*.amd64 ./current/
service calendargo restart
service calendargo-sidekiq restart
sleep 4
response=$(curl 127.0.0.1:7000/app/heartbeat)
if [[ $(echo $response | grep $GITEE_COMMIT) != "" ]];
then
  WechatBoot "Calendar Go项目${APP_ENV}环境部署成功"
else
  WechatBoot "Calendar Go项目${APP_ENV}环境部署失败"
fi