package main

import "strings"

const DEBUG = true

type Sprinkler rune

func (s Sprinkler) String() string {
	return string(s)
}

const (
	Working Sprinkler = '.'
	Damaged           = '#'
	Unknown           = '?'
)

type Group struct {
	Type       Sprinkler
	NumMembers int
	left       *Group
	right      *Group
}

func (g *Group) BorrowNeighbour(direction string, target int) (numBorrowed int, ok bool) {
	if !g.checkBorrow(direction, target) {
		return 0, false
	}
	switch direction {
	case "left":
		left := g.left
		// If left is damaged we have to take them all, or if it is just a single element we absorb it as well
		if left.Type == Damaged {
			g.left = left.left
			g.left.right = g
			g.NumMembers += left.NumMembers
			return left.NumMembers, true
		}
		if left.NumMembers == 1 && left.left.Type == Damaged {
			g.left = left.left
			n, borrowed := g.BorrowNeighbour(direction, target-1)
			if !borrowed {
				g.left = left
				return 0, false
			}
			g.NumMembers++
			return n + 1, true
		}
		left.NumMembers--
		g.NumMembers++
		return 1, true
	case "right":
		right := g.right
		if right.Type == Damaged || right.NumMembers == 1 {
			g.right = right.right
			g.right.left = g
			g.NumMembers += right.NumMembers
			return right.NumMembers, true
		}
		right.NumMembers--
		g.NumMembers++
		return 1, true
	}
	return 0, false
}

func (g *Group) checkBorrow(direction string, target int) bool {
	noLeft := g.left == nil
	if noLeft {
		return false
	}
	gWrongType := g.Type != Damaged
	leftWrongType := g.left.Type == Working
	overTarget := g.left.Type == Damaged && g.left.NumMembers+g.NumMembers > target
	return !(gWrongType || leftWrongType || overTarget)
}

func (g *Group) checkTerminated(direction string) bool {
	switch direction {
	case "left":
		if g.left.Type == Working {
			return true
		}
		if g.left.Type == Damaged {
			return false
		}
		if g.left.NumMembers == 1 {
			g.left = g.left.left
			g.left.right = g
			return true
		}
		g.left.NumMembers--
		return true
	case "right":
		if g.right.Type == Working {
			return true
		}
		if g.right.Type == Damaged {
			return false
		}
		if g.right.NumMembers == 1 {
			g.right = g.right.right
			g.right.left = g
			return true
		}
		g.right.NumMembers--
		return true
	}
	return false
}

func (g Group) String() string {
	return strings.Repeat(g.left.Type.String(), g.left.NumMembers) + "|" + strings.Repeat(g.Type.String(), g.NumMembers) + "|" + strings.Repeat(g.right.Type.String(), g.right.NumMembers)
}

type Sprinklers struct {
	Groups           *Group
	NumGroups        int
	ReferenceNumbers []int
}

// prefix and sufix checking with working collapse to a single working