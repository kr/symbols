package symbols

import (
	"encoding/binary"
	"testing"
)

var sink interface{}

func BenchmarkMakeString(b *testing.B) {
	sizes(b, func(b *testing.B, name []byte) {
		var v string
		for i := 0; i < b.N; i++ {
			// Make just a string.
			v = string(name)
		}
		sink = v
	})
}

func BenchmarkMakeSame(b *testing.B) {
	sizes(b, func(b *testing.B, name []byte) {
		// Preallocate the names and reset the timer
		// so we don't measure the time it took to make
		// the string values.
		names := make([]string, b.N)
		for i := range names {
			names[i] = string(name)
		}
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			// Make just a Symbol (using a new string for the sane.
			Make(names[i])
		}
	})
}

func BenchmarkMakeDiff(b *testing.B) {
	sizes(b, func(b *testing.B, name []byte) {
		// Preallocate the names and reset the timer
		// so we don't measure the time it took to make
		// the string values.
		names := make([]string, b.N)
		for i := range names {
			buf := make([]byte, len(name))
			binary.BigEndian.PutUint64(buf, uint64(i))
			names[i] = string(buf)
		}
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			// Make just a Symbol (using a new string for the sane.
			Make(names[i])
		}
	})
}

func BenchmarkSymbolNE(b *testing.B) {
	sizes(b, func(b *testing.B, name []byte) {
		x := Make(string(name))
		name[0] = ^name[0]
		y := Make(string(name))
		var v bool
		for i := 0; i < b.N; i++ {
			v = x == y
		}
		sink = v
	})
}

func BenchmarkSymbolEQ(b *testing.B) {
	sizes(b, func(b *testing.B, name []byte) {
		x := Make(string(name))
		y := Make(string(name))
		var v bool
		for i := 0; i < b.N; i++ {
			v = x == y
		}
		sink = v
	})
}

func BenchmarkStringID(b *testing.B) {
	sizes(b, func(b *testing.B, name []byte) {
		x := string(name)
		y := x // same string object
		var v bool
		for i := 0; i < b.N; i++ {
			v = x == y
		}
		sink = v
	})
}

func BenchmarkStringNE(b *testing.B) {
	sizes(b, func(b *testing.B, name []byte) {
		x := string(name)
		name[0] = ^name[0]
		y := string(name)
		var v bool
		for i := 0; i < b.N; i++ {
			v = x == y
		}
		sink = v
	})
}

func BenchmarkStringEQ(b *testing.B) {
	sizes(b, func(b *testing.B, name []byte) {
		x := string(name)
		y := string(name)
		var v bool
		for i := 0; i < b.N; i++ {
			v = x == y
		}
		sink = v
	})
}

func sizes(b *testing.B, f func(b *testing.B, name []byte)) {
	b.Run("10", func(b *testing.B) { f(b, make([]byte, 10)) })
	b.Run("100", func(b *testing.B) { f(b, make([]byte, 100)) })
	b.Run("1000", func(b *testing.B) { f(b, make([]byte, 1000)) })
}
