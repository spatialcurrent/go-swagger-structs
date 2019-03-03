// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================
package swagger

type Path struct {
	Summary     string    `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string    `json:"description,omitempty" yaml:"description,omitempty"`
	Get         Operation `json:"get,omitempty" yaml:"get,omitempty"`
	Post        Operation `json:"post,omitempty" yaml:"post,omitempty"`
	Delete      Operation `json:"delete,omitempty" yaml:"delete,omitempty"`
}
