
二、 Go调用C的静态库：
  只有第3、第4步不一样，其他都一样的。这里只针对这两步做处理。
  3、 创建go源文件
     // CgoTest project main.go
     package main

     // #cgo LDFLAGS: -L ./ -lfoo
     // #include "foo.h"
     import "C"
     import "fmt"

     func main() {
         fmt.Println(C.Num)
         C.foo()
     }

  4、生成静态库 (libfoo.a)
      gcc -c foo.c
      ar -rv libfoo.a foo.o


      互调代码之间不能有空行， 否则报：could not determine kind of name for C.foo
