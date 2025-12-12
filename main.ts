/**
 * @author Daren Ileleji
 * @version 1.0.0
 * @date 2025-12-10
 * @fileoverview This program will ask the user for a range of intgers, then ask the user to input intgers to calculate the average in that range
 */

// Variables 
let startingString: string = "";
let startingInter: number = 0;

let endingString: string = "";
let endingInter: number = 0;

let numberEnteredString: string = "";
let numberEntered: number = 0;

let insideSum: number = 0;
let outsideSum: number = 0;


// Input
startingString = prompt("Enter an interger to create a starting range") || "That is not a number";
startingInter = parseInt(startingString);
endingString = prompt("Enter an interger to create an ending range") || "That is not a number";
endingInter = parseInt(endingString);

// Processing 
do {
  numberEnteredString = prompt("Enter an interger in that range, if you wat to quit enter 0") || "That is not a number";
  numberEntered = parseInt(numberEnteredString);
  if (numberEntered != 0) {
    if (numberEntered >= startingInter && numberEntered <= endingInter) {
      insideSum = insideSum + numberEntered;
  } else {
    outsideSum = outsideSum + numberEntered;
  }

} 

} while (numberEntered != 0);

// Output
console.log(`The sum of intergers inside the range is ${insideSum}`)
console.log(`The sum of intergers outside the range is ${outsideSum}`)

console.log("\nDone.")



