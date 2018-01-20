package parsehtml

import (
	"testing"
	"fmt"
)

func TestGetUsersId(t *testing.T) {
	uids := GetUsersId("https://www.jianshu.com/recommendations/users?page=1")
	fmt.Println(len(uids))
	fmt.Println(uids[0])
}
