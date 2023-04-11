package chapter1

import (
	"testing"
)

func TestAddition(t *testing.T) {
	// 正常系のテストパターン
	success := map[string]struct {
		numA int
		numB int
		want int
	}{
		// FIXME: テストケースを追加
		"success test": {
			numA: 0,
			numB: 10,
			want: 10,
		},
	}
	// エラー系のテストパターン
	fail := map[string]struct {
		numA       int
		numB       int
		wantErrStr string
	}{
		// FIXME: テストケースを追加
		"fail test1": {
			numA:       -1,
			numB:       9,
			wantErrStr: "numA最小値エラー",
		},
		"fail test2": {
			numA:       0,
			numB:       9,
			wantErrStr: "numB最小値エラー",
		},
		"fail test3": {
			numA:       101,
			numB:       10,
			wantErrStr: "numA最大値エラー",
		},
		"fail test4": {
			numA:       100,
			numB:       201,
			wantErrStr: "numB最大値エラー",
		},
	}

	for tt, tc := range success {
		t.Run(tt, func(t *testing.T) {
			got, err := addtion(tc.numA, tc.numB)
			if err != nil {
				t.Errorf("err is not nil: %s", err)
			}
			if tc.want != got {
				t.Errorf("unexpected return. want:%d actual:%d", tc.want, got)
			}
		})
	}
	for tt, tc := range fail {
		t.Run(tt, func(t *testing.T) {
			got, err := addtion(tc.numA, tc.numB)
			if got != 0 {
				t.Errorf("unexpected return. want:0 actual:%d", got)
			}
			if tc.wantErrStr != err.Error() {
				t.Errorf("unexpected err. want:%s actual:%s", tc.wantErrStr, err)
			}
		})
	}
}
