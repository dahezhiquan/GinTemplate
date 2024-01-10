package encrypts

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"sync"
	"testing"
)

func layBinaryId() uint64 {
	var randomInt uint64
	_ = binary.Read(rand.Reader, binary.LittleEndian, &randomInt)
	return randomInt
}

func TestLayBinaryId(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id := layBinaryId()
			log.Println(id)
		}()
	}

	wg.Wait()
}

func TestFunnyWorkerId(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id := FunnyWorkerId()
			log.Println(id)
		}()
	}

	wg.Wait()
}

func TestFNV1a(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			hash := GetRandHex()
			log.Println(hash)
		}()
	}

	wg.Wait()
}

// 测试生成雪花❄ID生成
func TestLayId(t *testing.T) {
	log.Println("======测试生成雪花❄ID生成======")
	var wg sync.WaitGroup

	var mimi []string

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			hash := LayIDHash()
			found := false
			for _, v := range mimi {
				if v == hash {
					found = true
					break
				}
			}
			if found {
				fmt.Printf("%s is in the slice =================================================================\n", hash)
			}
			mimi = append(mimi, hash)
			log.Println(hash)
		}()
	}
	wg.Wait()
}
