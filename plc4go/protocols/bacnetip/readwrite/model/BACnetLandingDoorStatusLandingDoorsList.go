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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetLandingDoorStatusLandingDoorsList is the data-structure of this message
type BACnetLandingDoorStatusLandingDoorsList struct {
	OpeningTag   *BACnetOpeningTag
	LandingDoors []*BACnetLandingDoorStatusLandingDoorsListEntry
	ClosingTag   *BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

// IBACnetLandingDoorStatusLandingDoorsList is the corresponding interface of BACnetLandingDoorStatusLandingDoorsList
type IBACnetLandingDoorStatusLandingDoorsList interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetLandingDoors returns LandingDoors (property field)
	GetLandingDoors() []*BACnetLandingDoorStatusLandingDoorsListEntry
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetLandingDoorStatusLandingDoorsList) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetLandingDoorStatusLandingDoorsList) GetLandingDoors() []*BACnetLandingDoorStatusLandingDoorsListEntry {
	return m.LandingDoors
}

func (m *BACnetLandingDoorStatusLandingDoorsList) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLandingDoorStatusLandingDoorsList factory function for BACnetLandingDoorStatusLandingDoorsList
func NewBACnetLandingDoorStatusLandingDoorsList(openingTag *BACnetOpeningTag, landingDoors []*BACnetLandingDoorStatusLandingDoorsListEntry, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetLandingDoorStatusLandingDoorsList {
	return &BACnetLandingDoorStatusLandingDoorsList{OpeningTag: openingTag, LandingDoors: landingDoors, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetLandingDoorStatusLandingDoorsList(structType interface{}) *BACnetLandingDoorStatusLandingDoorsList {
	if casted, ok := structType.(BACnetLandingDoorStatusLandingDoorsList); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetLandingDoorStatusLandingDoorsList); ok {
		return casted
	}
	return nil
}

func (m *BACnetLandingDoorStatusLandingDoorsList) GetTypeName() string {
	return "BACnetLandingDoorStatusLandingDoorsList"
}

func (m *BACnetLandingDoorStatusLandingDoorsList) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetLandingDoorStatusLandingDoorsList) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Array field
	if len(m.LandingDoors) > 0 {
		for _, element := range m.LandingDoors {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetLandingDoorStatusLandingDoorsList) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLandingDoorStatusLandingDoorsListParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetLandingDoorStatusLandingDoorsList, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLandingDoorStatusLandingDoorsList"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Array field (landingDoors)
	if pullErr := readBuffer.PullContext("landingDoors", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	landingDoors := make([]*BACnetLandingDoorStatusLandingDoorsListEntry, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetLandingDoorStatusLandingDoorsListEntryParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'landingDoors' field")
			}
			landingDoors = append(landingDoors, CastBACnetLandingDoorStatusLandingDoorsListEntry(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("landingDoors", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetLandingDoorStatusLandingDoorsList"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetLandingDoorStatusLandingDoorsList(openingTag, landingDoors, closingTag, tagNumber), nil
}

func (m *BACnetLandingDoorStatusLandingDoorsList) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetLandingDoorStatusLandingDoorsList"); pushErr != nil {
		return pushErr
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return pushErr
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return popErr
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Array Field (landingDoors)
	if m.LandingDoors != nil {
		if pushErr := writeBuffer.PushContext("landingDoors", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.LandingDoors {
			_elementErr := _element.Serialize(writeBuffer)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'landingDoors' field")
			}
		}
		if popErr := writeBuffer.PopContext("landingDoors", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return pushErr
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return popErr
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLandingDoorStatusLandingDoorsList"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetLandingDoorStatusLandingDoorsList) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}