package service

// import (
// 	"testing"

// 	"../../infrastructure/entity/"
// )

// func TestTaskService_FindByID(tt *testing.T) {
// 	tt.Run(
// 		"正常系: エラーなし",
// 		func(t *testing.T) {
// 			task := &entity.Task{
// 				ID:        1,
// 				Name:      "Alice",
// 				CreatedAt: common.CurrentTime(),
// 				UpdatedAt: common.CurrentTime(),
// 			}

// 			// リポジトリのモックを作成
// 			r := new(mocks.IUserRepository)

// 			// モックの.FindByID(1)メソッドが呼び出されることを期待する。呼び出されたら(user, nil)を返す
// 			r.On("FindByID", user.ID).Return(user, nil)

// 			// テスト対象のサービス
// 			s := NewUserService(r)

// 			// サービスのメソッドを実行
// 			ret, err := s.FindByID(user.ID)
// 			// エラーがないことを期待する
// 			assert.NoError(t, err)

// 			// 各種フィールドが期待通りか確認
// 			assert.Equal(t, ret.ID, user.ID)
// 			assert.Equal(t, ret.Name, user.Name)
// 			assert.Equal(t, ret.CreatedAt, user.CreatedAt)
// 			assert.Equal(t, ret.UpdatedAt, user.UpdatedAt)

// 			// 上記で指定した通りの引数でメソッドが呼ばれることを期待する
// 			r.AssertExpectations(t)
// 		})
// }
