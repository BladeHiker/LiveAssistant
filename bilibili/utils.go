package bilibili

import (
	"bytes"
	"compress/zlib"
	"encoding/hex"
	"fmt"
	"io"
	"math"
)

// 返回字节数组表示数的十进制形式
func ByteArrToDecimal(src []byte) (sum int) {
	if src == nil {
		return 0
	}
	b := []byte(hex.EncodeToString(src))
	l := len(b)
	for i := l - 1; i >= 0; i-- {
		base := int(math.Pow(16, float64(l-i-1)))
		var mul int
		if int(b[i]) >= 97 {
			mul = int(b[i]) - 87
		} else {
			mul = int(b[i]) - 48
		}

		sum += base * mul
	}
	return
}

// gzip 格式数据的解压缩
func ZlibInflate(compress []byte) ([]byte, error) {
	var out bytes.Buffer
	c := bytes.NewReader(compress)
	r, err := zlib.NewReader(c)
	if err != zlib.ErrChecksum && err != zlib.ErrDictionary && err != zlib.ErrHeader && r != nil {
		_, _ = io.Copy(&out, r)
		if err := r.Close(); err != nil {
			fmt.Println("r.close err:", err)
			return nil, err
		}
		return out.Bytes(), nil
	}
	return nil, err
}
