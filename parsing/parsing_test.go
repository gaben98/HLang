package parsing

import "testing"

func TestTokenization(t *testing.T) {
	t.Log(TagLine("int sum = 0"))
	t.Log(TagLine("fmt.Println(SomeClass.SomeVal)"))
	t.Log(GTok("Char := \"'.'\";"))
	t.Error(LineTokenize("int sum = 0; fmt.Println(SomeClass.SomeVal);")) //, "{}()=.;\n", " 	"))
}

func TestGrammarization(t *testing.T) {
	L0 := MakeToken("'a'")
	L0.tag = "a-char"
	L0.next = Match{}
	vm := GVM{tokens: []string{"a"}, start: L0}
	/*L1 := Split{}
	L2 := MakeToken("'b'")
	L2.tag = "b-char"
	L3 := Split{}
	L4 := Match{}
	L0.next = L1
	L1.x = L0
	L1.y = L2
	L2.next = L3
	L3.x = L2
	L3.y = L4
	vm := GVM{tokens: []string{"a", "b"}, start: L0}*/
	t.Error(vm.Execute())
}
