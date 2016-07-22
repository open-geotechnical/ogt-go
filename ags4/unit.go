

package ags4

import (
	"encoding/json"
	"io/ioutil"
)

// Units are defined by AGS4 spec Rule #8
// Example Usage in an ags4 file
//
//	"GROUP","HORN"
//	"HEADING","LOCA_ID","HORN_TOP","HORN_BASE","HORN_ORNT","HORN_INCL","HORN_REM","FILE_FSET"
//	"UNIT","","m","m","deg","deg","",""
//	"TYPE","ID","2DP","2DP","0DP","0DP","X","X"
//
type Unit struct {

	// eg DegC, kN/m2  (latin remember)
	Unit string 		` json:"unit" db:"unit" ags:"UNIT_UNIT" `

	// eg kiloNewtons per square metre (latin)
	Description string 	` json:"description" db:"todo" `

	// This is daffo special with a proper with UTF-8 symbol (breaks spec)
	//Symbol string 	` json:"symbol" db:"symbol" `
}

// The memory cache variable loaded at startup (and relodable ?? )
var Units []Unit


func LoadUnitsFromFile(file_path string) error {

	bites, err := ioutil.ReadFile(file_path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bites, &Units)
	if err != nil {
		return err
	}

	return nil

}


// Returns true if the unit exists
func UnitExists(unit string) bool {
	if unit == "" {
		return false
	}
	for _, u := range Units {
		if u.Unit == unit {
			return true
		}
	}
	return false
}

// Unit autocomplete and valid an hints
// expect a request every  second.. whikst session engaged..
// or we cache to clinet.. Over to @bill
func UnitAutocomplete(txt string) (bool, []string){
	// imagine the user has types in a k/m when K/m was means
	// first we hekp with valid unit ?
	yipee := UnitExists(txt)

	hints := UnitsMatching(txt)

	return yipee, hints

}

func UnitsMatching(txt string) []string {

	matches := make([]string, 0, 0)
	// Do some clever searches
	// of the memtree and terun some matchin, case insensitive
	// meed some rapid magic..
	// et wahat tz ? planet..
	// U in USA ? imperial
	// Default = Metric ;;-)))


	return matches


}