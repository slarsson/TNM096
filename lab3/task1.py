import random
import copy

class Clause():
    def __init__(self, arg, empty=False):
        self.p = set()
        self.n = set()
        if not empty:
            self.__parse(arg)    

    def __hash__(self):
        return hash((frozenset(self.p), frozenset(self.n)))

    def __eq__(self, other):
        if not isinstance(other, type(self)):
            return NotImplemented
        return self.p == other.p and self.n == other.n

    def __parse(self, arg):
        arr = arg.split("V")
        assert len(arr) > 1
        for item in arr:
            c = item.strip()
            if len(c) == 0:
                continue
            if c[0] == "-":
                self.n.add(c[1:])
            else:
                self.p.add(c)

def resolution(a, b):
    a = copy.deepcopy(a)
    b = copy.deepcopy(b)

    x1 = a.p & b.n
    x2 = a.n & b.p
    if len(x1) == 0 and len(x2) == 0:
        return False
    if len(x1) != 0:
        item = random.sample(x1, 1)[0]
        a.p.remove(item)
        b.n.remove(item)
    else:
        item = random.sample(x2, 1)[0]
        a.n.remove(item)
        b.p.remove(item)

    c = Clause("", empty=True)
    c.p = a.p | b.p
    c.n = a.n | b.n
    if len(c.p & c.n) != 0:
        return False
    return c

# ???????
def solver(kb):
    xkb = copy.deepcopy(kb)
    s = set()

    #while xkb not kb:
    my_list = list(kb)
    for x in range(len(my_list)):
        for y in range(x + 1, len(my_list)):
            c = resolution(my_list[x], my_list[y])
            if c is not False:
                s.add(c)
            print(c)


q1 = Clause("a V b V -c")
q2 = Clause("c V b")
cc = resolution(q1, q2)
if cc is not False:
    print(cc.p)
    print(cc.n)

print(q1.p, q1.n)
print(q2.p, q2.n)

KB = set()
KB.add(q1)
KB.add(q2)
KB.add(q1)
assert len(KB) < 3

print(KB)

solver(KB)