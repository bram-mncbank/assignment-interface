package main

import (
	"fmt"
	"interface/libs"
	"sync"
)

func main() {
	var db []*libs.User

	svc1 := libs.MyUsers(db)

	people := []string{"bram", "adi", "budi", "sanjaya", "lisa", "rudi", "akri"}
	var wg sync.WaitGroup

	wg.Add(len(people))

	for _, v := range people {
		go func(n string) {
			r := svc1.Register(&libs.User{Nama: n})
			fmt.Printf("Inject N: %v\n", r)
			wg.Done()
		}(v)

	}
	wg.Wait()
	fmt.Println("-----------------------------------")

	listPeople := svc1.GetUser()

	wg.Add(len(listPeople))

	for _, v := range listPeople {
		go func(p *libs.User) {
			fmt.Println(p.Nama)
			wg.Done()
		}(v)
	}
	wg.Wait()

}
