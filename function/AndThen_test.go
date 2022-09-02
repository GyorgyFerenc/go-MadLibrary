package Function_test

import (
	"testing"

	f "github.com/GyorgyFerenc/go-MadLibrary/function"

	"github.com/stretchr/testify/assert"
)

func TestAndThen(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(f.AndThen(0, f.Adder(1), f.SubstracterRHS(1)), 0)
	assert.Equal(f.AndThen(68419, f.Adder(1000), f.Adder(1)), 69420)

	zero := f.AndThen(0, f.Adder(10), f.Adder(20), f.SubstracterRHS(10), f.SubstracterRHS(20))
	assert.Equal(zero, 0)

}

func TestAdder(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(f.Adder(1)(1), 2)
	assert.Equal(f.Adder(1)(2), 3)
	assert.Equal(f.Adder(1)(2), f.Adder(2)(1))

	add0 := f.Adder(0)

	for i := 0; i < 100; i++ {
		assert.Equal(add0(i), i)
	}

	assert.Equal(f.Adder[float32](0.0)(0.0), float32(0.0))

}

func TestMultiplier(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(f.Multiplier(1)(1), 1)
	assert.Equal(f.Multiplier(1)(2), 2)
	assert.Equal(f.Multiplier(1)(2), f.Multiplier(2)(1))

	mul1 := f.Multiplier(1)

	for i := 0; i < 100; i++ {
		assert.Equal(mul1(i), i)
	}

	assert.Equal(f.Multiplier[float32](0.0)(0.0), float32(0.0))

}

func TestSubstracter(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(f.SubstracterLHS(5)(2), 3)
	assert.NotEqual(f.SubstracterLHS(5)(2), f.SubstracterLHS(2)(5))

	assert.Equal(f.SubstracterLHS(5)(3), f.SubstracterRHS(3)(5))
	assert.Equal(f.SubstracterRHS(100)(200), 100)

	assert.Equal(f.SubstracterRHS(1)(0), 0-1)

}

func TestDivider(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(f.DividerLHS(5)(2), 2)
	assert.Equal(f.DividerLHS(4)(2), 2)

	assert.Equal(f.DividerLHS(5.0)(2.0), 2.5)
	assert.Equal(f.DividerLHS(5.0)(2.0), f.DividerRHS(2.0)(5.0))

}
