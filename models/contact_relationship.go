package models

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
)

// ContactRelationship is an object representing the database table.
type ContactRelationship struct {
	ContactRelationshipID int `boil:"contact_relationship_id" json:"contact_relationship_id" toml:"contact_relationship_id" yaml:"contact_relationship_id"`
	TypeID                int `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	SubjectID             int `boil:"subject_id" json:"subject_id" toml:"subject_id" yaml:"subject_id"`
	ObjectID              int `boil:"object_id" json:"object_id" toml:"object_id" yaml:"object_id"`

	R *contactRelationshipR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L contactRelationshipL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// contactRelationshipR is where relationships are stored.
type contactRelationshipR struct {
	Object  *Contact
	Subject *Contact
	Type    *Cvterm
}

// contactRelationshipL is where Load methods for each relationship are stored.
type contactRelationshipL struct{}

var (
	contactRelationshipColumns               = []string{"contact_relationship_id", "type_id", "subject_id", "object_id"}
	contactRelationshipColumnsWithoutDefault = []string{"type_id", "subject_id", "object_id"}
	contactRelationshipColumnsWithDefault    = []string{"contact_relationship_id"}
	contactRelationshipPrimaryKeyColumns     = []string{"contact_relationship_id"}
)

type (
	// ContactRelationshipSlice is an alias for a slice of pointers to ContactRelationship.
	// This should generally be used opposed to []ContactRelationship.
	ContactRelationshipSlice []*ContactRelationship
	// ContactRelationshipHook is the signature for custom ContactRelationship hook methods
	ContactRelationshipHook func(boil.Executor, *ContactRelationship) error

	contactRelationshipQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	contactRelationshipType                 = reflect.TypeOf(&ContactRelationship{})
	contactRelationshipMapping              = queries.MakeStructMapping(contactRelationshipType)
	contactRelationshipPrimaryKeyMapping, _ = queries.BindMapping(contactRelationshipType, contactRelationshipMapping, contactRelationshipPrimaryKeyColumns)
	contactRelationshipInsertCacheMut       sync.RWMutex
	contactRelationshipInsertCache          = make(map[string]insertCache)
	contactRelationshipUpdateCacheMut       sync.RWMutex
	contactRelationshipUpdateCache          = make(map[string]updateCache)
	contactRelationshipUpsertCacheMut       sync.RWMutex
	contactRelationshipUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var contactRelationshipBeforeInsertHooks []ContactRelationshipHook
var contactRelationshipBeforeUpdateHooks []ContactRelationshipHook
var contactRelationshipBeforeDeleteHooks []ContactRelationshipHook
var contactRelationshipBeforeUpsertHooks []ContactRelationshipHook

var contactRelationshipAfterInsertHooks []ContactRelationshipHook
var contactRelationshipAfterSelectHooks []ContactRelationshipHook
var contactRelationshipAfterUpdateHooks []ContactRelationshipHook
var contactRelationshipAfterDeleteHooks []ContactRelationshipHook
var contactRelationshipAfterUpsertHooks []ContactRelationshipHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ContactRelationship) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range contactRelationshipBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ContactRelationship) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range contactRelationshipBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ContactRelationship) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range contactRelationshipBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ContactRelationship) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range contactRelationshipBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ContactRelationship) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range contactRelationshipAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ContactRelationship) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range contactRelationshipAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ContactRelationship) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range contactRelationshipAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ContactRelationship) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range contactRelationshipAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ContactRelationship) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range contactRelationshipAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddContactRelationshipHook registers your hook function for all future operations.
func AddContactRelationshipHook(hookPoint boil.HookPoint, contactRelationshipHook ContactRelationshipHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		contactRelationshipBeforeInsertHooks = append(contactRelationshipBeforeInsertHooks, contactRelationshipHook)
	case boil.BeforeUpdateHook:
		contactRelationshipBeforeUpdateHooks = append(contactRelationshipBeforeUpdateHooks, contactRelationshipHook)
	case boil.BeforeDeleteHook:
		contactRelationshipBeforeDeleteHooks = append(contactRelationshipBeforeDeleteHooks, contactRelationshipHook)
	case boil.BeforeUpsertHook:
		contactRelationshipBeforeUpsertHooks = append(contactRelationshipBeforeUpsertHooks, contactRelationshipHook)
	case boil.AfterInsertHook:
		contactRelationshipAfterInsertHooks = append(contactRelationshipAfterInsertHooks, contactRelationshipHook)
	case boil.AfterSelectHook:
		contactRelationshipAfterSelectHooks = append(contactRelationshipAfterSelectHooks, contactRelationshipHook)
	case boil.AfterUpdateHook:
		contactRelationshipAfterUpdateHooks = append(contactRelationshipAfterUpdateHooks, contactRelationshipHook)
	case boil.AfterDeleteHook:
		contactRelationshipAfterDeleteHooks = append(contactRelationshipAfterDeleteHooks, contactRelationshipHook)
	case boil.AfterUpsertHook:
		contactRelationshipAfterUpsertHooks = append(contactRelationshipAfterUpsertHooks, contactRelationshipHook)
	}
}

// OneP returns a single contactRelationship record from the query, and panics on error.
func (q contactRelationshipQuery) OneP() *ContactRelationship {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single contactRelationship record from the query.
func (q contactRelationshipQuery) One() (*ContactRelationship, error) {
	o := &ContactRelationship{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for contact_relationship")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all ContactRelationship records from the query, and panics on error.
func (q contactRelationshipQuery) AllP() ContactRelationshipSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all ContactRelationship records from the query.
func (q contactRelationshipQuery) All() (ContactRelationshipSlice, error) {
	var o ContactRelationshipSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ContactRelationship slice")
	}

	if len(contactRelationshipAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all ContactRelationship records in the query, and panics on error.
func (q contactRelationshipQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all ContactRelationship records in the query.
func (q contactRelationshipQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count contact_relationship rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q contactRelationshipQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q contactRelationshipQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if contact_relationship exists")
	}

	return count > 0, nil
}

// ObjectG pointed to by the foreign key.
func (o *ContactRelationship) ObjectG(mods ...qm.QueryMod) contactQuery {
	return o.Object(boil.GetDB(), mods...)
}

// Object pointed to by the foreign key.
func (o *ContactRelationship) Object(exec boil.Executor, mods ...qm.QueryMod) contactQuery {
	queryMods := []qm.QueryMod{
		qm.Where("contact_id=$1", o.ObjectID),
	}

	queryMods = append(queryMods, mods...)

	query := Contacts(exec, queryMods...)
	queries.SetFrom(query.Query, "\"contact\"")

	return query
}

// SubjectG pointed to by the foreign key.
func (o *ContactRelationship) SubjectG(mods ...qm.QueryMod) contactQuery {
	return o.Subject(boil.GetDB(), mods...)
}

// Subject pointed to by the foreign key.
func (o *ContactRelationship) Subject(exec boil.Executor, mods ...qm.QueryMod) contactQuery {
	queryMods := []qm.QueryMod{
		qm.Where("contact_id=$1", o.SubjectID),
	}

	queryMods = append(queryMods, mods...)

	query := Contacts(exec, queryMods...)
	queries.SetFrom(query.Query, "\"contact\"")

	return query
}

// TypeG pointed to by the foreign key.
func (o *ContactRelationship) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *ContactRelationship) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// LoadObject allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (contactRelationshipL) LoadObject(e boil.Executor, singular bool, maybeContactRelationship interface{}) error {
	var slice []*ContactRelationship
	var object *ContactRelationship

	count := 1
	if singular {
		object = maybeContactRelationship.(*ContactRelationship)
	} else {
		slice = *maybeContactRelationship.(*ContactRelationshipSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &contactRelationshipR{}
		args[0] = object.ObjectID
	} else {
		for i, obj := range slice {
			obj.R = &contactRelationshipR{}
			args[i] = obj.ObjectID
		}
	}

	query := fmt.Sprintf(
		"select * from \"contact\" where \"contact_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Contact")
	}
	defer results.Close()

	var resultSlice []*Contact
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Contact")
	}

	if len(contactRelationshipAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Object = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ObjectID == foreign.ContactID {
				local.R.Object = foreign
				break
			}
		}
	}

	return nil
}

// LoadSubject allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (contactRelationshipL) LoadSubject(e boil.Executor, singular bool, maybeContactRelationship interface{}) error {
	var slice []*ContactRelationship
	var object *ContactRelationship

	count := 1
	if singular {
		object = maybeContactRelationship.(*ContactRelationship)
	} else {
		slice = *maybeContactRelationship.(*ContactRelationshipSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &contactRelationshipR{}
		args[0] = object.SubjectID
	} else {
		for i, obj := range slice {
			obj.R = &contactRelationshipR{}
			args[i] = obj.SubjectID
		}
	}

	query := fmt.Sprintf(
		"select * from \"contact\" where \"contact_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Contact")
	}
	defer results.Close()

	var resultSlice []*Contact
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Contact")
	}

	if len(contactRelationshipAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Subject = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.SubjectID == foreign.ContactID {
				local.R.Subject = foreign
				break
			}
		}
	}

	return nil
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (contactRelationshipL) LoadType(e boil.Executor, singular bool, maybeContactRelationship interface{}) error {
	var slice []*ContactRelationship
	var object *ContactRelationship

	count := 1
	if singular {
		object = maybeContactRelationship.(*ContactRelationship)
	} else {
		slice = *maybeContactRelationship.(*ContactRelationshipSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &contactRelationshipR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &contactRelationshipR{}
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

	if len(contactRelationshipAfterSelectHooks) != 0 {
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
			if local.TypeID == foreign.CvtermID {
				local.R.Type = foreign
				break
			}
		}
	}

	return nil
}

// SetObject of the contact_relationship to the related item.
// Sets o.R.Object to related.
// Adds o to related.R.ObjectContactRelationship.
func (o *ContactRelationship) SetObject(exec boil.Executor, insert bool, related *Contact) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"contact_relationship\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"object_id"}),
		strmangle.WhereClause("\"", "\"", 2, contactRelationshipPrimaryKeyColumns),
	)
	values := []interface{}{related.ContactID, o.ContactRelationshipID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ObjectID = related.ContactID

	if o.R == nil {
		o.R = &contactRelationshipR{
			Object: related,
		}
	} else {
		o.R.Object = related
	}

	if related.R == nil {
		related.R = &contactR{
			ObjectContactRelationship: o,
		}
	} else {
		related.R.ObjectContactRelationship = o
	}

	return nil
}

// SetSubject of the contact_relationship to the related item.
// Sets o.R.Subject to related.
// Adds o to related.R.SubjectContactRelationship.
func (o *ContactRelationship) SetSubject(exec boil.Executor, insert bool, related *Contact) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"contact_relationship\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"subject_id"}),
		strmangle.WhereClause("\"", "\"", 2, contactRelationshipPrimaryKeyColumns),
	)
	values := []interface{}{related.ContactID, o.ContactRelationshipID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.SubjectID = related.ContactID

	if o.R == nil {
		o.R = &contactRelationshipR{
			Subject: related,
		}
	} else {
		o.R.Subject = related
	}

	if related.R == nil {
		related.R = &contactR{
			SubjectContactRelationship: o,
		}
	} else {
		related.R.SubjectContactRelationship = o
	}

	return nil
}

// SetType of the contact_relationship to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeContactRelationship.
func (o *ContactRelationship) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"contact_relationship\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, contactRelationshipPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.ContactRelationshipID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &contactRelationshipR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeContactRelationship: o,
		}
	} else {
		related.R.TypeContactRelationship = o
	}

	return nil
}

// ContactRelationshipsG retrieves all records.
func ContactRelationshipsG(mods ...qm.QueryMod) contactRelationshipQuery {
	return ContactRelationships(boil.GetDB(), mods...)
}

// ContactRelationships retrieves all the records using an executor.
func ContactRelationships(exec boil.Executor, mods ...qm.QueryMod) contactRelationshipQuery {
	mods = append(mods, qm.From("\"contact_relationship\""))
	return contactRelationshipQuery{NewQuery(exec, mods...)}
}

// FindContactRelationshipG retrieves a single record by ID.
func FindContactRelationshipG(contactRelationshipID int, selectCols ...string) (*ContactRelationship, error) {
	return FindContactRelationship(boil.GetDB(), contactRelationshipID, selectCols...)
}

// FindContactRelationshipGP retrieves a single record by ID, and panics on error.
func FindContactRelationshipGP(contactRelationshipID int, selectCols ...string) *ContactRelationship {
	retobj, err := FindContactRelationship(boil.GetDB(), contactRelationshipID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindContactRelationship retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindContactRelationship(exec boil.Executor, contactRelationshipID int, selectCols ...string) (*ContactRelationship, error) {
	contactRelationshipObj := &ContactRelationship{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"contact_relationship\" where \"contact_relationship_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, contactRelationshipID)

	err := q.Bind(contactRelationshipObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from contact_relationship")
	}

	return contactRelationshipObj, nil
}

// FindContactRelationshipP retrieves a single record by ID with an executor, and panics on error.
func FindContactRelationshipP(exec boil.Executor, contactRelationshipID int, selectCols ...string) *ContactRelationship {
	retobj, err := FindContactRelationship(exec, contactRelationshipID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ContactRelationship) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *ContactRelationship) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *ContactRelationship) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *ContactRelationship) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no contact_relationship provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(contactRelationshipColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	contactRelationshipInsertCacheMut.RLock()
	cache, cached := contactRelationshipInsertCache[key]
	contactRelationshipInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			contactRelationshipColumns,
			contactRelationshipColumnsWithDefault,
			contactRelationshipColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(contactRelationshipType, contactRelationshipMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(contactRelationshipType, contactRelationshipMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"contact_relationship\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into contact_relationship")
	}

	if !cached {
		contactRelationshipInsertCacheMut.Lock()
		contactRelationshipInsertCache[key] = cache
		contactRelationshipInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single ContactRelationship record. See Update for
// whitelist behavior description.
func (o *ContactRelationship) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single ContactRelationship record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *ContactRelationship) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the ContactRelationship, and panics on error.
// See Update for whitelist behavior description.
func (o *ContactRelationship) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the ContactRelationship.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *ContactRelationship) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	contactRelationshipUpdateCacheMut.RLock()
	cache, cached := contactRelationshipUpdateCache[key]
	contactRelationshipUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(contactRelationshipColumns, contactRelationshipPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update contact_relationship, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"contact_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, contactRelationshipPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(contactRelationshipType, contactRelationshipMapping, append(wl, contactRelationshipPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update contact_relationship row")
	}

	if !cached {
		contactRelationshipUpdateCacheMut.Lock()
		contactRelationshipUpdateCache[key] = cache
		contactRelationshipUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q contactRelationshipQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q contactRelationshipQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for contact_relationship")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ContactRelationshipSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o ContactRelationshipSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o ContactRelationshipSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ContactRelationshipSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), contactRelationshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"contact_relationship\" SET %s WHERE (\"contact_relationship_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(contactRelationshipPrimaryKeyColumns), len(colNames)+1, len(contactRelationshipPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in contactRelationship slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ContactRelationship) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *ContactRelationship) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *ContactRelationship) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *ContactRelationship) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no contact_relationship provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(contactRelationshipColumnsWithDefault, o)

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

	contactRelationshipUpsertCacheMut.RLock()
	cache, cached := contactRelationshipUpsertCache[key]
	contactRelationshipUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			contactRelationshipColumns,
			contactRelationshipColumnsWithDefault,
			contactRelationshipColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			contactRelationshipColumns,
			contactRelationshipPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert contact_relationship, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(contactRelationshipPrimaryKeyColumns))
			copy(conflict, contactRelationshipPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"contact_relationship\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(contactRelationshipType, contactRelationshipMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(contactRelationshipType, contactRelationshipMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for contact_relationship")
	}

	if !cached {
		contactRelationshipUpsertCacheMut.Lock()
		contactRelationshipUpsertCache[key] = cache
		contactRelationshipUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single ContactRelationship record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *ContactRelationship) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single ContactRelationship record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ContactRelationship) DeleteG() error {
	if o == nil {
		return errors.New("models: no ContactRelationship provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single ContactRelationship record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *ContactRelationship) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single ContactRelationship record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ContactRelationship) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no ContactRelationship provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), contactRelationshipPrimaryKeyMapping)
	sql := "DELETE FROM \"contact_relationship\" WHERE \"contact_relationship_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from contact_relationship")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q contactRelationshipQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q contactRelationshipQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no contactRelationshipQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from contact_relationship")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o ContactRelationshipSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o ContactRelationshipSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no ContactRelationship slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o ContactRelationshipSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ContactRelationshipSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no ContactRelationship slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(contactRelationshipBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), contactRelationshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"contact_relationship\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, contactRelationshipPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(contactRelationshipPrimaryKeyColumns), 1, len(contactRelationshipPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from contactRelationship slice")
	}

	if len(contactRelationshipAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *ContactRelationship) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *ContactRelationship) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ContactRelationship) ReloadG() error {
	if o == nil {
		return errors.New("models: no ContactRelationship provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ContactRelationship) Reload(exec boil.Executor) error {
	ret, err := FindContactRelationship(exec, o.ContactRelationshipID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ContactRelationshipSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ContactRelationshipSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ContactRelationshipSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty ContactRelationshipSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ContactRelationshipSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	contactRelationships := ContactRelationshipSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), contactRelationshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"contact_relationship\".* FROM \"contact_relationship\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, contactRelationshipPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(contactRelationshipPrimaryKeyColumns), 1, len(contactRelationshipPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&contactRelationships)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ContactRelationshipSlice")
	}

	*o = contactRelationships

	return nil
}

// ContactRelationshipExists checks if the ContactRelationship row exists.
func ContactRelationshipExists(exec boil.Executor, contactRelationshipID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"contact_relationship\" where \"contact_relationship_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, contactRelationshipID)
	}

	row := exec.QueryRow(sql, contactRelationshipID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if contact_relationship exists")
	}

	return exists, nil
}

// ContactRelationshipExistsG checks if the ContactRelationship row exists.
func ContactRelationshipExistsG(contactRelationshipID int) (bool, error) {
	return ContactRelationshipExists(boil.GetDB(), contactRelationshipID)
}

// ContactRelationshipExistsGP checks if the ContactRelationship row exists. Panics on error.
func ContactRelationshipExistsGP(contactRelationshipID int) bool {
	e, err := ContactRelationshipExists(boil.GetDB(), contactRelationshipID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ContactRelationshipExistsP checks if the ContactRelationship row exists. Panics on error.
func ContactRelationshipExistsP(exec boil.Executor, contactRelationshipID int) bool {
	e, err := ContactRelationshipExists(exec, contactRelationshipID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
