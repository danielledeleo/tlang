#include <iostream>
#include <strstream>
#include <string>
#include "antlr4-runtime/antlr4-runtime.h"
#include "parser/TuringLexer.h"
#include "parser/TuringParser.h"
#include "parser/TuringBaseListener.h"
#include "parser/TuringBaseVisitor.h"
#include "types.h"

using namespace antlr4;
using namespace std;

class tlangVisitor: public TuringBaseVisitor {
public:
    map<string, tlangIdentifier*> variables;

    virtual antlrcpp::Any visitVariableDeclaration(TuringParser::VariableDeclarationContext *ctx) {
        auto start = ctx->getStart();
        auto variable = new tlangIdentifier();
        auto idents = ctx->variableIdentifierList();
        
        // variable->identifier = handleTypeSpec(ctx->typeSpec());
        variable->line_no = start->getLine();
        variable->char_no = start->getCharPositionInLine();
        variable->type = resolveType(ctx->typeSpec());
        
        for(size_t i = 0; idents->variableIdentifier(i); i++) {
            auto name = idents->variableIdentifier(i)->getText();
            if (!variables.insert(make_pair(name, variable)).second) {
                cerr << "(line:" << start->getLine() << ":" << start->getCharPositionInLine();
                cerr << ") parsing error -- '" << name << "' is already defined on (line:";
                cerr << variables[name]->line_no << ":" << variables[name]->char_no << ")" << endl;
            }
        }

        if (ctx->expression()) {
            visit(ctx->expression());
        }

        return NULL;
    }

    tlangType* resolveType(TuringParser::TypeSpecContext *ctx) {
        if (ctx->basicType()) {
            auto t = ctx->basicType();
            tlangPrimitive primitive;
            if (t->INT())             { primitive = tl_int; }
            else if (t->REAL())       { primitive = tl_real; }
            else if (t->BOOLEAN())    { primitive = tl_boolean; }
            else if (t->NAT())        { primitive = tl_nat; }
            else if (t->INTN())       { primitive = tl_intn; }
            else if (t->NATN())       { primitive = tl_natn; }
            else if (t->REALN())      { primitive = tl_realn; }
            else if (t->CHAR())       { primitive = tl_char; }

            return new tlangPrimitiveType(primitive);
        } else {
            tlangComposite composite;
            if (ctx->arrayDeclaration())      { composite = tl_array; }
            else if (ctx->classDeclaration()) { composite = tl_class; }
            else if (ctx->stringType())       { composite = tl_string; }
            else if (ctx->recordType())       { composite = tl_record; }

            return new tlangCompositeType(composite);
        }
        
        return NULL;
    }

    virtual antlrcpp::Any visitExpression(TuringParser::ExpressionContext *ctx) {
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

    string handleTypeSpec(TuringParser::TypeSpecContext *ctx) {
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

    virtual antlrcpp::Any visitTypeSpec(TuringParser::TypeSpecContext *ctx) {
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
    TuringLexer lexer(&input);
    CommonTokenStream tokens(&lexer);
    TuringParser parser(&tokens);
    tlangVisitor visitor;


    // parser.addParseListener(&listener);
    
    visitor.visit(parser.program());
    
    for (auto const& x : visitor.variables) {
        cout << x.first << " -> "
             << x.second->type->group()
             << endl;
    }

    stream.close();
}
