package games

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type UserFinder struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func UserFinderDestructJson(name string) UserFinder {
	httpGet, _ := http.Get("https://api.mojang.com/users/profiles/minecraft/" + name)
	var finder UserFinder
	reader, _ := ioutil.ReadAll(httpGet.Body)
	json.Unmarshal(reader, &finder)
	return finder
}
