// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewEventRequest new event request
//
// swagger:model NewEventRequest
type NewEventRequest struct {

	// event id
	EventID int64 `json:"event_id,omitempty"`

	// event time
	// Format: date-time
	EventTime strfmt.DateTime `json:"event_time,omitempty"`

	// event type
	EventType string `json:"event_type,omitempty"`

	// payload
	Payload string `json:"payload,omitempty"`

	// user id
	UserID int64 `json:"user_id,omitempty"`
}

// Validate validates this new event request
func (m *NewEventRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewEventRequest) validateEventTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EventTime) { // not required
		return nil
	}

	if err := validate.FormatOf("event_time", "body", "date-time", m.EventTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this new event request based on context it is used
func (m *NewEventRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewEventRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewEventRequest) UnmarshalBinary(b []byte) error {
	var res NewEventRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
