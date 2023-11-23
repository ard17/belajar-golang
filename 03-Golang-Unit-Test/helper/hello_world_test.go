package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldFail(t *testing.T) {
	result := HelloWorld("Name")
	if result != "Hello Name" {
		t.Fail()
	}

	fmt.Println("TestHelloWorldFail done")
}

func TestHelloWorldFailNow(t *testing.T) {
	result := HelloWorld("Name")
	if result != "Hello Name" {
		t.FailNow()
	}
	fmt.Println("TestHelloWorldFailNow done")
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("Name")
	if result != "Hello Name" {
		t.Error("Result must be 'Hello Name'")
	}

	fmt.Println("TestHelloWorldError done")
}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("Name")
	if result != "Hello Name" {
		t.Fatal("Result must be 'Hello Name'")
	}
	fmt.Println("TestHelloWorldFatal done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Name")
	assert.Equal(t, "Hello Name", result, "Result must be 'Hello Name'")
	fmt.Println("TestHelloWorldAssert done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Name")
	require.Equal(t, "Hello Name", result, "Result must be 'Hello Name'")
	fmt.Println("TestHelloWorldRequire done")
}

func TestHelloWorldSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can't run on linux")
	}

	result := HelloWorld("Name")
	require.Equal(t, "Hello Name", result, "Result must be 'Hello Name'")
	fmt.Println("TestHelloWorldRequire done")
}

func TestMain(m *testing.M) {
	// before unit test run
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after unit test run
	fmt.Println("AFTER UNIT TEST")
}

func TestSubtest(t *testing.T) {
	t.Run("Cloud", func(t *testing.T) {
		result := HelloWorld("Cloud")
		require.Equal(t, "Hello Cloud", result, "Result must be 'Hello Cloud'")
	})
	t.Run("Strife", func(t *testing.T) {
		result := HelloWorld("Strife")
		require.Equal(t, "Hello Strife", result, "Result must be 'Hello Strife'")
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Cloud",
			request:  "Cloud",
			expected: "Hello Cloud",
		},
		{
			name:     "Strife",
			request:  "Strife",
			expected: "Hello Strife",
		},
		{
			name:     "Tifa",
			request:  "Tifa",
			expected: "Hello Tifa",
		},
		{
			name:     "Lockhart",
			request:  "Lockhart",
			expected: "Hello Lockhart",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}

}

func BenchmarkHelloWorldBelajar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Belajar")

	}
}
func BenchmarkHelloWorldGolang(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Golang")

	}
}
func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Belajar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Belajar")

		}
	})
	b.Run("Golang", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Golang")

		}
	})
}
func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Belajar",
			request: "Belajar",
		},
		{
			name:    "Golang",
			request: "Golang",
		},
		{
			name:    "Benchmark",
			request: "Benchmark",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
