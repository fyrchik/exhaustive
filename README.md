# Preamble 
It happens, that there are some infinite sets on which we can perform exhaustive search in finite time.
For example, the Cantor set, which can be bijectively mapped to inifinite 0-1 sequences.

It is pretty easy to see, that if this predicate is `true` on _some_ inputs, then we can find them by brute-forcing.
However if our predicate is always false it's is not immediately obvious, _when_ should we return.

# WTF (I dont want math)
Every computable predicate over infinite 0-1 sequences must check only finite number of bits in it.

This number can be implicitly computed and it can be seen that our `Exists` terminates iff its argument terminates on every input.
 
# WTF (I want math)
There are some links below.
It is also true that for some function types equality predicate is decidable.

Haskell suits better (because it is lazy), however I've implemented it in Go, because why not?

# Links

http://math.andrej.com/2007/09/28/seemingly-impossible-functional-programs/

http://www.cs.bham.ac.uk/~mhe/papers/exhaustive.pdf

http://www.paultaylor.eu/ASD/
