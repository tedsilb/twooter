package daoutil

import (
	"math/rand"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateId() string {
	b := make([]rune, 20)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Now() *timestamppb.Timestamp {
	n := time.Now()
	return &timestamppb.Timestamp{
		Seconds: n.Unix(),
		Nanos:   int32(n.UnixNano() - (n.Unix() * 1000000000)),
	}
}
