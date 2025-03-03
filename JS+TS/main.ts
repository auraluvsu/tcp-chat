while (true) {
    console.log("Welcome to the calculator");
    let x: string | null = prompt("Please input a number:\n");
    let y: string | null = prompt("\nNow input the next number:\n");
    let opp: string | null = prompt(`\nNow choose your input:
           1. Addition
           2. Subtraction
           3. Multiplication
           4. Division
           5. Exponentiation`
          );
    if (x === null || y === null || opp === null) {
        console.log("Please enter a valid input")
        break
    } else {
        switch (opp) {
            case '1':
                const sum = Number(x) + Number(y);
                console.log(`The sum is ${sum}`);
                break
            case '2':
                const difference = Number(x) - Number(y);
                console.log(`The difference is ${difference}`);
                break
            case '3':
                const product = Number(x) * Number(y);
                console.log(`The product is ${product}`);
                break
            case '4':
                const quotient = Number(x)/Number(y);
                console.log(`The quotient is ${quotient}`);
                break
            case '5':
                const exponented = Number(x) ** Number(y);
                console.log(`The exponented product is ${exponented}`)
                break
            default:
                console.log("Error! Please choose an operation:\n")
                continue
        }
    }
    const tryAgain: string | null = prompt("Try again? y/n\n")
    if (tryAgain === null) {
        break
    } else if (tryAgain === 'n') {

}
