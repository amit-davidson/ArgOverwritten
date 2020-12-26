package testdata

func closeBody(body int) {
	body = 1 // want `"body" overwrites func parameter "body"`
}

func main() {
	closeBody(1)
}
