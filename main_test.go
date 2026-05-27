package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerServeHTTP(t *testing.T) {
	t.Parallel()

	subtests := map[string]struct {
		path     string
		wantCode int
	}{
		"returns 404 for /404": {
			path:     "/404",
			wantCode: http.StatusNotFound,
		},
		"returns 401 for /401": {
			path:     "/401",
			wantCode: http.StatusUnauthorized,
		},
		"returns 200 for /200": {
			path:     "/200",
			wantCode: http.StatusOK,
		},
		"returns 404 for root path": {
			path:     "/",
			wantCode: http.StatusNotFound,
		},
		"returns 404 for non-numeric path": {
			path:     "/foo",
			wantCode: http.StatusNotFound,
		},
		"returns 404 for unknown code": {
			path:     "/509",
			wantCode: http.StatusNotFound,
		},
		"returns 404 for numeric with trailing text": {
			path:     "/200abc",
			wantCode: http.StatusNotFound,
		},
		"returns 404 for numeric with extra segment": {
			path:     "/200/extra",
			wantCode: http.StatusNotFound,
		},
		"returns 404 for zero code": {
			path:     "/0",
			wantCode: http.StatusNotFound,
		},
		"returns 404 for negative code": {
			path:     "/-1",
			wantCode: http.StatusNotFound,
		},
	}
	h := &handler{}
	for subtestName, subtest := range subtests {
		t.Run(subtestName, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, subtest.path, nil)
			got := httptest.NewRecorder()
			h.ServeHTTP(got, req)
			if got.Code != subtest.wantCode {
				t.Fatalf(
					"GET %s => %d, want %d",
					subtest.path, got.Code, subtest.wantCode,
				)
			}
		},
		)
	}
}
