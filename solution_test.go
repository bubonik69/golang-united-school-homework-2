package square
import (
"testing"
)

const (
	SidesTriangle =3
	SidesSquare =4
	SidesCircle=0
)




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



