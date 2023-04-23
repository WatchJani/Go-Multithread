package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func Put(c chan int) {
	ran := rand.Intn(5) + 10

	time.Sleep(time.Duration(ran) * 100 * time.Millisecond)
	log.Println("Random number:", ran)

	for index := 0; index < ran; index++ {
		c <- rand.Intn(100)
	}
}

//asinhrono citanje
//zasto ako stavimo ovu funkciju u main nece raditi
//jer se ovdje blokira citavu ovu nit sve dok ne dobije
// dati chanel element koji ce procitati

func Get(c chan int) {
	for {
		fmt.Println("Reader:", <-c)
	}

	fmt.Println("blokira, nikad se ovo nece izvrsiti!!!")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	c, counter := make(chan int), 0

	go Put(c)
	go Put(c)
	// go Get(c)

	//ako zelimo da ne zablokiramo glavnu nit, koristimo select
	//kako bi izasli iz beskonacne petlje koristimo counter, kad su u pitanju unbufered chanel

	for {
		select {
		case b := <-c:
			fmt.Println(b)
		}
		counter++
		fmt.Println("counter:", counter)
		if counter == 10 {
			break
		}
	}

	time.Sleep(2*time.Second + 100*time.Millisecond)
	fmt.Println("Sad radi")
}
