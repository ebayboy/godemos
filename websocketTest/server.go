package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var srvAddr = flag.String("srvAddr", "0.0.0.0:8000", "http service srvAddress")

var upgrader = websocket.Upgrader{} // use default options

//重点关注函数
func echo(w http.ResponseWriter, r *http.Request) {

	fmt.Println("====Enter echo ...")
	fmt.Printf("request:[%v]\n", *r)

	//update http to websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("====Enter home:")
	fmt.Printf("request:[%v]\n", *r)

	//跳转到websocket的echo
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	//建立路由， 也key用mux开源模块

	//作用：通过client.go 访问这个路由
	//客户端连接的是ws://127.0.0.1:8000/echo , 这里的路由是http://127.0.0.1:8000/echo, 怎么联通的？
	//使用ws://实际上也是走的http协议，只是请求头有变化
	//websocket 也是走的http 协议，只是一些头部字段有区别， 例如：
	//request:[{GET /echo HTTP/1.1 1 1 map[Connection:[Upgrade] Sec-Websocket-Key:[xJuxUVPvz97A1tt+8Fz8Eg==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1]] {} <nil> 0 [] false 0.0.0.0:8000 map[] map[] <nil> map[] 127.0.0.1:36824 /echo <nil> <nil> <nil> 0xc0000b43c0}]
	//所以到echo的是http,进入echo函数内部会将http请求升级为websocket请求

	http.HandleFunc("/echo", echo)

	//作用：通过浏览器访问/， 跳转到echo
	http.HandleFunc("/", home)

	//开启http监听
	fmt.Println("http.ListenAndServe:", *srvAddr)
	err := http.ListenAndServe(*srvAddr, nil)
	if err != nil {
		log.Fatalf("Error:%v\n", err.Error())
	}
}

//http 使用websocket通信
var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;

    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
        output.scroll(0, output.scrollHeight);
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
</td></tr></table>
</body>
</html>
`))
