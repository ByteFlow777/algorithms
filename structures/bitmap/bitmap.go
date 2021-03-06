package bitmap

const (
	numLen = 8
)

type BitMap struct {
	n     uint64
	store []uint8
}

func NewBitmap(n uint64) *BitMap {
	return &BitMap{
		n:     n,
		store: make([]uint8, n/numLen+1),
	}
}

// GetRealIdx return the real index of the store slice
func (m *BitMap) GetRealIdx(num uint64) (arrIdx, bitIdx uint64) {
	arrIdx = num / numLen
	bitIdx = num - arrIdx*numLen
	return
}

// Add a num to the bitmap
func (m *BitMap) Add(num uint64) bool {
	if num > m.n {
		return false
	}

	arrIdx, bitIdx := m.GetRealIdx(num)
	m.store[arrIdx] |= 1 << bitIdx
	return true
}

// Contain check given num is existed
func (m *BitMap) Contain(num uint64) bool {
	if num > m.n {
		return false
	}

	arrIdx, bitIdx := m.GetRealIdx(num)
	return (m.store[arrIdx] & (1 << bitIdx)) != 0
}

// Del delete the given num
func (m *BitMap) Del(num uint64) {
	if num > m.n {
		return
	}

	arrIdx, bitIdx := m.GetRealIdx(num)
	m.store[arrIdx] &= ^(1 << bitIdx)
}

// Usage return the memory usage
func (m *BitMap) Usage() uint64 {
	bit := uint64(len(m.store)) * numLen
	return bit
}
