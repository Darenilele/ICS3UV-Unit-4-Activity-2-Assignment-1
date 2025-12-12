/**
 * @author Daren Ileleji
 * @version 1.0.0
 * @date 2025-12-10
 * @fileoverview This program will ask the user for a range of integers, then ask the user to input integers to calculate the sum in and out of that range
 */

package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {

    // Variables
    var startingString string = ""
    var startingInter int = 0

    var endingString string = ""
    var endingInter int = 0

    var numberEnteredString string = ""
    var numberEntered int = -1  // initialize to something not zero

    var insideSum int = 0
    var outsideSum int = 0

    // Input
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter an integer to create a starting range: ")
    startingString, _ = reader.ReadString('\n')
    startingString = strings.TrimSpace(startingString)
    startingInter, _ = strconv.Atoi(startingString)

    fmt.Print("Enter an integer to create an ending range: ")
    endingString, _ = reader.ReadString('\n')
    endingString = strings.TrimSpace(endingString)
    endingInter, _ = strconv.Atoi(endingString)

    // Processing - while-style loop
    for numberEntered != 0 {
      fmt.Print("Enter an integer in that range, if you want to quit enter 0: ")
      numberEnteredString, _ = reader.ReadString('\n')
      numberEnteredString = strings.TrimSpace(numberEnteredString)
      numberEntered, _ = strconv.Atoi(numberEnteredString)

        if numberEntered != 0 {
            if numberEntered >= startingInter && numberEntered <= endingInter {
            insideSum = insideSum + numberEntered
            } else {
            outsideSum = outsideSum + numberEntered
            }
        }
    }

    // Output
    fmt.Println("The sum of integers inside the range is", insideSum)
    fmt.Println("The sum of integers outside the range is", outsideSum)

    fmt.Println("\nDone.")
}
