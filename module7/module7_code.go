package module7

//go:generate goimports module7_code.go
func content() {
	fmt.Println(runtime.GOOS)
}
