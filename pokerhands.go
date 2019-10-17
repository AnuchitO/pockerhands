package pokerhands

import (
	"fmt"
)

type Face int

const (
	_2 Face = iota
	_3
	_4
	_5
	_6
	_7
	_8
	_9
	T
	J
	Q
	K
	A
)

type Suit int

const (
	S Suit = iota
	C
	H
	D
)

type Card struct {
	Face Face
	Suit Suit
}

func (c Card) isHigherThan(b Card) bool {
	return c.Face > b.Face
}

func Highest(hand []Card) Card {
	c, tail := hand[0], hand[1:]
	for _, h := range tail {
		if h.isHigherThan(c) {
			c = h
		}
	}

	return c
}

type Rank int

const (
	_ Rank = iota
	Pair
	TwoPairs
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

var ranks = [...]string{
	Pair:          "pair",
	TwoPairs:      "two pairs",
	ThreeOfAKind:  "three of a kind",
	Straight:      "straight",
	Flush:         "flush",
	FullHouse:     "full house",
	FourOfAKind:   "four of a kind",
	StraightFlush: "straight flush",
}

func (r Rank) String() string {
	return ranks[r]
}

func GetRank(hand []Card) Rank {
	tuples := []struct {
		r  Rank
		is func(hand []Card) bool
	}{
		{StraightFlush, IsStraightFlush},
		{FourOfAKind, IsFourOfAKind},
		{FullHouse, IsFullHouse},
		{Flush, IsFlush},
		{Straight, IsStraight},
		{ThreeOfAKind, IsThreeOfAKind},
		{TwoPairs, IsTwoPairs},
		{Pair, IsPair},
	}

	for _, t := range tuples {
		if t.is(hand) {
			return t.r
		}
	}

	return 0
}

func Duel(a, b []Card) string {
	rankB := GetRank(b)

	rankA := GetRank(a)

	if rankA > rankB {
		return fmt.Sprintf("%s > %s", rankA, rankB)
	}

	return fmt.Sprintf("a:%s,  b:%s", rankA, rankB)
}

func IsStraightFlush(hand []Card) bool {
	return IsStraight(hand) && IsFlush(hand)
}

func IsStraight(hand []Card) bool {
	c, tail := hand[0], hand[1:]
	for _, h := range tail {
		if h.Face != c.Face+1 {
			return false
		}
		c = h
	}
	return true
}

func IsFlush(hand []Card) bool {
	c, tail := hand[0], hand[1:]
	for _, h := range tail {
		if h.Suit != c.Suit {
			return false
		}
	}

	return true
}

func IsPair(hand []Card) bool {
	return len(FindRepeats(2, hand)) == 2
}

func IsTwoPairs(hand []Card) bool {
	return len(FindRepeats(2, hand)) == 4
}

func IsThreeOfAKind(hand []Card) bool {
	return len(FindRepeats(3, hand)) == 3
}

func IsFourOfAKind(hand []Card) bool {
	return len(FindRepeats(4, hand)) == 4
}

func IsFullHouse(hand []Card) bool {
	return IsThreeOfAKind(hand) && IsPair(hand)
}

func FindRepeats(r int, hand []Card) []Card {
	return Values(Having(repeats(r), GroupByFace(hand)))
}

func GroupByFace(hand []Card) map[Face][]Card {
	m := map[Face][]Card{}
	for _, h := range hand {
		v := m[h.Face]
		c := append(v, h)
		m[h.Face] = c
	}
	return m
}

func Having(predicate func(card []Card) bool, group map[Face][]Card) map[Face][]Card {
	result := map[Face][]Card{}
	for k, v := range group {
		if predicate(v) {
			result[k] = v
		}
	}

	return result
}

func Values(group map[Face][]Card) []Card {
	result := []Card{}
	for _, v := range group {
		result = append(result, v...)
	}

	return result
}

func repeats(size int) func(card []Card) bool {
	return func(card []Card) bool {
		return len(card) == size
	}
}
