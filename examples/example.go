package main

import "github.com/paul-lalonde/twitterstream"

func main() {
    stream := make(chan *twitterstream.Tweet)
    client := twitterstream.NewClient("username", "password")
    err := client.Track([]string{"miley"}, stream)
    if err != nil {
        println(err.Error())
    }
    for {
        tw := <-stream
        println(tw.User.Name, ": ", tw.Text)
    }
}
