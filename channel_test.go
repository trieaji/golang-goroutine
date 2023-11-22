package golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string) //membuat channel dengan tipe datanya adalah string
	defer close(channel)

	// channel <- "Laksa" //cara mengirim data ke channel

	// data := <- channel //cara menerima datanya (cara 1)

	// fmt.Println(<- channel)//cara jika ingin dikirim lngsung ke parameter

	//mengirim data goroutine menggunakan channel
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Ryokugyu"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	//menerima data goroutine yang menggunakan channel
	data := <-channel
	fmt.Println(data)

}

// channel sebagai parameter
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Aokiji"
}

func TestCreateChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// channel in (untuk parameter yg hanya berguna untuk mengirim data)
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Fujitora"
}

// channel out (untuk parameter yg hanya berguna untuk menerima data)
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// secara default channel hanya bisa menampung 1 data tapi kalau ingin menampung lebih dari 1 data, menggunakan buffered
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "kise"
	channel <- "akashi"
	channel <- "aomine"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println(cap(channel))
	fmt.Println(len(channel))

	// go func() {
	// 	channel <- "kise"
	// 	channel <- "akashi"
	// 	channel <- "aomine"
	// }()

	// go func() {
	// 	fmt.Println(<- channel)
	// 	fmt.Println(<- channel)
	// 	fmt.Println(<- channel)
	// }()

	// time.Sleep(2 * time.Second)

	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("selesai")
}

// select channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

// default select
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		default:
			fmt.Println("Menungggu data")
		}

		if counter == 2 {
			break
		}
	}
}
