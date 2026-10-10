// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pcommon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// string_value_strindex existed in the proto but no public accessor reached it,
// so profiles data using it could not be read or written through pdata (#15450).
func TestValueStrIndexRoundTrip(t *testing.T) {
	v := NewValueStrIndex(7)

	assert.Equal(t, ValueTypeStrIndex, v.Type())
	assert.Equal(t, int32(7), v.StrIndex())
	assert.Equal(t, "StrIndex", v.Type().String())
}

func TestValueSetStrIndexReplacesOtherTypes(t *testing.T) {
	v := NewValueStr("inline")
	v.SetStrIndex(3)

	assert.Equal(t, ValueTypeStrIndex, v.Type())
	assert.Equal(t, int32(3), v.StrIndex())
	assert.Empty(t, v.Str(), "the old inline string must not survive the switch")
}

func TestValueStrIndexEqual(t *testing.T) {
	assert.True(t, NewValueStrIndex(2).Equal(NewValueStrIndex(2)))
	assert.False(t, NewValueStrIndex(2).Equal(NewValueStrIndex(3)))
	// Different representations of a string are not the same value: one is an
	// index, the other is the text.
	assert.False(t, NewValueStrIndex(2).Equal(NewValueStr("2")))
}

func TestValueStrIndexAsRaw(t *testing.T) {
	assert.Equal(t, int32(5), NewValueStrIndex(5).AsRaw())
}

func TestValueStrIndexOnOtherTypeIsZero(t *testing.T) {
	assert.Equal(t, int32(0), NewValueStr("x").StrIndex())
}
