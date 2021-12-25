package services

import "github.com/herrera-ignacio/go-testing/src/api/utils/sort"

const (
	privateConst = "private"
	PublicConst  = "public"
)

func Sort(elements []int) {
	sort.BubbleSort(elements)
}
