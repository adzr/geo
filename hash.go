/*
Copyright 2017 Ahmed Zaher

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package geo

import (
	"fmt"
	"math"
)

const (
	// bitMask is a mask used in converting the hash bits into a Base32 string.
	bitMask = 0x1f
	// MaxHashBits is the maximum number of bits used to form geo-location point hash.
	MaxHashBits = 60
	// base32 is a string containing all the base32 characters.
	base32 = `0123456789bcdefghjkmnpqrstuvwxyz`
)

// Hash represents a wrapper that carries the necessary data about a certain geohash.
type Hash interface {
	fmt.Stringer
	// Value returns the visible integer value of the hash bits.
	Value() uint64
	// Bits returns a 64 bit integer that carries the geohash bits that starts at the most significant bit of the integer.
	Bits() uint64
	// Size returns the number of bits that represents the geohash value returned by Bits()
	// starting at the most significant bit and ending with the count returned by Size().
	Size() uint8
}

type hash struct {
	str  string
	val  uint64
	bits uint64
	size uint8
}

func (h *hash) Value() uint64 {
	if h != nil {
		return h.val
	}
	return 0
}

// String returns the Base32 string representation of the geohash.
func (h *hash) String() string {
	if h != nil {
		return h.str
	}
	return ""
}

func (h *hash) Bits() uint64 {
	if h != nil {
		return h.bits
	}
	return 0
}

func (h *hash) Size() uint8 {
	if h != nil {
		return h.size
	}
	return 0
}

// NewHash creates a new instance of Hash with the specified bits and its size, with maximum size of MaxHashBits.
// The function instantly calculates the integer value and the base32 string of the specified geohash bits and store
// them into the instance fields.
func NewHash(bits uint64, size uint8) Hash {

	size = uint8(math.Min(float64(size), MaxHashBits))
	val := bits >> (64 - size)
	str := ""

	bitCount := 0

	for i := 0; uint8(i) < size; i++ {

		bitCount++

		if bitCount == 5 {
			str += string(base32[(bits>>(64-uint8(i)-1))&bitMask])
			bitCount = 0
		}
	}

	return &hash{str: str, val: val, bits: bits, size: size}
}
