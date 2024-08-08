package fnord

import (
	"crypto/rand"
	"math/big"
	"regexp"
	"strings"
)

const magicWord = `JARVIS`

// const magicWord = `FOO`

func GetRandomRune() rune {
	max := big.NewInt(25)
	nBig, err := rand.Int(rand.Reader, max.Add(max, big.NewInt(1)))
	if err != nil {
		panic(err.Error())
	}

	n := nBig.Int64()

	return rune(n + 65)
}

func GetRandomString(length int) string {
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(GetRandomRune())
	}
	return b.String()
}

func StringContainsName(s string) bool {
	return strings.Contains(s, magicWord)
}

func StringRegexName(s string) bool {
	return regexp.MustCompile(magicWord).MatchString(s)
}
