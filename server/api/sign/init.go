package sign

import (
	"context"
	"strings"
	"time"

	lk "github.com/digisan/logkit"
	si "github.com/digisan/user-mgr/sign-in"
	so "github.com/digisan/user-mgr/sign-out"
	su "github.com/digisan/user-mgr/sign-up"
	u "github.com/digisan/user-mgr/user"
	vf "github.com/digisan/user-mgr/user/valfield"
)

var (
	ctx    context.Context
	Cancel context.CancelFunc
)

func init() {

	// set user db dir, activate ***[UserDB]***
	u.InitDB("./data/db-user")

	// set user validator
	su.SetValidator(map[string]func(o, v any) u.ValRst{
		vf.AvatarType: func(o, v any) u.ValRst {
			ok := v == "" || strings.HasPrefix(v.(string), "image/")
			return u.NewValRst(ok, "avatarType must have prefix - 'image/'")
		},
	})

	// monitor active users
	ctx, Cancel = context.WithCancel(context.Background())
	monitorUser(ctx, 3600*time.Second) // heartbeats checker timeout
}

func monitorUser(ctx context.Context, offlineTimeout time.Duration) {
	cInactive := make(chan string, 4096)
	si.MonitorInactive(ctx, cInactive, offlineTimeout, nil)
	go func() {
		for inactive := range cInactive {
			if so.Logout(inactive) == nil {
				if user, ok := UserCache.Load(inactive); ok {
					lk.Log("deleting token: [%v]", inactive)
					user.(*u.User).DeleteToken()
					UserCache.Delete(inactive)
				}
				lk.Log("offline: [%v]", inactive)
			}
		}
	}()
}
