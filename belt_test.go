package belt_test

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"testing"

	"github.com/Bowbaq/belt"
	"github.com/stretchr/testify/assert"
)

func TestCheckNil(t *testing.T) {
	if os.Getenv("TEST_CHECK_NIL") == "1" {
		belt.Check(nil)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestCheckNil")
	cmd.Env = append(os.Environ(), "TEST_CHECK_NIL=1")

	assert.NoError(t, cmd.Run())
}

func TestCheckErr(t *testing.T) {
	if os.Getenv("TEST_CHECK_ERR") == "1" {
		belt.Check(errors.New("This should crash"))
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestCheckErr")
	cmd.Env = append(os.Environ(), "TEST_CHECK_ERR=1")

	if err := cmd.Run(); assert.IsType(t, err, &exec.ExitError{}) {
		exit := err.(*exec.ExitError)
		assert.False(t, exit.Success(), "Process ran without error, expected exit code 1")
	}
}

func ExampleCheck() {
	_, err := http.Get("http://www.example.com/")
	belt.Check(err)
}

func TestContainsNilSlice(t *testing.T) {
	assert.False(t, belt.Contains(nil, "needle"))
}

func TestContainsEmptySlice(t *testing.T) {
	assert.False(t, belt.Contains([]string{}, "needle"))
}

func TestContainsPositiveFirst(t *testing.T) {
	assert.True(t, belt.Contains([]int{1, 2, 3, 4, 5}, 1))
}

func TestContainsPositiveLast(t *testing.T) {
	assert.True(t, belt.Contains([]int{1, 2, 3, 4, 5}, 5))
}

func TestContainsPositiveMiddle(t *testing.T) {
	assert.True(t, belt.Contains([]int{1, 2, 3, 4, 5}, 3))
}

func TestContainsNegative(t *testing.T) {
	assert.False(t, belt.Contains([]int{1, 2, 3, 4, 5}, 42))
}

func TestContainsNegativeTypeMismatch(t *testing.T) {
	assert.False(t, belt.Contains([]int{1, 2, 3, 4, 5}, "42"))
}

func benchmarkContains(l int, b *testing.B) {
	xs := make([]int, l)
	for i := 0; i < l; i++ {
		xs[i] = rand.Intn(1000)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		belt.Contains(xs, rand.Intn(1000))
	}
}

func BenchmarkContains10(b *testing.B)    { benchmarkContains(10, b) }
func BenchmarkContains100(b *testing.B)   { benchmarkContains(100, b) }
func BenchmarkContains1000(b *testing.B)  { benchmarkContains(1000, b) }
func BenchmarkContains10000(b *testing.B) { benchmarkContains(10000, b) }

func ExampleContains() {
	if belt.Contains(cakes, "Apple Pie") {
		fmt.Println("Apple Pie is a cake")
	} else {
		fmt.Println("Apple Pie is *not* a cake")
	}

	if belt.Contains(cakes, "Ice Cream") {
		fmt.Println("Ice Cream is a cake")
	} else {
		fmt.Println("Ice Cream is *not* a cake")
	}
	// Output:
	// Apple Pie is a cake
	// Ice Cream is *not* a cake
}

func TestHasPrefixNilSlice(t *testing.T) {
	assert.False(t, belt.HasPrefix("Apple Pie", nil))
}

func TestHasPrefixEmptySlice(t *testing.T) {
	assert.False(t, belt.HasPrefix("Apple Pie", []string{}))
}

func TestHasPrefixPositiveFirst(t *testing.T) {
	assert.True(t, belt.HasPrefix("Apple Pie", cakeFruits))
}

func TestHasPrefixPositiveMiddle(t *testing.T) {
	assert.True(t, belt.HasPrefix("Blueberry Pie", cakeFruits))
}

func TestHasPrefixPositiveLast(t *testing.T) {
	assert.True(t, belt.HasPrefix("Peach Cobbler", cakeFruits))
}

func TestHasPrefixPositiveSingle(t *testing.T) {
	assert.True(t, belt.HasPrefix("Peach Cobbler", []string{"Peach"}))
}

func TestHasPrefixNegative(t *testing.T) {
	assert.False(t, belt.HasPrefix("Ice Cream", cakeFruits))
}

func ExampleHasPrefix() {
	if belt.HasPrefix("Apple Pie", cakeFruits) {
		fmt.Println("Apples are a fruit")
	} else {
		fmt.Println("Apples are not fruit")
	}

	if belt.HasPrefix("Ice Cream", cakeFruits) {
		fmt.Println("Ice is a fruit")
	} else {
		fmt.Println("Ice is *not* a fruit")
	}
	// Output:
	// Apples are a fruit
	// Ice is *not* a fruit
}

func TestHasSuffixNilSlice(t *testing.T) {
	assert.False(t, belt.HasSuffix("Apple Pie", nil))
}

func TestHasSuffixEmptySlice(t *testing.T) {
	assert.False(t, belt.HasSuffix("Apple Pie", []string{}))
}

func TestHasSuffixPositiveFirst(t *testing.T) {
	assert.True(t, belt.HasSuffix("Apple Pie", cakeStyles))
}

func TestHasSuffixPositiveMiddle(t *testing.T) {
	assert.True(t, belt.HasSuffix("Apple Crumble", cakeStyles))
}

func TestHasSuffixPositiveLast(t *testing.T) {
	assert.True(t, belt.HasSuffix("Peach Cobbler", cakeStyles))
}

func TestHasSuffixPositiveSingle(t *testing.T) {
	assert.True(t, belt.HasSuffix("Peach Cobbler", []string{"Cobbler"}))
}

func TestHasSuffixNegative(t *testing.T) {
	assert.False(t, belt.HasSuffix("Ice Cream", cakeStyles))
}

func ExampleHasSuffix() {

	if belt.HasSuffix("Apple Pie", cakeStyles) {
		fmt.Println("Apple Pie is a cake")
	} else {
		fmt.Println("Apple Pie is *not* a cake")
	}

	if belt.HasSuffix("Ice Cream", cakeStyles) {
		fmt.Println("Ice Cream is a cake")
	} else {
		fmt.Println("Ice Cream is *not* a cake")
	}
	// Output:
	// Apple Pie is a cake
	// Ice Cream is *not* a cake
}

var (
	cakes      = []string{"Apple Pie", "Apple Crumble", "Peach Cobbler"}
	cakeStyles = []string{"Pie", "Crumble", "Cobbler"}
	cakeFruits = []string{"Apple", "Blueberry", "Peach"}
)
