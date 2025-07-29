package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/TimBerk/gophKeeper/internal/secret/domain"
)

// Client - структура для работы с HTTP-Запросами
type Client struct {
	BaseURL string
	HTTP    *http.Client
}

// New - инициализация нового клиента
func New(base string) *Client {
	return &Client{
		BaseURL: base,
		HTTP: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

// SaveToken - сохранение токена
func (c *Client) SaveToken(tok string) error {
	if err := os.MkdirAll(cacheDir(), 0o700); err != nil {
		return err
	}
	return os.WriteFile(tokenFile(), []byte(tok), 0o600)
}

// LoadToken - загрузка токена
func (c *Client) LoadToken() (string, error) {
	b, err := os.ReadFile(tokenFile())
	return string(b), err
}

// Register - выполнение запроса на регистрацию
func (c *Client) Register(user, pass string) error {
	body, _ := json.Marshal(map[string]string{"U": user, "P": pass})
	resp, err := c.HTTP.Post(c.BaseURL+"/register", "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		return fmt.Errorf("register: %s", resp.Status)
	}
	return nil
}

// Login - выполнение запроса на авторизацию
func (c *Client) Login(user, pass string) error {
	body, _ := json.Marshal(map[string]string{"U": user, "P": pass})
	resp, err := c.HTTP.Post(c.BaseURL+"/login", "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("login failed")
	}
	var out struct {
		Token string `json:"token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return err
	}
	return c.SaveToken(out.Token)
}

// Add - выполнение запроса на добавление записи
func (c *Client) Add(typ string, data []byte, meta map[string]string) error {
	tok, err := c.LoadToken()
	if err != nil {
		return err
	}
	payload, _ := json.Marshal(map[string]any{"type": typ, "data": data, "meta": meta})
	req, _ := http.NewRequest(http.MethodPost, c.BaseURL+"/secret", bytes.NewReader(payload))
	req.Header.Set("Authorization", "Bearer "+tok)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		return fmt.Errorf("add: %s", resp.Status)
	}
	return nil
}

// List - выполнение запроса на получение списка
func (c *Client) List() ([]*domain.Secret, error) {
	tok, err := c.LoadToken()
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest(http.MethodGet, c.BaseURL+"/secret/list", nil)
	req.Header.Set("Authorization", "Bearer "+tok)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var out []*domain.Secret
	return out, json.NewDecoder(resp.Body).Decode(&out)
}

// Sync - выполнение запроса для синхронизации локального хранлища
func (c *Client) Sync(repo domain.Repository) error {
	list, err := c.List()
	if err != nil {
		return err
	}
	for _, s := range list {
		if errSave := repo.Save(s); errSave != nil {
			return errSave
		}
	}
	return nil
}
