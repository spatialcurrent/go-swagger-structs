// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================
package swagger

type Definition struct {
	Type       string              `json:"type" yaml:"type"`
	Required   []string            `json:"required,omitempty" yaml:"required,omitempty"`
	Properties map[string]Property `json:"properties,omitempty" yaml:"properties,omitempty"`
}
