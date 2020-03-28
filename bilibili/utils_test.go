package bilibili

import "testing"

func TestByteArrToDecimal(t *testing.T) {
	data := "000000140010000100000003000000010001c133"
	sum := ByteArrToDecimal([]byte(data))
	if sum < 0 {
		t.Error("ByteArrToDecimal([]byte(data)) < 0 err")
	}
}

func TestZlibInflate(t *testing.T) {
	// 未压缩数据
	noDeflated := "000000140010000100000003000000010001c133"
	// 被压缩的数据
	deflated := "0000012200100002000000050000000078da548fbb4a03611085d7c2de6738f50466feeb663b41b48a2062b52e126384408c859b2a044c69e3052c0221082a8882454893264f13b2316f217f101667e0c00c87f3cd44d1d651b41385da0ed243e3ea0209f6760f6b2767b5e303105a9dcb6b2469ca24a42c89f3de2bb12436b646546cb43331559caa1aa7ad54890975ef7d3d8ecf411c3a23acee46a0545ba559694d283e87cbefc79fc1f36af452dcbf2e9f867f5e12660e2c20a3d41096ef1fc578b29ecd41580c1e8af1a4984e8bdbaff56cbe18bc8144395f154fceb053da101080a90871b93bed32eba60dd12116792b6f372be24d4540ffa66c7344a7db6e530ff90d92f24f42234702e3c55abd2fe897d68d70d6ff0d0000ffffa2dd6291"

	n, err := ZlibInflate([]byte(noDeflated))
	if n != nil && err == nil {
		t.Error("ZlibInflate([]byte(noDeflated)) err")
	}

	b, err := ZlibInflate([]byte(deflated))
	if b == nil || err != nil {
		t.Error("ZlibInflate([]byte(deflated)) err")
	}
}
