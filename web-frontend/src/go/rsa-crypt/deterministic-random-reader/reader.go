package deterministic_random_reader

import (
	"hash/fnv"
	"math/rand"

)

//This Package Contains a deterministic random reader
// that is used to generate deterministic RSA private key
// from user's passphrase and eth address so that users don't
// have to keep their long RSA key file.

type Reader struct {
	passphrase []byte
	address []byte
	generated int
}

func NewReader(pass, addr string) (*Reader){
	return &Reader{
		[]byte(pass),
		[]byte(addr),
		0,
	}
}

func (r *Reader) Read(b []byte) (n int, err error) {
	for index := range b {
		byte, err := r.ReadOneByte()
		if err != nil {
			return index, err
		}
		b[index] = byte
	}
	return len(b), nil
}

func getBytesHash(str []byte) uint64 {
	h := fnv.New64a()
	h.Write([]byte(str))
	return h.Sum64()
}

func rotateBytes(b []byte) []byte {
	res := make([]byte, len(b))
	copy(res[1:len(b)], b[:len(b)-1])
	res[0] = b[len(b)-1]
	return res
}

func (r *Reader) ReadOneByte() (byte, error) {
	if r.generated % 5 == 0 {
		seed := rotateBytes(r.passphrase)
		seed = append(seed, rotateBytes(r.address)...)
		hash := getBytesHash(seed)
		rand.Seed(int64(hash))
	}
	b := make([]byte, 1)
	_, err := rand.Read(b)
	if err != nil {
		return 0, err
	}
	return b[0], nil
}

