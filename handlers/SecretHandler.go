package handlers

import (
	"encoding/json"
	"github.com/bernardosecades/sharesecret/dto"
	"github.com/bernardosecades/sharesecret/services"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type SecretHandler struct {
	secretService services.SecretService
}

func NewSecretHandler(s services.SecretService) *SecretHandler {
	return &SecretHandler{s}
}

func (controller *SecretHandler) CreateSecret(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var cs dto.CreateSecretRequest
	_ = json.NewDecoder(r.Body).Decode(&cs)

	if len(cs.Content) == 0 {
		customError := dto.ErrorResponse{StatusCode: 400, Err: "Empty 'content'"}
		w.WriteHeader(customError.StatusCode)
		_ = json.NewEncoder(w).Encode(customError)
		return
	}

	pass := r.Header.Get("X-Password")
	secret, err := controller.secretService.CreateSecret(cs.Content, pass)

	if err != nil {
		customError := dto.ErrorResponse{StatusCode: 404, Err: err.Error()}
		w.WriteHeader(customError.StatusCode)
		_ = json.NewEncoder(w).Encode(customError)
		return
	}

	cr := dto.CreateSecretResponse{
		Url: os.Getenv("SERVER_URL") + ":" + os.Getenv("SERVER_PORT") + "/secret/" + secret.Id,
	}
	_ = json.NewEncoder(w).Encode(cr)
}

func (controller *SecretHandler) GetSecret(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	pass := r.Header.Get("X-Password")

	customPwd, err := controller.secretService.HasSecretWithCustomPwd(id)

	if err != nil {
		customError := dto.ErrorResponse{StatusCode: 404, Err: err.Error()}
		w.WriteHeader(customError.StatusCode)
		_ = json.NewEncoder(w).Encode(customError)
		return
	}

	if customPwd && pass == "" {
		customError := dto.ErrorResponse{StatusCode: 400, Err: "Missing X-Password"}
		w.WriteHeader(customError.StatusCode)
		_ = json.NewEncoder(w).Encode(customError)
		return
	}

	content, _ := controller.secretService.GetContentSecret(id, pass)

	if len(content) == 0 {
		customError := dto.ErrorResponse{StatusCode: 400, Err: "We can not decrypt the content"}
		w.WriteHeader(customError.StatusCode)
		_ = json.NewEncoder(w).Encode(customError)
		return
	}

	sr := dto.SecretResponse{content}
	_ = json.NewEncoder(w).Encode(sr)
}