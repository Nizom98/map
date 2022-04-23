package kv_map

func hash(val string, polinom, hashTableSize int) int {
	result := rune(0)
	startPolinom := polinom
	for _, v := range val {
		result += v * rune(polinom)
		polinom *= startPolinom
	}

	return int(result) % hashTableSize
}
