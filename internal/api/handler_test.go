package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"game-slots/internal/api"
	"game-slots/internal/storage"
)

func TestUploadConfig(t *testing.T) {
	s := storage.NewStorage() // Предполагается, что у вас есть функция NewStorage
	h := api.NewHandler(s)

	tests := []struct {
		name           string
		configType     string
		configName     string
		body           interface{}
		expectedStatus int
	}{
		{
			name:       "valid reels config",
			configType: "reels",
			configName: "testReels",
			body: [][]string{
				{"A", "B", "C", "D", "E"},
				{"F", "A", "F", "B", "C"},
				{"D", "E", "A", "G", "A"},
			},
			expectedStatus: http.StatusAccepted,
		},
		{
			name:       "valid lines config",
			configType: "lines",
			configName: "testLines",
			body: []map[string]interface{}{
				{
					"line": 1,
					"positions": []map[string]int{
						{"row": 0, "col": 0},
						{"row": 1, "col": 1},
						{"row": 2, "col": 2},
						{"row": 1, "col": 3},
						{"row": 0, "col": 4},
					},
				},
				{
					"line": 2,
					"positions": []map[string]int{
						{"row": 2, "col": 0},
						{"row": 1, "col": 1},
						{"row": 0, "col": 2},
						{"row": 1, "col": 3},
						{"row": 2, "col": 4},
					},
				},
				{
					"line": 3,
					"positions": []map[string]int{
						{"row": 1, "col": 0},
						{"row": 2, "col": 1},
						{"row": 1, "col": 2},
						{"row": 0, "col": 3},
						{"row": 1, "col": 4},
					},
				},
			},
			expectedStatus: http.StatusAccepted,
		},
		{
			name:       "valid payouts config",
			configType: "payouts",
			configName: "testPayouts",
			body: []map[string]interface{}{
				{
					"symbol": "A",
					"payout": []int{0, 0, 50, 100, 200},
				},
				{
					"symbol": "B",
					"payout": []int{0, 0, 40, 80, 160},
				},
				{
					"symbol": "C",
					"payout": []int{0, 0, 30, 60, 120},
				},
				{
					"symbol": "D",
					"payout": []int{0, 0, 20, 40, 80},
				},
				{
					"symbol": "E",
					"payout": []int{0, 0, 10, 20, 40},
				},
				{
					"symbol": "F",
					"payout": []int{0, 0, 5, 10, 20},
				},
				{
					"symbol": "G",
					"payout": []int{0, 0, 2, 5, 10},
				},
			},
			expectedStatus: http.StatusAccepted,
		},
		{
			name:           "missing config type",
			configType:     "",
			configName:     "testName",
			body:           []string{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "missing config name",
			configType:     "reels",
			configName:     "",
			body:           []string{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "invalid config type",
			configType:     "invalidType",
			configName:     "testName",
			body:           []string{},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "malformed json body",
			configType:     "reels",
			configName:     "testReels",
			body:           `{"malformed json"}`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, _ := json.Marshal(tt.body)
			req, err := http.NewRequest("POST", "/upload-config?type="+tt.configType+"&name="+tt.configName, bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			h.UploadConfig(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
		})
	}
}
