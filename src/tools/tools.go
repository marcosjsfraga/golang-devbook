package tools

// CheckError retorna erros
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
