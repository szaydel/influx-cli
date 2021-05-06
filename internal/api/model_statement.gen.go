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
	"errors"
)

// Statement - struct for Statement
type Statement struct {
	BadStatement        *BadStatement
	BuiltinStatement    *BuiltinStatement
	ExpressionStatement *ExpressionStatement
	MemberAssignment    *MemberAssignment
	OptionStatement     *OptionStatement
	ReturnStatement     *ReturnStatement
	TestStatement       *TestStatement
	VariableAssignment  *VariableAssignment
}

// BadStatementAsStatement is a convenience function that returns BadStatement wrapped in Statement
func BadStatementAsStatement(v *BadStatement) Statement {
	return Statement{BadStatement: v}
}

// BuiltinStatementAsStatement is a convenience function that returns BuiltinStatement wrapped in Statement
func BuiltinStatementAsStatement(v *BuiltinStatement) Statement {
	return Statement{BuiltinStatement: v}
}

// ExpressionStatementAsStatement is a convenience function that returns ExpressionStatement wrapped in Statement
func ExpressionStatementAsStatement(v *ExpressionStatement) Statement {
	return Statement{ExpressionStatement: v}
}

// MemberAssignmentAsStatement is a convenience function that returns MemberAssignment wrapped in Statement
func MemberAssignmentAsStatement(v *MemberAssignment) Statement {
	return Statement{MemberAssignment: v}
}

// OptionStatementAsStatement is a convenience function that returns OptionStatement wrapped in Statement
func OptionStatementAsStatement(v *OptionStatement) Statement {
	return Statement{OptionStatement: v}
}

// ReturnStatementAsStatement is a convenience function that returns ReturnStatement wrapped in Statement
func ReturnStatementAsStatement(v *ReturnStatement) Statement {
	return Statement{ReturnStatement: v}
}

// TestStatementAsStatement is a convenience function that returns TestStatement wrapped in Statement
func TestStatementAsStatement(v *TestStatement) Statement {
	return Statement{TestStatement: v}
}

// VariableAssignmentAsStatement is a convenience function that returns VariableAssignment wrapped in Statement
func VariableAssignmentAsStatement(v *VariableAssignment) Statement {
	return Statement{VariableAssignment: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Statement) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BadStatement
	err = json.Unmarshal(data, &dst.BadStatement)
	if err == nil {
		jsonBadStatement, _ := json.Marshal(dst.BadStatement)
		if string(jsonBadStatement) == "{}" { // empty struct
			dst.BadStatement = nil
		} else {
			match++
		}
	} else {
		dst.BadStatement = nil
	}

	// try to unmarshal data into BuiltinStatement
	err = json.Unmarshal(data, &dst.BuiltinStatement)
	if err == nil {
		jsonBuiltinStatement, _ := json.Marshal(dst.BuiltinStatement)
		if string(jsonBuiltinStatement) == "{}" { // empty struct
			dst.BuiltinStatement = nil
		} else {
			match++
		}
	} else {
		dst.BuiltinStatement = nil
	}

	// try to unmarshal data into ExpressionStatement
	err = json.Unmarshal(data, &dst.ExpressionStatement)
	if err == nil {
		jsonExpressionStatement, _ := json.Marshal(dst.ExpressionStatement)
		if string(jsonExpressionStatement) == "{}" { // empty struct
			dst.ExpressionStatement = nil
		} else {
			match++
		}
	} else {
		dst.ExpressionStatement = nil
	}

	// try to unmarshal data into MemberAssignment
	err = json.Unmarshal(data, &dst.MemberAssignment)
	if err == nil {
		jsonMemberAssignment, _ := json.Marshal(dst.MemberAssignment)
		if string(jsonMemberAssignment) == "{}" { // empty struct
			dst.MemberAssignment = nil
		} else {
			match++
		}
	} else {
		dst.MemberAssignment = nil
	}

	// try to unmarshal data into OptionStatement
	err = json.Unmarshal(data, &dst.OptionStatement)
	if err == nil {
		jsonOptionStatement, _ := json.Marshal(dst.OptionStatement)
		if string(jsonOptionStatement) == "{}" { // empty struct
			dst.OptionStatement = nil
		} else {
			match++
		}
	} else {
		dst.OptionStatement = nil
	}

	// try to unmarshal data into ReturnStatement
	err = json.Unmarshal(data, &dst.ReturnStatement)
	if err == nil {
		jsonReturnStatement, _ := json.Marshal(dst.ReturnStatement)
		if string(jsonReturnStatement) == "{}" { // empty struct
			dst.ReturnStatement = nil
		} else {
			match++
		}
	} else {
		dst.ReturnStatement = nil
	}

	// try to unmarshal data into TestStatement
	err = json.Unmarshal(data, &dst.TestStatement)
	if err == nil {
		jsonTestStatement, _ := json.Marshal(dst.TestStatement)
		if string(jsonTestStatement) == "{}" { // empty struct
			dst.TestStatement = nil
		} else {
			match++
		}
	} else {
		dst.TestStatement = nil
	}

	// try to unmarshal data into VariableAssignment
	err = json.Unmarshal(data, &dst.VariableAssignment)
	if err == nil {
		jsonVariableAssignment, _ := json.Marshal(dst.VariableAssignment)
		if string(jsonVariableAssignment) == "{}" { // empty struct
			dst.VariableAssignment = nil
		} else {
			match++
		}
	} else {
		dst.VariableAssignment = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BadStatement = nil
		dst.BuiltinStatement = nil
		dst.ExpressionStatement = nil
		dst.MemberAssignment = nil
		dst.OptionStatement = nil
		dst.ReturnStatement = nil
		dst.TestStatement = nil
		dst.VariableAssignment = nil

		return errors.New("data matches more than one schema in oneOf(Statement)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return errors.New("data failed to match schemas in oneOf(Statement)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Statement) MarshalJSON() ([]byte, error) {
	if src.BadStatement != nil {
		return json.Marshal(&src.BadStatement)
	}

	if src.BuiltinStatement != nil {
		return json.Marshal(&src.BuiltinStatement)
	}

	if src.ExpressionStatement != nil {
		return json.Marshal(&src.ExpressionStatement)
	}

	if src.MemberAssignment != nil {
		return json.Marshal(&src.MemberAssignment)
	}

	if src.OptionStatement != nil {
		return json.Marshal(&src.OptionStatement)
	}

	if src.ReturnStatement != nil {
		return json.Marshal(&src.ReturnStatement)
	}

	if src.TestStatement != nil {
		return json.Marshal(&src.TestStatement)
	}

	if src.VariableAssignment != nil {
		return json.Marshal(&src.VariableAssignment)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Statement) GetActualInstance() interface{} {
	if obj.BadStatement != nil {
		return obj.BadStatement
	}

	if obj.BuiltinStatement != nil {
		return obj.BuiltinStatement
	}

	if obj.ExpressionStatement != nil {
		return obj.ExpressionStatement
	}

	if obj.MemberAssignment != nil {
		return obj.MemberAssignment
	}

	if obj.OptionStatement != nil {
		return obj.OptionStatement
	}

	if obj.ReturnStatement != nil {
		return obj.ReturnStatement
	}

	if obj.TestStatement != nil {
		return obj.TestStatement
	}

	if obj.VariableAssignment != nil {
		return obj.VariableAssignment
	}

	// all schemas are nil
	return nil
}

type NullableStatement struct {
	value *Statement
	isSet bool
}

func (v NullableStatement) Get() *Statement {
	return v.value
}

func (v *NullableStatement) Set(val *Statement) {
	v.value = val
	v.isSet = true
}

func (v NullableStatement) IsSet() bool {
	return v.isSet
}

func (v *NullableStatement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatement(val *Statement) *NullableStatement {
	return &NullableStatement{value: val, isSet: true}
}

func (v NullableStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}