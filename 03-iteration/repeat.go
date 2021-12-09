package iteration

func Repeat(symbol string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += symbol
	}
	return repeated
}
