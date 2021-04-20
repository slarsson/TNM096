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

# VARFÖR FUNKAR INTE DETTA!?!?!?!?!?!
# ??
def solver(kb):
    kb = copy.deepcopy(kb)

    while True:
        s = set()
        my_list = list(kb)
        for i in range(len(my_list)):
            for j in range(i+1, len(my_list)):
                c = resolution(my_list[i], my_list[j])
                if c is not False:
                    s.add(c)
        
        if len(s) == 0:
            return kb

        kb_prime = incorporate(s, kb)

        if kb == kb_prime:
            break
        kb = kb_prime
    
    print(len(kb_prime))
    for item in kb_prime:
        print(item)
    return kb

def incorporate(s, kb):
    kb = copy.deepcopy(kb)
    for a in s:
        kb = incorporate_clause(a, kb)
    return kb

def incorporate_clause(a, kb):
    kb = copy.deepcopy(kb)

    for b in kb:
        if b.is_subset(a):
            return kb

    
    rm = set()
    for b in copy.deepcopy(kb):
        if a.is_strict_subset(b):
            rm.add(b)

    new_kb = set()
    new_kb.add(a)
    for item in kb:
        if item not in rm:
            new_kb.add(item)
    return new_kb


if __name__ == "__main__":
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
   
    solver(KB)
