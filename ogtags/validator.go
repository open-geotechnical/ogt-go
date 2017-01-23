package ogtags

import (
	"fmt"
	"io/ioutil"
)

func Rule_1(file_path string)[]error {

	errors := make([]error, 0, 0)

	bites, err := ioutil.ReadFile(file_path)
	if err != nil {
		errors = append(errors, err)
		return errors
	}
	fmt.Println("bites=", len(bites))
	//for b := range bites {
		// TODo check ascii

	//}

	return errors
}