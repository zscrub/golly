package lexer

import "github.com/zscrub/golly/token"

type Lexer struct {
        input         string
        position      int
        readPosition  int
        ch            byte
}

func New(input string) *Lexer {
        l := &Lexer{input: input}
        l.readChar()
        return l
}

func (l *Lexer) readChar() {
        if l.readPosition >= len(l.input) {
                l.ch = 0
        } else {
                l.ch = l.input[l.readPosition]
        }
        l.position = l.readPosition
        l.readPosition += 1
}


func (l *Lexer) peekChar() byte {
  if l.readPosition >= len(l.input) {
    return 0
  }

  return l.input[l.readPosition]
}

func (l *Lexer) NextToken() token.Token {
        var tok token.Token
        l.skipWhiteSpace()
        
        switch l.ch {
        case '=':
          if l.peekChar() == '=' {
            tok = l.comparisonToken(token.EQ)
          } else {
            tok = newToken(token.ASSIGN, l.ch)
          }
        case ';':
          tok = newToken(token.SEMICOLON, l.ch)
        case '(':
          tok = newToken(token.LPAREN, l.ch)
        case ')':
          tok = newToken(token.RPAREN, l.ch)
        case ',':
          tok = newToken(token.COMMA, l.ch) 
        case '+':
          tok = newToken(token.PLUS, l.ch)
        case '-':
          tok = newToken(token.MINUS, l.ch)
        case '{':
          tok = newToken(token.LBRACE, l.ch)
        case '}':
          tok = newToken(token.RBRACE, l.ch)
        case '*':
          tok = newToken(token.MULT, l.ch)
        case '/':
          tok = newToken(token.DIV, l.ch)
        case '^':
          tok = newToken(token.EXP, l.ch)
        case '!':
          if l.peekChar() == '=' {
            tok = l.comparisonToken(token.NOT_EQ)
          } else {
            tok = newToken(token.BANG, l.ch)
          }
        case '<':
          tok = newToken(token.LT, l.ch)
        case '>':
          tok = newToken(token.GT, l.ch)
        case 0:
          tok.Literal = ""
          tok.Type = token.EOF
        default:
          if isLetter(l.ch) {
            tok.Literal = l.readIdentifier(isLetter)
            tok.Type = token.LookupIdent(tok.Literal)
            return tok
          } else if isDigit(l.ch) { 
            tok.Type = token.INT
            tok.Literal = l.readIdentifier(isDigit)
            return tok
          } else {
            tok = newToken(token.ILLEGAL, l.ch)
          }
        }

        l.readChar()
        return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
        return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) comparisonToken(tokenType token.TokenType) token.Token {
  ch := l.ch
  l.readChar()
  literal := string(ch) + string(l.ch)
  return token.Token{Type: tokenType, Literal: literal}
}

func (l *Lexer) skipWhiteSpace() {
  for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
    l.readChar()
  }
}

func (l *Lexer) readIdentifier(checkIdentifier func(byte) bool) string {
  position := l.position
  for checkIdentifier(l.ch) {
    l.readChar()
  }
  return l.input[position:l.position]
}

func isLetter(ch byte) bool {
  return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
  return '0' <= ch && ch <= '9'
}
