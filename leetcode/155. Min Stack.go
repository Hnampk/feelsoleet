package leetcode

type MinStack struct {
	data []int
	min  []int
}

func Constructor() MinStack {
	return MinStack{
		data: []int{},
		min:  []int{},
	}
}

func (this *MinStack) Push(val int) {
	min := val
	if len(this.data) > 0 {
		if val > this.GetMin() {
			min = this.GetMin()
		}
	}

	this.data = append(this.data, val)
	this.min = append(this.min, min)
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
