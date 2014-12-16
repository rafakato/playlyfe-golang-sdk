package playlyfe

import (
	"io/ioutil"
	"log"
	"testing"
)

func getConfig() (config PlaylyfeClientConfiguration) {
	return PlaylyfeClientConfiguration{
		ClientId:     "Zjc0MWU0N2MtODkzNS00ZWNmLWEwNmYtY2M1MGMxNGQ1YmQ4",
		ClientSecret: "YzllYTE5NDQtNDMwMC00YTdkLWFiM2MtNTg0Y2ZkOThjYTZkMGIyNWVlNDAtNGJiMC0xMWU0LWI2NGEtYjlmMmFkYTdjOTI3",
		Type:         "client",
		CacheFile:    "./token_cache.json",
	}
}

func TestClient(t *testing.T) {
	config := getConfig()

	client, err := Client(config)

	resp, err := client.http.Get("https://api.playlyfe.com/v1/game/players")

	if err != nil {
		log.Fatalln(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Printf("Testing base client result: %s", string(body))
}

func TestClientGetRaw(t *testing.T) {
	config := getConfig()

	client, err := Client(config)

	result, err := client.GetRaw(Endpoint{Url: "/game/players"})
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Printf("Testing client raw get: %s", result)
}

func TestClientGetStruct(t *testing.T) {
	config := getConfig()

	client, err := Client(config)

	players := PlayersData{}

	err = client.Get(Endpoint{Url: "/game/players"}, &players)
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Printf("Testing client structured get: %+v", players)
}

func TestClientPostRaw(t *testing.T) {
	config := getConfig()

	client, err := Client(config)

	result, err := client.PostRaw(Endpoint{Url: "/game/players", RequestBody: Player{Id: "test1", Alias: "PlayerTest"}})

	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Printf("Testing client raw post: %s", result)
}

func TestClientGetWithQueryParameters(t *testing.T) {
	config := getConfig()

	client, err := Client(config)

	result, err := client.GetRaw(Endpoint{Url: "/notifications", QueryParameters: map[string]string{
		"player_id": "test1",
	}})
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Printf("Testing client get with query parameters: %s", result)
}

func TestClientPutRaw(t *testing.T) {
	config := getConfig()

	client, err := Client(config)

	result, err := client.PutRaw(Endpoint{Url: "/game/players/test1/reset"})

	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Printf("Testing client raw put: %s", result)
}

func TestClientDeleteRaw(t *testing.T) {
	config := getConfig()

	client, err := Client(config)

	result, err := client.DeleteRaw(Endpoint{Url: "/game/players/test1"})

	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Printf("Testing client raw delete: %s", result)
}
