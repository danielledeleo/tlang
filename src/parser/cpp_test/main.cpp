#include <iostream>
#include <strstream>
#include <string>
#include "antlr4-runtime.h"
#include "parser/tlangLexer.h"
#include "parser/tlangParser.h"
#include "parser/tlangBaseListener.h"
#include "parser/tlangBaseVisitor.h"
#include "TTypes.h"

using namespace antlr4;
using namespace std;

class compilerVisitor: public tlangBaseVisitor {
public:
    unordered_map<string, tlang::VariableDeclaration*> vars; 

    virtual antlrcpp::Any visitVariableDeclaration(tlangParser::VariableDeclarationContext *ctx) {
        auto idents = ctx->variableIdentifierList();
        auto type = resolveType(ctx->typeSpec());
        
        for(size_t i = 0; idents->variableIdentifier(i) != nullptr; i++) {
            auto ident = idents->variableIdentifier(i);
            auto name = ident->getText();
            auto t_ident = new tlang::Identifier(name);
            auto lineno = ident->getStart()->getLine();
            auto charno = ident->getStart()->getCharPositionInLine();
            auto variable = new tlang::VariableDeclaration(t_ident, type, lineno, charno);

            if (!vars.insert(make_pair(name, variable)).second) {
                cerr << "(line:" << lineno << ":" << charno;
                cerr << ") parsing error -- '" << name << "' is already defined on (line:";
                cerr << vars[name]->lineNo << ":" << vars[name]->charNo << ")" << endl;
                delete variable;
            }
        }

        if (ctx->expression()) {
            visit(ctx->expression());
        }

        return NULL;
    }

    tlang::Type* resolveType(tlangParser::TypeSpecContext *ctx) {
        if (ctx->basicType()) {
            auto t = ctx->basicType();
            if (t->INT())             { return new tlang::IntegerType(); }
            else if (t->REAL())       { return new tlang::RealType(); }
            else if (t->BOOLEAN())    { return new tlang::BooleanType(); }
            else if (t->NAT())        { return new tlang::NaturalType(); }
            else if (t->INTN())       { return new tlang::IntegerNType(); }
            else if (t->NATN())       { return new tlang::NaturalNType(); }
            else if (t->REALN())      { return new tlang::RealNType(); }
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

    virtual antlrcpp::Any visitExpression(tlangParser::ExpressionContext *ctx) {
        if (ctx->prefix != NULL) {
            cout << "prefix: " << ctx->prefix->getText() << endl;
            visit(ctx->expression(0));
        } else if (ctx->bop != NULL) {
            cout << "binop:" << ctx->bop->getText() << endl;
            visit(ctx->expression(0));
            visit(ctx->expression(1));
        } else if (ctx->primaryExpression() != NULL) {
            // wait, we might not be done
            if (ctx->primaryExpression()->L_PAREN()) {
                visit(ctx->primaryExpression());
            } else {
                cout << "we've hit the bottom: " << ctx->primaryExpression()->getText() << endl;
            }
        } else {
            cout << "method call: " << ctx->expression(0)->getText() << "(";
            cout << ctx->expressionList()->getText() << ")" << endl;
        }

        return NULL;
    }

    string handleTypeSpec(tlangParser::TypeSpecContext *ctx) {
        if (ctx->basicType()) {
            return ctx->basicType()->getText();

        } else if (ctx->arrayDeclaration()) {
            auto range = ctx->arrayDeclaration()->indexType();
            std::stringstream ss;
            ss << "[" << range->INTEGER_LITERAL(0)->getText() << ":";
            ss << range->INTEGER_LITERAL(1)->getText() << "] of ";
            ss << handleTypeSpec(ctx->arrayDeclaration()->typeSpec());
            return ss.str();
            
        } else if (ctx->stringType()) {
            std::stringstream ss;
            ss << "string";
            if (ctx->stringType()->INTEGER_LITERAL()) {
                ss << "(" << ctx->stringType()->INTEGER_LITERAL()->getText() << ")";
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

    std::ifstream stream;
    stream.open(filename);

    ANTLRInputStream input(stream);
    tlangLexer lexer(&input);
    CommonTokenStream tokens(&lexer);
    tlangParser parser(&tokens);
    compilerVisitor visitor;
    
    visitor.visit(parser.program());
    
    for (auto const& x : visitor.vars) {
        cout << x.first << " -> ";

        if (x.second->getType()) {
            cout << x.second->getType()->typeName();
        }
        cout << endl;
    }

    stream.close();
}
