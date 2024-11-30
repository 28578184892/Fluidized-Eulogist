package fbauth

import (
	fb_client "Eulogist/core/fb_auth/mv4/client"
	"bytes"
	"fmt"
	"io"
	"strconv"
)

// ...
type FNumRequest struct {
	Data string `json:"data"`
}

// ...
func TransferData(client *fb_client.Client, content string) string {

	fmt.Println("The buffer need to transfor:", content)

	resp, err := client.HttpClient.Post("http://127.0.0.1:23550/transfer", "application/json", bytes.NewBuffer([]byte(content)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() // 确保在函数结束时关闭响应体

	// 读取返回的内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ""
	}

	fmt.Println("The TransBuffer:" + string(body))

	return string(body)
}

// ...
func TransferCheckNum(client *fb_client.Client, mcp string, salt string, localPlayerId int64, memcheck []any) string {

	fmt.Println(localPlayerId)
	cntnt := "{\"DynamicMcpShit\":\"" + salt + "\", \"PlayerId\":" + strconv.FormatInt(localPlayerId, 10) + ",\"DynamicMcpData\":\"" + mcp + "\"}"

	fmt.Println("I will request as:", cntnt)
	resp, err := client.HttpClient.Post("http://127.0.0.1:23550/checknum", "application/json", bytes.NewBuffer([]byte(cntnt)))

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() // 确保在函数结束时关闭响应体

	// 读取返回的内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	fmt.Println("The TransBuffer:" + string(body))

	return string(body)
}
