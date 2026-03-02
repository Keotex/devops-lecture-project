package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Keotex/devops-lecture-project/checkout-service/pkg/token"
)

func TestCheckoutPlaceOrderHandler_Success(t *testing.T) {
	validToken, err := token.CreateToken("user")
	if err != nil {
		t.Fatalf("failed to create token: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/checkout/placeorder", nil)
	req.Header.Set("Authorization", "Bearer "+validToken)
	w := httptest.NewRecorder()

	checkoutPlaceOrderHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestCheckoutPlaceOrderHandler_MissingToken(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/checkout/placeorder", nil)
	w := httptest.NewRecorder()

	checkoutPlaceOrderHandler(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected status 401, got %d", w.Code)
	}
}

func TestCheckoutPlaceOrderHandler_InvalidToken(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/checkout/placeorder", nil)
	req.Header.Set("Authorization", "Bearer invalid.token.here")
	w := httptest.NewRecorder()

	checkoutPlaceOrderHandler(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected status 401, got %d", w.Code)
	}
}

func TestCheckoutPlaceOrderHandler_WrongMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/checkout/placeorder", nil)
	w := httptest.NewRecorder()

	checkoutPlaceOrderHandler(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status 405, got %d", w.Code)
	}
}
