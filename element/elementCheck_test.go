package elementcheck

import (
	"testing"
)

// TestRemoveElement テスト関数
// ref: https://qiita.com/saya713y/items/f7ee07e8f12ab85ed9bf
func TestRemoveElement(t *testing.T) {
	t.Run("len=7", func(t *testing.T) {
		// 期待する値
		exp := 5

		if RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2) != exp {
			// 下記だと処理は終了される
			// t.Fail() || t.FailNow()
			// t.Fatal("結果:", res, "期待:", exp) || t.Fatalf()

			// 処理は継続(t.Errorf()も同じ)
			t.Error("期待する値と一致しませんでした。")
			t.Fail()
		}
	})

	t.Run("len=4", func(t *testing.T) {
		// 期待する値
		exp := 2

		if RemoveElement([]int{3, 2, 2, 3}, 3) != exp {
			t.Error("期待する値と一致しませんでした。")
			t.Fail()
		}
	})

	// ログ出力
	t.Log("TestRemoveElementの確認終了")
}
