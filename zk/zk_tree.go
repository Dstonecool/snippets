package main

import (
	"flag"
	"github.com/samuel/go-zookeeper/zk"
	"path"
	"time"
)

func main() {
	path := ""

	flag.StringVar(&path, "path", "/", "zookeeper node path")
	flag.Parse()

	conn, _, err := zk.Connect(flag.Args(), time.Second)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	traverse(conn, path)
}

func traverse(conn *zk.Conn, root string) {
	children, _, err := conn.Children(root)
	if err != nil {
		panic(err)
	}
	for _, c := range children {
		p := path.Join(root, c)
		println(p)
		traverse(conn, p)
	}
}
