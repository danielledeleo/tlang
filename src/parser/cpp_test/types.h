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
};

class tlangPrimitiveType : public tlangType {
public:
    virtual string group() { return "primitive"; }
    tlangPrimitiveType(tlangPrimitive t) {
        primitive_type = t;
    }
private:
    tlangPrimitive primitive_type;
};

class tlangCompositeType : public tlangType {
public:
    virtual string group() { return "composite"; }
    tlangCompositeType(tlangComposite t) {
        composite_type = t;
    }
private:
    tlangComposite composite_type;
};


class tlangIdentifier {
public:
    // string          identifier;
    int             line_no, char_no;
    tlangType       *type;
};