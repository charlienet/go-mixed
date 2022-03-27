package crypto

import (
	"crypto/cipher"
	"crypto/des"
	"errors"
	"strconv"
)

type desInstance struct {
	blockCipher
	key []byte
}

func Des(key []byte) *desInstance {
	return &desInstance{key: key}
}

// 包含des和3des，根据不同的密钥长度计算
func (o *desInstance) newCipher() (cipher.Block, error) {
	switch len(o.key) {
	case 8:
		return des.NewCipher(o.key)
	case 24:
		return des.NewTripleDESCipher(o.key)
	default:
		return nil, errors.New("crypto/des: invalid key size " + strconv.Itoa(len(o.key)))
	}
}

type desEcb struct {
	*desInstance
}

func (o *desInstance) ECB() *desEcb {
	return &desEcb{desInstance: o}
}

func (o *desEcb) Encrypt(data []byte) ([]byte, error) {
	block, err := o.newCipher()
	if err != nil {
		return nil, err
	}

	return o.ecbEncrypt(block, data)
}

func (o *desEcb) Decrypt(cipherText []byte) ([]byte, error) {
	block, err := o.newCipher()
	if err != nil {
		return nil, err
	}

	return o.ecbDecrypt(block, cipherText)
}

var (
	defaultIV = make([]byte, des.BlockSize)
)

type desCbc struct {
	*desInstance
	iv []byte
}

func (o *desInstance) Cbc() *desCbc {
	return &desCbc{desInstance: o, iv: defaultIV}
}

func (o *desCbc) WithIV(iv []byte) *desCbc {
	o.iv = iv

	return o
}

func (o *desCbc) Encrypt(data []byte) ([]byte, error) {
	block, err := o.newCipher()
	if err != nil {
		return nil, err
	}

	e := cipher.NewCBCEncrypter(block, o.iv)
	return o.encrypt(e, data)
}

func (o *desCbc) Decrypt(cipherText []byte) ([]byte, error) {
	block, err := o.newCipher()
	if err != nil {
		return nil, err
	}

	d := cipher.NewCBCDecrypter(block, o.iv)

	return o.decrypt(d, cipherText)
}
