package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
)

//StructToMap ...
func StructToMap(data interface{}) (map[string]interface{}, error) {
	errMessage := "[structToMap()] > "
	m := make(map[string]interface{})

	j, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(errMessage + err.Error() + " :: [001]")
	}

	err = json.Unmarshal(j, &m)
	if err != nil {
		return nil, errors.New(errMessage + err.Error() + " :: [002]")
	}

	return m, nil
}

//ValueToMap ...
func ValueToMap(arg interface{}) (map[string]interface{}, error) {
	errMessage := "[ERR][Stage: Convert Data] >"
	m := make(map[string]interface{})

	j, err := json.Marshal(arg)
	if err != nil {

		return nil, fmt.Errorf("%+v %+v", errMessage, err.Error())
	}

	err = json.Unmarshal(j, &m)
	if err != nil {
		return nil, fmt.Errorf("%+v %+v", errMessage, err.Error())
	}

	return m, nil
}
