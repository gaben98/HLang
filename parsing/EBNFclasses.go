package parsing

import "regexp"

type gclass interface {
	MatchTokens([]string) bool
}

//GExpr represents an atom
type GExpr struct {
	matcher regexp.Regexp
}

//MakeGExpr creates and instance of GExpr
func MakeGExpr(regex string) GExpr {
	return GExpr{matcher: *regexp.MustCompile("^" + regex[1:len(regex)-1] + "$")}
}

//MatchTokens returns true iff the tokens array only holds one token and that token exactly matches the contained regex
func (exp *GExpr) MatchTokens(tokens []string) bool {
	if len(tokens) != 1 {
		return false
	}
	return exp.matcher.MatchString(tokens[0])
}

//Alternation represents an option of subexpressions
type Alternation struct {
	subexpressions []gclass
}

//MakeAlternation returns an instance of Alternation
func MakeAlternation(subexpr []gclass) Alternation {
	return Alternation{subexpressions: subexpr}
}

//MatchTokens will match the tokens if and only if one of it's subexpressions matches the token set exactly
func (alt *Alternation) MatchTokens(tokens []string) bool {
	for _, subexpr := range alt.subexpressions {
		if subexpr.MatchTokens(tokens) {
			return true
		}
	}
	return false
}
