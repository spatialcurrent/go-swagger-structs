package swagger

type Info struct {
	Version        string   `json:"version" yaml:"version"`
	Title          string   `json:"title" yaml:"title"`
	Description    string   `json:"description" yaml:"description"`
	TermsOfService string   `json:"termsOfService" yaml:"termsOfService"`
	Contact        *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License `json:"license,omitempty" yaml:"license,omitempty"`
}
