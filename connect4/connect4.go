package main

// from http://blog.notdot.net/2009/11/Recursion-and-concurrency-with-Go
// currently async version is broken on a dual-core laptop the kernel kills the process before it completes

import (
	"fmt"
	)

type state struct {
	board [7][6]int;
	heights [7]int;
}

type move struct {
	column int;
	score int;
}

func (s state) String() (ret string) {
	for y := 0; y < 6; y++ {
		for x := 0; x < 7; x++ {
			switch s.board[x][y] {
			case -1:
				ret += "x";
			case 1:
				ret += "o";
			case 0:
				ret += ".";
			}
		}
		ret += "\n"
	}
	return;
}

func (s *state) move(column, player int) {
	s.board[column][s.heights[column]] = player;
	s.heights[column]--;
}

func evaluate_one(a, b, c, d int) int {
	if a == b && a == c && a == d && a != 0 {
		return a * 10000000;
	}
	if a == b && a == c && a != 0 && d == 0 {
		return a * 10000;
	}
	if a == 0 && b == c && b == d && b != 0 {
		return b * 10000;
	}
	if a == b && a != 0 && c == 0 && d == 0 {
		return a * 100;
	}
	if a == 0 && b == c && b != 0 && d == 0 {
		return b * 100;
	}
	if a == 0 && b == 0 && c == d && c != 0 {
		return c * 100;
	}
	if a != 0 && b == 0 && c == 0 && d == 0 {
		return a;
	}
	if a == 0 && b != 0 && c == 0 && d == 0 {
		return b;
	}
	if a == 0 && b == 0 && c != 0 && d == 0 {
		return c;
	}
	if a == 0 && b == 0 && c == 0 && d != 0 {
		return d;
	}
	return 0;
}

func (s *state) evaluate() int {
	score := 0;
	
	// Horizontal
	for x:=0; x < 4; x ++ {
		for y := 0; y < 6; y++ {
			score += evaluate_one(s.board[x][y], s.board[x+1][y], s.board[x+2][y], s.board[x+3][y]);
		}
	}
	
	// Vertical
	for x:=0; x<7; x++ {
		for y := 0; y < 3; y++ {
			score += evaluate_one(s.board[x][y], s.board[x][y+1], s.board[x][y+2], s.board[x][y+3]);
		}
	}
	
	// Positive diagonal
	for x := 0; x < 4; x++ {
		for y := 0; y < 3; y++ {
			score += evaluate_one(s.board[x][y], s.board[x+1][y+1], s.board[x+2][y+2], s.board[x+3][y+3]);
		}
	}
	
	// Negative diagonal
	for x := 3; x < 7; x++ {
		for y := 0; y < 3; y++ {
			score += evaluate_one(s.board[x][y], s.board[x-1][y+1], s.board[x-2][y+2], s.board[x-3][y+3]);
		}
	}
	
	return score;
}

func negamax_sync(s state, column, player, depth int) move {
	s.move(column, player)

	score := s.evaluate() * -player;
	if score < -1000000 || score > 1000000 || depth == 0 {
		return move{column, score};
	}

	min_score := 1000000;
	for i := 0; i < 7; i++ {
		if s.heights[i] >= 0 {
			m := negamax_sync(s, i, -1 * player, depth -1);
			if m.score < min_score {
				min_score = m.score;
			}
		}
	}

	return move{column, -min_score};
}

func negamax_async(s state, column, player, depth int, ret chan move) {
	s.move(column, player);

	score := s.evaluate() * -player;
	if score < -1000000 || score > 1000000 || depth == 0 {
		ret <- move{column, score};
		return;
	}

	results := make(chan move);
	num_results := 0;

	for i := 0; i < 7; i++ {
		if s.heights[i] >= 0 {
			go negamax_async(s, i, -1 * player, depth - 1, results);
			num_results++;
		}
	}

	min_score := 1000000;
	for i:= 0; i < num_results; i++ {
		m := <-results;
		if m.score < min_score {
			min_score = m.score;
		}
	}

	ret <- move{column, -min_score};
}

func sync_move(s *state, player int) int {
	max := move{-1, -100000000};
	for i := 0; i < 7; i++ {
		if s.heights[i] >= 0 {
			if m := negamax_sync(*s, i, player, 6); m.score > max.score {
				max = m;
			}
		}
	}
	return max.column;
}

func async_move(s * state, player int) int {
	ret := make(chan move);

	num_moves := 0;
	for i := 0; i < 7; i++ {
		if s.heights[i] >= 0 {
			go negamax_async(*s, i, player, 6, ret);
			num_moves += 1;
		}
	}

	max := move{-1, -100000000};
	for i := 0; i < num_moves; i++ {
		if m := <-ret; m.score > max.score || (m.score == max.score && m.column < max.column) {
			max = m;
		}
	}
	return max.column;
}

func main() {
	var s state;

	s.heights = [7]int{ 5, 5, 5, 5, 5, 5, 5 };
	player := 1;
	score := 0;
	for score > -1000000 && score < 1000000 {
		move := sync_move(&s, player);
		s.move(move, player);
		score = s.evaluate();
		player = -1 * player;

		fmt.Printf("%s", s);
		fmt.Printf("%d\n\n", score);
	}
}