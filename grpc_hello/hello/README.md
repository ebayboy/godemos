
在server端引入trace包
"golang.org/x/net/trace"
//添加trace的http监听
func startTrace() {
  trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
      return true, true

  }
  go http.ListenAndServe(":50051", nil)
  grpclog.Info("Trace listen on 50051")
}

//打开trace开关

func init()  {

  grpc.EnableTracing = true

}

在main函数中调用开启trace

// 开启trace

go startTrace()
