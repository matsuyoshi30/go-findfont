// +build wasm

// This file for resolving build failure GOARCH=wasm when using this module

// Original License is below

// Copyright 2016 Florian Pigorsch. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package findfont

func getFontDirectories() (paths []string) {
	return []string{}
}
