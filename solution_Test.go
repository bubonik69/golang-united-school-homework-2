package square
//package main
import (
"testing"
)

func TestOk (in *testing.T){
	testCases :=[]struct{
		sideLen float64
		sidesNum sideNumType
		square	float64
	}{
		{
			sideLen: 10,
			sidesNum: 0,
			square:  314.1592653589793,
		},
		{
			sideLen: 10,
			sidesNum: 2,
			square : 100,
		},
		{
			sideLen: 10,
			sidesNum: 3,
			square : 43.30127018922193,
		},
	}
	for _,tc:=range testCases{
		tc:=tc
		if tc.square!= CalcSquare(tc.sideLen, tc.sidesNum){
			in.Errorf("test finished with error")
			return
		}

	}
	return

}

