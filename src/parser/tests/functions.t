function xyz (x, y, z: int) : int
    result x + y * z
end xyz

proc test
    put "test"
end test

var x : int := 12
x += 4

% comment

/* multiline comment
a couple more 
lines */

put xyz(1, 2, x)