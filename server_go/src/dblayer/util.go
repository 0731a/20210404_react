package dblayer

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(s *string) error {
	if s == nil {
		return errors.New("Reference provided for hashing password is nil")
	}
	//  bcrypt 패키지에서 사용할 수 있게 패스워드 문자열을 바이트 슬라이스로 변환한다.
	sBytes := []byte(*s)
	// GenerateFromPassword()  // 패스워드 해시를 반환
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// 패스워드 문자열을 해시값으로 바꾼다.
	*s = string(hashedBytes[:])
	return nil
}

func checkPassword(existingHash, incomingPass string) bool {
	// 해시와 패스워드 문자열이 일치하지 않을 경우 에러를 반환
	return bcrypt.CompareHashAndPassword([]byte(existingHash), []byte(incomingPass)) == nil

}
