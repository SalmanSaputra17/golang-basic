package math

import "testing"

type testpair struct {
	values []float64
	result float64
}

var avgSlice = []testpair{
  	{ []float64{1,2}, 1.5 },
  	{ []float64{1,1,1,1,1,1}, 1 },
  	{ []float64{-1,1}, 0 },
}

var maxSlice = []testpair{
  	{ []float64{1,2}, 2 },
  	{ []float64{1,5,7,8,4,1}, 8 },
  	{ []float64{-1,10}, 10 },
}
var minSlice = []testpair{
  	{ []float64{1,2}, 1 },
  	{ []float64{1,1,1,1,1,1}, 1 },
  	{ []float64{-1,1}, -1 },
}

func TestAverage(t *testing.T) {
  	for _, pair := range avgSlice {
  		if pair.values != nil && int(len(pair.values)) != 0 {
  			v := Average(pair.values)
	    	if v != pair.result {
		      	t.Error(
			        "For", pair.values,
			        "expected", pair.result,
			        "got", v,
			    )
			}
  		} else {
  			t.Error("Expected not empty slice, got", pair.result)
  		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range maxSlice {
  		if pair.values != nil && int(len(pair.values)) != 0 {
  			v := Max(pair.values)
	    	if v != pair.result {
		      	t.Error(
			        "For", pair.values,
			        "expected", pair.result,
			        "got", v,
			    )
			}
  		} else {
  			t.Error("Expected not empty slice, got", pair.result)
  		}
	}
}

func TestMin(t *testing.T) {
	for _, pair := range minSlice {
  		if pair.values != nil && int(len(pair.values)) != 0 {
  			v := Min(pair.values)
	    	if v != pair.result {
		      	t.Error(
			        "For", pair.values,
			        "expected", pair.result,
			        "got", v,
			    )
			}
  		} else {
  			t.Error("Expected not empty slice, got", pair.result)
  		}
	}	
}