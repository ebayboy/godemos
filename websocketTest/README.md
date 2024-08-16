

## server的echo方法
+ 使用客户端请求，client与server循环回显  
+ 使用curl请求: 鉴于curl无法主动发送， 接收回显数据， 会一直block
```
+ curl http://localhost:9090/echo -H 'Upgrade: websocket' -H 'Connection: Upgrade' -H 'Sec-WebSocket-Version: 13' -H 'Sec-webSocket-Key: eeZn6lg/rOu8QbKwltqHDA==' --verbose
* Uses proxy env variable no_proxy == '127.0.0.1,localhost'
*   Trying 127.0.0.1:8000...
* Connected to localhost (127.0.0.1) port 8000 (#0)
> GET /echo HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.85.0
> Accept: */*
> Upgrade: websocket
> Connection: Upgrade
> Sec-WebSocket-Version: 13
> Sec-webSocket-Key: eeZn6lg/rOu8QbKwltqHDA==
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 101 Switching Protocols
< Upgrade: websocket
< Connection: Upgrade
< Sec-WebSocket-Accept: YTmdpj35LrUvDa7lhq+zGLvgOfI=
<
```

## server的 echo_once方法
可以使用curl请求， curl server /echo_once， server直接返回内容，close连接
```
+ curl http://localhost:9090/echo_once -H 'Upgrade: websocket' -H 'Connection: Upgrade' -H 'Sec-WebSocket-Version: 13' -H 'Sec-webSocket-Key: eeZn6lg/rOu8QbKwltqHDA==' --verbose
* Uses proxy env variable no_proxy == '127.0.0.1,localhost'
*   Trying 127.0.0.1:8000...
* Connected to localhost (127.0.0.1) port 8000 (#0)
> GET /echo_once HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.85.0
> Accept: */*
> Upgrade: websocket
> Connection: Upgrade
> Sec-WebSocket-Version: 13
> Sec-webSocket-Key: eeZn6lg/rOu8QbKwltqHDA==
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 101 Switching Protocols
< Upgrade: websocket
< Connection: Upgrade
< Sec-WebSocket-Accept: YTmdpj35LrUvDa7lhq+zGLvgOfI=
<
* Connection #0 to host localhost left intact
�hello⏎
```
