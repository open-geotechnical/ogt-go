package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"bitbucket.org/daffodil/ags2go/ags4"
)


var Tabs = make(map[string][]ags4.Group)

func main() {

	// func rude(cache string)
	rude(os.Args[1])

	bytes, _ := json.MarshalIndent(Tabs, "", "  ")
	fmt.Printf("%s", bytes)
}



/*
Function rude is a rudimentary parser for AGS standard data. It takes as an argument the name of
the directory containing the AGS definition files.
*/
func rude(cache string) {

	// func ReadDir(dirname string) ([]os.FileInfo, error)
	files, err := ioutil.ReadDir(cache)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		name := file.Name()
		Tabs[name] = []Node{}

		// func ReadFile(filename string) ([]byte, error)
		bloc, err := ioutil.ReadFile(cache + "/" + name)
		if err != nil {
			log.Fatal(err)
		}

		// func NewBuffer(buf []byte) *Buffer
		buf := bytes.NewBuffer(bloc)
		node := new(Node)

		for {
			// func (b *Buffer) ReadString(delim byte) (line string, err error)
			// assumes \r\n not the reverse
			line, err := buf.ReadString(LF)
			if err != nil {
				break
			}

			// func Trim(s string, cutset string) string
			line = strings.Trim(line, " \r\n")

			// AGS format terminated with marking \r\n
			if len(line) == 0 {
				Tabs[name] = append(Tabs[name], *node)
				node = new(Node)
				continue
			}

			// used only during DATA cycle
			data := []string{}

			head := true
			mode := ""

			// func Split(s, sep string) []string
			for _, token := range strings.Split(line, ",") {
				// strip quotes
				token = strings.Trim(token, " \"")

				// set new mode
				if head {
					mode = token
					head = false
					continue
				}

				switch mode {
				case DATA:
					data = append(data, token)
				case GROUP:
					node.Group = token
				case HEADING:
					node.Heading = append(node.Heading, token)
				case TYPE:
					node.Type = append(node.Type, token)
				case UNIT:
					node.Unit = append(node.Unit, token)
				}
			}

			if mode == DATA {
				node.Data = append(node.Data, data)
			}
		}
	}
}
