package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func () interface{}  {
			return "not nil"
		},// untuk membuat data pool otomatis
	}

	pool.Put("Laksa") // memasukkan data
	pool.Put("Bayu") // memasukkan data
	pool.Put("Trie") // memasukkan data

	for i := 0; i < 10; i++ {
		go func ()  {
			data := pool.Get() // mengambil data
			fmt.Println(data)
			time.Sleep(1 * time.Second) // supaya pengambilan data nya sesuai
			pool.Put(data) // memasukkan data nya kembali
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("selesai")
}