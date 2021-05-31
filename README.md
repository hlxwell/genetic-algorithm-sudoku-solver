# genetic-algorithm-sudoku-solver

## How to run

```
go mod tidy
go run -v ./...
```

Result:

```
0th generation:
Matrix:  [R O D W W D O R R O W W W O R D] Valid Solution:  6
1th generation:
Matrix:  [R O D W W D O R R O W D W O R D] Valid Solution:  7
2th generation:
Matrix:  [R O D W W D O R R O W D D W R O] Valid Solution:  10
3th generation:
Matrix:  [R O D W W D O R R O W D D W R O] Valid Solution:  10
4th generation:
Matrix:  [R O D W W D O R R O W D D W R O] Valid Solution:  10
5th generation:
Matrix:  [R O D W W D O R R O W D D W R O] Valid Solution:  10
6th generation:
Result:
[R O D W]
[W D O R]
[O R W D]
[D W R O]
Total Valid Solutions:  12
```

## Places that you can tune.

1. CrossoverCount, which determine the crossover point.
2. Mutation method.
3. Crossover method.

## Pseudocode

```
TODO
```

## Findings

- After crossover several generations, all the individuals is going to be the same, so the mutation becomes to be important.

## References

- https://towardsdatascience.com/introduction-to-genetic-algorithms-including-example-code-e396e98d8bf3#:~:text=A%20genetic%20algorithm%20is%20a,offspring%20of%20the%20next%20generation.
- https://www.researchgate.net/publication/224180108_Solving_Sudoku_with_genetic_operations_that_preserve_building_blocks
- https://github.com/ctjacobs/sudoku-genetic-algorithm/blob/master/sudoku.py
