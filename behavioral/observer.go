package behavioral

import "fmt"

// The example is a channel that everytime a new
// news are added to it, its subscribers get notified

type User struct {
	name string
}

type UserI interface {
	getNotification(channelName string)
}

func (u User) getNotification(channelName string) {
	fmt.Printf("Hey %s! Channel %s has posted a new video\n", u.name, channelName)
}

type Channel struct {
	name            string
	registeredUsers []User
}

func newChannel(name string) *Channel {
	return &Channel{
		name: name,
	}
}

type ChannelI interface {
	register(user User)
	uregister(user User)
	notifyUsers()
}

func (c *Channel) register(u User) {
	c.registeredUsers = append(c.registeredUsers, u)
	fmt.Printf("User %s registered to the channel %s\n", u.name, c.name)
}

func (c *Channel) unregister(u User) {
	for i, user := range c.registeredUsers {
		if user.name == u.name {
			c.registeredUsers = append(c.registeredUsers[:i], c.registeredUsers[i+1])
			break
		}
	}

	fmt.Printf("User %s unregistered from the channel %s\n", u.name, c.name)
}

func (c *Channel) notifyUsers() {
	for _, u := range c.registeredUsers {
		u.getNotification(c.name)
	}
}

func (c *Channel) newEvent() {
	fmt.Println()
	fmt.Println("A new video has been posted, notifying all users in this channel")
	fmt.Println()
	c.notifyUsers()
}

func Observer() {
	channel := newChannel("Funny channel!")

	var users = []*User{
		{name: "Gabriel"},
		{name: "Junior"},
		{name: "Cleber"},
	}

	for _, u := range users {
		channel.register(*u)
	}

	channel.newEvent()
}
