module github.com/rtkrmccaw/kafka-go

go 1.15

require (
	github.com/klauspost/compress v1.15.9
	github.com/kr/pretty v0.1.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.15
	github.com/stretchr/testify v1.8.1
	github.com/xdg-go/scram v1.1.2
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/net v0.17.0
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

retract [v0.4.36, v0.4.37]
