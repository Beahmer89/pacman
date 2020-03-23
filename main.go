package main

import "bufio"
import "fmt"
import "log"
import "os"
import "os/exec"
import "pacman/movement"
import "pacman/players"

var maze []string
var pacman *players.Pacman
var ghosts []*players.Ghosts


func loadMaze() error {
    // read in txt file into mem
    f, err := os.Open("maze01.txt")
    if err != nil {
        return err
    }
    // close file reader
    defer f.Close()

    scanner := bufio.NewScanner(f)
    //return true as long as there is another line to read
    for scanner.Scan() {
        line := scanner.Text()
        maze = append(maze, line)
    }

    //figure out player position
    for row, line := range maze {
        for col, char := range line {
            switch char {
            case 'P':
                pacman = &players.Pacman{row, col}
            case 'G':
                ghosts = append(ghosts, &players.Ghosts{row, col})
            }
        }
    }

    return nil
}

func clearScreen() {
    fmt.Printf("\x1b[2J")
    movement.MoveCursor(0,0)
}

func printScreen() {
    clearScreen()
    for _, line := range maze {
        for _, chr := range line {
            switch chr {
            case '#':
                fmt.Printf("%c", chr)
            default:
                fmt.Printf(" ")
            }
        }
        fmt.Printf("\n")
    }
    movement.MoveCursor(pacman.Row, pacman.Col)
    fmt.Println("P")

    for _, ghost := range ghosts {
        movement.MoveCursor(ghost.Row, ghost.Col)
        fmt.Printf("G")
    }
}

func readInput() (string, error) {
    // make allocates and initializes objects
    buffer := make([]byte, 100)

    count, err := os.Stdin.Read(buffer)
    if err != nil {
        return "", err
    }

    // if user presses 1 key, is that key ESC
    if count == 1 && buffer[0] == 0x1b {
        return "ESC", nil
    } else if count >= 3 {
        if buffer[0] == 0x1b && buffer[1] == '[' {
            switch buffer[2]{
            case 'A':
                return "UP", nil
            case 'B':
                return "DOWN", nil
            case 'C':
                return "RIGHT", nil
            case 'D':
                return "LEFT", nil
            }
        }
    }
    return "", nil
}

//func movePlayer(direction string) {
//    pacman.Row, pacman.Col = movement.MakeMove(pacman.Row, pacman.Col, direction, maze)
//}

func init() {
    /** Enabling Cbreak Mode, which allows for some characters to be
    preprocessed and some to not be. Example Ctrl+C will abort,
    but arrow keys are passed
    **/
    cbTerm := exec.Command("/bin/stty", "cbreak", "-echo")
    cbTerm.Stdin = os.Stdin

    err := cbTerm.Run()
    if err != nil {
        log.Fatalf("Unable to activate cbreak mode terminal: %v\n", err)
    }
}

func cleanup() {
    /** Restoring Cooked Mode. Terminal we are used to, input preprocessed
    aka the system intercepts special characters to give them special meaning.
    **/
    cookedTerm := exec.Command("/bin/stty", "-cbreak", "echo")
    cookedTerm.Stdin = os.Stdin

    err := cookedTerm.Run()
    if err != nil {
        log.Fatalf("Unable to activate cooked mode terminal: %v\n", err)
    }
}

func main() {
    // initialize game
    defer cleanup()

    err := loadMaze()
    if err != nil {
        log.Printf("Error loading maze: %v\n", err)
        return
    }

    for {
        printScreen()
        input, err := readInput()
        if err != nil {
            log.Printf("Error reading input %v", err)
            break
        }

        movement.MovePlayer(input, maze, pacman)
        movement.MoveGhosts(maze, ghosts)

        if input == "ESC" {
            break
        }
    }
}
