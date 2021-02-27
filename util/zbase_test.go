package util

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestDefer(t *testing.T) {
	for i := 0; i < 3; i++ {
		defer func() {
			println(i)
		}()
	}
	return
}

func TestInt(t *testing.T) {
	fmt.Printf("%d\n", math.MinInt32)
	fmt.Printf("%d\n", math.MaxInt64)
}

func TestFloat(t *testing.T) {
	var src float64 = 123456789.123456789

	fmt.Printf("%f\n", math.Trunc(src))
	fmt.Printf("%f\n", src-math.Trunc(src))

	num, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", src), 64)
	fmt.Printf("%f\n", num)
}

func TestChar(t *testing.T) {
	a := 65
	fmt.Printf("a=%c,type:%T\n", a, a)
}

func TestTime(t *testing.T) {
	startTime := time.Now()
	for i := 0; i < 10000; i++ {
		fmt.Sprintf("%d", i)
	}
	fmt.Println(time.Now().Sub(startTime))

	startTime = time.Now()
	for i := 0; i < 10000; i++ {
		strconv.Itoa(i)
	}
	fmt.Println(time.Now().Sub(startTime))
}

func TestJoin(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))

	fmt.Printf("len:=%d\n", len("hello world"))
}

func TestMap(t *testing.T) {
	m := make(map[int]string)

	m[1] = "a"
	m[2] = "b"
	fmt.Printf("m:%#v,len(m):%d\n", m, len(m))

	delete(m, 1)
	delete(m, 3)

	fmt.Printf("m:%#v,len(m):%d\n", m, len(m))
}

func TestString(t *testing.T) {
	b := strings.Builder{}
	s := []string{"foo", "bar", "baz"}
	for k, v := range s {
		if k == 0 {
			b.WriteString(v)
		} else {
			b.WriteString("," + v)
		}
	}
	fmt.Printf(b.String())

	b.Reset()

	b.WriteString("hello world")
	fmt.Printf(b.String())
}

func TestAppend(t *testing.T) {
	src := []byte(nil)
	src = strconv.AppendBool(src, false)
	src = strconv.AppendInt(src, 3, 10)
	fmt.Printf("%b\n", src)
	fmt.Printf("%s\n", src)
}

func TestRegexp(t *testing.T) {

	src := "41.2 aa.22 88.1 cca 9887 123 abc.325"
	reg := regexp.MustCompile(`\d+\.\d+`)
	if reg == nil {
		return
	}
	result := reg.FindAllStringSubmatch(src, -1)
	fmt.Printf("%+v\n", result)
}

func TestSlice2Json(t *testing.T) {
	slice := []int32{1, 9, 8}

	buff, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("err:%+v\n", err)
		return
	}

	m := []int32(nil)
	err = json.Unmarshal(buff, &m)
	if err != nil {
		fmt.Printf("err:%+v\n", err)
		return
	}

	fmt.Printf("%+v\n", m)
}

func TestOsFile(t *testing.T) {
	f, err := os.OpenFile("./tmp.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0)
	if err != nil {
		fmt.Printf("err:%+v\n", err)
	}
	defer f.Close()

	b := new(strings.Builder)
	for i := 0; i < 10; i++ {
		b.WriteString(fmt.Sprintf("hello %d\n\n", i))
	}
	f.WriteString(b.String())

}

func TestGoroutine(t *testing.T) {

	go Task()

	for {
		fmt.Printf("main\n")
		time.Sleep(time.Second)
	}
}

func Task() {
	go subTask()

	for i := 0; i < 10; i++ {
		fmt.Printf("Task\n")
		time.Sleep(time.Second)
	}
}

func subTask() {
	for {
		fmt.Printf("subTask\n")
		time.Sleep(time.Second)
	}
}

func TestWait(t *testing.T) {
	start := time.Now()
	fmt.Printf("hello \n")
	wait()
	fmt.Printf("world,%v\n", time.Since(start))

}

func wait() {
	//time.Sleep(6 * time.Second)
	fmt.Printf("task\n")
	//for {
	select {
	case <-time.After(6 * time.Second):
		fmt.Println("wait 3 second")
		return
	}
	//}
}

func TestMyWait(t *testing.T) {
	c := make(chan bool, 1)
	time.Sleep(6 * time.Second)
	go func() {
		fmt.Printf("ggg\n")
		c <- true
	}()

	<-c
	fmt.Printf("done\n")
}

func TestRunTime(t *testing.T) {

	go func() {
		fmt.Printf("start\n")
		time.Sleep(5 * time.Second)
		//	for {
		//		select {
		//		case <-time.After(5 * time.Second):
		//			return
		//		}
		//	}
	}()

	runtime.Gosched()
	fmt.Printf("done\n")
	//time.Sleep(100)
}

func TestMaxProc(t *testing.T) {
	fmt.Printf("%v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < 5; i++ {
		go fmt.Printf("aaa\n")
		fmt.Printf("bbb\n")
	}
}

func speak(c chan bool) {
	fmt.Printf("hello")
	c <- true
}

func speak2(c chan bool) {
	<-c
	fmt.Printf("world\n")
}

func TestSpeak(t *testing.T) {
	c := make(chan bool)
	go speak(c)
	go speak2(c)
	time.Sleep(1000)

}

func TestChannel000(t *testing.T) {
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("i ===> %d\n", i)
			ch <- i
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Printf("var:%d len(ch):%d cap(ch):%d\n", val, len(ch), cap(ch))
	}
}

func TestChannel002(t *testing.T) {
	ch := produce()
	consume(ch)
}

func produce() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		defer close(ch)
	}()
	return ch
}

func consume(ch <-chan int) {
	for val := range ch {
		fmt.Printf("val:%d\n", val)
	}
}

func TestTimer(t *testing.T) {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Printf("时间到\n")
	}()

	timer.Stop()

	fmt.Printf("done\n")
}

func TestFei(t *testing.T) {
	ch := make(chan int)
	quit := make(chan struct{})

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d ", <-ch)
		}
		quit <- struct{}{}
	}()

	fei(ch, quit)
}

func fei(ch chan int, quit chan struct{}) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			return
		}
	}
}
