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
	c := make(chan int)

	go Put(c)
	go Put(c)
	go Get(c)

	//ako bi to uradili u main, funkcija bi cekala da joj se posalje neki drugi
	//chanel element, i nikad se ne bi zavrsila, dok kada to uradimo sa go rutin
	//to je druga nit, ona moze ostati blokirana do kraja programa, cekace sve
	//dok joj neko ne posalje neku novu vrijednost

	time.Sleep(2*time.Second + 100*time.Millisecond)
	fmt.Println("Sad radi")
}
