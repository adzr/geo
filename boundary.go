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

// Boundary represents a predefined geo-location polygon identified by two points.
// As a rule, the boundary lower bound always has to be on bottom left, while the upper bound has to be on
// the top right.
// All the points with latlng values between these two points' latlng
// are considered located inside this boundary polygon.
type Boundary interface {
	// Lower returns the lower bound point.
	Lower() Point
	// Upper returns the upper bound point.
	Upper() Point
}

type boundary struct {
	lower Point
	upper Point
}

func (b *boundary) Lower() Point {
	if b == nil || b.lower == nil {
		return &point{0, 0}
	}
	return b.lower
}

func (b *boundary) Upper() Point {
	if b == nil || b.upper == nil {
		return &point{0, 0}
	}
	return b.upper
}

func (b *boundary) String() string {
	return fmt.Sprintf("[%v -> %v]", b.Lower(), b.Upper())
}

// NewBoundary creates a new boundary instance with two specified geo-location points.
// As a rule, the boundary lower bound always has to be on bottom left, while the upper bound has to be on
// the top right.
// If the values passed to the NewBoundary don't satisfy this rule, then the values are auto-normalized to keep
// everything clean and in control.
// This way we won't have an issue dealing with bounds above 180 or below -180 longitude.
// Also we won't have an issue figuring out which direction (east or west) we're going to
// calculate the boundary starting from the lower bound till we reach the upper bound.
func NewBoundary(lower Point, upper Point) Boundary {
	// Normalizing the boundary.
	lowerLat := math.Min(lower.Latitude(), upper.Latitude())
	upperLat := math.Max(lower.Latitude(), upper.Latitude())
	var lowerLng float64
	var upperLng float64

	if math.Abs(lower.Latitude()) == NorthPoleLat || math.Abs(upper.Latitude()) == NorthPoleLat {
		lowerLng = 0
		upperLng = 0
	} else {
		lowerLng = lower.Longitude()
		upperLng = upper.Longitude()
	}

	chLower := make(chan Point, 1)
	chUpper := make(chan Point, 1)

	go func() {
		chLower <- NewPoint(lowerLat, lowerLng)
	}()

	go func() {
		chUpper <- NewPoint(upperLat, upperLng)
	}()

	return &boundary{lower: <-chLower, upper: <-chUpper}
}
