package main

import (
    "fmt"
    "os"
    "bufio"
    "time"
)

func printTime() {
        fmt.Print("\033[s") // save position
        fmt.Print("\033[0;0H") // goto 0 0
        fmt.Print("Current time: ", time.Now().Format("15:04:05"))
        fmt.Print("\033[u") // restore position
}


func main() {
    fmt.Print("\033[2 q") // set block style cursor
    fmt.Print("\033[2J\033[2;0H") // clear & goto 0 0

    file, err := os.Open("/dev/input/mice")
    defer file.Close()

    if err != nil {
        panic(err)
    }

    // clock
    go func() {
        for {
            printTime()

            time.Sleep(time.Second)
        }
    }()

    rd := bufio.NewReader(file)

    buffer := make([]byte, 3)
    for {
        rd.Read(buffer)

        /*
        var mouse string

        switch buffer[0] {
        case 9, 25, 41, 57:
            mouse = "Left"

        case 10, 26, 42, 58:
            mouse = "Right"

        case 11, 27, 43, 59:
            mouse = "Left + right"

        case 12, 28, 44, 60:
            mouse = "Middle"

        case 13, 29, 45, 61:
            mouse = "Middle + left"

        case 14, 30, 46, 62:
            mouse = "Middle + right"
          
        case 15, 31, 47, 63:
            mouse = "All"

        default:
            mouse="Rest"
        }
        */

        x, y := int8(buffer[1]), int8(buffer[2])

        if x > 0 {
            fmt.Printf("\033[%dC", x)

        } else if x < 0 {
            fmt.Printf("\033[%dD", -x)
        }

        if y > 0 {
            fmt.Printf("\033[%dA", y)
        } else if y < 0 {
            fmt.Printf("\033[%dB", -y)
        }

        switch buffer[0] {
        case 10, 26, 42, 58: // right
            fmt.Print("##\033[2D")

        case 9, 25, 41, 57: // left
            fmt.Print("  \033[2D")

        case 12, 28, 44, 60: // middle
            fmt.Print("  \033[2D")

        case 11, 27, 43, 59, 14, 30, 46, 62:  // left + right & middle + right
            fmt.Print("\033[2J")
            printTime()
        }

        time.Sleep(10*time.Millisecond)
    }
}
