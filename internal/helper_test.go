package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_onlyContain(t *testing.T) {
	as := assert.New(t)

	as.True(isOnlyContain("....", "."))
}

func Test_inText(t *testing.T) {
	type Case struct {
		Key  string
		Text string
		In   bool
	}

	// a in `aaa
	// abc in `aaa`bbb`ccc
	// 你好世界 in `你`好啊 ，哈哈，这个`世`界
	cases := []Case{
		{
			Key:  "a",
			Text: "aa",
			In:   true,
		},
		{
			Key:  "abc",
			Text: "aabbcc",
			In:   true,
		},
		{
			Key:  "ab",
			Text: "bbaa",
			In:   false,
		},
		{
			Key:  "你好",
			Text: "我叫你，请问好不好？",
			In:   true,
		},
	}

	for _, v := range cases {
		assert.Equal(t, v.In, inText(v.Key, v.Text), fmt.Sprintf("%#v", v))
	}
}

func Test_splitSpace(t *testing.T) {
	as := assert.New(t)

	as.Equal([]string{"a b", "c"}, splitSpace(`a\ b c`))
	as.Equal([]string{"a b", "c"}, splitSpace(`a\ b c  `))
}
