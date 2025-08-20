package coder

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/blowfish"
	"log/slog"
)

var blKey = []byte{'\xd8', '\x89', '\x0a', '\xf0', '\x66', '\xc9', '\x6b', '\x40', '\xd7', '\x01', '\xae', '\xfc', '\x43', '\x6f', '\xf9', '\xfe'}

type Coder struct {
	cipher *blowfish.Cipher
}

func (c *Coder) WithKey(key []byte) error {
	var err error
	if len(key) == 0 {
		key = blKey
	}

	c.cipher, err = blowfish.NewCipher(key)
	if err != nil {
		return fmt.Errorf("cannot create cipher: %w", err)
	}

	return nil
}

func (c *Coder) Decode(src []byte) []byte {
	slog.Debug("predecoded base64", "bytes", len(src), "raw", src)
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(src)))
	n, err := base64.StdEncoding.Decode(dst, src)
	if err != nil {
		slog.Error("cannot decode base64", "error", err.Error(), "src", src)
		return nil
	}
	dst = dst[:n]

	slog.Debug("decoding blowfish", "bytes", n, "raw", dst)
	bl := c.DecodeBlowfish(dst)
	slog.Debug("decoded blowfish", "data", bl, "len", len(bl))

	padValue := bl[len(bl)-1]
	slog.Debug("padding", "size", padValue)
	if padValue < 0x9 {
		bl = bl[:len(bl)-int(padValue)]
	}

	return bl
}

func (c *Coder) DecodeBlowfish(src []byte) []byte {
	dst := make([]byte, 0, len(src))
	tmp := make([]byte, 8)
	for i := 0; i < len(src); {
		start := i
		end := i + 8
		if end >= len(src) {
			end = len(src)
		}

		var input []byte
		for _, v := range src[start:end] {
			input = append(input, v)
		}

		for len(input) < 8 {
			input = append(input, 0)
		}

		c.cipher.Decrypt(tmp, input)
		i = i + 8
		dst = append(dst, tmp...)
		//fmt.Printf("%s % x\n", tmp, tmp)
	}

	return dst
}

func (c *Coder) EncodeBlowfish(src []byte) []byte {
	dst := make([]byte, 0, len(src))
	tmp := make([]byte, 8)
	for i := 0; i < len(src); {
		start := i
		end := i + 8
		if end >= len(src) {
			end = len(src)
		}

		var input []byte
		for _, v := range src[start:end] {
			input = append(input, v)
		}

		for len(input) < 8 {
			input = append(input, 0)
		}

		c.cipher.Encrypt(tmp, input)
		i = i + 8
		dst = append(dst, tmp...)
		//fmt.Printf("%s % x\n", tmp, tmp)
	}

	return dst
}

func (c *Coder) Encode(data []byte) []byte {
	slog.Debug("encode data", "len", len(data))
	padLen := 8 - len(data)&7
	if padLen <= 8 && padLen > 0 {
		padding := make([]byte, padLen)
		for i := range padding {
			padding[i] = byte(padLen)
		}
		slog.Debug("added padding", "size", padLen, "value", fmt.Sprintf("%02x", padding), "dataLen", len(data))
		data = append(data, padding...)
	}

	slog.Debug("data", "data", data, "len", len(data))
	data = c.EncodeBlowfish(data)
	slog.Debug("encoded blowfish", "bytes", len(data), "raw", data)

	dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dst, data)
	slog.Debug("encoded base64", "bytes", len(dst), "raw", dst)

	return dst
}
