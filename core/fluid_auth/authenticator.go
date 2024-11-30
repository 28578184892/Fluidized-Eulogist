package authenticator

import (
	"Eulogist/core/minecraft/protocol"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var ServerIP = ""
var ListenPort = "19132"

const minecraftAuthURL = `http://127.0.0.1:23550/authentication`

func MakeAuthentication(ctx context.Context, buffer []byte) (string, error) {
	body := `{"identityPublicKey":"` + base64.StdEncoding.EncodeToString(buffer) + `"}`

	request, _ := http.NewRequestWithContext(ctx, "POST", minecraftAuthURL, strings.NewReader(body))
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
