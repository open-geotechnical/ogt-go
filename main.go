package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	LF = 10

	DATA    = "DATA"
	GROUP   = "GROUP"
	HEADING = "HEADING"
	TYPE    = "TYPE"
	UNIT    = "UNIT"
)

// AGS order
type Tab struct {
	Group   string     `json:"GROUP"`
	Heading []string   `json:"HEADING"`
	Unit    []string   `json:"UNIT"`
	Type    []string   `json:"TYPE"`
	Data    [][]string `json:"DATA"`
}

var Tabs = make(map[string][]Tab)

func main() {

	// func rude(cache string)
	rude(os.Args[1])

	// func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
	bytes, _ := json.MarshalIndent(Tabs, "", "  ")
	fmt.Printf("%s", bytes)
}

/*
Function rude is a rudimentary parser for AGS standard data.
*/
func rude(archive string) {

	// func OpenReader(name string) (*ReadCloser, error)
	rc, err := zip.OpenReader(archive)
	if err != nil {
		log.Fatal(err)
	}
	defer rc.Close()

	for _, file := range rc.File {

		name := file.Name
		Tabs[name] = []Tab{}

		// func (f *File) Open() (rc io.ReadCloser, err error)
		rc1, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer rc1.Close()

		// func ReadAll(r io.Reader) ([]byte, error)
		bloc, err := ioutil.ReadAll(rc1)
		if err != nil {
			log.Fatal(err)
		}

		// func NewBuffer(buf []byte) *Buffer
		buf := bytes.NewBuffer(bloc)
		tab := new(Tab)

		for {
			// func (b *Buffer) ReadString(delim byte) (line string, err error)
			// assumes \r\n not the reverse
			line, err := buf.ReadString(LF)
			if err != nil {
				break
			}

			// func Trim(s string, cutset string) string
			line = strings.Trim(line, " \r\n")

			// AGS format terminated with redundant \r\n sentinel
			if len(line) == 0 {
				Tabs[name] = append(Tabs[name], *tab)
				tab = new(Tab)
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
					tab.Group = token
				case HEADING:
					tab.Heading = append(tab.Heading, token)
				case TYPE:
					tab.Type = append(tab.Type, token)
				case UNIT:
					tab.Unit = append(tab.Unit, token)
				}
			}

			if mode == DATA {
				tab.Data = append(tab.Data, data)
			}
		}
	}
}
