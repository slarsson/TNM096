from task1 import Clause, resolution, solver

# 1. Nobody else could have been involved other than A, B and C.
#   A or B or C => a V b V c
#
# 2. C never commits a crime without A’s participation.
#   C -> A => -c V a
#
# 3. B does not know how to drive.
#   => da V dc + -db 

KB = set()
KB.add(Clause("a V b V c"))
KB.add(Clause("-c V a"))
KB.add(Clause("da V dc"))
KB.add(Clause("-db"))

res = solver(KB)
for item in res:
    print(item)
