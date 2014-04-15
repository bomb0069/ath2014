package lib 

import (
  "encoding/json"
)


type Controller struct {
  dataAccess DataAccess
}

func CreateController() Controller { 
  return Controller{ DataAccess{} } 
}

func (c Controller) GetTopics() []byte {
  topics := c.dataAccess.GetTopics() 
  json, _ := json.Marshal(topics)
  return json
}
