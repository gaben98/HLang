package parsing

//Tokenize takes a string and tokenizes on a set of delimiters.  Delims in tdelims will be added as tokens
func Tokenize(s, tdelims, delims string) []string {
	buffer := ""
	tokens := make([]string, 0)
	for _, r := range s {
		if strcontains(tdelims, r) {
			if len(buffer) > 0 {
				tokens = append(tokens, buffer)
				buffer = ""
			}
			tokens = append(tokens, string(r))
		} else if strcontains(delims, r) {
			if len(buffer) > 0 {
				tokens = append(tokens, buffer)
				buffer = ""
			}
		} else {
			buffer += string(r)
		}
	}
	if len(buffer) > 0 {
		tokens = append(tokens, buffer)
		buffer = ""
	}
	return tokens
}

func strcontains(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}

//LineTokenize splits a string into individual lines
func LineTokenize(s string) []string {
	return Tokenize(s, "{}", ";")
}

//GTok tokenizes a grammar file line
func GTok(line string) []string {
	return gtokenize(line, "()*+? |;")
}

//delims are characters to tokenize on, but can be escaped
func gtokenize(line string, delims string) []string {
	buffer := ""
	isEscaping := false
	inQuote := false
	tokens := make([]string, 0)
	for _, c := range line {
		if strcontains(delims, c) && !isEscaping && !inQuote {
			if len(buffer) > 0 {
				tokens = append(tokens, buffer)
				buffer = ""
			}
			if c != ' ' {
				tokens = append(tokens, string(c))
			}
		} else {
			if c == '\\' {
				isEscaping = !isEscaping
			} else if c == '"' && !isEscaping {
				inQuote = !inQuote
			} else {
				isEscaping = false
			}
			buffer += string(c)
		}
	}
	if len(buffer) > 0 {
		tokens = append(tokens, buffer)
	}
	return tokens
}
