package token

type TokenType string


type Token struct{
    Type TokenType
    Literal string
}
var keywords = map[string]TokenType{
    "fn": FUNCTION,
    "let": LET,
    "true": TRUE,
    "false": FALSE,
    "if": IF,
    "else": ELSE,
    "return": RETURN,
}

func LookupIdent(ident string) TokenType{
    if tok, ok := keywords[ident]; ok{
        return tok
    }
    return IDENT
}

const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    //Indentifiers and Literals
    IDENT = "IDENT"
    INT   = "INT"
    

    EQ = "=="
    NOT_EQ = "!="

    //Operation
    ASSIGN = "="
    PLUS   = "+"
    MINUS  = "-"
    BANG   = "!"
    ASTERISK = "*"
    SLASH    = "/"

    LT = "<"
    GT = ">"

    //Delimiters
    COMMA   = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    //KEYWORDS
    FUNCTION = "FUNCTION"
    LET = "LET"
    TRUE = "TRUE"
    FALSE = "FALSE"
    IF = "IF"
    ELSE = "ELSE"
    RETURN = "RETURN"
)


