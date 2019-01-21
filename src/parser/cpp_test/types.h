#include <string>
using namespace std;

enum tlangPrimitive {
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

enum tlangComposite {
    tl_record,
    tl_array,
    tl_string,
    tl_class,
    tl_function
};

class tlangType {
public:
    virtual string group() = 0;
    virtual string typeName() = 0;
};

class tlangPrimitiveType : public tlangType {
public:
    virtual string group() { return "primitive"; }
    virtual string typeName() = 0;
};

class tlangCompositeType : public tlangType {
public:
    virtual string group() { return "composite"; }
    virtual string typeName() = 0;
};

class tlangIntegerType : public tlangPrimitiveType {
    virtual string typeName();
};

class tlangRealType : public tlangPrimitiveType {
    virtual string typeName();
};

class tlangArrayType : public tlangCompositeType {
public:
    tlangArrayType(tlangType *t);
    virtual string typeName();
    tlangType *memberType;
};

class tlangIdentifier {
public:
    // string          identifier;
    int             line_no, char_no;
    tlangType       *type;
};