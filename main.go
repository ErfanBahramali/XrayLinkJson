package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	libXray "github.com/xtls/libxray"
	xrayLog "github.com/xtls/xray-core/common/log"
)

type response struct {
	Success bool `json:"success"`
	Data    any  `json:"data,omitempty"`
	Error   any  `json:"error,omitempty"`
}

type silentLogHandler struct{}

func (silentLogHandler) Handle(xrayLog.Message) {}

func main() {
	// Silence xray-core warnings so stdout stays pure JSON.
	xrayLog.RegisterHandler(silentLogHandler{})

	// Check if command-line argument is provided
	if len(os.Args) < 2 {
		printJSON(response{Success: false, Error: "Please provide a Share link or Xray JSON"})
		return
	}

	input := os.Args[1] // Get the input from the second command-line argument

	// Determine if the input is a Share link or Xray JSON based on the first character
	if strings.HasPrefix(input, "{") {
		// Convert Xray JSON to Share links
		convertXrayJsonToShareLinks(input)
	} else {
		// Convert Share link to Xray JSON
		convertShareLinkToXrayJson(input)
	}
}

// Convert Share link to Xray JSON
func convertShareLinkToXrayJson(shareLink string) {
	// Convert the Share link to Base64 (if it's not already Base64)
	shareLinkBase64 := base64.StdEncoding.EncodeToString([]byte(shareLink))

	// Call the ConvertShareLinksToXrayJson function from the libXray package
	result := libXray.ConvertShareLinksToXrayJson(shareLinkBase64)

	// Check and display the result
	if result == "" {
		printJSON(response{Success: false, Error: "Error converting Share link to Xray JSON"})
		return
	}

	// Decode from Base64
	decodedBytes, err := base64.StdEncoding.DecodeString(result)
	if err != nil {
		printJSON(response{Success: false, Error: err.Error()})
		return
	}

	printSuccess(decodedBytes)
}

// Convert Xray JSON to Share links
func convertXrayJsonToShareLinks(xrayJson string) {
	// Base64 encode the Xray JSON string
	encodedJson := base64.StdEncoding.EncodeToString([]byte(xrayJson))

	// Call the ConvertXrayJsonToShareLinks function from the libXray package
	result := libXray.ConvertXrayJsonToShareLinks(encodedJson)

	// Check and display the result
	if result == "" {
		printJSON(response{Success: false, Error: "Error converting Xray JSON to Share links"})
		return
	}

	// Decode from Base64
	decodedBytes, err := base64.StdEncoding.DecodeString(result)
	if err != nil {
		printJSON(response{Success: false, Error: err.Error()})
		return
	}

	printSuccess(decodedBytes)
}

func printSuccess(decoded []byte) {
	if json.Valid(decoded) {
		fmt.Println(string(decoded))
		return
	}
	printJSON(response{Success: true, Data: string(decoded)})
}

func printJSON(v response) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Printf("{\"success\":false,\"error\":%q}\n", err.Error())
		return
	}
	fmt.Println(string(b))
}
