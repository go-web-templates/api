package logger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type TableLogger struct {}

func NewTableLogger() *TableLogger {
	return &TableLogger{}
}

func (t *TableLogger) getFormatedTime() string {
	now := time.Now()

	return fmt.Sprintf("%02d:%02d:%02d", now.Hour(), now.Minute(), now.Second())
}

func (t *TableLogger) getPrefixFormat(label string, labelColor string) string {
	return fmt.Sprintf("%s | %s%s%s | ", t.getFormatedTime(), labelColor, label, _COLOR_RESET)
}

func (t *TableLogger) getInfosFormat(infos []any) string {
	infosStrList := make([]string, len(infos))
	for i, info := range infos {
		infosStrList[i] = fmt.Sprintf("%v", info)
	}

	return strings.Join(infosStrList, " | ")
}

func (t *TableLogger) getFormatedLog(label string, labelColor string, infos []any) string {
	return t.getPrefixFormat(label, labelColor) + t.getInfosFormat(infos) + "\n"
}

func (t *TableLogger) Warning(infos ...any) {
	fmt.Printf(t.getFormatedLog("WAR", _COLOR_WAR, infos))
}

func (t *TableLogger) Info(infos ...any) {
	fmt.Printf(t.getFormatedLog("INF", _COLOR_INF, infos))
}

func (t *TableLogger) Error(infos ...any) {
	fmt.Printf(t.getFormatedLog("ERR", _COLOR_ERR, infos))
}

func (t *TableLogger) Fatal(infos ...any) {
	fmt.Printf(t.getFormatedLog("FAT", _COLOR_ERR, infos))
	os.Exit(1)
}

func (t *TableLogger) Format(format string, data ...any) string {
	return fmt.Sprintf(format, data...)
}
