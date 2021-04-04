# Task 1 - Sudoku

# a
For the Sudoku puzzle the number of values is equal to all squares (n = 9x9 = 81) and the domain size is 1 to 9, d = 9.

DFS is essentially bruteforce, we check all branches until we find a matching one, worst case complexity of O(d^n). If there are many missing numbers in the puzzle this might take a very long time and require a lot of memory.

Backtracking is DFS but we don’t visit branches that do not meet a constraint. DFS but with a reduced search space.

AC-3 uses arc consistency, it can in an early stage reduce the search space. Complexity = O(n^2 * d^3). This algorithm can only solve some of the easier puzzles, since it can't backtrack if there is someting wrong.

Min-conflicts is a bad choice. It tries with a random solution and then tries to change the conflicting values to find a new “state” with less conflicts. Could get stuck in a local-min.

# b
Forward Checking seems to be the most important option, since it can detect failure early it will reduce how deep the search has to go (use arc consistency to reduce the search space of the next variable). Minimum Remaining Values is better then First Unassigned Value since it will select the smallest domain (fewest remaining values), which hopefully is leading to a smaller search space.

Best choice: BT+FC+MRV
