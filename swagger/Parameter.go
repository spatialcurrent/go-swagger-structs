package swagger

type Parameter struct {
	Name        string      `json:"name" yaml:"name"`
	Required    bool        `json:"required" bson:"required"`
	Type        string      `json:"type" yaml:"type"`
	In          string      `json:"in" yaml:"in"`
	Description string      `json:"description,omitempty" yaml:"description,omitempty"`
	Default     interface{} `json:"default,omitempty" yaml:"default,omitempty"`
	Minimum     *int        `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum     *int        `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	Enumeration []string    `json:"enum,omitempty" yaml:"enum,omitempty"`
}
