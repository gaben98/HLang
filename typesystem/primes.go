package typesystem

func erastosthenes() func() int64 {
	pGen := prime()

	return func() int64 {
		num, nGen := pGen()
		pGen = nGen
		return num
	}
}

type generator func() (int64, generator)

func from(a int64) generator {
	return func() (int64, generator) {
		return a, from(a + 1)
	}
}

func filter(pred func(int64) bool, gen generator) generator {
	num, nGen := gen()
	if pred(num) {
		return func() (int64, generator) {
			return num, filter(pred, nGen)
		}
	}
	return filter(pred, nGen)
}

//sift returns a func lacking any multiples of n
func sift(n int64, gen generator) generator {
	return filter(func(x int64) bool { return x%n != 0 }, gen)
}

func phelp(gen generator) generator {
	num, _ := gen()
	return func() (int64, generator) {
		return num, phelp(sift(num, gen))
	}
}

func prime() generator {
	return func() (int64, generator) {
		return 2, phelp(sift(2, from(2)))
	}
}
