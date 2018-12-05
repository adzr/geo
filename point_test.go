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
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	PointCount       = 1000000
	DecimalPrecision = 0.000001
)

func TestPoint_ConstructorAndGetters(t *testing.T) {
	for i := 0; i < PointCount; i++ {
		lat := rand.Float64()*180 - 90
		lng := rand.Float64()*360 - 180
		p := NewPoint(lat, lng)
		assert.InDelta(t, lat, p.Latitude(), DecimalPrecision)

		if lng == -180 {
			assert.InDelta(t, 180, p.Longitude(), DecimalPrecision)
		} else {
			assert.InDelta(t, lng, p.Longitude(), DecimalPrecision)
		}
	}
}

func TestPoint_ConstructorAndGettersWithNonNormalizedLatLng(t *testing.T) {

	var lat, lng float64
	var p Point

	lat = 130
	lng = 200
	p = NewPoint(lat, lng)
	assert.InDelta(t, 90, p.Latitude(), 0)
	assert.InDelta(t, 0, p.Longitude(), 0)

	lat = -130
	lng = 200
	p = NewPoint(lat, lng)
	assert.InDelta(t, -90, p.Latitude(), 0)
	assert.InDelta(t, 0, p.Longitude(), 0)

	lat = 0
	lng = 200
	p = NewPoint(lat, lng)
	assert.InDelta(t, 0, p.Latitude(), 0)
	assert.InDelta(t, -160, p.Longitude(), 0)

	lat = 0
	lng = -200
	p = NewPoint(lat, lng)
	assert.InDelta(t, 0, p.Latitude(), 0)
	assert.InDelta(t, 160, p.Longitude(), 0)

	lat = 0
	lng = 180
	p = NewPoint(lat, lng)
	assert.InDelta(t, 0, p.Latitude(), 0)
	assert.InDelta(t, 180, p.Longitude(), 0)

	lat = 0
	lng = -180
	p = NewPoint(lat, lng)
	assert.InDelta(t, 0, p.Latitude(), 0)
	assert.InDelta(t, 180, p.Longitude(), 0)

	lat = 0
	lng = -360
	p = NewPoint(lat, lng)
	assert.InDelta(t, 0, p.Latitude(), 0)
	assert.InDelta(t, 0, p.Longitude(), 0)

	lat = 0
	lng = 360
	p = NewPoint(lat, lng)
	assert.InDelta(t, 0, p.Latitude(), 0)
	assert.InDelta(t, 0, p.Longitude(), 0)

	lat = 0
	lng = 0
	p = NewPoint(lat, lng)
	assert.InDelta(t, 0, p.Latitude(), 0)
	assert.InDelta(t, 0, p.Longitude(), 0)

	lat = 0
	lng = 380
	p = NewPoint(lat, lng)
	assert.InDelta(t, 0, p.Latitude(), 0)
	assert.InDelta(t, 20, p.Longitude(), 0)

	lat = 0
	lng = -380
	p = NewPoint(lat, lng)
	assert.InDelta(t, 0, p.Latitude(), 0)
	assert.InDelta(t, -20, p.Longitude(), 0)

	var pnt *point
	p = pnt

	assert.InDelta(t, 0, p.Latitude(), 0)
	assert.InDelta(t, 0, p.Longitude(), 0)
}

func TestPoint_String(t *testing.T) {
	var pnt *point
	var s fmt.Stringer = pnt

	assert.Equal(t, "(0, 0)", s.String())

	s = &point{latitude: 45.12345678, longitude: 35.87654321}

	assert.Equal(t, "(45.12345678, 35.87654321)", s.String())
}

func TestPoint_Precision(t *testing.T) {
	p := NewPoint(45.1234567851, 35.8765432138)
	assert.InDelta(t, 45.12345679, p.Latitude(), 0)
	assert.InDelta(t, 35.87654321, p.Longitude(), 0)

	p = NewPoint(-45.1234567851, -35.8765432138)
	assert.InDelta(t, -45.12345679, p.Latitude(), 0)
	assert.InDelta(t, -35.87654321, p.Longitude(), 0)
}
