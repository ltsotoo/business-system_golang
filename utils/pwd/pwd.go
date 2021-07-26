package pwd

import (
	"encoding/base64"

	"golang.org/x/crypto/scrypt"
)

//采用Scrypt加密密码
func ScryptPwd(password string) (string, error) {
	const pwdLen = 10
	salt := []byte{1, 6, 9, 10, 99, 100, 199, 233}

	HashPwd, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, pwdLen)
	if err != nil {
		return password, err
	}

	finalPwd := base64.StdEncoding.EncodeToString(HashPwd)
	return finalPwd, nil
}
