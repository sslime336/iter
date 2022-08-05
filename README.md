# iter

### About

Experimental implementation of iterators using Go for functional programming.

### Warn

Since all of these are experimental, this module should not be used in productive environment.

### Example

There are 2 ways to get the iterator.
1. Generate from existed slice, array or map.
2. Implement the Iterator interface. You can easily implement the interface by using the function Iter(any) from relevant package.


[1] Generate from existed slice, array or map.

```go


```


[2] Implement the Iterator interface.


```go


```

### Logger

A WeakLogger interface are provided for inner use, which means you can customize the log output using your favorite logger.

### Misc

Keep updating(maybe)...

The version supporting the production environment is under testing.
