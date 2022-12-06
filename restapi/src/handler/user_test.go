package handler

// import (
// 	"go-rest-api-example/app/common"
// 	"go-rest-api-example/app/common/dto"
// 	"go-rest-api-example/app/common/validation"
// 	"go-rest-api-example/mocks"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/go-playground/validator"
// 	"github.com/stretchr/testify/assert"

// 	"github.com/labstack/echo/v4"
// )

// func TestUserHandler_FindByID(tt *testing.T) {
// 	tt.Run(
// 		"正常系: エラーなし",
// 		func(t *testing.T) {
// 			// 現在時刻を取得
// 			now := common.CurrentTime()

// 			user := dto.NewUserModel(1, "Alice", now, now)

// 			// GETリクエストを作成
// 			req := httptest.NewRequest(http.MethodGet, "/user/detail?id=1", nil)

// 			// Echoインスタンスの作成
// 			e := echo.New()
// 			e.Validator = &validation.CustomValidator{V: validator.New()}

// 			//新しい コンテキストを作成
// 			rec := httptest.NewRecorder()
// 			c := e.NewContext(req, rec)

// 			// サービスのモックを作成
// 			s := new(mocks.IUserService)

// 			// モックの .FindByID(1)メソッドが呼ばれることを期待する。呼ばれたら(user, nil)を返す。
// 			s.On("FindByID", user.ID).Return(user, nil)

// 			// テスト対象のハンドラ
// 			h := NewUserHandler(s)

// 			// エラーがないことを期待する
// 			if assert.NoError(t, h.FindByID(c)) {
// 				// ステータスコードが200であることを期待する
// 				assert.Equal(t, http.StatusOK, rec.Code)

// 				// 上記で指定した通りにモックが呼び出されることを期待する
// 				s.AssertExpectations(t)
// 			}
// 		})
// 	tt.Run(
// 		"準正常系: idに0を指定→バリデーションに失敗",
// 		func(t *testing.T) {
// 			req := httptest.NewRequest(http.MethodGet, "/user/detail?id=0", nil)

// 			e := echo.New()
// 			e.Validator = &validation.CustomValidator{V: validator.New()}
// 			rec := httptest.NewRecorder()
// 			c := e.NewContext(req, rec)

// 			s := new(mocks.IUserService)
// 			h := NewUserHandler(s)

// 			assert.EqualError(t, h.FindByID(c), `code=400, message=failed to validation request`)
// 		})
// 	tt.Run(
// 		"準正常系: idに文字列を指定→バインドに失敗",
// 		func(t *testing.T) {
// 			req := httptest.NewRequest(http.MethodGet, "/user/detail?id=abc", nil)

// 			e := echo.New()
// 			e.Validator = &validation.CustomValidator{V: validator.New()}
// 			rec := httptest.NewRecorder()
// 			c := e.NewContext(req, rec)

// 			s := new(mocks.IUserService)
// 			h := NewUserHandler(s)

// 			// エラーメッセージが期待通りか確認
// 			assert.EqualError(t, h.FindByID(c), `code=400, message=failed to bind request`)
// 		})
// }
