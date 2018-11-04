package swagger

type Contact struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
	Url   string `json:"url,omitempty" yaml:"url,omitempty"`
}
