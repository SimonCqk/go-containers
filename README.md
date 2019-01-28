# go-containers
extension of standard "container" package, include a variety of commonly used data structures.

### queue

fast queue implementation based on ring buffer, and auto-scale, auto-shrink
are supported. Not thread-safe.

### stack

fast stack implementation based on slice, also support for auto-scale and 
auto-shrink.

### set

simple set implementation based on `map`.

### heap

a concrete encapsulation of "container/heap" interface, this is a minimum heap
implementation.

### priority queue

priority queue implementation, meanwhile it's a maximum heap, ordered by its priority
value.