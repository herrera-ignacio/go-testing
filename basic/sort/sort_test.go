package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSortOrderDESC(t *testing.T) {
	// Init
	elements := []int{9,7,5,3,1,2,4,6,8,0}
	fmt.Println("Before", elements)

	// Execution
	BubbleSort(elements)

	// Validation
	fmt.Println("After", elements)

	if elements[0] != 9 {
		t.Error("First element should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("Last element should be 0")
	}
}

func TestBubbleSort(t *testing.T) {
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
				elements: []int{9,7,5,3,1,2,4,6,8,0},
			},
			want: []int{9,8,7,6,5,4,3,2,1,0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("Before", tt.args.elements)

			// Execution
			BubbleSort(tt.args.elements)

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
