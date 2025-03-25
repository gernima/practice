package lenString

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringLength(t *testing.T) {
	a := assert.New(t)

	a.Equal(2, StringLength("Hi"))           // Обычный английский текст
	a.Equal(6, StringLength("Привет"))       // Русский текст
	a.Equal(8, StringLength("👋Привет👋"))     // Эмодзи + русский
	a.Equal(4, StringLength("👋Hi👋"))         // Эмодзи + английский
	a.Equal(0, StringLength(""))             // Пустая строка
	a.Equal(1, StringLength(" "))            // Один пробел
	a.Equal(6, StringLength("      "))       // Несколько пробелов
	a.Equal(2, StringLength("😊💡"))           // Два эмодзи
	a.Equal(3, StringLength("A Z"))          // Пробел между символами
	a.Equal(2, StringLength("\u200D\u200B")) // Zero-width joiner (ZWJ) и Zero-width space (ZWS) считаются символами
}
