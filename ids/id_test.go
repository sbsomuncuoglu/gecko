// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package ids

import (
	"bytes"
	"testing"
)

func TestID(t *testing.T) {
	hash := [32]byte{24}
	id := NewID(hash)

	if key := id.Key(); !bytes.Equal(hash[:], key[:]) {
		t.Fatalf("ID.Key returned wrong bytes")
	}

	prefixed := id.Prefix(0)

	if key := id.Key(); !bytes.Equal(hash[:], key[:]) {
		t.Fatalf("ID.Prefix mutated the ID")
	}

	if nextPrefix := id.Prefix(0); !prefixed.Equals(nextPrefix) {
		t.Fatalf("ID.Prefix not consistant")
	}

	if b := id.Bytes(); !bytes.Equal(hash[:], b) {
		t.Fatalf("ID.Bytes returned wrong bytes")
	}

	if str := id.String(); str != "Ba3mm8Ra8JYYebeZ9p7zw1ayorDbeD1euwxhgzSLsncKqGoNt" {
		t.Fatalf("ID.String returned wrong string: %s", str)
	}
}

func TestIDBit(t *testing.T) {
	id0 := NewID([32]byte{1 << 0})
	id1 := NewID([32]byte{1 << 1})
	id2 := NewID([32]byte{1 << 2})
	id3 := NewID([32]byte{1 << 3})
	id4 := NewID([32]byte{1 << 4})
	id5 := NewID([32]byte{1 << 5})
	id6 := NewID([32]byte{1 << 6})
	id7 := NewID([32]byte{1 << 7})
	id8 := NewID([32]byte{0, 1 << 0})

	if id0.Bit(0) != 1 {
		t.Fatalf("Wrong bit")
	} else if id1.Bit(1) != 1 {
		t.Fatalf("Wrong bit")
	} else if id2.Bit(2) != 1 {
		t.Fatalf("Wrong bit")
	} else if id3.Bit(3) != 1 {
		t.Fatalf("Wrong bit")
	} else if id4.Bit(4) != 1 {
		t.Fatalf("Wrong bit")
	} else if id5.Bit(5) != 1 {
		t.Fatalf("Wrong bit")
	} else if id6.Bit(6) != 1 {
		t.Fatalf("Wrong bit")
	} else if id7.Bit(7) != 1 {
		t.Fatalf("Wrong bit")
	} else if id8.Bit(8) != 1 {
		t.Fatalf("Wrong bit")
	}
}

func TestFromString(t *testing.T) {
	key := [32]byte{'a', 'v', 'a', ' ', 'l', 'a', 'b', 's'}
	id := NewID(key)
	idStr := id.String()
	id2, err := FromString(idStr)
	if err != nil {
		t.Fatal(err)
	}
	if id.Key() != id2.Key() {
		t.Fatal("Expected FromString to be inverse of String but it wasn't")
	}
}
