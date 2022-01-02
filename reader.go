package mandarinfcard

import (
	"bufio"
	"io"
	"strings"
)

type Word struct {
	Mandarin string
	Pinyin   string
	English  string
}

func ReadAll(src io.Reader) []Word {
	buf := bufio.NewScanner(src)

	dict := make([]Word, 0, 1000)
	for buf.Scan() {
		txt := buf.Text()

		cols := strings.Split(txt, ";")
		dict = append(dict, Word{
			Mandarin: cols[0],
			Pinyin:   cols[1],
			English:  cols[2],
		})
	}
	return dict
}
