# Concurrent Programming in Go
This repository holds codes of concurrent programming tools like mutexs, semaphores, wait groups, etc.

## Mutexes
Mutexes is short for *mutual exclusion*.
Mutexes make access to shared memory synchronous while modifying them. Thus blocking concurrency when modifying the shared resources and preventing race conditions. 

![Mutexes](images/mutexes.png)

## Conditional Variable
Condition variables work together with mutexes and give us the ability to block the current execution until we have a signal that a particular condition has changed.

![Conditional Variable](images/conditional_variable.png)