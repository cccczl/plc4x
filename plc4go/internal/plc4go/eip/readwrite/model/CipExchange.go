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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// Constant values.
const CipExchange_ITEMCOUNT uint16 = 0x02
const CipExchange_NULLPTR uint32 = 0x0
const CipExchange_UNCONNECTEDDATA uint16 = 0x00B2

// The data-structure of this message
type CipExchange struct {
	Service *CipService
}

// The corresponding interface
type ICipExchange interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

func NewCipExchange(service *CipService) *CipExchange {
	return &CipExchange{Service: service}
}

func CastCipExchange(structType interface{}) *CipExchange {
	castFunc := func(typ interface{}) *CipExchange {
		if casted, ok := typ.(CipExchange); ok {
			return &casted
		}
		if casted, ok := typ.(*CipExchange); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *CipExchange) GetTypeName() string {
	return "CipExchange"
}

func (m *CipExchange) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *CipExchange) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Const Field (itemCount)
	lengthInBits += 16

	// Const Field (nullPtr)
	lengthInBits += 32

	// Const Field (UnconnectedData)
	lengthInBits += 16

	// Implicit Field (size)
	lengthInBits += 16

	// Simple field (service)
	lengthInBits += m.Service.LengthInBits()

	return lengthInBits
}

func (m *CipExchange) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func CipExchangeParse(io utils.ReadBuffer, exchangeLen uint16) (*CipExchange, error) {
	if pullErr := io.PullContext("CipExchange"); pullErr != nil {
		return nil, pullErr
	}

	// Const Field (itemCount)
	itemCount, _itemCountErr := io.ReadUint16("itemCount", 16)
	if _itemCountErr != nil {
		return nil, errors.Wrap(_itemCountErr, "Error parsing 'itemCount' field")
	}
	if itemCount != CipExchange_ITEMCOUNT {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CipExchange_ITEMCOUNT) + " but got " + fmt.Sprintf("%d", itemCount))
	}

	// Const Field (nullPtr)
	nullPtr, _nullPtrErr := io.ReadUint32("nullPtr", 32)
	if _nullPtrErr != nil {
		return nil, errors.Wrap(_nullPtrErr, "Error parsing 'nullPtr' field")
	}
	if nullPtr != CipExchange_NULLPTR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CipExchange_NULLPTR) + " but got " + fmt.Sprintf("%d", nullPtr))
	}

	// Const Field (UnconnectedData)
	UnconnectedData, _UnconnectedDataErr := io.ReadUint16("UnconnectedData", 16)
	if _UnconnectedDataErr != nil {
		return nil, errors.Wrap(_UnconnectedDataErr, "Error parsing 'UnconnectedData' field")
	}
	if UnconnectedData != CipExchange_UNCONNECTEDDATA {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CipExchange_UNCONNECTEDDATA) + " but got " + fmt.Sprintf("%d", UnconnectedData))
	}

	// Implicit Field (size) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	size, _sizeErr := io.ReadUint16("size", 16)
	_ = size
	if _sizeErr != nil {
		return nil, errors.Wrap(_sizeErr, "Error parsing 'size' field")
	}

	if pullErr := io.PullContext("service"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (service)
	service, _serviceErr := CipServiceParse(io, uint16(exchangeLen)-uint16(uint16(10)))
	if _serviceErr != nil {
		return nil, errors.Wrap(_serviceErr, "Error parsing 'service' field")
	}
	if closeErr := io.CloseContext("service"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := io.CloseContext("CipExchange"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewCipExchange(service), nil
}

func (m *CipExchange) Serialize(io utils.WriteBuffer) error {
	if pushErr := io.PushContext("CipExchange"); pushErr != nil {
		return pushErr
	}

	// Const Field (itemCount)
	_itemCountErr := io.WriteUint16("itemCount", 16, 0x02)
	if _itemCountErr != nil {
		return errors.Wrap(_itemCountErr, "Error serializing 'itemCount' field")
	}

	// Const Field (nullPtr)
	_nullPtrErr := io.WriteUint32("nullPtr", 32, 0x0)
	if _nullPtrErr != nil {
		return errors.Wrap(_nullPtrErr, "Error serializing 'nullPtr' field")
	}

	// Const Field (UnconnectedData)
	_UnconnectedDataErr := io.WriteUint16("UnconnectedData", 16, 0x00B2)
	if _UnconnectedDataErr != nil {
		return errors.Wrap(_UnconnectedDataErr, "Error serializing 'UnconnectedData' field")
	}

	// Implicit Field (size) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	size := uint16(uint16(uint16(uint16(m.LengthInBytes()))-uint16(uint16(8))) - uint16(uint16(2)))
	_sizeErr := io.WriteUint16("size", 16, (size))
	if _sizeErr != nil {
		return errors.Wrap(_sizeErr, "Error serializing 'size' field")
	}

	// Simple Field (service)
	if pushErr := io.PushContext("service"); pushErr != nil {
		return pushErr
	}
	_serviceErr := m.Service.Serialize(io)
	if popErr := io.PopContext("service"); popErr != nil {
		return popErr
	}
	if _serviceErr != nil {
		return errors.Wrap(_serviceErr, "Error serializing 'service' field")
	}

	if popErr := io.PopContext("CipExchange"); popErr != nil {
		return popErr
	}
	return nil
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *CipExchange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "service":
				var dt *CipService
				if err := d.DecodeElement(&dt, &tok); err != nil {
					if err == io.EOF {
						continue
					}
					return err
				}
				m.Service = dt
			}
		}
	}
}

// Deprecated: the utils.WriteBufferReadBased should be used instead
func (m *CipExchange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := "org.apache.plc4x.java.eip.readwrite.CipExchange"
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Service, xml.StartElement{Name: xml.Name{Local: "service"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m CipExchange) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m CipExchange) Box(name string, width int) utils.AsciiBox {
	boxName := "CipExchange"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Const Field (itemCount)
	boxes = append(boxes, utils.BoxAnything("ItemCount", uint16(0x02), -1))
	// Const Field (nullPtr)
	boxes = append(boxes, utils.BoxAnything("NullPtr", uint32(0x0), -1))
	// Const Field (UnconnectedData)
	boxes = append(boxes, utils.BoxAnything("UnconnectedData", uint16(0x00B2), -1))
	// Implicit Field (size)
	size := uint16(uint16(uint16(uint16(m.LengthInBytes()))-uint16(uint16(8))) - uint16(uint16(2)))
	// uint16 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("Size", size, -1))
	// Complex field (case complex)
	boxes = append(boxes, m.Service.Box("service", width-2))
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}