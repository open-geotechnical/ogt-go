package main

import (
	//"bytes"
	//"encoding/json"
	//"fmt"
	//"io/ioutil"
	//"log"
	"flag"
	//"strings"

	"github.com/open-geotechnical/ogt-ags-go/server"
	"github.com/open-geotechnical/ogt-ags-go/ags4"
)


func main() {

	// TODO check listen is a valid address/port etc
	listen := flag.String("listen", "0.0.0.0:13777", "HTTP server address and port")

	//ags_data_dict := flag.String("ags_data_dict", "/ome/ags/workspace", "Path to `ags_data_dict` dir")
	ags_data_dict := flag.String("ags_data_dict", "c:\\z_pete_stuff\\src\\bitbucket.org\\daffodil\\ags-data-dict", "Path to `ags_data_dict` dir")

	ags4.InitLoad(*ags_data_dict)

	if false {
		server.Start(*listen)
	}

}

