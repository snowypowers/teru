package store

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Subscription struct {
	Name     string
	Endpoint string
	Key      string
	Interval int
}

type store struct {
	Client *http.Client
	Subs   []Subscription
}

var instance *store
var once sync.Once
var subs []Subscription
var db map[string][]byte

//Store Singleton instance method.
//Returns a reference to the store object
func Store() *store {
	once.Do(func() {
		client := &http.Client{
			Timeout: time.Second * 10,
		}
		subs := make([]Subscription, 1)
		instance = &store{client, subs}
		db = make(map[string][]byte)
	})
	return instance
}

//Listens to a API endpoint, calling it at interval, returning JSON body as string through out.
//Returns a
func (a store) Listen(sub Subscription) chan string {
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

func (a store) Value(sub Subscription) []byte {
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
