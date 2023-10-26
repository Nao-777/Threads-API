package validation

import (
	"log"
	"testing"
	"threadsAPI/model"
	"threadsAPI/utility"
)

func TestUserValidate(t *testing.T){
	uv:=NewUserValidation()
	img,err:=utility.NewUtility().ImgFileEndode("../../sampleImg/tester3_big.jpg")
	if err!=nil{
		log.Fatal(err)
	}
	// type testUser struct{
	// 	LoginID   string
	// 	Name      string
	// 	Password  string
	// }
	tests:=[]struct{
		name string
		args model.User
		want string
	}{
		{
			name:"testcase-login1",
			args: model.User{
				LoginID: "loginID",
				Name: "田中青",
				Password: "Password1",
			},
			want: "",
		},
		{
			name:"testcase-login2",
			args: model.User{
				LoginID: "",
				Name: "田中青",
				Password: "Password1",
			},
			want:"login_id: loginID is required.",
		},
		{
			name:"testcase-login3",
			args: model.User{
				LoginID: "login",
				Name: "田中青",
				Password: "Password1",
			},
			want:"login_id: limited min 6 max 24.",
		},
		{
			name:"testcase-login4",
			args: model.User{
				LoginID: "loginあ",
				Name: "田中青",
				Password: "Password1",
			},
			want:"login_id: 半角英数字(大文字あり)に一致していません.",
		},
		{
			name:"testcase-Name1",
			args: model.User{
				LoginID: "loginID",
				Name: "",
				Password: "Password1",
			},
			want:"name: Name is required.",
		},
		{
			name:"testcase-Name2",
			args: model.User{
				LoginID: "loginID",
				Name: "田中青aaaaaaaaaaaaaaaaaaaaaa",
				Password: "Password1",
			},
			want:"name: limited max 24.",
		},
		{
			name:"testcase-Password1",
			args: model.User{
				LoginID: "loginID",
				Name: "田中青",
				Password: "",
			},
			want:"password: password id required.",
		},
		{
			name:"testcase-Password2",
			args: model.User{
				LoginID: "loginID",
				Name: "田中青",
				Password: "Pass1",
			},
			want:"password: limited min 8 max 24.",
		},
		{
			name:"testcase-Password3",
			args: model.User{
				LoginID: "loginID",
				Name: "田中青",
				Password: "Pass1Pass1Pass1Pass1Pass1",
			},
			want:"password: limited min 8 max 24.",
		},
		{
			name:"testcase-Password4",
			args: model.User{
				LoginID: "loginID",
				Name: "田中青",
				Password: "Pass1あPass1",
			},
			want:"password: 半角英数字(大文字あり)に一致していません.",
		},
		{
			name:"testcase-Password5",
			args: model.User{
				LoginID: "loginID",
				Name: "田中青",
				Password: "passpass",
			},
			want:"password: 大文字の英字がありません.",
		},
		{
			name:"testcase-Password6",
			args: model.User{
				LoginID: "loginID",
				Name: "田中青",
				Password: "PASSPASS",
			},
			want:"password: 小文字の英字がありません.",
		},
		{
			name:"testcase-Password7",
			args: model.User{
				LoginID: "loginID",
				Name: "田中青",
				Password: "PAsSPASS",
			},
			want:"password: 数字がありません.",
		},
		{
			name:"testcase-Image1",
			args: model.User{
				LoginID: "loginID",
				Name: "田中青",
				Password: "Password1",
				ImageUrl: img,
			},
			want: "url: 画像サイズが大きすぎます.",
		},
		{
			name:"testcase-require",
			args: model.User{
			},
			want: "",
		},
	}
	for _,tt:=range tests{
		t.Run(tt.name,func(t *testing.T){
			var got string
			err:=uv.UserValidate(tt.args,true);
			if err==nil{
				got=""
			}else{
				got=err.Error()
			}
			
			if got!=tt.want{
				t.Errorf("want:%s |now:%s\n",tt.want,got)
			}
		})
	}
}