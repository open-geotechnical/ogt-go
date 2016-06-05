
package ags2go

// Units are defined in ags units file eg
//
//   gal = gallon
//   in2 = inch square
//   mg/kg = milligrams per kilogram
//   MN = megaNewton
//   ppm = parts per million


// Units lookup and memory cache
var Units[]AGSUnit


// Data Type Container
type AGSUnit struct {
	Unit string ` json:"unit" `
	UnitDescription string ` json:"unit_description" `
}




// Checks whether sting with unit is valid, ie does it exists in table
func IsValid(unit string) bool {
	// TODO
	return true
}





