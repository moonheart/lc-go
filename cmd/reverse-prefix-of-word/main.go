package reverse_prefix_of_word

func reversePrefix(word string, ch byte) string {
	bs := []byte(word)
	idx := -1
	for i, b := range bs {
		if b == ch {
			idx = i
			break
		}
	}
	if idx == -1 {
		return word
	}
	for i := 0; i <= idx/2; i++ {
		bs[i], bs[idx-i] = bs[idx-i], bs[i]
	}
	return string(bs)
}
