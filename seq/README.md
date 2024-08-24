seq - A Go Iterator Module
==========================

[![Go Reference](https://pkg.go.dev/badge/github.com/GodsBoss/g/seq.svg)](https://pkg.go.dev/github.com/GodsBoss/g/seq) [![Go Report Card](https://goreportcard.com/badge/github.com/GodsBoss/g/seq)](https://goreportcard.com/report/github.com/GodsBoss/g/seq)

A module built around [range funcs introduced in Go 1.23](https://go.dev/doc/go1.23#language). See [documentation](https://pkg.go.dev/github.com/GodsBoss/g/seq) for usage and examples.

Overview
--------

Most functionality in this package can be roughly grouped into three categories:

- __Producers__: Construct an iterator from non-iterator values. Example from the standard library: [`slices.All`](https://pkg.go.dev/slices#All).
- __Consumers__: Take iterators and condense them into non-iterator values. Example from the standard library: [`slices.Collect`](https://pkg.go.dev/slices#Collect).
- __Transformers__: Provide iterators backed by other iterators.

Packages of this module may contain any combination of producers, consumers and transformers.

Iterator properties
-------------------

### Finite vs. Infinite

Some iterators stop yielding items even if the `for` loop that ranges over them never `break`s or `return`s. Those are called __finite__ iterators. If an iterator produces values as long as `for` ranges over it, it is called __infinite__.

Infinite iterators can lead to problems in some situations, e.g. when trying to collect all values from the iterator into a slice.

### Reusability

Iterators may yield the same values when `range`d over twice. Such iterators are called __reusable__. Others may yield different values or no values at all.

### In `seq` Module

There is no in-code flag or label to mark iterators as finite/infinite or whether it is reusable. Many iterators are both, e.g. the [`Filter`](https://pkg.go.dev/github.com/GodsBoss/g/seq@v1.0.1/iterate#Filter) transformer produces iterators that inherit their properties from the iterator the values are taken from – if its finite, the filtered iterator is finite and so on. In general, the behaviour is documented or can be deduced fairly easily.
