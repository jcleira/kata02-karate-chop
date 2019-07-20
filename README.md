# codecata.com - kata02 - Karate Chop

![codecata.com](https://imgur.com/download/CJozxMr)

This repository contains my personal [kata02 Karate Chop](http://codekata.com/kata/kata02-karate-chop/) implementation.

For this kata I would like to include the though process involved while doing it (briefly).

#### entry 01

After reading the kata, I though oh ok! a binary search, but with five different implementations!.

Taking into account that I'll do it on Go I tried to think some possible implementations, I've got the following:

- Basic.
- Using concurrency (idiomatic one).
- Interface ready.
- Concurrent (using Mutexes) - not totally sure about this one. 
- Recursive - not something I would normally do (in Go).

#### entry 02

I would be creating a package for each implementation, and I would not include any `package main` (executable) as only I would need test to pass.

First step: create unit testing, I would be using testify (for simple asserts) and table tests that I think that would fit perfectly.

Second step: create benkchmarks, I would love to see how the different implementations relate on speed terms.

Third step: Create the `basic` implementation, I wouldn't use any special Go language's feature, just a plain binary search implementation.

#### entry 03

Creating the concurrency implementation was not as fun as I though on the beggining, it quickly became a complex solution, and I knew from the beggining that It wouldn't be performant.

The only thing that I enjoyed was checking for the goroutines and channels cleanup, as it's something that I always pay special attention (had some weird times with issues on that).

#### entry 04

The interfaces implementation, was a little funnier to develop. I wanted to extract the best interface that I could from Binary Search, ended creating one to update search's boundaries values.
