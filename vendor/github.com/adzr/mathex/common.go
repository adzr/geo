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

package mathex

import (
	"math"
)

// Round takes a float64 val and rounds it to the nearest decimal places taking into account the roundOn value and returning the new float64 result back.
//
// Example:
// Round(1.653265718543275, 5, 0.5)	=> 1.65327
// Round(0.62, 1, 0.5)			=> 0.6
// Round(0.62, 1, 0.2)			=> 0.7
// Round(0.62, 0, 0.5)			=> 1
// Round(92, -1, 0.5)			=> 90
// Round(92, -1, 0.1)			=> 100
// Round(-1.653265718543275, 5, 0.5)	=> -1.65327
// Round(-0.62, 1, 0.5)			=> -0.6
// Round(-0.62, 1, 0.2)			=> -0.7
// Round(-0.62, 0, 0.5)			=> -1
// Round(-92, -1, 0.5)			=> -90
// Round(-92, -1, 0.1)			=> -100
//
// The 'places' parameter indicates how many places the decimal point will move.
// The 'roundOn' parameter indicates on which value the round will happen.
//
// Inspired by the following gist:
//
// https://gist.github.com/DavidVaini/10308388
// https://gist.github.com/pelegm/c48cff315cd223f7cf7b
func Round(val float64, places int, roundOn float64) float64 {

	pow := math.Pow(10, float64(places))
	signed := math.Signbit(val)
	scaled := pow * math.Abs(val)
	_, div := math.Modf(scaled)

	var result float64

	if div >= roundOn {
		result = math.Ceil(scaled) / pow
	} else {
		result = math.Floor(scaled) / pow
	}

	if signed {
		return result * -1
	}

	return result
}
