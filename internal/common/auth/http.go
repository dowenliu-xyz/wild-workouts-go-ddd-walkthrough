package auth

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/pkg/errors"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/server/httperr"
)

type FirebaseHttpMiddleware struct {
	AuthClient *auth.Client
}

func (a FirebaseHttpMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		bearerToken := a.tokenFromHeader(r)
		if bearerToken == "" {
			httperr.Unauthorised("empty-bearer-token", nil, w, r)
			return
		}

		token, err := a.AuthClient.VerifyIDToken(ctx, bearerToken)
		if err != nil {
			httperr.Unauthorised("unable-to-verify-jwt", err, w, r)
			return
		}

		// it's always a good idea to use custom type as context key (in this case ctxKey)
		// because nobody from the outside of the package will be able to override/read this value
		ctx = context.WithValue(ctx, userContextKey, User{
			UUID:        token.UID,
			Email:       token.Claims["email"].(string),
			Role:        token.Claims["role"].(string),
			DisplayName: token.Claims["name"].(string),
		})
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (a FirebaseHttpMiddleware) tokenFromHeader(r *http.Request) string {
	headerValue := r.Header.Get("Authorization")

	if len(headerValue) > 7 && strings.ToLower(headerValue[0:7]) == "bearer " {
		return headerValue[7:]
	}

	return ""
}

type User struct {
	UUID  string
	Email string
	Role  string

	DisplayName string
}

type ctxKey int

const (
	userContextKey ctxKey = iota
)

// TODO 以下代码存在以下问题：
//		1. 使用的是 pkg/errors 包的方法，带的堆栈信息在启动时生成，不仅没有用处，还会在使用时干扰正常的堆栈输出。
//      2. 哨兵错误值。建议改成New+Is的不透明错误风格，把值依赖变更为函数依赖。

// 哨兵错误值的问题：1. 无法携带更多的上下文信息 2. 非常量的哨兵值可能在运行时被替换 3. 后续代码迭代时，必须保证前向兼容，不能轻易调整错误内容

var (
	// if we expect that the user of the function may be interested with concrete error,
	// it's a good idea to provide variable with this error
	NoUserInContextError = errors.New("no user in context")
)

func UserFromCtx(ctx context.Context) (User, error) {
	u, ok := ctx.Value(userContextKey).(User)
	if ok {
		return u, nil
	}

	return User{}, NoUserInContextError
}
