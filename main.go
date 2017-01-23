package main

// This is WIP and some r&d into AGS4 data format
// This is not official = unofficial playground
// but hopefully in due course a reliable lib in foss
//
// Important Note: Its r+d so not production ready.. yet ;-)

import (
	//"bytes"
	//"encoding/json"
	//"fmt"
	//"io/ioutil"
	//"log"
	"flag"
	//"strings"

	"github.com/open-geotechnical/ogt-ags-go/ogtags"
	"github.com/open-geotechnical/ogt-ags-go/server"
)

func main() {

	// TODO check listen is a valid address/port etc
	listen := flag.String("server", "0.0.0.0:13777", "HTTP server address and port")

	//ags_data_dict := flag.String("ags_data_dict", "/ome/ags/workspace", "Path to `ags_data_dict` dir")
	ags_data_dict := flag.String("ags_data_dict", "c:\\z_pete_stuff\\src\\bitbucket.org\\daffodil\\ags-data-dict", "Path to `ags_data_dict` dir")

	// Initialise the datadict stores
	ogtags.InitLoad(*ags_data_dict)

	// TODO make server a flag, for now its on for fun
	if true {
		server.Start(*listen)
	}


}
