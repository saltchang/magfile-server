package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func GetRandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func GetRandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func GetRandomEmail() string {
	return GetRandomString(6) + "@" + GetRandomString(int(GetRandomInt(4, 6))) + "." + GetRandomString(int(GetRandomInt(2, 4)))
}

func GetRandomStringArray(n int) []string {
	arr := []string{}
	for i := 0; i < n; i++ {
		arr = append(arr, GetRandomString(int(GetRandomInt(0, 6))))
	}

	return arr
}

func GetRandomBoolean() bool {
	return GetRandomInt(0, 1)%2 == 0
}

func GetRandomStringOption(options []string) string {
	return options[int(GetRandomInt(0, int64(len(options)-1)))]
}
