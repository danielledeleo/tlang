#include <iostream>
#include <strstream>
#include <string>
#include "antlr4-runtime/antlr4-runtime.h"
#include "parser/TuringLexer.h"
#include "parser/TuringParser.h"
#include "parser/TuringBaseListener.h"
#include "parser/TuringBaseVisitor.h"

using namespace antlr4;
using namespace std;

class tlangIdentifier {
public:
    string type;
    int line_no, char_no;
};

class tlangVisitor: public TuringBaseVisitor {
public:
    map<string, tlangIdentifier*> variables;

    virtual antlrcpp::Any visitVariableDeclaration(TuringParser::VariableDeclarationContext *ctx) {
        auto start = ctx->getStart();
        auto variable = new tlangIdentifier();
        auto idents = ctx->variableIdentifierList();
        
        variable->type = handleTypeSpec(ctx->typeSpec());
        variable->line_no = start->getLine();
        variable->char_no = start->getCharPositionInLine();
        
        while (idents != NULL) {
            auto name = idents->variableIdentifier()->getText();
            if (!variables.insert(make_pair(name, variable)).second) {
                cerr << "(line:" << start->getLine() << ":" << start->getCharPositionInLine();
                cerr << ") parsing error -- '" << name << "' is already defined on (line:";
                cerr << variables[name]->line_no << ":" << variables[name]->char_no << ")" << endl;
            }
            
            idents = idents->variableIdentifierList();
        }

        if (ctx->expression()) {
            visit(ctx->expression());
        }

        return NULL;
    }

    virtual antlrcpp::Any visitAdditiveExpression(TuringParser::AdditiveExpressionContext *ctx) {
        cout << "[" << ctx->additiveExpression(0)->getText() << " " << ctx->additiveOperator()->getText();
        cout << " " << ctx->additiveExpression(1)->getText() << "]" << endl;
        return NULL;
    }
    
    virtual antlrcpp::Any visitMultiplicativeExpression(TuringParser::MultiplicativeExpressionContext *ctx) {
        cout << "[" << ctx->multiplicativeExpression(0)->getText() << " " << ctx->multiplicativeOperator()->getText();
        cout << " " << ctx->multiplicativeExpression(1)->getText() << "]" << endl;
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

class tlangListener: public TuringBaseListener {
    virtual void exitVariableDeclaration(TuringParser::VariableDeclarationContext *ctx) {
        auto type = ctx->typeSpec();
        auto idents = ctx->variableIdentifierList();
        
        while (idents != NULL) {
            cout << idents->variableIdentifier()->getText() << " -> " << type << endl; 
            idents = idents->variableIdentifierList();
        }
    }

    virtual void exitTypeSpec(TuringParser::TypeSpecContext *ctx) {
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
    tlangListener listener;
    tlangVisitor visitor;

    // parser.addParseListener(&listener);
    
    visitor.visit(parser.program());
    
    stream.close();
}
