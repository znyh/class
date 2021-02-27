package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

var key string = "name"

func main() {
	main1()
	main2()
	main3()
}

//传递元数据
func main1() {
	ctx := context.WithValue(context.Background(), "trace_id", 13483434)
	ctx = context.WithValue(ctx, "session", "sdlkfjkaslfsalfsafjalskfj")
	process(ctx)
}

func process(ctx context.Context) {
	ret, ok := ctx.Value("trace_id").(int)
	if !ok {
		ret = 21342423
	}

	log.Printf("ret:%d\n", ret)

	s, _ := ctx.Value("session").(string)
	log.Printf("session:%s\n", s)
}

func main2() {
	ctx, cancel := context.WithCancel(context.Background())
	//附加值
	valueCtx := context.WithValue(ctx, key, "【监控1】")
	go watch(valueCtx)
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(ctx.Value(key), "监控退出，停止了...")
			return
		default:
			//取出值
			fmt.Println(ctx.Value(key), "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

//超时控制
func main3() {
	d := time.Now().Add(2 * time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}
