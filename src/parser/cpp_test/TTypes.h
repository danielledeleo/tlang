#include "llvm/IR/LLVMContext.h"
#include "llvm/IR/IRBuilder.h"
#include "llvm/IR/Module.h"

#include <string>
#include <vector>
using namespace std;

namespace tlang {
    class Type {
    public:
        virtual string group() = 0;
        virtual string typeName() = 0;
    };

    class PrimitiveType : public Type {
    public:
        virtual string group() { return "primitive"; }
        virtual string typeName() = 0;
        virtual size_t sizeOf() = 0;
    };

    class CompositeType : public Type {
    public:
        virtual string group() { return "composite"; }
        virtual string typeName() = 0;
    };

    // primitive subtype classes
    class IntegerType : public PrimitiveType {
        virtual string typeName();
        virtual size_t sizeOf();
    };

    class RealType : public PrimitiveType {
        virtual string typeName();
        virtual size_t sizeOf();
    };

    class BooleanType : public PrimitiveType {
        virtual string typeName();
        virtual size_t sizeOf();
    };

    class NaturalType : public PrimitiveType {
        virtual string typeName();
        virtual size_t sizeOf();
    };

    class Integer1Type : public PrimitiveType {
        virtual string typeName();
        virtual size_t sizeOf();
    };
    class Integer2Type : public PrimitiveType {
        virtual string typeName();
        virtual size_t sizeOf();
    };
    class Integer4Type : public PrimitiveType {
        virtual string typeName();
        virtual size_t sizeOf();
    };
    class Integer8Type : public PrimitiveType {
        virtual string typeName();
        virtual size_t sizeOf();
    };

    class Natural1Type : public PrimitiveType {
        virtual string typeName();
        virtual size_t sizeOf();
    };
    class Natural2Type : public PrimitiveType {
        virtual string typeName();
        virtual size_t sizeOf();
    };
    class Natural4Type : public PrimitiveType {
        virtual string typeName();
        virtual size_t sizeOf();
    };
    class Natural8Type : public PrimitiveType {
        virtual string typeName();
        virtual size_t sizeOf();
    };

    class Real4Type : public PrimitiveType {
        virtual string typeName();
        virtual size_t sizeOf();
    };
    class Real8Type : public PrimitiveType {
        virtual string typeName();
        virtual size_t sizeOf();
    };

    class CharType : public PrimitiveType {
        virtual string typeName();
        virtual size_t sizeOf();
    };

    // composite subtype classes
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

    class RecordType : public CompositeType {
    public:
        RecordType(vector<tlang::Type*> m);
        virtual string typeName();
    private:
        vector<tlang::Type*> members;
    };

    class Identifier {
    public:
        Identifier(string i);
        string toString();
    private:
        string identifier;
    };

    class Expression {
    };

    // todo: pick a better name for this, Turing uses "statementsAndDeclarations"
    class Phrase {
    public:
        size_t lineNo;
        size_t charNo;
    };

    class Statement : public Phrase {
    public:
        Statement();
    private:
        Expression* expression; // maybe LHS/RHS?
    };

    class Declaration : public Phrase {
    public:
        virtual Identifier* getIdentifier() = 0;
        virtual llvm::Value* codeGen() = 0;
    };

    // Block is a general container for groups of statements, expressions, and declarations. 
    class Block {
    public:
        Block();
        Block(vector<Phrase*> p);
        vector<Phrase*> passage;
    private:
    };

    class VariableDeclaration : public Declaration {
    public:
        VariableDeclaration(Identifier* i, Type* t, size_t ln, size_t cn);
        Identifier* getIdentifier();
        Type* getType();
        llvm::Value* codeGen();
    private:
        Identifier* identifier;
        Type*       type;
    };

    class TypeDeclaration : public Declaration {
    public:
        TypeDeclaration(Identifier* i, Type* t);
        Identifier* getIdentifier();
        llvm::Value* codeGen();
    private:
        Identifier* identifier;
        Type* type;
    };

    class FunctionBlock : public Block {
    public:
        FunctionBlock(vector<Phrase*> p);
    private:
        bool checkPhrases();
    };

    class FunctionDeclaration : public Declaration {
    public:
        FunctionDeclaration(Identifier* i, vector<Type*> a, Type* r, FunctionBlock* b);
        Identifier* getIdentifier();
        llvm::Value* codeGen();
    private:
        Identifier* identifier;
        vector<Type*> arguments;
        Type* returnType;
        FunctionBlock* body;
    };

    class ClassBlock : public Block {
    public:
        ClassBlock(vector<Phrase*> p);
    private:
        bool checkPhrases();
    };

    class ClassDeclaration : public Declaration {
    public:
        ClassDeclaration(Identifier* i, ClassBlock* b);
        Identifier* getIdentifier();
        llvm::Value* codeGen();
    private:
        Identifier* identifier;
        ClassBlock* body;
    };

    class Module {
    public:
        Module(string n, Block* b);
    private:
        // todo: exports
        string name;
        Block* body;
    };
} // end namespace 