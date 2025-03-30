package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testStringLength(t *testing.T) {
	a := assert.New(t)

	a.Equal(2, stringLength("Hi"))           // Обычный английский текст
	a.Equal(6, stringLength("Привет"))       // Русский текст
	a.Equal(8, stringLength("👋Привет👋"))     // Эмодзи + русский
	a.Equal(4, stringLength("👋Hi👋"))         // Эмодзи + английский
	a.Equal(0, stringLength(""))             // Пустая строка
	a.Equal(1, stringLength(" "))            // Один пробел
	a.Equal(6, stringLength("      "))       // Несколько пробелов
	a.Equal(2, stringLength("😊💡"))           // Два эмодзи
	a.Equal(3, stringLength("A Z"))          // Пробел между символами
	a.Equal(2, stringLength("\u200D\u200B")) // Zero-width joiner (ZWJ) и Zero-width space (ZWS) считаются символами
}
