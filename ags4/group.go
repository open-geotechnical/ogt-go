

package ags4

// AGS Group .. ta bill ..was Node
type Group struct {
	Group   string     `json:"GROUP"`
	Heading []string   `json:"HEADING"`
	Unit    []string   `json:"UNIT"`
	Type    []string   `json:"TYPE"`
	Data    [][]string `json:"DATA"`
}