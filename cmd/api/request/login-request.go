package request

import (
	"encoding/json"
)

type LoginRequest struct {
	Email    string
	Password string
}

func (r LoginRequest) Marshal() ([]byte, error) {
	marshal, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	return marshal, nil
}

func (r *LoginRequest) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &r)
}
