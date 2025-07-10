package arithmetic

type Arithmetic struct{}

func (a *Arithmetic) Add(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}

func (a *Arithmetic) Multiply(nums []int) int {
	s := 1
	for _, num := range nums {
		s *= num
	}
	return s
}
