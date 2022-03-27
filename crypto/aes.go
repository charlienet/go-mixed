package crypto

import (
	"crypto/aes"
	"crypto/cipher"
)

type aesBlock struct {
	blockCipher
	key []byte
}

func Aes(key []byte) *aesBlock {
	return &aesBlock{key: key}
}

type aesEcbBlock struct {
	*aesBlock
}

func (o *aesBlock) newCipher() (cipher.Block, error) {
	return aes.NewCipher(o.key)
}

func (o *aesBlock) ECB() *aesEcbBlock {
	return &aesEcbBlock{o}
}

var (
	defaultAesIV = make([]byte, aes.BlockSize)
)

type aesCbcBlock struct {
	*aesBlock
	iv []byte
}

func (o *aesBlock) CBC() *aesCbcBlock {
	return &aesCbcBlock{
		aesBlock: o,
		iv:       defaultAesIV,
	}
}

func (o *aesCbcBlock) WithIV(iv []byte) *aesCbcBlock {
	o.iv = iv

	return o
}

func (o *aesEcbBlock) Encrypt(msg []byte) ([]byte, error) {
	block, err := o.newCipher()
	if err != nil {
		return nil, err
	}

	return o.ecbEncrypt(block, msg)
}

func (o *aesEcbBlock) Decrypt(cipherText []byte) ([]byte, error) {
	block, err := o.newCipher()
	if err != nil {
		return nil, err
	}

	return o.ecbDecrypt(block, cipherText)
}

func (o *aesCbcBlock) Encrypt(msg []byte) ([]byte, error) {
	block, err := o.newCipher()
	if err != nil {
		return nil, err
	}

	c := cipher.NewCBCEncrypter(block, o.iv)
	return o.encrypt(c, msg)
}

func (o *aesCbcBlock) Decrypt(chiperText []byte) ([]byte, error) {
	block, err := o.newCipher()
	if err != nil {
		return nil, err
	}

	c := cipher.NewCBCEncrypter(block, o.iv)
	return o.decrypt(c, chiperText)
}
