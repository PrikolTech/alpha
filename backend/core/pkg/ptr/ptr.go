package ptr

func To[T any](v T) *T {
	return &v
}

func Value[T any](p *T) T {
	if p == nil {
		var v T
		return v
	}
	return *p
}
