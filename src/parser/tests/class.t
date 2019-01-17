class test
    export a, b

    function a (input : string) out : string
        out := input + input
    end a

    function b () : int
        result 4
    end b
end test