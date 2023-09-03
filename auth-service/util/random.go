package util

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) //0~max-min
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)] //0~k-1
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{USD, EUR, CAD, JPY, GBP, AUD, CHF, CNY, SEK, NZD}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomEmail() string {
	return RandomString(6) + "@gmail.com"
}

func RandomName() string {
	return RandomString(6)
}

func RandomPassword() string {
	return RandomString(6)
}

func RandomBool() bool {
	return rand.Intn(2) == 1
}

func RandomDescription() string {
	return RandomString(6)
}

func RandomPeriod() float64 {
	return float64(rand.Intn(100))
}

func RandomInt32() int32 {
	return int32(rand.Intn(100))
}

func RandomTime() time.Time {
	maxPast := 30 * 24 * time.Hour // 30 days in the past
	randomSeconds := rand.Int63n(int64(maxPast.Seconds()))
	randomDuration := time.Duration(randomSeconds) * time.Second
	randomTime := time.Now().Add(-randomDuration)
	return randomTime
}

func RandomUUID() uuid.UUID {
	u, _ := uuid.NewRandom()
	return u
}
