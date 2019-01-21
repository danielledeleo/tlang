#include <string>
using namespace std;

namespace tlang {
    enum Primitive {
        tl_void,
        tl_int,
        tl_real,
        tl_boolean,
        tl_nat,
        tl_intn,
        tl_natn,
        tl_realn,
        tl_char
    };

    enum Composite {
        tl_record,
        tl_array,
        tl_string,
        tl_class,
        tl_function
    };

    class Type {
    public:
        virtual string group() = 0;
        virtual string typeName() = 0;
    };

    class PrimitiveType : public Type {
    public:
        virtual string group() { return "primitive"; }
        virtual string typeName() = 0;
    };

    class CompositeType : public Type {
    public:
        virtual string group() { return "composite"; }
        virtual string typeName() = 0;
    };

    // primitive types
    class IntegerType : public PrimitiveType {
        virtual string typeName();
    };

    class RealType : public PrimitiveType {
        virtual string typeName();
    };

    class BooleanType : public PrimitiveType {
        virtual string typeName();
    };

    class NaturalType : public PrimitiveType {
        virtual string typeName();
    };

    class IntegerNType : public PrimitiveType {
        virtual string typeName();
    };

    class NaturalNType : public PrimitiveType {
        virtual string typeName();
    };

    class RealNType : public PrimitiveType {
        virtual string typeName();
    };

    class CharType : public PrimitiveType {
        virtual string typeName();
    };

    // composite types
    class ArrayType : public CompositeType {
    public:
        ArrayType(Type *t);
        virtual string typeName();
        Type *memberType;
    private:
        size_t capacity;
        size_t length;
        size_t offset; // offset of the zeroth index (e.g. array 1950 .. 2020 of int), offset would be 1950
    };

    class StringType : public CompositeType {
    public:
        StringType();
        StringType(size_t _cap);
        virtual string typeName();
    private:
        size_t capacity;
        size_t length;
    };

    class Identifier {
    public:
        // string          identifier;
        int             line_no, char_no;
        Type       *type;
    };
} // end namespace 