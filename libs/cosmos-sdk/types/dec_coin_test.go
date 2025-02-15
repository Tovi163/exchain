package types

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"testing"

	"github.com/okex/exchain/libs/cosmos-sdk/codec"

	"github.com/stretchr/testify/require"
)

func TestNewDecCoin(t *testing.T) {
	require.NotPanics(t, func() {
		NewInt64DecCoin(testDenom1, 5)
	})
	require.NotPanics(t, func() {
		NewInt64DecCoin(testDenom1, 0)
	})
	require.Panics(t, func() {
		NewInt64DecCoin(strings.ToUpper(testDenom1), 5)
	})
	require.Panics(t, func() {
		NewInt64DecCoin(testDenom1, -5)
	})
}

func TestNewDecCoinFromDec(t *testing.T) {
	require.NotPanics(t, func() {
		NewDecCoinFromDec(testDenom1, NewDec(5))
	})
	require.NotPanics(t, func() {
		NewDecCoinFromDec(testDenom1, ZeroDec())
	})
	require.Panics(t, func() {
		NewDecCoinFromDec(strings.ToUpper(testDenom1), NewDec(5))
	})
	require.Panics(t, func() {
		NewDecCoinFromDec(testDenom1, NewDec(-5))
	})
}

func TestNewDecCoinFromCoin(t *testing.T) {
	require.NotPanics(t, func() {
		NewDecCoinFromCoin(Coin{testDenom1, NewDec(5)})
	})
	require.NotPanics(t, func() {
		NewDecCoinFromCoin(Coin{testDenom1, NewDec(0)})
	})
	require.NotPanics(t, func() {
		NewDecCoinFromCoin(Coin{strings.ToUpper(testDenom1), NewDec(5)})
	})
	require.NotPanics(t, func() {
		NewDecCoinFromCoin(Coin{testDenom1, NewDec(-5)})
	})
}

func TestDecCoinIsPositive(t *testing.T) {
	dc := NewInt64DecCoin(testDenom1, 5)
	require.True(t, dc.IsPositive())

	dc = NewInt64DecCoin(testDenom1, 0)
	require.False(t, dc.IsPositive())
}

func TestAddDecCoin(t *testing.T) {
	decCoinA1 := NewDecCoinFromDec(testDenom1, NewDecWithPrec(11, 1))
	decCoinA2 := NewDecCoinFromDec(testDenom1, NewDecWithPrec(22, 1))
	decCoinB1 := NewDecCoinFromDec(testDenom2, NewDecWithPrec(11, 1))

	// regular add
	res := decCoinA1.Add(decCoinA1)
	require.Equal(t, decCoinA2, res, "sum of coins is incorrect")

	// bad denom add
	require.Panics(t, func() {
		decCoinA1.Add(decCoinB1)
	}, "expected panic on sum of different denoms")
}

func TestAddDecCoins(t *testing.T) {
	one := NewDec(1)
	zero := NewDec(0)
	two := NewDec(2)

	cases := []struct {
		inputOne DecCoins
		inputTwo DecCoins
		expected DecCoins
	}{
		{DecCoins{{testDenom1, one}, {testDenom2, one}}, DecCoins{{testDenom1, one}, {testDenom2, one}}, DecCoins{{testDenom1, two}, {testDenom2, two}}},
		{DecCoins{{testDenom1, zero}, {testDenom2, one}}, DecCoins{{testDenom1, zero}, {testDenom2, zero}}, DecCoins{{testDenom2, one}}},
		{DecCoins{{testDenom1, zero}, {testDenom2, zero}}, DecCoins{{testDenom1, zero}, {testDenom2, zero}}, DecCoins(nil)},
	}

	for tcIndex, tc := range cases {
		res := tc.inputOne.Add(tc.inputTwo...)
		require.Equal(t, tc.expected, res, "sum of coins is incorrect, tc #%d", tcIndex)
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		coin       DecCoin
		expectPass bool
		msg        string
	}{
		{
			NewDecCoin("mytoken", NewInt(10)),
			true,
			"valid coins should have passed",
		},
		{
			DecCoin{Denom: "BTC", Amount: NewDec(10)},
			false,
			"invalid denoms",
		},
		{
			DecCoin{Denom: "BTC", Amount: NewDec(-10)},
			false,
			"negative amount",
		},
	}

	for _, tc := range tests {
		tc := tc
		if tc.expectPass {
			require.True(t, tc.coin.IsValid(), tc.msg)
		} else {
			require.False(t, tc.coin.IsValid(), tc.msg)
		}
	}
}

func TestSubDecCoin(t *testing.T) {
	tests := []struct {
		coin       DecCoin
		expectPass bool
		msg        string
	}{
		{
			NewDecCoin("mytoken", NewInt(20)),
			true,
			"valid coins should have passed",
		},
		{
			NewDecCoin("othertoken", NewInt(20)),
			false,
			"denom mismatch",
		},
		{
			NewDecCoin("mytoken", NewInt(9)),
			false,
			"negative amount",
		},
	}

	decCoin := NewDecCoin("mytoken", NewInt(10))

	for _, tc := range tests {
		tc := tc
		if tc.expectPass {
			equal := tc.coin.Sub(decCoin)
			require.Equal(t, equal, decCoin, tc.msg)
		} else {
			require.Panics(t, func() { tc.coin.Sub(decCoin) }, tc.msg)
		}
	}
}

func TestSubDecCoins(t *testing.T) {
	tests := []struct {
		coins      DecCoins
		expectPass bool
		msg        string
	}{
		{
			NewDecCoinsFromCoins(NewCoin("mytoken", NewInt(10)), NewCoin("btc", NewInt(20)), NewCoin("eth", NewInt(30))),
			true,
			"sorted coins should have passed",
		},
		{
			DecCoins{NewDecCoin("mytoken", NewInt(10)), NewDecCoin("btc", NewInt(20)), NewDecCoin("eth", NewInt(30))},
			false,
			"unorted coins should panic",
		},
		{
			DecCoins{DecCoin{Denom: "BTC", Amount: NewDec(10)}, NewDecCoin("eth", NewInt(15)), NewDecCoin("mytoken", NewInt(5))},
			false,
			"invalid denoms",
		},
	}

	decCoins := NewDecCoinsFromCoins(NewCoin("btc", NewInt(10)), NewCoin("eth", NewInt(15)), NewCoin("mytoken", NewInt(5)))

	for _, tc := range tests {
		tc := tc
		if tc.expectPass {
			equal := tc.coins.Sub(decCoins)
			require.Equal(t, equal, decCoins, tc.msg)
		} else {
			require.Panics(t, func() { tc.coins.Sub(decCoins) }, tc.msg)
		}
	}
}

func TestSortDecCoins(t *testing.T) {
	good := DecCoins{
		NewInt64DecCoin("gas", 1),
		NewInt64DecCoin("mineral", 1),
		NewInt64DecCoin("tree", 1),
	}
	empty := DecCoins{
		NewInt64DecCoin("gold", 0),
	}
	badSort1 := DecCoins{
		NewInt64DecCoin("tree", 1),
		NewInt64DecCoin("gas", 1),
		NewInt64DecCoin("mineral", 1),
	}
	badSort2 := DecCoins{ // both are after the first one, but the second and third are in the wrong order
		NewInt64DecCoin("gas", 1),
		NewInt64DecCoin("tree", 1),
		NewInt64DecCoin("mineral", 1),
	}
	badAmt := DecCoins{
		NewInt64DecCoin("gas", 1),
		NewInt64DecCoin("tree", 0),
		NewInt64DecCoin("mineral", 1),
	}
	dup := DecCoins{
		NewInt64DecCoin("gas", 1),
		NewInt64DecCoin("gas", 1),
		NewInt64DecCoin("mineral", 1),
	}

	cases := []struct {
		coins         DecCoins
		before, after bool // valid before/after sort
	}{
		{good, true, true},
		{empty, false, false},
		{badSort1, false, true},
		{badSort2, false, true},
		{badAmt, false, false},
		{dup, false, false},
	}

	for tcIndex, tc := range cases {
		require.Equal(t, tc.before, tc.coins.IsValid(), "coin validity is incorrect before sorting, tc #%d", tcIndex)
		tc.coins.Sort()
		require.Equal(t, tc.after, tc.coins.IsValid(), "coin validity is incorrect after sorting, tc #%d", tcIndex)
	}
}

func TestDecCoinsIsValid(t *testing.T) {
	testCases := []struct {
		input    DecCoins
		expected bool
	}{
		{DecCoins{}, true},
		{DecCoins{DecCoin{testDenom1, NewDec(5)}}, true},
		{DecCoins{DecCoin{testDenom1, NewDec(5)}, DecCoin{testDenom2, NewDec(100000)}}, true},
		{DecCoins{DecCoin{testDenom1, NewDec(-5)}}, false},
		{DecCoins{DecCoin{"AAA", NewDec(5)}}, false},
		{DecCoins{DecCoin{testDenom1, NewDec(5)}, DecCoin{"B", NewDec(100000)}}, false},
		{DecCoins{DecCoin{testDenom1, NewDec(5)}, DecCoin{testDenom2, NewDec(-100000)}}, false},
		{DecCoins{DecCoin{testDenom1, NewDec(-5)}, DecCoin{testDenom2, NewDec(100000)}}, false},
		{DecCoins{DecCoin{"AAA", NewDec(5)}, DecCoin{testDenom2, NewDec(100000)}}, false},
	}

	for i, tc := range testCases {
		res := tc.input.IsValid()
		require.Equal(t, tc.expected, res, "unexpected result for test case #%d, input: %v", i, tc.input)
	}
}

func TestParseDecCoins(t *testing.T) {
	testCases := []struct {
		input          string
		expectedResult DecCoins
		expectedErr    bool
	}{
		{"", nil, false},
		{"4stake", NewDecCoinsFromDec("stake", NewDec(4)), false},
		{"5.5atom,4stake", NewDecCoins(NewDecCoinFromDec("atom", MustNewDecFromStr("5.5")), NewDecCoinFromDec("stake", NewDec(4))), false},
		{"0.0stake", nil, true},
		{"0.004STAKE", nil, true},
		{
			"0.004stake",
			DecCoins{NewDecCoinFromDec("stake", NewDecWithPrec(4000000000000000, Precision))},
			false,
		},
		{
			"5.04atom,0.004stake",
			DecCoins{
				NewDecCoinFromDec("atom", NewDecWithPrec(5040000000000000000, Precision)),
				NewDecCoinFromDec("stake", NewDecWithPrec(4000000000000000, Precision)),
			},
			false,
		},
	}

	for i, tc := range testCases {
		res, err := ParseDecCoins(tc.input)
		if tc.expectedErr {
			require.Error(t, err, "expected error for test case #%d, input: %v", i, tc.input)
		} else {
			require.NoError(t, err, "unexpected error for test case #%d, input: %v", i, tc.input)
			require.Equal(t, tc.expectedResult, res, "unexpected result for test case #%d, input: %v", i, tc.input)
		}
	}
}

func TestDecCoinsString(t *testing.T) {
	testCases := []struct {
		input    DecCoins
		expected string
	}{
		{DecCoins{}, ""},
		{
			DecCoins{
				NewDecCoinFromDec("atom", NewDecWithPrec(5040000000000000000, Precision)),
				NewDecCoinFromDec("stake", NewDecWithPrec(4000000000000000, Precision)),
			},
			"5.040000000000000000atom,0.004000000000000000stake",
		},
	}

	for i, tc := range testCases {
		out := tc.input.String()
		require.Equal(t, tc.expected, out, "unexpected result for test case #%d, input: %v", i, tc.input)
	}
}

func TestDecCoinsIntersect(t *testing.T) {
	testCases := []struct {
		input1         string
		input2         string
		expectedResult string
	}{
		{"", "", ""},
		{"1.0stake", "", ""},
		{"1.0stake", "1.0stake", "1.0stake"},
		{"", "1.0stake", ""},
		{"1.0stake", "", ""},
		{"2.0stake,1.0trope", "1.9stake", "1.9stake"},
		{"2.0stake,1.0trope", "2.1stake", "2.0stake"},
		{"2.0stake,1.0trope", "0.9trope", "0.9trope"},
		{"2.0stake,1.0trope", "1.9stake,0.9trope", "1.9stake,0.9trope"},
		{"2.0stake,1.0trope", "1.9stake,0.9trope,20.0other", "1.9stake,0.9trope"},
		{"2.0stake,1.0trope", "1.0other", ""},
	}

	for i, tc := range testCases {
		in1, err := ParseDecCoins(tc.input1)
		require.NoError(t, err, "unexpected parse error in %v", i)
		in2, err := ParseDecCoins(tc.input2)
		require.NoError(t, err, "unexpected parse error in %v", i)
		exr, err := ParseDecCoins(tc.expectedResult)
		require.NoError(t, err, "unexpected parse error in %v", i)

		require.True(t, in1.Intersect(in2).IsEqual(exr), "in1.cap(in2) != exr in %v", i)
	}
}

func TestDecCoinsTruncateDecimal(t *testing.T) {
	decCoinA := NewDecCoinFromDec("bar", MustNewDecFromStr("5.41"))
	decCoinB := NewDecCoinFromDec("foo", MustNewDecFromStr("6.00"))

	testCases := []struct {
		input          DecCoins
		truncatedCoins Coins
		changeCoins    DecCoins
	}{
		{DecCoins{}, Coins(nil), DecCoins(nil)},
		{
			DecCoins{decCoinA, decCoinB},
			Coins{NewInt64Coin(decCoinA.Denom, 5), NewInt64Coin(decCoinB.Denom, 6)},
			DecCoins{NewDecCoinFromDec(decCoinA.Denom, MustNewDecFromStr("0.41"))},
		},
		{
			DecCoins{decCoinB},
			Coins{NewInt64Coin(decCoinB.Denom, 6)},
			DecCoins(nil),
		},
	}

	for i, tc := range testCases {
		truncatedCoins, changeCoins := tc.input.TruncateDecimal()
		require.Equal(
			t, tc.truncatedCoins, truncatedCoins,
			"unexpected truncated coins; tc #%d, input: %s", i, tc.input,
		)
		require.Equal(
			t, tc.changeCoins, changeCoins,
			"unexpected change coins; tc #%d, input: %s", i, tc.input,
		)
	}
}

func TestDecCoinsQuoDecTruncate(t *testing.T) {
	x := MustNewDecFromStr("1.00")
	y := MustNewDecFromStr("10000000000000000000.00")

	testCases := []struct {
		coins  DecCoins
		input  Dec
		result DecCoins
		panics bool
	}{
		{DecCoins{}, ZeroDec(), DecCoins(nil), true},
		{DecCoins{NewDecCoinFromDec("foo", x)}, y, DecCoins(nil), false},
		{DecCoins{NewInt64DecCoin("foo", 5)}, NewDec(2), DecCoins{NewDecCoinFromDec("foo", MustNewDecFromStr("2.5"))}, false},
	}

	for i, tc := range testCases {
		tc := tc
		if tc.panics {
			require.Panics(t, func() { tc.coins.QuoDecTruncate(tc.input) })
		} else {
			res := tc.coins.QuoDecTruncate(tc.input)
			require.Equal(t, tc.result, res, "unexpected result; tc #%d, coins: %s, input: %s", i, tc.coins, tc.input)
		}
	}
}

func TestNewDecCoinsWithIsValid(t *testing.T) {
	fake1 := append(NewDecCoins(NewDecCoin("mytoken", NewInt(10))), DecCoin{Denom: "BTC", Amount: NewDec(10)})
	fake2 := append(NewDecCoins(NewDecCoin("mytoken", NewInt(10))), DecCoin{Denom: "BTC", Amount: NewDec(-10)})

	tests := []struct {
		coin       DecCoins
		expectPass bool
		msg        string
	}{
		{
			NewDecCoins(NewDecCoin("mytoken", NewInt(10))),
			true,
			"valid coins should have passed",
		},
		{
			fake1,
			false,
			"invalid denoms",
		},
		{
			fake2,
			false,
			"negative amount",
		},
	}

	for _, tc := range tests {
		tc := tc
		if tc.expectPass {
			require.True(t, tc.coin.IsValid(), tc.msg)
		} else {
			require.False(t, tc.coin.IsValid(), tc.msg)
		}
	}
}

func TestDecCoins_AddDecCoinWithIsValid(t *testing.T) {
	lengthTestDecCoins := NewDecCoins().Add(NewDecCoin("mytoken", NewInt(10))).Add(DecCoin{Denom: "BTC", Amount: NewDec(10)})
	require.Equal(t, 2, len(lengthTestDecCoins), "should be 2")

	tests := []struct {
		coin       DecCoins
		expectPass bool
		msg        string
	}{
		{
			NewDecCoins().Add(NewDecCoin("mytoken", NewInt(10))),
			true,
			"valid coins should have passed",
		},
		{
			NewDecCoins().Add(NewDecCoin("mytoken", NewInt(10))).Add(DecCoin{Denom: "BTC", Amount: NewDec(10)}),
			false,
			"invalid denoms",
		},
		{
			NewDecCoins().Add(NewDecCoin("mytoken", NewInt(10))).Add(DecCoin{Denom: "BTC", Amount: NewDec(-10)}),
			false,
			"negative amount",
		},
	}

	for _, tc := range tests {
		tc := tc
		if tc.expectPass {
			require.True(t, tc.coin.IsValid(), tc.msg)
		} else {
			require.False(t, tc.coin.IsValid(), tc.msg)
		}
	}
}

func TestDecCoinAmino(t *testing.T) {
	cdc := codec.New()
	amount := Dec{big.NewInt(1234567890123456)}
	amount = amount.Mul(Dec{big.NewInt(1234567890)})
	tests := []DecCoin{
		{},
		{"test", Dec{}},
		{"test", amount},
		{"test", Dec{big.NewInt(100000)}},
		{"", Dec{big.NewInt(math.MinInt64)}},
	}

	for i, test := range tests {
		expect, err := cdc.MarshalBinaryBare(test)
		require.NoError(t, err)
		actual, err := test.MarshalToAmino(cdc)
		require.NoError(t, err)
		require.EqualValues(t, expect, actual)

		t.Log(fmt.Sprintf("%d marshal pass", i))

		var dc DecCoin
		err = cdc.UnmarshalBinaryBare(expect, &dc)
		require.NoError(t, err)
		var actualDc DecCoin
		err = actualDc.UnmarshalFromAmino(cdc, expect)
		require.NoError(t, err)
		require.EqualValues(t, dc, actualDc)

		t.Log(fmt.Sprintf("%d unmarshal pass", i))
	}
}

func BenchmarkDecCoinAmino(b *testing.B) {
	cdc := codec.New()
	coin := DecCoin{"my coin", Dec{big.NewInt(1234567890123456)}}
	coin.Amount.Mul(Dec{big.NewInt(1234567890)})
	b.ResetTimer()

	b.Run("amino", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = cdc.MarshalBinaryBare(coin)
		}
	})

	b.Run("marshaller", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = coin.MarshalToAmino(cdc)
		}
	})
}
