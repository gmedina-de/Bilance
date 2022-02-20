package server

import (
	"genuine/core/authenticator"
	"genuine/core/log"
	"github.com/emersion/go-vcard"
	"github.com/emersion/go-webdav/carddav"
	"net/http"
	"strings"
	"time"
)

func Carddav(auth authenticator.Authenticator, log log.Log) any {
	addr := "0.0.0.0:8082"

	b := &backend{make(map[string]carddav.AddressObject)}
	card := vcard.Card{}
	card.Set(vcard.FieldVersion, &vcard.Field{Value: "3.0"})
	card.AddName(&vcard.Name{
		Field:           nil,
		FamilyName:      "John",
		GivenName:       "Doe",
		AdditionalName:  "Johnny",
		HonorificPrefix: "Dr.",
		HonorificSuffix: "Master of the Universe",
	})
	card.AddAddress(&vcard.Address{
		Field:           nil,
		PostOfficeBox:   "1234",
		ExtendedAddress: "asdfasdf",
		StreetAddress:   "Down Street",
		Locality:        "London",
		Region:          "London",
		PostalCode:      "12345",
		Country:         "UK",
	})

	b.PutAddressObject("/admin", card)

	dav := &carddav.Handler{
		b,
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || !auth.Authenticate(username, password) || !strings.HasPrefix(r.URL.Path, "/"+username) {
			w.Header().Set("WWW-Basic", `Basic realm="davfs"`)
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}
		dav.ServeHTTP(w, r)
	})

	log.Info("carddav server started http://%v", addr)
	go func() {
		log.Critical(http.ListenAndServe(addr, handler).Error())
	}()
	return nil
}

type backend struct {
	contacts map[string]carddav.AddressObject
}

func (b *backend) AddressBook() (*carddav.AddressBook, error) {
	return &carddav.AddressBook{
		Path:                 "admin",
		Name:                 "TestAdressBook",
		Description:          "MyAdressBook",
		MaxResourceSize:      1000,
		SupportedAddressData: nil,
	}, nil
}

func (b *backend) GetAddressObject(path string, req *carddav.AddressDataRequest) (*carddav.AddressObject, error) {
	result := b.contacts[path]
	return &result, nil
}

func (b *backend) ListAddressObjects(req *carddav.AddressDataRequest) ([]carddav.AddressObject, error) {
	var result []carddav.AddressObject
	for _, v := range b.contacts {
		result = append(result, v)
	}
	return result, nil
}

func (b *backend) QueryAddressObjects(query *carddav.AddressBookQuery) ([]carddav.AddressObject, error) {
	return nil, nil
}

func (b *backend) PutAddressObject(path string, card vcard.Card) (loc string, err error) {
	b.contacts[path] = carddav.AddressObject{
		Path:    path,
		ModTime: time.Time{},
		ETag:    "",
		Card:    card,
	}
	return path, nil
}

func (b *backend) DeleteAddressObject(path string) error {
	delete(b.contacts, path)
	return nil
}
