package services

import (
	"errors"
	"log"
)

type LoginService struct{}

func NewLoginService() *LoginService {
	return &LoginService{}
}

func (s *LoginService) LoginCert(systemLawyerID, idSystem, idLawyer string) (map[string]interface{}, error) {
	if systemLawyerID == "" && idLawyer == "" {
		return nil, errors.New("Id do advogado inválido")
	}
	idTemp := systemLawyerID
	if idTemp == "" {
		idTemp = idLawyer
	}

	log.Printf("Login realizado com sucesso para cert ID: %s", idTemp)
	response := map[string]interface{}{
		"message": "Login realizado com sucesso",
		"cert_id": idTemp,
	}
	return response, nil
}
