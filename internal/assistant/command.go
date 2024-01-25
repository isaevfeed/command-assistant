package assistant

type Command struct {
	Command string
	Args    []string `json:"arguments"`
}
