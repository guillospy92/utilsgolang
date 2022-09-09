package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const FileName = "students.csv"

var count int32

type user struct {
	Id, Name, LastName, Email, Phone string
	FriendIds                        []string
}

func main() {
	start := time.Now()
	users, err := processCSV()

	if err != nil {
		fmt.Println("error processing CSV", err)
		return
	}

	// sequentialProcessingSequential(users)

	chUser := make(chan *user)

	go sequentialProcessingConcurrency(users, chUser)

	sendNotificationConcurrency(chUser)

	fmt.Printf("Total Time: %v", time.Since(start).Seconds())

	fmt.Println("total register", count)

}

func sequentialProcessingSequential(users []*user) {
	visited := make(map[string]bool)

	for _, user := range users {
		if visited[user.Id] {
			continue
		}

		visited[user.Id] = true

		sendNotificationSequential(user)

		for _, friend := range user.FriendIds {
			user, err := findUserById(friend, users)

			if err != nil {
				fmt.Println("user no exists", user)
				continue
			}
			sendNotificationSequential(user)
		}
	}
}

func sendNotificationSequential(user *user) {
	time.Sleep(12 * time.Millisecond)
	fmt.Println("send notification \n", user.Phone)
	count++
}

func sequentialProcessingConcurrency(users []*user, chUser chan *user) {
	visited := make(map[string]bool)
	defer close(chUser)

	for _, user := range users {
		if visited[user.Id] {
			continue
		}

		visited[user.Id] = true

		chUser <- user

		for _, friend := range user.FriendIds {
			user, err := findUserById(friend, users)

			if err != nil {
				fmt.Println("user no exists", user)
				continue
			}

			chUser <- user
		}
	}
}

func sendNotificationConcurrency(userChan <-chan *user) {

	var wg = sync.WaitGroup{}

	sendNotification := func(worker int) {
		defer wg.Done()
		for user := range userChan {
			time.Sleep(12 * time.Millisecond)
			fmt.Println("send notification \n", user.Phone)
			fmt.Println("worker", worker)
			atomic.AddInt32(&count, 1)
		}
	}

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go sendNotification(i)
	}

	wg.Wait()
}

func findUserById(userId string, users []*user) (*user, error) {
	for _, user := range users {
		if user.Id == userId {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user not found with id %v", userId)
}

func processCSV() ([]*user, error) {
	f, err := os.Open(FileName)
	var users []*user

	if err != nil {
		log.Fatal("error open file", err)
		return users, err
	}

	s := bufio.NewScanner(f)

	for s.Scan() {
		// delete spacing
		line := strings.Trim(s.Text(), " ")

		// convert text in array
		lineArray := strings.Split(line, ",")

		// convert text in array colum 5
		ids := strings.Split(lineArray[5], " ")

		// delete double [[]] by []
		ids = ids[1 : len(ids)-1]

		user := &user{
			Id:        lineArray[0],
			Name:      lineArray[1],
			LastName:  lineArray[2],
			Email:     lineArray[3],
			Phone:     lineArray[4],
			FriendIds: ids,
		}

		users = append(users, user)
	}

	return users, nil
}
