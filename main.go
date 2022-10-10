package main

import (
	"fmt"
	"net/http"
)

type Server_info struct {
	ip   string
	prot int
	path string
}

func (info *Server_info) address() string {
	if info.ip == "" || info.prot == 0 {

		info.ip = "127.0.0.1"
		info.prot = 9000
	}
	// url := info.ip + ":" + strconv.Itoa(info.prot)
	url := fmt.Sprintf("%s:%d", info.ip, info.prot)
	fmt.Printf("服务地址设定为:%s\n", url)
	return url
}

func main() {
	server_info := Server_info{
		ip:   "127.0.0.1",
		prot: 9000,
		path: "./",
	}

	http.Handle("/", http.FileServer(http.Dir(server_info.path)))
	e := http.ListenAndServe(server_info.address(), nil)
	if e != nil {
		fmt.Println(e.Error())
	}
}
