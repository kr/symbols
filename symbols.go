// Package symbols provides Lisp-style primitive symbol objects.
package symbols

import "sync"

// A Symbol is an object uniquely identified by its name,
// which can be any string.
// If two symbols have equal names,
// the symbols themselves are also equal.
//
// A Symbol takes longer to create than the equivalent string.
// But once they're created, comparing two Symbols for equality
// is faster than comparing strings.
//
// A Symbol uses memory proportional to the size of its name.
// Note: once allocated, a Symbol is never freed,
// even if it becomes unreachable.
// Avoid allocating many large symbols.
//
// Programs using symbols should typically
// store and pass them as values, not pointers.
// That is, symbol variables and struct fields
// should be of type symbols.Symbol, not *symbols.Symbol.
//
// The zero value of Symbol is the symbol named "nil".
type Symbol struct{ name *string }

// We have to pick something.
// Another reasonable choice would be the empty string.
// But I think "nil" is more helpful.
const zeroName = "nil"

var (
	allSyms = map[string]Symbol{}
	mu      sync.Mutex
)

// Make returns the symbol named name.
func Make(name string) Symbol {
	if name == zeroName {
		// We must not create an entry for zeroName in the table,
		// because then there would be two unequal Symbol objects
		// (the zero value and the table entry) for the same name.
		return Symbol{}
	}
	mu.Lock()
	defer mu.Unlock()
	s, ok := allSyms[name]
	if !ok {
		s = Symbol{&name}
		allSyms[name] = s
	}
	return s
}

// String returns the name of s.
func (s Symbol) String() string {
	if s.name == nil {
		return zeroName
	}
	return *s.name
}
