/*
 *  Author: Joseph Harvey
 *  email: j3harvey@gmail.com
 */

package main

import (
  "bufio"
  "fmt"
  //"list"
  //"log"
  "os"
  "strconv"
  "strings"
  //"unicode/utf8"
)

// The name of this program is:
var NAME string = "goose"
// Version:
var VERSION string = "0.1"

var SIZE int = 19
var BOARD [][]rune = empty_board(SIZE)

func empty_board (size int) (board [][]rune) {
  // returns a slice of rune slices for an empty board
  if size > 25 || size < 2 {
    reporterror("", "unacceptable size")
  }
  board = make([][]rune, size, 25)
  for i := 0 ; i < size ; i++ {
    board[i] = make([]rune, size, 25)
    for j := 0 ; j < size ; j++ {
      board[i][j] = '0'
    }
  }
  return
}


/*
// default to a board size of 9
var board_size int = 9
// each point on the board has one of 4 states:
//    0: empty
//    1: black
//    2: white
//    3: illegal move due to ko
var board_configuration [9][9]byte
var w_captures, b_captures int
var komi float32
var time Time
var colour string // either 'w', 'b', 'white' or 'black'

type Time struct {
  t float
}


type Player struct {

}
*/


/*
 *  Utility commands for reading commands, and writing responses and errors
 */

func readline() (line string, err error) {
  r := bufio.NewReader(os.Stdin)
  line, err = r.ReadString('\n')
  if err != nil{
    fmt.Println(err)
    panic(err)
  }
  return
}

func tellmewhattodo() (id string, command string, args []string) {
  line, _ := readline()
  words := strings.Fields(line)
  _, err := strconv.Atoi(words[0])
  if err == nil {
    // The id of a gtp command is an integer, but not all commandsnhave an id.
    // Internally to this program a valid id is represented by a decimal string.
    // Commands received with no id are given an id that cannot be converted to 
    // an integer with Atoi, for internal use, the default being the empty string ""
    id, command, args = words[0], words[1], words[2:]
  } else {
    id, command, args = "", words[0], words[1:]
  }
  return
}

func sendresponse(id string, command string, args []string) {
  _, err := strconv.Atoi(id)
  if err == nil{
    fmt.Printf("=" + id + " " + command + " " + strings.Join(args, " ") + "\n\n")
  } else {
    fmt.Printf("= " + command + " " + strings.Join(args, " ") + "\n\n")
  }
}

func reporterror(id string, message string) {
  _, err := strconv.Atoi(id)
  if err == nil{
    fmt.Printf("?" + id + " " + message + "\n\n")
  } else {
    fmt.Printf("? " + message + "\n\n")
  }
}


/*
 *  Custom Commands (for testing and debugging)
 */

func empty_board_rune(i int, j int, size int) string {
  // returns a 2D array of unicode character which print prettily to a board
  switch {
  case size > 12 && (i == 4 || i == size - 3) && (j == 4 || j == size - 3):
    return "▪"
  case size > 12 && size % 2 == 1 && (i == 4 || i == size - 3 || i == (size + 1)/2) && (j == 4 || j == size - 3 || j == (size + 1)/2):
    return "▪"
  case i == 1 && j == 1:
    return "┌"
  case i == 1 && j == SIZE:
    return "┐"
  case i == 1:
    return "┬"
  case i == SIZE && j == 1:
    return "└"
  case i == SIZE && j == SIZE:
    return "┘"
  case i == SIZE:
    return "┴"
  case j == 1:
    return "├"
  case j == SIZE:
    return "┤"
  default:
    return "┼"
  }
}

func show_board() {
  //  Print a pretty board with unicode
  //
  //  Here are some useful runes to copy and paste:
  //   ┌┐└┘├┤┬┴┼═║╒╓╔╕╖╗╘╙╚╛╜╝╞╟╠╡╢╣╤╥╦╧╨╩╪╫╬▀▄▌▐■□▪▫◊○◌●◦
  //
  for i := 1 ; i <= SIZE ; i++ {
    for j := 1 ; j <= SIZE ; j++ {
      switch BOARD[i-1][j-1] {
      case 'b':
        fmt.Printf("○")
      case 'w':
        fmt.Printf("●")
      case '0':
        //if board position is empty, show the empty board symbol
        fmt.Printf(empty_board_rune(i, j, SIZE))
      case '*':
        //if board position is involved in a ko
        fmt.Printf("x")
      }
    }
    fmt.Printf("\n")
  }
  fmt.Printf("\n")
}

/*
 *  Go Text Protocol Commands
 *
 *  All supported commands:
 *    protocol_version    TODO
 *    name                TODO
 *    version             TODO
 *    known_command       TODO
 *    list_commands       TODO
 *    quit                TODO
 *    boardsize           TODO
 *    clear_board         TODO
 *    komi                TODO
 *    play                TODO
 *    genmove             TODO
 */

/*
func dispatch( id string, command string, args []string) {
  switch command {
    case "protocol_version":
      protocol_version(id, args)
    case "name":
      name(id)
    case "version":
      version(id)
    //case "known_command":
      //known_command(id, args)
    //case "list_commands":
      //list_commands(id, args)
    case "boardsize":
      boardsize(id, args)
    case "clear_board":
      clear_board(id, args)
    case "komi":
      komi(id, args)
    case "play":
      play(id, args)
    case "genmove":
      genmove(id, args)
    case "quit":
      os.Exit(0)
     default:
      reporterror(id, "unknown command")
  }
}
*/


func protocol_version(id string) {
  sendresponse(id, "2", make([]string, 0, 0))
}

func name(id string) {
  sendresponse(id, NAME, make([]string, 0, 0))
}

func version(id string) {
  sendresponse(id, VERSION, make([]string, 0, 0))
}

func boardsize(id string, size string) {
  _, err := strconv.Atoi(size)
  if err != nil{
    reporterror(id, "unacceptable size") //this error message is part of GTP 2
  }
  new_size, _ := strconv.Atoi(size)
  if new_size != SIZE {
    SIZE = new_size
    //BOARD = make([][]byte, 0, 0)
  }
  //sendresponse(id, SIZE, make([]string, 0, 0))
}

//func clear_board(id string) {
  //BOARD := [][]
//}

/*
 *  The main loop
 */

func main() {
  show_board()

  for {
    id, command, args := tellmewhattodo()
    sendresponse(id, command, args)
  }
}

