package bitmap

import (
	"testing"
)

func TestNewBitmap(t *testing.T) {
	bitmap := NewBitmap(0)
	if len(bitmap.store) != 1 {
		t.Errorf("len of inner store, need: %d, but: %d",
			1, len(bitmap.store))
	}

	bitmap = NewBitmap(15)
	if len(bitmap.store) != 2 {
		t.Errorf("len of inner store, need: %d, but: %d",
			2, len(bitmap.store))
	}
}

func TestBitMap_Add(t *testing.T) {
	bitmap := NewBitmap(16)
	if f := bitmap.Add(0); !f {
		t.Errorf("add num error")
	}

	if f := bitmap.Add(256); f {
		t.Errorf("test for overflow num errors")
	}
}

func TestBitMap_Contain(t *testing.T) {
	bitmap := NewBitmap(16)

	_ = bitmap.Add(0)
	if !bitmap.Contain(0) {
		t.Errorf("the bitmap dont't contains 0")
	}

	if bitmap.Contain(256) {
		t.Errorf("the bitmap should't contains 256")
	}
}

func TestBitMap_Del(t *testing.T) {
	bitmap := NewBitmap(16)

	// test for overflow num
	bitmap.Del(256)

	_ = bitmap.Add(16)
	bitmap.Del(16)
	if bitmap.Contain(16) {
		t.Errorf("the bitmap should't contains deleted num 16")
	}
}

func TestBitMap_GetRealIdx(t *testing.T) {
	bitmap := NewBitmap(16)
	arrIdx, bitIdx := bitmap.GetRealIdx(0)
	if arrIdx != 0 && bitIdx != 0 {
		t.Errorf("want: %d %d, but: %d %d", 0, 0, arrIdx, bitIdx)
	}
}

func TestBitMap_Usage(t *testing.T) {
	bitmap := NewBitmap(16)
	if bitmap.Usage() != 8*3 {
		print(bitmap.Usage())
		t.Errorf("result of usage errors")
	}
}
