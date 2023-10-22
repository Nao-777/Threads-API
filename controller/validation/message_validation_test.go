package validation

import (
	"log"
	"testing"
	"threadsAPI/model"
	"threadsAPI/utility"
)

func TestMessageValidate(t *testing.T) {
	mv:=NewMessageValidation()
	img,err:=utility.NewUtility().ImgFileEndode("../../sampleImg/tester3_big.jpg")
	if err!=nil{
		log.Fatal(err)
	}
	img2,err:=utility.NewUtility().ImgFileEndode("../../sampleImg/tester2.jpg")
	if err!=nil{
		log.Fatal(err)
	}
	var words string
	var words2 string
	for i:=0;i<500;i++{
		words+="a"
	}
	for i:=0;i<501;i++{
		words2+="a"
	}
	tests:=[]struct{
		name string
		args model.Message
		want string
	}{
		{
			name:"case-pass",
			args:model.Message{
				Message: words,
				ImageUrl: img2,
			},
			want: "",
		},
		{
			name:"case-msg",
			args:model.Message{
				Message: "",
				ImageUrl: "",
			},
			want: "message: message is required.",
		},
		{
			name:"case-msg",
			args:model.Message{
				Message: words2,
				ImageUrl: "",
			},
			want: "message: limited min 1 max 500.",
		},
		{
			name:"case-img",
			args:model.Message{
				Message: words,
				ImageUrl: img2,
			},
			want: "",
		},
		{
			name:"case-img2",
			args:model.Message{
				Message: words,
				ImageUrl: img,
			},
			want: "url: 画像サイズが大きすぎます.",
		},
		{
			name:"case-img",
			args:model.Message{
				Message: words,
				ImageUrl: "",
			},
			want: "",
		},
	}
	for _,test:=range tests{
		t.Run(test.name,func(t *testing.T){
			var got string
			err:=mv.MessageValidate(test.args);
			if err==nil{
				got=""
			}else{
				got=err.Error()
			}
			
			if got!=test.want{
				t.Errorf("want:%s |now:%s\n",test.want,got)
			}
		})
	}
}