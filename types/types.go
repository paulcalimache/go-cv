package types

import "errors"

type Format string

const (
	html Format = "html"
	pdf  Format = "pdf"
)

// Set implements pflag.Value.
func (f *Format) Set(v string) error {
	if v == string(html) || v == string(pdf) {
		*f = Format(v)
		return nil
	} else {
		return errors.New("must be 'html' or 'pdf'")
	}
}

// String implements pflag.Value.
func (f *Format) String() string {
	return string(*f)
}

// Type implements pflag.Value.
func (f *Format) Type() string {
	return "format"
}
