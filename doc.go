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

/*
Package geo is a library package that provide utilities to process calculations on geographical points, hashes & boundaries.

Brief

The package has a few features that might be handy, including:
- Geo-point latlng normalization.
- Geo-point 8 decimal places precision handling.
- Measuring the distance between two given geo-points.
- Calculating the geo-hash of a given geo-point with a defined precision.
- Calculating the geo-point representing a given geo-hash.
- Calculating the needed geo-hash precision given the radius in kilometers.
- Calculating the neighbour geo-hashes of the given geo-hash.
- Creating normalized boundary boxes given two geo-points.

Usage

  $ go get -u github.com/adzr/geo

Then, import the package:

  import (
    "github.com/adzr/geo"
  )

Examples

Creating a normalized geo-point:

  // lat = 130
  // lng = 200
  // This code will output (90, 0)
  // because there is no lat above 90
  // and lng is always 0 when lat is 90.
  fmt.Printf("%v\n", geo.NewPoint(130, 200))

Geo-point precision is rounded to 8 places on 0.5:

  // lat = -45.1234567851
  // lng = -35.8765432138
  // This code will output (-45.12345679, -35.87654321)
  fmt.Printf("%v\n", geo.NewPoint(-45.1234567851, -35.8765432138))

Distance measurement is done using Haversine formula (https://en.wikipedia.org/wiki/Haversine_formula):

  // latlng1 => (50.432356, 83.873793)
  // latlng2 => (58.124521, 63.735753)
  // This code will output the distance in kilometers which is around 1553.
  fmt.Printf("%v\n", geo.GetDistance(geo.NewPoint(50.432356, 83.873793), geo.NewPoint(58.124521, 63.735753))

Now let's say we want to get the geo-hash of a given geo-point within 1km precision:

  // latlng => (50.432356, 83.873793)
  // This code will output "vbgw" which is the hash bits encoded into a base32 string.
  fmt.Printf("%v\n", geo.GetHash(geo.NewPoint(50.432356, 83.873793), geo.GetNeededPrecision(1)))

Probably, it's easy now to go playing around with the rest of the API, I hope the apidocs explain it well.

*/
package geo
