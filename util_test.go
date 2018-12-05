/*
Copyright 2018 Ahmed Zaher

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
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDistance(t *testing.T) {
	tolerance := 0.05
	assert.InDelta(t, 0, GetDistance(NewPoint(0, 0), NewPoint(0, 0)), tolerance)
	assert.InDelta(t, 111.2, GetDistance(NewPoint(0, 0), NewPoint(0, 1)), tolerance)
	assert.InDelta(t, 111.2, GetDistance(NewPoint(0, 0), NewPoint(1, 0)), tolerance)
	assert.InDelta(t, 111.2, GetDistance(NewPoint(0, 1), NewPoint(0, 0)), tolerance)
	assert.InDelta(t, 111.2, GetDistance(NewPoint(1, 0), NewPoint(0, 0)), tolerance)
	assert.InDelta(t, 855.7, GetDistance(NewPoint(50.432356, 0.873793), NewPoint(58.124521, 0.735753)), tolerance)
	assert.InDelta(t, 1553, GetDistance(NewPoint(50.432356, 83.873793), NewPoint(58.124521, 63.735753)), tolerance)
}

func TestGetNeighbour(t *testing.T) {
	assert.Equal(t, "", GetNeighbour("", North))
	assert.Equal(t, "", GetNeighbour("ub188qkx0", 128))
	assert.Equal(t, "ub188qkx2", GetNeighbour("ub188qkx0", North))
	assert.Equal(t, "ub188qkwb", GetNeighbour("ub188qkx0", South))
	assert.Equal(t, "ub188qkx1", GetNeighbour("ub188qkx0", East))
	assert.Equal(t, "ub188qkrp", GetNeighbour("ub188qkx0", West))
	assert.Equal(t, "ub188qkrr", GetNeighbour(GetNeighbour("ub188qkx0", North), West))
	assert.Equal(t, "ub188qkx3", GetNeighbour(GetNeighbour("ub188qkx0", North), East))
	assert.Equal(t, "ub188qkqz", GetNeighbour(GetNeighbour("ub188qkx0", South), West))
	assert.Equal(t, "ub188qkwc", GetNeighbour(GetNeighbour("ub188qkx0", South), East))
}

func TestReverseHash(t *testing.T) {

	var p Point

	p = ReverseHash("ub188qkx0")
	assert.InDelta(t, 45.123446, p.Latitude(), DecimalPrecision)
	assert.InDelta(t, 35.876563, p.Longitude(), DecimalPrecision)

	p = ReverseHash("ub188qkx3")
	assert.InDelta(t, 45.123489, p.Latitude(), DecimalPrecision)
	assert.InDelta(t, 35.876606, p.Longitude(), DecimalPrecision)

	p = ReverseHash("ub188qkwc")
	assert.InDelta(t, 45.123403, p.Latitude(), DecimalPrecision)
	assert.InDelta(t, 35.876606, p.Longitude(), DecimalPrecision)

	p = ReverseHash("")
	assert.Nil(t, p)

	p = ReverseHash("a")
	assert.Nil(t, p)
}

func TestGetHashFromString(t *testing.T) {
	assert.Equal(t, nil, GetHashFromString(""))
	assert.Equal(t, NewHash(uint64(0xaf777bb800000000), 30), GetHashFromString("pxvrrf"))
	assert.Equal(t, "ub188qkx", GetHashFromString("ub188qkx").String())
}

func TestGetNeededPrecision(t *testing.T) {

	d := EarthRadiusInKM

	for i := 0; i <= 64; i += 2 {
		assert.InDelta(t, i, GetNeededPrecision(d), 0)
		d /= 2
	}
}

func BenchmarkGetDistance(b *testing.B) {

	for i := 0; i < b.N; i++ {
		lat1 := rand.Float64()*180 - 90
		lng1 := rand.Float64()*360 - 180
		p1 := NewPoint(lat1, lng1)

		lat2 := rand.Float64()*180 - 90
		lng2 := rand.Float64()*360 - 180
		p2 := NewPoint(lat2, lng2)

		GetDistance(p1, p2)
	}
}

func BenchmarkGetDistance_Parallel(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lat1 := rand.Float64()*180 - 90
			lng1 := rand.Float64()*360 - 180
			p1 := NewPoint(lat1, lng1)

			lat2 := rand.Float64()*180 - 90
			lng2 := rand.Float64()*360 - 180
			p2 := NewPoint(lat2, lng2)

			GetDistance(p1, p2)
		}
	})
}

func TestHash(t *testing.T) {
	var p Point

	s := GetHash(nil, 16)
	assert.Empty(t, s)

	p = NewPoint(45.12345678, 35.87654321)
	s = GetHash(p, 60)
	assert.Equal(t, "ub188qkx0n18", s.String())

	p = NewPoint(48.6667, -4.334)
	s = GetHash(p, 30)
	assert.Equal(t, "gbsuv7", s.String())

	p = NewPoint(0, 33.22315)
	s = GetHash(p, 45)
	assert.Equal(t, "kxzxupbrg", s.String())

	p = NewPoint(0, 23)
	s = GetHash(p, 5)
	assert.Equal(t, "k", s.String())

	p = NewPoint(0, 179.999979)
	s = GetHash(p, 45)
	assert.Equal(t, "rzzzzzzzz", s.String())

	p = NewPoint(90, 33.220017)
	s = GetHash(p, 45)
	assert.Equal(t, "gzzzzzzzz", s.String())
}

func TestHashAndNeighbours(t *testing.T) {

	var pnt *point
	var p Point = pnt

	assert.Empty(t, GetHashAndNeighbours(nil, 0))

	assert.Empty(t, GetHashAndNeighbours(p, 0))

	assert.Empty(t, GetHashAndNeighbours(NewPoint(45.12354, 35.8766), 0))

	assert.Empty(t, GetHashAndNeighbours(NewPoint(45.12354, 35.8766), 0))

	hashes := GetHashAndNeighbours(NewPoint(45.12354, 35.8766), 40)

	hashMap := make(map[string]Hash)

	for _, hash := range hashes {
		hashMap[hash.String()] = hash
	}

	_, h := hashMap["ub188qkx"]
	_, n := hashMap["ub188qs8"]
	_, w := hashMap["ub188qkr"]
	_, nw := hashMap["ub188qs2"]

	assert.Len(t, hashes, 4)
	assert.True(t, h && n && w && nw)

	hashes = GetHashAndNeighbours(NewPoint(45.12343, 35.8768), 40)

	hashMap = make(map[string]Hash)

	for _, hash := range hashes {
		hashMap[hash.String()] = hash
	}

	_, h = hashMap["ub188qkx"]
	_, e := hashMap["ub188qkz"]
	_, se := hashMap["ub188qky"]
	_, s := hashMap["ub188qkw"]

	assert.Len(t, hashes, 4)
	assert.True(t, h && e && s && se)

	hashes = GetHashAndNeighbours(NewPoint(45.12354, 35.8768), 40)

	hashMap = make(map[string]Hash)

	for _, hash := range hashes {
		hashMap[hash.String()] = hash
	}

	_, h = hashMap["ub188qkx"]
	_, e = hashMap["ub188qkz"]
	_, ne := hashMap["ub188qsb"]
	_, n = hashMap["ub188qs8"]

	assert.Len(t, hashes, 4)
	assert.True(t, h && e && n && ne)

	hashes = GetHashAndNeighbours(NewPoint(45.12343, 35.8766), 40)

	hashMap = make(map[string]Hash)

	for _, hash := range hashes {
		hashMap[hash.String()] = hash
	}

	_, h = hashMap["ub188qkx"]
	_, w = hashMap["ub188qkr"]
	_, sw := hashMap["ub188qkq"]
	_, s = hashMap["ub188qkw"]

	assert.Len(t, hashes, 4)
	assert.True(t, h && w && s && sw)
}

func BenchmarkGetHashAndNeighbours(b *testing.B) {

	for i := 0; i < b.N; i++ {
		lat := rand.Float64()*180 - 90
		lng := rand.Float64()*360 - 180
		c := rand.Intn(60)
		p := NewPoint(lat, lng)
		GetHashAndNeighbours(p, uint8(c))
	}
}

func BenchmarkGetHashAndNeighbours_Parallel(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lat := rand.Float64()*180 - 90
			lng := rand.Float64()*360 - 180
			c := rand.Intn(60)
			p := NewPoint(lat, lng)
			GetHashAndNeighbours(p, uint8(c))
		}
	})
}
