package main

import (
	rsaCrypt "./rsa-crypt"
	"encoding/base64"
	"fmt"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("RsaCrypt", map[string]interface{}{
		"New": func (addr, pass string) *js.Object {
			crypt, err := rsaCrypt.NewChatCrypto(addr, pass)
			if err != nil {
				fmt.Println(err)
			}
			return js.MakeWrapper(crypt)
		},
		"GetPublicKey": func (crypt *rsaCrypt.ChatCrypto) string {
			return base64.StdEncoding.EncodeToString(crypt.GetPublicKey())
		},
		"EncryptWithPubKey": func (crypt *rsaCrypt.ChatCrypto, publicKey string, plainText string) (text string, status bool) {
			res, err := crypt.EncryptString(plainText, publicKey)
			if err != nil {
				fmt.Println(err)
				return "", false
			}
			return res, true
		},
		"Decrypt": func (crypt *rsaCrypt.ChatCrypto, cipherText string) (text string, status bool) {
			res, err := crypt.DecryptString(cipherText)
			if err != nil {
				fmt.Println(err)
				return "", false
			}
			return res, true
		},
	})
}
