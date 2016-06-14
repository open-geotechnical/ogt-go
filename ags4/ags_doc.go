

package ags4


// The Ags4Doc represents the data structure
// for an ags file
// The "struct should be
// cool and funky and serialiabled into json, yaml and AGS
// <NO xml please>
//



type AGS4Doc struct {

	// Project is included and main "ref"
	Project Project ` json:"project"  ags:"PROJ" `

	// Units are required, eg %, degC, gal, Ltr
	Units []Unit  ` json:"units"  ags:"UNITS" `


	// wow
	Groups []Group ` json:"groups"  ags:"groups" `

}

