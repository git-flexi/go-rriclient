package rri

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"golang.org/x/net/idna"
)

// LatestVersion denotes the latest RRI version supported by the client.
const LatestVersion Version = "5.0"

const (
	QueryEntityVerificationInformation QueryFieldEntity = "VerificationInformation"
)

const (
	QueryFieldNameEntity QueryFieldName = "entity"
	// QueryFieldNameVersion denotes the query field name for version.
	QueryFieldNameVersion QueryFieldName = "version"
	// QueryFieldNameAction denotes the query field name for action.
	QueryFieldNameAction QueryFieldName = "action"
	// QueryFieldNameUser denotes the query field name for login user.
	QueryFieldNameUser QueryFieldName = "user"
	// QueryFieldNamePassword denotes the query field name for login password.
	QueryFieldNamePassword QueryFieldName = "password"
	// QueryFieldNameDomainIDN denotes the query field name for IDN domain name.
	QueryFieldNameDomainIDN QueryFieldName = "domain"
	// QueryFieldNameDomainACE denotes the query field name for ACE domain name.
	QueryFieldNameDomainACE QueryFieldName = "domain-ace"
	// QueryFieldNameHolder denotes the query field name for holder handle.
	QueryFieldNameHolder QueryFieldName = "holder"
	// QueryFieldNameGeneralRequest denotes the query field name for general request handle.
	QueryFieldNameGeneralRequest QueryFieldName = "generalrequest"
	// QueryFieldNameAbuseContact denotes the query field name for abuse contact handle.
	QueryFieldNameAbuseContact QueryFieldName = "abusecontact"
	// QueryFieldNameNameServer denotes the query field name for name servers.
	QueryFieldNameNameServer QueryFieldName = "nserver"
	// QueryFieldNameHandle denotes the query field name for denic handles.
	QueryFieldNameHandle QueryFieldName = "handle"
	// QueryFieldNameDisconnect denotes the query field name for disconnect.
	QueryFieldNameDisconnect QueryFieldName = "disconnect"
	// QueryFieldNameAuthInfoHash denotes the query field name for auth info hash.
	QueryFieldNameAuthInfoHash QueryFieldName = "authinfohash"
	// QueryFieldNameAuthInfoExpire denotes the query field name for auth info expire.
	QueryFieldNameAuthInfoExpire QueryFieldName = "authinfoexpire"
	// QueryFieldNameAuthInfo denotes the query field name for auth info hash.
	QueryFieldNameAuthInfo QueryFieldName = "authinfo"
	// QueryFieldNameType denotes the query field name for type.
	QueryFieldNameType QueryFieldName = "type"
	// QueryFieldNameName denotes the query field name for name.
	QueryFieldNameName QueryFieldName = "name"
	// QueryFieldNameOrganisation denotes the query field name for organisation.
	QueryFieldNameOrganisation QueryFieldName = "organisation"
	// QueryFieldNameAddress denotes the query field name for address.
	QueryFieldNameAddress QueryFieldName = "address"
	// QueryFieldNamePostalCode denotes the query field name for postalcode.
	QueryFieldNamePostalCode QueryFieldName = "postalcode"
	// QueryFieldNameCity denotes the query field name for city.
	QueryFieldNameCity QueryFieldName = "city"
	// QueryFieldNameCountryCode denotes the query field name for countrycode.
	QueryFieldNameCountryCode QueryFieldName = "countrycode"
	// QueryFieldNameEMail denotes the query field name for email.
	QueryFieldNameEMail QueryFieldName = "email"
	// QueryFieldNameMsgID denotes the query field name for a message id.
	QueryFieldNameMsgID QueryFieldName = "msgid"
	// QueryFieldNameMsgType denotes the query field name for a message type.
	QueryFieldNameMsgType QueryFieldName = "msgtype"
	// QueryFieldNamePhone denotes the query field name for phone.
	QueryFieldNamePhone QueryFieldName = "phone"
	// QueryFieldNameVerifiedClaim denotes the query field name for verified claim.
	QueryFieldNameVerifiedClaim QueryFieldName = "VerifiedClaim"
	// QueryFieldNameVerificationResult denotes the query field name for verification result.
	QueryFieldNameVerificationResult QueryFieldName = "VerificationResult"
	// QueryFieldNameVerificationReference denotes the query field name for verification reference.
	QueryFieldNameVerificationReference QueryFieldName = "VerificationReference"
	// QueryFieldNameVerificationTimestamp denotes the query field name for verification timestamp.
	QueryFieldNameVerificationTimestamp QueryFieldName = "VerificationTimestamp"
	// QueryFieldNameVerificationEvidence denotes the query field name for verification evidence.
	QueryFieldNameVerificationEvidence QueryFieldName = "VerificationEvidence"
	// QueryFieldNameVerificationMethod denotes the query field name for verification method.
	QueryFieldNameVerificationMethod QueryFieldName = "VerificationMethod"
	// QueryFieldNameTrustFramework denotes the query field name for trust framework.
	QueryFieldNameTrustFramework QueryFieldName = "TrustFramework"
)

const (
	// ActionLogin denotes the action value for login.
	ActionLogin QueryAction = "LOGIN"
	// ActionLogout denotes the action value for logout.
	ActionLogout QueryAction = "LOGOUT"
	// ActionCheck denotes the action value for check.
	ActionCheck QueryAction = "CHECK"
	// ActionInfo denotes the action value for info.
	ActionInfo QueryAction = "INFO"
	// ActionCreate denotes the action value for create.
	ActionCreate QueryAction = "CREATE"
	// ActionUpdate denotes the action value for update.
	ActionUpdate QueryAction = "UPDATE"
	// ActionChangeHolder denotes the action value for change holder.
	ActionChangeHolder QueryAction = "CHHOLDER"
	// ActionDelete deontes the action value for delete.
	ActionDelete QueryAction = "DELETE"
	// ActionRestore deontes the action value for restore.
	ActionRestore QueryAction = "RESTORE"
	// ActionTransit deontes the action value for transit.
	ActionTransit QueryAction = "TRANSIT"
	// ActionCreateAuthInfo1 denotes the action value for create AuthInfo1.
	ActionCreateAuthInfo1 QueryAction = "CREATE-AUTHINFO1"
	// ActionCreateAuthInfo2 denotes the action value for create AuthInfo2.
	ActionCreateAuthInfo2 QueryAction = "CREATE-AUTHINFO2"
	// ActionChangeProvider denotes the action value for change provider.
	ActionChangeProvider QueryAction = "CHPROV"
	// ActionQueueRead denotes the action value to read from the registry message queue.
	ActionQueueRead QueryAction = "QUEUE-READ"
	// ActionQueueDelete denotes the action value to delete from the registry message queue.
	ActionQueueDelete QueryAction = "QUEUE-DELETE"
)

// Version represents the RRI protocol version.
type Version string

// Normalize returns the normalized representation of the given Version.
func (v Version) Normalize() Version {
	return v
}

// QueryAction represents the action of a RRI query.
type QueryAction string

// Normalize returns the normalized representation of the given QueryAction.
func (q QueryAction) Normalize() QueryAction {
	return QueryAction(strings.ToUpper(string(q)))
}

// QueryFieldName represents a single data field of a query.
type QueryFieldName string

// Normalize returns the normalized representation of the given QueryFieldName.
func (q QueryFieldName) Normalize() QueryFieldName {
	return QueryFieldName(strings.ToLower(string(q)))
}

type QueryFieldEntity string

func (q QueryFieldEntity) String() string {
	if q == "" {
		return ""
	}
	return fmt.Sprintf("[%s]", string(q))
}

// Normalize returns the normalized representation of the given QueryFieldEntity.
func (q QueryFieldEntity) Normalize() QueryFieldEntity {
	return QueryFieldEntity(strings.ToLower(string(q)))
}

func splitLines(str string) []string {
	return strings.Split(strings.ReplaceAll(strings.ReplaceAll(str, "\r\n", "\n"), "\r", "\n"), "\n")
}

// Query represents a RRI request.
type Query struct {
	fields QueryFieldList
}

// Version returns the query version.
func (q *Query) Version() Version {
	return Version(q.FirstField(QueryFieldNameVersion)).Normalize()
}

// Action returns the query action.
func (q *Query) Action() QueryAction {
	return QueryAction(q.FirstField(QueryFieldNameAction)).Normalize()
}

// String returns a human readable representation of the query.
func (q *Query) String() string {
	var sb strings.Builder
	switch q.Action() {
	case ActionLogin:
		sb.WriteString(fmt.Sprintf("%q", q.FirstField(QueryFieldNameUser)))
	}
	return fmt.Sprintf("%s{%s}", q.Action(), sb.String())
}

// EncodeKV returns the Key-Value representation as used for RRI communication.
func (q *Query) EncodeKV() string {
	var sb strings.Builder
	for _, f := range q.fields {
		if sb.Len() > 0 {
			sb.WriteString("\n")
		}

		if f.Name != QueryFieldNameEntity {
			sb.WriteString(string(f.Name))
			sb.WriteString(": ")
		}

		sb.WriteString(f.Value)
	}
	return sb.String()
}

// Fields returns all additional response fields.
func (q *Query) Fields() QueryFieldList {
	return q.fields
}

// Field returns all values defined for a field name.
func (q *Query) Field(fieldName QueryFieldName) []string {
	return q.fields.Values(fieldName)
}

// FirstField returns the first field value or an empty string for a field name.
func (q *Query) FirstField(fieldName QueryFieldName) string {
	return q.fields.FirstValue(fieldName)
}

// NewQuery returns a query with the given parameters.
func NewQuery(version Version, action QueryAction, fields QueryFieldList) *Query {
	newFields := NewQueryFieldList()
	newFields.Add(QueryFieldNameVersion, string(version.Normalize()))
	newFields.Add(QueryFieldNameAction, string(action.Normalize()))
	if fields != nil {
		fields.CopyTo(&newFields)
	}
	return &Query{newFields}
}

// NewLoginQuery returns a login query for the given credentials.
func NewLoginQuery(username, password string) *Query {
	fields := NewQueryFieldList()
	fields.Add(QueryFieldNameUser, username)
	fields.Add(QueryFieldNamePassword, password)
	return NewQuery(LatestVersion, ActionLogin, fields)
}

// NewLogoutQuery returns a logout query.
func NewLogoutQuery() *Query {
	return NewQuery(LatestVersion, ActionLogout, nil)
}

// NewCreateContactQuery returns a check query.
func NewCreateContactQuery(handle DenicHandle, contactData ContactData) *Query {
	fields := NewQueryFieldList()
	fields.Add(QueryFieldNameHandle, handle.String())
	contactData.PutToQueryFields(&fields)
	return NewQuery(LatestVersion, ActionCreate, fields)
}

// NewCheckHandleQuery returns a check query for a contact or request contact handle.
func NewCheckHandleQuery(handle DenicHandle) *Query {
	fields := NewQueryFieldList()
	fields.Add(QueryFieldNameHandle, handle.String())
	return NewQuery(LatestVersion, ActionCheck, fields)
}

// NewInfoHandleQuery returns an info query for a contact or request contact handle.
func NewInfoHandleQuery(handle DenicHandle) *Query {
	fields := NewQueryFieldList()
	fields.Add(QueryFieldNameHandle, handle.String())
	return NewQuery(LatestVersion, ActionInfo, fields)
}

func PutDomainToQueryFields(fields *QueryFieldList, domain string) {
	if strings.HasPrefix(strings.ToLower(domain), "xn--") {
		fields.Add(QueryFieldNameDomainACE, domain)
		if idn, err := idna.ToUnicode(domain); err == nil {
			fields.Add(QueryFieldNameDomainIDN, idn)
		}

	} else {
		fields.Add(QueryFieldNameDomainIDN, domain)
		if ace, err := idna.ToASCII(domain); err == nil {
			// TODO only add ace string if it differs from idn
			fields.Add(QueryFieldNameDomainACE, ace)
		}
	}
}

// NewCreateDomainQuery returns a query to create a domain.
func NewCreateDomainQuery(domain string, domainData DomainData) *Query {
	fields := NewQueryFieldList()
	PutDomainToQueryFields(&fields, domain)
	domainData.PutToQueryFields(&fields)
	return NewQuery(LatestVersion, ActionCreate, fields)
}

// NewCheckDomainQuery returns a check query.
func NewCheckDomainQuery(domain string) *Query {
	fields := NewQueryFieldList()
	PutDomainToQueryFields(&fields, domain)
	return NewQuery(LatestVersion, ActionCheck, fields)
}

// NewInfoDomainQuery returns an info query.
func NewInfoDomainQuery(domain string) *Query {
	fields := NewQueryFieldList()
	PutDomainToQueryFields(&fields, domain)
	return NewQuery(LatestVersion, ActionInfo, fields)
}

// NewUpdateDomainQuery returns a query to update a domain.
func NewUpdateDomainQuery(domain string, domainData DomainData) *Query {
	fields := NewQueryFieldList()
	PutDomainToQueryFields(&fields, domain)
	domainData.PutToQueryFields(&fields)
	return NewQuery(LatestVersion, ActionUpdate, fields)
}

// NewChangeHolderQuery returns a query to update a domain.
func NewChangeHolderQuery(domain string, domainData DomainData) *Query {
	fields := NewQueryFieldList()
	PutDomainToQueryFields(&fields, domain)
	domainData.PutToQueryFields(&fields)
	return NewQuery(LatestVersion, ActionChangeHolder, fields)
}

// NewDeleteDomainQuery returns a delete query.
func NewDeleteDomainQuery(domain string) *Query {
	fields := NewQueryFieldList()
	PutDomainToQueryFields(&fields, domain)
	return NewQuery(LatestVersion, ActionDelete, fields)
}

// NewRestoreDomainQuery returns a restore query.
func NewRestoreDomainQuery(domain string) *Query {
	fields := NewQueryFieldList()
	PutDomainToQueryFields(&fields, domain)
	return NewQuery(LatestVersion, ActionRestore, fields)
}

// NewTransitDomainQuery returns a restore query.
func NewTransitDomainQuery(domain string, disconnect bool) *Query {
	fields := NewQueryFieldList()
	PutDomainToQueryFields(&fields, domain)
	if disconnect {
		fields.Add(QueryFieldNameDisconnect, "true")
	} else {
		fields.Add(QueryFieldNameDisconnect, "false")
	}
	return NewQuery(LatestVersion, ActionTransit, fields)
}

// NewCreateAuthInfo1Query returns a create AuthInfo1 query.
func NewCreateAuthInfo1Query(domain, authInfo string, expireDay time.Time) *Query {
	fields := NewQueryFieldList()
	PutDomainToQueryFields(&fields, domain)
	fields.Add(QueryFieldNameAuthInfoHash, computeHashSHA256(authInfo))
	fields.Add(QueryFieldNameAuthInfoExpire, expireDay.Format("20060102"))
	return NewQuery(LatestVersion, ActionCreateAuthInfo1, fields)
}

func computeHashSHA256(str string) string {
	hasher := sha256.New()
	hasher.Write([]byte(str))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

// NewCreateAuthInfo2Query returns a create AuthInfo2 query.
func NewCreateAuthInfo2Query(domain string) *Query {
	fields := NewQueryFieldList()
	PutDomainToQueryFields(&fields, domain)
	return NewQuery(LatestVersion, ActionCreateAuthInfo2, fields)
}

// NewChangeProviderQuery returns a query to create a domain.
func NewChangeProviderQuery(domain, authInfo string, domainData DomainData) *Query {
	fields := NewQueryFieldList()
	PutDomainToQueryFields(&fields, domain)
	domainData.PutToQueryFields(&fields)
	fields.Add(QueryFieldNameAuthInfo, authInfo)
	return NewQuery(LatestVersion, ActionChangeProvider, fields)
}

// NewQueueReadQuery returns a query to read from the registry message queue. Use msgType to filter for specific message types or use an empty string to process all message types.
func NewQueueReadQuery(msgType string) *Query {
	fields := NewQueryFieldList()
	if len(msgType) > 0 {
		fields.Add(QueryFieldNameMsgType, msgType)
	}
	return NewQuery(LatestVersion, ActionQueueRead, fields)
}

// NewQueueReadQuery returns a query to read from the registry message queue. Use msgType to delete only specific message types or use an empty string to process all message types. This is required if you want to delete the oldest message of a specific type that is not the oldest in your full queue.
func NewQueueDeleteQuery(msgID, msgType string) *Query {
	fields := NewQueryFieldList()
	fields.Add(QueryFieldNameMsgID, msgID)
	if len(msgType) > 0 {
		fields.Add(QueryFieldNameMsgType, msgType)
	}
	return NewQuery(LatestVersion, ActionQueueDelete, fields)
}

// ParseQueryKV parses a single key-value encoded query.
func ParseQueryKV(str string) (*Query, error) {
	lines := strings.Split(str, "\n")
	fields := NewQueryFieldList()
	for _, line := range lines {
		// trim spaces and ignore empty lines
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("query line must be key-value separated by ':'")
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		fields.Add(QueryFieldName(key), value)
	}

	versionValues := fields.Values(QueryFieldNameVersion)
	if len(versionValues) == 0 {
		return nil, fmt.Errorf("%s key is missing", QueryFieldNameVersion)
	}
	if len(versionValues) > 1 {
		return nil, fmt.Errorf("multiple %s values", QueryFieldNameVersion)
	}

	actionValues := fields.Values(QueryFieldNameAction)
	if len(actionValues) == 0 {
		return nil, fmt.Errorf("%s key is missing", QueryFieldNameAction)
	}
	if len(actionValues) > 1 {
		return nil, fmt.Errorf("multiple %s values", QueryFieldNameAction)
	}

	return &Query{fields}, nil
}

// ParseQuery tries to detect the query format (KV or XML) and returns the parsed query.
func ParseQuery(str string) (*Query, error) {
	// TODO detect type
	return ParseQueryKV(str)
}
