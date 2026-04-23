package alias

type Alias struct {
	Typ string
}

func (p *Alias) alias() {}

type Type interface {
	alias()
}

type Chan struct {
	Elem Type
}

func (p *Chan) alias() {}

type Struct struct {
	Fields []Type
}

func (p *Struct) alias() {}

type Map struct {
	Key  Type
	Elem Type
}

func (p *Map) alias() {}

type Array struct {
	Elem Type
}

func (p *Array) alias() {}

type Slice struct {
	Elem Type
}

func (p *Slice) alias() {}

type Func struct {
	Params  []Type
	Results []Type
}

func (p *Func) alias() {}

type Named struct {
	Underlying Type
	Methods    map[string]*Func
}

func (p *Named) alias() {}

type Interface struct {
	Methods map[string]*Func
}

func (p *Interface) alias() {}
