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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash_Nil(t *testing.T) {

	var hsh *hash
	var h Hash = hsh

	assert.Equal(t, uint64(0), uint64(h.Size()))
	assert.Equal(t, uint64(0), uint64(h.Bits()))
	assert.Equal(t, uint64(0), uint64(h.Value()))
	assert.Equal(t, "", h.String())
}

func TestHash_Getters(t *testing.T) {

	var bits uint64 = 0xaf777bbb00000000

	h := NewHash(bits, 32)

	assert.Equal(t, uint8(32), h.Size())
	assert.Equal(t, bits, h.Bits())
	assert.Equal(t, bits>>32, h.Value())
	assert.Equal(t, "pxvrrf", h.String())
}
