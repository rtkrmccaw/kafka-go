package kafka

import (
	"go.uber.org/zap"
)

func logInfo(fname string, seq int, fields ...zap.Field) {
	rlog(zap.L().Info, fname, seq, fields...)
}

func logError(fname string, seq int, fields ...zap.Field) {
	rlog(zap.L().Error, fname, seq, fields...)
}

func rlog(lfunc func(string, ...zap.Field), fname string, seq int, fields ...zap.Field) {
	newFields := make([]zap.Field, (len(fields) + 2))
	newFields[0] = zap.String("func", fname)
	newFields[1] = zap.Int("seq", seq)
	a := append(newFields, fields...)

	lfunc("--RYAN-- kafka-go", a...)
}
