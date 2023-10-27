package channel_singleton

type singelton struct{}

var instance singelton

func GetInstance() *singelton {
	return &instance
}

func (s *singelton) AddOne() {
	addCh <- true
}

func (s *singelton) GetCount() int {

	resCh := make(chan int)
	defer close(resCh)

	getCountCh <- resCh
	return <-resCh
}

func (s *singelton) Stop() {

	quitCh <- true
	close(addCh)
	close(getCountCh)
	close(quitCh)
}

// Bidirectional Channel - used as reciever from
var addCh chan bool = make(chan bool)

// // Bidirectional Channel - used as reciever from // gives a channel in return
var getCountCh chan chan int = make(chan chan int)

// // Bidirectional Channel - used as reciever from
var quitCh chan bool = make(chan bool)

func init() {

	var count int

	go func(addCh <-chan bool, getCountCh <-chan chan int, quitCh <-chan bool) {

		for {
			select {
			case <-addCh:
				count++
			case ch := <-getCountCh:
				ch <- count //sending count value to channel

			case <-quitCh:
				break
			}
		}

	}(addCh, getCountCh, quitCh)
}
