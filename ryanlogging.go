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
	newFields := make([]zap.Field, 0)
	newFields = append(newFields, zap.String("func", fname))
	newFields = append(newFields, zap.Int("seq", seq))
	newFields = append(newFields, fields...)

	lfunc("--RYAN-- IAM", newFields...)
}
