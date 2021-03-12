package main

import (
	"io"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/pkg/log"
)

///*
//	 v1 -----------------------------------
//*/
//func HomeFunc(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("hello"))
//}
//
//func main() {
//	http.HandleFunc("/", HomeFunc)
//	http.ListenAndServe("127.0.0.1:8000", nil)
//}

///*
//	 v2 -----------------------------------
//*/
//func main() {
//	mux := http.NewServeMux()
//
//	mux.Handle("/", &HomeHandler{})
//	mux.HandleFunc("/hello", Hello)
//
//	if err := http.ListenAndServe(":8000", mux); err != nil {
//		log.Error("ListenAndServe error(%+v)", err)
//	}
//
//}
//
//type HomeHandler struct {
//}
//
//func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	_, _ = io.WriteString(w, "[/]")
//
//}
//
//func Hello(w http.ResponseWriter, r *http.Request) {
//	_, _ = io.WriteString(w, "[Hello], url:"+r.URL.String())
//}

///*
//	 v3 -----------------------------------
//*/
//func
var (
	mux map[string]func(w http.ResponseWriter, r *http.Request)
)

func main() {
	s := http.Server{
		Addr:              "127.0.0.1:8000",
		Handler:           &myHandler{},
		TLSConfig:         nil,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	mux = map[string]func(w http.ResponseWriter, r *http.Request){}
	mux["/hello"] = sayhello
	mux["/bye"] = saybye

	if err := s.ListenAndServe(); err != nil {
		log.Error("ListenAndServe error(%+v)", err)
	}
}

type myHandler struct {
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//_, _ = io.WriteString(w, "URL:"+r.URL.String())

	h, ok := mux[r.URL.String()]
	if !ok {
		log.Warn("mux can not find handler,Url:%+v", r.URL.String())
		return
	}
	h(w, r)
	log.Info("ServeHttp, Url:%s", r.URL.String())
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "hello, URL:"+r.URL.String())
}
func saybye(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "bye, URL:"+r.URL.String())
}
