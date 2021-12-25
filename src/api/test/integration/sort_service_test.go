package integration

import (
	"fmt"
	"github.com/herrera-ignacio/go-testing/src/api/services"
	"sort"
	"testing"
)

func TestConstants(t *testing.T) {
	// We can only test public constants from outside the package
	if services.PublicConst != "public" {
		t.Error("privateConst should be 'private'")
	}
}

func TestSort(t *testing.T) {
	// Init
	type args struct {
		elements []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testing desc",
			args: args{
				elements: []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0},
			},
			want: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("Before", tt.args.elements)

			// Execution
			services.Sort(tt.args.elements)

			fmt.Println("After", tt.args.elements)
			fmt.Println(tt.want)

			// Validation
			if len(tt.args.elements) != len(tt.want) {
				t.Error("Length should match")
			}

			for i, val := range tt.args.elements {
				if val != tt.want[i] {
					t.Errorf("Should match %d and %d", val, tt.want[i])
				}
			}
		})
	}
}

func BenchmarkTestSort(b *testing.B) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	for i := 0; i < b.N; i++ {
		services.Sort(elements)
	}
}

func BenchmarkTestSTLSort(b *testing.B) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
}
