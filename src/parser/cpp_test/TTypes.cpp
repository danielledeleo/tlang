#include "TTypes.h"
#include <vector>
#include <string>
#include <iostream>
#include <sstream>

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
    VariableDeclaration::VariableDeclaration(Identifier* i, Type* t, size_t ln, size_t cn) {
        identifier = i;
        type = t;
        lineNo = ln;
        charNo = cn;
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

