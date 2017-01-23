
package ogtags

import (
	"io/ioutil"
)

func GetExamples()([]string, error){

	lst := make([]string, 0, 0)
	files, err := ioutil.ReadDir(examplesDir)
	if err != nil {
		return lst, err
	}

	for _, f := range files {
		lst = append(lst, f.Name())
	}

	return lst, nil
}



func ParseExample(ex_file string) (*Document, error){

	file_path := examplesDir + "/" + ex_file

	return NewDocumentFromFile(file_path)

}
