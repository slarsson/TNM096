# Task 2 - NQueens

# a
Backtracking worst case grows exponentially, at ~250 it starts taking a very long time for my computer.

AC-3 can’t solve it at all. Without backtracking it has to start with the correct value, otherwise it might take the wrong path.

Min-conflict can probably be very large, since it should solve the problem in around 50 steps. But large input will result in many checks, slow..

# b
The BT+FC+MRV works best. Most important MRV seems to be, since it selects the next variable with the most constraints (most likely to fail path) and therefore finds failure earlier (reduces the search space).

# c
Around 50 steps, but it depends on the randomness. If the initial state is selected perfectly it could be done in a single step. Since randomness is involved you can’t find an exact solution.

# d
In state-based search the heuristics needs to be specific to the problem (such as the manhattan distance in lab 1). For the CSP the heuristics is more general, such as MRV, selecting the most constrained value. The heuristics could be used on several problems.
