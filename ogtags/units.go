

package ogtags

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"strings"
)

// Represent and AGS unit and its description.
type Unit struct {

	// eg DegC, kN/m2  (latin remember)
	Unit string 		` json:"unit"  ags:"UNIT_UNIT" `

	// eg kiloNewtons per square metre (latin)
	Description string 	` json:"description"  `

	// This is idea with an UTF-8 symbol (breaks spec)
	//Symbol string 	` json:"symbol" `
}

var UnitsMap map[string]Unit


func LoadUnitsDDFromFile(file_path string) error {

	bites, err := ioutil.ReadFile(file_path)
	if err != nil {
		return err
	}

	var units = make([]Unit, 0, 0)
	err = json.Unmarshal(bites, &units)
	if err != nil {
		return err
	}
	var mutex = &sync.Mutex{}
	mutex.Lock()
	for _, u := range units {
		UnitsMap[u.Unit] = u
	}
	mutex.Unlock()
	return nil

}


// Returns true if the unit exists
func UnitExists(unit_str string) bool {
	unit_str = strings.TrimSpace(unit_str)
	if unit_str == "" {
		return false
	}
	_, found := UnitsMap[unit_str]
	return found
}


func GetUnits() []Unit {
	lst := make([]Unit, 0, len(UnitsMap))
	for _, u := range UnitsMap {
		lst = append(lst, u)
	}
	return lst
}