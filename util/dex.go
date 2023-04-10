// film/dex.go
package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"unicode"
)

func generateKey(salt []byte, output int) ([]byte, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/enimax-anime/key/e4/key.txt")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error closing", err)
		}
	}(resp.Body)

	secret, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	key := md5.Sum(append(secret, salt...))
	currentKey := key[:]
	for len(currentKey) < output {
		key = md5.Sum(append(key[:], secret...))
		key = md5.Sum(append(key[:], salt...))
		currentKey = append(currentKey, key[:]...)
	}

	return currentKey[:output], nil
}

func decipher(encodedURL string) (map[string]interface{}, error) {
	s1, err := base64.StdEncoding.DecodeString(encodedURL)
	if err != nil {
		return nil, err
	}

	key, err := generateKey(s1[8:16], 48)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key[:32])
	if err != nil {
		return nil, err
	}

	decrypted := make([]byte, len(s1[16:]))
	mode := cipher.NewCBCDecrypter(block, key[32:])
	mode.CryptBlocks(decrypted, s1[16:])

	decrypted = bytes.TrimLeftFunc(decrypted, unicode.IsSpace)
	fmt.Printf("Decoded Data: %s\n", string(decrypted))

	var result map[string]interface{}
	fmt.Printf("Decoded Data Before Unmarshal: %s\n", string(decrypted))
	err = json.Unmarshal(decrypted, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func titties(url string) (map[string]interface{}, error) {
	decryptedURLs, err := decipher(url)
	if err != nil {
		return nil, err
	}
	return decryptedURLs, nil
}

func Boobies(encoded string) (string, error) {
	cum, err := titties(encoded)
	if err != nil {
		fmt.Println("error:", err)
		return "", err
	}
	fmt.Println("decrypted URLs:", cum)

	decryptedSources, err := json.Marshal(cum)
	if err != nil {
		return "", err
	}
	return string(decryptedSources), nil
}
