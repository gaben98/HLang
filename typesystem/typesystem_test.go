package typesystem

import "testing"

func TestPrimeGen(t *testing.T) {
	pnxt := erastosthenes()
	nums := make([]int64, 0)
	for i := 0; i < 5; i++ {
		nums = append(nums, pnxt())
	}
	t.Error(nums)
}

func TestObjectTree(t *testing.T) {
	hInt := Object.Define("integer")
	hFloat := Object.Define("float")
	subInt := hInt.Define("subint")
	subInt2 := hInt.Define("short")
	t.Log(PrintObjectTree(&Object))
	t.Log(subInt.Is(Object))
	t.Log(subInt.Is(*hFloat))
	t.Log(hInt.Is(*subInt2))
	t.Error(subInt2.Is(*hInt))
}

func TestTuples(t *testing.T) {
	htpl := DefineTuple([]*IHType{hptr(HInt), hptr(HChar)})
	htpl2 := DefineTuple([]*IHType{hptr(&Object), hptr(HChar)})
	htpl3 := DefineTuple([]*IHType{hptr(&Object), hptr(HChar), hptr(HInt)})
	t.Log(htpl.Is(htpl2))
	t.Log(htpl2.Is(htpl3))
	t.Error(htpl2.Is(htpl))
}
