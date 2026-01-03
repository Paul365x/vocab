module github.com/vocab

go 1.24.11

require github.com/spf13/cobra v1.10.2

require (
	github.com/PuerkitoBio/goquery v1.11.0 // indirect
	github.com/andybalholm/cascadia v1.3.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
	golang.org/x/net v0.47.0 // indirect
)

require (
	github.com/dariubs/percent v1.0.0
	internal/config v1.0.0
	internal/translate v0.0.0-00010101000000-000000000000
	internal/util v1.0.0
)

replace internal/config => ./internal/config

replace internal/util => ./internal/util

replace internal/translate => ./internal/translate
