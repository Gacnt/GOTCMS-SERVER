package main

import (
	"fmt"

	"github.com/Gacnt/otcmsserver/database"
	"github.com/graphql-go/graphql"
)

var authTypeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Auth",
		Fields: graphql.Fields{
			"role_level": &graphql.Field{Type: graphql.Int},
			"token":      &graphql.Field{Type: graphql.String},
			"isLogged":   &graphql.Field{Type: graphql.Boolean},
			"name":       &graphql.Field{Type: graphql.String},
		},
	},
)

var tokenType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Auth",
		Fields: graphql.Fields{
			"role_level": &graphql.Field{Type: graphql.Int},
			"token":      &graphql.Field{Type: graphql.String},
			"name":       &graphql.Field{Type: graphql.String},
		},
	},
)

var accountType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Account",
		Fields: graphql.Fields{
			"id":        &graphql.Field{Type: graphql.ID},
			"name":      &graphql.Field{Type: graphql.String},
			"password":  &graphql.Field{Type: graphql.String},
			"secret":    &graphql.Field{Type: graphql.String},
			"type":      &graphql.Field{Type: graphql.String},
			"premdays":  &graphql.Field{Type: graphql.Int},
			"email":     &graphql.Field{Type: graphql.String},
			"creation":  &graphql.Field{Type: graphql.Int},
			"rolelevel": &graphql.Field{Type: graphql.Int},
		},
	},
)

var playerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Player",
		Fields: graphql.Fields{
			"id":                    &graphql.Field{Type: graphql.ID},
			"name":                  &graphql.Field{Type: graphql.String},
			"group_id":              &graphql.Field{Type: graphql.String},
			"account_id":            &graphql.Field{Type: graphql.String},
			"level":                 &graphql.Field{Type: graphql.Int},
			"vocation":              &graphql.Field{Type: graphql.Int},
			"health":                &graphql.Field{Type: graphql.Int},
			"healthmax":             &graphql.Field{Type: graphql.Int},
			"experience":            &graphql.Field{Type: graphql.Int},
			"lookbody":              &graphql.Field{Type: graphql.Int},
			"lookfeet":              &graphql.Field{Type: graphql.Int},
			"looklegs":              &graphql.Field{Type: graphql.Int},
			"looktype":              &graphql.Field{Type: graphql.Int},
			"lookaddons":            &graphql.Field{Type: graphql.Int},
			"maglevel":              &graphql.Field{Type: graphql.Int},
			"mana":                  &graphql.Field{Type: graphql.Int},
			"manamax":               &graphql.Field{Type: graphql.Int},
			"manaspent":             &graphql.Field{Type: graphql.Int},
			"soul":                  &graphql.Field{Type: graphql.Int},
			"town_id":               &graphql.Field{Type: graphql.Int},
			"posx":                  &graphql.Field{Type: graphql.Int},
			"posy":                  &graphql.Field{Type: graphql.Int},
			"posz":                  &graphql.Field{Type: graphql.Int},
			"cap":                   &graphql.Field{Type: graphql.Int},
			"sex":                   &graphql.Field{Type: graphql.Int},
			"lastlogin":             &graphql.Field{Type: graphql.Int},
			"blessings":             &graphql.Field{Type: graphql.Int},
			"onlinetime":            &graphql.Field{Type: graphql.Int},
			"deletion":              &graphql.Field{Type: graphql.Int},
			"balance":               &graphql.Field{Type: graphql.Int},
			"offlinetraining_time":  &graphql.Field{Type: graphql.Int},
			"offlinetraining_skill": &graphql.Field{Type: graphql.Int},
			"stamina":               &graphql.Field{Type: graphql.Int},
			"skill_fist":            &graphql.Field{Type: graphql.Int},
			"skill_fist_tries":      &graphql.Field{Type: graphql.Int},
			"skill_club":            &graphql.Field{Type: graphql.Int},
			"skill_club_tries":      &graphql.Field{Type: graphql.Int},
			"skill_sword":           &graphql.Field{Type: graphql.Int},
			"skill_sword_tries":     &graphql.Field{Type: graphql.Int},
			"skill_axe":             &graphql.Field{Type: graphql.Int},
			"skill_axe_tries":       &graphql.Field{Type: graphql.Int},
			"skill_dist":            &graphql.Field{Type: graphql.Int},
			"skill_dist_tries":      &graphql.Field{Type: graphql.Int},
			"skill_shielding":       &graphql.Field{Type: graphql.Int},
			"skill_shielding_tries": &graphql.Field{Type: graphql.Int},
			"skill_fishing":         &graphql.Field{Type: graphql.Int},
			"skill_fishing_tries":   &graphql.Field{Type: graphql.Int},
		},
	},
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"accounts": &graphql.Field{
			Type:        graphql.NewList(accountType),
			Description: "Load all player accounts",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				acc := []otdb.Account{}
				err := db.Select(&acc, "SELECT * FROM `accounts`")
				if err != nil {
					fmt.Printf("Accounts Query -- Error: %v", err)
				}

				return acc, nil
			},
		},
	},
})
