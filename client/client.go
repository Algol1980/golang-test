package client

type Avatar struct {
	Url  string
	Size int64
}

type Client struct {
	id   int64
	name string
	age  int
	img  Avatar
}

func (c Client) HasAvatar() bool {
	if c.img.Url != "" {
		return true
	}
	return false
}

func (c *Client) UpdateAvatar() {
	c.img.Url = "https://loremflickr.com/640/480/abstract"
}

func NewClient(name string, age int, img Avatar) Client {
	return Client{
		id:   87,
		name: name,
		age:  age,
		img:  img,
	}
}
