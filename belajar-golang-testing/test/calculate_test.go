package test

import (
	"errors"
	"go-testing/controller"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculateSum(t *testing.T) {
	result := controller.CalculateSum(100, 200, 300)
	if result != (100 + 200 + 300) {
		panic("Summation error!")
	}
}

func TestCalculateSumNegative(t *testing.T) {
	result := controller.CalculateSum(100, -200, -300)
	if result != (100 - 200 - 300) {
		panic("Negative summation error!")
	}
}

/*
	- Menjalankan unit test
		go test
	- Melihat detail function apa aja yg di test
		go test -v -> v itu verbose
	- Ingin memilih function yang di test
		go test -v -run=TestCalculateSumNegative
	- Run semua test di semua package
		go test ./...

		go test -v ./...
		go test -v ./... -run=TestCalculateSum	-> menjalankan TestCalculateSum dan TestCalculateSumNegative
		go test -v ./... -run=TestCalculateSumNegative
*/

func TestCalculateDivision(t *testing.T) {
	t.Run("ZeroDivisor", func(t *testing.T) {
		result, err := controller.CalculateDivision(9, 0)
		errTrue := errors.New("Division by zero!")

		require.Equal(t, errTrue, err, err.Error())
		assert.Equal(t, 0, result)
	})
	t.Run("Negative", func(t *testing.T) {
		result1, _ := controller.CalculateDivision(9, -9)
		require.Equal(t, -1, result1)

		result2, _ := controller.CalculateDivision(-9, 9)
		assert.Equal(t, -1, result2)
	})
}

/*
	go test -v ./... -run=TestCalculateDivision/ZeroDivisor
	go test -v ./... -run=/Negative
*/

func TestCalculateDivisionTable(t *testing.T) {
	cases := []struct {
		name            string
		requestDividend int
		requestDivisor  int
		expectedResult  int
		expectedError   error
	}{
		{
			name:            "ZeroDividend",
			requestDividend: 0,
			requestDivisor:  9,
			expectedResult:  0,
			expectedError:   nil,
		},
		{
			name:            "ZeroDivisor",
			requestDividend: 9,
			requestDivisor:  0,
			expectedResult:  0,
			expectedError:   errors.New("Division by zero!"),
		},
		{
			name:            "NegativeDivident",
			requestDividend: -9,
			requestDivisor:  3,
			expectedResult:  -3,
			expectedError:   nil,
		},
		{
			name:            "NegativeDivisor",
			requestDividend: 10,
			requestDivisor:  -2,
			expectedResult:  -5,
			expectedError:   nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, err := controller.CalculateDivision(c.requestDividend, c.requestDivisor)

			require.Equal(t, c.expectedError, err)
			assert.Equal(t, c.expectedResult, result)
		})
	}
}

func BenchmarkCalculateSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomizer := rand.New(rand.NewSource(10))
		randomLength := randomizer.Intn(10000)
		var randomSlice []int
		for i := 0; i < randomLength; i++ {
			randomSlice = append(randomSlice, i)
		}

		controller.CalculateSum(randomSlice...)
	}
}

/*
	Menjalankan benchmark dengan semua unit test
		go test -v -bench=. ./...

	Menjalankan benchmark tanpa unit test
		go test -v -run=TestX -bench=. ./...

	Menjalankan spesifik benchmark tanpa unit test
		go test -v -run=TestX -bench=BenchmarkCalculateSum ./...
*/

func BenchmarkCalculateSumSub(b *testing.B) {
	b.Run("ZeroDivisor", func(b *testing.B) {
		randomizer := rand.New(rand.NewSource(9))
		randomDividend := randomizer.Intn(100000)

		for i := 0; i < b.N; i++ {
			controller.CalculateDivision(randomDividend, 0)
		}
	})

	b.Run("ZeroDividend", func(b *testing.B) {
		randomizer := rand.New(rand.NewSource(8))
		randomDivisor := randomizer.Intn(100000)

		for i := 0; i < b.N; i++ {
			controller.CalculateDivision(0, randomDivisor)
		}
	})
}

/*
	- Menjalankan hanya benchmark
		go test -v ./... -run=TestX -bench=BenchmarkCalculateSumSub/ZeroDivisor
*/

func BenchmarkCalculateSumTable(b *testing.B) {
	cases := []struct {
		name string
		a    int
		b    int
	}{
		{
			name: "ZeroDivisor",
			a:    9,
			b:    0,
		},
		{
			name: "ZeroDividend",
			a:    0,
			b:    9,
		},
		{
			name: "AllPositives",
			a:    120,
			b:    3,
		},
		{
			name: "AllNegatives",
			a:    -45,
			b:    -15,
		},
	}
	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			for i := 1; i < b.N; i++ {
				controller.CalculateDivision(c.a, c.b)
			}
		})
	}
}
