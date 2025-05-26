package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertText(mes string) string {
	if len(mes) == 0 {
		err := errors.New("File empty")
		return fmt.Sprintf("Error: %v\n", err)
	}
	russianAlphabet := "абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
	symbolMorse := "-."

	if strings.ContainsAny(mes, russianAlphabet) {
		result := morse.ToMorse(mes)
		return result
	}
	if strings.ContainsAny(mes, symbolMorse) {
		result := morse.ToText(mes)
		return result
	}

	err := errors.New("Error: invalid type text")
	return fmt.Sprintf("Error: %v\n", err)
}
