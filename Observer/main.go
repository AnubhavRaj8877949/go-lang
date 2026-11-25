package main

import "fmt"

type Subscriber struct {
	userName string
}

func (s *Subscriber) Update(channelName string, videoTitle string) {
	fmt.Printf("Notification for %s: %s ne '%s' upload kiya hai!\n", s.userName, channelName, videoTitle)
}

type YouTubeChannel struct {
	name        string
	subscribers []*Subscriber // Yeh ek list (slice) hai jisme saare subscribers store honge
}

func (c *YouTubeChannel) subscribe(user *Subscriber) {
	c.subscribers = append(c.subscribers, user)
}
func (c *YouTubeChannel) unsubscribe(user *Subscriber) {
	for i, subscriber := range c.subscribers {
		if subscriber == user {
			c.subscribers = append(c.subscribers[:i], c.subscribers[i+1:]...)
			break
		}
	}

}

func (c *YouTubeChannel) notifyAll(videoTitle string) {

	for _, user := range c.subscribers {
		user.Update(c.name, videoTitle)
	}

}
func (c *YouTubeChannel) uploadVideo(videoTitle string) {
	fmt.Printf("\n[%s] ne naya video upload kiya: '%s'\n", c.name, videoTitle)
	c.notifyAll(videoTitle)
}

func main() {

	// create channel
	myChannel := &YouTubeChannel{name: "Code with go"}

	//Do subscriber banye
	user1 := &Subscriber{userName: "Amit"}
	user2 := &Subscriber{userName: "rohit"}

	myChannel.subscribe(user1)
	myChannel.subscribe(user2)

	myChannel.uploadVideo("Go Interfaces Explained")
	myChannel.unsubscribe(user1)
	myChannel.uploadVideo("Observer Pattern in Go")

}

/*
Kal agar tumhe:
- Analytics System ko update karna hai.
- Email Service ko update karna hai.
- Mobile Notification Service ko update karna hai.
Toh tumhe YouTubeChannel ke andar changes karne padenge.

Ye tight coupling hai.
eg.
  type YouTubeChannel struct {
    name        string
    subscribers []*Subscriber
    analytics   *Analytics
    email       *EmailService
}

func (c *YouTubeChannel) notifyAll(videoTitle string) {
    for _, user := range c.subscribers {
        user.Update(c.name, videoTitle)
    }

    c.analytics.Update(c.name, videoTitle)
    c.email.Update(c.name, videoTitle)
}




*/
