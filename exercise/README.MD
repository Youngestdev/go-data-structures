# Exercises

## 1. Write a sequential search that looks a key passed as an argument in an array also passed as an argument.

#### solution: [Sequential Search Solution](./SeqSearch.go)

## 2. Understand the a maximum character sequence function found in [MaxCharSequence](./MaxCharSequence.go)

## 3. Rewrite the bubble sort algorithm to incorporate a check to see whether the array in sorted after each pass. and to stop processing when this occurs.

#### solution: [Bubble Sort](./BubbleSort.go)

## 4. Write a complete Go program to solve the towers of Hanoi problem as outlined in the text. Have the `MoveDisk()` method write the state of the game after it moves a disk, so that when towers are moved the sequence of steps can be traced. Number the disks from 1 for the smallest to `n` for the largest. The `HanoiState` struct will need private fields to keep track of the contents of the tower A, B and C (what would be a good data structure for this task?), and a `String()` method to create a `string` representation of the three towers. The main program can create an instance of `HanoiState` with four disks and then call `MoveTower` ( Peg.A, Peg.C, Peg.B, 4) to move them.
#### solution: [Tower of Hanoi](./HanoiTower.go)