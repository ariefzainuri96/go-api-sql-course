package main

import (
	"encoding/json"
	"net/http"

	"github.com/ariefzainuri96/go-api-blogging/cmd/api/request"
	"github.com/ariefzainuri96/go-api-blogging/cmd/api/response"
)

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	var baseResp response.BaseResponse

	var data request.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = "Invalid request"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	resp, err := app.store.Auth.Login(r.Context(), data)

	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = err.Error()
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusInternalServerError)
		return
	}

	baseResp.Status = http.StatusOK
	baseResp.Message = "Success login!"
	resp.BaseResponse = baseResp
	loginResp, _ := resp.Marshal()

	w.WriteHeader(http.StatusOK)
	w.Write(loginResp)
}

func (app *application) AuthRouter() *http.ServeMux {
	authRouter := http.NewServeMux()

	authRouter.HandleFunc("POST /login", app.login)

	return authRouter
}
