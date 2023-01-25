package main

import "fmt"

var public_key int
var private_key int
var n int

func Encoder(input_string string) []int64 {
	encode_list := []int64{}
	for i := 0; i < len(input_string); i++ {
		encode_list = append(encode_list, encrypt(int64(input_string[i])))
	}
	return encode_list
}

func Decoder(encoded []int64) string {
	s := ""
	for i := 0; i < len(encoded); i++ {
		s += fmt.Sprintf("%c", decrypt(encoded[i]))
	}

	return s
}

func encrypt(message int64) int64 {
	var encrypted_text int64 = 1
	for e := public_key; e > 0; e-- {
		encrypted_text *= int64(message)
	}
	encrypted_text %= int64(n)

	return encrypted_text
}

func decrypt(encrypted_text int64) int64 {
	var decrypted int64 = 1
	for d := private_key; d > 0; d-- {
		decrypted *= int64(encrypted_text)
		decrypted %= int64(n)

	}
	return int64(decrypted)
}

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}

func init() {
	p := 13
	q := 17
	n = p * q
	fi := (p - 1) * (q - 1)
	e := 2
	for {
		if gcd(e, fi) == 1 {
			break
		}
		e++
	}
	public_key = e

	d := 2
	for {
		if (d*e)%fi == 1 {
			break
		}
		d++
	}
	private_key = d
}

// func main() {
// 	// setKeys()
// 	s := "abc"
// 	encoded := encoder(s)
// 	fmt.Println(encoded)
// 	decoded := decoder(encoded)
// 	fmt.Println(decoded)
// }
