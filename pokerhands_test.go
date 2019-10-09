package pokerhands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPockerXx(t *testing.T) {
	a := Card{Face: _2, Suit: D}
	b := Card{Face: K, Suit: D}

	result := b.isHigherThan(a)

	assert.True(t, result)
}

func TestGetHighestCard(t *testing.T) {
	expected := Card{Face: A, Suit: C}
	hand := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _4, Suit: C},
		Card{Face: _8, Suit: C},
		Card{Face: A, Suit: C},
		Card{Face: K, Suit: C},
	}

	actual := Highest(hand)

	assert.Equal(t, actual, expected)
}

func TestFindRepeats(t *testing.T) {
	expected := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _2, Suit: S},
	}
	hand := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _2, Suit: S},
		Card{Face: _8, Suit: C},
		Card{Face: _9, Suit: C},
		Card{Face: K, Suit: C},
	}

	actual := FindRepeats(2, hand)

	assert.Equal(t, actual, expected)
}

func TestIsPair(t *testing.T) {
	hand := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _2, Suit: S},
		Card{Face: _8, Suit: C},
		Card{Face: _9, Suit: C},
		Card{Face: K, Suit: C},
	}

	actual := IsPair(hand)

	assert.True(t, actual)
}

func TestIsTwoPairs(t *testing.T) {
	hand := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _2, Suit: S},
		Card{Face: _8, Suit: C},
		Card{Face: _8, Suit: S},
		Card{Face: K, Suit: C},
	}

	actual := IsTwoPairs(hand)

	assert.True(t, actual)
}

func TestIsThreeOfAKind(t *testing.T) {
	hand := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _2, Suit: S},
		Card{Face: _2, Suit: D},
		Card{Face: _8, Suit: C},
		Card{Face: K, Suit: C},
	}

	actual := IsThreeOfAKind(hand)

	assert.True(t, actual)
}

func TestIsFourOfAKind(t *testing.T) {
	hand := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _2, Suit: S},
		Card{Face: _2, Suit: D},
		Card{Face: _2, Suit: H},
		Card{Face: K, Suit: C},
	}

	actual := IsFourOfAKind(hand)

	assert.True(t, actual)
}

func TestIsFullHouse(t *testing.T) {
	hand := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _2, Suit: S},
		Card{Face: _2, Suit: D},
		Card{Face: K, Suit: H},
		Card{Face: K, Suit: C},
	}

	actual := IsFullHouse(hand)

	assert.True(t, actual)
}

func TestIsStraight(t *testing.T) {
	hand := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _3, Suit: S},
		Card{Face: _4, Suit: D},
		Card{Face: _5, Suit: H},
		Card{Face: _6, Suit: C},
	}

	actual := IsStraight(hand)

	assert.True(t, actual)
}

func TestIsFlush(t *testing.T) {
	hand := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _8, Suit: C},
		Card{Face: _4, Suit: C},
		Card{Face: K, Suit: C},
		Card{Face: _6, Suit: C},
	}

	actual := IsFlush(hand)

	assert.True(t, actual)
}

func TestStraightFlush(t *testing.T) {
	hand := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _3, Suit: C},
		Card{Face: _4, Suit: C},
		Card{Face: _5, Suit: C},
		Card{Face: _6, Suit: C},
	}

	actual := IsStraightFlush(hand)

	assert.True(t, actual)
}

func TestGroupingByFace(t *testing.T) {
	hand := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _2, Suit: S},
		Card{Face: _8, Suit: C},
		Card{Face: _8, Suit: D},
		Card{Face: K, Suit: C},
	}

	actual := GroupByFace(hand)

	expected := map[Face][]Card{
		_2: []Card{
			{Face: _2, Suit: C},
			{Face: _2, Suit: S},
		},
		_8: []Card{
			{Face: _8, Suit: C},
			{Face: _8, Suit: D},
		},
		K: []Card{
			{Face: K, Suit: C},
		},
	}

	assert.Equal(t, actual, expected)
}

func TestHavingFilter(t *testing.T) {
	group := map[Face][]Card{
		_2: []Card{
			{Face: _2, Suit: C},
			{Face: _2, Suit: S},
		},
		_8: []Card{
			{Face: _8, Suit: C},
			{Face: _8, Suit: D},
		},
		K: []Card{
			{Face: K, Suit: C},
		},
	}

	expected := map[Face][]Card{
		_2: []Card{
			{Face: _2, Suit: C},
			{Face: _2, Suit: S},
		},
		_8: []Card{
			{Face: _8, Suit: C},
			{Face: _8, Suit: D},
		},
	}

	actual := Having(repeats(2), group)

	assert.Equal(t, actual, expected)

}

func TestStraightFlushWinFourOfAkind(t *testing.T) {
	sf := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _3, Suit: C},
		Card{Face: _4, Suit: C},
		Card{Face: _5, Suit: C},
		Card{Face: _6, Suit: C},
	}

	fok := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _2, Suit: S},
		Card{Face: _2, Suit: D},
		Card{Face: _2, Suit: H},
		Card{Face: K, Suit: C},
	}

	actual := Duel(sf, fok)

	assert.Equal(t, actual, "straight flush > four of a kind")
}

func TestGetRankStraightFlush(t *testing.T) {
	h := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _3, Suit: C},
		Card{Face: _4, Suit: C},
		Card{Face: _5, Suit: C},
		Card{Face: _6, Suit: C},
	}

	actual := GetRank(h)

	assert.Equal(t, actual, StraightFlush)
}
func TestGetRankFourOfAKind(t *testing.T) {
	h := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _2, Suit: S},
		Card{Face: _2, Suit: D},
		Card{Face: _2, Suit: H},
		Card{Face: K, Suit: C},
	}

	actual := GetRank(h)

	assert.Equal(t, actual, FourOfAKind)
}
