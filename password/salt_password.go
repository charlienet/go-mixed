package password

import (
	"github.com/charlienet/go-mixed/bytesconv"
	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return bytesconv.BytesToString(hash), nil
}

func ComparePassword(hashed string, plain []byte) bool {
	byteHash := bytesconv.StringToBytes(hashed)
	err := bcrypt.CompareHashAndPassword(byteHash, plain)

	return err == nil
}
