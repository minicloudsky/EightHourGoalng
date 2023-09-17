package method

type User struct {
	Id   uint64
	Name string
	Age  int8
	Sex  string
}

func (u User) GetAge() int8 {
	return u.Age
}

func (u *User) SetAge(age int8) {
	u.Age = age
}

// Student 继承
type Student struct {
	User  User
	Score int8
	Class string
	Grade string
}

func (s *Student) SetScore(score int8) {
	s.Score = score
}

func (s *Student) GetScore() int8 {
	return s.Score
}
