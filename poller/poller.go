//Package poller provides a polling service to collect data from external API endpoints at regular intervals.
package poller

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Subscription struct {
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
	Key      string `json:"key"`
	Interval int    `json:"interval"`
}

type poller struct {
	Client *http.Client
	Subs   []Subscription
}

var instance *poller
var once sync.Once
var subs []Subscription
var db map[string][]byte

//Poller Singleton instance method.
//Returns a reference to the poller object
func Poller() *poller {
	once.Do(func() {
		client := &http.Client{
			Timeout: time.Second * 10,
		}
		subs := make([]Subscription, 1)
		instance = &poller{client, subs}
		db = make(map[string][]byte)
	})
	return instance
}

//Listens to a API endpoint, calling it at interval, returning JSON body as string through out.
//Returns a
func (a poller) Listen(sub Subscription) chan string {
	subs = append(subs, sub)
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
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return quit
}

func (a poller) Value(sub Subscription) []byte {
	return db[sub.Name]
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
