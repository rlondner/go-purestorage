// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

// Array type describes the array object returned by the API
type Array struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Model   string `json:"model,omitempty"`
	OS      string `json:"os,omitempty"`
	Version string `json:"version,omitempty"`
	AsOf    int    `json:"_as_of,omitempty"`
}

// Tag type describes the array object returned by the API
type Tag struct {
	Key               string            `json:"key,omitempty"`
	Namespace         string            `json:"namespace,omitempty"`
	Resource          map[string]string `json:"resource,omitempty"`
	TagOrganizationID int               `json:"tag_organization_id,omitempty"`
	Value             string            `json:"value,omitempty"`
}
