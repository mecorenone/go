// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// A Type represents a type of Go.
// All types implement the Type interface.
type Type interface {
	// Underlying returns the underlying type of a type
	// w/o following forwarding chains. Only used by
	// client packages.
	Underlying() Type

	// String returns a string representation of a type.
	String() string
}

// under returns the true expanded underlying type.
// If it doesn't exist, the result is Typ[Invalid].
// under must only be called when a type is known
// to be fully set up.
func under(t Type) Type {
	if n := asNamed(t); n != nil {
		return n.under()
	}
	return t
}

// If the argument to asNamed, or asTypeParam is of the respective type
// (possibly after resolving a *Named type), these methods return that type.
// Otherwise the result is nil.

func asNamed(t Type) *Named {
	e, _ := t.(*Named)
	if e != nil {
		e.resolve(nil)
	}
	return e
}

func asTypeParam(t Type) *TypeParam {
	u, _ := under(t).(*TypeParam)
	return u
}
