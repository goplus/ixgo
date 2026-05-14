/*
 * Copyright (c) 2022 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ixgo

import "strings"

const (
	errDeclaredNotUsed     = "declared and not used"
	errImportedNotUsed     = "imported and not used"
	errAppendOutOfRange    = "len out of range"
	errSliceToArrayPointer = "cannot convert slice with length %v to array or pointer to array with length %v"
)

func hasTypesNotUsedError(msg string) bool {
	return strings.HasPrefix(msg, errDeclaredNotUsed+":") || strings.HasSuffix(msg, errImportedNotUsed)
}

func isTypesDeclaredNotUsed(msg string) (string, bool) {
	if strings.HasPrefix(msg, errDeclaredNotUsed+":") {
		return msg[len(errDeclaredNotUsed)+2:], true
	}
	return "", false
}
