#include <iostream>
#include <strstream>
#include <string>
#include "antlr4-runtime.h"
#include "parser/tlangLexer.h"
#include "parser/tlangParser.h"
#include "parser/tlangBaseListener.h"
#include "parser/tlangBaseVisitor.h"
#include "TTypes.h"

#include "llvm/ADT/StringRef.h"
#include "llvm/IR/LLVMContext.h"
#include "llvm/IR/IRBuilder.h"
#include "llvm/IR/Module.h"
#include "llvm/IR/Verifier.h"


using namespace antlr4;
using namespace std;

static llvm::LLVMContext TheContext;
static llvm::IRBuilder<> Builder(TheContext);

void printErr(size_t lineno, size_t charno, string err_type, string err_message) {
    cerr << "(line:" << lineno << ":" << charno;
    cerr << ") " << err_type << " error -- " << err_message << endl;
}

class compilerVisitor: public tlangBaseVisitor {
public:
    unique_ptr<llvm::Module> TheModule;
    unordered_map<string, tlang::VariableDeclaration*> vars;

    virtual antlrcpp::Any visitProgram(tlangParser::ProgramContext *ctx) {
        auto ft = llvm::FunctionType::get(llvm::Type::getInt64Ty(TheContext), false);
        auto f = llvm::Function::Create(ft, llvm::Function::ExternalLinkage, "__entry_point", TheModule.get());

        llvm::BasicBlock *BB = llvm::BasicBlock::Create(TheContext, "entry", f);
        Builder.SetInsertPoint(BB);
        
        for(size_t i = 0; ctx->topLevel(i) != nullptr; i++) {
            auto tl = ctx->topLevel(i);
            // cout << " -- " << i << " -- " << endl;
            
            if (tl->statementOrDeclaration()) {
                auto decl = tl->statementOrDeclaration()->declaration();
                auto stmt = tl->statementOrDeclaration()->statement();
                if (decl) {
                    if (decl->subprogramDeclaration()) {
                    } else if (decl->variableDeclaration()) {
                        auto vars = handleVariableDeclaration(decl->variableDeclaration());
                    }

                } else if (stmt) {
                    if (stmt->expression()) {
                        Builder.CreateRet(handleExpression(stmt->expression()));
                        llvm::verifyFunction(*f);
                    }
                }
            }
        }

        return 0;
    }


    vector<tlang::VariableDeclaration*>* handleVariableDeclaration(tlangParser::VariableDeclarationContext *ctx) {
        auto idents = ctx->variableIdentifierList();
        auto type = resolveType(ctx->typeSpec());
        auto variables = new vector<tlang::VariableDeclaration*>();
        
        for(size_t i = 0; idents->variableIdentifier(i) != nullptr; i++) {
            auto ident = idents->variableIdentifier(i);
            auto name = ident->getText();
            auto t_ident = new tlang::Identifier(name);
            auto lineno = ident->getStart()->getLine();
            auto charno = ident->getStart()->getCharPositionInLine();
            auto val = createValue(ctx->typeSpec(), ctx->expression());
            auto variable = new tlang::VariableDeclaration(t_ident, type, lineno, charno, val);
            variables->push_back(variable);

            if (!vars.insert(make_pair(name, variable)).second) {
                stringstream err_message;
                err_message << name << "' is already defined on (line:" <<
                    vars[name]->lineNo << ":" << vars[name]->charNo << ")";
                
                printErr(lineno, charno, "semantic", err_message.str());
                delete variable;
                // exit(1);
            }   
        }

        return variables;
    }

    llvm::Value *createValue(tlangParser::TypeSpecContext *ts_ctx, tlangParser::ExpressionContext *e_ctx) {
        if (ts_ctx->basicType()) {
            auto t = ts_ctx->basicType();
            if (t->INT()) {
                uint64_t v = 0;
                if (e_ctx) {
                    v = atoll(e_ctx->getText().c_str());
                }
                return llvm::ConstantInt::get(TheContext, llvm::APInt(32, v, true));
            } else if (t->REAL()) {
                double v = 0;
                if (e_ctx) {
                    v = atof(e_ctx->getText().c_str());
                }
                return llvm::ConstantFP::get(TheContext, llvm::APFloat(v));
            }
        }
        return 0;
    }

    tlang::Type* resolveType(tlangParser::TypeSpecContext *ctx) {
        if (ctx->basicType()) {
            auto t = ctx->basicType();
            if (t->INT())             { return new tlang::IntegerType(); }
            else if (t->REAL())       { return new tlang::RealType(); }
            else if (t->BOOLEAN())    { return new tlang::BooleanType(); }
            else if (t->NAT())        { return new tlang::NaturalType(); }
            else if (t->INT1())       { return new tlang::Integer1Type(); }
            else if (t->INT2())       { return new tlang::Integer2Type(); }
            else if (t->INT4())       { return new tlang::Integer4Type(); }
            else if (t->INT8())       { return new tlang::Integer8Type(); }
            else if (t->NAT1())       { return new tlang::Natural1Type(); }
            else if (t->NAT2())       { return new tlang::Natural2Type(); }
            else if (t->NAT4())       { return new tlang::Natural4Type(); }
            else if (t->NAT8())       { return new tlang::Natural8Type(); }
            else if (t->REAL4())      { return new tlang::Real4Type(); }
            else if (t->REAL8())      { return new tlang::Real8Type(); }
            else if (t->CHAR())       { return new tlang::CharType(); }
            
        } else {
            if (ctx->arrayDeclaration()) { 
                return new tlang::ArrayType(resolveType(ctx->arrayDeclaration()->typeSpec()));
            }
            else if (ctx->classDeclaration()) { }
            else if (ctx->stringType())       { return new tlang::StringType(); }

            // iterates through record fields, resolves each type, and constructs a RecordType object
            else if (ctx->recordType()) {
                vector<tlang::Type*> members;            
                auto fields = ctx->recordType()->recordField();
                for_each(fields.begin(), fields.end(), [this, &members](tlangParser::RecordFieldContext *ctx) -> void {
                    members.push_back(resolveType(ctx->typeSpec()));
                });
                return new tlang::RecordType(members);
            }

            // return new tlang::CompositeType(composite);
        }
        
        return NULL;
    }

    llvm::Value* handleExpression(tlangParser::ExpressionContext *ctx) {
        
        if (ctx->prefix) {
            // cout << "prefix: " << ctx->prefix->getText() << endl;
            visit(ctx->expression(0));
        } else if (ctx->bop) {
            auto v = handleBinOp(ctx->bop, ctx->expression(0), ctx->expression(1));
            return v;
        } else if (ctx->primaryExpression()) {
            // wait, we might not be done
            if (ctx->primaryExpression()->literal()) {
                return handleLiteral(ctx->primaryExpression()->literal());
            } else if (ctx->primaryExpression()->L_PAREN()) {
                return handleExpression(ctx->primaryExpression()->expression());
            } else {
                // cout << ctx->primaryExpression()->getText();
            }
        } else {
            // cout << "method call: " << ctx->expression(0)->getText() << "(";
            // cout << ctx->expressionList()->getText() << ")" << endl;
        }

        return NULL;
    }

    llvm::Value *handleBinOp(Token *bop, tlangParser::ExpressionContext *LHS, tlangParser::ExpressionContext *RHS) {        
        auto op = bop->getType();
        auto lhs_v = handleExpression(LHS);
        auto rhs_v = handleExpression(RHS);

        switch (op) {
            case tlangParser::PLUS:
                return Builder.CreateAdd(lhs_v, rhs_v, "add"); break;
            case tlangParser::MINUS:
                return Builder.CreateSub(lhs_v, rhs_v, "minus"); break;
            default:
                break;
        }
        return nullptr;
    }

    // all integer literals treated at 64 bit to start and casted down later
    llvm::Value* handleIntLiteral(tlangParser::Integer_literalContext *l, const char* str, int base, bool is_signed) {
        char *strtolchar;
        uint64_t i = strtoull(str, &strtolchar, base);

        if (errno == ERANGE) {
            printErr(l->getStart()->getLine(), 
                l->getStart()->getCharPositionInLine(), 
                "semantic", "integer literal too big");
        }
        return llvm::ConstantInt::get(TheContext, llvm::APInt(64, i, is_signed)); 
    }

    llvm::Value* handleLiteral(tlangParser::LiteralContext *ctx) {
        char *strtolchar;
        if (ctx->integer_literal()) {
            auto l = ctx->integer_literal();
            if (l->DECIMAL_LITERAL()) {
                if (l->sign) {
                    // signed integer
                    if (l->sign->getType() == tlangParser::MINUS) {
                        return handleIntLiteral(l, l->getText().c_str(), 10, true);
                    } 
                } else {
                    // unsigned integer
                    return handleIntLiteral(l, l->getText().c_str(), 10, false);
                }
            } else if (l->HEX_LITERAL()) {
                // shift start of string by length of prefix
                return handleIntLiteral(l, l->HEX_LITERAL()->getText().c_str() + 2, 16, false);
            } else if (l->OCTAL_LITERAL()) {
                return handleIntLiteral(l, l->OCTAL_LITERAL()->getText().c_str() + 1, 8, false);
            } else if (l->BINARY_LITERAL()) {
                return handleIntLiteral(l, l->BINARY_LITERAL()->getText().c_str() + 2, 2, false);
            }
        } else if (ctx->REAL_LITERAL()) {
            double r = atof(ctx->getText().c_str());
            return llvm::ConstantFP::get(TheContext, llvm::APFloat(r));;
        } else if (ctx->STRING_LITERAL()) {
            cout << "String literal!!!" << endl;
            auto str = ctx->STRING_LITERAL()->getText();
            return Builder.CreateGlobalString(llvm::StringRef(str), "someStr");
        }
        return nullptr;
    }

    string handleTypeSpec(tlangParser::TypeSpecContext *ctx) {
        if (ctx->basicType()) {
            return ctx->basicType()->getText();

        } else if (ctx->arrayDeclaration()) {
            auto range = ctx->arrayDeclaration()->indexType();
            stringstream ss;
            ss << "[" << range->integer_literal(0)->getText() << ":";
            ss << range->integer_literal(1)->getText() << "] of ";
            ss << handleTypeSpec(ctx->arrayDeclaration()->typeSpec());
            return ss.str();
            
        } else if (ctx->stringType()) {
            stringstream ss;
            ss << "string";
            if (ctx->stringType()->integer_literal()) {
                ss << "(" << ctx->stringType()->integer_literal()->getText() << ")";
            }
            return ss.str();
        } else {
            return "[ unimplemented ]";
        }
    }

    virtual antlrcpp::Any visitTypeSpec(tlangParser::TypeSpecContext *ctx) {
        return NULL;
    }
};

int main(int argc, const char* argv[]) {
    string filename;
    if (argc > 1) {
        filename = argv[1];
    } else {
        cerr << "missing source filename" << endl;
        exit(-1);
    }

    ifstream stream;
    stream.open(filename);

    ANTLRInputStream input(stream);
    tlangLexer lexer(&input);
    CommonTokenStream tokens(&lexer);
    tlangParser parser(&tokens);
    compilerVisitor visitor;
    visitor.TheModule = llvm::make_unique<llvm::Module>("tlang", TheContext);
    
    visitor.visit(parser.program());
    
    for (auto const& x : visitor.vars) {
        cout << x.first << " -> ";

        if (x.second->getType()) {
            cout << x.second->getType()->typeName();
            // if (x.second->value) {
            if (x.second->value) 
                cout << " | " <<x.second->value->getValueID();
            // }
            // if (x.second->getType()->group() == "primitive") {
            //     cout << "<" << ((tlang::PrimitiveType*)x.second->getType())->sizeOf() * 8 << ">";
            // }
        }
        cout << endl;
    }
    cout << "--------------" << endl;

    visitor.TheModule->print(llvm::errs(), nullptr);

    stream.close();
}
