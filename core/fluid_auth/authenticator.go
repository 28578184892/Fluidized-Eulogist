package authenticator

import (
	"Eulogist/core/minecraft/protocol"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AuthRequest struct {
	PublicKey string `json:"identityPublicKey"`
}

var ServerIP = ""
var ListenPort = "19132"

const FluidAuthBaseUrl = "http://127.0.0.1:23550/"

func MakeAuthentication(ctx context.Context, buffer []byte) (string, error) {
	minecraftAuthURL := FluidAuthBaseUrl + "authentication"

	authRequest := AuthRequest{
		PublicKey: base64.StdEncoding.EncodeToString(buffer),
	}
	jsonData, _ := json.Marshal(authRequest)

	request, _ := http.NewRequestWithContext(ctx, "POST", minecraftAuthURL, bytes.NewReader(jsonData))
	request.Header.Set("Content-Type", "application/json")

	request.Header.Set("User-Agent", "Fluid/GameService")
	request.Header.Set("Client-Version", protocol.CurrentVersion)

	c := &http.Client{}
	resp, err := c.Do(request)
	if err != nil {
		return "", fmt.Errorf("POST %v: %v", minecraftAuthURL, err)
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("POST %v: %v", minecraftAuthURL, resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	c.CloseIdleConnections()
	return string(data), err
}
