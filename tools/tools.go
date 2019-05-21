package tools

import (
	"strconv"
	"sync"
	"time"
)

/**
*@Author: PantaZheng
*@CreateAt: 2019/5/20 13:22
*@Title: tools.go
*@Package: tools
*@Description: (用一句话描述该文件该做什么)
@Software: GoLand
*/

type Tools struct {
}

var (
	Tool = New()
	once sync.Once
)

func New() (t *Tools) {
	once.Do(func() {
		t = &Tools{}
	})
	return t
}

func (t *Tools) ParseInt(s string) int {
	if id, err := strconv.Atoi(s); err == nil {
		return id
	} else {
		return 0
	}
}

func (t *Tools) ParseString(i int) string {
	return strconv.Itoa(i)
}

func (t *Tools) TimeFormat(time *time.Time) string {
	return time.Format("2016-01-02")
}
