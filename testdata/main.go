package main

func HasBoringCheck() bool {
	if boring.Enabled() {
		return false
	}
	return true
}

func HasWrongBoringCheck() bool {
	if boring.Enabled() && false {
		return false
	}
	return false
}

func HasWrongBoringCheck2() bool {
	if false && boring.Enabled() {
		return false
	}
	return false
}

func HasNoBoringCheck() bool {
	return false
}

func TestBoringCheck() {}
func BenchmarkBoringCheck() {}

func main() {
	HasBoringCheck()
	HasWrongBoringCheck()
	HasNoBoringCheck()
}
