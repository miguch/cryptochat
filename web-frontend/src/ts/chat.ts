/** This file includes the code associated with controlling of the
 CryptoChat web interface */

import cryptoUtils from "./rsa_crypto"

window.cryptoUtils = cryptoUtils || undefined;

class Chat {
    public readonly selfAddress: string;
    private passphrase: string | undefined = undefined;
    private crypt: cryptoUtils | undefined = undefined;
    constructor(address: string) {
        this.selfAddress = address;
    }

    public attemptLogin(pass: string) {
        let testCrypt = new cryptoUtils(this.selfAddress, pass);

    }
}

export default Chat;

