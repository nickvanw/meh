# meh
--
    import "github.com/nickvanw/meh"

Package meh provides a simple client for the meh.com API

## Usage

#### func  WithClient

```go
func WithClient(client *http.Client) func(c *Client)
```

#### func  WithKey

```go
func WithKey(key string) func(c *Client)
```

#### func  WithURL

```go
func WithURL(urlStr string) func(c *Client)
```

#### type Client

```go
type Client struct {
}
```


#### func  NewClient

```go
func NewClient(conf ...func(c *Client)) *Client
```
NewClient returns a new Meh API Client

#### func (*Client) Current

```go
func (c *Client) Current() (*Current, error)
```
Current returns the current deal on the meh.com homepage

#### type Current

```go
type Current struct {
	Deal struct {
		Features string `json:"features"`
		ID       string `json:"id"`
		Items    []struct {
			Attributes []interface{} `json:"attributes"`
			Condition  string        `json:"condition"`
			ID         string        `json:"id"`
			Photo      string        `json:"photo"`
			Price      int           `json:"price"`
		} `json:"items"`
		Photos         []string  `json:"photos"`
		SoldOutAt      time.Time `json:"soldOutAt"`
		Specifications string    `json:"specifications"`
		Story          struct {
			Body  string `json:"body"`
			Title string `json:"title"`
		} `json:"story"`
		Theme struct {
			AccentColor     string `json:"accentColor"`
			BackgroundColor string `json:"backgroundColor"`
			BackgroundImage string `json:"backgroundImage"`
			Foreground      string `json:"foreground"`
		} `json:"theme"`
		Title string `json:"title"`
		Topic struct {
			CommentCount int       `json:"commentCount"`
			CreatedAt    time.Time `json:"createdAt"`
			ID           string    `json:"id"`
			ReplyCount   int       `json:"replyCount"`
			URL          string    `json:"url"`
			VoteCount    int       `json:"voteCount"`
		} `json:"topic"`
		URL string `json:"url"`
	} `json:"deal"`
	Poll struct {
		Answers []struct {
			ID        string `json:"id"`
			Text      string `json:"text"`
			VoteCount int    `json:"voteCount"`
		} `json:"answers"`
		ID        string    `json:"id"`
		StartDate time.Time `json:"startDate"`
		Title     string    `json:"title"`
		Topic     struct {
			CommentCount int       `json:"commentCount"`
			CreatedAt    time.Time `json:"createdAt"`
			ID           string    `json:"id"`
			ReplyCount   int       `json:"replyCount"`
			URL          string    `json:"url"`
			VoteCount    int       `json:"voteCount"`
		} `json:"topic"`
	} `json:"poll"`
	Video struct {
		ID        string    `json:"id"`
		StartDate time.Time `json:"startDate"`
		Title     string    `json:"title"`
		Topic     struct {
			CommentCount int       `json:"commentCount"`
			CreatedAt    time.Time `json:"createdAt"`
			ID           string    `json:"id"`
			ReplyCount   int       `json:"replyCount"`
			URL          string    `json:"url"`
			VoteCount    int       `json:"voteCount"`
		} `json:"topic"`
		URL string `json:"url"`
	} `json:"video"`
}
```

Current represents the current deal on the front page of meh.com

