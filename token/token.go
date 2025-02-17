package token

type TokenType string

type Token struct {
  Type     TokenType
  Literal  string
  // FileName   string
  // LineNumber io.Reader
}

const (
  ILLEGAL = "ILLEGAL"
  EOF     = "EOF"

  // Identifiers and Literals
  IDENT = "IDENT" // add, foobar, x, y
  INT   = "INT"

  // Operators
  ASSIGN  = "="
  PLUS    = "+"
  MINUS   = "-"
  MULT    = "*"
  DIV     = "/"
  EXP     = "^"
  BANG    = "!"
  LT      = "<"
  GT      = ">"
  EQ      = "=="
  NOT_EQ  = "!="

  // Delimiters
  COMMA     = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  // Keywords
  VAR = "VAR"
  FN  = "FN"
  IF  = "IF"
  // AND = "AND"
  // OR  = "OR"

  RETURN = "RETURN"
  ELSE   = "ELSE"

  TRUE = "TRUE"
  FALSE = "FALSE"
)

var keywords = map[string]TokenType{
  "fn": FN,
  "var": VAR,
  "if": IF,
  "else": ELSE,
  "return": RETURN,
  "true": TRUE,
  "false": FALSE,
}

func LookupIdent(ident string) TokenType {
  if tok, ok := keywords[ident]; ok {
    return tok
  }
  return IDENT
}
