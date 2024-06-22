package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/badoux/checkmail"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/davidmytton/url-verifier"
	"github.com/synapsecns/sanguine/contrib/opbot/signoz"
	"github.com/synapsecns/sanguine/contrib/opbot/signoz/example/keychain"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"os"
)

func main() {
	client := loadConfig()

	resp, err := client.SearchTraces(context.Background(), signoz.Last1Hr, map[string]string{})
	if err != nil {
		panic(err)
	}

	_ = resp
}

const (
	// ConfigPath for plaintext config files.
	ConfigPath = "~/.signoz/"
	// EmailPath path to email file.
	EmailPath = ConfigPath + "email"
	// URLPath path to url file.
	URLPath = ConfigPath + "url"
	// KeychainServiceName (service name for key chain: ios only).
	KeychainServiceName = "signoz-example"
)

type config struct {
	email    string
	url      string
	password string
}

func loadConfig() *signoz.Client {
	cfg := config{}

	var configFormFields []huh.Field

	if cfg.url = readPath(URLPath); cfg.url == "" {
		configFormFields = append(configFormFields, huh.NewInput().
			Value(&cfg.url).
			Title("URL").
			Validate(func(s string) error {
				verifier := urlverifier.NewVerifier()
				if _, err := verifier.Verify(s); err != nil {
					return fmt.Errorf("invalid URL: %w", err)
				}
				return nil
			}).
			Placeholder("URL Path"))
	}

	if cfg.email = readPath(EmailPath); cfg.email == "" {
		configFormFields = append(configFormFields, huh.NewInput().
			Value(&cfg.email).
			Title("Email").
			Validate(func(s string) error {
				err := checkmail.ValidateFormat(s)
				if err != nil {
					return fmt.Errorf("invalid email address: %w", err)
				}
				return nil
			}).
			Placeholder("Email Address"))
	}

	var err error
	cfg.password, err = getPassword()
	if err != nil {
		panic(err)
	}

	if cfg.password == "" {
		configFormFields = append(configFormFields, huh.NewInput().
			Value(&cfg.password).
			EchoMode(huh.EchoModePassword).
			Title("Password").
			Placeholder("Password"))
	}

	// if either key is not set, prepend the note to the form
	if len(configFormFields) > 0 {
		configFormFields = append([]huh.Field{
			huh.NewNote().
				Title("Setup Signoz Configuration").
				Next(true).
				Description("Signoz requires an API key to authenticate with the server. You can find your API key in the Signoz dashboard.")}, configFormFields...)

		group := huh.NewGroup(configFormFields...)
		err = huh.NewForm(group).Run()
		if err != nil {
			panic(err)
		}
	}

	// try to login, if it works, save creds otherwise panic!
	var clientErr error
	var client *signoz.Client

	err = spinner.New().
		Title("Logging in...").
		Action(func() {
			client = signoz.NewClientFromUser(metrics.NewNullHandler(), cfg.url, cfg.email, cfg.password)

			_, err := client.Services(context.Background(), signoz.Last1Hr)
			if err != nil {
				clientErr = fmt.Errorf("could not login: %w", err)
			}
		}).Run()
	if err != nil {
		panic(err)
	}

	if clientErr != nil {
		panic(clientErr)
	}

	storeConfig(cfg)

	return client
}

func storeConfig(cfg config) {
	_ = os.MkdirAll(core.ExpandOrReturnPath(ConfigPath), 0700)
	err := os.WriteFile(core.ExpandOrReturnPath(EmailPath), []byte(cfg.email), 0600)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(core.ExpandOrReturnPath(URLPath), []byte(cfg.url), 0600)
	if err != nil {
		panic(err)
	}

	_, err = getPassword()
	if err != nil {
		item := keychain.NewGenericPassword(KeychainServiceName, KeychainServiceName, "signoz password", []byte(cfg.password), KeychainServiceName)
		err := keychain.AddItem(item)
		if err != nil {
			panic(err)
		}
	}
}

// readPath reads a path.
func readPath(path string) string {
	fileBytes, err := os.ReadFile(core.ExpandOrReturnPath(path))
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		panic(err)
	}

	return string(fileBytes)
}

func getPassword() (string, error) {
	res, err := keychain.GetGenericPassword(KeychainServiceName, KeychainServiceName, "signoz password", KeychainServiceName)
	return string(res), err
}
