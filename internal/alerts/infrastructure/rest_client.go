package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"ws-server/internal/alerts/domain/entities"

)

type RESTClient struct {
	BaseURL string
}

func (c *RESTClient) RegisterDanger(usuarioID int) (*entities.AlertResponse, error) {
	body, err := json.Marshal(map[string]int{"usuario_id": usuarioID})
	if err != nil {
		return nil, fmt.Errorf("error marshaling body: %w", err)
	}
	
	endpoint := fmt.Sprintf("%s/alertas", c.BaseURL)
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("error calling external API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	var alertResponse entities.AlertResponse
	if err := json.NewDecoder(resp.Body).Decode(&alertResponse); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}
	
	return &alertResponse, nil
}