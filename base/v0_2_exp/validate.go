// Copyright 2019 Red Hat, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.)

package v0_2_exp

import (
	"errors"

	"github.com/coreos/vcontext/path"
	"github.com/coreos/vcontext/report"
)

var (
	ErrTooManyResourceSources = errors.New("only one of the following can be set: inline, local, source")
	ErrMountUnitNoPath        = errors.New("path is required if with_mount_unit is true")
	ErrMountUnitNoFormat      = errors.New("format is required if with_mount_unit is true")
)

func (rs Resource) Validate(c path.ContextPath) (r report.Report) {
	var field string
	sources := 0
	if rs.Local != nil {
		sources++
		field = "local"
	}
	if rs.Inline != nil {
		sources++
		field = "inline"
	}
	if rs.Source != nil {
		sources++
		field = "source"
	}
	if sources > 1 {
		r.AddOnError(c.Append(field), ErrTooManyResourceSources)
	}
	return
}

func (fs Filesystem) Validate(c path.ContextPath) (r report.Report) {
	if fs.WithMountUnit == nil || !*fs.WithMountUnit {
		return
	}
	if fs.Path == nil || *fs.Path == "" {
		r.AddOnError(c.Append("path"), ErrMountUnitNoPath)
	}
	if fs.Format == nil || *fs.Format == "" {
		r.AddOnError(c.Append("format"), ErrMountUnitNoFormat)
	}
	return
}