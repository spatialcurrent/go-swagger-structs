package swagger

type Definition struct {
	Type       string              `json:"type" yaml:"type"`
	Required   []string            `json:"required,omitempty" yaml:"required,omitempty"`
	Properties map[string]Property `json:"properties,omitempty" yaml:"properties,omitempty"`
}
