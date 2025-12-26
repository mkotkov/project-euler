const readline = require("readline");
const {PROBLEMS} = require("./registry");

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

console.log("Available Project Euler problems:");

rl.question("Number of problems Project Euler: ", (problemInput) => {
  const problemNumber = Number(problemInput);
  const solver = PROBLEMS[problemNumber];

  if (!solver) {
    console.log("Invalid problem number");
    rl.close();
    return;
  }

  rl.question("Enter number (n): ", (nInput) => {
    const n = Number(nInput);

    if (Number.isNaN(n)) {
      console.log("Input must be a number");
      rl.close();
      return;
    }

    const result = solver(n);
    console.log("Answer:", result);
    rl.close();
  });
});
