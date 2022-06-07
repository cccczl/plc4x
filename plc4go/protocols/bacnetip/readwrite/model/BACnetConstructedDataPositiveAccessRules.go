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

// BACnetConstructedDataPositiveAccessRules is the data-structure of this message
type BACnetConstructedDataPositiveAccessRules struct {
	*BACnetConstructedData
	PositiveAccessRules []*BACnetAccessRule

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataPositiveAccessRules is the corresponding interface of BACnetConstructedDataPositiveAccessRules
type IBACnetConstructedDataPositiveAccessRules interface {
	IBACnetConstructedData
	// GetPositiveAccessRules returns PositiveAccessRules (property field)
	GetPositiveAccessRules() []*BACnetAccessRule
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

func (m *BACnetConstructedDataPositiveAccessRules) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataPositiveAccessRules) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_POSITIVE_ACCESS_RULES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataPositiveAccessRules) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataPositiveAccessRules) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataPositiveAccessRules) GetPositiveAccessRules() []*BACnetAccessRule {
	return m.PositiveAccessRules
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataPositiveAccessRules factory function for BACnetConstructedDataPositiveAccessRules
func NewBACnetConstructedDataPositiveAccessRules(positiveAccessRules []*BACnetAccessRule, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataPositiveAccessRules {
	_result := &BACnetConstructedDataPositiveAccessRules{
		PositiveAccessRules:   positiveAccessRules,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataPositiveAccessRules(structType interface{}) *BACnetConstructedDataPositiveAccessRules {
	if casted, ok := structType.(BACnetConstructedDataPositiveAccessRules); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPositiveAccessRules); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataPositiveAccessRules(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataPositiveAccessRules(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataPositiveAccessRules) GetTypeName() string {
	return "BACnetConstructedDataPositiveAccessRules"
}

func (m *BACnetConstructedDataPositiveAccessRules) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataPositiveAccessRules) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.PositiveAccessRules) > 0 {
		for _, element := range m.PositiveAccessRules {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataPositiveAccessRules) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataPositiveAccessRulesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataPositiveAccessRules, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPositiveAccessRules"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (positiveAccessRules)
	if pullErr := readBuffer.PullContext("positiveAccessRules", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	positiveAccessRules := make([]*BACnetAccessRule, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetAccessRuleParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'positiveAccessRules' field")
			}
			positiveAccessRules = append(positiveAccessRules, CastBACnetAccessRule(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("positiveAccessRules", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPositiveAccessRules"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataPositiveAccessRules{
		PositiveAccessRules:   positiveAccessRules,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataPositiveAccessRules) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPositiveAccessRules"); pushErr != nil {
			return pushErr
		}

		// Array Field (positiveAccessRules)
		if m.PositiveAccessRules != nil {
			if pushErr := writeBuffer.PushContext("positiveAccessRules", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.PositiveAccessRules {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'positiveAccessRules' field")
				}
			}
			if popErr := writeBuffer.PopContext("positiveAccessRules", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPositiveAccessRules"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataPositiveAccessRules) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}