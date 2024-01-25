package assistant

type ProjectCommand struct {
	Decode *Command `json:"decode"`
	Encode *Command `json:"encode"`
	Prod   string   `json:"prod"`
	Stage  string   `json:"stage"`
	Test   string   `json:"test"`
}
