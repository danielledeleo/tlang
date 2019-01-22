# tlang
A new implementation of Turing built on LLVM with some help from ANTLR

# Building
tlang is in heavy development, so things are always changing.

You'll need:
- cmake
- ninja (for super fast builds)
- LLVM (including clang, g++ works for some parts for now)
- antlr4 (with C++ runtime libraries)
- Go for certain test modules

First, `cd tlang/src/parser`

Run `./configure.sh` (just runs cmake, for now)

Run `ninja` to build. 

Then try it out! `tlangcppparser tests/<somefile.t>`

# License
This is a reimplementation of a previously closed source programming language. I chose a new name and licensed it under GPLv3.
