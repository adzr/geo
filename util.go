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
	"math"
	"strings"
)

// Direction represents the geographic direction, being it North, East, South or West.
type Direction byte

const (
	// Pi is the infamous value that represents the result of 22/7.
	Pi float64 = 22.0 / 7.0
	// DegToRad is the conversion factor of angle from degrees to radians.
	DegToRad float64 = Pi / 180.0
	// EarthRadiusInKM is the radius of earth measured in kilometers.
	EarthRadiusInKM float64 = 6371.009

	// North direction index.
	North Direction = 0x1
	// East direction index.
	East Direction = 0x2
	// South direction index.
	South Direction = 0x8
	// West direction index.
	West Direction = 0x10
)

var neighbours = map[Direction][]string{
	North: []string{"p0r21436x8zb9dcf5h7kjnmqesgutwvy", "bc01fg45238967deuvhjyznpkmstqrwx"},
	East:  []string{"bc01fg45238967deuvhjyznpkmstqrwx", "p0r21436x8zb9dcf5h7kjnmqesgutwvy"},
	South: []string{"14365h7k9dcfesgujnmqp0r2twvyx8zb", "238967debc01fg45kmstqrwxuvhjyznp"},
	West:  []string{"238967debc01fg45kmstqrwxuvhjyznp", "14365h7k9dcfesgujnmqp0r2twvyx8zb"},
}

var borders = map[Direction][]string{
	North: []string{"prxz", "bcfguvyz"},
	East:  []string{"bcfguvyz", "prxz"},
	South: []string{"028b", "0145hjnp"},
	West:  []string{"0145hjnp", "028b"},
}

// GetDistance returns the distance in kilometers between two given geo-location points using haversine formula.
func GetDistance(p1 Point, p2 Point) float64 {
	lat1 := p1.Latitude() * DegToRad
	lat2 := p2.Latitude() * DegToRad
	dLat := math.Abs(lat2 - lat1)
	dLng := math.Abs((p2.Longitude() - p1.Longitude()) * DegToRad)
	a := math.Pow(math.Sin(dLat/2), 2) + math.Pow(math.Sin(dLng/2), 2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Asin(math.Sqrt(a))
	return EarthRadiusInKM * c
}

// GetHash calculates and returns the geohash of a given geo-location point with a given precision,
// taking into consideration that the precision is the number of bits desired to represent the returned hash.
func GetHash(point Point, precision uint8) Hash {

	if point == nil {
		return nil
	}

	precision = uint8(math.Min(float64(precision), MaxHashBits))

	minLat := SouthPoleLat
	maxLat := NorthPoleLat

	minLng := -HalfLongitude
	maxLng := HalfLongitude

	var hash uint64
	var isEvenBit = true

	for i := 0; uint8(i) < precision; i++ {

		hash <<= 1

		if isEvenBit {
			d := (maxLng + minLng) / 2
			if point.Longitude() > d {
				hash |= 1
				minLng = d
			} else {
				maxLng = d
			}
		} else {
			d := (maxLat + minLat) / 2
			if point.Latitude() > d {
				hash |= 1
				minLat = d
			} else {
				maxLat = d
			}
		}

		isEvenBit = !isEvenBit
	}

	hash <<= (64 - precision)

	return NewHash(hash, precision)
}

// GetHashAndNeighbours returns the geohash of the given geo-location point and all its neighbour geohashes,
// based on the precision given.
// Considering P is the point given, and N is a neighbour, the function will return the following:
//
// N N N
// N P N
// N N N
//
// This can help in locating points in nearby areas.
func GetHashAndNeighbours(point Point, precision uint8) []Hash {

	var hashes = []Hash{}

	if point == nil {
		return hashes
	}

	precision = uint8(math.Min(float64(precision), MaxHashBits))

	hash := GetHash(point, precision)

	if len(hash.String()) == 0 {
		return hashes
	}

	center := ReverseHash(hash.String())

	if point.Latitude() > center.Latitude() {
		north := GetNeighbour(hash.String(), North)
		hashes = append(hashes, GetHashFromString(north))

		if point.Longitude() > center.Longitude() {
			hashes = append(hashes, GetHashFromString(GetNeighbour(north, East)))
		} else if point.Longitude() < center.Longitude() {
			hashes = append(hashes, GetHashFromString(GetNeighbour(north, West)))
		}
	} else if point.Latitude() < center.Latitude() {
		south := GetNeighbour(hash.String(), South)
		hashes = append(hashes, GetHashFromString(south))

		if point.Longitude() > center.Longitude() {
			hashes = append(hashes, GetHashFromString(GetNeighbour(south, East)))
		} else if point.Longitude() < center.Longitude() {
			hashes = append(hashes, GetHashFromString(GetNeighbour(south, West)))
		}
	}

	if point.Longitude() > center.Longitude() {
		hashes = append(hashes, GetHashFromString(GetNeighbour(hash.String(), East)))
	} else if point.Longitude() < center.Longitude() {
		hashes = append(hashes, GetHashFromString(GetNeighbour(hash.String(), West)))
	}

	hashes = append(hashes, hash)

	return hashes
}

// GetNeighbour returns the neighbour geohash of given a geohash in the given direction.
//
// Based on an MIT licensed implementation by Dave Troy.
// http://github.com/davetroy/geohash-js
func GetNeighbour(hash string, direction Direction) string {

	hash = strings.ToLower(hash)

	length := len(hash)

	if length == 0 {
		return ""
	}

	invalidDirection := true

	for _, dir := range []Direction{North, East, South, West} {
		if direction == dir {
			invalidDirection = false
		}
	}

	if invalidDirection {
		return ""
	}

	category := int(math.Mod(float64(length), 2))
	lastChar := string(hash[length-1])
	parent := string(hash[0 : length-1])

	if strings.Index(borders[direction][category], lastChar) != -1 && len(parent) > 0 {
		parent = GetNeighbour(parent, direction)
	}

	return parent + string(base32[strings.Index(neighbours[direction][category], lastChar)])
}

// GetHashFromString returns the geohash instance based on a given Base32 geohash string.
func GetHashFromString(str string) Hash {

	str = strings.ToLower(str)

	length := len(str)

	if length == 0 {
		return nil
	}

	var bits uint64

	for i := 0; i < length; i++ {
		bits <<= 5
		bits |= uint64(strings.Index(base32, string(str[i])))
	}

	size := uint8(length * 5)

	bits <<= (64 - size)

	return NewHash(bits, size)
}

// ReverseHash returns the point that represents the given Base32 geohash string.
// Notice that the point returned is the point that represents the geohash zone,
// in which any point in this zone will have the same hash value considering the
// same precision specified, which means that the point returned may not represent
// the value desired in certain scenarios, so use with caution.
func ReverseHash(hash string) Point {

	length := len(hash)

	if length == 0 {
		return nil
	}

	hash = strings.ToLower(hash)

	minLat := SouthPoleLat
	maxLat := NorthPoleLat
	minLng := -HalfLongitude
	maxLng := HalfLongitude

	var isEvenBit = true

	for i := 0; i < length; i++ {
		char := string(hash[i])
		index := strings.Index(base32, char)

		if index == -1 {
			return nil
		}

		for j := 4; j >= 0; j-- {
			bit := (index >> uint(j)) & 1

			if isEvenBit {
				var lng = (minLng + maxLng) / 2
				if bit == 1 {
					minLng = lng
				} else {
					maxLng = lng
				}
			} else {
				var lat = (minLat + maxLat) / 2
				if bit == 1 {
					minLat = lat
				} else {
					maxLat = lat
				}
			}

			isEvenBit = !isEvenBit
		}
	}

	return NewPoint((minLat+maxLat)/2, (minLng+maxLng)/2)
}

// GetNeededPrecision return the precision needed to used in geohashing
// to attain the desired accuracy with a given radius in kilometers.
func GetNeededPrecision(radiusInKM float64) uint8 {

	var d = EarthRadiusInKM
	var bits uint8

	d /= 2

	for d >= radiusInKM {
		d /= 2
		bits += 2
	}

	return bits
}
