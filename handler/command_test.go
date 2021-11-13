package handler

import (
	"encoding/xml"
	"fmt"
	"runrun_uncle/model"
	"testing"
)

var command = "<xml><ToUserName><![CDATA[gh_b5f45f5985e6]]></ToUserName>\n<FromUserName><![CDATA[oiNVetw1hs1HIyp_FrVkzA7xT93U]]></FromUserName>\n<CreateTime>1636807903</CreateTime>\n<MsgType><![CDATA[text]]></MsgType>\n<Content><![CDATA[hello_world]]></Content>\n<MsgId>23433453374041788</MsgId>\n</xml>"

func TestMarshalUnmarshalXaml(t *testing.T) {
	x := model.MsgStruct{}
	err := xml.Unmarshal([]byte(command), &x)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(x)
		fmt.Println(x.Content)
	}

	bytes, _ := xml.Marshal(x)
	fmt.Println(string(bytes))

}
