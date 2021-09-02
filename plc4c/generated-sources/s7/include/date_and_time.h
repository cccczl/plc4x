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

#ifndef PLC4C_S7_READ_WRITE_DATE_AND_TIME_H_
#define PLC4C_S7_READ_WRITE_DATE_AND_TIME_H_

#include <stdbool.h>
#include <stdint.h>
#include <plc4c/driver_s7_static.h>
#include <plc4c/spi/read_buffer.h>
#include <plc4c/spi/write_buffer.h>
#include <plc4c/utils/list.h>

// Code generated by code-generation. DO NOT EDIT.

#ifdef __cplusplus
extern "C" {
#endif


struct plc4c_s7_read_write_date_and_time {
  /* Properties */
  uint8_t year;
  uint8_t month;
  uint8_t day;
  uint8_t hour;
  uint8_t minutes;
  uint8_t seconds;
  uint16_t msec : 12;
  uint8_t dow : 4;
};
typedef struct plc4c_s7_read_write_date_and_time plc4c_s7_read_write_date_and_time;

// Create an empty NULL-struct
plc4c_s7_read_write_date_and_time plc4c_s7_read_write_date_and_time_null();

plc4c_return_code plc4c_s7_read_write_date_and_time_parse(plc4c_spi_read_buffer* readBuffer, plc4c_s7_read_write_date_and_time** message);

plc4c_return_code plc4c_s7_read_write_date_and_time_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_date_and_time* message);

uint16_t plc4c_s7_read_write_date_and_time_length_in_bytes(plc4c_s7_read_write_date_and_time* message);

uint16_t plc4c_s7_read_write_date_and_time_length_in_bits(plc4c_s7_read_write_date_and_time* message);

#ifdef __cplusplus
}
#endif
#endif  // PLC4C_S7_READ_WRITE_DATE_AND_TIME_H_