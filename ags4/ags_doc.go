

package ags4


// The Ags4Doc represents the data structure
// for an ags file
// The "struct should be
// cool and funky and serialiabled into json, yaml and AGS
// <NO xml.hell please ie no arrays said Bigll>
//



type AGS4Doc struct {

	// Project is included and main "ref"
	// on each "file=" project = spreashsheet name in a way
	Project Project ` json:"project"  ags:"PROJ" `

	// Units are required, eg %, degC, gal, Ltr
	Units []Unit  ` json:"units"  ags:"TYPE" `

	// Types ?? eg 2dp, text, date
	// BTW its GROUP in AGS
	Types[]string 	` json:"units"  ags:"UNITS" `

	// The groups in this Doc
	Groups []Group ` json:"groups"  ags:"groups" `

}



