//Declare file of only the RSA part of the cryptico library

declare class RSAKey {
    n: any;
    e: any;
    d: any;
    p: any;
    q: any;
    dmp1: any;
    dmq1: any;
    coeff: any;
}

declare class cryptico {
    generateRSAKey(passphrase: string, bitLength: number): RSAKey;
    publicKeyString(rsakey: string): string;
    publicKeyID(publicKeyString: string): string;
    encrypt(plaintext: string, publicKeyString: string, signingKey: RSAKey): {
        status: string;
        cipher: string|undefined;
    };
    decrypt(ciphertext: string, key: RSAKey): {
        status: string;
        plaintext: string|undefined;
        signature: string|undefined;
        publicKeyString: string|undefined;
    }

}