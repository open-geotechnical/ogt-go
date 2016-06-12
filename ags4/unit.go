

package ags4

// AGS Units are defined by AGS..
// eg
// "%" = "percentage"
// "MPN/100ml" = "most probable number per 100 millilitres",
// "DegC" = "degree Celsius"
//
// NOte that AGS is latin chars, so not of them windows
// and unicodec haracters in there.. (which crashed my system)
// So we need to validate that,
// even though golang is UTF8.. SIGH!
// But if we go to UTF-* then we can have a proper Degree^C symbol
// uneditable by a windows users... sigh!!..
//
//
// Dreaded SPEC and DEF
// AGS4 "UNITS", spec/def will be loaded into
// this var at startup.. later we need hot..reload
var Units = map[string]Unit


type Unit struct {

	// eg DegC, kN/m2  (latin remember)
	Unit string 		` json:"unit" db:"unit" `

	// eg kiloNewtons per square metre (latin)
	Description string 	` json:"description" db:"description" `

	// This is daffo special of OC. proper with UTF-8 symbol
	Symbol string 	` json:"symbol" db:"symbol" `
}

// Returns true if the unit exists
// Whist this is a "definition endevour"
// we also want a typs and auto complete
func UnitExists(unit string) bool {

	// TO do is look up the map for key

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

	matches := make(string,0,0)
	// Do some clever searches
	// of the memtree and terun some matchin, case insensitive
	// meed some rapid magic..
	// et wahat tz ? planet..
	// U in USA ? imperial
	// Default = Metric ;;-)))


	return matches


}