package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"reflect"
)

// UserSessionInfo ...
type UserSessionInfo struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// JSONRaw ...
type JSONRaw json.RawMessage

// Value ...
func (j JSONRaw) Value() (driver.Value, error) {
	byteArr := []byte(j)
	return driver.Value(byteArr), nil
}

// Scan ...
func (j *JSONRaw) Scan(src interface{}) error {
	asBytes, ok := src.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []bytes"))
	}
	err := json.Unmarshal(asBytes, &j)
	if err != nil {
		return error(errors.New("Scan could not unmarshal to []string"))
	}
	return nil
}

// MarshalJSON ...
func (j *JSONRaw) MarshalJSON() ([]byte, error) {
	return *j, nil
}

// UnmarshalJSON ...
func (j *JSONRaw) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*j = append((*j)[0:0], data...)
	return nil
}

// ConvertToInt64 ...
func ConvertToInt64(number interface{}) int64 {
	if reflect.TypeOf(number).String() == "int" {
		return int64(number.(int))
	}
	return number.(int64)
}

func UintSliceToJsonString(slice []uint) string {
	sliceJson, err := json.Marshal(slice)
	if err != nil {
		panic(err)
	}
	return string(sliceJson)
}

func JsonStringToUintSlice(str string) (slice []uint) {
	byteSlice := []byte(str)
	if err := json.Unmarshal(byteSlice, &slice); err != nil {
		panic(err)
	}
	return slice
}

func StringUintMapToJsonString(dict map[string]uint) string {
	jsonMap, err := json.Marshal(dict)
	if err != nil {
		panic(err)
	}
	return string(jsonMap)
}

func JsonStringToStringUintMap(strMap string) (mp map[string]uint) {
	byteSlice := []byte(strMap)
	if err := json.Unmarshal(byteSlice, &mp); err != nil {
		panic(err)
	}
	return mp
}
