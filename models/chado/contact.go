package chado

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
	"gopkg.in/nullbio/null.v5"
)

// Contact is an object representing the database table.
type Contact struct {
	ContactID   int         `boil:"contact_id" json:"contact_id" toml:"contact_id" yaml:"contact_id"`
	TypeID      null.Int    `boil:"type_id" json:"type_id,omitempty" toml:"type_id" yaml:"type_id,omitempty"`
	Name        string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Description null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`

	R *contactR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L contactL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// contactR is where relationships are stored.
type contactR struct {
	Type                       *Cvterm
	ObjectContactRelationship  *ContactRelationship
	SubjectContactRelationship *ContactRelationship
	Stockcollections           StockcollectionSlice
}

// contactL is where Load methods for each relationship are stored.
type contactL struct{}

var (
	contactColumns               = []string{"contact_id", "type_id", "name", "description"}
	contactColumnsWithoutDefault = []string{"type_id", "name", "description"}
	contactColumnsWithDefault    = []string{"contact_id"}
	contactPrimaryKeyColumns     = []string{"contact_id"}
)

type (
	// ContactSlice is an alias for a slice of pointers to Contact.
	// This should generally be used opposed to []Contact.
	ContactSlice []*Contact
	// ContactHook is the signature for custom Contact hook methods
	ContactHook func(boil.Executor, *Contact) error

	contactQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	contactType                 = reflect.TypeOf(&Contact{})
	contactMapping              = queries.MakeStructMapping(contactType)
	contactPrimaryKeyMapping, _ = queries.BindMapping(contactType, contactMapping, contactPrimaryKeyColumns)
	contactInsertCacheMut       sync.RWMutex
	contactInsertCache          = make(map[string]insertCache)
	contactUpdateCacheMut       sync.RWMutex
	contactUpdateCache          = make(map[string]updateCache)
	contactUpsertCacheMut       sync.RWMutex
	contactUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var contactBeforeInsertHooks []ContactHook
var contactBeforeUpdateHooks []ContactHook
var contactBeforeDeleteHooks []ContactHook
var contactBeforeUpsertHooks []ContactHook

var contactAfterInsertHooks []ContactHook
var contactAfterSelectHooks []ContactHook
var contactAfterUpdateHooks []ContactHook
var contactAfterDeleteHooks []ContactHook
var contactAfterUpsertHooks []ContactHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Contact) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range contactBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Contact) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range contactBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Contact) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range contactBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Contact) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range contactBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Contact) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range contactAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Contact) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range contactAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Contact) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range contactAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Contact) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range contactAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Contact) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range contactAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddContactHook registers your hook function for all future operations.
func AddContactHook(hookPoint boil.HookPoint, contactHook ContactHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		contactBeforeInsertHooks = append(contactBeforeInsertHooks, contactHook)
	case boil.BeforeUpdateHook:
		contactBeforeUpdateHooks = append(contactBeforeUpdateHooks, contactHook)
	case boil.BeforeDeleteHook:
		contactBeforeDeleteHooks = append(contactBeforeDeleteHooks, contactHook)
	case boil.BeforeUpsertHook:
		contactBeforeUpsertHooks = append(contactBeforeUpsertHooks, contactHook)
	case boil.AfterInsertHook:
		contactAfterInsertHooks = append(contactAfterInsertHooks, contactHook)
	case boil.AfterSelectHook:
		contactAfterSelectHooks = append(contactAfterSelectHooks, contactHook)
	case boil.AfterUpdateHook:
		contactAfterUpdateHooks = append(contactAfterUpdateHooks, contactHook)
	case boil.AfterDeleteHook:
		contactAfterDeleteHooks = append(contactAfterDeleteHooks, contactHook)
	case boil.AfterUpsertHook:
		contactAfterUpsertHooks = append(contactAfterUpsertHooks, contactHook)
	}
}

// OneP returns a single contact record from the query, and panics on error.
func (q contactQuery) OneP() *Contact {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single contact record from the query.
func (q contactQuery) One() (*Contact, error) {
	o := &Contact{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for contact")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Contact records from the query, and panics on error.
func (q contactQuery) AllP() ContactSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Contact records from the query.
func (q contactQuery) All() (ContactSlice, error) {
	var o ContactSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to Contact slice")
	}

	if len(contactAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Contact records in the query, and panics on error.
func (q contactQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Contact records in the query.
func (q contactQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count contact rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q contactQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q contactQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if contact exists")
	}

	return count > 0, nil
}

// TypeG pointed to by the foreign key.
func (o *Contact) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Contact) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// ObjectContactRelationshipG pointed to by the foreign key.
func (o *Contact) ObjectContactRelationshipG(mods ...qm.QueryMod) contactRelationshipQuery {
	return o.ObjectContactRelationship(boil.GetDB(), mods...)
}

// ObjectContactRelationship pointed to by the foreign key.
func (o *Contact) ObjectContactRelationship(exec boil.Executor, mods ...qm.QueryMod) contactRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("object_id=$1", o.ContactID),
	}

	queryMods = append(queryMods, mods...)

	query := ContactRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"contact_relationship\"")

	return query
}

// SubjectContactRelationshipG pointed to by the foreign key.
func (o *Contact) SubjectContactRelationshipG(mods ...qm.QueryMod) contactRelationshipQuery {
	return o.SubjectContactRelationship(boil.GetDB(), mods...)
}

// SubjectContactRelationship pointed to by the foreign key.
func (o *Contact) SubjectContactRelationship(exec boil.Executor, mods ...qm.QueryMod) contactRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("subject_id=$1", o.ContactID),
	}

	queryMods = append(queryMods, mods...)

	query := ContactRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"contact_relationship\"")

	return query
}

// StockcollectionsG retrieves all the stockcollection's stockcollection.
func (o *Contact) StockcollectionsG(mods ...qm.QueryMod) stockcollectionQuery {
	return o.Stockcollections(boil.GetDB(), mods...)
}

// Stockcollections retrieves all the stockcollection's stockcollection with an executor.
func (o *Contact) Stockcollections(exec boil.Executor, mods ...qm.QueryMod) stockcollectionQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"contact_id\"=$1", o.ContactID),
	)

	query := Stockcollections(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stockcollection\" as \"a\"")
	return query
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (contactL) LoadType(e boil.Executor, singular bool, maybeContact interface{}) error {
	var slice []*Contact
	var object *Contact

	count := 1
	if singular {
		object = maybeContact.(*Contact)
	} else {
		slice = *maybeContact.(*ContactSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &contactR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &contactR{}
			args[i] = obj.TypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvterm\" where \"cvterm_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Cvterm")
	}
	defer results.Close()

	var resultSlice []*Cvterm
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Cvterm")
	}

	if len(contactAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Type = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.TypeID.Int == foreign.CvtermID {
				local.R.Type = foreign
				break
			}
		}
	}

	return nil
}

// LoadObjectContactRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (contactL) LoadObjectContactRelationship(e boil.Executor, singular bool, maybeContact interface{}) error {
	var slice []*Contact
	var object *Contact

	count := 1
	if singular {
		object = maybeContact.(*Contact)
	} else {
		slice = *maybeContact.(*ContactSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &contactR{}
		args[0] = object.ContactID
	} else {
		for i, obj := range slice {
			obj.R = &contactR{}
			args[i] = obj.ContactID
		}
	}

	query := fmt.Sprintf(
		"select * from \"contact_relationship\" where \"object_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ContactRelationship")
	}
	defer results.Close()

	var resultSlice []*ContactRelationship
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ContactRelationship")
	}

	if len(contactAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.ObjectContactRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ContactID == foreign.ObjectID {
				local.R.ObjectContactRelationship = foreign
				break
			}
		}
	}

	return nil
}

// LoadSubjectContactRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (contactL) LoadSubjectContactRelationship(e boil.Executor, singular bool, maybeContact interface{}) error {
	var slice []*Contact
	var object *Contact

	count := 1
	if singular {
		object = maybeContact.(*Contact)
	} else {
		slice = *maybeContact.(*ContactSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &contactR{}
		args[0] = object.ContactID
	} else {
		for i, obj := range slice {
			obj.R = &contactR{}
			args[i] = obj.ContactID
		}
	}

	query := fmt.Sprintf(
		"select * from \"contact_relationship\" where \"subject_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ContactRelationship")
	}
	defer results.Close()

	var resultSlice []*ContactRelationship
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ContactRelationship")
	}

	if len(contactAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.SubjectContactRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ContactID == foreign.SubjectID {
				local.R.SubjectContactRelationship = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockcollections allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (contactL) LoadStockcollections(e boil.Executor, singular bool, maybeContact interface{}) error {
	var slice []*Contact
	var object *Contact

	count := 1
	if singular {
		object = maybeContact.(*Contact)
	} else {
		slice = *maybeContact.(*ContactSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &contactR{}
		args[0] = object.ContactID
	} else {
		for i, obj := range slice {
			obj.R = &contactR{}
			args[i] = obj.ContactID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stockcollection\" where \"contact_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load stockcollection")
	}
	defer results.Close()

	var resultSlice []*Stockcollection
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice stockcollection")
	}

	if len(stockcollectionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Stockcollections = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ContactID == foreign.ContactID.Int {
				local.R.Stockcollections = append(local.R.Stockcollections, foreign)
				break
			}
		}
	}

	return nil
}

// SetType of the contact to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeContacts.
func (o *Contact) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"contact\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, contactPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.ContactID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID.Int = related.CvtermID
	o.TypeID.Valid = true

	if o.R == nil {
		o.R = &contactR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeContacts: ContactSlice{o},
		}
	} else {
		related.R.TypeContacts = append(related.R.TypeContacts, o)
	}

	return nil
}

// RemoveType relationship.
// Sets o.R.Type to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Contact) RemoveType(exec boil.Executor, related *Cvterm) error {
	var err error

	o.TypeID.Valid = false
	if err = o.Update(exec, "type_id"); err != nil {
		o.TypeID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Type = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.TypeContacts {
		if o.TypeID.Int != ri.TypeID.Int {
			continue
		}

		ln := len(related.R.TypeContacts)
		if ln > 1 && i < ln-1 {
			related.R.TypeContacts[i] = related.R.TypeContacts[ln-1]
		}
		related.R.TypeContacts = related.R.TypeContacts[:ln-1]
		break
	}
	return nil
}

// SetObjectContactRelationship of the contact to the related item.
// Sets o.R.ObjectContactRelationship to related.
// Adds o to related.R.Object.
func (o *Contact) SetObjectContactRelationship(exec boil.Executor, insert bool, related *ContactRelationship) error {
	var err error

	if insert {
		related.ObjectID = o.ContactID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"contact_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"object_id"}),
			strmangle.WhereClause("\"", "\"", 2, contactRelationshipPrimaryKeyColumns),
		)
		values := []interface{}{o.ContactID, related.ContactRelationshipID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.ObjectID = o.ContactID

	}

	if o.R == nil {
		o.R = &contactR{
			ObjectContactRelationship: related,
		}
	} else {
		o.R.ObjectContactRelationship = related
	}

	if related.R == nil {
		related.R = &contactRelationshipR{
			Object: o,
		}
	} else {
		related.R.Object = o
	}
	return nil
}

// SetSubjectContactRelationship of the contact to the related item.
// Sets o.R.SubjectContactRelationship to related.
// Adds o to related.R.Subject.
func (o *Contact) SetSubjectContactRelationship(exec boil.Executor, insert bool, related *ContactRelationship) error {
	var err error

	if insert {
		related.SubjectID = o.ContactID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"contact_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"subject_id"}),
			strmangle.WhereClause("\"", "\"", 2, contactRelationshipPrimaryKeyColumns),
		)
		values := []interface{}{o.ContactID, related.ContactRelationshipID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.SubjectID = o.ContactID

	}

	if o.R == nil {
		o.R = &contactR{
			SubjectContactRelationship: related,
		}
	} else {
		o.R.SubjectContactRelationship = related
	}

	if related.R == nil {
		related.R = &contactRelationshipR{
			Subject: o,
		}
	} else {
		related.R.Subject = o
	}
	return nil
}

// AddStockcollections adds the given related objects to the existing relationships
// of the contact, optionally inserting them as new records.
// Appends related to o.R.Stockcollections.
// Sets related.R.Contact appropriately.
func (o *Contact) AddStockcollections(exec boil.Executor, insert bool, related ...*Stockcollection) error {
	var err error
	for _, rel := range related {
		rel.ContactID.Int = o.ContactID
		rel.ContactID.Valid = true
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "contact_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &contactR{
			Stockcollections: related,
		}
	} else {
		o.R.Stockcollections = append(o.R.Stockcollections, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &stockcollectionR{
				Contact: o,
			}
		} else {
			rel.R.Contact = o
		}
	}
	return nil
}

// SetStockcollections removes all previously related items of the
// contact replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Contact's Stockcollections accordingly.
// Replaces o.R.Stockcollections with related.
// Sets related.R.Contact's Stockcollections accordingly.
func (o *Contact) SetStockcollections(exec boil.Executor, insert bool, related ...*Stockcollection) error {
	query := "update \"stockcollection\" set \"contact_id\" = null where \"contact_id\" = $1"
	values := []interface{}{o.ContactID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.Stockcollections {
			rel.ContactID.Valid = false
			if rel.R == nil {
				continue
			}

			rel.R.Contact = nil
		}

		o.R.Stockcollections = nil
	}
	return o.AddStockcollections(exec, insert, related...)
}

// RemoveStockcollections relationships from objects passed in.
// Removes related items from R.Stockcollections (uses pointer comparison, removal does not keep order)
// Sets related.R.Contact.
func (o *Contact) RemoveStockcollections(exec boil.Executor, related ...*Stockcollection) error {
	var err error
	for _, rel := range related {
		rel.ContactID.Valid = false
		if rel.R != nil {
			rel.R.Contact = nil
		}
		if err = rel.Update(exec, "contact_id"); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Stockcollections {
			if rel != ri {
				continue
			}

			ln := len(o.R.Stockcollections)
			if ln > 1 && i < ln-1 {
				o.R.Stockcollections[i] = o.R.Stockcollections[ln-1]
			}
			o.R.Stockcollections = o.R.Stockcollections[:ln-1]
			break
		}
	}

	return nil
}

// ContactsG retrieves all records.
func ContactsG(mods ...qm.QueryMod) contactQuery {
	return Contacts(boil.GetDB(), mods...)
}

// Contacts retrieves all the records using an executor.
func Contacts(exec boil.Executor, mods ...qm.QueryMod) contactQuery {
	mods = append(mods, qm.From("\"contact\""))
	return contactQuery{NewQuery(exec, mods...)}
}

// FindContactG retrieves a single record by ID.
func FindContactG(contactID int, selectCols ...string) (*Contact, error) {
	return FindContact(boil.GetDB(), contactID, selectCols...)
}

// FindContactGP retrieves a single record by ID, and panics on error.
func FindContactGP(contactID int, selectCols ...string) *Contact {
	retobj, err := FindContact(boil.GetDB(), contactID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindContact retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindContact(exec boil.Executor, contactID int, selectCols ...string) (*Contact, error) {
	contactObj := &Contact{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"contact\" where \"contact_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, contactID)

	err := q.Bind(contactObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from contact")
	}

	return contactObj, nil
}

// FindContactP retrieves a single record by ID with an executor, and panics on error.
func FindContactP(exec boil.Executor, contactID int, selectCols ...string) *Contact {
	retobj, err := FindContact(exec, contactID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Contact) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Contact) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Contact) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Contact) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no contact provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(contactColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	contactInsertCacheMut.RLock()
	cache, cached := contactInsertCache[key]
	contactInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			contactColumns,
			contactColumnsWithDefault,
			contactColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(contactType, contactMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(contactType, contactMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"contact\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

		if len(cache.retMapping) != 0 {
			cache.query += fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "chado: unable to insert into contact")
	}

	if !cached {
		contactInsertCacheMut.Lock()
		contactInsertCache[key] = cache
		contactInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Contact record. See Update for
// whitelist behavior description.
func (o *Contact) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Contact record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Contact) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Contact, and panics on error.
// See Update for whitelist behavior description.
func (o *Contact) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Contact.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Contact) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	contactUpdateCacheMut.RLock()
	cache, cached := contactUpdateCache[key]
	contactUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(contactColumns, contactPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update contact, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"contact\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, contactPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(contactType, contactMapping, append(wl, contactPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update contact row")
	}

	if !cached {
		contactUpdateCacheMut.Lock()
		contactUpdateCache[key] = cache
		contactUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q contactQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q contactQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for contact")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ContactSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o ContactSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o ContactSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ContactSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("chado: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), contactPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"contact\" SET %s WHERE (\"contact_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(contactPrimaryKeyColumns), len(colNames)+1, len(contactPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in contact slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Contact) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Contact) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Contact) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Contact) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no contact provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(contactColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	contactUpsertCacheMut.RLock()
	cache, cached := contactUpsertCache[key]
	contactUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			contactColumns,
			contactColumnsWithDefault,
			contactColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			contactColumns,
			contactPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert contact, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(contactPrimaryKeyColumns))
			copy(conflict, contactPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"contact\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(contactType, contactMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(contactType, contactMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "chado: unable to upsert for contact")
	}

	if !cached {
		contactUpsertCacheMut.Lock()
		contactUpsertCache[key] = cache
		contactUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Contact record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Contact) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Contact record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Contact) DeleteG() error {
	if o == nil {
		return errors.New("chado: no Contact provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Contact record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Contact) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Contact record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Contact) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Contact provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), contactPrimaryKeyMapping)
	sql := "DELETE FROM \"contact\" WHERE \"contact_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from contact")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q contactQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q contactQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no contactQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from contact")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o ContactSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o ContactSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no Contact slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o ContactSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ContactSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Contact slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(contactBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), contactPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"contact\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, contactPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(contactPrimaryKeyColumns), 1, len(contactPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from contact slice")
	}

	if len(contactAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Contact) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Contact) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Contact) ReloadG() error {
	if o == nil {
		return errors.New("chado: no Contact provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Contact) Reload(exec boil.Executor) error {
	ret, err := FindContact(exec, o.ContactID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ContactSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ContactSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ContactSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty ContactSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ContactSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	contacts := ContactSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), contactPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"contact\".* FROM \"contact\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, contactPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(contactPrimaryKeyColumns), 1, len(contactPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&contacts)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in ContactSlice")
	}

	*o = contacts

	return nil
}

// ContactExists checks if the Contact row exists.
func ContactExists(exec boil.Executor, contactID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"contact\" where \"contact_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, contactID)
	}

	row := exec.QueryRow(sql, contactID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if contact exists")
	}

	return exists, nil
}

// ContactExistsG checks if the Contact row exists.
func ContactExistsG(contactID int) (bool, error) {
	return ContactExists(boil.GetDB(), contactID)
}

// ContactExistsGP checks if the Contact row exists. Panics on error.
func ContactExistsGP(contactID int) bool {
	e, err := ContactExists(boil.GetDB(), contactID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ContactExistsP checks if the Contact row exists. Panics on error.
func ContactExistsP(exec boil.Executor, contactID int) bool {
	e, err := ContactExists(exec, contactID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
