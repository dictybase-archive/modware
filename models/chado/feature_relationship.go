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

// FeatureRelationship is an object representing the database table.
type FeatureRelationship struct {
	FeatureRelationshipID int         `boil:"feature_relationship_id" json:"feature_relationship_id" toml:"feature_relationship_id" yaml:"feature_relationship_id"`
	SubjectID             int         `boil:"subject_id" json:"subject_id" toml:"subject_id" yaml:"subject_id"`
	ObjectID              int         `boil:"object_id" json:"object_id" toml:"object_id" yaml:"object_id"`
	TypeID                int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Value                 null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	Rank                  int         `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *featureRelationshipR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L featureRelationshipL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// featureRelationshipR is where relationships are stored.
type featureRelationshipR struct {
	Type                    *Cvterm
	Object                  *Feature
	Subject                 *Feature
	FeatureRelationshipprop *FeatureRelationshipprop
	FeatureRelationshipPub  *FeatureRelationshipPub
}

// featureRelationshipL is where Load methods for each relationship are stored.
type featureRelationshipL struct{}

var (
	featureRelationshipColumns               = []string{"feature_relationship_id", "subject_id", "object_id", "type_id", "value", "rank"}
	featureRelationshipColumnsWithoutDefault = []string{"subject_id", "object_id", "type_id", "value"}
	featureRelationshipColumnsWithDefault    = []string{"feature_relationship_id", "rank"}
	featureRelationshipPrimaryKeyColumns     = []string{"feature_relationship_id"}
)

type (
	// FeatureRelationshipSlice is an alias for a slice of pointers to FeatureRelationship.
	// This should generally be used opposed to []FeatureRelationship.
	FeatureRelationshipSlice []*FeatureRelationship
	// FeatureRelationshipHook is the signature for custom FeatureRelationship hook methods
	FeatureRelationshipHook func(boil.Executor, *FeatureRelationship) error

	featureRelationshipQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	featureRelationshipType                 = reflect.TypeOf(&FeatureRelationship{})
	featureRelationshipMapping              = queries.MakeStructMapping(featureRelationshipType)
	featureRelationshipPrimaryKeyMapping, _ = queries.BindMapping(featureRelationshipType, featureRelationshipMapping, featureRelationshipPrimaryKeyColumns)
	featureRelationshipInsertCacheMut       sync.RWMutex
	featureRelationshipInsertCache          = make(map[string]insertCache)
	featureRelationshipUpdateCacheMut       sync.RWMutex
	featureRelationshipUpdateCache          = make(map[string]updateCache)
	featureRelationshipUpsertCacheMut       sync.RWMutex
	featureRelationshipUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var featureRelationshipBeforeInsertHooks []FeatureRelationshipHook
var featureRelationshipBeforeUpdateHooks []FeatureRelationshipHook
var featureRelationshipBeforeDeleteHooks []FeatureRelationshipHook
var featureRelationshipBeforeUpsertHooks []FeatureRelationshipHook

var featureRelationshipAfterInsertHooks []FeatureRelationshipHook
var featureRelationshipAfterSelectHooks []FeatureRelationshipHook
var featureRelationshipAfterUpdateHooks []FeatureRelationshipHook
var featureRelationshipAfterDeleteHooks []FeatureRelationshipHook
var featureRelationshipAfterUpsertHooks []FeatureRelationshipHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FeatureRelationship) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshipBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FeatureRelationship) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshipBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FeatureRelationship) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshipBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FeatureRelationship) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshipBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FeatureRelationship) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshipAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FeatureRelationship) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshipAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FeatureRelationship) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshipAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FeatureRelationship) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshipAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FeatureRelationship) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureRelationshipAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFeatureRelationshipHook registers your hook function for all future operations.
func AddFeatureRelationshipHook(hookPoint boil.HookPoint, featureRelationshipHook FeatureRelationshipHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		featureRelationshipBeforeInsertHooks = append(featureRelationshipBeforeInsertHooks, featureRelationshipHook)
	case boil.BeforeUpdateHook:
		featureRelationshipBeforeUpdateHooks = append(featureRelationshipBeforeUpdateHooks, featureRelationshipHook)
	case boil.BeforeDeleteHook:
		featureRelationshipBeforeDeleteHooks = append(featureRelationshipBeforeDeleteHooks, featureRelationshipHook)
	case boil.BeforeUpsertHook:
		featureRelationshipBeforeUpsertHooks = append(featureRelationshipBeforeUpsertHooks, featureRelationshipHook)
	case boil.AfterInsertHook:
		featureRelationshipAfterInsertHooks = append(featureRelationshipAfterInsertHooks, featureRelationshipHook)
	case boil.AfterSelectHook:
		featureRelationshipAfterSelectHooks = append(featureRelationshipAfterSelectHooks, featureRelationshipHook)
	case boil.AfterUpdateHook:
		featureRelationshipAfterUpdateHooks = append(featureRelationshipAfterUpdateHooks, featureRelationshipHook)
	case boil.AfterDeleteHook:
		featureRelationshipAfterDeleteHooks = append(featureRelationshipAfterDeleteHooks, featureRelationshipHook)
	case boil.AfterUpsertHook:
		featureRelationshipAfterUpsertHooks = append(featureRelationshipAfterUpsertHooks, featureRelationshipHook)
	}
}

// OneP returns a single featureRelationship record from the query, and panics on error.
func (q featureRelationshipQuery) OneP() *FeatureRelationship {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single featureRelationship record from the query.
func (q featureRelationshipQuery) One() (*FeatureRelationship, error) {
	o := &FeatureRelationship{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for feature_relationship")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all FeatureRelationship records from the query, and panics on error.
func (q featureRelationshipQuery) AllP() FeatureRelationshipSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all FeatureRelationship records from the query.
func (q featureRelationshipQuery) All() (FeatureRelationshipSlice, error) {
	var o FeatureRelationshipSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to FeatureRelationship slice")
	}

	if len(featureRelationshipAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all FeatureRelationship records in the query, and panics on error.
func (q featureRelationshipQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all FeatureRelationship records in the query.
func (q featureRelationshipQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count feature_relationship rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q featureRelationshipQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q featureRelationshipQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if feature_relationship exists")
	}

	return count > 0, nil
}

// TypeG pointed to by the foreign key.
func (o *FeatureRelationship) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *FeatureRelationship) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// ObjectG pointed to by the foreign key.
func (o *FeatureRelationship) ObjectG(mods ...qm.QueryMod) featureQuery {
	return o.Object(boil.GetDB(), mods...)
}

// Object pointed to by the foreign key.
func (o *FeatureRelationship) Object(exec boil.Executor, mods ...qm.QueryMod) featureQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.ObjectID),
	}

	queryMods = append(queryMods, mods...)

	query := Features(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature\"")

	return query
}

// SubjectG pointed to by the foreign key.
func (o *FeatureRelationship) SubjectG(mods ...qm.QueryMod) featureQuery {
	return o.Subject(boil.GetDB(), mods...)
}

// Subject pointed to by the foreign key.
func (o *FeatureRelationship) Subject(exec boil.Executor, mods ...qm.QueryMod) featureQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.SubjectID),
	}

	queryMods = append(queryMods, mods...)

	query := Features(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature\"")

	return query
}

// FeatureRelationshippropG pointed to by the foreign key.
func (o *FeatureRelationship) FeatureRelationshippropG(mods ...qm.QueryMod) featureRelationshippropQuery {
	return o.FeatureRelationshipprop(boil.GetDB(), mods...)
}

// FeatureRelationshipprop pointed to by the foreign key.
func (o *FeatureRelationship) FeatureRelationshipprop(exec boil.Executor, mods ...qm.QueryMod) featureRelationshippropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_relationship_id=$1", o.FeatureRelationshipID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureRelationshipprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_relationshipprop\"")

	return query
}

// FeatureRelationshipPubG pointed to by the foreign key.
func (o *FeatureRelationship) FeatureRelationshipPubG(mods ...qm.QueryMod) featureRelationshipPubQuery {
	return o.FeatureRelationshipPub(boil.GetDB(), mods...)
}

// FeatureRelationshipPub pointed to by the foreign key.
func (o *FeatureRelationship) FeatureRelationshipPub(exec boil.Executor, mods ...qm.QueryMod) featureRelationshipPubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_relationship_id=$1", o.FeatureRelationshipID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureRelationshipPubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_relationship_pub\"")

	return query
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureRelationshipL) LoadType(e boil.Executor, singular bool, maybeFeatureRelationship interface{}) error {
	var slice []*FeatureRelationship
	var object *FeatureRelationship

	count := 1
	if singular {
		object = maybeFeatureRelationship.(*FeatureRelationship)
	} else {
		slice = *maybeFeatureRelationship.(*FeatureRelationshipSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureRelationshipR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &featureRelationshipR{}
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

	if len(featureRelationshipAfterSelectHooks) != 0 {
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

// LoadObject allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureRelationshipL) LoadObject(e boil.Executor, singular bool, maybeFeatureRelationship interface{}) error {
	var slice []*FeatureRelationship
	var object *FeatureRelationship

	count := 1
	if singular {
		object = maybeFeatureRelationship.(*FeatureRelationship)
	} else {
		slice = *maybeFeatureRelationship.(*FeatureRelationshipSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureRelationshipR{}
		args[0] = object.ObjectID
	} else {
		for i, obj := range slice {
			obj.R = &featureRelationshipR{}
			args[i] = obj.ObjectID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature\" where \"feature_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Feature")
	}
	defer results.Close()

	var resultSlice []*Feature
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Feature")
	}

	if len(featureRelationshipAfterSelectHooks) != 0 {
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
			if local.ObjectID == foreign.FeatureID {
				local.R.Object = foreign
				break
			}
		}
	}

	return nil
}

// LoadSubject allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureRelationshipL) LoadSubject(e boil.Executor, singular bool, maybeFeatureRelationship interface{}) error {
	var slice []*FeatureRelationship
	var object *FeatureRelationship

	count := 1
	if singular {
		object = maybeFeatureRelationship.(*FeatureRelationship)
	} else {
		slice = *maybeFeatureRelationship.(*FeatureRelationshipSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureRelationshipR{}
		args[0] = object.SubjectID
	} else {
		for i, obj := range slice {
			obj.R = &featureRelationshipR{}
			args[i] = obj.SubjectID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature\" where \"feature_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Feature")
	}
	defer results.Close()

	var resultSlice []*Feature
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Feature")
	}

	if len(featureRelationshipAfterSelectHooks) != 0 {
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
			if local.SubjectID == foreign.FeatureID {
				local.R.Subject = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeatureRelationshipprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureRelationshipL) LoadFeatureRelationshipprop(e boil.Executor, singular bool, maybeFeatureRelationship interface{}) error {
	var slice []*FeatureRelationship
	var object *FeatureRelationship

	count := 1
	if singular {
		object = maybeFeatureRelationship.(*FeatureRelationship)
	} else {
		slice = *maybeFeatureRelationship.(*FeatureRelationshipSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureRelationshipR{}
		args[0] = object.FeatureRelationshipID
	} else {
		for i, obj := range slice {
			obj.R = &featureRelationshipR{}
			args[i] = obj.FeatureRelationshipID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_relationshipprop\" where \"feature_relationship_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureRelationshipprop")
	}
	defer results.Close()

	var resultSlice []*FeatureRelationshipprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureRelationshipprop")
	}

	if len(featureRelationshipAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeatureRelationshipprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureRelationshipID == foreign.FeatureRelationshipID {
				local.R.FeatureRelationshipprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeatureRelationshipPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureRelationshipL) LoadFeatureRelationshipPub(e boil.Executor, singular bool, maybeFeatureRelationship interface{}) error {
	var slice []*FeatureRelationship
	var object *FeatureRelationship

	count := 1
	if singular {
		object = maybeFeatureRelationship.(*FeatureRelationship)
	} else {
		slice = *maybeFeatureRelationship.(*FeatureRelationshipSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureRelationshipR{}
		args[0] = object.FeatureRelationshipID
	} else {
		for i, obj := range slice {
			obj.R = &featureRelationshipR{}
			args[i] = obj.FeatureRelationshipID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_relationship_pub\" where \"feature_relationship_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureRelationshipPub")
	}
	defer results.Close()

	var resultSlice []*FeatureRelationshipPub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureRelationshipPub")
	}

	if len(featureRelationshipAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeatureRelationshipPub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureRelationshipID == foreign.FeatureRelationshipID {
				local.R.FeatureRelationshipPub = foreign
				break
			}
		}
	}

	return nil
}

// SetType of the feature_relationship to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeFeatureRelationship.
func (o *FeatureRelationship) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_relationship\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureRelationshipPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.FeatureRelationshipID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &featureRelationshipR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeFeatureRelationship: o,
		}
	} else {
		related.R.TypeFeatureRelationship = o
	}

	return nil
}

// SetObject of the feature_relationship to the related item.
// Sets o.R.Object to related.
// Adds o to related.R.ObjectFeatureRelationship.
func (o *FeatureRelationship) SetObject(exec boil.Executor, insert bool, related *Feature) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_relationship\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"object_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureRelationshipPrimaryKeyColumns),
	)
	values := []interface{}{related.FeatureID, o.FeatureRelationshipID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ObjectID = related.FeatureID

	if o.R == nil {
		o.R = &featureRelationshipR{
			Object: related,
		}
	} else {
		o.R.Object = related
	}

	if related.R == nil {
		related.R = &featureR{
			ObjectFeatureRelationship: o,
		}
	} else {
		related.R.ObjectFeatureRelationship = o
	}

	return nil
}

// SetSubject of the feature_relationship to the related item.
// Sets o.R.Subject to related.
// Adds o to related.R.SubjectFeatureRelationship.
func (o *FeatureRelationship) SetSubject(exec boil.Executor, insert bool, related *Feature) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_relationship\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"subject_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureRelationshipPrimaryKeyColumns),
	)
	values := []interface{}{related.FeatureID, o.FeatureRelationshipID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.SubjectID = related.FeatureID

	if o.R == nil {
		o.R = &featureRelationshipR{
			Subject: related,
		}
	} else {
		o.R.Subject = related
	}

	if related.R == nil {
		related.R = &featureR{
			SubjectFeatureRelationship: o,
		}
	} else {
		related.R.SubjectFeatureRelationship = o
	}

	return nil
}

// SetFeatureRelationshipprop of the feature_relationship to the related item.
// Sets o.R.FeatureRelationshipprop to related.
// Adds o to related.R.FeatureRelationship.
func (o *FeatureRelationship) SetFeatureRelationshipprop(exec boil.Executor, insert bool, related *FeatureRelationshipprop) error {
	var err error

	if insert {
		related.FeatureRelationshipID = o.FeatureRelationshipID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_relationshipprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"feature_relationship_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureRelationshippropPrimaryKeyColumns),
		)
		values := []interface{}{o.FeatureRelationshipID, related.FeatureRelationshippropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.FeatureRelationshipID = o.FeatureRelationshipID

	}

	if o.R == nil {
		o.R = &featureRelationshipR{
			FeatureRelationshipprop: related,
		}
	} else {
		o.R.FeatureRelationshipprop = related
	}

	if related.R == nil {
		related.R = &featureRelationshippropR{
			FeatureRelationship: o,
		}
	} else {
		related.R.FeatureRelationship = o
	}
	return nil
}

// SetFeatureRelationshipPub of the feature_relationship to the related item.
// Sets o.R.FeatureRelationshipPub to related.
// Adds o to related.R.FeatureRelationship.
func (o *FeatureRelationship) SetFeatureRelationshipPub(exec boil.Executor, insert bool, related *FeatureRelationshipPub) error {
	var err error

	if insert {
		related.FeatureRelationshipID = o.FeatureRelationshipID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_relationship_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"feature_relationship_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureRelationshipPubPrimaryKeyColumns),
		)
		values := []interface{}{o.FeatureRelationshipID, related.FeatureRelationshipPubID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.FeatureRelationshipID = o.FeatureRelationshipID

	}

	if o.R == nil {
		o.R = &featureRelationshipR{
			FeatureRelationshipPub: related,
		}
	} else {
		o.R.FeatureRelationshipPub = related
	}

	if related.R == nil {
		related.R = &featureRelationshipPubR{
			FeatureRelationship: o,
		}
	} else {
		related.R.FeatureRelationship = o
	}
	return nil
}

// FeatureRelationshipsG retrieves all records.
func FeatureRelationshipsG(mods ...qm.QueryMod) featureRelationshipQuery {
	return FeatureRelationships(boil.GetDB(), mods...)
}

// FeatureRelationships retrieves all the records using an executor.
func FeatureRelationships(exec boil.Executor, mods ...qm.QueryMod) featureRelationshipQuery {
	mods = append(mods, qm.From("\"feature_relationship\""))
	return featureRelationshipQuery{NewQuery(exec, mods...)}
}

// FindFeatureRelationshipG retrieves a single record by ID.
func FindFeatureRelationshipG(featureRelationshipID int, selectCols ...string) (*FeatureRelationship, error) {
	return FindFeatureRelationship(boil.GetDB(), featureRelationshipID, selectCols...)
}

// FindFeatureRelationshipGP retrieves a single record by ID, and panics on error.
func FindFeatureRelationshipGP(featureRelationshipID int, selectCols ...string) *FeatureRelationship {
	retobj, err := FindFeatureRelationship(boil.GetDB(), featureRelationshipID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFeatureRelationship retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeatureRelationship(exec boil.Executor, featureRelationshipID int, selectCols ...string) (*FeatureRelationship, error) {
	featureRelationshipObj := &FeatureRelationship{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"feature_relationship\" where \"feature_relationship_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, featureRelationshipID)

	err := q.Bind(featureRelationshipObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from feature_relationship")
	}

	return featureRelationshipObj, nil
}

// FindFeatureRelationshipP retrieves a single record by ID with an executor, and panics on error.
func FindFeatureRelationshipP(exec boil.Executor, featureRelationshipID int, selectCols ...string) *FeatureRelationship {
	retobj, err := FindFeatureRelationship(exec, featureRelationshipID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *FeatureRelationship) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *FeatureRelationship) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *FeatureRelationship) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *FeatureRelationship) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no feature_relationship provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureRelationshipColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	featureRelationshipInsertCacheMut.RLock()
	cache, cached := featureRelationshipInsertCache[key]
	featureRelationshipInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			featureRelationshipColumns,
			featureRelationshipColumnsWithDefault,
			featureRelationshipColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(featureRelationshipType, featureRelationshipMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(featureRelationshipType, featureRelationshipMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"feature_relationship\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into feature_relationship")
	}

	if !cached {
		featureRelationshipInsertCacheMut.Lock()
		featureRelationshipInsertCache[key] = cache
		featureRelationshipInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single FeatureRelationship record. See Update for
// whitelist behavior description.
func (o *FeatureRelationship) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single FeatureRelationship record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *FeatureRelationship) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the FeatureRelationship, and panics on error.
// See Update for whitelist behavior description.
func (o *FeatureRelationship) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the FeatureRelationship.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *FeatureRelationship) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	featureRelationshipUpdateCacheMut.RLock()
	cache, cached := featureRelationshipUpdateCache[key]
	featureRelationshipUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(featureRelationshipColumns, featureRelationshipPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update feature_relationship, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"feature_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, featureRelationshipPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(featureRelationshipType, featureRelationshipMapping, append(wl, featureRelationshipPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update feature_relationship row")
	}

	if !cached {
		featureRelationshipUpdateCacheMut.Lock()
		featureRelationshipUpdateCache[key] = cache
		featureRelationshipUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q featureRelationshipQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q featureRelationshipQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for feature_relationship")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FeatureRelationshipSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FeatureRelationshipSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FeatureRelationshipSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FeatureRelationshipSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureRelationshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"feature_relationship\" SET %s WHERE (\"feature_relationship_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featureRelationshipPrimaryKeyColumns), len(colNames)+1, len(featureRelationshipPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in featureRelationship slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *FeatureRelationship) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *FeatureRelationship) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *FeatureRelationship) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *FeatureRelationship) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no feature_relationship provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureRelationshipColumnsWithDefault, o)

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

	featureRelationshipUpsertCacheMut.RLock()
	cache, cached := featureRelationshipUpsertCache[key]
	featureRelationshipUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			featureRelationshipColumns,
			featureRelationshipColumnsWithDefault,
			featureRelationshipColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			featureRelationshipColumns,
			featureRelationshipPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert feature_relationship, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(featureRelationshipPrimaryKeyColumns))
			copy(conflict, featureRelationshipPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"feature_relationship\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(featureRelationshipType, featureRelationshipMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(featureRelationshipType, featureRelationshipMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for feature_relationship")
	}

	if !cached {
		featureRelationshipUpsertCacheMut.Lock()
		featureRelationshipUpsertCache[key] = cache
		featureRelationshipUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single FeatureRelationship record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeatureRelationship) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single FeatureRelationship record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *FeatureRelationship) DeleteG() error {
	if o == nil {
		return errors.New("chado: no FeatureRelationship provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single FeatureRelationship record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeatureRelationship) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single FeatureRelationship record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FeatureRelationship) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no FeatureRelationship provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), featureRelationshipPrimaryKeyMapping)
	sql := "DELETE FROM \"feature_relationship\" WHERE \"feature_relationship_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from feature_relationship")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q featureRelationshipQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q featureRelationshipQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no featureRelationshipQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from feature_relationship")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FeatureRelationshipSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FeatureRelationshipSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no FeatureRelationship slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FeatureRelationshipSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FeatureRelationshipSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no FeatureRelationship slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(featureRelationshipBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureRelationshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"feature_relationship\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featureRelationshipPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featureRelationshipPrimaryKeyColumns), 1, len(featureRelationshipPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from featureRelationship slice")
	}

	if len(featureRelationshipAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *FeatureRelationship) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *FeatureRelationship) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *FeatureRelationship) ReloadG() error {
	if o == nil {
		return errors.New("chado: no FeatureRelationship provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FeatureRelationship) Reload(exec boil.Executor) error {
	ret, err := FindFeatureRelationship(exec, o.FeatureRelationshipID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureRelationshipSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureRelationshipSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureRelationshipSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty FeatureRelationshipSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureRelationshipSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	featureRelationships := FeatureRelationshipSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureRelationshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"feature_relationship\".* FROM \"feature_relationship\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featureRelationshipPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(featureRelationshipPrimaryKeyColumns), 1, len(featureRelationshipPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&featureRelationships)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in FeatureRelationshipSlice")
	}

	*o = featureRelationships

	return nil
}

// FeatureRelationshipExists checks if the FeatureRelationship row exists.
func FeatureRelationshipExists(exec boil.Executor, featureRelationshipID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"feature_relationship\" where \"feature_relationship_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, featureRelationshipID)
	}

	row := exec.QueryRow(sql, featureRelationshipID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if feature_relationship exists")
	}

	return exists, nil
}

// FeatureRelationshipExistsG checks if the FeatureRelationship row exists.
func FeatureRelationshipExistsG(featureRelationshipID int) (bool, error) {
	return FeatureRelationshipExists(boil.GetDB(), featureRelationshipID)
}

// FeatureRelationshipExistsGP checks if the FeatureRelationship row exists. Panics on error.
func FeatureRelationshipExistsGP(featureRelationshipID int) bool {
	e, err := FeatureRelationshipExists(boil.GetDB(), featureRelationshipID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FeatureRelationshipExistsP checks if the FeatureRelationship row exists. Panics on error.
func FeatureRelationshipExistsP(exec boil.Executor, featureRelationshipID int) bool {
	e, err := FeatureRelationshipExists(exec, featureRelationshipID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
