package http

import (
	"encoding/json"
	"net/http"

	"github.com/TimBerk/gophKeeper/internal/platform/jwt"
)

// Handler инкапсулирует зависимости и отдаёт готовые chi-handler’ы.
type Handler struct {
	auth  Auth
	vault Vault
}

// NewHandler создает новый обработчик
func NewHandler(a Auth, v Vault) *Handler { return &Handler{auth: a, vault: v} }

// Register отвечает за регистрацию пользователя
func (h *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct{ U, P string }
		_ = json.NewDecoder(r.Body).Decode(&req)
		if err := h.auth.Register(req.U, req.P); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

// Login отвечает за авторизацию пользователя
func (h *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct{ U, P string }
		_ = json.NewDecoder(r.Body).Decode(&req)
		tok, err := h.auth.Login(req.U, req.P)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]string{"token": tok})
	}
}

// AddSecret отвечает за добавление новой secret-записи
func (h *Handler) AddSecret() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := r.Context().Value(jwt.CtxKey("uid")).(string)
		var req struct {
			Type string            `json:"type"`
			Data []byte            `json:"data"`
			Meta map[string]string `json:"meta"`
		}
		_ = json.NewDecoder(r.Body).Decode(&req)
		if err := h.vault.Add(uid, req.Type, req.Data, req.Meta); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

// ListSecrets отвечает за получение списка secret-записей
func (h *Handler) ListSecrets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := r.Context().Value(jwt.CtxKey("uid")).(string)
		list, _ := h.vault.List(uid)
		_ = json.NewEncoder(w).Encode(list)
	}
}

// MetaVersion отвечает за получение версии
func (h *Handler) MetaVersion(ai SchemaInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		v, _ := ai.Current()
		_ = json.NewEncoder(w).Encode(map[string]int64{"version": v})
	}
}
