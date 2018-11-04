package swagger

type Operation struct {
	Tags        []string            `json:"tags" yaml:"tags"`
	Summary     string              `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string              `json:"description,omitempty" yaml:"description,omitempty"`
	Consumes    []string            `json:"consumes,omitempty" yaml:"consumes,omitempty"`
	Produces    []string            `json:"produces,omitempty" yaml:"produces,omitempty"`
	Parameters  []Parameter         `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Responses   map[string]Response `json:"responses" yaml:"responses"`
}
