package crypto

import (
	"sync"

	"github.com/tjfoc/gmsm/sm4"
)

type sm4Instance struct {
	key []byte
}

func Sm4(key []byte) *sm4Instance {
	return &sm4Instance{key: key}
}

type sm4EcbInstance struct {
	*sm4Instance
}

func (o *sm4Instance) ECB() *sm4EcbInstance {
	return &sm4EcbInstance{
		sm4Instance: o,
	}
}

func (o *sm4EcbInstance) Encrypt(msg []byte) ([]byte, error) {
	return sm4.Sm4Ecb(o.key, msg, true)
}

func (o *sm4EcbInstance) Decrypt(cipherText []byte) ([]byte, error) {
	return sm4.Sm4Ecb(o.key, cipherText, false)
}

type sm4CbcInstance struct {
	*sm4Instance
	iv   []byte
	lock sync.Mutex
}

func (o *sm4Instance) CBC() *sm4CbcInstance {
	return &sm4CbcInstance{
		sm4Instance: o,
	}
}

func (o *sm4CbcInstance) WithIV(iv []byte) *sm4CbcInstance {
	o.iv = iv

	return o
}

func (o *sm4CbcInstance) Encrypt(msg []byte) ([]byte, error) {
	o.lock.Lock()
	defer o.lock.Unlock()

	if err := sm4.SetIV(o.iv); err != nil {
		return nil, err
	}
	defer resetIV()

	return sm4.Sm4Cbc(o.key, msg, true)
}

func (o *sm4CbcInstance) Decrypt(cipherText []byte) ([]byte, error) {
	o.lock.Lock()
	defer o.lock.Unlock()

	if err := sm4.SetIV(o.iv); err != nil {
		return nil, err
	}
	defer resetIV()

	return sm4.Sm4Cbc(o.key, cipherText, false)
}

var emptyIV = make([]byte, sm4.BlockSize)

func resetIV() {
	_ = sm4.SetIV(emptyIV)
}
