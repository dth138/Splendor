package main

//RunSimulation Plays the Game
func RunSimulation(deck [][]GemCard, board [][]GemCard, bank Bank, players []Player) {

	//For Each Player
	pc := 0
	gameOver := false
	winCondition := false
	for !gameOver {
		//Take Turn
		PlayerTurn(players, &bank, deck, pc)
		//Evaluate Win condition
		if !winCondition {
			winCondition = EvalWinCondition(players, &bank, deck)
		}
		//Evaluate Game End (because game isnt' over at Win)
		gameOver = EvalGameEnd()
		//Update Turn
		pc = nextPlayer(pc, len(players))
		//Upkeep Board state

	}

}

func nextPlayer(i int, num int) int {
	i++
	return i % (num - 1)
}

//PlayerTurn is the control loop for a single turn
func PlayerTurn(players []Player, bank *Bank, deck [][]GemCard, pc int) {
	//Calculate Purchase power
	//Evaluate Moves
	//Move
}

//EvalWinCondition Evaluates whether or not to start the countdown timer
//  for end of the game
func EvalWinCondition(players []Player, bank *Bank, deck [][]GemCard) bool {
	return false
}

//EvalGameEnd determines when to stop running the simulation
func EvalGameEnd() bool {
	return false
}
