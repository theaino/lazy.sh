package cache

import "crypto/md5"

func Sum(data []byte) []byte {
	sum := md5.Sum(data)
	return sum[:]
}
