package mod

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"regexp"
	"time"
)

const (
	Unknown = iota
	Danbooru
	Konachan
	Yandere
	Rule34
)

func GetImagesByBoard(tags string, board int64) []map[string]interface{} {
	switch board {
	case Konachan:
		return GetImagesByTags(tags, "https://konachan.com/post.json?tags=")
	case Yandere:
		return GetImagesByTags(tags, "https://yande.re/post.json?tags=")
	case Rule34:
		return GetImagesByTags(tags, "https://rule34.xxx/index.php?page=dapi&s=post&q=index&json=1&tags=")
	default:
		return GetImagesByTags(tags, "https://danbooru.donmai.us/posts.json?tags=")
	}
}

func GetImagesByTags(tags string, url string) (images []map[string]interface{}) {
	tags = regexp.MustCompile(`\s+`).ReplaceAllString(tags, "_")
	httpGet, _ := http.Get(url + tags)
	json.NewDecoder(httpGet.Body).Decode(&images)
	defer httpGet.Body.Close()
	reg := regexp.MustCompile(`(?i)\.jpg$|\.png$|\.jpeg$`)
	for i, post := range images {
		if !reg.MatchString(post["file_url"].(string)) {
			images = append(images[:i], images[i+1:]...)
		}
	}
	return images
}

func GetRandImage(tags string, board int64) map[string]interface{} {
	images := GetImagesByBoard(tags, board)
	if len(images) == 0 || images != nil {
		return make(map[string]interface{})
	}
	rand.Seed(time.Now().UnixNano())
	return images[rand.Intn(len(images))]
}
