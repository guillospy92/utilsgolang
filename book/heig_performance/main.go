package main

import (
	"fmt"
	"sync"
	"time"
)

type order struct {
	ID       int
	Total    float64
	Subtotal float64
	Delivery float64
	Discount float64
	StoreID  int
}

func main() {
	t := time.Now()
	chOrder := make(chan order)

	var wg sync.WaitGroup
	wg.Add(2)
	go getOrderInformation(chOrder, &wg)
	go discountOrder(chOrder, &wg)

	wg.Wait()
	fmt.Println(time.Since(t).Seconds())
}

func getOrderInformation(o chan order, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	order := order{ID: 1, Total: 0}
	order.ID = 1
	order.Total = 10_000
	o <- order
	close(o)
}

func discountOrder(o chan order, wg *sync.WaitGroup) order {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	orderChan := <-o
	orderChan.Total = orderChan.Total / 2

	fmt.Println(orderChan, "Hello")
	return orderChan
}

func deliveryAmount(order *order) {
	time.Sleep(1 * time.Second)
	order.Total += 3_000
}
