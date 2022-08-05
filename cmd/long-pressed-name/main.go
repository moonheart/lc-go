package long_pressed_name

func isLongPressedName(name string, typed string) bool {
	i := 0
	for j := 0; j < len(typed); j++ {
		if i < len(name) && typed[j] == name[i] {
			i++
		} else if j > 0 && typed[j] == typed[j-1] {

		} else {
			return false
		}
	}
	return i == len(name)
}
