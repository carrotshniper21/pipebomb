// pipebomb/util/dex.go
package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"fmt"
)

func Dechiper(encoded_url string) []byte {
	resp, err := http.Get("https://raw.githubusercontent.com/carrotshniper21/key/e4/key.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	key := string(body)

	result, _ := DecryptString(key, encoded_url)
	return result
}

var openSSLSaltHeader string = "Salted_" 

type OpenSSLCreds struct {
	key []byte
	iv  []byte
}

func DecryptString(passphrase, encryptedBase64String string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(encryptedBase64String)
	if err != nil {
		return nil, err
	}
	saltHeader := data[:aes.BlockSize]
	if string(saltHeader[:7]) != openSSLSaltHeader {
		return nil, fmt.Errorf("Does not appear to have been encrypted with OpenSSL, salt header missing.")
	}
	salt := saltHeader[8:]
	creds, err := extractOpenSSLCreds([]byte(passphrase), salt)
	if err != nil {
		return nil, err
	}
	return decrypt(creds.key, creds.iv, data)
}

func decrypt(key, iv, data []byte) ([]byte, error) {
	if len(data) == 0 || len(data)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("bad blocksize(%v), aes.BlockSize = %v\n", len(data), aes.BlockSize)
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	cbc := cipher.NewCBCDecrypter(c, iv)
	cbc.CryptBlocks(data[aes.BlockSize:], data[aes.BlockSize:])
	out, err := pkcs7Unpad(data[aes.BlockSize:], aes.BlockSize)
	if out == nil {
		return nil, err
	}
	return out, nil
}

func extractOpenSSLCreds(password, salt []byte) (OpenSSLCreds, error) {
	m := make([]byte, 48)
	prev := []byte{}
	for i := 0; i < 3; i++ {
		prev = hash(prev, password, salt)
		copy(m[i*16:], prev)
	}
	return OpenSSLCreds{key: m[:32], iv: m[32:]}, nil
}

func hash(prev, password, salt []byte) []byte {
	a := make([]byte, len(prev)+len(password)+len(salt))
	copy(a, prev)
	copy(a[len(prev):], password)
	copy(a[len(prev)+len(password):], salt)
	return md5sum(a)
}

func md5sum(data []byte) []byte {
	h := md5.New()
	h.Write(data)
	return h.Sum(nil)
}

func pkcs7Unpad(data []byte, blocklen int) ([]byte, error) {
	if blocklen <= 0 {
		return nil, fmt.Errorf("invalid blocklen %d", blocklen)
	}
	if len(data)%blocklen != 0 || len(data) == 0 {
		return nil, fmt.Errorf("invalid data len %d", len(data))
	}
	padlen := int(data[len(data)-1])
	if padlen > blocklen || padlen == 0 {
		return nil, fmt.Errorf("invalid padding")
	}
	pad := data[len(data)-padlen:]
	for i := 0; i < padlen; i++ {
		if pad[i] != byte(padlen) {
			return nil, fmt.Errorf("invalid padding")
		}
	}
	return data[:len(data)-padlen], nil
}
