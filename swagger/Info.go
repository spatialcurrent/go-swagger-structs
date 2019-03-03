// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================
package swagger

type Info struct {
	Version        string   `json:"version" yaml:"version"`
	Title          string   `json:"title" yaml:"title"`
	Description    string   `json:"description" yaml:"description"`
	TermsOfService string   `json:"termsOfService" yaml:"termsOfService"`
	Contact        *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License `json:"license,omitempty" yaml:"license,omitempty"`
}
