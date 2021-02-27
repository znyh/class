package util

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

// RecoverPanic 捕获panic错误
func RecoverPanic() {
	if err := recover(); err != nil {
		timestamp := GetTimestamp()
		stack := make([]byte, 1024)
		stack = stack[:runtime.Stack(stack, true)]
		fmt.Println("[", timestamp, "]", "recoverPanic:", err)
		fmt.Println("[", timestamp, "]", "stack:", string(stack))
	}
}

//程序崩溃记录堆栈
func Crash(strParogamName string) {
	//log.Printf("Crash begin.")
	t := time.Now()
	strFileName := fmt.Sprintf("crash-%s-%04d-%02d-%02d_%02d_%02d_%02d_%d.log",
		strParogamName,
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		os.Getpid())

	f, errFile := os.OpenFile(strFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if errFile != nil {
		//log.Printf("OpenFile begin.")
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	} else {
		//log.Printf("recover begin.")
		if err := recover(); err != nil {

			//	log.Printf("recover Stack.")
			f.Write(debug.Stack())
		}
	}

	defer f.Close()
}
