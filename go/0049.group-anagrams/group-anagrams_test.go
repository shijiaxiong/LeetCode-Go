package problem0049

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	one []string
}

type ans struct {
	one [][]string
}

func Test_Problem0049(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			ans{[][]string{
				[]string{"bat"},
				[]string{"nat", "tan"},
				[]string{"ate", "eat", "tea"},
			}},
		},
	}

	for _,q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		res:= groupAnagrams(p.one)
		for _,v  := range res {
			// 排序是为了和构造的ans保持一致
			sort.Strings(v)
			ast.Equal(a.one[len(v)-1], v,"输出：%v", p)
		}

	}
}
