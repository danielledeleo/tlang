antlr=`which antlr`
found=$?
if [ $found != 0 ];
	then antlr=`which antlr4`;
	found4=$?;
	if [ $found4 != 0 ]; 
		then echo "antlr not found in path"; 
		exit 1;
	fi
fi

$antlr -Dlanguage=Cpp -listener -visitor tlang.g4 -o ./cpp_test/parser
$antlr -Dlanguage=Go -listener -visitor tlang.g4 -o ./go_test/parser
