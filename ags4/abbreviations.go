

package ags4

//
type Abbrev struct {
	Group   string     `json:"GROUP"`
	Heading []string   `json:"HEADING"`
	Unit    []string   `json:"UNIT"`
	Type    []string   `json:"TYPE"`
	Data    [][]string `json:"DATA"`
}