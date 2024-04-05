package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkCRUDOperations(b *testing.B) {
	handlers := map[string]struct {
		method string
		path   string
		body   []byte
	}{
		"GetItems":   {"GET", "/items", nil},
		"GetItem":    {"GET", "/items/1", nil},
		"CreateItem": {"POST", "/items", []byte(`{"id": "1", "name": "Test Item"}`)},
		"UpdateItem": {"PUT", "/items/1", []byte(`{"id": "1", "name": "Updated Item"}`)},
		"DeleteItem": {"DELETE", "/items/1", nil},
	}

	for name, h := range handlers {
		b.Run(name, func(b *testing.B) {
			req, err := http.NewRequest(h.method, h.path, bytes.NewBuffer(h.body))
			if err != nil {
				b.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(getHandlerFuncByName(name))

			for i := 0; i < b.N; i++ {
				handler.ServeHTTP(rr, req)
			}
		})
	}
}

func getHandlerFuncByName(name string) http.HandlerFunc {
	switch name {
	case "GetItems":
		return GetItems
	case "GetItem":
		return GetItem
	case "CreateItem":
		return CreateItem
	case "UpdateItem":
		return UpdateItem
	case "DeleteItem":
		return DeleteItem
	default:
		panic("Unknown handler name")
	}
}
