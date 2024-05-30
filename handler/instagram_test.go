package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func TestSubscription(t *testing.T) {
	type query struct {
		mode        string
		verifyToken string
		challenge   string
	}
	tests := []struct {
		name   string
		token  string
		query  query
		status int
	}{
		{
			name:  "test1",
			token: "",
			query: query{
				mode:        "",
				verifyToken: "",
				challenge:   "",
			},
			status: http.StatusBadRequest,
		},
		{
			name:  "test2",
			token: "token",
			query: query{
				mode:        "subscribe",
				verifyToken: "token",
				challenge:   "1158201444",
			},
			status: http.StatusOK,
		},
	}

	router := gin.Default()
	ig := Instagram{}
	ig.Router(router)
	viper.SetDefault("verify_token", "token")
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(
				http.MethodGet,
				"/api/v1/instagram/webhook",
				nil,
			)

			value := url.Values{}
			value.Add("hub.mode", tt.query.mode)
			value.Add("hub.verify_token", tt.query.verifyToken)
			value.Add("hub.challenge", tt.query.challenge)
			r.URL.RawQuery = value.Encode()
			router.ServeHTTP(w, r)
			if w.Code != tt.status {
				t.Errorf("got: %v, want: %v", w.Code, tt.status)
			}
		})
	}
}
