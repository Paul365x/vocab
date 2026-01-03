module github.com/vocab/internal/translate

go 1.24.11

require github.com/PuerkitoBio/goquery v1.11.0

require (
	github.com/andybalholm/cascadia v1.3.3 // indirect
	golang.org/x/net v0.47.0 // indirect
)

require (
	internal/config v1.0.0
)
replace internal/config => ../config
