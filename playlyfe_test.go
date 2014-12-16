package playlyfe

import (
	"fmt"
	"io/ioutil"

	"net/http"
	"os"
	"testing"
)

func getConfigClient() (config PlaylyfeClientConfiguration) {
	return PlaylyfeClientConfiguration{
		ClientId:     "Zjc0MWU0N2MtODkzNS00ZWNmLWEwNmYtY2M1MGMxNGQ1YmQ4",
		ClientSecret: "YzllYTE5NDQtNDMwMC00YTdkLWFiM2MtNTg0Y2ZkOThjYTZkMGIyNWVlNDAtNGJiMC0xMWU0LWI2NGEtYjlmMmFkYTdjOTI3",
		Type:         "client",
		CacheFile:    "./token_cache.json",
	}
}

func getConfigCode() (config PlaylyfeClientConfiguration) {
	return PlaylyfeClientConfiguration{
		ClientId:     "Zjc0MWU0N2MtODkzNS00ZWNmLWEwNmYtY2M1MGMxNGQ1YmQ4",
		ClientSecret: "YzllYTE5NDQtNDMwMC00YTdkLWFiM2MtNTg0Y2ZkOThjYTZkMGIyNWVlNDAtNGJiMC0xMWU0LWI2NGEtYjlmMmFkYTdjOTI3",
		Type:         "code",
		RedirectUrl:  "http://localhost:3000/code",
		CacheFile:    "./token_cache.json",
	}
}

func startWebServer(port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	})
	http.ListenAndServe(":"+string(port), nil)
}

func TestClientTypeCode(t *testing.T) {
	startWebServer(3000)
	config := getConfigCode()

	os.Remove("./token_cache.json")
	client, err := Client(config)
	if err != nil {
		t.Fatal(err)
		return
	}

	resp, err := client.http.Get("https://api.playlyfe.com/v1/game/players")

	if err != nil {
		t.Fatal(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Logf("Testing base client result: %s", string(body))
}

/*func TestClientTypeClient(t *testing.T) {
	config := getConfigClient()

	os.Remove("./token_cache.json")
	client, err := Client(config)

	resp, err := client.http.Get("https://api.playlyfe.com/v1/game/players")

	if err != nil {
		t.Fatal(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Logf("Testing base client result: %s", string(body))
}

func TestClientGetRaw(t *testing.T) {
	config := getConfigClient()

	os.Remove("./token_cache.json")
	client, err := Client(config)

	result, err := client.GetRaw(Endpoint{Url: "/game/players"})
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Logf("Testing client raw get: %s", result)
}

func TestClientGetStruct(t *testing.T) {
	config := getConfigClient()

	os.Remove("./token_cache.json")
	client, err := Client(config)

	players := PlayersData{}

	err = client.Get(Endpoint{Url: "/game/players"}, &players)
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Logf("Testing client structured get: %+v", players)
}

func TestClientPostRaw(t *testing.T) {
	config := getConfigClient()

	os.Remove("./token_cache.json")
	client, err := Client(config)

	result, err := client.PostRaw(Endpoint{Url: "/game/players", RequestBody: Player{Id: "test1", Alias: "PlayerTest"}})

	if err != nil {
		t.Fatal(err)
		return
	}

	t.Logf("Testing client raw post: %s", result)
}

func TestClientGetWithQueryParameters(t *testing.T) {
	config := getConfigClient()

	os.Remove("./token_cache.json")
	client, err := Client(config)

	result, err := client.GetRaw(Endpoint{Url: "/notifications", QueryParameters: map[string]string{
		"player_id": "test1",
	}})
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Logf("Testing client get with query parameters: %s", result)
}

func TestClientPutRaw(t *testing.T) {
	config := getConfigClient()

	os.Remove("./token_cache.json")
	client, err := Client(config)

	result, err := client.PutRaw(Endpoint{Url: "/game/players/test1/reset"})

	if err != nil {
		t.Fatal(err)
		return
	}

	t.Logf("Testing client raw put: %s", result)
}

func TestClientDeleteRaw(t *testing.T) {
	config := getConfigClient()

	os.Remove("./token_cache.json")
	client, err := Client(config)

	result, err := client.DeleteRaw(Endpoint{Url: "/game/players/test1"})

	if err != nil {
		t.Fatal(err)
		return
	}

	t.Logf("Testing client raw delete: %s", result)
}
*/
