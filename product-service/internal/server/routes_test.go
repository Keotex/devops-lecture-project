package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Keotex/devops-lecture-project/product-service/pkg/product"
)

func TestProductListHandler_Success(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()

	productListHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var products []product.Product
	if err := json.Unmarshal(w.Body.Bytes(), &products); err != nil {
		t.Errorf("expected valid JSON, got error: %v", err)
	}
	if len(products) == 0 {
		t.Error("expected at least one product")
	}
}

func TestProductListHandler_WrongMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/products", nil)
	w := httptest.NewRecorder()

	productListHandler(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status 405, got %d", w.Code)
	}
}

func TestProductDetailHandler_Success(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	req.SetPathValue("id", "1")
	w := httptest.NewRecorder()

	productDetailHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var p product.Product
	if err := json.Unmarshal(w.Body.Bytes(), &p); err != nil {
		t.Errorf("expected valid JSON product, got error: %v", err)
	}
	if p.ID != 1 {
		t.Errorf("expected product ID 1, got %d", p.ID)
	}
}

func TestProductDetailHandler_NotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/products/999", nil)
	req.SetPathValue("id", "999")
	w := httptest.NewRecorder()

	productDetailHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", w.Code)
	}
}

func TestProductDetailHandler_InvalidID(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/products/abc", nil)
	req.SetPathValue("id", "abc")
	w := httptest.NewRecorder()

	productDetailHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestProductDetailHandler_WrongMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/products/1", nil)
	req.SetPathValue("id", "1")
	w := httptest.NewRecorder()

	productDetailHandler(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status 405, got %d", w.Code)
	}
}

func TestProductDetailHandler_ContentType(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	req.SetPathValue("id", "1")
	w := httptest.NewRecorder()

	productDetailHandler(w, req)

	ct := w.Header().Get("Content-Type")
	if !strings.Contains(ct, "application/json") {
		t.Errorf("expected Content-Type application/json, got %s", ct)
	}
}
