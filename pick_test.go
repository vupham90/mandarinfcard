package mandarinfcard

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestPick(t *testing.T) {
	now := time.Date(2021, time.January, 10, 0, 0, 0, 0, time.Local)

	picked := Pick(now, 0, 10)

	if !reflect.DeepEqual(picked, 10) {
		t.Errorf("picked = %v", picked)
	}
}

func TestPickerDistribution(t *testing.T) {
	size := 100
	data := make(map[int]int, size)

	now := time.Now()
	for i := 0; i < 10000; i++ {
		newDate := now.Add(time.Hour * 24 * time.Duration(i))

		picked := Pick(newDate, 0, size)

		if _, ok := data[picked]; !ok {
			data[picked] = 0
		}
		data[picked]++
	}

	fmt.Println(data)
}
