package model

type Direction uint8

const (
	Up Direction = iota
	Down
	Left
	Right
)

type BodyPart struct {
	X int
	Y int
}

type Snake struct {
	Dir  Direction
	Body *LinkedList[BodyPart]
}

func NewSnake(startX int, startY int) *Snake {
	body := NewLinkedList[BodyPart](30)
	node := NewNode[BodyPart](BodyPart{startX, startY})
	node2 := NewNode[BodyPart](BodyPart{startX - 1, startY})
	body.Add(*node)
	body.Add(*node2)
	return &Snake{Right, body}
}

func (s *Snake) Move() {
	head := s.Body.Head()
	tail := s.Body.Tail()
	switch s.Dir {
	case Up:
		tail.Value.Y = head.Value.Y - 1
		tail.Value.X = head.Value.X
	case Down:
		tail.Value.Y = head.Value.Y + 1
		tail.Value.X = head.Value.X
	case Left:
		tail.Value.X = head.Value.X - 1
		tail.Value.Y = head.Value.Y
	case Right:
		tail.Value.X = head.Value.X + 1
		tail.Value.Y = head.Value.Y
	}

	head.Prev = s.Body.TailIndex
	tail.Next = s.Body.HeadIndex
	s.Body.HeadIndex = s.Body.TailIndex
	s.Body.TailIndex = tail.Prev

	tail.Prev = -1
	head.Next = -1
}

func (s *Snake) Grow() {
	tail := s.Body.Tail()
	tailNode := NewNode[BodyPart](BodyPart{tail.Value.X, tail.Value.Y})
	s.Body.Add(*tailNode)
}

func (s *Snake) CollidesWithFood(f Point) bool {
	head := s.Body.Head().Value
	return head.X == f.X && head.Y == f.Y
}

func (s *Snake) CollidesWithSelf() bool {
	head := s.Body.Head().Value
	for i, bp := range s.Body.Nodes {
		if s.Body.HeadIndex == i {
			continue
		}
		if bp.Value.X == head.X && bp.Value.Y == head.Y {
			return true
		}
	}
	return false
}

func (s *Snake) CollidesWithWall(maxX int, maxY int) bool {
	head := s.Body.Head().Value
	return head.X < 0 || head.X >= maxX || head.Y < 0 || head.Y >= maxY
}
