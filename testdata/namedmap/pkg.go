package namedmap

type Values map[string]Element

type Element struct{}

func (Element) Back() Values {
	return nil
}
