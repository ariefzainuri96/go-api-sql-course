package response

import (
	"encoding/json"
)

type LoginResponse struct {
	BaseResponse
	Data LoginData `json:"data"`
}

func (r LoginResponse) Marshal() ([]byte, error) {
	marshal, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	return marshal, nil
}

func (r *LoginResponse) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &r)
}

type LoginData struct {
	ID        int64  `json:"id"`
	Token     string `json:"token"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}
