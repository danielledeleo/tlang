#include "TTypes.h"

namespace tlang {
    // IntegerType implementations
    string IntegerType::typeName() { return PrimitiveType::group() + ":" + "int"; }

    // RealType implementations
    string RealType::typeName() { return PrimitiveType::group() + ":" + "real"; }

    // BooleanType implementations
    string BooleanType::typeName() { return PrimitiveType::group() + ":" + "boolean"; }

    // NaturalType implemenations
    string NaturalType::typeName() { return PrimitiveType::group() + ":" + "nat"; }

    // IntegerNType implementations
    string IntegerNType::typeName() { return PrimitiveType::group() + ":" + "intn"; } 

    // NaturalNType implementations
    string NaturalNType::typeName() { return PrimitiveType::group() + ":" + "natn"; }

    // RealNType implementations
    string RealNType::typeName() { return PrimitiveType::group() + ":" + "realn"; }

    // CharType implementations
    string CharType::typeName() { return PrimitiveType::group() + ":" + "char"; }

    // ArrayType implementations
    string ArrayType::typeName() { 
        return CompositeType::group() + ":" + "array" + ":[" + memberType->typeName() + "]"; 
    }
    ArrayType::ArrayType(Type *t) { memberType = t; }

    // StringType implementations
    string StringType::typeName() {
        return CompositeType::group() + ":" + "string";
    }
    StringType::StringType() {}
    StringType::StringType(size_t _cap) { capacity = _cap; }

}