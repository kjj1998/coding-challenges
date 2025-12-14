package algo

import (
	"encoding/binary"
	"hash/fnv"
	"math/rand"
	"sort"
	"time"
)

type hashedString struct {
	value string
	hash  uint64
}

func RandomSort(strs []string) []string {
	var hashedStrings []hashedString
	seed := rand.New(rand.NewSource(time.Now().UnixNano())).Uint64()

	for _, str := range strs {
		hashedStrings = append(hashedStrings, hashedString{
			value: str,
			hash:  hashWithSeed(str, seed),
		})
	}

	sort.Slice(hashedStrings, func(i, j int) bool {
		return hashedStrings[i].hash < hashedStrings[j].hash
	})

	result := make([]string, len(strs))
	for i, hashedString := range hashedStrings {
		result[i] = hashedString.value
	}

	return result
}

func hashWithSeed(s string, seed uint64) uint64 {
	h := fnv.New64a()

	binary.Write(h, binary.LittleEndian, seed)
	h.Write([]byte(s))

	return h.Sum64()
}
