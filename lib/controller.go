package lib

import (
	"encoding/json"
)

type Controller struct {
	dataAccess DataAccess
}

type BlankResult struct{}

func CreateController() Controller {
	return Controller{DataAccess{}}
}

func (c Controller) InsertTopic(topic Topic) {
	c.dataAccess.InsertTopic(topic)
}

func (c Controller) GetTopic(permalink string) []byte {
	topic := c.dataAccess.GetTopic(permalink)
	if topic.Permalink != "" {
		json, _ := json.Marshal(topic)
		return json
	} else {
		return []byte("{}")
	}
}

func (c Controller) GetTopics() []byte {
	topics := c.dataAccess.GetTopics()
	json, _ := json.Marshal(topics)
	return json
}

func (c Controller) UpdateTopic(topic Topic) {
	c.dataAccess.UpdateTopic(topic)
}
