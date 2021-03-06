/*
 * Swagger Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package petstore

import (
	"time"
)

type FormatTest struct {
	Integer int32 `json:"integer,omitempty" xml:"integer"`

	Int32_ int32 `json:"int32,omitempty" xml:"int32"`

	Int64_ int64 `json:"int64,omitempty" xml:"int64"`

	Number float32 `json:"number" xml:"number"`

	Float float32 `json:"float,omitempty" xml:"float"`

	Double float64 `json:"double,omitempty" xml:"double"`

	String_ string `json:"string,omitempty" xml:"string"`

	Byte_ string `json:"byte" xml:"byte"`

	Binary string `json:"binary,omitempty" xml:"binary"`

	Date string `json:"date" xml:"date"`

	DateTime time.Time `json:"dateTime,omitempty" xml:"dateTime"`

	Uuid string `json:"uuid,omitempty" xml:"uuid"`

	Password string `json:"password" xml:"password"`
}
