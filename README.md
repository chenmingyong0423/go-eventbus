# EventBus
Implement eventbus in go.
# Install
```go
go get github.com/chenmingyong0423/go-eventbus
```

# Usage
```go
type PostInfo struct {
    PostId string `json:"post_id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

eventBus := eventbus.NewEventBus()

subscribe := eventBus.Subscribe("post")

go func() {
    postInfo := PostInfo{
        PostId: "1",
        Title:  "Implement eventbus in Go",
        Author: "陈明勇",
    }
    bytes, err := json.Marshal(postInfo)
    if err != nil {
        panic(err)
    }
    eventBus.Publish("post", eventbus.Event{Payload: bytes})
}()

event := <-subscribe

var postInfo PostInfo
err := json.Unmarshal(event.Payload, &postInfo)
if err != nil {
    panic(err)
}
fmt.Println(postInfo)
```
