//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type FirmataCommandSetDigitalPinValue struct {
	Pin    uint8
	On     bool
	Parent *FirmataCommand
}

// The corresponding interface
type IFirmataCommandSetDigitalPinValue interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *FirmataCommandSetDigitalPinValue) CommandCode() uint8 {
	return 0x5
}

func (m *FirmataCommandSetDigitalPinValue) InitializeParent(parent *FirmataCommand) {
}

func NewFirmataCommandSetDigitalPinValue(pin uint8, on bool) *FirmataCommand {
	child := &FirmataCommandSetDigitalPinValue{
		Pin:    pin,
		On:     on,
		Parent: NewFirmataCommand(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastFirmataCommandSetDigitalPinValue(structType interface{}) *FirmataCommandSetDigitalPinValue {
	castFunc := func(typ interface{}) *FirmataCommandSetDigitalPinValue {
		if casted, ok := typ.(FirmataCommandSetDigitalPinValue); ok {
			return &casted
		}
		if casted, ok := typ.(*FirmataCommandSetDigitalPinValue); ok {
			return casted
		}
		if casted, ok := typ.(FirmataCommand); ok {
			return CastFirmataCommandSetDigitalPinValue(casted.Child)
		}
		if casted, ok := typ.(*FirmataCommand); ok {
			return CastFirmataCommandSetDigitalPinValue(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *FirmataCommandSetDigitalPinValue) GetTypeName() string {
	return "FirmataCommandSetDigitalPinValue"
}

func (m *FirmataCommandSetDigitalPinValue) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *FirmataCommandSetDigitalPinValue) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (pin)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (on)
	lengthInBits += 1

	return lengthInBits
}

func (m *FirmataCommandSetDigitalPinValue) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func FirmataCommandSetDigitalPinValueParse(io utils.ReadBuffer) (*FirmataCommand, error) {
	if pullErr := io.PullContext("FirmataCommandSetDigitalPinValue"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (pin)
	pin, _pinErr := io.ReadUint8("pin", 8)
	if _pinErr != nil {
		return nil, errors.Wrap(_pinErr, "Error parsing 'pin' field")
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := io.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (on)
	on, _onErr := io.ReadBit("on")
	if _onErr != nil {
		return nil, errors.Wrap(_onErr, "Error parsing 'on' field")
	}

	if closeErr := io.CloseContext("FirmataCommandSetDigitalPinValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &FirmataCommandSetDigitalPinValue{
		Pin:    pin,
		On:     on,
		Parent: &FirmataCommand{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *FirmataCommandSetDigitalPinValue) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("FirmataCommandSetDigitalPinValue"); pushErr != nil {
			return pushErr
		}

		// Simple Field (pin)
		pin := uint8(m.Pin)
		_pinErr := io.WriteUint8("pin", 8, (pin))
		if _pinErr != nil {
			return errors.Wrap(_pinErr, "Error serializing 'pin' field")
		}

		// Reserved Field (reserved)
		{
			_err := io.WriteUint8("reserved", 7, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (on)
		on := bool(m.On)
		_onErr := io.WriteBit("on", (on))
		if _onErr != nil {
			return errors.Wrap(_onErr, "Error serializing 'on' field")
		}

		if popErr := io.PopContext("FirmataCommandSetDigitalPinValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *FirmataCommandSetDigitalPinValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "pin":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Pin = data
			case "on":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.On = data
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

// Deprecated: the utils.WriteBufferReadBased should be used instead
func (m *FirmataCommandSetDigitalPinValue) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Pin, xml.StartElement{Name: xml.Name{Local: "pin"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.On, xml.StartElement{Name: xml.Name{Local: "on"}}); err != nil {
		return err
	}
	return nil
}

func (m FirmataCommandSetDigitalPinValue) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m FirmataCommandSetDigitalPinValue) Box(name string, width int) utils.AsciiBox {
	boxName := "FirmataCommandSetDigitalPinValue"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("Pin", m.Pin, -1))
		// Reserved Field (reserved)
		// reserved field can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("reserved", uint8(0x00), -1))
		// Simple field (case simple)
		// bool can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("On", m.On, -1))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}