// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExportTemplate export template
//
// swagger:model ExportTemplate
type ExportTemplate struct {

	// Content type
	// Required: true
	ContentType *string `json:"content_type"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// File extension
	//
	// Extension to append to the rendered filename
	// Max Length: 15
	FileExtension string `json:"file_extension,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// MIME type
	//
	// Defaults to <code>text/plain</code>
	// Max Length: 50
	MimeType string `json:"mime_type,omitempty"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// Template code
	//
	// The list of objects being exported is passed as a context variable named <code>queryset</code>.
	// Required: true
	// Min Length: 1
	TemplateCode *string `json:"template_code"`

	// template language
	TemplateLanguage *ExportTemplateTemplateLanguage `json:"template_language,omitempty"`
}

// Validate validates this export template
func (m *ExportTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileExtension(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMimeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExportTemplate) validateContentType(formats strfmt.Registry) error {

	if err := validate.Required("content_type", "body", m.ContentType); err != nil {
		return err
	}

	return nil
}

func (m *ExportTemplate) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 200); err != nil {
		return err
	}

	return nil
}

func (m *ExportTemplate) validateFileExtension(formats strfmt.Registry) error {

	if swag.IsZero(m.FileExtension) { // not required
		return nil
	}

	if err := validate.MaxLength("file_extension", "body", string(m.FileExtension), 15); err != nil {
		return err
	}

	return nil
}

func (m *ExportTemplate) validateMimeType(formats strfmt.Registry) error {

	if swag.IsZero(m.MimeType) { // not required
		return nil
	}

	if err := validate.MaxLength("mime_type", "body", string(m.MimeType), 50); err != nil {
		return err
	}

	return nil
}

func (m *ExportTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	return nil
}

func (m *ExportTemplate) validateTemplateCode(formats strfmt.Registry) error {

	if err := validate.Required("template_code", "body", m.TemplateCode); err != nil {
		return err
	}

	if err := validate.MinLength("template_code", "body", string(*m.TemplateCode), 1); err != nil {
		return err
	}

	return nil
}

func (m *ExportTemplate) validateTemplateLanguage(formats strfmt.Registry) error {

	if swag.IsZero(m.TemplateLanguage) { // not required
		return nil
	}

	if m.TemplateLanguage != nil {
		if err := m.TemplateLanguage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template_language")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExportTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExportTemplate) UnmarshalBinary(b []byte) error {
	var res ExportTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ExportTemplateTemplateLanguage Template language
//
// swagger:model ExportTemplateTemplateLanguage
type ExportTemplateTemplateLanguage struct {

	// label
	// Required: true
	// Enum: [Django Jinja2]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [django jinja2]
	Value *string `json:"value"`
}

func (m *ExportTemplateTemplateLanguage) UnmarshalJSON(b []byte) error {
	type ExportTemplateTemplateLanguageAlias ExportTemplateTemplateLanguage
	var t ExportTemplateTemplateLanguageAlias
	if err := json.Unmarshal([]byte("{\"id\":20,\"label\":\"Jinja2\",\"value\":\"jinja2\"}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = ExportTemplateTemplateLanguage(t)
	return nil
}

// Validate validates this export template template language
func (m *ExportTemplateTemplateLanguage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var exportTemplateTemplateLanguageTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Django","Jinja2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		exportTemplateTemplateLanguageTypeLabelPropEnum = append(exportTemplateTemplateLanguageTypeLabelPropEnum, v)
	}
}

const (

	// ExportTemplateTemplateLanguageLabelDjango captures enum value "Django"
	ExportTemplateTemplateLanguageLabelDjango string = "Django"

	// ExportTemplateTemplateLanguageLabelJinja2 captures enum value "Jinja2"
	ExportTemplateTemplateLanguageLabelJinja2 string = "Jinja2"
)

// prop value enum
func (m *ExportTemplateTemplateLanguage) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, exportTemplateTemplateLanguageTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExportTemplateTemplateLanguage) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("template_language"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("template_language"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var exportTemplateTemplateLanguageTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["django","jinja2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		exportTemplateTemplateLanguageTypeValuePropEnum = append(exportTemplateTemplateLanguageTypeValuePropEnum, v)
	}
}

const (

	// ExportTemplateTemplateLanguageValueDjango captures enum value "django"
	ExportTemplateTemplateLanguageValueDjango string = "django"

	// ExportTemplateTemplateLanguageValueJinja2 captures enum value "jinja2"
	ExportTemplateTemplateLanguageValueJinja2 string = "jinja2"
)

// prop value enum
func (m *ExportTemplateTemplateLanguage) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, exportTemplateTemplateLanguageTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExportTemplateTemplateLanguage) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("template_language"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("template_language"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExportTemplateTemplateLanguage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExportTemplateTemplateLanguage) UnmarshalBinary(b []byte) error {
	var res ExportTemplateTemplateLanguage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}