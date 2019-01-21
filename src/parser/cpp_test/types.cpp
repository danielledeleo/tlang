#include "types.h"

// tlangIntegerType implementations
string tlangIntegerType::typeName() { return tlangPrimitiveType::group() + ":" + "int"; }

// tlangRealType implementations
string tlangRealType::typeName() { return tlangPrimitiveType::group() + ":" + "real"; }

// tlangArrayType implementations
string tlangArrayType::typeName() { 
    return tlangCompositeType::group() + ":" + "array" + ":" + memberType->typeName(); 
}
tlangArrayType::tlangArrayType(tlangType *t) { memberType = t; }