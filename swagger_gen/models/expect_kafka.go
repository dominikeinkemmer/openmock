// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ExpectKafka expect kafka
// swagger:model ExpectKafka
type ExpectKafka struct {

	// kafka topic to listen on
	Topic string `json:"topic,omitempty"`
}

// Validate validates this expect kafka
func (m *ExpectKafka) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExpectKafka) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExpectKafka) UnmarshalBinary(b []byte) error {
	var res ExpectKafka
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
