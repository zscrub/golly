module github.com/zscrub/golly/ast

go 1.19

require (
	github.com/zscrub/golly/token v0.0.0-20250217000543-75df7e8c5be8
)

replace (
	github.com/zscrub/golly/token => ../token
	github.com/zscrub/golly/lexer => ../lexer
)