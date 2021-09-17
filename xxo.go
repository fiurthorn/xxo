package main

type Pos struct {
	X, Y int
}

func Minimax(player Player, move []Pos) int {
	return -1
	// if board.won || board.remaining == 0 {
	// 	return rating(player)
	// }

	// var bestScore = -1000;
	// for (var i = 0, len = Board.length; i < len; i++) {
	//   var xy = board.pos(i);
	//   if (board.isEmpty(xy)) {
	//     board[xy] = player;
	//     final score = -minimax(_oppoite(player), null);
	//     board.unset(xy);
	//     if (score > bestScore) {
	//       if (move != null) move.clear();
	//       bestScore = score;
	//     }
	//     if (bestScore == score && move != null) {
	//       move.add(xy);
	//     }
	//   }
	// }
	// return bestScore;
}
