/*
Model Registry REST API

REST API for Model Registry to create and manage ML model metadata

API version: v1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model_registry_client

import (
	"encoding/json"
	"fmt"
)

// OrderByField Supported fields for ordering result entities.
type OrderByField string

// List of OrderByField
const (
	ORDERBYFIELD_CREATE_TIME OrderByField = "CREATE_TIME"
	ORDERBYFIELD_LAST_UPDATE_TIME OrderByField = "LAST_UPDATE_TIME"
	ORDERBYFIELD_ID OrderByField = "ID"
)

// All allowed values of OrderByField enum
var AllowedOrderByFieldEnumValues = []OrderByField{
	"CREATE_TIME",
	"LAST_UPDATE_TIME",
	"ID",
}

func (v *OrderByField) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderByField(value)
	for _, existing := range AllowedOrderByFieldEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderByField", value)
}

// NewOrderByFieldFromValue returns a pointer to a valid OrderByField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderByFieldFromValue(v string) (*OrderByField, error) {
	ev := OrderByField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderByField: valid values are %v", v, AllowedOrderByFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderByField) IsValid() bool {
	for _, existing := range AllowedOrderByFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderByField value
func (v OrderByField) Ptr() *OrderByField {
	return &v
}

type NullableOrderByField struct {
	value *OrderByField
	isSet bool
}

func (v NullableOrderByField) Get() *OrderByField {
	return v.value
}

func (v *NullableOrderByField) Set(val *OrderByField) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderByField) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderByField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderByField(val *OrderByField) *NullableOrderByField {
	return &NullableOrderByField{value: val, isSet: true}
}

func (v NullableOrderByField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderByField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

