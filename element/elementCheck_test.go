package elementcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRemoveElement テスト関数
func TestRemoveElement(t *testing.T) {
	// TableDriven
	// PHPUnitのdataProviderのように振る舞わせる
	// ref: https://qiita.com/atotto/items/f6b8c773264a3183a53c
	cases := []struct {
		nums []int
		val  int
		exp  int
	}{
		{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, exp: 5},
		{nums: []int{3, 2, 2, 3}, val: 2, exp: 2},
	}

	for _, c := range cases {
		ret := new(ElementCheck).RemoveElement(c.nums, c.val)
		// ref: https://qiita.com/hein946/items/a4a4b9241b14bbff6299
		assert.Equal(t, ret, c.exp, "返り値と期待する値が一緒か確認する")
	}

	// ref: https://qiita.com/saya713y/items/f7ee07e8f12ab85ed9bf
	// t.Run("len=7", func(t *testing.T) {
	// 	// 期待する値
	// 	exp := 5

	// 	if RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2) != exp {
	// 		// 下記だと処理は終了される
	// 		// t.Fail() || t.FailNow()
	// 		// t.Fatal("結果:", res, "期待:", exp) || t.Fatalf()

	// 		// 処理は継続(t.Errorf()も同じ)
	// 		t.Error("期待する値と一致しませんでした。")
	// 		t.Fail()
	// 	}
	// })

	// ログ出力
	t.Log("TestRemoveElementの確認終了")
}
