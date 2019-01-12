package compare

import "v2ray.com/v2ray-study/common/errors"

func BytesEqualWithDetail(a []byte, b []byte) error {
	if len(a) != len(b) {
		return errors.New("mismatch array length ", len(a), " vs ", len(b))
	}
	for idx, v := range a {
		if b[idx] != v {
			return errors.New("mismatch array value at index [", idx, "]: ", v, " vs ", b[idx])
		}
	}
	return nil
}

// Note: compare two byte slice
func BytesEqual(a []byte, b []byte) bool {
	return BytesEqualWithDetail(a, b) == nil
}

// Note: check if elements of the byte array is same as target
func BytesAll(arr []byte, value byte) bool {
	for _, v := range arr {
		if v != value {
			return false
		}
	}

	return true
}
