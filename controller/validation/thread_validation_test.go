package validation

import (
	"log"
	"testing"
	"threadsAPI/model"
	"threadsAPI/utility"
)

func TestThreadValidate(t *testing.T) {
	tv:=NewThreadValidation()
	img,err:=utility.NewUtility().ImgFileEndode("../../sampleImg/tester3_big.jpg")
	img2,err:=utility.NewUtility().ImgFileEndode("../../sampleImg/tester2.jpg")
	if err!=nil{
		log.Fatal(err)
	}
	var words string
	var words2 string
	var contents string
	var contents2 string
	for i:=0;i<100;i++{
		words+="a"
	}
	for i:=0;i<101;i++{
		words2+="a"
	}
	for i:=0;i<500;i++{
		contents+="b"
	}
	for i:=0;i<501;i++{
		contents2+="b"
	}
	log.Println(words)
	tests:=[]struct{
		name string
		args model.Thread
		want string
	}{
		{
			name :"case-pass",
			args: model.Thread{
				Title: "title",
				Contents: "contents",
				ImageUrl: "",
			},
			want: "",
		},
		{
			name :"case-title",
			args: model.Thread{
				Title: "",
				Contents: "contents",
				ImageUrl: "",
			},
			want: "title: Title is required.",
		},
		{
			name :"case-title",
			args: model.Thread{
				Title: words,
				Contents: "contents",
				ImageUrl: "",
			},
			want: "",
		},
		{
			name :"case-title",
			args: model.Thread{
				Title: words2,
				Contents: "contents",
				ImageUrl: "",
			},
			want: "title: limited min 1 max 100.",
		},
		{
			name :"case-contents",
			args: model.Thread{
				Title: "words2",
				Contents: "",
				ImageUrl: "",
			},
			want: "contents: Contents is required.",
		},
		{
			name :"case-contents",
			args: model.Thread{
				Title: "words2",
				Contents: contents,
				ImageUrl: "",
			},
			want: "",
		},
		{
			name :"case-contents",
			args: model.Thread{
				Title: "words2",
				Contents: contents2,
				ImageUrl: "",
			},
			want: "contents: limited min 1 max 500.",
		},
		{
			name :"case-img",
			args: model.Thread{
				Title: "words2",
				Contents: "contents2",
				ImageUrl: img,
			},
			want: "url: 画像サイズが大きすぎます.",
		},
		{
			name :"case-img",
			args: model.Thread{
				Title: "words2",
				Contents: "contents2",
				ImageUrl: img2,
			},
			want: "",
		},
	}
	for _,test:=range tests{
		t.Run(test.name,func(t *testing.T){
			var got string
			err:=tv.ThreadValidate(test.args);
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