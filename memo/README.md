# Memoization of Functions

*reference: The Go Programming Language - Alan A. A. Donavan, Brian W. Kernighan*

This tree contains two methods for enabling memoization. Additionally, these two methods are
representative of two methods of solving typical synchronization problems.

## monitor/

Solve memiozation problem by using a goroutine to act as monitor and guard
access to resource (in this case map of cached function calls).

## mutex/

Solve memoization problem by using a mutex to guard calls to shared resource
(again, in this case map of cached function calls).

