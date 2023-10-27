package marvel

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Execute() {
	const apiPublicKey = "12f13954a4dd77b17c27da4c76f7dcc3"
	const apiPrivateKey = "001d34388348d0f60f86addf054a350aab4da1cf"

	url := generateUrl(apiPrivateKey, apiPublicKey)

	// Send GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print the response
	fmt.Println(string(body))
}

func generateUrl(apiPrivateKey string, apiPublicKey string) string {
	timestamp := time.Now().Unix()
	hasher := md5.New()
	hasher.Write([]byte(fmt.Sprintf("%d%s%s", timestamp, apiPrivateKey, apiPublicKey)))
	hash := hex.EncodeToString(hasher.Sum(nil))

	baseURL := "https://gateway.marvel.com/v1/public"
	endpoint := "/characters" // Replace this with the appropriate endpoint
	url := fmt.Sprintf("%s%s?ts=%d&apikey=%s&hash=%s", baseURL, endpoint, timestamp, apiPublicKey, hash)
	return url
}
