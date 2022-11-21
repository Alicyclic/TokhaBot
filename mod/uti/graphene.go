package uti

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type border struct {
	Thickness uint8  `json:"thickness"`
	Colour    string `json:"colour"`
}

type graphene struct {
	Code       string `json:"code"`
	Theme      string `json:"theme"`
	Format     string `json:"format"`
	Lang       string `json:"lang"`
	Upscale    uint8  `json:"upscale"`
	Font       string `json:"font"`
	LineNumber bool   `json:"lineNumber"`
	Border     border `json:"border"`
}

func NewGraphene() *graphene {
	return &graphene{
		Code:       "hello world",
		Theme:      "one-dark-pro",
		Format:     "png",
		LineNumber: true,
		Upscale:    5,
		Font:       "iosevka",
		Border: border{
			Thickness: 30,
			Colour:    "#d1d1d1",
		},
	}
}

func (g *graphene) SetCode(c string) {
	g.Code = c
}

func ReturnIfErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func (g graphene) Request() (img []byte, err error) {
	json_data, err := json.Marshal(g)
	ReturnIfErr(err)
	res, err := http.Post("https://teknologi-umum-graphene.fly.dev/api", "application.json", bytes.NewBuffer(json_data))
	ReturnIfErr(err)
	defer res.Body.Close()
	img, err = ioutil.ReadAll(res.Body)
	ReturnIfErr(err)
	return
}
