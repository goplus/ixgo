/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
 * Licensed under the Apache License, Version 2.0.
 */

// Package mixedrunner enables optional mixed numeric and boolean runners.
// Import it for side effects:
//
//	import _ "github.com/goplus/ixgo/mixedrunner"
package mixedrunner

import "github.com/goplus/ixgo"

func init() {
	ixgo.RegisterMixedRunner(MakeRunner)
}
