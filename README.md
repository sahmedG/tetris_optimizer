Tetris Optimizer

![image](https://ars.els-cdn.com/content/image/1-s2.0-S0926580521000522-gr2.jpg)

Overview

Tetris Optimizer is a Go program designed to compile successfully, assemble Tetris tetrominoes in order to create the smallest square possible, and identify each tetromino in the solution using uppercase Latin letters (A for the first one, B for the second, etc). The program expects at least one tetromino in the input text file. In case of bad format on the tetrominoes or a bad file format, it prints ERROR.

Features
Compilation: The program compiles successfully and is written in Go.
Optimization: Assembles tetrominoes to create the smallest square possible.
Identification: Each tetromino in the solution is identified with uppercase Latin letters (A, B, etc).
Error Handling: In case of bad format on the tetrominoes or a bad file format, the program prints ERROR.
Test Files: Unit testing is encouraged, and sample test files are provided.
Usage
bash
Copy code
go run main.go input.txt
Replace input.txt with the path to your tetromino input file.

Example
Input file (input.txt):

bash

  ....
  .##.
  .##.
  ....

Run the program:

bash

go run main.go input.txt
Output:

A

Contributing

Contributions are welcome! Please follow the contribution guidelines when submitting pull requests.

License
This project is licensed under the MIT License - see the LICENSE file for details.

Acknowledgments

Inspired by the classic game Tetris.
Special thanks to the Go community for their valuable contributions.
Contact
For issues or suggestions, please open an issue.

Happy optimizing! 🎮
