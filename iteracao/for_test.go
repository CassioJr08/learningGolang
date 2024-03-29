package iteracao

import (
	"fmt"
	"testing"
)

func ExampleRepetir() {
	repeat := Repetir("a")
	fmt.Println(repeat)
	// Output: aaaaa
}

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a")
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticoes)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a")
	}
}
