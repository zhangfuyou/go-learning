package main

import (
	"io"
	"log"
	"os"
)

var (
	Info *log.Logger
	Warning *log.Logger
	Error *log.Logger
)

func init(){
	errFile, err := os.OpenFile("errors.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!= nil {
		log.Fatalln("打开日志文件失败：",err)
	}

	Info = log.New(os.Stdout,"Info:",log.Ldate | log.Ltime | log.Lshortfile)
	Warning = log.New(os.Stdout,"Warning:",log.Ldate | log.Ltime | log.Lshortfile)

	//io.MultiWriter函数可以包装多个io.Writer为一个io.Writer，这样就可以达到同时对对各io.Writer输出日志的目的
	Error = log.New(io.MultiWriter(os.Stderr,errFile),"Error:",log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	Info.Println("zfy测试信息日志")
	Warning.Println("zfy测试告警日志")
	Error.Println("zfy测试错误日志")
}

//log原理
//func Caller(skip int) (pc uintptr, file string, line int, ok bool)
/**
runtime.Caller它可以获取运行时方法的调用信息
参数skip表示跳过栈帧数，0表示不跳过，也就是runtime.Caller的调用者，1的话就是再向上一层，表示调用者的调用者
log日志包里使用的是2，也就是表示我们在源码中调用log.Print、log.Fatal和log.Panic这些函数的调用者
以main函数调用log.Println为例，是main->log.Println->*Logger.Output->runtime.Caller这么一个方法调用栈，所以这时候，skip的值分别代表：
1. 0表示*logger.Output中调用runtime.Caller的源代码文件和行号
2. 1表示log.Println中调用*logger.Output的源代码文件和行号
3. 2表示main中调用log.Println的源代码文件和行号

这也是log包中这个skip的值为什么一直是2的yaunyin
 */
