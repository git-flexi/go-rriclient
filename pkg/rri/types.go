package rri

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// ContactTypePerson denotes a person.
	ContactTypePerson ContactType = "PERSON"
	// ContactTypeOrganisation denotes an organisation.
	ContactTypeOrganisation ContactType = "ORG"
	// ContactTypeRequest denotes a request contact.
	ContactTypeRequest ContactType = "REQUEST"
)

// ContactType represents the type of a contact handle.
type ContactType string

// Normalize returns the normalized representation of the given ContactType.
func (t ContactType) Normalize() ContactType {
	return ContactType(strings.ToUpper(string(t)))
}

// ParseContactType parses a contact type from string.
func ParseContactType(str string) (ContactType, error) {
	switch strings.ToUpper(str) {
	case "PERSON":
		return ContactTypePerson, nil
	case "ORG":
		return ContactTypeOrganisation, nil
	default:
		return "", fmt.Errorf("invalid contact type")
	}
}

// DenicHandle represents a handle like DENIC-1000006-SOME-CODE
type DenicHandle struct {
	RegAccID    int
	ContactCode string
}

func (h DenicHandle) String() string {
	if h.IsEmpty() {
		return ""
	}
	return fmt.Sprintf("DENIC-%d-%s", h.RegAccID, strings.ToUpper(h.ContactCode))
}

// IsEmpty returns true when the given denic handle is unset.
func (h DenicHandle) IsEmpty() bool {
	return h.RegAccID == 0 && len(h.ContactCode) == 0
}

// NewDenicHandle assembles a new denic handle.
func NewDenicHandle(regAccID int, contactCode string) DenicHandle {
	return DenicHandle{
		regAccID,
		strings.ToUpper(contactCode),
	}
}

// EmptyDenicHandle returns an empty denic handle.
func EmptyDenicHandle() DenicHandle {
	return DenicHandle{}
}

// ParseDenicHandle tries to parse a handle like DENIC-1000006-SOME-CODE. Returns an empty denic handle if str is empty.
func ParseDenicHandle(str string) (DenicHandle, error) {
	if len(str) == 0 {
		return EmptyDenicHandle(), nil
	}

	parts := strings.SplitN(str, "-", 3)
	if len(parts) != 3 {
		return DenicHandle{}, fmt.Errorf("invalid handle")
	}

	if strings.ToUpper(parts[0]) != "DENIC" {
		return DenicHandle{}, fmt.Errorf("invalid handle")
	}

	regAccID, err := strconv.Atoi(parts[1])
	if err != nil {
		return DenicHandle{}, fmt.Errorf("invalid handle")
	}

	return NewDenicHandle(regAccID, strings.ToUpper(parts[2])), nil
}

// DomainData holds domain information.
type DomainData struct {
	HolderHandles         []DenicHandle
	GeneralRequestHandles []DenicHandle
	AbuseContactHandles   []DenicHandle
	NameServers           []string
}

func (domainData *DomainData) PutToQueryFields(fields *QueryFieldList) {
	putHandlesToQueryFields := func(fieldName QueryFieldName, handles []DenicHandle) {
		for _, h := range handles {
			if !h.IsEmpty() {
				fields.Add(fieldName, h.String())
			}
		}
	}

	putHandlesToQueryFields(QueryFieldNameHolder, domainData.HolderHandles)
	putHandlesToQueryFields(QueryFieldNameGeneralRequest, domainData.GeneralRequestHandles)
	putHandlesToQueryFields(QueryFieldNameAbuseContact, domainData.AbuseContactHandles)
	fields.Add(QueryFieldNameNameServer, domainData.NameServers...)
}

// ContactData holds information of a contact handle.
type ContactData struct {
	Type         ContactType
	Name         string
	Organisation string
	Address      string
	PostalCode   string
	City         string
	CountryCode  string
	EMail        []string
	Phone        string

	VerificationInformation []VerificationInformation
}

func (contactData *ContactData) PutToQueryFields(fields *QueryFieldList) {
	fields.Add(QueryFieldNameType, string(contactData.Type.Normalize()))
	fields.Add(QueryFieldNameName, contactData.Name)
	fields.Add(QueryFieldNameOrganisation, splitLines(contactData.Organisation)...)
	fields.Add(QueryFieldNameAddress, splitLines(contactData.Address)...)
	fields.Add(QueryFieldNamePostalCode, contactData.PostalCode)
	fields.Add(QueryFieldNameCity, contactData.City)
	fields.Add(QueryFieldNameCountryCode, contactData.CountryCode)
	fields.Add(QueryFieldNameEMail, contactData.EMail...)
	fields.Add(QueryFieldNamePhone, contactData.Phone)

	for _, verificationInfo := range contactData.VerificationInformation {
		verificationInfo.PutToQueryFields(fields)
	}
}
