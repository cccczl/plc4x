/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataReliability is the data-structure of this message
type BACnetConstructedDataReliability struct {
	*BACnetConstructedData
	Reliability *BACnetConstructedDataReliabilityEntry

	// Arguments.
	TagNumber                  uint8
	PropertyIdentifierArgument BACnetContextTagPropertyIdentifier
}

// IBACnetConstructedDataReliability is the corresponding interface of BACnetConstructedDataReliability
type IBACnetConstructedDataReliability interface {
	IBACnetConstructedData
	// GetReliability returns Reliability (property field)
	GetReliability() *BACnetConstructedDataReliabilityEntry
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataReliability) GetObjectType() BACnetObjectType {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataReliability) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataReliability) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataReliability) GetReliability() *BACnetConstructedDataReliabilityEntry {
	return m.Reliability
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataReliability factory function for BACnetConstructedDataReliability
func NewBACnetConstructedDataReliability(reliability *BACnetConstructedDataReliabilityEntry, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8, propertyIdentifierArgument BACnetContextTagPropertyIdentifier) *BACnetConstructedDataReliability {
	_result := &BACnetConstructedDataReliability{
		Reliability:           reliability,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber, propertyIdentifierArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataReliability(structType interface{}) *BACnetConstructedDataReliability {
	if casted, ok := structType.(BACnetConstructedDataReliability); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataReliability); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataReliability(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataReliability(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataReliability) GetTypeName() string {
	return "BACnetConstructedDataReliability"
}

func (m *BACnetConstructedDataReliability) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataReliability) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (reliability)
	lengthInBits += m.Reliability.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataReliability) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataReliabilityParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectType BACnetObjectType, propertyIdentifierArgument *BACnetContextTagPropertyIdentifier) (*BACnetConstructedDataReliability, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataReliability"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (reliability)
	if pullErr := readBuffer.PullContext("reliability"); pullErr != nil {
		return nil, pullErr
	}
	_reliability, _reliabilityErr := BACnetConstructedDataReliabilityEntryParse(readBuffer)
	if _reliabilityErr != nil {
		return nil, errors.Wrap(_reliabilityErr, "Error parsing 'reliability' field")
	}
	reliability := CastBACnetConstructedDataReliabilityEntry(_reliability)
	if closeErr := readBuffer.CloseContext("reliability"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataReliability"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataReliability{
		Reliability:           CastBACnetConstructedDataReliabilityEntry(reliability),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataReliability) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataReliability"); pushErr != nil {
			return pushErr
		}

		// Simple Field (reliability)
		if pushErr := writeBuffer.PushContext("reliability"); pushErr != nil {
			return pushErr
		}
		_reliabilityErr := m.Reliability.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("reliability"); popErr != nil {
			return popErr
		}
		if _reliabilityErr != nil {
			return errors.Wrap(_reliabilityErr, "Error serializing 'reliability' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataReliability"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataReliability) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}