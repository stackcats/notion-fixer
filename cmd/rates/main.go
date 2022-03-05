package main

import (
	"context"
	"log"
	"os"

	"github.com/stackcats/notion-fixer/pkg/fixer"
	"github.com/stackcats/notion-fixer/pkg/notion"
)

func main() {
	fixerAccessKey := os.Getenv("FIXER_ACCESS_KEY")
	ctx := context.Background()
	rates, err := fixer.GetRates(ctx, fixerAccessKey, "usd")
	if err != nil {
		log.Fatal(err)
	}

	notionToken := os.Getenv("NOTION_TOKEN")
	notionDatabaseID := os.Getenv("NOTION_DATABASE_ID")
	err = notion.UpdateRates(ctx, notionToken, notionDatabaseID, rates)
	if err != nil {
		log.Fatal(err)
	}
}
