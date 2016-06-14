

package ags4


// The "Project"(PROJ)
// is the "main sheet/tab" for ags
//
/*
"GROUP","PROJ"
"HEADING","PROJ_ID","PROJ_NAME","PROJ_LOC","PROJ_CLNT","PROJ_CONT","PROJ_ENG"
"UNIT","","","","","",""
"TYPE","ID","X","X","X","X","X"
"DATA","Prj-AG","Acme Gasworks","Greenwich, London","Acme Enterprises","Acme Monitoring Ltd","Acme Consulting"
*/

type Project struct {

	// The ID of this job.
	// as the docs will be flying around and
	// differnt parts.. this ID is probably
	// the
	ID string ` json:"PROG_ID" ags:"PROG_ID" `

	Name string ` json:"PROJ_NAME" ags:"PROJ_NAME" `

	// TODO
	//Location string
}

func NewProjectFromAGS(ags_block string ) Project {

	p := new(Project)
	p.ID = "Project factorty ?"
	// Do Rude Parser
	return p
}
