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
type FirmataCommandSysex struct {
	Command *SysexCommand
	Parent  *FirmataCommand
}

// The corresponding interface
type IFirmataCommandSysex interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *FirmataCommandSysex) CommandCode() uint8 {
	return 0x0
}

func (m *FirmataCommandSysex) InitializeParent(parent *FirmataCommand) {
}

func NewFirmataCommandSysex(command *SysexCommand) *FirmataCommand {
	child := &FirmataCommandSysex{
		Command: command,
		Parent:  NewFirmataCommand(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastFirmataCommandSysex(structType interface{}) *FirmataCommandSysex {
	castFunc := func(typ interface{}) *FirmataCommandSysex {
		if casted, ok := typ.(FirmataCommandSysex); ok {
			return &casted
		}
		if casted, ok := typ.(*FirmataCommandSysex); ok {
			return casted
		}
		if casted, ok := typ.(FirmataCommand); ok {
			return CastFirmataCommandSysex(casted.Child)
		}
		if casted, ok := typ.(*FirmataCommand); ok {
			return CastFirmataCommandSysex(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *FirmataCommandSysex) GetTypeName() string {
	return "FirmataCommandSysex"
}

func (m *FirmataCommandSysex) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *FirmataCommandSysex) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (command)
	lengthInBits += m.Command.LengthInBits()

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *FirmataCommandSysex) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func FirmataCommandSysexParse(io utils.ReadBuffer, response bool) (*FirmataCommand, error) {
	if pullErr := io.PullContext("FirmataCommandSysex"); pullErr != nil {
		return nil, pullErr
	}

	if pullErr := io.PullContext("command"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (command)
	command, _commandErr := SysexCommandParse(io, response)
	if _commandErr != nil {
		return nil, errors.Wrap(_commandErr, "Error parsing 'command' field")
	}
	if closeErr := io.CloseContext("command"); closeErr != nil {
		return nil, closeErr
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := io.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0xF7) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0xF7),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	if closeErr := io.CloseContext("FirmataCommandSysex"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &FirmataCommandSysex{
		Command: command,
		Parent:  &FirmataCommand{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *FirmataCommandSysex) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("FirmataCommandSysex"); pushErr != nil {
			return pushErr
		}

		// Simple Field (command)
		if pushErr := io.PushContext("command"); pushErr != nil {
			return pushErr
		}
		_commandErr := m.Command.Serialize(io)
		if popErr := io.PopContext("command"); popErr != nil {
			return popErr
		}
		if _commandErr != nil {
			return errors.Wrap(_commandErr, "Error serializing 'command' field")
		}

		// Reserved Field (reserved)
		{
			_err := io.WriteUint8("reserved", 8, uint8(0xF7))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		if popErr := io.PopContext("FirmataCommandSysex"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *FirmataCommandSysex) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "command":
				var dt *SysexCommand
				if err := d.DecodeElement(&dt, &tok); err != nil {
					if err == io.EOF {
						continue
					}
					return err
				}
				m.Command = dt
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
func (m *FirmataCommandSysex) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Command, xml.StartElement{Name: xml.Name{Local: "command"}}); err != nil {
		return err
	}
	return nil
}

func (m FirmataCommandSysex) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m FirmataCommandSysex) Box(name string, width int) utils.AsciiBox {
	boxName := "FirmataCommandSysex"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Complex field (case complex)
		boxes = append(boxes, m.Command.Box("command", width-2))
		// Reserved Field (reserved)
		// reserved field can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("reserved", uint8(0xF7), -1))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}