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
	"gopkg.in/nullbio/null.v5"
)

// FeatureRelationshipprop is an object representing the database table.
type FeatureRelationshipprop struct {
	FeatureRelationshippropID int         `boil:"feature_relationshipprop_id" json:"feature_relationshipprop_id" toml:"feature_relationshipprop_id" yaml:"feature_relationshipprop_id"`
	FeatureRelationshipID     int         `boil:"feature_relationship_id" json:"feature_relationship_id" toml:"feature_relationship_id" yaml:"feature_relationship_id"`
	TypeID                    int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Value                     null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	Rank                      int         `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *featureRelationshippropR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L featureRelationshippropL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// featureRelationshippropR is where relationships are stored.
type featureRelationshippropR struct {
	FeatureRelationship        *FeatureRelationship
	Type                       *Cvterm
	FeatureRelationshippropPub *FeatureRelationshippropPub
}

// featureRelationshippropL is where Load methods for each relationship are stored.
type featureRelationshippropL struct{}

var (
	featureRelationshippropColumns               = []string{"feature_relationshipprop_id", "feature_relationship_id", "type_id", "value", "rank"}
	featureRelationshippropColumnsWithoutDefault = []string{"feature_relationship_id", "type_id", "value"}
	featureRelationshippropColumnsWithDefault    = []string{"feature_relationshipprop_id", "rank"}
	featureRelationshippropPrimaryKeyColumns     = []string{"feature_relationshipprop_id"}
)

type (
	// FeatureRelationshippropSlice is an alias for a slice of pointers to FeatureRelationshipprop.
	// This should generally be used opposed to []FeatureRelationshipprop.
	FeatureRelationshippropSlice []*FeatureRelationshipprop
	// FeatureRelationshippropHook is the signature for custom FeatureRelationshipprop hook methods
	FeatureRelationshippropHook func(boil.Executor, *FeatureRelationshipprop) error

	featureRelationshippropQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	featureRelationshippropType                 = reflect.TypeOf(&FeatureRelationshipprop{})
	featureRelationshippropMapping              = queries.MakeStructMapping(featureRelationshippropType)
	featureRelationshippropPrimaryKeyMapping, _ = queries.BindMapping(featureRelationshippropType, featureRelationshippropMapping, featureRelationshippropPrimaryKeyColumns)
	featureRelationshippropInsertCacheMut       sync.RWMutex
	featureRelationshippropInsertCache          = make(map[string]insertCache)
	featureRelationshippropUpdateCacheMut       sync.RWMutex
	featureRelationshippropUpdateCache          = make(map[string]updateCache)
	featureRelationshippropUpsertCacheMut       sync.RWMutex
	featureRelationshippropUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var featureRelationshippropBeforeInsertHooks []FeatureRelationshippropHook
var featureRelationshippropBeforeUpdateHooks []FeatureRelationshippropHook
var featureRelationshippropBeforeDeleteHooks []FeatureRelationshippropHook
var featureRelationshippropBeforeUpsertHooks []FeatureRelationshippropHook

var featureRelationshippropAfterInsertHooks []FeatureRelationshippropHook
var featureRelationshippropAfterSelectHooks []FeatureRelationshippropHook
var featureRelationshippropAfterUpdateHooks []FeatureRelationshippropHook
var featureRelationshippropAfterDeleteHooks []FeatureRelationshippropHook
var featureRelationshippropAfterUpsertHooks []FeatureRelationshippropHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FeatureRelationshipprop) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshippropBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FeatureRelationshipprop) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshippropBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FeatureRelationshipprop) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshippropBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FeatureRelationshipprop) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshippropBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FeatureRelationshipprop) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshippropAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FeatureRelationshipprop) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshippropAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FeatureRelationshipprop) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshippropAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FeatureRelationshipprop) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshippropAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FeatureRelationshipprop) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshippropAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFeatureRelationshippropHook registers your hook function for all future operations.
func AddFeatureRelationshippropHook(hookPoint boil.HookPoint, featureRelationshippropHook FeatureRelationshippropHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		featureRelationshippropBeforeInsertHooks = append(featureRelationshippropBeforeInsertHooks, featureRelationshippropHook)
	case boil.BeforeUpdateHook:
		featureRelationshippropBeforeUpdateHooks = append(featureRelationshippropBeforeUpdateHooks, featureRelationshippropHook)
	case boil.BeforeDeleteHook:
		featureRelationshippropBeforeDeleteHooks = append(featureRelationshippropBeforeDeleteHooks, featureRelationshippropHook)
	case boil.BeforeUpsertHook:
		featureRelationshippropBeforeUpsertHooks = append(featureRelationshippropBeforeUpsertHooks, featureRelationshippropHook)
	case boil.AfterInsertHook:
		featureRelationshippropAfterInsertHooks = append(featureRelationshippropAfterInsertHooks, featureRelationshippropHook)
	case boil.AfterSelectHook:
		featureRelationshippropAfterSelectHooks = append(featureRelationshippropAfterSelectHooks, featureRelationshippropHook)
	case boil.AfterUpdateHook:
		featureRelationshippropAfterUpdateHooks = append(featureRelationshippropAfterUpdateHooks, featureRelationshippropHook)
	case boil.AfterDeleteHook:
		featureRelationshippropAfterDeleteHooks = append(featureRelationshippropAfterDeleteHooks, featureRelationshippropHook)
	case boil.AfterUpsertHook:
		featureRelationshippropAfterUpsertHooks = append(featureRelationshippropAfterUpsertHooks, featureRelationshippropHook)
	}
}

// OneP returns a single featureRelationshipprop record from the query, and panics on error.
func (q featureRelationshippropQuery) OneP() *FeatureRelationshipprop {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single featureRelationshipprop record from the query.
func (q featureRelationshippropQuery) One() (*FeatureRelationshipprop, error) {
	o := &FeatureRelationshipprop{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for feature_relationshipprop")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all FeatureRelationshipprop records from the query, and panics on error.
func (q featureRelationshippropQuery) AllP() FeatureRelationshippropSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all FeatureRelationshipprop records from the query.
func (q featureRelationshippropQuery) All() (FeatureRelationshippropSlice, error) {
	var o FeatureRelationshippropSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to FeatureRelationshipprop slice")
	}

	if len(featureRelationshippropAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all FeatureRelationshipprop records in the query, and panics on error.
func (q featureRelationshippropQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all FeatureRelationshipprop records in the query.
func (q featureRelationshippropQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count feature_relationshipprop rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q featureRelationshippropQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q featureRelationshippropQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if feature_relationshipprop exists")
	}

	return count > 0, nil
}

// FeatureRelationshipG pointed to by the foreign key.
func (o *FeatureRelationshipprop) FeatureRelationshipG(mods ...qm.QueryMod) featureRelationshipQuery {
	return o.FeatureRelationship(boil.GetDB(), mods...)
}

// FeatureRelationship pointed to by the foreign key.
func (o *FeatureRelationshipprop) FeatureRelationship(exec boil.Executor, mods ...qm.QueryMod) featureRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_relationship_id=$1", o.FeatureRelationshipID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_relationship\"")

	return query
}

// TypeG pointed to by the foreign key.
func (o *FeatureRelationshipprop) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *FeatureRelationshipprop) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// FeatureRelationshippropPubG pointed to by the foreign key.
func (o *FeatureRelationshipprop) FeatureRelationshippropPubG(mods ...qm.QueryMod) featureRelationshippropPubQuery {
	return o.FeatureRelationshippropPub(boil.GetDB(), mods...)
}

// FeatureRelationshippropPub pointed to by the foreign key.
func (o *FeatureRelationshipprop) FeatureRelationshippropPub(exec boil.Executor, mods ...qm.QueryMod) featureRelationshippropPubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_relationshipprop_id=$1", o.FeatureRelationshippropID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureRelationshippropPubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_relationshipprop_pub\"")

	return query
}

// LoadFeatureRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureRelationshippropL) LoadFeatureRelationship(e boil.Executor, singular bool, maybeFeatureRelationshipprop interface{}) error {
	var slice []*FeatureRelationshipprop
	var object *FeatureRelationshipprop

	count := 1
	if singular {
		object = maybeFeatureRelationshipprop.(*FeatureRelationshipprop)
	} else {
		slice = *maybeFeatureRelationshipprop.(*FeatureRelationshippropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureRelationshippropR{}
		args[0] = object.FeatureRelationshipID
	} else {
		for i, obj := range slice {
			obj.R = &featureRelationshippropR{}
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

	if len(featureRelationshippropAfterSelectHooks) != 0 {
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

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureRelationshippropL) LoadType(e boil.Executor, singular bool, maybeFeatureRelationshipprop interface{}) error {
	var slice []*FeatureRelationshipprop
	var object *FeatureRelationshipprop

	count := 1
	if singular {
		object = maybeFeatureRelationshipprop.(*FeatureRelationshipprop)
	} else {
		slice = *maybeFeatureRelationshipprop.(*FeatureRelationshippropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureRelationshippropR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &featureRelationshippropR{}
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

	if len(featureRelationshippropAfterSelectHooks) != 0 {
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

// LoadFeatureRelationshippropPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureRelationshippropL) LoadFeatureRelationshippropPub(e boil.Executor, singular bool, maybeFeatureRelationshipprop interface{}) error {
	var slice []*FeatureRelationshipprop
	var object *FeatureRelationshipprop

	count := 1
	if singular {
		object = maybeFeatureRelationshipprop.(*FeatureRelationshipprop)
	} else {
		slice = *maybeFeatureRelationshipprop.(*FeatureRelationshippropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureRelationshippropR{}
		args[0] = object.FeatureRelationshippropID
	} else {
		for i, obj := range slice {
			obj.R = &featureRelationshippropR{}
			args[i] = obj.FeatureRelationshippropID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_relationshipprop_pub\" where \"feature_relationshipprop_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureRelationshippropPub")
	}
	defer results.Close()

	var resultSlice []*FeatureRelationshippropPub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureRelationshippropPub")
	}

	if len(featureRelationshippropAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeatureRelationshippropPub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureRelationshippropID == foreign.FeatureRelationshippropID {
				local.R.FeatureRelationshippropPub = foreign
				break
			}
		}
	}

	return nil
}

// SetFeatureRelationship of the feature_relationshipprop to the related item.
// Sets o.R.FeatureRelationship to related.
// Adds o to related.R.FeatureRelationshipprop.
func (o *FeatureRelationshipprop) SetFeatureRelationship(exec boil.Executor, insert bool, related *FeatureRelationship) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_relationshipprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"feature_relationship_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureRelationshippropPrimaryKeyColumns),
	)
	values := []interface{}{related.FeatureRelationshipID, o.FeatureRelationshippropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FeatureRelationshipID = related.FeatureRelationshipID

	if o.R == nil {
		o.R = &featureRelationshippropR{
			FeatureRelationship: related,
		}
	} else {
		o.R.FeatureRelationship = related
	}

	if related.R == nil {
		related.R = &featureRelationshipR{
			FeatureRelationshipprop: o,
		}
	} else {
		related.R.FeatureRelationshipprop = o
	}

	return nil
}

// SetType of the feature_relationshipprop to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeFeatureRelationshipprop.
func (o *FeatureRelationshipprop) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_relationshipprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureRelationshippropPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.FeatureRelationshippropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &featureRelationshippropR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeFeatureRelationshipprop: o,
		}
	} else {
		related.R.TypeFeatureRelationshipprop = o
	}

	return nil
}

// SetFeatureRelationshippropPub of the feature_relationshipprop to the related item.
// Sets o.R.FeatureRelationshippropPub to related.
// Adds o to related.R.FeatureRelationshipprop.
func (o *FeatureRelationshipprop) SetFeatureRelationshippropPub(exec boil.Executor, insert bool, related *FeatureRelationshippropPub) error {
	var err error

	if insert {
		related.FeatureRelationshippropID = o.FeatureRelationshippropID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_relationshipprop_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"feature_relationshipprop_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureRelationshippropPubPrimaryKeyColumns),
		)
		values := []interface{}{o.FeatureRelationshippropID, related.FeatureRelationshippropPubID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.FeatureRelationshippropID = o.FeatureRelationshippropID

	}

	if o.R == nil {
		o.R = &featureRelationshippropR{
			FeatureRelationshippropPub: related,
		}
	} else {
		o.R.FeatureRelationshippropPub = related
	}

	if related.R == nil {
		related.R = &featureRelationshippropPubR{
			FeatureRelationshipprop: o,
		}
	} else {
		related.R.FeatureRelationshipprop = o
	}
	return nil
}

// FeatureRelationshippropsG retrieves all records.
func FeatureRelationshippropsG(mods ...qm.QueryMod) featureRelationshippropQuery {
	return FeatureRelationshipprops(boil.GetDB(), mods...)
}

// FeatureRelationshipprops retrieves all the records using an executor.
func FeatureRelationshipprops(exec boil.Executor, mods ...qm.QueryMod) featureRelationshippropQuery {
	mods = append(mods, qm.From("\"feature_relationshipprop\""))
	return featureRelationshippropQuery{NewQuery(exec, mods...)}
}

// FindFeatureRelationshippropG retrieves a single record by ID.
func FindFeatureRelationshippropG(featureRelationshippropID int, selectCols ...string) (*FeatureRelationshipprop, error) {
	return FindFeatureRelationshipprop(boil.GetDB(), featureRelationshippropID, selectCols...)
}

// FindFeatureRelationshippropGP retrieves a single record by ID, and panics on error.
func FindFeatureRelationshippropGP(featureRelationshippropID int, selectCols ...string) *FeatureRelationshipprop {
	retobj, err := FindFeatureRelationshipprop(boil.GetDB(), featureRelationshippropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFeatureRelationshipprop retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeatureRelationshipprop(exec boil.Executor, featureRelationshippropID int, selectCols ...string) (*FeatureRelationshipprop, error) {
	featureRelationshippropObj := &FeatureRelationshipprop{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"feature_relationshipprop\" where \"feature_relationshipprop_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, featureRelationshippropID)

	err := q.Bind(featureRelationshippropObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from feature_relationshipprop")
	}

	return featureRelationshippropObj, nil
}

// FindFeatureRelationshippropP retrieves a single record by ID with an executor, and panics on error.
func FindFeatureRelationshippropP(exec boil.Executor, featureRelationshippropID int, selectCols ...string) *FeatureRelationshipprop {
	retobj, err := FindFeatureRelationshipprop(exec, featureRelationshippropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *FeatureRelationshipprop) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *FeatureRelationshipprop) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *FeatureRelationshipprop) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *FeatureRelationshipprop) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no feature_relationshipprop provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureRelationshippropColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	featureRelationshippropInsertCacheMut.RLock()
	cache, cached := featureRelationshippropInsertCache[key]
	featureRelationshippropInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			featureRelationshippropColumns,
			featureRelationshippropColumnsWithDefault,
			featureRelationshippropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(featureRelationshippropType, featureRelationshippropMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(featureRelationshippropType, featureRelationshippropMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"feature_relationshipprop\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into feature_relationshipprop")
	}

	if !cached {
		featureRelationshippropInsertCacheMut.Lock()
		featureRelationshippropInsertCache[key] = cache
		featureRelationshippropInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single FeatureRelationshipprop record. See Update for
// whitelist behavior description.
func (o *FeatureRelationshipprop) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single FeatureRelationshipprop record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *FeatureRelationshipprop) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the FeatureRelationshipprop, and panics on error.
// See Update for whitelist behavior description.
func (o *FeatureRelationshipprop) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the FeatureRelationshipprop.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *FeatureRelationshipprop) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	featureRelationshippropUpdateCacheMut.RLock()
	cache, cached := featureRelationshippropUpdateCache[key]
	featureRelationshippropUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(featureRelationshippropColumns, featureRelationshippropPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update feature_relationshipprop, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"feature_relationshipprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, featureRelationshippropPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(featureRelationshippropType, featureRelationshippropMapping, append(wl, featureRelationshippropPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update feature_relationshipprop row")
	}

	if !cached {
		featureRelationshippropUpdateCacheMut.Lock()
		featureRelationshippropUpdateCache[key] = cache
		featureRelationshippropUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q featureRelationshippropQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q featureRelationshippropQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for feature_relationshipprop")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FeatureRelationshippropSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FeatureRelationshippropSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FeatureRelationshippropSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FeatureRelationshippropSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureRelationshippropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"feature_relationshipprop\" SET %s WHERE (\"feature_relationshipprop_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featureRelationshippropPrimaryKeyColumns), len(colNames)+1, len(featureRelationshippropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in featureRelationshipprop slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *FeatureRelationshipprop) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *FeatureRelationshipprop) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *FeatureRelationshipprop) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *FeatureRelationshipprop) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no feature_relationshipprop provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureRelationshippropColumnsWithDefault, o)

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

	featureRelationshippropUpsertCacheMut.RLock()
	cache, cached := featureRelationshippropUpsertCache[key]
	featureRelationshippropUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			featureRelationshippropColumns,
			featureRelationshippropColumnsWithDefault,
			featureRelationshippropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			featureRelationshippropColumns,
			featureRelationshippropPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert feature_relationshipprop, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(featureRelationshippropPrimaryKeyColumns))
			copy(conflict, featureRelationshippropPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"feature_relationshipprop\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(featureRelationshippropType, featureRelationshippropMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(featureRelationshippropType, featureRelationshippropMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for feature_relationshipprop")
	}

	if !cached {
		featureRelationshippropUpsertCacheMut.Lock()
		featureRelationshippropUpsertCache[key] = cache
		featureRelationshippropUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single FeatureRelationshipprop record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeatureRelationshipprop) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single FeatureRelationshipprop record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *FeatureRelationshipprop) DeleteG() error {
	if o == nil {
		return errors.New("models: no FeatureRelationshipprop provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single FeatureRelationshipprop record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeatureRelationshipprop) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single FeatureRelationshipprop record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FeatureRelationshipprop) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no FeatureRelationshipprop provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), featureRelationshippropPrimaryKeyMapping)
	sql := "DELETE FROM \"feature_relationshipprop\" WHERE \"feature_relationshipprop_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from feature_relationshipprop")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q featureRelationshippropQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q featureRelationshippropQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no featureRelationshippropQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from feature_relationshipprop")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FeatureRelationshippropSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FeatureRelationshippropSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no FeatureRelationshipprop slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FeatureRelationshippropSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FeatureRelationshippropSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no FeatureRelationshipprop slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(featureRelationshippropBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureRelationshippropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"feature_relationshipprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featureRelationshippropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featureRelationshippropPrimaryKeyColumns), 1, len(featureRelationshippropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from featureRelationshipprop slice")
	}

	if len(featureRelationshippropAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *FeatureRelationshipprop) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *FeatureRelationshipprop) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *FeatureRelationshipprop) ReloadG() error {
	if o == nil {
		return errors.New("models: no FeatureRelationshipprop provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FeatureRelationshipprop) Reload(exec boil.Executor) error {
	ret, err := FindFeatureRelationshipprop(exec, o.FeatureRelationshippropID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureRelationshippropSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureRelationshippropSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureRelationshippropSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty FeatureRelationshippropSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureRelationshippropSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	featureRelationshipprops := FeatureRelationshippropSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureRelationshippropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"feature_relationshipprop\".* FROM \"feature_relationshipprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featureRelationshippropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(featureRelationshippropPrimaryKeyColumns), 1, len(featureRelationshippropPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&featureRelationshipprops)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FeatureRelationshippropSlice")
	}

	*o = featureRelationshipprops

	return nil
}

// FeatureRelationshippropExists checks if the FeatureRelationshipprop row exists.
func FeatureRelationshippropExists(exec boil.Executor, featureRelationshippropID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"feature_relationshipprop\" where \"feature_relationshipprop_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, featureRelationshippropID)
	}

	row := exec.QueryRow(sql, featureRelationshippropID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if feature_relationshipprop exists")
	}

	return exists, nil
}

// FeatureRelationshippropExistsG checks if the FeatureRelationshipprop row exists.
func FeatureRelationshippropExistsG(featureRelationshippropID int) (bool, error) {
	return FeatureRelationshippropExists(boil.GetDB(), featureRelationshippropID)
}

// FeatureRelationshippropExistsGP checks if the FeatureRelationshipprop row exists. Panics on error.
func FeatureRelationshippropExistsGP(featureRelationshippropID int) bool {
	e, err := FeatureRelationshippropExists(boil.GetDB(), featureRelationshippropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FeatureRelationshippropExistsP checks if the FeatureRelationshipprop row exists. Panics on error.
func FeatureRelationshippropExistsP(exec boil.Executor, featureRelationshippropID int) bool {
	e, err := FeatureRelationshippropExists(exec, featureRelationshippropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
