
package ogtags

import (
	"io/ioutil"
)

func GetExamples()([]string, error){
	return LoadExamplesFromDir()
}


// load groups.json file to mem
func LoadExamplesFromDir() ([]string, error) {

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