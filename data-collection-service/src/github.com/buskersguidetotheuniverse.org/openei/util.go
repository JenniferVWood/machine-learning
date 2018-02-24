package openei

import (
	"fmt"
	"github.com/buskersguidetotheuniverse.org/types"
	"time"
)

func MakeKeyFromTimeStampAndGeo(geometry types.Geometry) string {
	t := time.Now().Unix()
	key := fmt.Sprintf("%v", t)

	// reverse
	runes := []rune(key)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	coords := fmt.Sprintf("%v_%v", geometry.Coordinates[0], geometry.Coordinates[1])
	key = fmt.Sprintf("%v_%v", coords, string(runes))
	return key
}
