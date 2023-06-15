package test_stu

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestHelloTom(t *testing.T) {
	output := HelloTom()
	expectOutput := "Jerry!"
	assert.Equal(t, expectOutput, output)
}

func TestJudgePassLineTrue(t *testing.T) {
	isPass := JudgePassLine(70)
	assert.Equal(t, true, isPass)
}

func TestJudgePassLineFlase(t *testing.T) {
	isPass := JudgePassLine(50)
	assert.Equal(t, true, isPass)
}