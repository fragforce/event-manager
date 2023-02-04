package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/parkervcp/discord-oauth2"
	"golang.org/x/oauth2"
	"io"

	"context"
	"encoding/json"
	"log"
	"net/http"
)

var (
	state     = "random"
	oauthConf discordOauth

	discordConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/discord/callback",
		ClientID:     "353011020467404810",
		ClientSecret: "CDIqEJ6oFMghx2xY0-JPTG4Q6fo6B1KO",
		Scopes:       []string{discord.ScopeIdentify, discord.ScopeGuilds, discord.ScopeGuildsJoin},
		Endpoint:     discord.Endpoint,
	}
)

func discordRedirect(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, discordConfig.AuthCodeURL(state))
}

func discordCallback(c *gin.Context) {
	if c.Query("state") != state {
		c.String(400, "invalid state")
	}

	token, err := discordConfig.Exchange(context.Background(), c.Query("code"))
	if err != nil {
		log.Println(err)
	}

	dUser, err := getDiscordUser(token)

	if err != nil {
		log.Println(err)
	}

	getDiscordUserGuilds(token)

	c.String(200, "hello %s", dUser)
}

func getDiscordUser(token *oauth2.Token) (dUser discordUser, err error) {
	res, err := discordConfig.Client(context.Background(), token).Get("https://discord.com/api/users/@me")

	if err != nil || res.StatusCode != 200 {
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return
	}

	json.Unmarshal(body, &dUser)

	return
}

func getDiscordUserGuilds(token *oauth2.Token) (guilds []string) {
	res, err := discordConfig.Client(context.Background(), token).Get("https://discord.com/api/users/@me/guilds")

	if err != nil || res.StatusCode != 200 {
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return
	}

	log.Println(string(body))

	return
}

func joinDiscordGuild(token *oauth2.Token, guild string) (err error) {
	res, err := discordConfig.Client(context.Background(), token).Get("https://discord.com/api/users/@me/guilds")

	if err != nil || res.StatusCode != 200 {
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return
	}

	log.Println(string(body))

	return
}
