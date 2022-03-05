package notion

import (
	"context"
	"strings"

	"github.com/jomei/notionapi"
)

func UpdateRates(ctx context.Context, token, databaseID string, rates map[string]float64) error {
	client := notionapi.NewClient(notionapi.Token(token))

	db, err := client.Database.Query(ctx, notionapi.DatabaseID(databaseID), nil)
	if err != nil {
		return err
	}

	for _, page := range db.Results {
		title, ok := page.Properties["Name"].(*notionapi.TitleProperty)
		if !ok {
			continue
		}

		currency := strings.ToUpper(title.Title[0].PlainText)
		rate, ok := rates[currency]
		if !ok {
			continue
		}

		toUpdate := &notionapi.PageUpdateRequest{
			Properties: notionapi.Properties{
				"Rates": notionapi.NumberProperty{
					Number: rate,
				},
			},
		}

		id := notionapi.PageID(page.ID.String())

		client.Page.Update(ctx, id, toUpdate)
	}

	return nil
}
