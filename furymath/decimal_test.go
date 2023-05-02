package furymath_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"

	"github.com/merlins-labs/merlins/furymath"
)

type decimalTestSuite struct {
	suite.Suite
}

var (
	zeroAdditiveErrTolerance = furymath.ErrTolerance{
		AdditiveTolerance: sdk.ZeroDec(),
	}
)

func TestDecimalTestSuite(t *testing.T) {
	suite.Run(t, new(decimalTestSuite))
}

// assertMutResult given expected value after applying a math operation, a start value,
// mutative and non mutative results with start values, asserts that mutation are only applied
// to the mutative versions. Also, asserts that both results match the expected value.
func (s *decimalTestSuite) assertMutResult(expectedResult, startValue, mutativeResult, nonMutativeResult, mutativeStartValue, nonMutativeStartValue furymath.BigDec) {
	// assert both results are as expected.
	s.Require().Equal(expectedResult, mutativeResult)
	s.Require().Equal(expectedResult, nonMutativeResult)

	// assert that mutative method mutated the receiver
	s.Require().Equal(mutativeStartValue, expectedResult)
	// assert that non-mutative method did not mutate the receiver
	s.Require().Equal(nonMutativeStartValue, startValue)
}

func (s *decimalTestSuite) TestAddMut() {
	toAdd := furymath.MustNewDecFromStr("10")
	tests := map[string]struct {
		startValue        furymath.BigDec
		expectedMutResult furymath.BigDec
	}{
		"0":  {furymath.NewBigDec(0), furymath.NewBigDec(10)},
		"1":  {furymath.NewBigDec(1), furymath.NewBigDec(11)},
		"10": {furymath.NewBigDec(10), furymath.NewBigDec(20)},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			startMut := tc.startValue.Clone()
			startNonMut := tc.startValue.Clone()

			resultMut := startMut.AddMut(toAdd)
			resultNonMut := startNonMut.Add(toAdd)

			s.assertMutResult(tc.expectedMutResult, tc.startValue, resultMut, resultNonMut, startMut, startNonMut)
		})
	}
}

func (s *decimalTestSuite) TestQuoMut() {
	quoBy := furymath.MustNewDecFromStr("2")
	tests := map[string]struct {
		startValue        furymath.BigDec
		expectedMutResult furymath.BigDec
	}{
		"0":  {furymath.NewBigDec(0), furymath.NewBigDec(0)},
		"1":  {furymath.NewBigDec(1), furymath.MustNewDecFromStr("0.5")},
		"10": {furymath.NewBigDec(10), furymath.NewBigDec(5)},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			startMut := tc.startValue.Clone()
			startNonMut := tc.startValue.Clone()

			resultMut := startMut.QuoMut(quoBy)
			resultNonMut := startNonMut.Quo(quoBy)

			s.assertMutResult(tc.expectedMutResult, tc.startValue, resultMut, resultNonMut, startMut, startNonMut)
		})
	}
}
func TestDecApproxEq(t *testing.T) {
	// d1 = 0.55, d2 = 0.6, tol = 0.1
	d1 := furymath.NewDecWithPrec(55, 2)
	d2 := furymath.NewDecWithPrec(6, 1)
	tol := furymath.NewDecWithPrec(1, 1)

	require.True(furymath.DecApproxEq(t, d1, d2, tol))

	// d1 = 0.55, d2 = 0.6, tol = 1E-5
	d1 = furymath.NewDecWithPrec(55, 2)
	d2 = furymath.NewDecWithPrec(6, 1)
	tol = furymath.NewDecWithPrec(1, 5)

	require.False(furymath.DecApproxEq(t, d1, d2, tol))

	// d1 = 0.6, d2 = 0.61, tol = 0.01
	d1 = furymath.NewDecWithPrec(6, 1)
	d2 = furymath.NewDecWithPrec(61, 2)
	tol = furymath.NewDecWithPrec(1, 2)

	require.True(furymath.DecApproxEq(t, d1, d2, tol))
}

// create a decimal from a decimal string (ex. "1234.5678")
func (s *decimalTestSuite) MustNewDecFromStr(str string) (d furymath.BigDec) {
	d, err := furymath.NewDecFromStr(str)
	s.Require().NoError(err)

	return d
}

func (s *decimalTestSuite) TestNewDecFromStr() {
	largeBigInt, success := new(big.Int).SetString("3144605511029693144278234343371835", 10)
	s.Require().True(success)

	tests := []struct {
		decimalStr string
		expErr     bool
		exp        furymath.BigDec
	}{
		{"", true, furymath.BigDec{}},
		{"0.-75", true, furymath.BigDec{}},
		{"0", false, furymath.NewBigDec(0)},
		{"1", false, furymath.NewBigDec(1)},
		{"1.1", false, furymath.NewDecWithPrec(11, 1)},
		{"0.75", false, furymath.NewDecWithPrec(75, 2)},
		{"0.8", false, furymath.NewDecWithPrec(8, 1)},
		{"0.11111", false, furymath.NewDecWithPrec(11111, 5)},
		{"314460551102969.31442782343433718353144278234343371835", true, furymath.NewBigDec(3141203149163817869)},
		{
			"314460551102969314427823434337.18357180924882313501835718092488231350",
			true, furymath.NewDecFromBigIntWithPrec(largeBigInt, 4),
		},
		{
			"314460551102969314427823434337.1835",
			false, furymath.NewDecFromBigIntWithPrec(largeBigInt, 4),
		},
		{".", true, furymath.BigDec{}},
		{".0", true, furymath.NewBigDec(0)},
		{"1.", true, furymath.NewBigDec(1)},
		{"foobar", true, furymath.BigDec{}},
		{"0.foobar", true, furymath.BigDec{}},
		{"0.foobar.", true, furymath.BigDec{}},
		{"179769313486231590772930519078902473361797697894230657273430081157732675805500963132708477322407536021120113879871393357658789768814416622492847430639474124377767893424865485276302219601246094119453082952085005768838150682342462881473913110540827237163350510684586298239947245938479716304835356329624224137216", true, furymath.BigDec{}},
	}

	for tcIndex, tc := range tests {
		res, err := furymath.NewDecFromStr(tc.decimalStr)
		if tc.expErr {
			s.Require().NotNil(err, "error expected, decimalStr %v, tc %v", tc.decimalStr, tcIndex)
		} else {
			s.Require().Nil(err, "unexpected error, decimalStr %v, tc %v", tc.decimalStr, tcIndex)
			s.Require().True(res.Equal(tc.exp), "equality was incorrect, res %v, exp %v, tc %v", res, tc.exp, tcIndex)
		}

		// negative tc
		res, err = furymath.NewDecFromStr("-" + tc.decimalStr)
		if tc.expErr {
			s.Require().NotNil(err, "error expected, decimalStr %v, tc %v", tc.decimalStr, tcIndex)
		} else {
			s.Require().Nil(err, "unexpected error, decimalStr %v, tc %v", tc.decimalStr, tcIndex)
			exp := tc.exp.Mul(furymath.NewBigDec(-1))
			s.Require().True(res.Equal(exp), "equality was incorrect, res %v, exp %v, tc %v", res, exp, tcIndex)
		}
	}
}

func (s *decimalTestSuite) TestDecString() {
	tests := []struct {
		d    furymath.BigDec
		want string
	}{
		{furymath.NewBigDec(0), "0.000000000000000000000000000000000000"},
		{furymath.NewBigDec(1), "1.000000000000000000000000000000000000"},
		{furymath.NewBigDec(10), "10.000000000000000000000000000000000000"},
		{furymath.NewBigDec(12340), "12340.000000000000000000000000000000000000"},
		{furymath.NewDecWithPrec(12340, 4), "1.234000000000000000000000000000000000"},
		{furymath.NewDecWithPrec(12340, 5), "0.123400000000000000000000000000000000"},
		{furymath.NewDecWithPrec(12340, 8), "0.000123400000000000000000000000000000"},
		{furymath.NewDecWithPrec(1009009009009009009, 17), "10.090090090090090090000000000000000000"},
		{furymath.MustNewDecFromStr("10.090090090090090090090090090090090090"), "10.090090090090090090090090090090090090"},
	}
	for tcIndex, tc := range tests {
		s.Require().Equal(tc.want, tc.d.String(), "bad String(), index: %v", tcIndex)
	}
}

func (s *decimalTestSuite) TestDecFloat64() {
	tests := []struct {
		d    furymath.BigDec
		want float64
	}{
		{furymath.NewBigDec(0), 0.000000000000000000},
		{furymath.NewBigDec(1), 1.000000000000000000},
		{furymath.NewBigDec(10), 10.000000000000000000},
		{furymath.NewBigDec(12340), 12340.000000000000000000},
		{furymath.NewDecWithPrec(12340, 4), 1.234000000000000000},
		{furymath.NewDecWithPrec(12340, 5), 0.123400000000000000},
		{furymath.NewDecWithPrec(12340, 8), 0.000123400000000000},
		{furymath.NewDecWithPrec(1009009009009009009, 17), 10.090090090090090090},
	}
	for tcIndex, tc := range tests {
		value, err := tc.d.Float64()
		s.Require().Nil(err, "error getting Float64(), index: %v", tcIndex)
		s.Require().Equal(tc.want, value, "bad Float64(), index: %v", tcIndex)
		s.Require().Equal(tc.want, tc.d.MustFloat64(), "bad MustFloat64(), index: %v", tcIndex)
	}
}

func (s *decimalTestSuite) TestSdkDec() {
	tests := []struct {
		d        furymath.BigDec
		want     sdk.Dec
		expPanic bool
	}{
		{furymath.NewBigDec(0), sdk.MustNewDecFromStr("0.000000000000000000"), false},
		{furymath.NewBigDec(1), sdk.MustNewDecFromStr("1.000000000000000000"), false},
		{furymath.NewBigDec(10), sdk.MustNewDecFromStr("10.000000000000000000"), false},
		{furymath.NewBigDec(12340), sdk.MustNewDecFromStr("12340.000000000000000000"), false},
		{furymath.NewDecWithPrec(12340, 4), sdk.MustNewDecFromStr("1.234000000000000000"), false},
		{furymath.NewDecWithPrec(12340, 5), sdk.MustNewDecFromStr("0.123400000000000000"), false},
		{furymath.NewDecWithPrec(12340, 8), sdk.MustNewDecFromStr("0.000123400000000000"), false},
		{furymath.NewDecWithPrec(1009009009009009009, 17), sdk.MustNewDecFromStr("10.090090090090090090"), false},
	}
	for tcIndex, tc := range tests {
		if tc.expPanic {
			s.Require().Panics(func() { tc.d.SDKDec() })
		} else {
			value := tc.d.SDKDec()
			s.Require().Equal(tc.want, value, "bad SdkDec(), index: %v", tcIndex)
		}
	}
}

func (s *decimalTestSuite) TestBigDecFromSdkDec() {
	tests := []struct {
		d        sdk.Dec
		want     furymath.BigDec
		expPanic bool
	}{
		{sdk.MustNewDecFromStr("0.000000000000000000"), furymath.NewBigDec(0), false},
		{sdk.MustNewDecFromStr("1.000000000000000000"), furymath.NewBigDec(1), false},
		{sdk.MustNewDecFromStr("10.000000000000000000"), furymath.NewBigDec(10), false},
		{sdk.MustNewDecFromStr("12340.000000000000000000"), furymath.NewBigDec(12340), false},
		{sdk.MustNewDecFromStr("1.234000000000000000"), furymath.NewDecWithPrec(12340, 4), false},
		{sdk.MustNewDecFromStr("0.123400000000000000"), furymath.NewDecWithPrec(12340, 5), false},
		{sdk.MustNewDecFromStr("0.000123400000000000"), furymath.NewDecWithPrec(12340, 8), false},
		{sdk.MustNewDecFromStr("10.090090090090090090"), furymath.NewDecWithPrec(1009009009009009009, 17), false},
	}
	for tcIndex, tc := range tests {
		if tc.expPanic {
			s.Require().Panics(func() { furymath.BigDecFromSDKDec(tc.d) })
		} else {
			value := furymath.BigDecFromSDKDec(tc.d)
			s.Require().Equal(tc.want, value, "bad furymath.BigDecFromSdkDec(), index: %v", tcIndex)
		}
	}
}

func (s *decimalTestSuite) TestBigDecFromSdkDecSlice() {
	tests := []struct {
		d        []sdk.Dec
		want     []furymath.BigDec
		expPanic bool
	}{
		{[]sdk.Dec{sdk.MustNewDecFromStr("0.000000000000000000")}, []furymath.BigDec{furymath.NewBigDec(0)}, false},
		{[]sdk.Dec{sdk.MustNewDecFromStr("0.000000000000000000"), sdk.MustNewDecFromStr("1.000000000000000000")}, []furymath.BigDec{furymath.NewBigDec(0), furymath.NewBigDec(1)}, false},
		{[]sdk.Dec{sdk.MustNewDecFromStr("1.000000000000000000"), sdk.MustNewDecFromStr("0.000000000000000000"), sdk.MustNewDecFromStr("0.000123400000000000")}, []furymath.BigDec{furymath.NewBigDec(1), furymath.NewBigDec(0), furymath.NewDecWithPrec(12340, 8)}, false},
		{[]sdk.Dec{sdk.MustNewDecFromStr("10.000000000000000000")}, []furymath.BigDec{furymath.NewBigDec(10)}, false},
		{[]sdk.Dec{sdk.MustNewDecFromStr("12340.000000000000000000")}, []furymath.BigDec{furymath.NewBigDec(12340)}, false},
		{[]sdk.Dec{sdk.MustNewDecFromStr("1.234000000000000000"), sdk.MustNewDecFromStr("12340.000000000000000000")}, []furymath.BigDec{furymath.NewDecWithPrec(12340, 4), furymath.NewBigDec(12340)}, false},
		{[]sdk.Dec{sdk.MustNewDecFromStr("0.123400000000000000"), sdk.MustNewDecFromStr("12340.000000000000000000")}, []furymath.BigDec{furymath.NewDecWithPrec(12340, 5), furymath.NewBigDec(12340)}, false},
		{[]sdk.Dec{sdk.MustNewDecFromStr("0.000123400000000000"), sdk.MustNewDecFromStr("10.090090090090090090")}, []furymath.BigDec{furymath.NewDecWithPrec(12340, 8), furymath.NewDecWithPrec(1009009009009009009, 17)}, false},
		{[]sdk.Dec{sdk.MustNewDecFromStr("10.090090090090090090"), sdk.MustNewDecFromStr("10.090090090090090090")}, []furymath.BigDec{furymath.NewDecWithPrec(1009009009009009009, 17), furymath.NewDecWithPrec(1009009009009009009, 17)}, false},
	}
	for tcIndex, tc := range tests {
		if tc.expPanic {
			s.Require().Panics(func() { furymath.BigDecFromSDKDecSlice(tc.d) })
		} else {
			value := furymath.BigDecFromSDKDecSlice(tc.d)
			s.Require().Equal(tc.want, value, "bad furymath.BigDecFromSdkDec(), index: %v", tcIndex)
		}
	}
}

func (s *decimalTestSuite) TestEqualities() {
	tests := []struct {
		d1, d2     furymath.BigDec
		gt, lt, eq bool
	}{
		{furymath.NewBigDec(0), furymath.NewBigDec(0), false, false, true},
		{furymath.NewDecWithPrec(0, 2), furymath.NewDecWithPrec(0, 4), false, false, true},
		{furymath.NewDecWithPrec(100, 0), furymath.NewDecWithPrec(100, 0), false, false, true},
		{furymath.NewDecWithPrec(-100, 0), furymath.NewDecWithPrec(-100, 0), false, false, true},
		{furymath.NewDecWithPrec(-1, 1), furymath.NewDecWithPrec(-1, 1), false, false, true},
		{furymath.NewDecWithPrec(3333, 3), furymath.NewDecWithPrec(3333, 3), false, false, true},

		{furymath.NewDecWithPrec(0, 0), furymath.NewDecWithPrec(3333, 3), false, true, false},
		{furymath.NewDecWithPrec(0, 0), furymath.NewDecWithPrec(100, 0), false, true, false},
		{furymath.NewDecWithPrec(-1, 0), furymath.NewDecWithPrec(3333, 3), false, true, false},
		{furymath.NewDecWithPrec(-1, 0), furymath.NewDecWithPrec(100, 0), false, true, false},
		{furymath.NewDecWithPrec(1111, 3), furymath.NewDecWithPrec(100, 0), false, true, false},
		{furymath.NewDecWithPrec(1111, 3), furymath.NewDecWithPrec(3333, 3), false, true, false},
		{furymath.NewDecWithPrec(-3333, 3), furymath.NewDecWithPrec(-1111, 3), false, true, false},

		{furymath.NewDecWithPrec(3333, 3), furymath.NewDecWithPrec(0, 0), true, false, false},
		{furymath.NewDecWithPrec(100, 0), furymath.NewDecWithPrec(0, 0), true, false, false},
		{furymath.NewDecWithPrec(3333, 3), furymath.NewDecWithPrec(-1, 0), true, false, false},
		{furymath.NewDecWithPrec(100, 0), furymath.NewDecWithPrec(-1, 0), true, false, false},
		{furymath.NewDecWithPrec(100, 0), furymath.NewDecWithPrec(1111, 3), true, false, false},
		{furymath.NewDecWithPrec(3333, 3), furymath.NewDecWithPrec(1111, 3), true, false, false},
		{furymath.NewDecWithPrec(-1111, 3), furymath.NewDecWithPrec(-3333, 3), true, false, false},
	}

	for tcIndex, tc := range tests {
		s.Require().Equal(tc.gt, tc.d1.GT(tc.d2), "GT result is incorrect, tc %d", tcIndex)
		s.Require().Equal(tc.lt, tc.d1.LT(tc.d2), "LT result is incorrect, tc %d", tcIndex)
		s.Require().Equal(tc.eq, tc.d1.Equal(tc.d2), "equality result is incorrect, tc %d", tcIndex)
	}
}

func (s *decimalTestSuite) TestDecsEqual() {
	tests := []struct {
		d1s, d2s []furymath.BigDec
		eq       bool
	}{
		{[]furymath.BigDec{furymath.NewBigDec(0)}, []furymath.BigDec{furymath.NewBigDec(0)}, true},
		{[]furymath.BigDec{furymath.NewBigDec(0)}, []furymath.BigDec{furymath.NewBigDec(1)}, false},
		{[]furymath.BigDec{furymath.NewBigDec(0)}, []furymath.BigDec{}, false},
		{[]furymath.BigDec{furymath.NewBigDec(0), furymath.NewBigDec(1)}, []furymath.BigDec{furymath.NewBigDec(0), furymath.NewBigDec(1)}, true},
		{[]furymath.BigDec{furymath.NewBigDec(1), furymath.NewBigDec(0)}, []furymath.BigDec{furymath.NewBigDec(1), furymath.NewBigDec(0)}, true},
		{[]furymath.BigDec{furymath.NewBigDec(1), furymath.NewBigDec(0)}, []furymath.BigDec{furymath.NewBigDec(0), furymath.NewBigDec(1)}, false},
		{[]furymath.BigDec{furymath.NewBigDec(1), furymath.NewBigDec(0)}, []furymath.BigDec{furymath.NewBigDec(1)}, false},
		{[]furymath.BigDec{furymath.NewBigDec(1), furymath.NewBigDec(2)}, []furymath.BigDec{furymath.NewBigDec(2), furymath.NewBigDec(4)}, false},
		{[]furymath.BigDec{furymath.NewBigDec(3), furymath.NewBigDec(18)}, []furymath.BigDec{furymath.NewBigDec(1), furymath.NewBigDec(6)}, false},
	}

	for tcIndex, tc := range tests {
		s.Require().Equal(tc.eq, furymath.DecsEqual(tc.d1s, tc.d2s), "equality of decional arrays is incorrect, tc %d", tcIndex)
		s.Require().Equal(tc.eq, furymath.DecsEqual(tc.d2s, tc.d1s), "equality of decional arrays is incorrect (converse), tc %d", tcIndex)
	}
}

func (s *decimalTestSuite) TestArithmetic() {
	tests := []struct {
		d1, d2                                furymath.BigDec
		expMul, expMulTruncate                furymath.BigDec
		expQuo, expQuoRoundUp, expQuoTruncate furymath.BigDec
		expAdd, expSub                        furymath.BigDec
	}{
		//  d1         d2         MUL    MulTruncate    QUO    QUORoundUp QUOTrunctate  ADD         SUB
		{furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0)},
		{furymath.NewBigDec(1), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(1), furymath.NewBigDec(1)},
		{furymath.NewBigDec(0), furymath.NewBigDec(1), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(1), furymath.NewBigDec(-1)},
		{furymath.NewBigDec(0), furymath.NewBigDec(-1), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(-1), furymath.NewBigDec(1)},
		{furymath.NewBigDec(-1), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(0), furymath.NewBigDec(-1), furymath.NewBigDec(-1)},

		{furymath.NewBigDec(1), furymath.NewBigDec(1), furymath.NewBigDec(1), furymath.NewBigDec(1), furymath.NewBigDec(1), furymath.NewBigDec(1), furymath.NewBigDec(1), furymath.NewBigDec(2), furymath.NewBigDec(0)},
		{furymath.NewBigDec(-1), furymath.NewBigDec(-1), furymath.NewBigDec(1), furymath.NewBigDec(1), furymath.NewBigDec(1), furymath.NewBigDec(1), furymath.NewBigDec(1), furymath.NewBigDec(-2), furymath.NewBigDec(0)},
		{furymath.NewBigDec(1), furymath.NewBigDec(-1), furymath.NewBigDec(-1), furymath.NewBigDec(-1), furymath.NewBigDec(-1), furymath.NewBigDec(-1), furymath.NewBigDec(-1), furymath.NewBigDec(0), furymath.NewBigDec(2)},
		{furymath.NewBigDec(-1), furymath.NewBigDec(1), furymath.NewBigDec(-1), furymath.NewBigDec(-1), furymath.NewBigDec(-1), furymath.NewBigDec(-1), furymath.NewBigDec(-1), furymath.NewBigDec(0), furymath.NewBigDec(-2)},

		{
			furymath.NewBigDec(3), furymath.NewBigDec(7), furymath.NewBigDec(21), furymath.NewBigDec(21),
			furymath.MustNewDecFromStr("0.428571428571428571428571428571428571"), furymath.MustNewDecFromStr("0.428571428571428571428571428571428572"), furymath.MustNewDecFromStr("0.428571428571428571428571428571428571"),
			furymath.NewBigDec(10), furymath.NewBigDec(-4),
		},
		{
			furymath.NewBigDec(2), furymath.NewBigDec(4), furymath.NewBigDec(8), furymath.NewBigDec(8), furymath.NewDecWithPrec(5, 1), furymath.NewDecWithPrec(5, 1), furymath.NewDecWithPrec(5, 1),
			furymath.NewBigDec(6), furymath.NewBigDec(-2),
		},

		{furymath.NewBigDec(100), furymath.NewBigDec(100), furymath.NewBigDec(10000), furymath.NewBigDec(10000), furymath.NewBigDec(1), furymath.NewBigDec(1), furymath.NewBigDec(1), furymath.NewBigDec(200), furymath.NewBigDec(0)},

		{
			furymath.NewDecWithPrec(15, 1), furymath.NewDecWithPrec(15, 1), furymath.NewDecWithPrec(225, 2), furymath.NewDecWithPrec(225, 2),
			furymath.NewBigDec(1), furymath.NewBigDec(1), furymath.NewBigDec(1), furymath.NewBigDec(3), furymath.NewBigDec(0),
		},
		{
			furymath.NewDecWithPrec(3333, 4), furymath.NewDecWithPrec(333, 4), furymath.NewDecWithPrec(1109889, 8), furymath.NewDecWithPrec(1109889, 8),
			furymath.MustNewDecFromStr("10.009009009009009009009009009009009009"), furymath.MustNewDecFromStr("10.009009009009009009009009009009009010"), furymath.MustNewDecFromStr("10.009009009009009009009009009009009009"),
			furymath.NewDecWithPrec(3666, 4), furymath.NewDecWithPrec(3, 1),
		},
	}

	for tcIndex, tc := range tests {
		tc := tc
		resAdd := tc.d1.Add(tc.d2)
		resSub := tc.d1.Sub(tc.d2)
		resMul := tc.d1.Mul(tc.d2)
		resMulTruncate := tc.d1.MulTruncate(tc.d2)
		s.Require().True(tc.expAdd.Equal(resAdd), "exp %v, res %v, tc %d", tc.expAdd, resAdd, tcIndex)
		s.Require().True(tc.expSub.Equal(resSub), "exp %v, res %v, tc %d", tc.expSub, resSub, tcIndex)
		s.Require().True(tc.expMul.Equal(resMul), "exp %v, res %v, tc %d", tc.expMul, resMul, tcIndex)
		s.Require().True(tc.expMulTruncate.Equal(resMulTruncate), "exp %v, res %v, tc %d", tc.expMulTruncate, resMulTruncate, tcIndex)

		if tc.d2.IsZero() { // panic for divide by zero
			s.Require().Panics(func() { tc.d1.Quo(tc.d2) })
		} else {
			resQuo := tc.d1.Quo(tc.d2)
			s.Require().True(tc.expQuo.Equal(resQuo), "exp %v, res %v, tc %d", tc.expQuo.String(), resQuo.String(), tcIndex)

			resQuoRoundUp := tc.d1.QuoRoundUp(tc.d2)
			s.Require().True(tc.expQuoRoundUp.Equal(resQuoRoundUp), "exp %v, res %v, tc %d",
				tc.expQuoRoundUp.String(), resQuoRoundUp.String(), tcIndex)

			resQuoTruncate := tc.d1.QuoTruncate(tc.d2)
			s.Require().True(tc.expQuoTruncate.Equal(resQuoTruncate), "exp %v, res %v, tc %d",
				tc.expQuoTruncate.String(), resQuoTruncate.String(), tcIndex)
		}
	}
}

func (s *decimalTestSuite) TestBankerRoundChop() {
	tests := []struct {
		d1  furymath.BigDec
		exp int64
	}{
		{s.MustNewDecFromStr("0.25"), 0},
		{s.MustNewDecFromStr("0"), 0},
		{s.MustNewDecFromStr("1"), 1},
		{s.MustNewDecFromStr("0.75"), 1},
		{s.MustNewDecFromStr("0.5"), 0},
		{s.MustNewDecFromStr("7.5"), 8},
		{s.MustNewDecFromStr("1.5"), 2},
		{s.MustNewDecFromStr("2.5"), 2},
		{s.MustNewDecFromStr("0.545"), 1}, // 0.545-> 1 even though 5 is first decimal and 1 not even
		{s.MustNewDecFromStr("1.545"), 2},
	}

	for tcIndex, tc := range tests {
		resNeg := tc.d1.Neg().RoundInt64()
		s.Require().Equal(-1*tc.exp, resNeg, "negative tc %d", tcIndex)

		resPos := tc.d1.RoundInt64()
		s.Require().Equal(tc.exp, resPos, "positive tc %d", tcIndex)
	}
}

func (s *decimalTestSuite) TestTruncate() {
	tests := []struct {
		d1  furymath.BigDec
		exp int64
	}{
		{s.MustNewDecFromStr("0"), 0},
		{s.MustNewDecFromStr("0.25"), 0},
		{s.MustNewDecFromStr("0.75"), 0},
		{s.MustNewDecFromStr("1"), 1},
		{s.MustNewDecFromStr("1.5"), 1},
		{s.MustNewDecFromStr("7.5"), 7},
		{s.MustNewDecFromStr("7.6"), 7},
		{s.MustNewDecFromStr("7.4"), 7},
		{s.MustNewDecFromStr("100.1"), 100},
		{s.MustNewDecFromStr("1000.1"), 1000},
	}

	for tcIndex, tc := range tests {
		resNeg := tc.d1.Neg().TruncateInt64()
		s.Require().Equal(-1*tc.exp, resNeg, "negative tc %d", tcIndex)

		resPos := tc.d1.TruncateInt64()
		s.Require().Equal(tc.exp, resPos, "positive tc %d", tcIndex)
	}
}

func (s *decimalTestSuite) TestStringOverflow() {
	// two random 64 bit primes
	dec1, err := furymath.NewDecFromStr("51643150036226787134389711697696177267")
	s.Require().NoError(err)
	dec2, err := furymath.NewDecFromStr("-31798496660535729618459429845579852627")
	s.Require().NoError(err)
	dec3 := dec1.Add(dec2)
	s.Require().Equal(
		"19844653375691057515930281852116324640.000000000000000000000000000000000000",
		dec3.String(),
	)
}

func (s *decimalTestSuite) TestDecMulInt() {
	tests := []struct {
		sdkDec furymath.BigDec
		sdkInt furymath.BigInt
		want   furymath.BigDec
	}{
		{furymath.NewBigDec(10), furymath.NewInt(2), furymath.NewBigDec(20)},
		{furymath.NewBigDec(1000000), furymath.NewInt(100), furymath.NewBigDec(100000000)},
		{furymath.NewDecWithPrec(1, 1), furymath.NewInt(10), furymath.NewBigDec(1)},
		{furymath.NewDecWithPrec(1, 5), furymath.NewInt(20), furymath.NewDecWithPrec(2, 4)},
	}
	for i, tc := range tests {
		got := tc.sdkDec.MulInt(tc.sdkInt)
		s.Require().Equal(tc.want, got, "Incorrect result on test case %d", i)
	}
}

func (s *decimalTestSuite) TestDecCeil() {
	testCases := []struct {
		input    furymath.BigDec
		expected furymath.BigDec
	}{
		{furymath.MustNewDecFromStr("0.001"), furymath.NewBigDec(1)},   // 0.001 => 1.0
		{furymath.MustNewDecFromStr("-0.001"), furymath.ZeroDec()},     // -0.001 => 0.0
		{furymath.ZeroDec(), furymath.ZeroDec()},                       // 0.0 => 0.0
		{furymath.MustNewDecFromStr("0.9"), furymath.NewBigDec(1)},     // 0.9 => 1.0
		{furymath.MustNewDecFromStr("4.001"), furymath.NewBigDec(5)},   // 4.001 => 5.0
		{furymath.MustNewDecFromStr("-4.001"), furymath.NewBigDec(-4)}, // -4.001 => -4.0
		{furymath.MustNewDecFromStr("4.7"), furymath.NewBigDec(5)},     // 4.7 => 5.0
		{furymath.MustNewDecFromStr("-4.7"), furymath.NewBigDec(-4)},   // -4.7 => -4.0
	}

	for i, tc := range testCases {
		res := tc.input.Ceil()
		s.Require().Equal(tc.expected, res, "unexpected result for test case %d, input: %v", i, tc.input)
	}
}

func (s *decimalTestSuite) TestApproxRoot() {
	testCases := []struct {
		input    furymath.BigDec
		root     uint64
		expected furymath.BigDec
	}{
		{furymath.OneDec(), 10, furymath.OneDec()},                                                                            // 1.0 ^ (0.1) => 1.0
		{furymath.NewDecWithPrec(25, 2), 2, furymath.NewDecWithPrec(5, 1)},                                                    // 0.25 ^ (0.5) => 0.5
		{furymath.NewDecWithPrec(4, 2), 2, furymath.NewDecWithPrec(2, 1)},                                                     // 0.04 ^ (0.5) => 0.2
		{furymath.NewDecFromInt(furymath.NewInt(27)), 3, furymath.NewDecFromInt(furymath.NewInt(3))},                          // 27 ^ (1/3) => 3
		{furymath.NewDecFromInt(furymath.NewInt(-81)), 4, furymath.NewDecFromInt(furymath.NewInt(-3))},                        // -81 ^ (0.25) => -3
		{furymath.NewDecFromInt(furymath.NewInt(2)), 2, furymath.MustNewDecFromStr("1.414213562373095048801688724209698079")}, // 2 ^ (0.5) => 1.414213562373095048801688724209698079
		{furymath.NewDecWithPrec(1005, 3), 31536000, furymath.MustNewDecFromStr("1.000000000158153903837946258002096839")},    // 1.005 ^ (1/31536000) ≈ 1.000000000158153903837946258002096839
		{furymath.SmallestDec(), 2, furymath.NewDecWithPrec(1, 18)},                                                           // 1e-36 ^ (0.5) => 1e-18
		{furymath.SmallestDec(), 3, furymath.MustNewDecFromStr("0.000000000001000000000000000002431786")},                     // 1e-36 ^ (1/3) => 1e-12
		{furymath.NewDecWithPrec(1, 8), 3, furymath.MustNewDecFromStr("0.002154434690031883721759293566519280")},              // 1e-8 ^ (1/3) ≈ 0.002154434690031883721759293566519
	}

	// In the case of 1e-8 ^ (1/3), the result repeats every 5 iterations starting from iteration 24
	// (i.e. 24, 29, 34, ... give the same result) and never converges enough. The maximum number of
	// iterations (100) causes the result at iteration 100 to be returned, regardless of convergence.

	for i, tc := range testCases {
		res, err := tc.input.ApproxRoot(tc.root)
		s.Require().NoError(err)
		s.Require().True(tc.expected.Sub(res).Abs().LTE(furymath.SmallestDec()), "unexpected result for test case %d, input: %v", i, tc.input)
	}
}

func (s *decimalTestSuite) TestApproxSqrt() {
	testCases := []struct {
		input    furymath.BigDec
		expected furymath.BigDec
	}{
		{furymath.OneDec(), furymath.OneDec()},                                                                             // 1.0 => 1.0
		{furymath.NewDecWithPrec(25, 2), furymath.NewDecWithPrec(5, 1)},                                                    // 0.25 => 0.5
		{furymath.NewDecWithPrec(4, 2), furymath.NewDecWithPrec(2, 1)},                                                     // 0.09 => 0.3
		{furymath.NewDecFromInt(furymath.NewInt(9)), furymath.NewDecFromInt(furymath.NewInt(3))},                           // 9 => 3
		{furymath.NewDecFromInt(furymath.NewInt(-9)), furymath.NewDecFromInt(furymath.NewInt(-3))},                         // -9 => -3
		{furymath.NewDecFromInt(furymath.NewInt(2)), furymath.MustNewDecFromStr("1.414213562373095048801688724209698079")}, // 2 => 1.414213562373095048801688724209698079
	}

	for i, tc := range testCases {
		res, err := tc.input.ApproxSqrt()
		s.Require().NoError(err)
		s.Require().Equal(tc.expected, res, "unexpected result for test case %d, input: %v", i, tc.input)
	}
}

func (s *decimalTestSuite) TestDecSortableBytes() {
	tests := []struct {
		d    furymath.BigDec
		want []byte
	}{
		{furymath.NewBigDec(0), []byte("000000000000000000000000000000000000.000000000000000000000000000000000000")},
		{furymath.NewBigDec(1), []byte("000000000000000000000000000000000001.000000000000000000000000000000000000")},
		{furymath.NewBigDec(10), []byte("000000000000000000000000000000000010.000000000000000000000000000000000000")},
		{furymath.NewBigDec(12340), []byte("000000000000000000000000000000012340.000000000000000000000000000000000000")},
		{furymath.NewDecWithPrec(12340, 4), []byte("000000000000000000000000000000000001.234000000000000000000000000000000000")},
		{furymath.NewDecWithPrec(12340, 5), []byte("000000000000000000000000000000000000.123400000000000000000000000000000000")},
		{furymath.NewDecWithPrec(12340, 8), []byte("000000000000000000000000000000000000.000123400000000000000000000000000000")},
		{furymath.NewDecWithPrec(1009009009009009009, 17), []byte("000000000000000000000000000000000010.090090090090090090000000000000000000")},
		{furymath.NewDecWithPrec(-1009009009009009009, 17), []byte("-000000000000000000000000000000000010.090090090090090090000000000000000000")},
		{furymath.MustNewDecFromStr("1000000000000000000000000000000000000"), []byte("max")},
		{furymath.MustNewDecFromStr("-1000000000000000000000000000000000000"), []byte("--")},
	}
	for tcIndex, tc := range tests {
		s.Require().Equal(tc.want, furymath.SortableDecBytes(tc.d), "bad String(), index: %v", tcIndex)
	}

	s.Require().Panics(func() { furymath.SortableDecBytes(furymath.MustNewDecFromStr("1000000000000000000000000000000000001")) })
	s.Require().Panics(func() {
		furymath.SortableDecBytes(furymath.MustNewDecFromStr("-1000000000000000000000000000000000001"))
	})
}

func (s *decimalTestSuite) TestDecEncoding() {
	testCases := []struct {
		input   furymath.BigDec
		rawBz   string
		jsonStr string
		yamlStr string
	}{
		{
			furymath.NewBigDec(0), "30",
			"\"0.000000000000000000000000000000000000\"",
			"\"0.000000000000000000000000000000000000\"\n",
		},
		{
			furymath.NewDecWithPrec(4, 2),
			"3430303030303030303030303030303030303030303030303030303030303030303030",
			"\"0.040000000000000000000000000000000000\"",
			"\"0.040000000000000000000000000000000000\"\n",
		},
		{
			furymath.NewDecWithPrec(-4, 2),
			"2D3430303030303030303030303030303030303030303030303030303030303030303030",
			"\"-0.040000000000000000000000000000000000\"",
			"\"-0.040000000000000000000000000000000000\"\n",
		},
		{
			furymath.MustNewDecFromStr("1.414213562373095048801688724209698079"),
			"31343134323133353632333733303935303438383031363838373234323039363938303739",
			"\"1.414213562373095048801688724209698079\"",
			"\"1.414213562373095048801688724209698079\"\n",
		},
		{
			furymath.MustNewDecFromStr("-1.414213562373095048801688724209698079"),
			"2D31343134323133353632333733303935303438383031363838373234323039363938303739",
			"\"-1.414213562373095048801688724209698079\"",
			"\"-1.414213562373095048801688724209698079\"\n",
		},
	}

	for _, tc := range testCases {
		bz, err := tc.input.Marshal()
		s.Require().NoError(err)
		s.Require().Equal(tc.rawBz, fmt.Sprintf("%X", bz))

		var other furymath.BigDec
		s.Require().NoError((&other).Unmarshal(bz))
		s.Require().True(tc.input.Equal(other))

		bz, err = json.Marshal(tc.input)
		s.Require().NoError(err)
		s.Require().Equal(tc.jsonStr, string(bz))
		s.Require().NoError(json.Unmarshal(bz, &other))
		s.Require().True(tc.input.Equal(other))

		bz, err = yaml.Marshal(tc.input)
		s.Require().NoError(err)
		s.Require().Equal(tc.yamlStr, string(bz))
	}
}

// Showcase that different orders of operations causes different results.
func (s *decimalTestSuite) TestOperationOrders() {
	n1 := furymath.NewBigDec(10)
	n2 := furymath.NewBigDec(1000000010)
	s.Require().Equal(n1.Mul(n2).Quo(n2), furymath.NewBigDec(10))
	s.Require().NotEqual(n1.Mul(n2).Quo(n2), n1.Quo(n2).Mul(n2))
}

func BenchmarkMarshalTo(b *testing.B) {
	b.ReportAllocs()
	bis := []struct {
		in   furymath.BigDec
		want []byte
	}{
		{
			furymath.NewBigDec(1e8), []byte{
				0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
				0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
				0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
			},
		},
		{furymath.NewBigDec(0), []byte{0x30}},
	}
	data := make([]byte, 100)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, bi := range bis {
			if n, err := bi.in.MarshalTo(data); err != nil {
				b.Fatal(err)
			} else {
				if !bytes.Equal(data[:n], bi.want) {
					b.Fatalf("Mismatch\nGot:  % x\nWant: % x\n", data[:n], bi.want)
				}
			}
		}
	}
}

func (s *decimalTestSuite) TestLog2() {
	var expectedErrTolerance = furymath.MustNewDecFromStr("0.000000000000000000000000000000000100")

	tests := map[string]struct {
		initialValue furymath.BigDec
		expected     furymath.BigDec

		expectedPanic bool
	}{
		"log_2{-1}; invalid; panic": {
			initialValue:  furymath.OneDec().Neg(),
			expectedPanic: true,
		},
		"log_2{0}; invalid; panic": {
			initialValue:  furymath.ZeroDec(),
			expectedPanic: true,
		},
		"log_2{0.001} = -9.965784284662087043610958288468170528": {
			initialValue: furymath.MustNewDecFromStr("0.001"),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+0.999912345+with+33+digits
			expected: furymath.MustNewDecFromStr("-9.965784284662087043610958288468170528"),
		},
		"log_2{0.56171821941421412902170941} = -0.832081497183140708984033250637831402": {
			initialValue: furymath.MustNewDecFromStr("0.56171821941421412902170941"),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+0.56171821941421412902170941+with+36+digits
			expected: furymath.MustNewDecFromStr("-0.832081497183140708984033250637831402"),
		},
		"log_2{0.999912345} = -0.000126464976533858080645902722235833": {
			initialValue: furymath.MustNewDecFromStr("0.999912345"),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+0.999912345+with+37+digits
			expected: furymath.MustNewDecFromStr("-0.000126464976533858080645902722235833"),
		},
		"log_2{1} = 0": {
			initialValue: furymath.NewBigDec(1),
			expected:     furymath.NewBigDec(0),
		},
		"log_2{2} = 1": {
			initialValue: furymath.NewBigDec(2),
			expected:     furymath.NewBigDec(1),
		},
		"log_2{7} = 2.807354922057604107441969317231830809": {
			initialValue: furymath.NewBigDec(7),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+7+37+digits
			expected: furymath.MustNewDecFromStr("2.807354922057604107441969317231830809"),
		},
		"log_2{512} = 9": {
			initialValue: furymath.NewBigDec(512),
			expected:     furymath.NewBigDec(9),
		},
		"log_2{580} = 9.179909090014934468590092754117374938": {
			initialValue: furymath.NewBigDec(580),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+600+37+digits
			expected: furymath.MustNewDecFromStr("9.179909090014934468590092754117374938"),
		},
		"log_2{1024} = 10": {
			initialValue: furymath.NewBigDec(1024),
			expected:     furymath.NewBigDec(10),
		},
		"log_2{1024.987654321} = 10.001390817654141324352719749259888355": {
			initialValue: furymath.NewDecWithPrec(1024987654321, 9),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+1024.987654321+38+digits
			expected: furymath.MustNewDecFromStr("10.001390817654141324352719749259888355"),
		},
		"log_2{912648174127941279170121098210.92821920190204131121} = 99.525973560175362367047484597337715868": {
			initialValue: furymath.MustNewDecFromStr("912648174127941279170121098210.92821920190204131121"),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+912648174127941279170121098210.92821920190204131121+38+digits
			expected: furymath.MustNewDecFromStr("99.525973560175362367047484597337715868"),
		},
		"log_2{Max Spot Price} = 128": {
			initialValue: furymath.BigDecFromSDKDec(furymath.MaxSpotPrice), // 2^128 - 1
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+%28%282%5E128%29+-+1%29+38+digits
			expected: furymath.MustNewDecFromStr("128"),
		},
		// The value tested below is: gammtypes.MaxSpotPrice * 0.99 = (2^128 - 1) * 0.99
		"log_2{336879543251729078828740861357450529340.45} = 127.98550043030488492336620207564264562": {
			initialValue: furymath.MustNewDecFromStr("336879543251729078828740861357450529340.45"),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+%28%28%282%5E128%29+-+1%29*0.99%29++38+digits
			expected: furymath.MustNewDecFromStr("127.98550043030488492336620207564264562"),
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			furymath.ConditionalPanic(s.T(), tc.expectedPanic, func() {
				// Create a copy to test that the original was not modified.
				// That is, that LogbBase2() is non-mutative.
				initialCopy := tc.initialValue.Clone()

				res := tc.initialValue.LogBase2()
				require.True(furymath.DecApproxEq(s.T(), tc.expected, res, expectedErrTolerance))
				require.Equal(s.T(), initialCopy, tc.initialValue)
			})
		})
	}
}

func (s *decimalTestSuite) TestLn() {
	var expectedErrTolerance = furymath.MustNewDecFromStr("0.000000000000000000000000000000000100")

	tests := map[string]struct {
		initialValue furymath.BigDec
		expected     furymath.BigDec

		expectedPanic bool
	}{
		"log_e{-1}; invalid; panic": {
			initialValue:  furymath.OneDec().Neg(),
			expectedPanic: true,
		},
		"log_e{0}; invalid; panic": {
			initialValue:  furymath.ZeroDec(),
			expectedPanic: true,
		},
		"log_e{0.001} = -6.90775527898213705205397436405309262": {
			initialValue: furymath.MustNewDecFromStr("0.001"),
			// From: https://www.wolframalpha.com/input?i=log0.001+to+36+digits+with+36+decimals
			expected: furymath.MustNewDecFromStr("-6.90775527898213705205397436405309262"),
		},
		"log_e{0.56171821941421412902170941} = -0.576754943768592057376050794884207180": {
			initialValue: furymath.MustNewDecFromStr("0.56171821941421412902170941"),
			// From: https://www.wolframalpha.com/input?i=log0.56171821941421412902170941+to+36+digits
			expected: furymath.MustNewDecFromStr("-0.576754943768592057376050794884207180"),
		},
		"log_e{0.999912345} = -0.000087658841924023373535614212850888": {
			initialValue: furymath.MustNewDecFromStr("0.999912345"),
			// From: https://www.wolframalpha.com/input?i=log0.999912345+to+32+digits
			expected: furymath.MustNewDecFromStr("-0.000087658841924023373535614212850888"),
		},
		"log_e{1} = 0": {
			initialValue: furymath.NewBigDec(1),
			expected:     furymath.NewBigDec(0),
		},
		"log_e{e} = 1": {
			initialValue: furymath.MustNewDecFromStr("2.718281828459045235360287471352662498"),
			// From: https://www.wolframalpha.com/input?i=e+with+36+decimals
			expected: furymath.NewBigDec(1),
		},
		"log_e{7} = 1.945910149055313305105352743443179730": {
			initialValue: furymath.NewBigDec(7),
			// From: https://www.wolframalpha.com/input?i=log7+up+to+36+decimals
			expected: furymath.MustNewDecFromStr("1.945910149055313305105352743443179730"),
		},
		"log_e{512} = 6.238324625039507784755089093123589113": {
			initialValue: furymath.NewBigDec(512),
			// From: https://www.wolframalpha.com/input?i=log512+up+to+36+decimals
			expected: furymath.MustNewDecFromStr("6.238324625039507784755089093123589113"),
		},
		"log_e{580} = 6.36302810354046502061849560850445238": {
			initialValue: furymath.NewBigDec(580),
			// From: https://www.wolframalpha.com/input?i=log580+up+to+36+decimals
			expected: furymath.MustNewDecFromStr("6.36302810354046502061849560850445238"),
		},
		"log_e{1024.987654321} = 6.93243584693509415029056534690631614": {
			initialValue: furymath.NewDecWithPrec(1024987654321, 9),
			// From: https://www.wolframalpha.com/input?i=log1024.987654321+to+36+digits
			expected: furymath.MustNewDecFromStr("6.93243584693509415029056534690631614"),
		},
		"log_e{912648174127941279170121098210.92821920190204131121} = 68.986147965719214790400745338243805015": {
			initialValue: furymath.MustNewDecFromStr("912648174127941279170121098210.92821920190204131121"),
			// From: https://www.wolframalpha.com/input?i=log912648174127941279170121098210.92821920190204131121+to+38+digits
			expected: furymath.MustNewDecFromStr("68.986147965719214790400745338243805015"),
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			furymath.ConditionalPanic(s.T(), tc.expectedPanic, func() {
				// Create a copy to test that the original was not modified.
				// That is, that Ln() is non-mutative.
				initialCopy := tc.initialValue.Clone()

				res := tc.initialValue.Ln()
				require.True(furymath.DecApproxEq(s.T(), tc.expected, res, expectedErrTolerance))
				require.Equal(s.T(), initialCopy, tc.initialValue)
			})
		})
	}
}

func (s *decimalTestSuite) TestTickLog() {
	tests := map[string]struct {
		initialValue furymath.BigDec
		expected     furymath.BigDec

		expectedErrTolerance furymath.BigDec
		expectedPanic        bool
	}{
		"log_1.0001{-1}; invalid; panic": {
			initialValue:  furymath.OneDec().Neg(),
			expectedPanic: true,
		},
		"log_1.0001{0}; invalid; panic": {
			initialValue:  furymath.ZeroDec(),
			expectedPanic: true,
		},
		"log_1.0001{0.001} = -69081.006609899112313305835611219486392199": {
			initialValue: furymath.MustNewDecFromStr("0.001"),
			// From: https://www.wolframalpha.com/input?i=log_1.0001%280.001%29+to+41+digits
			expectedErrTolerance: furymath.MustNewDecFromStr("0.000000000000000000000000000143031879"),
			expected:             furymath.MustNewDecFromStr("-69081.006609899112313305835611219486392199"),
		},
		"log_1.0001{0.999912345} = -0.876632247930741919880461740717176538": {
			initialValue: furymath.MustNewDecFromStr("0.999912345"),
			// From: https://www.wolframalpha.com/input?i=log_1.0001%280.999912345%29+to+36+digits
			expectedErrTolerance: furymath.MustNewDecFromStr("0.000000000000000000000000000000138702"),
			expected:             furymath.MustNewDecFromStr("-0.876632247930741919880461740717176538"),
		},
		"log_1.0001{1} = 0": {
			initialValue: furymath.NewBigDec(1),

			expectedErrTolerance: furymath.ZeroDec(),
			expected:             furymath.NewBigDec(0),
		},
		"log_1.0001{1.0001} = 1": {
			initialValue: furymath.MustNewDecFromStr("1.0001"),

			expectedErrTolerance: furymath.MustNewDecFromStr("0.000000000000000000000000000000152500"),
			expected:             furymath.OneDec(),
		},
		"log_1.0001{512} = 62386.365360724158196763710649998441051753": {
			initialValue: furymath.NewBigDec(512),
			// From: https://www.wolframalpha.com/input?i=log_1.0001%28512%29+to+41+digits
			expectedErrTolerance: furymath.MustNewDecFromStr("0.000000000000000000000000000129292137"),
			expected:             furymath.MustNewDecFromStr("62386.365360724158196763710649998441051753"),
		},
		"log_1.0001{1024.987654321} = 69327.824629506998657531621822514042777198": {
			initialValue: furymath.NewDecWithPrec(1024987654321, 9),
			// From: https://www.wolframalpha.com/input?i=log_1.0001%281024.987654321%29+to+41+digits
			expectedErrTolerance: furymath.MustNewDecFromStr("0.000000000000000000000000000143836264"),
			expected:             furymath.MustNewDecFromStr("69327.824629506998657531621822514042777198"),
		},
		"log_1.0001{912648174127941279170121098210.92821920190204131121} = 689895.972156319183538389792485913311778672": {
			initialValue: furymath.MustNewDecFromStr("912648174127941279170121098210.92821920190204131121"),
			// From: https://www.wolframalpha.com/input?i=log_1.0001%28912648174127941279170121098210.92821920190204131121%29+to+42+digits
			expectedErrTolerance: furymath.MustNewDecFromStr("0.000000000000000000000000001429936067"),
			expected:             furymath.MustNewDecFromStr("689895.972156319183538389792485913311778672"),
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			furymath.ConditionalPanic(s.T(), tc.expectedPanic, func() {
				// Create a copy to test that the original was not modified.
				// That is, that Ln() is non-mutative.
				initialCopy := tc.initialValue.Clone()

				res := tc.initialValue.TickLog()
				fmt.Println(name, res.Sub(tc.expected).Abs())
				require.True(furymath.DecApproxEq(s.T(), tc.expected, res, tc.expectedErrTolerance))
				require.Equal(s.T(), initialCopy, tc.initialValue)
			})
		})
	}
}

func (s *decimalTestSuite) TestCustomBaseLog() {
	tests := map[string]struct {
		initialValue furymath.BigDec
		base         furymath.BigDec

		expected             furymath.BigDec
		expectedErrTolerance furymath.BigDec

		expectedPanic bool
	}{
		"log_2{-1}: normal base, invalid argument - panics": {
			initialValue:  furymath.NewBigDec(-1),
			base:          furymath.NewBigDec(2),
			expectedPanic: true,
		},
		"log_2{0}: normal base, invalid argument - panics": {
			initialValue:  furymath.NewBigDec(0),
			base:          furymath.NewBigDec(2),
			expectedPanic: true,
		},
		"log_(-1)(2): invalid base, normal argument - panics": {
			initialValue:  furymath.NewBigDec(2),
			base:          furymath.NewBigDec(-1),
			expectedPanic: true,
		},
		"log_1(2): base cannot equal to 1 - panics": {
			initialValue:  furymath.NewBigDec(2),
			base:          furymath.NewBigDec(1),
			expectedPanic: true,
		},
		"log_30(100) = 1.353984985057691049642502891262784015": {
			initialValue: furymath.NewBigDec(100),
			base:         furymath.NewBigDec(30),
			// From: https://www.wolframalpha.com/input?i=log_30%28100%29+to+37+digits
			expectedErrTolerance: furymath.ZeroDec(),
			expected:             furymath.MustNewDecFromStr("1.353984985057691049642502891262784015"),
		},
		"log_0.2(0.99) = 0.006244624769837438271878639001855450": {
			initialValue: furymath.MustNewDecFromStr("0.99"),
			base:         furymath.MustNewDecFromStr("0.2"),
			// From: https://www.wolframalpha.com/input?i=log_0.2%280.99%29+to+34+digits
			expectedErrTolerance: furymath.MustNewDecFromStr("0.000000000000000000000000000000000013"),
			expected:             furymath.MustNewDecFromStr("0.006244624769837438271878639001855450"),
		},

		"log_0.0001(500000) = -1.424742501084004701196565276318876743": {
			initialValue: furymath.NewBigDec(500000),
			base:         furymath.NewDecWithPrec(1, 4),
			// From: https://www.wolframalpha.com/input?i=log_0.0001%28500000%29+to+37+digits
			expectedErrTolerance: furymath.MustNewDecFromStr("0.000000000000000000000000000000000003"),
			expected:             furymath.MustNewDecFromStr("-1.424742501084004701196565276318876743"),
		},

		"log_500000(0.0001) = -0.701881216598197542030218906945601429": {
			initialValue: furymath.NewDecWithPrec(1, 4),
			base:         furymath.NewBigDec(500000),
			// From: https://www.wolframalpha.com/input?i=log_500000%280.0001%29+to+36+digits
			expectedErrTolerance: furymath.MustNewDecFromStr("0.000000000000000000000000000000000001"),
			expected:             furymath.MustNewDecFromStr("-0.701881216598197542030218906945601429"),
		},

		"log_10000(5000000) = 1.674742501084004701196565276318876743": {
			initialValue: furymath.NewBigDec(5000000),
			base:         furymath.NewBigDec(10000),
			// From: https://www.wolframalpha.com/input?i=log_10000%285000000%29+to+37+digits
			expectedErrTolerance: furymath.MustNewDecFromStr("0.000000000000000000000000000000000002"),
			expected:             furymath.MustNewDecFromStr("1.674742501084004701196565276318876743"),
		},
		"log_0.123456789(1) = 0": {
			initialValue: furymath.OneDec(),
			base:         furymath.MustNewDecFromStr("0.123456789"),

			expectedErrTolerance: furymath.ZeroDec(),
			expected:             furymath.ZeroDec(),
		},
		"log_1111(1111) = 1": {
			initialValue: furymath.NewBigDec(1111),
			base:         furymath.NewBigDec(1111),

			expectedErrTolerance: furymath.ZeroDec(),
			expected:             furymath.OneDec(),
		},

		"log_1.123{1024.987654321} = 59.760484327223888489694630378785099461": {
			initialValue: furymath.NewDecWithPrec(1024987654321, 9),
			base:         furymath.NewDecWithPrec(1123, 3),
			// From: https://www.wolframalpha.com/input?i=log_1.123%281024.987654321%29+to+38+digits
			expectedErrTolerance: furymath.MustNewDecFromStr("0.000000000000000000000000000000007686"),
			expected:             furymath.MustNewDecFromStr("59.760484327223888489694630378785099461"),
		},

		"log_1.123{912648174127941279170121098210.92821920190204131121} = 594.689327867863079177915648832621538986": {
			initialValue: furymath.MustNewDecFromStr("912648174127941279170121098210.92821920190204131121"),
			base:         furymath.NewDecWithPrec(1123, 3),
			// From: https://www.wolframalpha.com/input?i=log_1.123%28912648174127941279170121098210.92821920190204131121%29+to+39+digits
			expectedErrTolerance: furymath.MustNewDecFromStr("0.000000000000000000000000000000077705"),
			expected:             furymath.MustNewDecFromStr("594.689327867863079177915648832621538986"),
		},
	}
	for name, tc := range tests {
		s.Run(name, func() {
			furymath.ConditionalPanic(s.T(), tc.expectedPanic, func() {
				// Create a copy to test that the original was not modified.
				// That is, that Ln() is non-mutative.
				initialCopy := tc.initialValue.Clone()
				res := tc.initialValue.CustomBaseLog(tc.base)
				require.True(furymath.DecApproxEq(s.T(), tc.expected, res, tc.expectedErrTolerance))
				require.Equal(s.T(), initialCopy, tc.initialValue)
			})
		})
	}
}

func (s *decimalTestSuite) TestPowerInteger() {
	var expectedErrTolerance = furymath.MustNewDecFromStr("0.000000000000000000000000000000100000")

	tests := map[string]struct {
		base           furymath.BigDec
		exponent       uint64
		expectedResult furymath.BigDec

		expectedToleranceOverwrite furymath.BigDec
	}{
		"0^2": {
			base:     furymath.ZeroDec(),
			exponent: 2,

			expectedResult: furymath.ZeroDec(),
		},
		"1^2": {
			base:     furymath.OneDec(),
			exponent: 2,

			expectedResult: furymath.OneDec(),
		},
		"4^4": {
			base:     furymath.MustNewDecFromStr("4"),
			exponent: 4,

			expectedResult: furymath.MustNewDecFromStr("256"),
		},
		"5^3": {
			base:     furymath.MustNewDecFromStr("5"),
			exponent: 4,

			expectedResult: furymath.MustNewDecFromStr("625"),
		},
		"e^10": {
			base:     furymath.EulersNumber,
			exponent: 10,

			// https://www.wolframalpha.com/input?i=e%5E10+41+digits
			expectedResult: furymath.MustNewDecFromStr("22026.465794806716516957900645284244366354"),
		},
		"geom twap overflow: 2^log_2{max spot price + 1}": {
			base: furymath.TwoBigDec,
			// add 1 for simplicity of calculation to isolate overflow.
			exponent: uint64(furymath.BigDecFromSDKDec(furymath.MaxSpotPrice).Add(furymath.OneDec()).LogBase2().TruncateInt().Uint64()),

			// https://www.wolframalpha.com/input?i=2%5E%28floor%28+log+base+2+%282%5E128%29%29%29+++39+digits
			expectedResult: furymath.MustNewDecFromStr("340282366920938463463374607431768211456"),
		},
		"geom twap overflow: 2^log_2{max spot price}": {
			base:     furymath.TwoBigDec,
			exponent: uint64(furymath.BigDecFromSDKDec(furymath.MaxSpotPrice).LogBase2().TruncateInt().Uint64()),

			// https://www.wolframalpha.com/input?i=2%5E%28floor%28+log+base+2+%282%5E128+-+1%29%29%29+++39+digits
			expectedResult: furymath.MustNewDecFromStr("170141183460469231731687303715884105728"),
		},
		"geom twap overflow: 2^log_2{max spot price / 2 - 2017}": { // 2017 is prime.
			base:     furymath.TwoBigDec,
			exponent: uint64(furymath.BigDecFromSDKDec(furymath.MaxSpotPrice.Quo(sdk.NewDec(2)).Sub(sdk.NewDec(2017))).LogBase2().TruncateInt().Uint64()),

			// https://www.wolframalpha.com/input?i=e%5E10+41+digits
			expectedResult: furymath.MustNewDecFromStr("85070591730234615865843651857942052864"),
		},

		// sdk.Dec test vectors copied from merlins-labs/cosmos-sdk:

		"1.0 ^ (10) => 1.0": {
			base:     furymath.OneDec(),
			exponent: 10,

			expectedResult: furymath.OneDec(),
		},
		"0.5 ^ 2 => 0.25": {
			base:     furymath.NewDecWithPrec(5, 1),
			exponent: 2,

			expectedResult: furymath.NewDecWithPrec(25, 2),
		},
		"0.2 ^ 2 => 0.04": {
			base:     furymath.NewDecWithPrec(2, 1),
			exponent: 2,

			expectedResult: furymath.NewDecWithPrec(4, 2),
		},
		"3 ^ 3 => 27": {
			base:     furymath.NewBigDec(3),
			exponent: 3,

			expectedResult: furymath.NewBigDec(27),
		},
		"-3 ^ 4 = 81": {
			base:     furymath.NewBigDec(-3),
			exponent: 4,

			expectedResult: furymath.NewBigDec(81),
		},
		"-3 ^ 50 = 717897987691852588770249": {
			base:     furymath.NewBigDec(-3),
			exponent: 50,

			expectedResult: furymath.MustNewDecFromStr("717897987691852588770249"),
		},
		"-3 ^ 51 = -2153693963075557766310747": {
			base:     furymath.NewBigDec(-3),
			exponent: 51,

			expectedResult: furymath.MustNewDecFromStr("-2153693963075557766310747"),
		},
		"1.414213562373095049 ^ 2 = 2": {
			base:     furymath.NewDecWithPrec(1414213562373095049, 18),
			exponent: 2,

			expectedResult:             furymath.NewBigDec(2),
			expectedToleranceOverwrite: furymath.MustNewDecFromStr("0.0000000000000000006"),
		},
	}

	for name, tc := range tests {
		tc := tc
		s.Run(name, func() {

			tolerance := expectedErrTolerance
			if !tc.expectedToleranceOverwrite.IsNil() {
				tolerance = tc.expectedToleranceOverwrite
			}

			// Main system under test
			actualResult := tc.base.PowerInteger(tc.exponent)
			require.True(furymath.DecApproxEq(s.T(), tc.expectedResult, actualResult, tolerance))

			// Secondary system under test.
			// To reduce boilerplate from the same test cases when exponent is a
			// positive integer, we also test Power().
			// Negative exponent and base are not supported for Power()
			if tc.exponent >= 0 && !tc.base.IsNegative() {
				actualResult2 := tc.base.Power(furymath.NewDecFromInt(furymath.NewIntFromUint64(tc.exponent)))
				require.True(furymath.DecApproxEq(s.T(), tc.expectedResult, actualResult2, tolerance))
			}
		})
	}
}

func (s *decimalTestSuite) TestClone() {
	tests := map[string]struct {
		startValue furymath.BigDec
	}{
		"1.1": {
			startValue: furymath.MustNewDecFromStr("1.1"),
		},
		"-3": {
			startValue: furymath.MustNewDecFromStr("-3"),
		},
		"0": {
			startValue: furymath.MustNewDecFromStr("-3"),
		},
	}

	for name, tc := range tests {
		tc := tc
		s.Run(name, func() {

			copy := tc.startValue.Clone()

			s.Require().Equal(tc.startValue, copy)

			copy.MulMut(furymath.NewBigDec(2))
			// copy and startValue do not share internals.
			s.Require().NotEqual(tc.startValue, copy)
		})
	}
}

// TestMul_Mutation tests that MulMut mutates the receiver
// while Mut is not.
func (s *decimalTestSuite) TestMul_Mutation() {

	mulBy := furymath.MustNewDecFromStr("2")

	tests := map[string]struct {
		startValue        furymath.BigDec
		expectedMulResult furymath.BigDec
	}{
		"1.1": {
			startValue:        furymath.MustNewDecFromStr("1.1"),
			expectedMulResult: furymath.MustNewDecFromStr("2.2"),
		},
		"-3": {
			startValue:        furymath.MustNewDecFromStr("-3"),
			expectedMulResult: furymath.MustNewDecFromStr("-6"),
		},
		"0": {
			startValue:        furymath.ZeroDec(),
			expectedMulResult: furymath.ZeroDec(),
		},
	}

	for name, tc := range tests {
		tc := tc
		s.Run(name, func() {
			startMut := tc.startValue.Clone()
			startNonMut := tc.startValue.Clone()

			resultMut := startMut.MulMut(mulBy)
			resultNonMut := startNonMut.Mul(mulBy)

			s.assertMutResult(tc.expectedMulResult, tc.startValue, resultMut, resultNonMut, startMut, startNonMut)
		})
	}
}

// TestPowerInteger_Mutation tests that PowerIntegerMut mutates the receiver
// while PowerInteger is not.
func (s *decimalTestSuite) TestPowerInteger_Mutation() {

	exponent := uint64(2)

	tests := map[string]struct {
		startValue     furymath.BigDec
		expectedResult furymath.BigDec
	}{
		"1": {
			startValue:     furymath.OneDec(),
			expectedResult: furymath.OneDec(),
		},
		"-3": {
			startValue:     furymath.MustNewDecFromStr("-3"),
			expectedResult: furymath.MustNewDecFromStr("9"),
		},
		"0": {
			startValue:     furymath.ZeroDec(),
			expectedResult: furymath.ZeroDec(),
		},
		"4": {
			startValue:     furymath.MustNewDecFromStr("4.5"),
			expectedResult: furymath.MustNewDecFromStr("20.25"),
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {

			startMut := tc.startValue.Clone()
			startNonMut := tc.startValue.Clone()

			resultMut := startMut.PowerIntegerMut(exponent)
			resultNonMut := startNonMut.PowerInteger(exponent)

			s.assertMutResult(tc.expectedResult, tc.startValue, resultMut, resultNonMut, startMut, startNonMut)
		})
	}
}

func (s *decimalTestSuite) TestPower() {
	tests := map[string]struct {
		base           furymath.BigDec
		exponent       furymath.BigDec
		expectedResult furymath.BigDec
		expectPanic    bool
		errTolerance   furymath.ErrTolerance
	}{
		// N.B.: integer exponents are tested under TestPowerInteger.

		"3 ^ 2 = 9 (integer base and integer exponent)": {
			base:     furymath.NewBigDec(3),
			exponent: furymath.NewBigDec(2),

			expectedResult: furymath.NewBigDec(9),

			errTolerance: zeroAdditiveErrTolerance,
		},
		"2^0.5 (base of 2 and non-integer exponent)": {
			base:     furymath.MustNewDecFromStr("2"),
			exponent: furymath.MustNewDecFromStr("0.5"),

			// https://www.wolframalpha.com/input?i=2%5E0.5+37+digits
			expectedResult: furymath.MustNewDecFromStr("1.414213562373095048801688724209698079"),

			errTolerance: furymath.ErrTolerance{
				AdditiveTolerance: minDecTolerance,
				RoundingDir:       furymath.RoundDown,
			},
		},
		"3^0.33 (integer base other than 2 and non-integer exponent)": {
			base:     furymath.MustNewDecFromStr("3"),
			exponent: furymath.MustNewDecFromStr("0.33"),

			// https://www.wolframalpha.com/input?i=3%5E0.33+37+digits
			expectedResult: furymath.MustNewDecFromStr("1.436977652184851654252692986409357265"),

			errTolerance: furymath.ErrTolerance{
				AdditiveTolerance: minDecTolerance,
				RoundingDir:       furymath.RoundDown,
			},
		},
		"e^0.98999 (non-integer base and non-integer exponent)": {
			base:     furymath.EulersNumber,
			exponent: furymath.MustNewDecFromStr("0.9899"),

			// https://www.wolframalpha.com/input?i=e%5E0.9899+37+digits
			expectedResult: furymath.MustNewDecFromStr("2.690965362357751196751808686902156603"),

			errTolerance: furymath.ErrTolerance{
				AdditiveTolerance: minDecTolerance,
				RoundingDir:       furymath.RoundUnconstrained,
			},
		},
		"10^0.001 (small non-integer exponent)": {
			base:     furymath.NewBigDec(10),
			exponent: furymath.MustNewDecFromStr("0.001"),

			// https://www.wolframalpha.com/input?i=10%5E0.001+37+digits
			expectedResult: furymath.MustNewDecFromStr("1.002305238077899671915404889328110554"),

			errTolerance: furymath.ErrTolerance{
				AdditiveTolerance: minDecTolerance,
				RoundingDir:       furymath.RoundUnconstrained,
			},
		},
		"13^100.7777 (large non-integer exponent)": {
			base:     furymath.NewBigDec(13),
			exponent: furymath.MustNewDecFromStr("100.7777"),

			// https://www.wolframalpha.com/input?i=13%5E100.7777+37+digits
			expectedResult: furymath.MustNewDecFromStr("1.822422110233759706998600329118969132").Mul(furymath.NewBigDec(10).PowerInteger(112)),

			errTolerance: furymath.ErrTolerance{
				MultiplicativeTolerance: minDecTolerance,
				RoundingDir:             furymath.RoundDown,
			},
		},
		"large non-integer exponent with large non-integer base - panics": {
			base:     furymath.MustNewDecFromStr("169.137"),
			exponent: furymath.MustNewDecFromStr("100.7777"),

			expectPanic: true,
		},
		"negative base - panic": {
			base:     furymath.NewBigDec(-3),
			exponent: furymath.MustNewDecFromStr("4"),

			expectPanic: true,
		},
		"negative exponent - panic": {
			base:     furymath.NewBigDec(1),
			exponent: furymath.MustNewDecFromStr("-4"),

			expectPanic: true,
		},
		"base < 1 - panic (see godoc)": {
			base:     furymath.NewBigDec(1).Sub(furymath.SmallestDec()),
			exponent: furymath.OneDec(),

			expectPanic: true,
		},
	}

	for name, tc := range tests {
		tc := tc
		s.Run(name, func() {
			furymath.ConditionalPanic(s.T(), tc.expectPanic, func() {
				actualResult := tc.base.Power(tc.exponent)
				s.Require().Equal(0, tc.errTolerance.CompareBigDec(tc.expectedResult, actualResult))
			})
		})
	}
}
