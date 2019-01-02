package parsing

//Tag is an enum for a tagged token
type Tag int

const (
	//TType is the tag for a type token
	TType Tag = iota
	//TVar is the tag for a variable identifier token
	TVar Tag = iota
	//TVal is the tag for a literal value
	TVal Tag = iota
	//TAssign is the tag for =
	TAssign Tag = iota
)

//HToken represents a tagged token
type HToken struct {
	token string
	tag   Tag
}

//TagLine classifies tokens within a line of HLang code
func TagLine(line string) []HToken {
	tokens := Tokenize(line, ",.=()", " 	")
	htokens := make([]HToken, 0)
	for i, token := range tokens {
		t := TType
		if i == 1 {
			t = TVar
		}
		if token == "=" {
			t = TAssign
		}
		if '0' <= token[0] && token[0] <= '9' {
			t = TVal
		}
		htokens = append(htokens, HToken{token: token, tag: t})
	}
	return htokens
}
