package main

import (
	//"bytes"
	//"encoding/json"
	//"fmt"
	//"io/ioutil"
	//"log"
	"flag"
	//"strings"

	"bitbucket.org/daf0dil/ags2go/server"
	"bitbucket.org/daf0dil/ags2go/ags4"
)


//var Tabs = make(map[string][]ags4.Group)

func main() {

	// TODO check listen is a valid address/port etc
	listen := flag.String("listen", "0.0.0.0:13777", "HTTP server address and port")

	workspace_dir := flag.String("workspace_dir", "/home/ags/workspace", "Path to `workspace` dir")

	go ags4.InitLoad(*workspace_dir + "/ags-def-json/4")

	server.Start(*listen)


}

