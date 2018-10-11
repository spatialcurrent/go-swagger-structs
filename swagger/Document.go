package swagger

type Document struct {
	Version  string          `json:"swagger" yaml:"swagger"`
	Schemes  []string        `json:"schemes" yaml:"schemes"`
	BasePath string          `json:"basePath" yaml:"basePath"`
	Host     string          `json:"host" yaml:"host"`
	External *External       `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Info     *Info           `json:"info,omitempty" yaml:"info,omitempty"`
	Paths    map[string]Path `json:"paths" yaml:"paths"`
}
