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
	"no rank",
	"pair",
	"two pairs",
	"three of a kind",
	"straight",
	"flush",
	"full house",
	"four of a kind",
	"straight flush",
}

func (r Rank) String() string {
	return ranks[r]
}

func GetRank(hand []Card) Rank {
	m := map[Rank]func(hand []Card) bool{
		StraightFlush: IsStraightFlush,
		FourOfAKind:   IsRankFourOfAKind,
		FullHouse:     IsRankFullHouse,
		Flush:         IsFlush,
		Straight:      IsStraight,
		ThreeOfAKind:  IsRankThreeOfAKind,
		TwoPairs:      IsRankTwoPairs,
		Pair:          IsRankPair,
	}

	for r, isRank := range m {
		if isRank(hand) {
			return r
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

func IsRankPair(hand []Card) bool {
	return len(FindRepeats(2, hand)) == 2
}

func IsRankTwoPairs(hand []Card) bool {
	return len(FindRepeats(2, hand)) == 4
}

func IsRankThreeOfAKind(hand []Card) bool {
	return len(FindRepeats(3, hand)) == 3
}

func IsRankFourOfAKind(hand []Card) bool {
	return len(FindRepeats(4, hand)) == 4
}

func IsRankFullHouse(hand []Card) bool {
	return IsRankThreeOfAKind(hand) && IsRankPair(hand)
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
