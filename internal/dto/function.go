package dto

type Function struct {
	Name string `json:"name"`
	File string `json:"file"`
	Line int    `json:"line"`
}
