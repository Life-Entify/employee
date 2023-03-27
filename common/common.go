package employee

import (
	"encoding/json"
)

func ToJSONStruct(value interface{}, out any) error {
	valueByte, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = json.Unmarshal(valueByte, out)
	if err != nil {
		return err
	}
	return nil
}
