module github.com/Chara-X/search

go 1.23.1

replace github.com/Chara-X/util => ../util

replace github.com/Chara-X/priority => ../priority

require (
	github.com/Chara-X/priority v0.0.0-00010101000000-000000000000
	github.com/Chara-X/util v0.0.0-00010101000000-000000000000
	github.com/kljensen/snowball v0.9.0
)
