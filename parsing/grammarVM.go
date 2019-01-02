package parsing

import "regexp"

//GVM is a grammar virtual machine, for evaluating if a text matches a grammar
type GVM struct {
	tokens []string //tokenized text
	start  Inst     //instruction to start at
}

type threadcontext struct {
	ip    Inst
	match []string
}

func (tc *threadcontext) dupe() threadcontext {
	tmp := make([]string, len(tc.match))
	copy(tmp, tc.match)
	return threadcontext{ip: tc.ip, match: tmp}
}

//Execute runs the gvm and returns a classification array
func (vm *GVM) Execute() []string {
	clist := make([]threadcontext, 0) //fresh threads
	nlist := make([]threadcontext, 0) //done threads
	clist = append(clist, threadcontext{ip: vm.start})
	for _, token := range vm.tokens {
		for _, tc := range clist {
			if tc.ip.Eval(&tc, token, &clist, &nlist) {
				return tc.match
			}
		}
		clist = make([]threadcontext, len(nlist))
		copy(clist, nlist)
		nlist = make([]threadcontext, 0)
	}
	return nil
}

//Inst represents a GVM instruction
type Inst interface {
	Eval(*threadcontext, string /*token*/, *[]threadcontext /*pointer to clist*/, *[]threadcontext /*pointer to nlist*/) bool //returns if terminal match
}

//Token represents an instruction that matches a token to a regular expression
type Token struct {
	matcher regexp.Regexp
	next    Inst
	tag     string
}

//MakeToken makes an instance of ExprMatch
func MakeToken(regex string) Token {
	return Token{matcher: *regexp.MustCompile("^" + regex[1:len(regex)-1] + "$")}
}

//Eval updates the vm state to reflect a match of a single token
func (atom Token) Eval(tc *threadcontext, token string, clist, nlist *[]threadcontext) bool {
	if atom.matcher.MatchString(token) {
		next := tc.dupe()
		next.ip = atom.next
		next.match = append(next.match, atom.tag)
		ctx := (append(*nlist, next))
		nlist = &ctx
	}
	return false
}

//Match represents a successful match instruction
type Match struct{}

//Eval terminates successfully
func (match Match) Eval(tc *threadcontext, token string, clist, nlist *[]threadcontext) bool {
	return true
}

//Jmp represents a jump instruction
type Jmp struct {
	next Inst
}

//Eval jumps the instruction pointer of the vm to the one specified, then resumes execution
func (jmp Jmp) Eval(tc *threadcontext, token string, clist, nlist *[]threadcontext) bool {
	next := tc.dupe()
	next.ip = jmp.next
	ctx := append(*clist, next)
	clist = &ctx
	return false
}

//Split represents the split instruction, which creates a new VM thread
type Split struct {
	x, y Inst
}

//Eval computes the result of a split operation
func (splt Split) Eval(tc *threadcontext, token string, clist, nlist *[]threadcontext) bool {
	t1 := tc.dupe()
	t2 := tc.dupe()
	t1.ip = splt.x
	t2.ip = splt.y
	ctx := append(*clist, t1, t2)
	clist = &ctx
	return false
}
