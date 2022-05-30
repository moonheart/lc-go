package main

func replaceSpace(s string) string {
	var byts []byte
	for _, i := range s {
		if i == ' ' {
			byts = append(byts, '%', '2', '0')
		} else {
			byts = append(byts, byte(i))
		}
	}
	return string(byts)
}
