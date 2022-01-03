/*
本周作业
编写一个 HTTP 服务器
1.接收客户端 request，并将 request 中带的 header 写入 response header
2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4.当访问 localhost/healthz 时，应返回 200
*/

package httpserver

import (
	"log"
	"net/http"
	"net/http/pprof"
	"os"
)

func ServerMain() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexFunc)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/debug/pprof", pprof.Index)
	mux.HandleFunc("/debug/Profile", pprof.Profile)
	mux.HandleFunc("/debug/Symbol", pprof.Symbol)
	mux.HandleFunc("/debug/Trace", pprof.Trace)
	if err := http.ListenAndServe(":80", mux); err != nil {
		return
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("server alive")); err != nil {
	}
	return
}

func indexFunc(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	for k, v := range r.Header {
		header.Set(k, v[0])
	}
	header.Set("Version", os.Getenv("VERSION"))
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("hello world")); err != nil {
	}

	log.Printf("IP:%s, HTTP Response Code:%d\n", r.RemoteAddr, http.StatusOK)

	return
}
