package hash

// WithPolinom вычисляет полиномиальную хеш и возврашает остаток от деления hashTableSize.
// val - строка полиномиальная хеш которой вычисляется.
// polinom - произволольное число (полином).
// hashTableSize - для вычисление остатка от деления.
//!!!Желательно, чтобы числа polinom и hashTableSize были больше нуля!!!
func WithPolinom(val string, polinom, hashTableSize uint) uint {
	if len(val) == 0 || hashTableSize == 0 {
		return 0
	} 
	valRunes, polinomRune := []rune(val), rune(polinom)
	sum := valRunes[0]
	startPolinom := polinomRune
	valRunes = valRunes[1:]

	for _, v := range valRunes {
		sum += v * polinomRune
		polinomRune *= startPolinom
	}

	return uint(sum) % hashTableSize
}