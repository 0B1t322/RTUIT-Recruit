package hash

type HashUint map[uint]struct{}

func (h *HashUint) Add(u uint) {
	func(m map[uint]struct{}) {
		if _, find := m[u]; !find {
			m[u] = struct{}{}
		}
	}(*h)
}

func (h *HashUint) Delete(u uint) {
	func(m map[uint]struct{}) {
		delete(m, u)
	}(*h)
}

func (h *HashUint) Contains(u uint) bool {
	return func(m map[uint]struct{}) bool {
		_, find := m[u]
		return find
	}(*h)
}

func (h *HashUint) GetKeys() []uint {
	return func(m map[uint]struct{}) []uint {
		var keys []uint

		for k, _ := range m {
			keys = append(keys, k)
		}

		return keys
	}(*h)
}

func New() HashUint {
	return make(map[uint]struct{})
}