package account_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"strings"
)

type AccountService struct {
	URL string
	*http.Client
}

func NewAccountService(URL string) AccountService {
	return AccountService{
		URL:    URL,
		Client: &http.Client{},
	}
}

func (s *AccountService) Auth(token string, skipAuth bool) (AuthResponse, error) {
	if skipAuth {
		return AuthResponse{}, nil
	}

	req, err := http.NewRequest("GET", s.URL+"/autenticacao/token-valido", nil)
	if err != nil {
		return AuthResponse{}, err
	}

	q := req.URL.Query()
	q.Add("token", token)
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())
	req.Header.Set("Authorization", token)

	res, err := s.Client.Do(req)
	if err != nil {
		return AuthResponse{}, err
	}

	if res.StatusCode != http.StatusOK {
		return AuthResponse{}, errors.New("unauthorized")
	}

	var response struct {
		Data struct {
			ID     string
			Perfil string
		}
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return AuthResponse{}, err
	}

	return AuthResponse{
		ID:   response.Data.ID,
		Role: Role(response.Data.Perfil),
	}, nil
}

func (s *AccountService) ExistsTeacher(token string, uuid uuid.UUID) error {
	req, err := http.NewRequest("GET", s.URL+"/professor/id", nil)
	if err != nil {
		return err
	}

	if !strings.Contains(token, "Bearer") {
		token = "Bearer " + token
	}

	req.Header.Set("Authorization", token)

	q := req.URL.Query()
	q.Add("id", uuid.String())
	req.URL.RawQuery = q.Encode()

	res, err := s.Client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("teacher not found")
	}

	return nil
}

func (s *AccountService) ExistsStudent(token string, uuid uuid.UUID) error {
	req, err := http.NewRequest("GET", s.URL+"/aluno/id", nil)
	if err != nil {
		return err
	}

	if !strings.Contains(token, "Bearer") {
		token = "Bearer " + token
	}

	req.Header.Set("Authorization", token)

	q := req.URL.Query()
	q.Add("id", uuid.String())
	req.URL.RawQuery = q.Encode()

	res, err := s.Client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("student not found")
	}

	return nil
}
