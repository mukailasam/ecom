package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonRequest struct {
	Data map[string]interface{}
}

/*
func JsonRequestDecoderTwo(r *http.Request) (*userDetails, error) {
	ud := &userDetails{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &ud)
	if err != nil {
		return nil, errors.New("unable")
	}

	return ud, nil
}
*/

func JsonRequestDecoder(r *http.Request) (*JsonRequest, error) {
	jr := &JsonRequest{}
	err := json.NewDecoder(r.Body).Decode(&jr.Data)
	if err != nil {
		return nil, err
	}

	return jr, nil
}

func GetValue(jr *JsonRequest, key string) (string, bool) {
	data, ok := jr.Data[key]
	if !ok {
		return "", false
	}

	return fmt.Sprintf("%v", data), true
}
