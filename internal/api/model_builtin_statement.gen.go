/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// BuiltinStatement Declares a builtin identifier and its type
type BuiltinStatement struct {
	// Type of AST node
	Type *string     `json:"type,omitempty"`
	Id   *Identifier `json:"id,omitempty"`
}

// NewBuiltinStatement instantiates a new BuiltinStatement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuiltinStatement() *BuiltinStatement {
	this := BuiltinStatement{}
	return &this
}

// NewBuiltinStatementWithDefaults instantiates a new BuiltinStatement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuiltinStatementWithDefaults() *BuiltinStatement {
	this := BuiltinStatement{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BuiltinStatement) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuiltinStatement) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BuiltinStatement) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BuiltinStatement) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BuiltinStatement) GetId() Identifier {
	if o == nil || o.Id == nil {
		var ret Identifier
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuiltinStatement) GetIdOk() (*Identifier, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BuiltinStatement) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given Identifier and assigns it to the Id field.
func (o *BuiltinStatement) SetId(v Identifier) {
	o.Id = &v
}

func (o BuiltinStatement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableBuiltinStatement struct {
	value *BuiltinStatement
	isSet bool
}

func (v NullableBuiltinStatement) Get() *BuiltinStatement {
	return v.value
}

func (v *NullableBuiltinStatement) Set(val *BuiltinStatement) {
	v.value = val
	v.isSet = true
}

func (v NullableBuiltinStatement) IsSet() bool {
	return v.isSet
}

func (v *NullableBuiltinStatement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuiltinStatement(val *BuiltinStatement) *NullableBuiltinStatement {
	return &NullableBuiltinStatement{value: val, isSet: true}
}

func (v NullableBuiltinStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuiltinStatement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}