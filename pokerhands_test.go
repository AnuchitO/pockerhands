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

func TestIsRankPair(t *testing.T) {
	hand := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _2, Suit: S},
		Card{Face: _8, Suit: C},
		Card{Face: _9, Suit: C},
		Card{Face: K, Suit: C},
	}

	actual := IsRankPair(hand)

	assert.True(t, actual)
}

func TestIsRankTwoPairs(t *testing.T) {
	hand := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _2, Suit: S},
		Card{Face: _8, Suit: C},
		Card{Face: _8, Suit: S},
		Card{Face: K, Suit: C},
	}

	actual := IsRankTwoPairs(hand)

	assert.True(t, actual)
}

func TestIsRankThreeOfAKind(t *testing.T) {
	hand := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _2, Suit: S},
		Card{Face: _2, Suit: D},
		Card{Face: _8, Suit: C},
		Card{Face: K, Suit: C},
	}

	actual := IsRankThreeOfAKind(hand)

	assert.True(t, actual)
}

func TestIsRankFourOfAKind(t *testing.T) {
	hand := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _2, Suit: S},
		Card{Face: _2, Suit: D},
		Card{Face: _2, Suit: H},
		Card{Face: K, Suit: C},
	}

	actual := IsRankFourOfAKind(hand)

	assert.True(t, actual)
}

func TestIsRankFullHouse(t *testing.T) {
	hand := []Card{
		Card{Face: _2, Suit: C},
		Card{Face: _2, Suit: S},
		Card{Face: _2, Suit: D},
		Card{Face: K, Suit: H},
		Card{Face: K, Suit: C},
	}

	actual := IsRankFullHouse(hand)

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
