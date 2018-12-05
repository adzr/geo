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
	"math"

	"github.com/adzr/mathex"
)

const (
	// DecimalPlaces is the number of decimal places considered in the geo-location point latlng values, to indicate the
	// precision of the coordinates.
	DecimalPlaces = 8
	// RoundOn is the decimal value considered when rounding the geo-location point latlng values.
	RoundOn = 0.5
	// TotalLongitude is a constant representing the maximum longitude value, ignoring the negative values.
	TotalLongitude float64 = 360.0
	// HalfLongitude is a constant representing the maximum longitude value, considering the negative values.
	HalfLongitude float64 = 180.0
	// NorthPoleLat is a constant representing the maximum latitude value, always at the center of the north pole.
	NorthPoleLat float64 = 90.0
	// SouthPoleLat is a constant representing the maximum latitude value, always at the center of the south pole.
	SouthPoleLat float64 = -90.0
)

// Point is a geo-location point represented by latitude and longitude (latlng) values.
type Point interface {
	// Latitude is an angle which ranges from 0° at the Equator to +90° at the North Pole or -90° at the South Pole.
	Latitude() float64
	// Longitude is an angle which ranges from 0° at the Prime Meridian to +180° eastward and −180° westward.
	Longitude() float64
}

type point struct {
	latitude  float64
	longitude float64
}

func (p *point) String() string {
	return fmt.Sprintf("(%v, %v)", p.Latitude(), p.Longitude())
}

func (p *point) Latitude() float64 {
	if p != nil {
		return p.latitude
	}
	return 0.0
}

func (p *point) Longitude() float64 {
	if p != nil {
		return p.longitude
	}
	return 0.0
}

// NewPoint creates a new geo-location point instance, given the latlng values.
// The function auto-normalizes the values to void over-calculations.
// It keeps the latitude value always capped between -90 and +90.
// It sets the longitude to 0 incase if the latitude value is exactly equals to -90 or +90,
// since the longitude has no meaning if the latitude is at one of the two poles.
// It also keeps the longitude value always capped between almost -180 and +180.
func NewPoint(latitude float64, longitude float64) Point {

	// create the point.
	p := &point{}

	// normalize latitude.
	if latitude > 0 {
		p.latitude = math.Min(latitude, NorthPoleLat)
	} else {
		p.latitude = math.Max(latitude, SouthPoleLat)
	}

	p.latitude = mathex.Round(p.latitude, DecimalPlaces, RoundOn)

	// normalize longitude.
	if math.Abs(p.latitude) == NorthPoleLat {
		p.longitude = 0
	} else {
		remainder := math.Mod(longitude, TotalLongitude)

		if remainder <= -HalfLongitude {
			remainder += TotalLongitude
		} else if remainder > HalfLongitude {
			remainder -= TotalLongitude
		}

		p.longitude = remainder
	}

	p.longitude = mathex.Round(p.longitude, DecimalPlaces, RoundOn)

	return p
}
