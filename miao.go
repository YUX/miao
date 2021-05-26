package miao

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"strings"
)

const b64 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/="

var hashTable = map[string]string{}
var reverseTable = map[string]string{}

func buildTable() map[string]string {
	var z = [3]string{`\u200b`, `\u200c`, `\u200d`}
	for i, c := range b64 {
		key := string(c)
		value := ""
		for _, d := range fmt.Sprintf("%04s", big.NewInt(int64(i)).Text(3)) {
			value += z[int(d)-48]
		}
		hashTable[key] = value
		reverseTable[value] = key
	}
	return hashTable
}

func init() {
	buildTable()
}

func Encode(humainLang string) string {
	humainLangB64 := base64.StdEncoding.EncodeToString([]byte(humainLang))
	table := hashTable
	miaoLang := `喵`
	for _, c := range humainLangB64 {
		if rand.Float64() < 0.2 {
			miaoLang += "喵"
		}
		miaoLang += table[fmt.Sprintf("%c", c)]
	}
	miaoLang += "喵"
	log.Println(miaoLang)
	return miaoLang
}
func Decode(miaoLang string) string {
	miaoLang = strings.ReplaceAll(miaoLang, "喵", "")
	table := reverseTable
	humainLangB64 := ``
	for i := 0; i < len(miaoLang); i = i + 24 {
		humainLangB64 += table[miaoLang[i:i+24]]
	}
	humainLangByte, _ := base64.StdEncoding.DecodeString(humainLangB64)
	return string(humainLangByte)
}
