package infrastructure

// import (
// 	"go-rest-api-example/app/common"
// 	"go-rest-api-example/app/infrastructure/mysql/entity"
// 	"testing"

// 	"github.com/stretchr/testify/assert"

// 	"github.com/DATA-DOG/go-sqlmock"
// )

// func TestUserRepository_Create(tt *testing.T) {
// 	tt.Run(
// 		"正常系: エラーなし",
// 		func(t *testing.T) {
// 			user := &entity.User{
// 				Name:      "Alice",
// 				CreatedAt: common.CurrentTime(),
// 			}

// 			// モック用のコネクションを作成
// 			db, mock, err := sqlmock.New()
// 			assert.NoError(t, err)
// 			defer db.Close()

// 			// SQｌ、引数、戻り値が意図したものであることを期待する
// 			mock.ExpectPrepare(`INSERT INTO users`).
// 				ExpectExec().
// 				WithArgs(user.Name, user.CreatedAt).
// 				WillReturnResult(sqlmock.NewResult(1, 1))

// 				// テスト対象のリポジトリを作成
// 			r := NewUserRepository(db)

// 			// エラーが発生しないことを期待する
// 			assert.NoError(t, r.Create(user))

// 			// 上記で指定した通りにモックが呼ばれることを期待する
// 			assert.NoError(t, mock.ExpectationsWereMet())
// 		})
// }
