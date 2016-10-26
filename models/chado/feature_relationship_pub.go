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
)

// FeatureRelationshipPub is an object representing the database table.
type FeatureRelationshipPub struct {
	FeatureRelationshipPubID int `boil:"feature_relationship_pub_id" json:"feature_relationship_pub_id" toml:"feature_relationship_pub_id" yaml:"feature_relationship_pub_id"`
	FeatureRelationshipID    int `boil:"feature_relationship_id" json:"feature_relationship_id" toml:"feature_relationship_id" yaml:"feature_relationship_id"`
	PubID                    int `boil:"pub_id" json:"pub_id" toml:"pub_id" yaml:"pub_id"`

	R *featureRelationshipPubR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L featureRelationshipPubL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// featureRelationshipPubR is where relationships are stored.
type featureRelationshipPubR struct {
	FeatureRelationship *FeatureRelationship
	Pub                 *Pub
}

// featureRelationshipPubL is where Load methods for each relationship are stored.
type featureRelationshipPubL struct{}

var (
	featureRelationshipPubColumns               = []string{"feature_relationship_pub_id", "feature_relationship_id", "pub_id"}
	featureRelationshipPubColumnsWithoutDefault = []string{"feature_relationship_id", "pub_id"}
	featureRelationshipPubColumnsWithDefault    = []string{"feature_relationship_pub_id"}
	featureRelationshipPubPrimaryKeyColumns     = []string{"feature_relationship_pub_id"}
)

type (
	// FeatureRelationshipPubSlice is an alias for a slice of pointers to FeatureRelationshipPub.
	// This should generally be used opposed to []FeatureRelationshipPub.
	FeatureRelationshipPubSlice []*FeatureRelationshipPub
	// FeatureRelationshipPubHook is the signature for custom FeatureRelationshipPub hook methods
	FeatureRelationshipPubHook func(boil.Executor, *FeatureRelationshipPub) error

	featureRelationshipPubQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	featureRelationshipPubType                 = reflect.TypeOf(&FeatureRelationshipPub{})
	featureRelationshipPubMapping              = queries.MakeStructMapping(featureRelationshipPubType)
	featureRelationshipPubPrimaryKeyMapping, _ = queries.BindMapping(featureRelationshipPubType, featureRelationshipPubMapping, featureRelationshipPubPrimaryKeyColumns)
	featureRelationshipPubInsertCacheMut       sync.RWMutex
	featureRelationshipPubInsertCache          = make(map[string]insertCache)
	featureRelationshipPubUpdateCacheMut       sync.RWMutex
	featureRelationshipPubUpdateCache          = make(map[string]updateCache)
	featureRelationshipPubUpsertCacheMut       sync.RWMutex
	featureRelationshipPubUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var featureRelationshipPubBeforeInsertHooks []FeatureRelationshipPubHook
var featureRelationshipPubBeforeUpdateHooks []FeatureRelationshipPubHook
var featureRelationshipPubBeforeDeleteHooks []FeatureRelationshipPubHook
var featureRelationshipPubBeforeUpsertHooks []FeatureRelationshipPubHook

var featureRelationshipPubAfterInsertHooks []FeatureRelationshipPubHook
var featureRelationshipPubAfterSelectHooks []FeatureRelationshipPubHook
var featureRelationshipPubAfterUpdateHooks []FeatureRelationshipPubHook
var featureRelationshipPubAfterDeleteHooks []FeatureRelationshipPubHook
var featureRelationshipPubAfterUpsertHooks []FeatureRelationshipPubHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FeatureRelationshipPub) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshipPubBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FeatureRelationshipPub) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshipPubBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FeatureRelationshipPub) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshipPubBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FeatureRelationshipPub) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshipPubBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FeatureRelationshipPub) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshipPubAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FeatureRelationshipPub) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshipPubAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FeatureRelationshipPub) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshipPubAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FeatureRelationshipPub) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshipPubAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FeatureRelationshipPub) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshipPubAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFeatureRelationshipPubHook registers your hook function for all future operations.
func AddFeatureRelationshipPubHook(hookPoint boil.HookPoint, featureRelationshipPubHook FeatureRelationshipPubHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		featureRelationshipPubBeforeInsertHooks = append(featureRelationshipPubBeforeInsertHooks, featureRelationshipPubHook)
	case boil.BeforeUpdateHook:
		featureRelationshipPubBeforeUpdateHooks = append(featureRelationshipPubBeforeUpdateHooks, featureRelationshipPubHook)
	case boil.BeforeDeleteHook:
		featureRelationshipPubBeforeDeleteHooks = append(featureRelationshipPubBeforeDeleteHooks, featureRelationshipPubHook)
	case boil.BeforeUpsertHook:
		featureRelationshipPubBeforeUpsertHooks = append(featureRelationshipPubBeforeUpsertHooks, featureRelationshipPubHook)
	case boil.AfterInsertHook:
		featureRelationshipPubAfterInsertHooks = append(featureRelationshipPubAfterInsertHooks, featureRelationshipPubHook)
	case boil.AfterSelectHook:
		featureRelationshipPubAfterSelectHooks = append(featureRelationshipPubAfterSelectHooks, featureRelationshipPubHook)
	case boil.AfterUpdateHook:
		featureRelationshipPubAfterUpdateHooks = append(featureRelationshipPubAfterUpdateHooks, featureRelationshipPubHook)
	case boil.AfterDeleteHook:
		featureRelationshipPubAfterDeleteHooks = append(featureRelationshipPubAfterDeleteHooks, featureRelationshipPubHook)
	case boil.AfterUpsertHook:
		featureRelationshipPubAfterUpsertHooks = append(featureRelationshipPubAfterUpsertHooks, featureRelationshipPubHook)
	}
}

// OneP returns a single featureRelationshipPub record from the query, and panics on error.
func (q featureRelationshipPubQuery) OneP() *FeatureRelationshipPub {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single featureRelationshipPub record from the query.
func (q featureRelationshipPubQuery) One() (*FeatureRelationshipPub, error) {
	o := &FeatureRelationshipPub{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for feature_relationship_pub")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all FeatureRelationshipPub records from the query, and panics on error.
func (q featureRelationshipPubQuery) AllP() FeatureRelationshipPubSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all FeatureRelationshipPub records from the query.
func (q featureRelationshipPubQuery) All() (FeatureRelationshipPubSlice, error) {
	var o FeatureRelationshipPubSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to FeatureRelationshipPub slice")
	}

	if len(featureRelationshipPubAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all FeatureRelationshipPub records in the query, and panics on error.
func (q featureRelationshipPubQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all FeatureRelationshipPub records in the query.
func (q featureRelationshipPubQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count feature_relationship_pub rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q featureRelationshipPubQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q featureRelationshipPubQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if feature_relationship_pub exists")
	}

	return count > 0, nil
}

// FeatureRelationshipG pointed to by the foreign key.
func (o *FeatureRelationshipPub) FeatureRelationshipG(mods ...qm.QueryMod) featureRelationshipQuery {
	return o.FeatureRelationship(boil.GetDB(), mods...)
}

// FeatureRelationship pointed to by the foreign key.
func (o *FeatureRelationshipPub) FeatureRelationship(exec boil.Executor, mods ...qm.QueryMod) featureRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_relationship_id=$1", o.FeatureRelationshipID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_relationship\"")

	return query
}

// PubG pointed to by the foreign key.
func (o *FeatureRelationshipPub) PubG(mods ...qm.QueryMod) pubQuery {
	return o.Pub(boil.GetDB(), mods...)
}

// Pub pointed to by the foreign key.
func (o *FeatureRelationshipPub) Pub(exec boil.Executor, mods ...qm.QueryMod) pubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub\"")

	return query
}

// LoadFeatureRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureRelationshipPubL) LoadFeatureRelationship(e boil.Executor, singular bool, maybeFeatureRelationshipPub interface{}) error {
	var slice []*FeatureRelationshipPub
	var object *FeatureRelationshipPub

	count := 1
	if singular {
		object = maybeFeatureRelationshipPub.(*FeatureRelationshipPub)
	} else {
		slice = *maybeFeatureRelationshipPub.(*FeatureRelationshipPubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureRelationshipPubR{}
		args[0] = object.FeatureRelationshipID
	} else {
		for i, obj := range slice {
			obj.R = &featureRelationshipPubR{}
			args[i] = obj.FeatureRelationshipID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_relationship\" where \"feature_relationship_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureRelationship")
	}
	defer results.Close()

	var resultSlice []*FeatureRelationship
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureRelationship")
	}

	if len(featureRelationshipPubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeatureRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureRelationshipID == foreign.FeatureRelationshipID {
				local.R.FeatureRelationship = foreign
				break
			}
		}
	}

	return nil
}

// LoadPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureRelationshipPubL) LoadPub(e boil.Executor, singular bool, maybeFeatureRelationshipPub interface{}) error {
	var slice []*FeatureRelationshipPub
	var object *FeatureRelationshipPub

	count := 1
	if singular {
		object = maybeFeatureRelationshipPub.(*FeatureRelationshipPub)
	} else {
		slice = *maybeFeatureRelationshipPub.(*FeatureRelationshipPubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureRelationshipPubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &featureRelationshipPubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"pub\" where \"pub_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Pub")
	}
	defer results.Close()

	var resultSlice []*Pub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Pub")
	}

	if len(featureRelationshipPubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Pub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PubID == foreign.PubID {
				local.R.Pub = foreign
				break
			}
		}
	}

	return nil
}

// SetFeatureRelationship of the feature_relationship_pub to the related item.
// Sets o.R.FeatureRelationship to related.
// Adds o to related.R.FeatureRelationshipPub.
func (o *FeatureRelationshipPub) SetFeatureRelationship(exec boil.Executor, insert bool, related *FeatureRelationship) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_relationship_pub\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"feature_relationship_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureRelationshipPubPrimaryKeyColumns),
	)
	values := []interface{}{related.FeatureRelationshipID, o.FeatureRelationshipPubID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FeatureRelationshipID = related.FeatureRelationshipID

	if o.R == nil {
		o.R = &featureRelationshipPubR{
			FeatureRelationship: related,
		}
	} else {
		o.R.FeatureRelationship = related
	}

	if related.R == nil {
		related.R = &featureRelationshipR{
			FeatureRelationshipPub: o,
		}
	} else {
		related.R.FeatureRelationshipPub = o
	}

	return nil
}

// SetPub of the feature_relationship_pub to the related item.
// Sets o.R.Pub to related.
// Adds o to related.R.FeatureRelationshipPub.
func (o *FeatureRelationshipPub) SetPub(exec boil.Executor, insert bool, related *Pub) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_relationship_pub\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureRelationshipPubPrimaryKeyColumns),
	)
	values := []interface{}{related.PubID, o.FeatureRelationshipPubID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PubID = related.PubID

	if o.R == nil {
		o.R = &featureRelationshipPubR{
			Pub: related,
		}
	} else {
		o.R.Pub = related
	}

	if related.R == nil {
		related.R = &pubR{
			FeatureRelationshipPub: o,
		}
	} else {
		related.R.FeatureRelationshipPub = o
	}

	return nil
}

// FeatureRelationshipPubsG retrieves all records.
func FeatureRelationshipPubsG(mods ...qm.QueryMod) featureRelationshipPubQuery {
	return FeatureRelationshipPubs(boil.GetDB(), mods...)
}

// FeatureRelationshipPubs retrieves all the records using an executor.
func FeatureRelationshipPubs(exec boil.Executor, mods ...qm.QueryMod) featureRelationshipPubQuery {
	mods = append(mods, qm.From("\"feature_relationship_pub\""))
	return featureRelationshipPubQuery{NewQuery(exec, mods...)}
}

// FindFeatureRelationshipPubG retrieves a single record by ID.
func FindFeatureRelationshipPubG(featureRelationshipPubID int, selectCols ...string) (*FeatureRelationshipPub, error) {
	return FindFeatureRelationshipPub(boil.GetDB(), featureRelationshipPubID, selectCols...)
}

// FindFeatureRelationshipPubGP retrieves a single record by ID, and panics on error.
func FindFeatureRelationshipPubGP(featureRelationshipPubID int, selectCols ...string) *FeatureRelationshipPub {
	retobj, err := FindFeatureRelationshipPub(boil.GetDB(), featureRelationshipPubID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFeatureRelationshipPub retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeatureRelationshipPub(exec boil.Executor, featureRelationshipPubID int, selectCols ...string) (*FeatureRelationshipPub, error) {
	featureRelationshipPubObj := &FeatureRelationshipPub{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"feature_relationship_pub\" where \"feature_relationship_pub_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, featureRelationshipPubID)

	err := q.Bind(featureRelationshipPubObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from feature_relationship_pub")
	}

	return featureRelationshipPubObj, nil
}

// FindFeatureRelationshipPubP retrieves a single record by ID with an executor, and panics on error.
func FindFeatureRelationshipPubP(exec boil.Executor, featureRelationshipPubID int, selectCols ...string) *FeatureRelationshipPub {
	retobj, err := FindFeatureRelationshipPub(exec, featureRelationshipPubID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *FeatureRelationshipPub) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *FeatureRelationshipPub) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *FeatureRelationshipPub) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *FeatureRelationshipPub) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no feature_relationship_pub provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureRelationshipPubColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	featureRelationshipPubInsertCacheMut.RLock()
	cache, cached := featureRelationshipPubInsertCache[key]
	featureRelationshipPubInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			featureRelationshipPubColumns,
			featureRelationshipPubColumnsWithDefault,
			featureRelationshipPubColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(featureRelationshipPubType, featureRelationshipPubMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(featureRelationshipPubType, featureRelationshipPubMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"feature_relationship_pub\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into feature_relationship_pub")
	}

	if !cached {
		featureRelationshipPubInsertCacheMut.Lock()
		featureRelationshipPubInsertCache[key] = cache
		featureRelationshipPubInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single FeatureRelationshipPub record. See Update for
// whitelist behavior description.
func (o *FeatureRelationshipPub) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single FeatureRelationshipPub record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *FeatureRelationshipPub) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the FeatureRelationshipPub, and panics on error.
// See Update for whitelist behavior description.
func (o *FeatureRelationshipPub) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the FeatureRelationshipPub.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *FeatureRelationshipPub) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	featureRelationshipPubUpdateCacheMut.RLock()
	cache, cached := featureRelationshipPubUpdateCache[key]
	featureRelationshipPubUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(featureRelationshipPubColumns, featureRelationshipPubPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update feature_relationship_pub, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"feature_relationship_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, featureRelationshipPubPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(featureRelationshipPubType, featureRelationshipPubMapping, append(wl, featureRelationshipPubPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update feature_relationship_pub row")
	}

	if !cached {
		featureRelationshipPubUpdateCacheMut.Lock()
		featureRelationshipPubUpdateCache[key] = cache
		featureRelationshipPubUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q featureRelationshipPubQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q featureRelationshipPubQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for feature_relationship_pub")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FeatureRelationshipPubSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FeatureRelationshipPubSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FeatureRelationshipPubSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FeatureRelationshipPubSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureRelationshipPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"feature_relationship_pub\" SET %s WHERE (\"feature_relationship_pub_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featureRelationshipPubPrimaryKeyColumns), len(colNames)+1, len(featureRelationshipPubPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in featureRelationshipPub slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *FeatureRelationshipPub) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *FeatureRelationshipPub) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *FeatureRelationshipPub) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *FeatureRelationshipPub) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no feature_relationship_pub provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureRelationshipPubColumnsWithDefault, o)

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

	featureRelationshipPubUpsertCacheMut.RLock()
	cache, cached := featureRelationshipPubUpsertCache[key]
	featureRelationshipPubUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			featureRelationshipPubColumns,
			featureRelationshipPubColumnsWithDefault,
			featureRelationshipPubColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			featureRelationshipPubColumns,
			featureRelationshipPubPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert feature_relationship_pub, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(featureRelationshipPubPrimaryKeyColumns))
			copy(conflict, featureRelationshipPubPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"feature_relationship_pub\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(featureRelationshipPubType, featureRelationshipPubMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(featureRelationshipPubType, featureRelationshipPubMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for feature_relationship_pub")
	}

	if !cached {
		featureRelationshipPubUpsertCacheMut.Lock()
		featureRelationshipPubUpsertCache[key] = cache
		featureRelationshipPubUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single FeatureRelationshipPub record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeatureRelationshipPub) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single FeatureRelationshipPub record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *FeatureRelationshipPub) DeleteG() error {
	if o == nil {
		return errors.New("chado: no FeatureRelationshipPub provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single FeatureRelationshipPub record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeatureRelationshipPub) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single FeatureRelationshipPub record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FeatureRelationshipPub) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no FeatureRelationshipPub provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), featureRelationshipPubPrimaryKeyMapping)
	sql := "DELETE FROM \"feature_relationship_pub\" WHERE \"feature_relationship_pub_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from feature_relationship_pub")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q featureRelationshipPubQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q featureRelationshipPubQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no featureRelationshipPubQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from feature_relationship_pub")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FeatureRelationshipPubSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FeatureRelationshipPubSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no FeatureRelationshipPub slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FeatureRelationshipPubSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FeatureRelationshipPubSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no FeatureRelationshipPub slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(featureRelationshipPubBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureRelationshipPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"feature_relationship_pub\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featureRelationshipPubPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featureRelationshipPubPrimaryKeyColumns), 1, len(featureRelationshipPubPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from featureRelationshipPub slice")
	}

	if len(featureRelationshipPubAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *FeatureRelationshipPub) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *FeatureRelationshipPub) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *FeatureRelationshipPub) ReloadG() error {
	if o == nil {
		return errors.New("chado: no FeatureRelationshipPub provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FeatureRelationshipPub) Reload(exec boil.Executor) error {
	ret, err := FindFeatureRelationshipPub(exec, o.FeatureRelationshipPubID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureRelationshipPubSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureRelationshipPubSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureRelationshipPubSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty FeatureRelationshipPubSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureRelationshipPubSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	featureRelationshipPubs := FeatureRelationshipPubSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureRelationshipPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"feature_relationship_pub\".* FROM \"feature_relationship_pub\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featureRelationshipPubPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(featureRelationshipPubPrimaryKeyColumns), 1, len(featureRelationshipPubPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&featureRelationshipPubs)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in FeatureRelationshipPubSlice")
	}

	*o = featureRelationshipPubs

	return nil
}

// FeatureRelationshipPubExists checks if the FeatureRelationshipPub row exists.
func FeatureRelationshipPubExists(exec boil.Executor, featureRelationshipPubID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"feature_relationship_pub\" where \"feature_relationship_pub_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, featureRelationshipPubID)
	}

	row := exec.QueryRow(sql, featureRelationshipPubID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if feature_relationship_pub exists")
	}

	return exists, nil
}

// FeatureRelationshipPubExistsG checks if the FeatureRelationshipPub row exists.
func FeatureRelationshipPubExistsG(featureRelationshipPubID int) (bool, error) {
	return FeatureRelationshipPubExists(boil.GetDB(), featureRelationshipPubID)
}

// FeatureRelationshipPubExistsGP checks if the FeatureRelationshipPub row exists. Panics on error.
func FeatureRelationshipPubExistsGP(featureRelationshipPubID int) bool {
	e, err := FeatureRelationshipPubExists(boil.GetDB(), featureRelationshipPubID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FeatureRelationshipPubExistsP checks if the FeatureRelationshipPub row exists. Panics on error.
func FeatureRelationshipPubExistsP(exec boil.Executor, featureRelationshipPubID int) bool {
	e, err := FeatureRelationshipPubExists(exec, featureRelationshipPubID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
