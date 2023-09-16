package util

import (
	"crypto/rand"
	"math"
	"math/big"
)

type IdGenerator struct{
}
func (ig IdGenerator) GenerateSurrogateKey() string {
	//a~Zをランダムに生成して文字列にする
	var surrogateKey string
	var randChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < 10; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))//最大値はでかければなんでもいい
		if err != nil {
			panic(err)
		}
		surrogateKey += string(randChars[int(n.Int64()) % len(randChars)])
	}
	return surrogateKey
}