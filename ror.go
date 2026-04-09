package ascon

func ROR(in []string, count int) []string {
	for i := 0; i < count; i++ {
		tail := in[len(in)-1]
		remaining := in[:len(in)-1]
		in = append([]string{tail}, remaining...)
	}
	return in
}
