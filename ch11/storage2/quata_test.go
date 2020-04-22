package storage2

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	// 保留留待恢复的 notifyUser
	saved := notifyUser
	defer func() {
		notifyUser = saved
	}()

	// 设置测试的伪通知 notifyUser
	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}

	// 模拟已经使用 980MB 的情况
	const user = "joe@example.org"
	usage[user] = 980000000 // simulate a 980MB-used condition
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}
	const wantSubstring = "%98 of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, want substring %q", notifiedMsg, notifiedUser)
	}
}
