module github.com/vocab

go 1.23.2

require github.com/spf13/cobra v1.10.2

require (
	github.com/dariubs/percent v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
)

require (
	internal/config v1.0.0
	internal/util v1.0.0
)

replace internal/config => ./internal/config

replace internal/util => ./internal/util
