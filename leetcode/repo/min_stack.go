package repo

type MinStack struct {
	array []int
	stack []int
}

func (this *MinStack) Push(val int) {
	this.array = append(this.array, val)
	if len(this.stack) == 0 || this.stack[len(this.stack)-1] >= val {
		this.stack = append(this.stack, val)
	}
}

func (this *MinStack) Pop() {
	val := this.array[len(this.array)-1]
	this.array = this.array[:len(this.array)-1]
	if val == this.stack[len(this.stack)-1] {
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.array[len(this.array)-1]
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
