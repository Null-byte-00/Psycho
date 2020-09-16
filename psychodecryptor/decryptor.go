package main

/*** psycho decryptor program ***/
import (
	"encoding/pem"
	"encoding/hex"
	"crypto/rsa"
	"crypto/x509"
	_"crypto/sha256"
	_"crypto/rand"
	"crypto"
	_"io/ioutil"
	"errors"
	"fmt"
	"os"
)

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

func PemToprivatekey(privPEM string) (*rsa.PrivateKey, error) {
    block, _ := pem.Decode([]byte(privPEM))
    if block == nil {
            return nil, errors.New("failed to parse PEM block containing the key")
    }

    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
            return nil, err
    }

    return priv, nil
}

func Rsadecrypt(prikeystr string, text string) string {
	prikey, _ := PemToprivatekey(prikeystr)
	decryptedtext, err := prikey.Decrypt(nil, []byte(text), &rsa.OAEPOptions{Hash: crypto.MD5})
	if err != nil {
		fmt.Println("hahahahah:", err)
	}
	return string(decryptedtext)
}

func main() {
	prikeypem := "-----BEGIN RSA PRIVATE KEY-----\n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\n-----END RSA PRIVATE KEY-----"
	prikey, _ := PemToprivatekey(prikeypem)
	encryptedkey, _ := hex.DecodeString(os.Args[1])
	decryptedkey, _ := prikey.Decrypt(nil, []byte(encryptedkey), &rsa.OAEPOptions{Hash: crypto.MD5})
	fmt.Println(string(decryptedkey))
}
