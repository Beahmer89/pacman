package movement

import "fmt"
import "math/rand"
import "pacman/players"


func moveUp(newRow, newCol int, maze[]string) (int, int) {
    newRow = newRow - 1
    if newRow < 0 {
        newRow = len(maze) - 1
    }
    return newRow, newCol
}

func moveDown(newRow, newCol int, maze[]string) (int, int) {
    newRow = newRow + 1
    if newRow == len(maze) {
        newRow = 0
    }
    return newRow, newCol
}

func moveLeft(newRow, newCol int, maze[]string) (int, int) {
    newCol = newCol - 1
    if newCol < 0 {
        newCol = len(maze[0]) - 1
    }
    return newRow, newCol
}

func moveRight(newRow, newCol int, maze[]string) (int, int) {
    newCol = newCol + 1
    if newCol == len(maze[0]) {
        newCol = 0
    }
    return newRow, newCol
}

func makeMove(oldRow, oldCol int, direction string, maze[]string) (newRow, newCol int){
    var movementMap = map[string]func(row, col int, maze[]string) (int, int) {
        "UP": moveUp,
        "DOWN": moveDown,
        "LEFT": moveLeft,
        "RIGHT": moveRight,
    }

    newRow, newCol = movementMap[direction](oldRow, oldCol, maze)

    if maze[newRow][newCol] == '#' {
        newRow = oldRow
        newCol = oldCol
    }

    return newRow, newCol
}

func determineDirection() string {
    //Genereate random number to pick direction
    dir := rand.Intn(4)
    move := map[int]string {
        0: "UP",
        1: "DOWN",
        2: "RIGHT",
        3: "LEFT",
    }
    return move[dir]
}

func MovePlayer(direction string, maze[]string, pacman *players.Pacman) {
    pacman.Row, pacman.Col = makeMove(pacman.Row, pacman.Col, direction, maze)
}

func MoveGhosts(maze[]string, ghosts[]*players.Ghosts) {
    for _, ghost := range ghosts {
        direction := determineDirection()
        ghost.Row, ghost.Col = makeMove(ghost.Row, ghost.Col, direction, maze)
    }
}

func MoveCursor(row, col int) {
    fmt.Printf("\x1b[%d;%df", row+1, col+1)
}
