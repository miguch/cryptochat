package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/miguch/cryptochat/web-frontend/src/go/rsa-crypt"
)

func main() {
	js.Global.Set("rsaCrypt", js.MakeWrapper(rsa_crypt.NewChatCrypto))
}
