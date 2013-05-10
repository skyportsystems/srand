// Copyright 2013 by Dobrosław Żybort. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package srand simplify choosing random values, rand helper.
package srand

import (
	"math/rand"
	"time"
)

//=============================================================================

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Int returns a non-negative random number from min to max
// (0 <= min <= nr <= max).
// Will return 0 if max < 0
func Int(min, max int) int {
	var choice int

	if max < 0 {
		return 0
	}
	if min < 0 {
		min = 0
	}

	if min < max {
		choice = min + rand.Intn(max-min+1)
	} else {
		choice = min
	}
	return choice
}

//func IntChoice(min, max int) int {
//	return Int(min, max int)
//}

// IntMany returns a slice of non-negative random numbers from min to max.
func IntMany(min, max, quantity int) (out []int) {
	for i := 0; i < quantity; i++ {
		out = append(out, Int(min, max))
	}
	return out
}

// IntSample returns a slice of non-negative random, unique numbers
// from min to max. If the quantity is bigger than max - min sample size
// will be the size of max - min.
//func IntSample(min, max, quantity int) (sample []int) {
//	for i := 0; i < quantity; i++ {
//		sample = append(sample, Int(min, max))
//	}
//	return sample
//}
