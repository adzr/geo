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
Package mathex provides some common math functions that complements the standard math package.

Brief

This library is to wrap any math functions that's missing from the standard Go math package.

Usage

	$ go get -u github.com/adzr/mathex

Then, import the package:

  import (
    "github.com/adzr/mathex"
  )

Example

  // This will output 1.65326
  fmt.Printf("%v\n", mathex.Round(1.653264418543275, 5, 0.5))

  // Same with a negative number, this will output -1.65326
  fmt.Printf("%v\n", mathex.Round(-1.653264418543275, 5, 0.5))

*/
package mathex
