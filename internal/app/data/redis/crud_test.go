package redis

import (
	"context"
	"testing"
)

// go test -v oauth-server/redis
// go test oauth-server/redis -v -run TestRedisCrud
func TestRedisCrud(t *testing.T) {
	// 해당 패키지의 init() 함수를 통하여 전달됩니다.
	if Client == nil {
		t.Error("redis client가 할당되지 않았습니다.")
		return
	}

	ctx := context.Background()

	key := "redisCrudTest"
	createValue := "테스트 값 입니다"
	editValue := "수정된 값 입니다"

	// 생성
	err := Client.Set(ctx, key, createValue, 0).Err()
	if err != nil {
		t.Errorf("redis 생성에 실패하였습니다: " + err.Error())
		return
	}

	// 읽기
	val, err := Client.Get(ctx, key).Result()
	if err != nil {
		t.Errorf("redis 읽기에 실패하였습니다: " + err.Error())
		return
	}
	if val != createValue {
		t.Errorf("redis 생성 값과 읽은 값이 다릅니다.\n")
		t.Errorf("\t읽은 값: %s\n", val)
		t.Errorf("\t생성 값: %s\n", editValue)
		return
	}

	// 수정
	err = Client.Set(ctx, key, editValue, 0).Err()
	if err != nil {
		t.Errorf("redis 수정에 실패하였습니다: " + err.Error())
		return
	}
	val, err = Client.Get(ctx, key).Result()
	if err != nil {
		t.Errorf("redis 읽기에 실패하였습니다: " + err.Error())
		return
	}
	if val != editValue {
		t.Errorf("redis 수정 값이 다릅니다.\n")
		t.Errorf("\t읽은 값: %s\n", val)
		t.Errorf("\t수정 값: %s\n", editValue)
		return
	}

	// 삭제
	err = Client.Del(ctx, key).Err()
	if err != nil {
		t.Error("redis 삭제에 실패하였습니다.", err.Error())
		return
	}
}
