package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxy"

func init(){
	rand.Seed(time.Now().UnixNano())
}


// RandomInt generate a random integer 64 between min and max
func RandomInt64(min, max int64) int64{
	return min + rand.Int63n(max-min+1)
}

// RandomInt generate a random integer 32 between min and max
func RandomInt32(min, max int32) int32{
	return min + rand.Int31n(max-min+1)
}

// RandomInt generate a random string of n chars
func RandomString(n int) string{
	var sb strings.Builder
	k := len(alphabet)

	for i:=0; i<n ; i++{
		c:= alphabet[ rand.Intn(k) ]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generates a random name
func RandomOwner() string{
	return RandomString(6)
}

// Generates a random phone
func RandomPhone() int64{
	return RandomInt64(0,9899358239)
 }

// Generates a random type
func RandomType() int32{
	return RandomInt32(0,2)
 }
