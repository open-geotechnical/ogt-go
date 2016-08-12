
package ags4

import (

	"io/ioutil"
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
	doc.FileName = file_path
	doc.Source = string(bites)
	err = doc.Parse()
	return doc, err
}




// RudeParser is a fast parser and a than rubs the dub
// Clothethes off. here it all is...
// and unless its kinda proper, then it thows a towel..
//
// So the rude parser is designed to be fast and eats it in an out
// like there's no tomorrow..
// So by design we thig mayne this is the "OUTER"
// ie the file makes sense and is eatable..
// and mistakes here we crash
// Next step would be rude validator...
func RudeParser(file_path string)  (Document, error) {

	doci := Document{}

	// do things
	// if it crashed, then error and expception

	return doci, nil

}

// Elongated parse, that passes errors
// and gives debug info..by predro
func Rude(file string)  (Document, error) {

	doci := Document{}

	// do things
	// real slow line by lin
	// and is error then try and continue until impossible..

	return doci, nil

}


