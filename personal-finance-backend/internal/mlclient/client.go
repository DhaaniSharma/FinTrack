package mlclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const MLURL = "http://127.0.0.1:8000/predict/spender-type"

func PredictSpender(input ExpenseInput) (*PredictionData, error) {

	// 1 : Encode the input data to JSON :
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	// 2 : Create HTTP Client :
	client := &http.Client{Timeout: 10 * time.Second}

	// 3: Create HTTP Request :
	req, err := http.NewRequest(http.MethodPost, MLURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// 4. Send Request :
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 5. Decode Response :
	var mlResp MLResponse
	if err := json.NewDecoder(resp.Body).Decode(&mlResp); err != nil {
		return nil, err
	}

	// 6. Handle ML-Level Errors :
	if mlResp.Status != "success" || mlResp.Data == nil {
		return nil, fmt.Errorf("ML Prediction Failed: %v", mlResp.Error)
	}
	return mlResp.Data, nil
}
