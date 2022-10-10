package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	ip   string
	prot int
	path string
	help bool
)

type Server_info struct {
	ip   string
	prot int
	path string
}

func (info *Server_info) address() string {

	// url := info.ip + ":" + strconv.Itoa(info.prot)
	url := fmt.Sprintf("%s:%d", info.ip, info.prot)
	fmt.Printf("启动服务，地址为:%s\n", url)
	return url
}

func main() {
	init_flag()
	server_info := Server_info{
		ip:   ip,
		prot: prot,
		path: path,
	}

	http.Handle("/", http.FileServer(http.Dir(server_info.path)))
	e := http.ListenAndServe(server_info.address(), nil)
	if e != nil {
		fmt.Println(e.Error())
	}
}

func init_flag() {
	flag.StringVar(&ip, "addr", "127.0.0.1", "ip地址,默认为127.0.0.1")
	flag.IntVar(&prot, "p", 9000, "端口号,默认为9000")
	flag.StringVar(&path, "path", "./", "路径,默认为./当前路径")
	flag.BoolVar(&help, "h", false, "帮助")
	flag.Parse()
	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}
}
