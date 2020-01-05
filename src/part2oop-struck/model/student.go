package model

//演示go的工厂模式

// Student 结构体
type Student struct {
	Name  string
	Score int
}

type student struct {
	name  string
	score int
}

// NewStudent 手动创建一个学生对象
func NewStudent(name string, score int) *student {
	return &student{
		name:  name,
		score: score,
	}
}

//GetStudent 手动获取一个学生对象
func (s *student) GetName() string {
	return s.name
}

//// 通过goland创建getter setter
//func (s *student) Score() int {
//	return s.score
//}
//
//func (s *student) SetScore(score int) {
//	s.score = score
//}
//
//func (s *student) Name() string {
//	return s.name
//}
//
//func (s *student) SetName(name string) {
//	s.name = name
//}
