package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/clients/responses"
)

type AuthClientInterface interface {
	GetUserInfo(token string) (*responses.UserInfo, error)
}

type AuthClient struct {
	client  *http.Client
	baseURL string
}

func NewAuthClient() *AuthClient {
	return &AuthClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: "http://w220066.ferozo.com/tp_prog2/api",
	}
}

func (auth *AuthClient) GetUserInfo(token string) (*responses.UserInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", auth.baseURL+"/Account/UserInfo", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Authorization", token)

	response, err := auth.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", response.StatusCode, body)
	}

	var userInfo responses.UserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return &userInfo, nil
}
