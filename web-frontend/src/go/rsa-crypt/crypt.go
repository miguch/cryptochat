package rsa_crypt

import (
	drr "./deterministic-random-reader"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"io"
)

type ChatCrypto struct {
	Passphrase string
	UserAddress string
	RsaPrivateKey *rsa.PrivateKey
	randReader io.Reader
}

func NewChatCrypto(address, passphrase string) (*ChatCrypto, error) {
	randReader := drr.NewReader(passphrase, address)
	key, err := rsa.GenerateKey(randReader, 1024)
	if err != nil {
		return nil, err
	}
	return &ChatCrypto{
		Passphrase: passphrase,
		UserAddress: address,
		RsaPrivateKey: key,
		randReader: randReader,
	}, nil
}

func (cc *ChatCrypto) Encrypt(plainText, publicKey []byte) ([]byte, error) {
	pubKey, err := x509.ParsePKCS1PublicKey(publicKey)
	if err != nil {
		return []byte{}, err
	}
	return rsa.EncryptPKCS1v15(cc.randReader, pubKey, plainText)
}

func (cc *ChatCrypto) Decrypt(cipherText []byte) ([]byte, error) {
	return cc.RsaPrivateKey.Decrypt(cc.randReader, cipherText, nil)
}

func (cc *ChatCrypto) GetPublicKey() []byte {
	keyBytes := x509.MarshalPKCS1PublicKey(&cc.RsaPrivateKey.PublicKey)
	return keyBytes
}

func (cc *ChatCrypto) EncryptString(plainText, publicKey string) (string, error) {
	key, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return "", err
	}
	res, err := cc.Encrypt([]byte(plainText), key)
	return base64.StdEncoding.EncodeToString(res), err
}

func (cc *ChatCrypto) DecryptString(cipherText string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	res, err := cc.Decrypt(data)
	return string(res), err
}

