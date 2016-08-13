
package ags4

import (

	"io/ioutil"
	"path"
)

func ParseExample(ex_file string) (*Document, error){

	file_path := examplesDir + "/" + ex_file

	return ParseFile(file_path)

}

func ParseFile(file_path string) (*Document, error){

	doc := NewDocument()
	bites, err := ioutil.ReadFile(file_path)
	if err != nil {
		return doc, err
	}
	doc.FileName = path.Base(file_path)
	doc.Source = string(bites)
	err = doc.Parse()
	return doc, err
}


