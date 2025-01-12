// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package audit

import (
	"encoding/json"
	"fmt"
	"io"
)

var _ Writer = (*JSONWriter)(nil)

// JSONWriter is a Writer implementation that structures data into a JSON format.
type JSONWriter struct {
	Prefix string
}

func (f *JSONWriter) WriteRequest(w io.Writer, req *AuditRequestEntry) error {
	if req == nil {
		return fmt.Errorf("request entry was nil, cannot encode")
	}

	if len(f.Prefix) > 0 {
		_, err := w.Write([]byte(f.Prefix))
		if err != nil {
			return err
		}
	}

	enc := json.NewEncoder(w)
	return enc.Encode(req)
}

func (f *JSONWriter) WriteResponse(w io.Writer, resp *AuditResponseEntry) error {
	if resp == nil {
		return fmt.Errorf("response entry was nil, cannot encode")
	}

	if len(f.Prefix) > 0 {
		_, err := w.Write([]byte(f.Prefix))
		if err != nil {
			return err
		}
	}

	enc := json.NewEncoder(w)
	return enc.Encode(resp)
}
