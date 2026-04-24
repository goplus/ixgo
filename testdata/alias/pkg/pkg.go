package pkg

import (
	"github.com/goplus/ixgo/testdata/alias/msg"
)

type Int = int
type myInt = int
type M = map[int]string
type S = []int
type Msg = msg.Message

const (
	A1 Int = 100
	A2 int = 200
)

var (
	V1 Int = 100
	V2     = M{0: "hello"}
	V3     = []Int{100, 200}
	V4     = struct {
		X Int
		Y [2]Int
		Z chan Int
		M int
		S func(Int)
		P *struct{ Int }
	}{}
)

func Demo1(Int, int) Int {
	return 0
}

func Demo2(myInt, int) Int {
	return 0
}

func Demo3(any) {}

func Demo4(int, M) {}

func Demo5(int, map[Int]string) {}

func Demo6(string, S, Msg) {}

func Demo7(string, ...Int) {}

type Addable interface {
	Add(Int, int) int
}

type Point struct {
	X Int
	Y int
}

func (p *Point) Demo1(Int, int) Int {
	return 0
}

func (p *Point) Demo2(myInt, int) Int {
	return 0
}

func (p *Point) Demo3(any) {}

func (p *Point) Demo4(int, M) {}

func (p *Point) Demo5(int, map[Int]msg.Message) (int, msg.Message) {
	return 0, ""
}
