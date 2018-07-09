package opts

import (
	"reflect"
	"testing"
)

func TestValidateCPUPeriod(t *testing.T) {
	cases := []struct {
          in, want int64
      }{
          {0, 0},
	  {500,0},
          {1000,0}, 
      }
      for _, c := range cases {
          got := ValidateCPUPeriod(c.in)
          if got != nil {
              t.Errorf("CPU cfs period  %d cannot be less than 1ms (i.e. 1000) or larger than 1s (i.e. 1000000)", period)
          }
}
}

func TestValidateCPUQuota(t *testing.T) {
	// TODO
}
