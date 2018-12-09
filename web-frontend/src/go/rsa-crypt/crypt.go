package rsa_crypt

import (
	drr "./deterministic-random-reader"
	"crypto/rsa"
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
	key, err := rsa.GenerateKey(randReader, 2048)
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

func (cc *ChatCrypto) Encrypt(plainText []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(cc.randReader, &cc.RsaPrivateKey.PublicKey, plainText)
}

func (cc *ChatCrypto) Decrypt(cipherText []byte) ([]byte, error) {
	return cc.RsaPrivateKey.Decrypt(cc.randReader, cipherText, nil)
}
