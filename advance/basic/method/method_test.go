package method

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	user := User{
		Id:   10,
		Name: "tony",
		Age:  100,
		Sex:  "male",
	}
	fmt.Println(user.GetAge())
	user.SetAge(20)
	fmt.Println(user.GetAge())
}

func TestStudent(t *testing.T) {
	s := Student{
		User: User{
			Id:   100,
			Name: "susan",
			Age:  23,
			Sex:  "female",
		},
		Score: 30,
		Class: "five",
		Grade: "nine",
	}
	fmt.Println(s)
	s.Score = 10
	fmt.Println(s.Score)
	s.SetScore(90)
	fmt.Println(s.Score)
}
