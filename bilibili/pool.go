package bilibili

type Pool struct {
	DanMu        chan []byte
	Gift         chan []byte
	WelCome      chan []byte
	WelComeGuard chan []byte
	Online       chan int
	Fans         chan int
}

func NewPool() *Pool {
	return &Pool{
		DanMu:        make(chan []byte, 10),
		Gift:         make(chan []byte, 10),
		WelCome:      make(chan []byte, 10),
		WelComeGuard: make(chan []byte, 10),
		Online:       make(chan int, 10),
		Fans:         make(chan int, 10),
	}
}
