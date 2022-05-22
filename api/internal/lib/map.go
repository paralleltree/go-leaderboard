package lib

func Map[TSrc any, TDest any](src []TSrc, selector func(TSrc) TDest) []TDest {
	res := make([]TDest, 0, len(src))
	for _, item := range src {
		res = append(res, selector(item))
	}
	return res
}
