module github.com/zscrub/golly

go 1.19

require github.com/zscrub/golly/repl v0.0.0

require (
	github.com/zscrub/golly/lexer v0.0.0-20250217000543-75df7e8c5be8 // indirect
	github.com/zscrub/golly/token v0.0.0-20250217000543-75df7e8c5be8 // indirect
)

replace (
	github.com/zscrub/golly/lexer => ./lexer

	github.com/zscrub/golly/repl => ./repl

	github.com/zscrub/golly/token => ./token
)
