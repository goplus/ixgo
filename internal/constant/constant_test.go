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

package constant

import (
	"go/constant"
	"go/token"
	"testing"
)

func TestExactConstant(t *testing.T) {
	const c = "3.14159265358979323846264338327950288419716939937510582097494459"
	x := constant.MakeFromLiteral(c, token.FLOAT, 0)
	if v := ExactConstant(x); v != c {
		t.Fatalf("%v", v)
	}
}

func TestExactConstantEx1(t *testing.T) {
	const c = "3.14159265358979323846264338327950288419716939937510582097494459"
	x := constant.MakeFromLiteral(c, token.FLOAT, 0)
	if v, exact := ExactConstantEx(x, true); v != c || exact != true {
		t.Fatalf("%v", v)
	}
}

func TestExactConstantEx2(t *testing.T) {
	const c = "1.401298464324817e-45"
	x := constant.BinaryOp(constant.MakeFromLiteral("1", token.INT, 0), token.QUO, constant.MakeFromLiteral("713623846352979940529142984724747568191373312", token.INT, 0))
	if v, exact := ExactConstantEx(x, true); v != c || exact != false {
		t.Fatalf("%v", v)
	}
}

func TestExactConstantEx3(t *testing.T) {
	const c = "1.01"
	x := constant.BinaryOp(constant.MakeFromLiteral("101", token.INT, 0), token.QUO, constant.MakeFromLiteral("100", token.INT, 0))
	if v, exact := ExactConstantEx(x, true); v != c || exact != true {
		t.Fatalf("%v", v)
	}
}

func TestExactConstantEx4(t *testing.T) {
	const c = "3.14e+100"
	x := constant.MakeFromLiteral("3.14e100", token.FLOAT, 0)
	if v, exact := ExactConstantEx(x, true); v != c || exact != true {
		t.Fatalf("%v", v)
	}
}

func TestExactConstantEx5(t *testing.T) {
	const c = "(11/10 + 5i)"
	x := constant.BinaryOp(constant.MakeFromLiteral("1.1", token.FLOAT, 0), token.ADD, constant.MakeImag(constant.MakeInt64(5)))
	if v, exact := ExactConstantEx(x, true); v != c || exact != true {
		t.Fatalf("%v", v)
	}
}

func TestExactConstantEx6(t *testing.T) {
	const c = "100000"
	x := constant.MakeFromLiteral("100000.0", token.FLOAT, 0)
	if v, exact := ExactConstantEx(x, true); v != c || exact != true {
		t.Fatalf("%v", v)
	}
}

func TestExactConstantEx7(t *testing.T) {
	const c = "1.2e+6"
	x := constant.MakeFromLiteral("1200000.0", token.FLOAT, 0)
	if v, exact := ExactConstantEx(x, true); v != c || exact != true {
		t.Fatalf("%v", v)
	}
}

func TestExactConstantEx8(t *testing.T) {
	const c = "1e-5"
	x := constant.MakeFromLiteral("0.00001", token.FLOAT, 0)
	if v, exact := ExactConstantEx(x, true); v != c || exact != true {
		t.Fatalf("%v", v)
	}
}

func TestExactConstantEx9(t *testing.T) {
	const c = "0.0001"
	x := constant.MakeFromLiteral("0.0001", token.FLOAT, 0)
	if v, exact := ExactConstantEx(x, true); v != c || exact != true {
		t.Fatalf("%v", v)
	}
}

func TestExactConstantEx10(t *testing.T) {
	const c = "0.1"
	x := constant.MakeFromLiteral("0.1", token.FLOAT, 0)
	if v, exact := ExactConstantEx(x, true); v != c || exact != true {
		t.Fatalf("%v", v)
	}
}
