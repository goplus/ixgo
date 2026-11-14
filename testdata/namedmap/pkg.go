package namedmap

type Values map[string]string

type Template struct{}

func (t *Template) Match(s string) Values {
	return Values{"k": s}
}
