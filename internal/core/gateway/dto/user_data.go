package dto

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

type UserAuthData struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Checksum string `json:"checksum,omitempty"`
}

func (dto *UserAuthData) GenerateChecksum() string {
	h := md5.Sum([]byte(dto.Email + dto.encodePassword() + dto.Name))
	return fmt.Sprintf("%x", h)
}

func (dto *UserAuthData) encodePassword() string {
	return base64.StdEncoding.EncodeToString([]byte(dto.Password))
}

func (dto *UserAuthData) DecodePassword() string {
	return base64.StdEncoding.EncodeToString([]byte(dto.Password))
}
