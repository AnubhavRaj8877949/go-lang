package main

import "fmt"

// 1. INTERFACE (Common Rule)
type Observer interface {
	Update(channelName string, videoTitle string)
}

// Observer 1: Normal User
type Subscriber struct {
	userName string
}

func (s *Subscriber) Update(channelName string, videoTitle string) {
	fmt.Printf("Notification for %s: %s ne '%s' upload kiya hai!\n", s.userName, channelName, videoTitle)
}

// Observer 2: Analytics System (Naya feature bina YouTubeChannel ko chhede!)
type AnalyticsSystem struct{}

func (a *AnalyticsSystem) Update(channelName string, videoTitle string) {
	fmt.Printf("[Analytics] Database mein '%s' ka tracking shuru ho gaya.\n", videoTitle)
}

// Observer 3: Email Service (Ek aur naya feature!)
type EmailService struct {
	adminEmail string
}

func (e *EmailService) Update(channelName string, videoTitle string) {
	fmt.Printf("[Email] %s par mail bhej diya gaya hai.\n", e.adminEmail)
}

//----------------------------------------------------

type YouTubeChannel struct {
	name        string
	subscribers []Observer // Yeh ek list (slice) hai jisme saare subscribers store honge
}

func (c *YouTubeChannel) subscribe(o Observer) {
	c.subscribers = append(c.subscribers, o)
}
func (c *YouTubeChannel) unsubscribe(o Observer) {
	for i, subscriber := range c.subscribers {
		if subscriber == o {
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

	analytics := &AnalyticsSystem{}
	email := &EmailService{adminEmail: "admin@youtube.com"}

	myChannel.subscribe(user1)
	myChannel.subscribe(user2)

	myChannel.subscribe(analytics)
	myChannel.subscribe(email)

	myChannel.uploadVideo("Go Interfaces Explained")
	myChannel.unsubscribe(user1)
	myChannel.uploadVideo("Observer Pattern in Go")
}
