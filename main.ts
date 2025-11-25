/**
 * @author Rabira Motuma
 * @version 1.0.0
 * @date 2025-11-22
 * @fileoverview This porgram asks the user for their name and age and then returns that information.
 */

// variables
let userName: string;
let ageAsString: string;
let ageAsNumber: number;
let ageFiveYearsAgo: number;

// input
userName = prompt("What is your name?") || "No name entered!";
ageAsString = prompt("How old are you?") || "-1";

// process
ageAsNumber = parseInt(ageAsString);
ageFiveYearsAgo = ageAsNumber - 5;

// output
console.log("\n")
console.log("Hello, " + userName + "!");
console.log("You are " + ageAsNumber + " years old.");
console.log(
  "Five years ago, you were " + ageFiveYearsAgo.toString() + " years old."
);

console.log("\nDone.");
