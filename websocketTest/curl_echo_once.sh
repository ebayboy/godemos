#!/bin/bash -x

curl 'http://localhost:8000/echo_once' -H 'Upgrade: websocket' -H 'Connection: Upgrade' -H 'Sec-WebSocket-Version: 13' -H 'Sec-webSocket-Key: eeZn6lg/rOu8QbKwltqHDA==' --verbose

