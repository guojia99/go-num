package _55

// Constructor 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
// 实现 MinStack 类:
//
// MinStack() 初始化堆栈对象。
// void push(int val) 将元素val推入堆栈。
// void pop() 删除堆栈顶部的元素。
// int top() 获取堆栈顶部的元素。
// int getMin() 获取堆栈中的最小元素。
//
// 示例 1:
//
// 输入：
// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
//
// 输出：
// [null,null,null,null,-3,null,0,-2]
//
// 解释：
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin();   --> 返回 -3.
// minStack.pop();
// minStack.top();      --> 返回 0.
// minStack.getMin();   --> 返回 -2.
//
// 提示：
//
// -231 <= val <= 231 - 1
// pop、top 和 getMin 操作总是在 非空栈 上调用
// push, pop, top, and getMin最多被调用 3 * 104 次
func Constructor() MinStack {
	return MinStack{
		val: make([][2]int, 0),
	}
}

type MinStack struct {
	val [][2]int // [i][0] 该元素最大值 [i][1] 该元素推入时栈的最小值
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (c *MinStack) Push(val int) {
	if len(c.val) == 0 {
		c.val = append(c.val, [2]int{val, val})
		return
	}
	c.val = append(c.val, [2]int{val, min(val, c.val[len(c.val)-1][1])})
}

func (c *MinStack) Pop() {
	c.val = c.val[:len(c.val)-1]
}

func (c *MinStack) Top() int {
	return c.val[len(c.val)-1][0]
}

func (c *MinStack) GetMin() int {
	return c.val[len(c.val)-1][1]
}
