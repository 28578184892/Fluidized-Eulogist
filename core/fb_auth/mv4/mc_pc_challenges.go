package fbauth

import (
	fb_client "Eulogist/core/fb_auth/mv4/client"
	authenticator "Eulogist/core/fluid_auth"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type FNumRequest struct {
	Data string `json:"data"`
}

type CheckNumRequest struct {
	DynamicMcpShit string `json:"DynamicMcpShit"`
	PlayerId       int64  `json:"PlayerId"`
	DynamicMcpData string `json:"DynamicMcpData"`
}

func TransferData(client *fb_client.Client, content string) string {

	resp, err := client.HttpClient.Post(authenticator.FluidAuthBaseUrl+"transfer", "application/json", bytes.NewBuffer([]byte(content)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ""
	}

	return string(body)
}

func TransferCheckNum(client *fb_client.Client, mcp string, salt string, localPlayerId int64, memcheck []any) string {

	data := CheckNumRequest{
		DynamicMcpShit: salt,
		PlayerId:       localPlayerId,
		DynamicMcpData: mcp,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}

	resp, err := client.HttpClient.Post(authenticator.FluidAuthBaseUrl+"checknum", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	return string(body)
}
