package io

import "gopkg.in/yaml.v3"

/*
Common abstract struct so we can quickly indicate whether this struct is
exportable or not.
*/
type Marshalable struct{}

func (m *Marshalable) Marshal() (out []byte, err error) {
	return yaml.Marshal(m)
}

func (m *Marshalable) Unmarshal(in []byte) (err error) {
	return yaml.Unmarshal(in, m)
}
