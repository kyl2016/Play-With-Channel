package main

import "fmt"

func main() {
	t := IntegerConstant{&Match{KEYWORD, "wizard"}, 2}
	fmt.Println(t.Type(), t.Lexeme(), t.Value())

	x := Token(t)
	fmt.Println(x.Type(), x.Lexeme())
}

type TokenType uint16

const (
	KEYWORD TokenType = iota
	IDENTIFIER
	LBRACKET
	RBRACKET
	INT
)

type Token interface {
	Type() TokenType
	Lexeme() string
}

type Match struct {
	toktype TokenType
	lexeme string
}

func (m *Match) Type() TokenType{
	return m.toktype
}

func (m *Match) Lexeme() string{
	return m.lexeme
}

type IntegerConstant struct {
	Token
	value uint64
}


//
//
//type IntegerConstant struct {
//	token Token
//	value uint64
//}

//
//func (i *IntegerConstant) Type() TokenType{
//	return i.token.Type()
//}
//
//func (i *IntegerConstant) Lexeme() string{
//	return i.token.Lexeme()
//}

func (i *IntegerConstant) Value() uint64{
	return i.value
}

