package dbkv

import (
	"errors"
	"strconv"
)

type DBKVModel struct {
	Id          int64  `json:"id" gorm:"primaryKey"`
	Key         string `json:"key" gorm:"index;unique"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

func (model *DBKVModel) ToString() string {
	return model.Value
}

func (model *DBKVModel) ToBool() (bool, error) {
	switch model.Value {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, errors.New("format err, str value:" + model.Value)
	}
}

func (model *DBKVModel) ToInt() (int, error) {
	return strconv.Atoi(model.Value)
}

func (model *DBKVModel) ToInt32() (int32, error) {
	result, err := strconv.ParseInt(model.Value, 10, 32)
	if err != nil {
		return 0, err
	} else {
		return int32(result), nil
	}
}

func (model *DBKVModel) ToInt64() (int64, error) {
	return strconv.ParseInt(model.Value, 10, 64)
}

func (model *DBKVModel) ToUInt64() (uint64, error) {
	return strconv.ParseUint(model.Value, 10, 64)
}

func (model *DBKVModel) ToFloat64() (float64, error) {
	return strconv.ParseFloat(model.Value, 64)
}

func (model *DBKVModel) ToFloat32() (float32, error) {
	result, err := strconv.ParseFloat(model.Value, 32)
	if err != nil {
		return 0, err
	} else {
		return float32(result), nil
	}
}
