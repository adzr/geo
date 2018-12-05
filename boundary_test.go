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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoundary_ConstructorAndGetters(t *testing.T) {

	var b Boundary

	b = &boundary{lower: &point{latitude: 0, longitude: 1}, upper: &point{latitude: 2, longitude: 3}}

	assert.InDelta(t, 0, b.Lower().Latitude(), 0)
	assert.InDelta(t, 1, b.Lower().Longitude(), 0)
	assert.InDelta(t, 2, b.Upper().Latitude(), 0)
	assert.InDelta(t, 3, b.Upper().Longitude(), 0)

	var bnd *boundary

	b = bnd

	assert.InDelta(t, 0, b.Lower().Latitude(), 0)
	assert.InDelta(t, 0, b.Lower().Longitude(), 0)
	assert.InDelta(t, 0, b.Upper().Latitude(), 0)
	assert.InDelta(t, 0, b.Upper().Longitude(), 0)

	bnd = &boundary{lower: &point{latitude: -90, longitude: 0}, upper: &point{latitude: 0, longitude: 50}}

	assert.Equal(t, "[(-90, 0) -> (0, 50)]", bnd.String())
}

func TestBoundary_ConstructorAndGettersWithNonNormalizedPoints(t *testing.T) {

	var b Boundary

	b = NewBoundary(NewPoint(0, 1), NewPoint(2, 3))
	assert.InDelta(t, 0, b.Lower().Latitude(), 0)
	assert.InDelta(t, 1, b.Lower().Longitude(), 0)
	assert.InDelta(t, 2, b.Upper().Latitude(), 0)
	assert.InDelta(t, 3, b.Upper().Longitude(), 0)

	b = NewBoundary(NewPoint(2, 3), NewPoint(0, 1))
	assert.InDelta(t, 0, b.Lower().Latitude(), 0)
	assert.InDelta(t, 3, b.Lower().Longitude(), 0)
	assert.InDelta(t, 2, b.Upper().Latitude(), 0)
	assert.InDelta(t, 1, b.Upper().Longitude(), 0)

	b = NewBoundary(NewPoint(0, 3), NewPoint(2, 1))
	assert.InDelta(t, 0, b.Lower().Latitude(), 0)
	assert.InDelta(t, 3, b.Lower().Longitude(), 0)
	assert.InDelta(t, 2, b.Upper().Latitude(), 0)
	assert.InDelta(t, 1, b.Upper().Longitude(), 0)

	b = NewBoundary(NewPoint(0, 3), NewPoint(2, -130))
	assert.InDelta(t, 0, b.Lower().Latitude(), 0)
	assert.InDelta(t, 3, b.Lower().Longitude(), 0)
	assert.InDelta(t, 2, b.Upper().Latitude(), 0)
	assert.InDelta(t, -130, b.Upper().Longitude(), 0)

	b = NewBoundary(NewPoint(90, 3), NewPoint(2, -130))
	assert.InDelta(t, 2, b.Lower().Latitude(), 0)
	assert.InDelta(t, 0, b.Lower().Longitude(), 0)
	assert.InDelta(t, 90, b.Upper().Latitude(), 0)
	assert.InDelta(t, 0, b.Upper().Longitude(), 0)

	b = NewBoundary(NewPoint(2, 3), NewPoint(-90, -130))
	assert.InDelta(t, -90, b.Lower().Latitude(), 0)
	assert.InDelta(t, 0, b.Lower().Longitude(), 0)
	assert.InDelta(t, 2, b.Upper().Latitude(), 0)
	assert.InDelta(t, 0, b.Upper().Longitude(), 0)
}
