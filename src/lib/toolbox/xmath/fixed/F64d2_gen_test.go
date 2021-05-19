// Code created from "fixed_test.go.tmpl" - don't edit by hand
//
// Copyright ©2016-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package fixed_test

import (
	"encoding/json"
	"testing"

	"lib/toolbox/xmath/fixed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gopkg.in/yaml.v2"
)

type embedded64d2 struct {
	Field fixed.F64d2
}

func TestConversion64d2(t *testing.T) {
	assert.Equal(t, "0.1", fixed.F64d2FromFloat64(0.1).String())
	assert.Equal(t, "0.2", fixed.F64d2FromFloat64(0.2).String())
	assert.Equal(t, "0.3", fixed.F64d2FromStringForced("0.3").String())
	assert.Equal(t, "-0.1", fixed.F64d2FromFloat64(-0.1).String())
	assert.Equal(t, "-0.2", fixed.F64d2FromFloat64(-0.2).String())
	assert.Equal(t, "-0.3", fixed.F64d2FromStringForced("-0.3").String())
	assert.Equal(t, "0.33", fixed.F64d2FromStringForced("0.3333").String())
	assert.Equal(t, "-0.33", fixed.F64d2FromStringForced("-0.3333").String())
	assert.Equal(t, "0.66", fixed.F64d2FromStringForced("0.6666").String())
	assert.Equal(t, "-0.66", fixed.F64d2FromStringForced("-0.6666").String())
	assert.Equal(t, "1", fixed.F64d2FromFloat64(1.004).String())
	assert.Equal(t, "1", fixed.F64d2FromFloat64(1.0049).String())
	assert.Equal(t, "1", fixed.F64d2FromFloat64(1.005).String())
	assert.Equal(t, "1", fixed.F64d2FromFloat64(1.009).String())
	assert.Equal(t, "-1", fixed.F64d2FromFloat64(-1.004).String())
	assert.Equal(t, "-1", fixed.F64d2FromFloat64(-1.0049).String())
	assert.Equal(t, "-1", fixed.F64d2FromFloat64(-1.005).String())
	assert.Equal(t, "-1", fixed.F64d2FromFloat64(-1.009).String())
	assert.Equal(t, "0.04", fixed.F64d2FromStringForced("0.0405").String())
	assert.Equal(t, "-0.04", fixed.F64d2FromStringForced("-0.0405").String())

	v, err := fixed.F64d2FromString("33.0")
	assert.NoError(t, err)
	assert.Equal(t, v, fixed.F64d2FromInt64(33))

	v, err = fixed.F64d2FromString("33.00000000000000000000")
	assert.NoError(t, err)
	assert.Equal(t, v, fixed.F64d2FromInt64(33))
}

func TestAddSub64d2(t *testing.T) {
	oneThird := fixed.F64d2FromStringForced("0.33")
	negTwoThirds := fixed.F64d2FromStringForced("-0.66")
	one := fixed.F64d2FromInt64(1)
	oneAndTwoThirds := fixed.F64d2FromStringForced("1.66")
	nineThousandSix := fixed.F64d2FromInt64(9006)
	ninetyPointZeroSix := fixed.F64d2FromStringForced("90.06")
	twelvePointThirtyFour := fixed.F64d2FromStringForced("12.34")
	two := fixed.F64d2FromInt64(2)
	assert.Equal(t, "0.99", (oneThird + oneThird + oneThird).String())
	assert.Equal(t, "0.67", (one - oneThird).String())
	assert.Equal(t, "-1.66", (negTwoThirds - one).String())
	assert.Equal(t, "0", (negTwoThirds - one + oneAndTwoThirds).String())
	assert.Equal(t, fixed.F64d2FromInt64(10240), fixed.F64d2FromInt64(1234)+nineThousandSix)
	assert.Equal(t, "10240", (fixed.F64d2FromInt64(1234) + nineThousandSix).String())
	assert.Equal(t, fixed.F64d2FromStringForced("102.4"), twelvePointThirtyFour+ninetyPointZeroSix)
	assert.Equal(t, "102.4", (twelvePointThirtyFour + ninetyPointZeroSix).String())
	assert.Equal(t, "-1.5", (fixed.F64d2FromFloat64(0.5) - two).String())
}

func TestMulDiv64d2(t *testing.T) {
	pointThree := fixed.F64d2FromStringForced("0.3")
	negativePointThree := fixed.F64d2FromStringForced("-0.3")
	assert.Equal(t, "0.33", fixed.F64d2FromInt64(1).Div(fixed.F64d2FromInt64(3)).String())
	assert.Equal(t, "-0.33", fixed.F64d2FromInt64(1).Div(fixed.F64d2FromInt64(-3)).String())
	assert.Equal(t, "0.1", pointThree.Div(fixed.F64d2FromInt64(3)).String())
	assert.Equal(t, "0.9", pointThree.Mul(fixed.F64d2FromInt64(3)).String())
	assert.Equal(t, "-0.9", negativePointThree.Mul(fixed.F64d2FromInt64(3)).String())
}

func TestTrunc64d2(t *testing.T) {
	assert.Equal(t, fixed.F64d2FromInt64(0), fixed.F64d2FromStringForced("0.3333").Trunc())
	assert.Equal(t, fixed.F64d2FromInt64(2), fixed.F64d2FromStringForced("2.6789").Trunc())
	assert.Equal(t, fixed.F64d2FromInt64(3), fixed.F64d2FromInt64(3).Trunc())
	assert.Equal(t, fixed.F64d2FromInt64(0), fixed.F64d2FromStringForced("-0.3333").Trunc())
	assert.Equal(t, fixed.F64d2FromInt64(-2), fixed.F64d2FromStringForced("-2.6789").Trunc())
	assert.Equal(t, fixed.F64d2FromInt64(-3), fixed.F64d2FromInt64(-3).Trunc())
}

func TestJSON64d2(t *testing.T) {
	for i := int64(-25000); i < 25001; i += 13 {
		testJSON64d2(t, fixed.F64d2FromInt64(i))
	}
	testJSON64d2(t, fixed.F64d2FromInt64(1844674407371259000))
}

func testJSON64d2(t *testing.T, v fixed.F64d2) {
	t.Helper()
	e1 := embedded64d2{Field: v}
	data, err := json.Marshal(&e1)
	assert.NoError(t, err)
	var e2 embedded64d2
	err = json.Unmarshal(data, &e2)
	assert.NoError(t, err)
	require.Equal(t, e1, e2)
}

func TestYAML64d2(t *testing.T) {
	for i := int64(-25000); i < 25001; i += 13 {
		testYAML64d2(t, fixed.F64d2FromInt64(i))
	}
	testYAML64d2(t, fixed.F64d2FromInt64(1844674407371259000))
}

func testYAML64d2(t *testing.T, v fixed.F64d2) {
	t.Helper()
	e1 := embedded64d2{Field: v}
	data, err := yaml.Marshal(&e1)
	assert.NoError(t, err)
	var e2 embedded64d2
	err = yaml.Unmarshal(data, &e2)
	assert.NoError(t, err)
	require.Equal(t, e1, e2)
}
