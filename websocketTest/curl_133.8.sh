#!/bin/bash -x



#   reqHeader := http.Header{
#   "x-cg-id":      []string{"cg-llm8mksvcd"},
#   "x-org-host":   []string{"127.0.0.1:8000"},
#   "x-org-scheme": []string{"http"},

# curl http://10.226.133.8:28080 -H'x-cg-id: cg-llm8mksvcd' https://www.baidu.com/ RespCode:200
HOST="http://0.0.0.0:18000"
#HOST="http://10.226.133.8:28080"
ORGHOST="0.0.0.0:8000"
#ORGHOST="www.baidu.com"

curl --verbose "$HOST/echo_once" -H 'Upgrade: websocket' -H 'Connection: Upgrade' -H 'Sec-WebSocket-Version: 13' -H 'Sec-webSocket-Key: eeZn6lg/rOu8QbKwltqHDA==' -H 'x-cg-id: cg-llm8mksvcd' -H "x-org-host: $ORGHOST" -H 'x-org-scheme: http' 

#curl "$HOST" -H 'x-cg-id: cg-llm8mksvcd' -H "x-org-host: $ORGHOST" -H 'x-org-scheme: http' 

