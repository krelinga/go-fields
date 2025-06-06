package main

import (
	"encoding/json"
)

var (
	AdultFieldNil     = AdultField{}
	AdultFieldDefault = NewAdultField(true)
)

type AdultField struct {
	ptr *bool
}

func (f AdultField) GetDefault() bool {
	out, _ := f.GetSafe()
	return out
}

func (f AdultField) GetSafe() (bool, bool) {
	if f.ptr == nil {
		return *AdultFieldDefault.ptr, false
	}
	return *f.ptr, true
}

func (f AdultField) GetPanic() bool {
	out, ok := f.GetSafe()
	if !ok {
		panic("nil AdultField")
	}
	return out
}

func (f *AdultField) UnmarshalJSON(in []byte) error {
	var val *bool
	if err := json.Unmarshal(in, &val); err != nil {
		return err
	}
	if val == nil {
		*f = AdultFieldDefault
	} else {
		*f = NewAdultField(*val)
	}
	return nil
}

func (f *AdultField) setDefault() {
	*f = AdultFieldDefault
}

func NewAdultField(value bool) AdultField {
	return AdultField{ptr: func() *bool {
		out := new(bool)
		*out = value
		return out
	}()}
}

func MergeAdultFields(inputs ...AdultField) AdultField {
	var result AdultField
	for _, input := range inputs {
		if input == AdultFieldNil {
			continue
		}
		if result == AdultFieldNil || (result == AdultFieldDefault && input != AdultFieldDefault) {
			result = input
			continue
		}
	}
	return result
}

func (f AdultField) MarshalJSON() ([]byte, error) {
	if f == AdultFieldNil || f == AdultFieldDefault {
		return json.Marshal(nil)
	}
	return json.Marshal(f.GetPanic())
}

type setDefaulter interface {
	setDefault()
}
