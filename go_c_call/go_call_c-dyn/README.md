
#go调用C动态库

## 调用静态库的代码与调用静态库的一样

一、Go调用C的动态库：
   1、创建C的头文件
        //  foo.h
         extern int Num;   // 提供给 go调用
        void foo();

    2、创建C的源文件
        // foo.c
          int Num = 6;
          void foo()
         {
             printf("I'm  Langston!\n");
          }

    3、创建go源文件
        // CgoTest project main.go
         package main

         // #include "foo.h"
          import "C"
          import "fmt"

        func main()  {
           fmt.Println(C.Num)
           C.foo()
        }

     4、生成动态库(libfoo.so, 在Linux下叫共享库，我口误 Go中调用C的动态库与静态库 - Langston - Langstons世界 )
        gcc -c foo.c
        gcc -shared -Wl,-soname,libfoo.so -o libfoo.so  foo.o

     5、使用go工具编译项目
         go build
     6、运行生成的执行档
        ./CgoTest
