function xyz (x, y, z: int) : int
    result x + y * z
end xyz

proc test
    put "test"
end test

var x : int := 12
x += 4

put xyz(1, 2, x)