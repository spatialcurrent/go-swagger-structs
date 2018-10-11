package swagger

type Path struct {
	Summary     string    `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string    `json:"description,omitempty" yaml:"description,omitempty"`
	Get         Operation `json:"get,omitempty" yaml:"get,omitempty"`
	Post        Operation `json:"post,omitempty" yaml:"post,omitempty"`
}
