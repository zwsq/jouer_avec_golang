package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

// User represents a user with an id and first name.
type User struct {
	ID    int
	First string
}

type MockDatastore struct {
	Users map[int]User
}

func (md MockDatastore) GetUser(id int) (User, error) {
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf("User %v not found", id)
	}
	return user, nil
}

func (md MockDatastore) SaveUser(u User) error {
	_, ok := md.Users[u.ID]
	if ok {
		return fmt.Errorf("User %v already exists", u.ID)
	}
	md.Users[u.ID] = u
	return nil
}

type Datastore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

type Service struct {
	ds Datastore
}

func (s Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}

func (s Service) SaveUser(u User) error {
	return s.ds.SaveUser(u)
}

func main() {
	db := MockDatastore{
		Users: make(map[int]User),
	}

	srvc := Service{
		ds: db,
	}

	u1 := User{
		ID:    1,
		First: "James",
	}

	err := srvc.SaveUser(u1)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	u1Returned, err := srvc.GetUser(u1.ID)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	fmt.Println(u1)
	fmt.Println(u1Returned)
	now := time.Now()
	callGoPower(2)
	elapsed := time.Since(now)
	print(elapsed)

}
func goPower(n float64, p float64) float64 {
	return math.Pow(n, p)
}
func callGoPower(n float64) {
	var i float64 = 0
	for i <= 1000 {
		i++
		fmt.Println(goPower(n, i))
	}
}
