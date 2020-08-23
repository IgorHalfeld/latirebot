package handlers

import (
	"testing"

	"github.com/igorhalfeld/latirebot/structs"
)

func TestSort(t *testing.T) {
	p := []structs.Product{
		structs.Product{
			NormalPrice:   10.0,
			DiscountPrice: 9,
		},
		structs.Product{
			NormalPrice:   20.0,
			DiscountPrice: 9,
		},
		structs.Product{
			NormalPrice:   10.0,
			DiscountPrice: 5,
		},
		structs.Product{
			NormalPrice:   10.0,
			DiscountPrice: 7,
		},
		structs.Product{
			NormalPrice:   10.0,
			DiscountPrice: 8,
		},
	}

	pSorted := []structs.Product{
		structs.Product{
			NormalPrice:   20.0,
			DiscountPrice: 9,
		},
		structs.Product{
			NormalPrice:   10.0,
			DiscountPrice: 5,
		},
		structs.Product{
			NormalPrice:   10.0,
			DiscountPrice: 7,
		},
		structs.Product{
			NormalPrice:   10.0,
			DiscountPrice: 8,
		},
		structs.Product{
			NormalPrice:   10.0,
			DiscountPrice: 9,
		},
	}
	p = sortDiscount(p)
	for k := range pSorted {
		if p[k].DiscountPrice != pSorted[k].DiscountPrice {
			t.Log("It was:", p[k])
			t.Log("Should Be: ", pSorted[k])
			t.Fail()
		}
	}
}
