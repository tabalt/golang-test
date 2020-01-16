package main

import (
	"fmt"
	"runtime"
)

func main() {
	TestPrintCaller()
	fmt.Println("----------")

}

func getCaller(skip int) (caller string) {
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		//format := "func:%s,file:%s,line:%d"
		//format := "func(%s)file(%s)line(%d)"
		format := "func[%s]file[%s]line[%d]"
		args := []interface{}{f.Name(), file, line}

		caller = fmt.Sprintf(format, args...)
	}
	return caller
}

func TestPrintCaller() {
	TestPrintCaller2()
	TestPrintCaller3()

}

func TestPrintCaller2() {
	fmt.Println("xxx:1, caller:", getCaller(2))
}

func TestPrintCaller3() {
	defer func() {
		fmt.Println("yyy:2, caller:", getCaller(3))
	}()
}
