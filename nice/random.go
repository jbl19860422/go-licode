package nice

import (
	"math/rand"
	"time"
)

type NiceRNG struct {
	priv 	interface{}
}

func NewNiceRNG() *NiceRNG {
	return &NiceRNG{}
}

func (this *NiceRNG) rng_generate_bytes(len uint) []byte {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, len)
	if n, err := rand.Read(b); err != nil || uint(n) != len {
		return nil
	}
	return b
}

func (this *NiceRNG) rng_generate_int(low uint, high uint) uint {
	if low > high {
		panic("low must little than high")
		return 0
	}

	return low + uint(rand.Intn(int(high - low)))
}

/*
 * Generates a stream of octets containing only characters
 * with ASCII codecs of 0x41-5A (A-Z), 0x61-7A (a-z),
 * 0x30-39 (0-9), 0x2b (+) and 0x2f (/). This matches
 * the definition of 'ice-char' in ICE Ispecification,
 * section 15.1 (ID-16).
 *
 * @param rng context
 * @param len number of octets to product
 * @param buf buffer to store the results
 */
func (this *NiceRNG) nice_rng_generate_bytes_print(l uint) []byte {
	var i uint
	chars := string("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	buf := make([]byte, l)
	for i = 0; i < l; i++ {
		buf[i] = chars[this.rng_generate_int(0, uint(len(chars)))];
	}
	return buf
}