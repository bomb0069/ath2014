package lib

import (
	"encoding/json"
)

type Controller struct {
	dataAccess DataAccess
}


func CreateController() Controller {
	return Controller{DataAccess{}}
}

func (c Controller) InsertTopic(topic Topic) {
	c.dataAccess.InsertTopic(topic)
}

func (c Controller) GetTopics() []byte {
	topics := c.dataAccess.GetTopics() 
	json, _ := json.Marshal(topics)
	return json
}

func (c Controller) UpdateTopic(topic Topic) {
	c.dataAccess.UpdateTopic(topic)
}
