package twotoone

func TwoToOne(s1 string, s2 string) string {
	// concatenate two strings to one
	s3 := s1 + s2
	var seen []byte = make([]byte, 26)
	// hold a counter to know, if we have seen all letters, then we can stop parsing the string.
	var alphabetCounter int
	for _, char := range s3 {
		if seen[char-97] == 0 {
			seen[char-97] = byte(char)
			alphabetCounter += 1
			// if we have seen all letters, then we can return the whole alphabet
			if alphabetCounter > 26 {
				return "abcdefghijklmnopqrstuvwxyz"
			}
		}
	}
	// construct sorted string
	var result string
	for i := 0; i < 26; i++ {
		if seen[i] != 0 {
			result += string(seen[i])
		}
	}
	return result
}
