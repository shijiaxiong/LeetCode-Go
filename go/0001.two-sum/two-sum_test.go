package problem0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	p para
	a ans
}

type para struct {
	one []int
	two int
}

type ans struct {
	one []int
}

func Test_Problem001(t *testing.T) {
	ast := assert.New(t)

	question := []question{
		{
			p: para{
				one: []int{3, 2, 4},
				two: 6,
			},
			a: ans{
				one: []int{1, 2},
			},
		},
		question{
			p: para{
				one: []int{3, 2, 4},
				two: 8,
			},
			a: ans{
				one: nil,
			},
		},
	}

	for _, ques := range question {
		p,a := ques.p, ques.a

		ast.Equal(a.one, twoSum(p.one, p.two), "输入:%v", p)
	}
}