// Value and Pointer semantics

// When it come to built-in type (numeric, string, bool) we should always be using value
// semantics. Instead of creating garbage by allocating them on the heap we should keep
// them on the stack. There are exceptions but until we know it is okay to take the
// exception, we should follow the guideline.

// The reference type (slice, map, channel, interface) also focuses on using value semantics.
// The only time we want to take the address of a slice is when we are sharing it down the
// call stack to Unmarshal function since it always requires the address of a value.

// When declaring a type, we must ask ourselves immediately:
// - Does this type require value semantic or pointer semantic?
// - If I need to modify this value, should we create a new value or should we modify
// the value itself so everyone can see it?
// It needs to be consistent. It is okay to guess it wrong the first time and refactor later.

package main

import (
	"sync/atomic"
	"syscall"
)

// --------------
// Value semantic
// --------------

// Since we use value semantics for reference types,
// the implementations is using value semantics for both.
type IP []byte
type IPMask []byte

//
