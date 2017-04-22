
package ogtags

import (
	"fmt"
	"path/filepath"
)

// The directory where `ags-data-dict` is
// these  variables set with ogtags.InitLoad()
var DataDictPath string = ""

var groupsFile string = ""
var abbrsFile string = ""
var unitsFile string = ""
var dataTypesFile string = ""

var examplesDir string = ""



// Ide here is to read all the def files..
// This needs to be a spperate function and reloadable with mutex
//
// The initialize loop lock sout any requests bases on update..
// we expect ags to updates..
// and in memory foroever as in a count..
// However we can be fast by keeping th top 20 or so
// hit list.. TO do this as its a process we keep  here..
//
func InitLoad(ags_data_dict_dir string) {

	fmt.Println("Loading AGS", ags_data_dict_dir)

	DataDictPath = ags_data_dict_dir

	groupsFile = DataDictPath + "/ags4/groups.json"
	abbrsFile = DataDictPath + "/ags4/abbreviations.json"
	unitsFile = DataDictPath + "/ags4/units.json"
	dataTypesFile = DataDictPath + "/ags4/data_types.json"

	examplesDir = DataDictPath + "/ags4_examples"

	var all_errors []error
	var err error

	err = LoadUnitsDDFromFile(filepath.FromSlash(unitsFile))
	if err != nil {
		all_errors = append(all_errors, err)
	}

	err = LoadDataTypesDDFromFile(filepath.FromSlash(dataTypesFile))
	if err != nil {
		all_errors = append(all_errors, err)
	}

	err = LoadGroupsDDFromFile(filepath.FromSlash(groupsFile))
	if err != nil {
		all_errors = append(all_errors, err)
	}

	err = LoadAbbrsDDFromFile(filepath.FromSlash(abbrsFile))
	if err != nil {
		all_errors = append(all_errors, err)
	}


	if len(all_errors) > 0 {
		fmt.Println("errors")
		for _, ee := range all_errors {
			fmt.Println("\t", ee.Error())
		}
	}
	fmt.Println("Loading AGS DONE")


}

