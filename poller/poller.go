//Package poller provides a polling service to collect data from external API endpoints at regular intervals.
//Poller also serves as a cache for the latest data.
package poller

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

//Subscription is a struct that poller accepts to start listening to an external endpoint
type Subscription struct {
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
	Key      string `json:"key"`
	Interval int    `json:"interval"`
}

type Poller struct {
	Client *http.Client
	Subs   []Subscription
}

var instance *Poller
var once sync.Once
var subscribers map[string][]func([]byte)
var db map[string][]byte

//Poller Singleton instance method.
//Returns a reference to the poller object
func NewPoller( c *http.Client) *Poller {
	once.Do(func() {
		endpoints := make([]Subscription, 1)
		instance = &Poller{c, endpoints}
		db = make(map[string][]byte)
		subscribers = make(map[string][]func([]byte))
	})
	return instance
}

//Listens to a API endpoint, calling it at interval, returning JSON body as string through out.
//Returns a
func (a Poller) Listen(sub Subscription) chan string {
	instance.Subs = append(instance.Subs, sub)
	quit := make(chan string)
	ticker := time.NewTicker(time.Second * time.Duration(sub.Interval))
	//Initial retrieval
	res := retrieve(sub)
	db[sub.Name] = res
	go func() {
		for {
			select {
			case <-ticker.C:
				res := retrieve(sub)
				db[sub.Name] = res
				subs := subscribers[sub.Name]
				for _,i := range subs {
					i(res)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return quit
}

//Subscribes to a update using a name.
func (a Poller) Subscribe(name string, f func([]byte)) {
	subscribers[name] = append(subscribers[name], f)
}

//Returns the last value of the subscription
func (a Poller) Value(sub Subscription) []byte {
	return db[sub.Name]
}

//Same as Value but accepts a string
func (a Poller) ValueByName(name string) []byte {
	return db[name]
}

//Retrieve immediately sends an API request and attempts to retrieve the information available.
//Returns the JSON body in string format.
func retrieve(sub Subscription) []byte {
	req, err := http.NewRequest("GET", sub.Endpoint, nil)
	req.Header.Add("api-key", sub.Key)
	res, err := instance.Client.Do(req)

	if err != nil {
		log.Fatal(err)
		return nil
	} else {
		defer res.Body.Close()
		bs, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return bs
	}
}
