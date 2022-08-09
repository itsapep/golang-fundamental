package intro

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	expected := "Selamat Pagi!"
	actual := "Selamat Pagi!"
	if actual != expected {
		t.Errorf("Greeting() = %q, want %q", actual, expected)
	}
}

func TestGreetingOther(t *testing.T) {
	expected := "Selamat Pagi!"
	actual := "Selamat Pagi!"
	if expected != actual {
		t.Errorf("Greeting() = %q, want %q", actual, expected)
	}
}

func TestSayHi(t *testing.T) {
	expected := "Hai salam kenal!\n"
	actual := SayHi()
	if actual != expected {
		t.Errorf("SayHi() = %q, want %q", actual, expected)
	}
}

func BenchmarkGreetingTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greeting()
	}
}
