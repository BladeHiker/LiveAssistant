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
		DanMu:        make(chan []byte, 15),
		Gift:         make(chan []byte, 15),
		WelCome:      make(chan []byte, 15),
		WelComeGuard: make(chan []byte, 15),
		GreatSailing: make(chan []byte, 15),
		Online:       make(chan int, 5),
		Fans:         make(chan []byte, 5),
	}
}

func init() {
	P = NewPool()
}