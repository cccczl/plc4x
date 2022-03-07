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

type GAVState uint8

type IGAVState interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	GAVState_DOES_NOT_EXIST GAVState = 0
	GAVState_ON             GAVState = 1
	GAVState_OFF            GAVState = 2
	GAVState_ERROR          GAVState = 3
)

var GAVStateValues []GAVState

func init() {
	_ = errors.New
	GAVStateValues = []GAVState{
		GAVState_DOES_NOT_EXIST,
		GAVState_ON,
		GAVState_OFF,
		GAVState_ERROR,
	}
}

func GAVStateByValue(value uint8) GAVState {
	switch value {
	case 0:
		return GAVState_DOES_NOT_EXIST
	case 1:
		return GAVState_ON
	case 2:
		return GAVState_OFF
	case 3:
		return GAVState_ERROR
	}
	return 0
}

func GAVStateByName(value string) GAVState {
	switch value {
	case "DOES_NOT_EXIST":
		return GAVState_DOES_NOT_EXIST
	case "ON":
		return GAVState_ON
	case "OFF":
		return GAVState_OFF
	case "ERROR":
		return GAVState_ERROR
	}
	return 0
}

func GAVStateKnows(value uint8) bool {
	for _, typeValue := range GAVStateValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastGAVState(structType interface{}) GAVState {
	castFunc := func(typ interface{}) GAVState {
		if sGAVState, ok := typ.(GAVState); ok {
			return sGAVState
		}
		return 0
	}
	return castFunc(structType)
}

func (m GAVState) GetLengthInBits() uint16 {
	return 2
}

func (m GAVState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func GAVStateParse(readBuffer utils.ReadBuffer) (GAVState, error) {
	val, err := readBuffer.ReadUint8("GAVState", 2)
	if err != nil {
		return 0, nil
	}
	return GAVStateByValue(val), nil
}

func (e GAVState) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("GAVState", 2, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e GAVState) name() string {
	switch e {
	case GAVState_DOES_NOT_EXIST:
		return "DOES_NOT_EXIST"
	case GAVState_ON:
		return "ON"
	case GAVState_OFF:
		return "OFF"
	case GAVState_ERROR:
		return "ERROR"
	}
	return ""
}

func (e GAVState) String() string {
	return e.name()
}