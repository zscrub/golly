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
  // AND = "AND"
  // OR  = "OR"
)

var keywords = map[string]TokenType{
  "fn": FN,
  "var": VAR,
}

func LookupIdent(ident string) TokenType {
  if tok, ok := keywords[ident]; ok {
    return tok
  }
  return IDENT
}
