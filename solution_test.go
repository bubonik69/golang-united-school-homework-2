package square
import (
"testing"
)

const (
	SidesTriangle =3
	SidesSquare =4
	SidesCircle=0
)



func TestCalc (t *testing.T){
	testCases :=[]struct{
		sideLen float64
		sidesNum sideNumType
		square	float64
	}{
		{
			sideLen: 10,
			sidesNum: SidesCircle,
			square:  314.1592653589793,
		},
		{
			sideLen: 10,
			sidesNum: SidesSquare,
			square : 100,
		},
		{
			sideLen: 10,
			sidesNum: SidesTriangle,
			square : 43.30127018922193,
		},
	}
	for _,tc:=range testCases{
		//tc:=tc
		actual:=CalcSquare(tc.sideLen, tc.sidesNum)
		if tc.square-actual>0.01{
			t.Errorf("test finished with error, result of test: ")
		}
	}
	t.Log("OK!")
}

func TestCalcSquare(t *testing.T) {
	var expects float64 = 100
	actual := CalcSquare(100, SidesSquare)

	diff(t, expects, actual)
}

func diff(t *testing.T, expects, actual float64) {
	if expects-actual > 0.01 {
		t.Errorf("actual %f expect %f", actual, expects)
	}

	t.Log("OK!")
}



