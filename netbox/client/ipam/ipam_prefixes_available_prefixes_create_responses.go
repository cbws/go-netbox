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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// IpamPrefixesAvailablePrefixesCreateReader is a Reader for the IpamPrefixesAvailablePrefixesCreate structure.
type IpamPrefixesAvailablePrefixesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesAvailablePrefixesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamPrefixesAvailablePrefixesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamPrefixesAvailablePrefixesCreateOK creates a IpamPrefixesAvailablePrefixesCreateOK with default headers values
func NewIpamPrefixesAvailablePrefixesCreateOK() *IpamPrefixesAvailablePrefixesCreateOK {
	return &IpamPrefixesAvailablePrefixesCreateOK{}
}

/*IpamPrefixesAvailablePrefixesCreateOK handles this case with default header values.

IpamPrefixesAvailablePrefixesCreateOK ipam prefixes available prefixes create o k
*/
type IpamPrefixesAvailablePrefixesCreateOK struct {
	Payload []*models.AvailablePrefix
}

func (o *IpamPrefixesAvailablePrefixesCreateOK) Error() string {
	return fmt.Sprintf("[POST /ipam/prefixes/{id}/available-prefixes/][%d] ipamPrefixesAvailablePrefixesCreateOK  %+v", 200, o.Payload)
}

func (o *IpamPrefixesAvailablePrefixesCreateOK) GetPayload() []*models.AvailablePrefix {
	return o.Payload
}

func (o *IpamPrefixesAvailablePrefixesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
