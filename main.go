package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func main() {
	url := "https://api.foolstack.net/v1/poster"

	payload := map[string]interface{}{
		"url": "https://foolstack.net",
		"option": map[string]interface{}{
			"type":      "png",
			"width":     320,
			"height":    480,
			"waitUntil": "networkidle0",
			"quality":   100,
			"fullPage":  "true",
		},
	}

	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImNsa3Q5Yjh1ZTAwMmFweWV6bXR6a2h4b3EiLCJzZXJ2aWNlIjoiUE9TVEVSIiwiaWF0IjoxNjkwOTUyNDQ2LCJleHAiOjQ4NDY3MTI0NDZ9._SJsQhUnHhsBXAm-Dw1DxzwbknHRVjjyjbOBZAUNPlc")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}
