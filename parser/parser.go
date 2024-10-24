package parser

import (
    "monkey/ast"
    "monkey/lexer"
    "monkey/token"
)


//Struct for our Parser

type Parser struct {
    l *lexer.Lexer
    curToken token.Token
    peekToken token.Token
}

// Constructor
func New(l *lexer.Lexer) *Parser {
    p := &Parser{l: l}


    // Read two tokens so curr and peek both are setup
    p.nextToken()
    p.nextToken()
    
    return p
}

// Function to itterate over tokens and set peek and curr
func (p *Parser) nextToken() {
    p.curToken = p.peekToken

    p.peekToken = p.l.NextToken()

}

func (p *Parser) ParseProgram() *ast.Program {
    return nil
}
