package goroutine_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

//

func RunHelloWorld() {
	fmt.Println("Hello World!")
}

func TestCreateGoroutines(t *testing.T) {
	go RunHelloWorld()                   // dijalankan no (2). Tidak bisa menerima return value dari fungsi
	fmt.Println("UpsUpsUpsUpsUpsUpsUps") // dijalankan no (3) karena strinynya besar
	fmt.Println("Ups")                   // dijalankan no (1) karena stringnya kecil

	time.Sleep(1 * time.Second)
}

//

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Test" // Apabila channel tidak diisi maka program akan tetap menunggu
		fmt.Println("Selesai mengirimkan data ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(2 * time.Second)
}

//

func GenerateRandomNumber(channel chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Mulai menghasilkan random number")

	randomNumber := rand.Intn(10000)
	channel <- randomNumber

	fmt.Println("Selesai menghasilkan random number")
}

func TestGenerateRandomNumber(t *testing.T) {
	channel := make(chan int)
	defer close(channel)

	go GenerateRandomNumber(channel)

	fmt.Println("Memproses,")

	randomNumberResult := <-channel
	fmt.Println(randomNumberResult)

	fmt.Println("Program selesai")
	time.Sleep(2 * time.Second)
}

//

func InChannel(channel chan<- string) {
	// time.Sleep(2 * time.Second)
	channel <- "Test..."
	fmt.Println("Selesai memberikan data ke channel!")
}

func OutChannel(channel <-chan string) {
	data := <-channel
	fmt.Println("Data: ", data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go InChannel(channel)
	go OutChannel(channel)

	time.Sleep(3 * time.Second)
}

//

func TestBufferedChannel(t *testing.T) {
	/*
		- channel := make(chan string)
			maka akan blocking -> program berjalan terus sampai ada yang mengambil
		- channel := make(chan string, 1)
			tidak akan blocking karena masuk ke buffer
	*/
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "A"
		channel <- "B"
		channel <- "C"
	}()
	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

//

func TestRangeChannel(t *testing.T) {
	capacity := 10
	channel := make(chan string, capacity)

	go func() {
		for i := 0; i < capacity; i++ {
			channel <- "Perulangan ke-" + strconv.Itoa(i)
		}
		defer close(channel)
	}()
	go func() { // kalau for di luar goroutine dan channel tidak di close maka program berjalan terus
		for data := range channel {
			fmt.Println(data)
		}
	}()

	time.Sleep(2 * time.Second)
}

//

/*
Untuk mengambil data pada channel dengan benar
*/
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan int)
	channel2 := make(chan int)
	defer close(channel1)
	defer close(channel2)

	go GenerateRandomNumber(channel1)
	go GenerateRandomNumber(channel2)

	go func() {
		counter := 0
		for {
			if counter >= 2 {
				break
			}
			select {
			case data := <-channel1:
				fmt.Println("Data: ", data)
				counter++
			case data := <-channel2:
				fmt.Println("Data: ", data)
				counter++
			default:
				fmt.Println("Menunggu data....") // dijalankan apabila channel belum berisi data
			}
		}
	}()

	time.Sleep(5 * time.Second)
}

//

func TestRaceCondition(t *testing.T) {
	x := 0
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 0; j <= 100; j++ {
				x += 1
			}
		}()
	}

	time.Sleep(4 * time.Second)
	fmt.Println(x)
}

//

func TestMutexRaceCondition(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 10; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(4 * time.Second)
	fmt.Println(x)
}

//

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance uint64
}

func (ba *BankAccount) AddBalance(n uint64) {
	ba.RWMutex.Lock()
	ba.Balance += n
	ba.RWMutex.Unlock()
}

func (ba *BankAccount) ReadBalance() uint64 {
	var balance uint64
	ba.RWMutex.RLock()
	balance = ba.Balance
	ba.RWMutex.RUnlock()

	return balance
}

func TestReadWriteMutex(t *testing.T) {
	bankAccount := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			bankAccount.AddBalance(uint64(i))
			fmt.Println(bankAccount.ReadBalance())
		}()
	}

	time.Sleep(4 * time.Second)
	fmt.Println("Final balance", bankAccount.Balance)
}

//

type UserBalance struct {
	sync.Mutex
	Nama    string
	Balance uint64
}

func (ub *UserBalance) Lock() {
	ub.Mutex.Lock()
}

func (ub *UserBalance) Unlock() {
	ub.Mutex.Unlock()
}

func (ub *UserBalance) Change(amount uint64) {
	ub.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount uint64, i int) {
	fmt.Println("Transfer ke-" + strconv.Itoa(i))

	user1.Lock()
	fmt.Println("Lock user1 ", user1.Nama)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2 ", user2.Nama)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	fmt.Println("Unlock user 1", user1.Nama)
	user1.Unlock()

	fmt.Println("Unlock user 2", user2.Nama)
	user2.Unlock()
}

func TestTransfer(t *testing.T) {
	user1 := UserBalance{
		Nama:    "A",
		Balance: 1000,
	}
	user2 := UserBalance{
		Nama:    "B",
		Balance: 1000,
	}

	// deadlock: saling menunggu
	go Transfer(&user1, &user2, 100, 1) // hasil: user1=900 user2=1100
	go Transfer(&user2, &user1, 300, 2) // hasil: user1=1200 user2=800
	// hasil deadlock: user1=900 user2=700

	time.Sleep(4 * time.Second)

	fmt.Println("User1 ", user1.Nama, " balance: ", user1.Balance)
	fmt.Println("User2 ", user2.Nama, " balance: ", user2.Balance)
}

//

/*
RunAsynchronus
  - Membuat goroutine yang melakukan print Hello
*/
func RunAsynchronus(group *sync.WaitGroup, i int) {
	defer group.Done()

	// group.Add(1)

	fmt.Printf("Hello %d \n", i)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		group.Add(1)
		go RunAsynchronus(group, i)
	}

	group.Wait()
	fmt.Println("Complete")
}

//

var onceCounter = 0

func OnlyOnce() {
	fmt.Println("OnlyOnce called!")
	onceCounter += 1
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		fmt.Println("Begin waitgroup...")
		group.Add(1)
		go func() {
			once.Do(OnlyOnce) // hanya dijalankan sekali walaupun masuk ke 100 goroutine
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("onceConter = ", onceCounter)
}

//

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return "New"
		},
	}
	group := sync.WaitGroup{}

	pool.Put("A")
	pool.Put("B")
	pool.Put("C")

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {

			data := pool.Get()
			fmt.Println("Pool data =. ", data)

			time.Sleep(1 * time.Second)
			pool.Put(data)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("TestPool finished!")
}

//

func AddToSyncMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	data.Store(value, value*10)
}

func TestSyncMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		group.Add(1)
		go AddToSyncMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, ": ", value)
		return true
	})
}

//

func WaitCondition(cond *sync.Cond, value int, group *sync.WaitGroup) {
	defer group.Done()

	cond.L.Lock()
	cond.Wait() // setelah saya locking saya boleh lanjut tidak ke proses selanjutnya?
	fmt.Println("Done ", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	var locker = sync.Mutex{}
	var cond = sync.NewCond(&locker)
	var group = sync.WaitGroup{}
	var numGo = 10

	for i := 0; i < numGo; i++ {
		group.Add(1)
		go WaitCondition(cond, i, &group)
	}

	go func() {
		for i := 0; i < numGo; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast() // mensignal semua goroutine
	// }()

	group.Wait()
}

//

func TestAtomic(t *testing.T) {
	var counter uint64 = 0
	var waitgroup = sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		waitgroup.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				atomic.AddUint64(&counter, 1) // counter += 10
				fmt.Printf("Counter %d \n", counter)
			}
			waitgroup.Done()
		}()
	}

	waitgroup.Wait()
	fmt.Println(counter)
}

//

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	channel := time.After(1 * time.Second)

	tick := <-channel // butuh channelnya saja
	fmt.Println(tick)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}

	group.Add(1)
	time.AfterFunc(5*time.Second, func() { // nunggu lima detik dulu
		fmt.Println("Time inside func() ", time.Now()) // baru ini dieksekusi
		group.Done()
	})

	fmt.Println("Time outside func() ", time.Now())
	group.Wait()
}

//

func StopTicker(ticker *time.Ticker) {
	fmt.Println("Stopping the ticker...")
	ticker.Stop()
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second) // ticker itu interval
	maxIter := 5

	for time := range ticker.C {
		if maxIter == 0 {
			StopTicker(ticker)
		} else {
			fmt.Println(time)
			maxIter--
		}
	}
}

func TesTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}

func TestGoMaxProcs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println(totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println(totalThread)

	// total goroutine yang berjalan saat ini
	totalGoroutine := runtime.NumGoroutine() // default 2: 1 untuk menjalankan fungsinya, 1 untuk garbage collection
	fmt.Println(totalGoroutine)
}
