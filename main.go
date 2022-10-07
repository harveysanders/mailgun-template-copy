package main

import (
	"context"
	"fmt"
	"log"
	"os"

	mailgun "github.com/mailgun/mailgun-go/v4"
)

var MG_API_KEY = os.Getenv("MG_API_KEY")

func main() {
	mgOldDomain := os.Getenv("MG_OLD_MAIL_DOMAIN")
	mgNewDomain := os.Getenv("MG_NEW_MAIL_DOMAIN")
	templateName := os.Getenv("MG_TEMPLATE_NAME")
	ctx := context.Background()

	err := copyTemplate(ctx, templateName, mgOldDomain, mgNewDomain)
	log.Fatal(err)
}

func copyTemplate(ctx context.Context, templateName, fromDomain, toDomain string) error {
	mgOld := mailgun.NewMailgun(fromDomain, MG_API_KEY)
	t, err := mgOld.GetTemplate(ctx, templateName)
	if err != nil {
		return fmt.Errorf("getTemplate: %v", err)
	}

	mgNew := mailgun.NewMailgun(toDomain, MG_API_KEY)

	err = mgNew.CreateTemplate(ctx, &t)
	if err != nil {
		return fmt.Errorf("createTemplate: %v", err)
	}
	return nil
}
