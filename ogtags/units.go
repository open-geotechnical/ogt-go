
package ogtags

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"strings"
)

// Represent's an AGS defined unit and its description, etc, eg
//   DegC = Degrees Centigrade
//   yyymmdd = Date in Internatinayl formati
//
// Note that the developeri is interestedo in this in returning
// different measures.. eg x.Kg, x.Lbs, x.m, x.Ki, x.convert()

type Unit struct {

	// eg DegC, kN/m2  (latin remember)
	Unit string 		` json:"unit"  ags:"UNIT_UNIT" `

	// eg kiloNewtons per square metre (latin)
	Description string 	` json:"description"  `

	// This is idea with an UTF-8 symbol (breaks spec)
	// TODO Symbol could be interesting to return degc, html, etc
	Symbol string 	` json:"symbol" `
}

// Memory cache for an unit loop up, automcomplete and validation etc
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


// Returns true if the unit exists/reconginsed..
// Note is case sensitive.. as in k / K  = kilo / Kelvin
func UnitExists(unit_str string) bool {
	unit_str = strings.TrimSpace(unit_str)
	if unit_str == "" {
		return false
	}
	_, found := UnitsMap[unit_str]
	return found
}

// Returns a list of recognised units of measuring stuff
func Units() []Unit {
	lst := make([]Unit, 0, len(UnitsMap))
	for _, u := range UnitsMap {
		lst = append(lst, u)
	}
	return lst
}