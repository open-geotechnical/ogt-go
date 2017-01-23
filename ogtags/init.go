
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

	err = LoadUnitsFromFile(filepath.FromSlash(unitsFile))
	if err != nil {
		all_errors = append(all_errors, err)
	}

	// groups and their heading need to be loaded first
	err = LoadGroupsFromFile(filepath.FromSlash(groupsFile))
	if err != nil {
		all_errors = append(all_errors, err)
	}

	err = LoadAbbrsFromFile(filepath.FromSlash(abbrsFile))
	if err != nil {
		all_errors = append(all_errors, err)
	}


	if len(all_errors) > 0 {
		fmt.Println("WTF?", all_errors)
		// Tantrum()
	}

	fmt.Println("Loading AGS DONE")


}

