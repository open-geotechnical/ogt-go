

package ogtags

import (
	"encoding/json"
	"io/ioutil"
)

// Represent and AGS unit and its description.
type Unit struct {

	// eg DegC, kN/m2  (latin remember)
	Unit string 		` json:"unit"  ags:"UNIT_UNIT" `

	// eg kiloNewtons per square metre (latin)
	Description string 	` json:"description"  `

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
func UnitAutocomplete(txt string) (bool, []string){

	yipee := UnitExists(txt)

	hints := UnitsMatching(txt)

	return yipee, hints

}

func UnitsMatching(txt string) []string {

	matches := make([]string, 0, 0)
	// Do some clever searches
	// of the memtree and terun some matchin, case insensitive


	return matches


}