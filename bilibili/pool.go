package bilibili

type Pool struct {
	DanMu        chan []byte
	Gift         chan []byte
	WelCome      chan []byte
	WelComeGuard chan []byte
	GreatSailing chan []byte
	Online       chan int
	Fans         chan []byte
}

func NewPool() *Pool {
	return &Pool{
		DanMu:        make(chan []byte, 10),
		Gift:         make(chan []byte, 10),
		WelCome:      make(chan []byte, 10),
		WelComeGuard: make(chan []byte, 10),
		GreatSailing: make(chan []byte, 10),
		Online:       make(chan int, 10),
		Fans:         make(chan []byte, 10),
	}
}