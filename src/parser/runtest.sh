#!/bin/bash
mkdir -p .antlr
antlr tlang.g4 -o .antlr

# macOS only
javac -classpath /usr/local/Cellar/antlr/4.7.2/antlr-4.7.2-complete.jar .antlr/*.java
cd .antlr
grun Turing program -gui ../$1
