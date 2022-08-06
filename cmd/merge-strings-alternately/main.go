package merge_strings_alternately

func mergeAlternately(word1 string, word2 string) string {
	res := []byte{}

	for i, j := 0, 0; i < len(word1) || j < len(word2); i, j = i+1, j+1 {
		if i < len(word1) {
			res = append(res, word1[i])
		}
		if j < len(word2) {
			res = append(res, word2[j])
		}
	}
	return string(res)
}
