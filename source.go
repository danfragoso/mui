package mui

type Source struct {
	rawSource  string
	sourceLen  int
	currentIdx int
}

func createSource(raw string) *Source {
	return &Source{
		rawSource:  raw,
		sourceLen:  len(raw),
		currentIdx: 0,
	}
}

func (source *Source) getChar() rune {
	source.currentIdx++
	return rune(source.rawSource[source.currentIdx-1])
}

func (source *Source) hasChar() bool {
	if source.currentIdx == source.sourceLen {
		return false
	}

	return true
}
