package server

import (
	"math/rand"
	"testing"
	"time"
)

const cns01Count = 67
const cns03Count = 138

func TestCalculateContent(t *testing.T) {

	ens01 := make([]byte, 160)
	ens02 := make([]byte, 1)
	ens03 := make([]byte, 321)

	cns01 := make([]rune, cns01Count)
	cns02 := make([]rune, 1)
	cns03 := make([]rune, cns03Count)
	for i := 0; i < 160; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		rd := r.Intn(26) + 65
		ens01[i] = byte(rd)
	}
	for i := 0; i < 1; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		rd := r.Intn(26) + 65
		ens02[i] = byte(rd)
	}
	for i := 0; i < 321; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		rd := r.Intn(26) + 65
		ens03[i] = byte(rd)
	}
	// 含有中文测试
	for i := 0; i < cns01Count; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		rd := r.Intn(26) + 65
		cns01[i] = rune(rd)
	}

	for i := 0; i < 1; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		rd := r.Intn(26) + 65
		cns02[i] = rune(rd)
	}
	for i := 0; i < cns03Count; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		rd := r.Intn(26) + 65
		cns03[i] = rune(rd)
	}
	cns01 = append(cns01, '我')
	cns02 = append(cns02, '我')
	cns03 = append(cns03, '我')

	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "test case 01",
			args: args{string(ens01)},
			want: 1,
		},
		{name: "test case 02",
			args: args{string(ens02)},
			want: 1,
		},
		{name: "test case 03",
			args: args{string(ens03)},
			want: 3,
		},
		// en case end
		// cn case start
		{name: "test case 01",
			args: args{string(cns01)},
			want: 1,
		},
		{name: "test case 02",
			args: args{string(cns02)},
			want: 1,
		},
		{name: "test case 03",
			args: args{string(cns03)},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateContent(tt.args.content); got != tt.want {
				t.Errorf("CalculateContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
