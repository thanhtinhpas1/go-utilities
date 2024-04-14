package main

func zigzagEncode(val int32) int32 {
	return int32(uint32(val>>31)) ^ (val << 1)
}

func zigzadDecode(val int32) int32 {
	return (val >> 1) ^ -(val & 1)
}
