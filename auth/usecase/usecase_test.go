package authUsecase

import (
	"fmt"
	"testing"
	"time"
)


func TestGenerateAuthToken(t *testing.T){

	a := NewAuthUseCase(nil, "test-salt", []byte("test-signing-key"), 5 * time.Second)
	test_user_id := "id-user-test"
	token, err := a.GenerateAuthToken(test_user_id)
	if err != nil || token == ""{
		t.Error(err)
		t.Fail()
	}
	
	fmt.Println(token)
}


func TestParseToken(t *testing.T){

	a := NewAuthUseCase(nil, "test-salt", []byte("test-signing-key"), 5 * time.Second)
	test_user_id := "id-user-test"
	token, err := a.GenerateAuthToken(test_user_id)

	user_id, err := a.ParseToken(token)
	if err != nil{
		t.Fail()
	}
	fmt.Println(user_id)
	if user_id != "id-user-test"{
		t.Fail()
	}
}