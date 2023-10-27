package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "nickname") {
		return errors.New("Nickname ja esta sendo usado")
	}
	if strings.Contains(err, "email") {
		return errors.New("Email ja esta sendo usado")
	}
	if strings.Contains(err, "title") {
		return errors.New("Titulo ja esta sendo usado")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("Senha incorreta")
	}
	return errors.New("Algo esta incorreto")
}
