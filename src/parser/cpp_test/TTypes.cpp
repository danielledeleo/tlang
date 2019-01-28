#include "TTypes.h"
#include <vector>
#include <string>
#include <iostream>
#include <sstream>
#include <cstdint>

#include "llvm/IR/LLVMContext.h"
#include "llvm/IR/IRBuilder.h"
#include "llvm/IR/Module.h"

namespace tlang {
    // IntegerType implementations
    string IntegerType::typeName() { return PrimitiveType::group() + ":" + "int"; }
    size_t IntegerType::sizeOf() { return sizeof(int); }
    

    // RealType implementations
    string RealType::typeName() { return PrimitiveType::group() + ":" + "real"; }
    size_t RealType::sizeOf() { return sizeof(float); }

    // BooleanType implementations
    string BooleanType::typeName() { return PrimitiveType::group() + ":" + "boolean"; }
    size_t BooleanType::sizeOf() { return sizeof(bool); }

    // NaturalType implemenations
    string NaturalType::typeName() { return PrimitiveType::group() + ":" + "nat"; }
    size_t NaturalType::sizeOf() { return sizeof(uint); }

    // Integer1Type implementations
    string Integer1Type::typeName() { return PrimitiveType::group() + ":" + "int1"; } 
    size_t Integer1Type::sizeOf() { return sizeof(int8_t); } 
    // Integer2Type implementations
    string Integer2Type::typeName() { return PrimitiveType::group() + ":" + "int2"; } 
    size_t Integer2Type::sizeOf() { return sizeof(int16_t); } 
    // Integer4Type implementations
    string Integer4Type::typeName() { return PrimitiveType::group() + ":" + "int4"; } 
    size_t Integer4Type::sizeOf() { return sizeof(int32_t); } 
    // Integer8Type implementations
    string Integer8Type::typeName() { return PrimitiveType::group() + ":" + "int8"; } 
    size_t Integer8Type::sizeOf() { return sizeof(int64_t); } 

    // Natural1Type implementations
    string Natural1Type::typeName() { return PrimitiveType::group() + ":" + "nat1"; }
    size_t Natural1Type::sizeOf() { return sizeof(uint8_t); }
    // Natural2Type implementations
    string Natural2Type::typeName() { return PrimitiveType::group() + ":" + "nat2"; }
    size_t Natural2Type::sizeOf() { return sizeof(uint16_t); }
    // Natural4Type implementations
    string Natural4Type::typeName() { return PrimitiveType::group() + ":" + "nat4"; }
    size_t Natural4Type::sizeOf() { return sizeof(uint32_t); }
    // Natural8Type implementations
    string Natural8Type::typeName() { return PrimitiveType::group() + ":" + "nat8"; }
    size_t Natural8Type::sizeOf() { return sizeof(uint64_t); }

    // Real4Type implementations
    string Real4Type::typeName() { return PrimitiveType::group() + ":" + "real4"; }
    size_t Real4Type::sizeOf() { return sizeof(float); }
    // Real8Type implementations
    string Real8Type::typeName() { return PrimitiveType::group() + ":" + "real8"; }
    size_t Real8Type::sizeOf() { return sizeof(double); }

    // CharType implementations
    string CharType::typeName() { return PrimitiveType::group() + ":" + "char"; }
    size_t CharType::sizeOf() { return sizeof(char); }

    // ArrayType implementations
    string ArrayType::typeName() { 
        return CompositeType::group() + ":" + "array" + ":[" + memberType->typeName() + "]"; 
    }
    ArrayType::ArrayType(Type *t) { memberType = t; }
    size_t ArrayType::sizeOf() {
        return sizeof(capacity) + sizeof(length) + sizeof(raw_array);
    }
    // StringType implementations
    string StringType::typeName() {
        return CompositeType::group() + ":" + "string";
    }
    StringType::StringType() {}
    StringType::StringType(size_t _cap) { capacity = _cap; }
    size_t StringType::sizeOf() {
        return sizeof(capacity) + sizeof(length) + sizeof(raw_array);
    }

    // RecordType implementations
    RecordType::RecordType(std::vector<Type*> m) {
        members = m;
    }
    string RecordType::typeName() {
        std::stringstream ss;

        for_each(members.begin(), members.end(), [&ss](Type *t) -> void {
            ss << std::endl << "    " << t->typeName();
        });

        return CompositeType::group() + ":" + "record" + ":[" + ss.str() + "\n" + "]";
    }

    size_t RecordType::sizeOf() {
        size_t sum = 0;

        for (auto n : members) {
            sum += n->sizeOf();
        }

        return sum;
    }

    // Identifier implementations
    Identifier::Identifier(string i) {
        identifier = i;
    }
    string Identifier::toString() {
        return identifier;
    }

    // Statement implementations
    Statement::Statement() {}

    // Declaration implementations

    // Expression implementations

    // Block implementations
    Block::Block() {};
    Block::Block(vector<Phrase*> p) {
        passage = p;
    }

    // VariableDeclaration implementations
    VariableDeclaration::VariableDeclaration(Identifier* i, Type* t, size_t ln, size_t cn, llvm::Value* val) {
        identifier = i;
        type = t;
        lineNo = ln;
        charNo = cn;
        value = val;
    }

    Identifier* VariableDeclaration::getIdentifier() {
        return identifier;
    }

    Type* VariableDeclaration::getType() {
        return type;
    }

    // TypeDeclaration implementations
    TypeDeclaration::TypeDeclaration(Identifier* i, Type* t) {
        identifier = i;
        type = t;
    }

    Identifier* TypeDeclaration::getIdentifier() {
        return identifier;
    }

    // FunctionDeclaration implementations
    FunctionDeclaration::FunctionDeclaration(Identifier* i, vector<Type*> a, Type* r, FunctionBlock* b) {
        identifier = i;
        arguments = a;
        returnType = r;
        body = b;
    }
    Identifier* FunctionDeclaration::getIdentifier() {
        return identifier;
    }

    // FunctionBlock implementations
    FunctionBlock::FunctionBlock(vector<Phrase*> p) {
        passage = p;
        if (!checkPhrases()) {
            throw new exception(); // todo
        }
    }
    bool FunctionBlock::checkPhrases() {
        return true; // todo
    }
    
    // ClassDeclaration implementations
    ClassDeclaration::ClassDeclaration(Identifier* i, ClassBlock* b) {
        identifier = i;
        body = b;
    }
    Identifier* ClassDeclaration::getIdentifier() {
        return identifier;
    }

    // ClassBlock implementations
    ClassBlock::ClassBlock(vector<Phrase*> p) {
        passage = p;
        if (!checkPhrases()) {
            throw new exception(); // todo
        }
    }

    bool ClassBlock::checkPhrases() {
        return true; // todo
    }

    // Module implementations
    Module::Module(string n, Block* b) {
        name = n;
        body = b;
    }
}

