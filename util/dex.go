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

func main(url string) (map[string]interface{}, error) {
	url = "U2FsdGVkX1/bv0J3BD6JIOW7s2Me7qYvfGF7uD5yJSHTLIdj3cvnRpzsxMibTgCwGcNhi3YonZGTcObeSm3A7L3PR1I/WyG0MM6/qVoW1l+EW1v9VmyrETY6IRcDcQ9FnRTtfUVAIogkkmnKc5s0ABtQCGl6ZwC1H5hXJaWru19VIMxAKc/vU8tS4HA8eCmKr4vI5H+sL5cjb3RBu2Abv3WX/PgxlYgsq77xBSWS8PvnI9QHNK3weTuPekTubR9c2qoftbDZXytP2QSRAFksZaRRM2PXUof+GjHgPOUW6JpdZRK/Uc8UOI1xRWxjGcr2kMrvG8nlaB02aa+hJwUCh4/O2LaodllSUYsz1zrwrrEzIKNUOHbroOM0czUxNnPMmL+PWffbJbAKJcPV8wE4/E1h5i7k/HWR3guOKUPwxZe7tvyLPcRYZOI8hWvxCZGCPqTDHL4DiZkTGdNEjk9YjglLKzdnHvRuPcP9jezGhOQ="
	decryptedURLs, err := decipher(url)
	fmt.Println(decryptedURLs)
	if err != nil {
		return nil, err
	}
	return decryptedURLs, nil
}

