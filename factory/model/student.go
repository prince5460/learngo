package model

type student struct {
	Name  string
	Age   int
	score float64
}

//工厂模式
func NewStudent(name string, age int, score float64) *student {
	return &student{
		Name:  name,
		Age:   age,
		score: score,
	}
}

func (s *student) GetScore() float64 {
	return s.score
}
