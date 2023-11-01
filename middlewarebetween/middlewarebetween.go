package middlewarebetween

import (
	"context"
	"fmt"
	"net/http"

	helper "github.com/piyush/golangtodobackend/helpers"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientToken := r.Header.Get("token")
		if clientToken == "" {
			http.Error(w, fmt.Sprintf("No Authorization header provided"), http.StatusInternalServerError)
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			http.Error(w, err, http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), "email", claims.Email)
		ctx = context.WithValue(ctx, "first_name", claims.First_name)
		ctx = context.WithValue(ctx, "last_name", claims.Last_name)
		ctx = context.WithValue(ctx, "uid", claims.Uid)
		ctx = context.WithValue(ctx, "user_type", claims.User_type)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
