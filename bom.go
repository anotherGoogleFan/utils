package utils

import "bytes"

func AddBom(s []byte) []byte {
	return append([]byte{0xef, 0xbb, 0xbf}, s...)
}

func RemoveBom(s []byte) []byte {
	return bytes.Trim(s, "\xef\xbb\xbf")
}
