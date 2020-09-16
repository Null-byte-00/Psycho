package decryption

import (
	"io/ioutil"
	"crypto/cipher"
	"crypto/aes"
	"encoding/base64"
	"strings"
	"os"
)
//decryption struct
type Decryption struct {
	Key string
	Ext string
}

//create new decryption object
func NewDecryption(key string, ext string) Decryption {
	return Decryption {
		Key: key,
		Ext: ext,
	}
}

//checks the password with test.PSYCHO file
func (dec *Decryption) Checktest() bool {
	readdata, _ := ioutil.ReadFile("C:\\Psychodata\\test.PSYCHO")
	block, _ := aes.NewCipher([]byte(dec.Key))
	gcm, _ := cipher.NewGCM(block)
	nonce, encryptedtext := readdata[:gcm.NonceSize()], readdata[gcm.NonceSize():]
	text, _ := gcm.Open(nil, nonce, encryptedtext, nil)
	finaltext, _ := base64.StdEncoding.DecodeString(string(text))
	if string(finaltext) == "psycho" {
		return true
	}
	return false
}

//decrypts files with a rsa key
func (dec *Decryption) Decryptfile(filename string) {
	readdata, _ := ioutil.ReadFile(filename)
	block, _ := aes.NewCipher([]byte(dec.Key))
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce, encryptedtext := readdata[:nonceSize], readdata[nonceSize:]
	text, _ := gcm.Open(nil, nonce, encryptedtext, nil)
	finaltext, _ := base64.StdEncoding.DecodeString(string(text))
	ioutil.WriteFile(strings.Replace(filename, dec.Ext, "", -1 ), finaltext, 0644)
	os.Remove(filename)
}