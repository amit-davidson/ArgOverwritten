package testdata

func body(a, b int, c int) {
	a = 5 // want `"a" overwrites func parameter "a"`
	_ = b
	c = 3 // want `"c" overwrites func parameter "c"`
}

func main() {
	body(1, 2, 3)
}
