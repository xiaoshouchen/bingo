package helper

func BytesToInt(bytes []byte) uint64 {
	var num uint64 = 0
	for i := 0; i < len(bytes); i++ {
		num |= uint64(bytes[i]) << (i * 8)
	}
	return num
}
