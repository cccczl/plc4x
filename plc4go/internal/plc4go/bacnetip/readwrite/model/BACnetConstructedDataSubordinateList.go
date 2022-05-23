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

// BACnetConstructedDataSubordinateList is the data-structure of this message
type BACnetConstructedDataSubordinateList struct {
	*BACnetConstructedData
	SubordinateList []*BACnetDeviceObjectReference

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataSubordinateList is the corresponding interface of BACnetConstructedDataSubordinateList
type IBACnetConstructedDataSubordinateList interface {
	IBACnetConstructedData
	// GetSubordinateList returns SubordinateList (property field)
	GetSubordinateList() []*BACnetDeviceObjectReference
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

func (m *BACnetConstructedDataSubordinateList) GetObjectType() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataSubordinateList) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SUBORDINATE_LIST
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataSubordinateList) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataSubordinateList) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataSubordinateList) GetSubordinateList() []*BACnetDeviceObjectReference {
	return m.SubordinateList
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSubordinateList factory function for BACnetConstructedDataSubordinateList
func NewBACnetConstructedDataSubordinateList(subordinateList []*BACnetDeviceObjectReference, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataSubordinateList {
	_result := &BACnetConstructedDataSubordinateList{
		SubordinateList:       subordinateList,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataSubordinateList(structType interface{}) *BACnetConstructedDataSubordinateList {
	if casted, ok := structType.(BACnetConstructedDataSubordinateList); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSubordinateList); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataSubordinateList(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataSubordinateList(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataSubordinateList) GetTypeName() string {
	return "BACnetConstructedDataSubordinateList"
}

func (m *BACnetConstructedDataSubordinateList) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataSubordinateList) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.SubordinateList) > 0 {
		for _, element := range m.SubordinateList {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataSubordinateList) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSubordinateListParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectType BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataSubordinateList, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSubordinateList"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (subordinateList)
	if pullErr := readBuffer.PullContext("subordinateList", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	subordinateList := make([]*BACnetDeviceObjectReference, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetDeviceObjectReferenceParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'subordinateList' field")
			}
			subordinateList = append(subordinateList, CastBACnetDeviceObjectReference(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("subordinateList", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSubordinateList"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataSubordinateList{
		SubordinateList:       subordinateList,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataSubordinateList) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSubordinateList"); pushErr != nil {
			return pushErr
		}

		// Array Field (subordinateList)
		if m.SubordinateList != nil {
			if pushErr := writeBuffer.PushContext("subordinateList", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.SubordinateList {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'subordinateList' field")
				}
			}
			if popErr := writeBuffer.PopContext("subordinateList", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSubordinateList"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataSubordinateList) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}