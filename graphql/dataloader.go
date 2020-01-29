package graphql

import (
	"context"
	"net/http"
	"time"

	"github.com/confus1on/meetmeup/ent"
	"github.com/confus1on/meetmeup/ent/users"
	"github.com/confus1on/meetmeup/models"
)

const userLoaderKey = "myUserLoader"

func DataLoaderMiddleware(ctx context.Context, client *ent.Client, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userLoader := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Second,
			fetch: func(ids []int) ([]*models.User, []error) {
				var userss []*models.User

				id := make([]interface{}, len(ids))

				for i := range ids {
					id[i] = ids[i]
				}

				us, err := client.Users.
					Query().
					Where(users.IDIn(ids...)).
					All(ctx)

				if err != nil {
					return nil, []error{err}
				}

				for _, u := range us {
					user := &models.User{
						ID:       u.ID,
						Username: u.Username,
						Email:    u.Email,
					}

					userss = append(userss, user)
				}

				u := make(map[int]*models.User, len(userss))

				for _, user := range userss {
					u[user.ID] = user
				}

				result := make([]*models.User, len(ids))

				for i, id := range ids {
					result[i] = u[id]
				}

				return result, nil
			},
		}

		ctx := context.WithValue(r.Context(), userLoaderKey, &userLoader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userLoaderKey).(*UserLoader)
}
