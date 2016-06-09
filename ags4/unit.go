

package ags4

// AGS Units are defined in



type Unit struct {
	Unit string 		` json:"unit" db:"unit" `
	Description string 	` json:"description" db:"description" `
}

// Returns true if the unit exists
func UnitExists(unit string) bool {

	// TO do is look up the

	return false
}