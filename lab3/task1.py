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
        assert len(arr) >= 1
        for item in arr:
            c = item.strip()
            if len(c) == 0:
                continue
            if c[0] == "-":
                self.n.add(c[1:])
            else:
                self.p.add(c)

    def __str__(self):
        values = []
        for value in self.p:
            values.append(value)
        for value in self.n:
            values.append("-" + value)
        return " V ".join(values)

    def __len__(self):
        return len(self.p) + len(self.n) 

    def is_subset(self, other):
        return self.p.issubset(other.p) and self.n.issubset(other.n)

    def is_strict_subset(self, other):
        sub = self.is_subset(other)
        l1 = len(self.p) + len(self.n)
        l2 = len(other.p) + len(other.n)
        return sub and l1 < l2

def resolution(a, b):
    a = copy.deepcopy(a)
    b = copy.deepcopy(b)

    x1 = a.p & b.n # intersection
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
    c.p = a.p | b.p # union
    c.n = a.n | b.n
    if len(c.p & c.n) != 0:
        return False
    return c

def solver(kb):
    while True:
        kb = copy.deepcopy(kb)
        kb_prim = copy.deepcopy(kb)
        s = set()

        my_list = list(kb)
        for i in range(len(kb)-1):
            for j in range(i+1, len(kb)):
                c = resolution(my_list[i], my_list[j])
                if c is not False:
                    s.add(c)
        
        # nothing to do
        if len(s) == 0:
            return kb

        # if something in s (a) is a strict subset to something in kb (b), meaning: b have everything in a + more
        # => remove b and a instead
        for a in s:
            item_to_remove = set()
            for b in kb:
                if a.is_strict_subset(b):
                    item_to_remove.add(b)
            for item in item_to_remove:
                kb.remove(item)
            kb.add(a)

        if kb == kb_prim:
            return kb


if __name__ == "__main__":
    print("TASK 1:")
    # 1. 
    # The resolution of two clauses in CNF. That is, given two clauses the 
    # program calculates their resolvent by applying one resolution step.
    A = Clause("a V b V -c")
    B = Clause("c V b")
    res = resolution(A, B)
    print(res)
    assert res == Clause("a V b")

    A = Clause("a V b V -c")
    B = Clause("d V b V -g")
    res = resolution(A, B)
    print(res)
    assert res == False

    A = Clause("-b V c V t")
    B = Clause("-c V z V b")
    res = resolution(A, B)
    print(res)
    assert res == False

    print("")
    print("TASK 2:")
    # 2.
    # The resolution mechanism applied to a given set S of clauses.
    # Given S, the program selectstwo arbitrary clauses from S, or any previously calculated resolvent, and calculates thenew resolvents.
    # The program applies the resolution step until no new resolvent can bederived.
    KB = set()
    KB.add(Clause("-sun V -money V ice"))
    KB.add(Clause("-money V ice V movie"))
    KB.add(Clause("-movie V money"))
    KB.add(Clause("-movie V -ice"))
    KB.add(Clause("movie"))
    res = solver(KB)
    for item in res:
        print(item)
    assert len(res) == 4
    items = [item.__str__() for item in res]
    assert "money" in items
    assert "movie" in items
    assert "-ice" in items
    assert "-sun" in items
