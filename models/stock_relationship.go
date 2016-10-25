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

// StockRelationship is an object representing the database table.
type StockRelationship struct {
	StockRelationshipID int         `boil:"stock_relationship_id" json:"stock_relationship_id" toml:"stock_relationship_id" yaml:"stock_relationship_id"`
	SubjectID           int         `boil:"subject_id" json:"subject_id" toml:"subject_id" yaml:"subject_id"`
	ObjectID            int         `boil:"object_id" json:"object_id" toml:"object_id" yaml:"object_id"`
	TypeID              int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Value               null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	Rank                int         `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *stockRelationshipR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockRelationshipL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// stockRelationshipR is where relationships are stored.
type stockRelationshipR struct {
	Type                     *Cvterm
	Object                   *Stock
	Subject                  *Stock
	StockRelationshipPub     *StockRelationshipPub
	StockRelationshipCvterms StockRelationshipCvtermSlice
}

// stockRelationshipL is where Load methods for each relationship are stored.
type stockRelationshipL struct{}

var (
	stockRelationshipColumns               = []string{"stock_relationship_id", "subject_id", "object_id", "type_id", "value", "rank"}
	stockRelationshipColumnsWithoutDefault = []string{"subject_id", "object_id", "type_id", "value"}
	stockRelationshipColumnsWithDefault    = []string{"stock_relationship_id", "rank"}
	stockRelationshipPrimaryKeyColumns     = []string{"stock_relationship_id"}
)

type (
	// StockRelationshipSlice is an alias for a slice of pointers to StockRelationship.
	// This should generally be used opposed to []StockRelationship.
	StockRelationshipSlice []*StockRelationship
	// StockRelationshipHook is the signature for custom StockRelationship hook methods
	StockRelationshipHook func(boil.Executor, *StockRelationship) error

	stockRelationshipQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockRelationshipType                 = reflect.TypeOf(&StockRelationship{})
	stockRelationshipMapping              = queries.MakeStructMapping(stockRelationshipType)
	stockRelationshipPrimaryKeyMapping, _ = queries.BindMapping(stockRelationshipType, stockRelationshipMapping, stockRelationshipPrimaryKeyColumns)
	stockRelationshipInsertCacheMut       sync.RWMutex
	stockRelationshipInsertCache          = make(map[string]insertCache)
	stockRelationshipUpdateCacheMut       sync.RWMutex
	stockRelationshipUpdateCache          = make(map[string]updateCache)
	stockRelationshipUpsertCacheMut       sync.RWMutex
	stockRelationshipUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var stockRelationshipBeforeInsertHooks []StockRelationshipHook
var stockRelationshipBeforeUpdateHooks []StockRelationshipHook
var stockRelationshipBeforeDeleteHooks []StockRelationshipHook
var stockRelationshipBeforeUpsertHooks []StockRelationshipHook

var stockRelationshipAfterInsertHooks []StockRelationshipHook
var stockRelationshipAfterSelectHooks []StockRelationshipHook
var stockRelationshipAfterUpdateHooks []StockRelationshipHook
var stockRelationshipAfterDeleteHooks []StockRelationshipHook
var stockRelationshipAfterUpsertHooks []StockRelationshipHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *StockRelationship) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *StockRelationship) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *StockRelationship) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *StockRelationship) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *StockRelationship) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *StockRelationship) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *StockRelationship) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *StockRelationship) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *StockRelationship) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStockRelationshipHook registers your hook function for all future operations.
func AddStockRelationshipHook(hookPoint boil.HookPoint, stockRelationshipHook StockRelationshipHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stockRelationshipBeforeInsertHooks = append(stockRelationshipBeforeInsertHooks, stockRelationshipHook)
	case boil.BeforeUpdateHook:
		stockRelationshipBeforeUpdateHooks = append(stockRelationshipBeforeUpdateHooks, stockRelationshipHook)
	case boil.BeforeDeleteHook:
		stockRelationshipBeforeDeleteHooks = append(stockRelationshipBeforeDeleteHooks, stockRelationshipHook)
	case boil.BeforeUpsertHook:
		stockRelationshipBeforeUpsertHooks = append(stockRelationshipBeforeUpsertHooks, stockRelationshipHook)
	case boil.AfterInsertHook:
		stockRelationshipAfterInsertHooks = append(stockRelationshipAfterInsertHooks, stockRelationshipHook)
	case boil.AfterSelectHook:
		stockRelationshipAfterSelectHooks = append(stockRelationshipAfterSelectHooks, stockRelationshipHook)
	case boil.AfterUpdateHook:
		stockRelationshipAfterUpdateHooks = append(stockRelationshipAfterUpdateHooks, stockRelationshipHook)
	case boil.AfterDeleteHook:
		stockRelationshipAfterDeleteHooks = append(stockRelationshipAfterDeleteHooks, stockRelationshipHook)
	case boil.AfterUpsertHook:
		stockRelationshipAfterUpsertHooks = append(stockRelationshipAfterUpsertHooks, stockRelationshipHook)
	}
}

// OneP returns a single stockRelationship record from the query, and panics on error.
func (q stockRelationshipQuery) OneP() *StockRelationship {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single stockRelationship record from the query.
func (q stockRelationshipQuery) One() (*StockRelationship, error) {
	o := &StockRelationship{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stock_relationship")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all StockRelationship records from the query, and panics on error.
func (q stockRelationshipQuery) AllP() StockRelationshipSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all StockRelationship records from the query.
func (q stockRelationshipQuery) All() (StockRelationshipSlice, error) {
	var o StockRelationshipSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to StockRelationship slice")
	}

	if len(stockRelationshipAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all StockRelationship records in the query, and panics on error.
func (q stockRelationshipQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all StockRelationship records in the query.
func (q stockRelationshipQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stock_relationship rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q stockRelationshipQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q stockRelationshipQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stock_relationship exists")
	}

	return count > 0, nil
}

// TypeG pointed to by the foreign key.
func (o *StockRelationship) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *StockRelationship) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// ObjectG pointed to by the foreign key.
func (o *StockRelationship) ObjectG(mods ...qm.QueryMod) stockQuery {
	return o.Object(boil.GetDB(), mods...)
}

// Object pointed to by the foreign key.
func (o *StockRelationship) Object(exec boil.Executor, mods ...qm.QueryMod) stockQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_id=$1", o.ObjectID),
	}

	queryMods = append(queryMods, mods...)

	query := Stocks(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock\"")

	return query
}

// SubjectG pointed to by the foreign key.
func (o *StockRelationship) SubjectG(mods ...qm.QueryMod) stockQuery {
	return o.Subject(boil.GetDB(), mods...)
}

// Subject pointed to by the foreign key.
func (o *StockRelationship) Subject(exec boil.Executor, mods ...qm.QueryMod) stockQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_id=$1", o.SubjectID),
	}

	queryMods = append(queryMods, mods...)

	query := Stocks(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock\"")

	return query
}

// StockRelationshipPubG pointed to by the foreign key.
func (o *StockRelationship) StockRelationshipPubG(mods ...qm.QueryMod) stockRelationshipPubQuery {
	return o.StockRelationshipPub(boil.GetDB(), mods...)
}

// StockRelationshipPub pointed to by the foreign key.
func (o *StockRelationship) StockRelationshipPub(exec boil.Executor, mods ...qm.QueryMod) stockRelationshipPubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_relationship_id=$1", o.StockRelationshipID),
	}

	queryMods = append(queryMods, mods...)

	query := StockRelationshipPubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_relationship_pub\"")

	return query
}

// StockRelationshipCvtermsG retrieves all the stock_relationship_cvterm's stock relationship cvterm.
func (o *StockRelationship) StockRelationshipCvtermsG(mods ...qm.QueryMod) stockRelationshipCvtermQuery {
	return o.StockRelationshipCvterms(boil.GetDB(), mods...)
}

// StockRelationshipCvterms retrieves all the stock_relationship_cvterm's stock relationship cvterm with an executor.
func (o *StockRelationship) StockRelationshipCvterms(exec boil.Executor, mods ...qm.QueryMod) stockRelationshipCvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"stock_relationship_id\"=$1", o.StockRelationshipID),
	)

	query := StockRelationshipCvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_relationship_cvterm\" as \"a\"")
	return query
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockRelationshipL) LoadType(e boil.Executor, singular bool, maybeStockRelationship interface{}) error {
	var slice []*StockRelationship
	var object *StockRelationship

	count := 1
	if singular {
		object = maybeStockRelationship.(*StockRelationship)
	} else {
		slice = *maybeStockRelationship.(*StockRelationshipSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockRelationshipR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &stockRelationshipR{}
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

	if len(stockRelationshipAfterSelectHooks) != 0 {
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
func (stockRelationshipL) LoadObject(e boil.Executor, singular bool, maybeStockRelationship interface{}) error {
	var slice []*StockRelationship
	var object *StockRelationship

	count := 1
	if singular {
		object = maybeStockRelationship.(*StockRelationship)
	} else {
		slice = *maybeStockRelationship.(*StockRelationshipSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockRelationshipR{}
		args[0] = object.ObjectID
	} else {
		for i, obj := range slice {
			obj.R = &stockRelationshipR{}
			args[i] = obj.ObjectID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock\" where \"stock_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Stock")
	}
	defer results.Close()

	var resultSlice []*Stock
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Stock")
	}

	if len(stockRelationshipAfterSelectHooks) != 0 {
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
			if local.ObjectID == foreign.StockID {
				local.R.Object = foreign
				break
			}
		}
	}

	return nil
}

// LoadSubject allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockRelationshipL) LoadSubject(e boil.Executor, singular bool, maybeStockRelationship interface{}) error {
	var slice []*StockRelationship
	var object *StockRelationship

	count := 1
	if singular {
		object = maybeStockRelationship.(*StockRelationship)
	} else {
		slice = *maybeStockRelationship.(*StockRelationshipSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockRelationshipR{}
		args[0] = object.SubjectID
	} else {
		for i, obj := range slice {
			obj.R = &stockRelationshipR{}
			args[i] = obj.SubjectID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock\" where \"stock_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Stock")
	}
	defer results.Close()

	var resultSlice []*Stock
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Stock")
	}

	if len(stockRelationshipAfterSelectHooks) != 0 {
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
			if local.SubjectID == foreign.StockID {
				local.R.Subject = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockRelationshipPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockRelationshipL) LoadStockRelationshipPub(e boil.Executor, singular bool, maybeStockRelationship interface{}) error {
	var slice []*StockRelationship
	var object *StockRelationship

	count := 1
	if singular {
		object = maybeStockRelationship.(*StockRelationship)
	} else {
		slice = *maybeStockRelationship.(*StockRelationshipSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockRelationshipR{}
		args[0] = object.StockRelationshipID
	} else {
		for i, obj := range slice {
			obj.R = &stockRelationshipR{}
			args[i] = obj.StockRelationshipID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_relationship_pub\" where \"stock_relationship_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockRelationshipPub")
	}
	defer results.Close()

	var resultSlice []*StockRelationshipPub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockRelationshipPub")
	}

	if len(stockRelationshipAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.StockRelationshipPub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockRelationshipID == foreign.StockRelationshipID {
				local.R.StockRelationshipPub = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockRelationshipCvterms allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockRelationshipL) LoadStockRelationshipCvterms(e boil.Executor, singular bool, maybeStockRelationship interface{}) error {
	var slice []*StockRelationship
	var object *StockRelationship

	count := 1
	if singular {
		object = maybeStockRelationship.(*StockRelationship)
	} else {
		slice = *maybeStockRelationship.(*StockRelationshipSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockRelationshipR{}
		args[0] = object.StockRelationshipID
	} else {
		for i, obj := range slice {
			obj.R = &stockRelationshipR{}
			args[i] = obj.StockRelationshipID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_relationship_cvterm\" where \"stock_relationship_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load stock_relationship_cvterm")
	}
	defer results.Close()

	var resultSlice []*StockRelationshipCvterm
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice stock_relationship_cvterm")
	}

	if len(stockRelationshipCvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.StockRelationshipCvterms = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockRelationshipID == foreign.StockRelationshipID {
				local.R.StockRelationshipCvterms = append(local.R.StockRelationshipCvterms, foreign)
				break
			}
		}
	}

	return nil
}

// SetType of the stock_relationship to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeStockRelationship.
func (o *StockRelationship) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_relationship\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockRelationshipPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.StockRelationshipID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &stockRelationshipR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeStockRelationship: o,
		}
	} else {
		related.R.TypeStockRelationship = o
	}

	return nil
}

// SetObject of the stock_relationship to the related item.
// Sets o.R.Object to related.
// Adds o to related.R.ObjectStockRelationship.
func (o *StockRelationship) SetObject(exec boil.Executor, insert bool, related *Stock) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_relationship\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"object_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockRelationshipPrimaryKeyColumns),
	)
	values := []interface{}{related.StockID, o.StockRelationshipID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ObjectID = related.StockID

	if o.R == nil {
		o.R = &stockRelationshipR{
			Object: related,
		}
	} else {
		o.R.Object = related
	}

	if related.R == nil {
		related.R = &stockR{
			ObjectStockRelationship: o,
		}
	} else {
		related.R.ObjectStockRelationship = o
	}

	return nil
}

// SetSubject of the stock_relationship to the related item.
// Sets o.R.Subject to related.
// Adds o to related.R.SubjectStockRelationship.
func (o *StockRelationship) SetSubject(exec boil.Executor, insert bool, related *Stock) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_relationship\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"subject_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockRelationshipPrimaryKeyColumns),
	)
	values := []interface{}{related.StockID, o.StockRelationshipID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.SubjectID = related.StockID

	if o.R == nil {
		o.R = &stockRelationshipR{
			Subject: related,
		}
	} else {
		o.R.Subject = related
	}

	if related.R == nil {
		related.R = &stockR{
			SubjectStockRelationship: o,
		}
	} else {
		related.R.SubjectStockRelationship = o
	}

	return nil
}

// SetStockRelationshipPub of the stock_relationship to the related item.
// Sets o.R.StockRelationshipPub to related.
// Adds o to related.R.StockRelationship.
func (o *StockRelationship) SetStockRelationshipPub(exec boil.Executor, insert bool, related *StockRelationshipPub) error {
	var err error

	if insert {
		related.StockRelationshipID = o.StockRelationshipID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock_relationship_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"stock_relationship_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockRelationshipPubPrimaryKeyColumns),
		)
		values := []interface{}{o.StockRelationshipID, related.StockRelationshipPubID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.StockRelationshipID = o.StockRelationshipID

	}

	if o.R == nil {
		o.R = &stockRelationshipR{
			StockRelationshipPub: related,
		}
	} else {
		o.R.StockRelationshipPub = related
	}

	if related.R == nil {
		related.R = &stockRelationshipPubR{
			StockRelationship: o,
		}
	} else {
		related.R.StockRelationship = o
	}
	return nil
}

// AddStockRelationshipCvterms adds the given related objects to the existing relationships
// of the stock_relationship, optionally inserting them as new records.
// Appends related to o.R.StockRelationshipCvterms.
// Sets related.R.StockRelationship appropriately.
func (o *StockRelationship) AddStockRelationshipCvterms(exec boil.Executor, insert bool, related ...*StockRelationshipCvterm) error {
	var err error
	for _, rel := range related {
		rel.StockRelationshipID = o.StockRelationshipID
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "stock_relationship_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &stockRelationshipR{
			StockRelationshipCvterms: related,
		}
	} else {
		o.R.StockRelationshipCvterms = append(o.R.StockRelationshipCvterms, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &stockRelationshipCvtermR{
				StockRelationship: o,
			}
		} else {
			rel.R.StockRelationship = o
		}
	}
	return nil
}

// StockRelationshipsG retrieves all records.
func StockRelationshipsG(mods ...qm.QueryMod) stockRelationshipQuery {
	return StockRelationships(boil.GetDB(), mods...)
}

// StockRelationships retrieves all the records using an executor.
func StockRelationships(exec boil.Executor, mods ...qm.QueryMod) stockRelationshipQuery {
	mods = append(mods, qm.From("\"stock_relationship\""))
	return stockRelationshipQuery{NewQuery(exec, mods...)}
}

// FindStockRelationshipG retrieves a single record by ID.
func FindStockRelationshipG(stockRelationshipID int, selectCols ...string) (*StockRelationship, error) {
	return FindStockRelationship(boil.GetDB(), stockRelationshipID, selectCols...)
}

// FindStockRelationshipGP retrieves a single record by ID, and panics on error.
func FindStockRelationshipGP(stockRelationshipID int, selectCols ...string) *StockRelationship {
	retobj, err := FindStockRelationship(boil.GetDB(), stockRelationshipID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindStockRelationship retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStockRelationship(exec boil.Executor, stockRelationshipID int, selectCols ...string) (*StockRelationship, error) {
	stockRelationshipObj := &StockRelationship{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stock_relationship\" where \"stock_relationship_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, stockRelationshipID)

	err := q.Bind(stockRelationshipObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stock_relationship")
	}

	return stockRelationshipObj, nil
}

// FindStockRelationshipP retrieves a single record by ID with an executor, and panics on error.
func FindStockRelationshipP(exec boil.Executor, stockRelationshipID int, selectCols ...string) *StockRelationship {
	retobj, err := FindStockRelationship(exec, stockRelationshipID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *StockRelationship) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *StockRelationship) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *StockRelationship) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *StockRelationship) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stock_relationship provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockRelationshipColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	stockRelationshipInsertCacheMut.RLock()
	cache, cached := stockRelationshipInsertCache[key]
	stockRelationshipInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			stockRelationshipColumns,
			stockRelationshipColumnsWithDefault,
			stockRelationshipColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(stockRelationshipType, stockRelationshipMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockRelationshipType, stockRelationshipMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"stock_relationship\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into stock_relationship")
	}

	if !cached {
		stockRelationshipInsertCacheMut.Lock()
		stockRelationshipInsertCache[key] = cache
		stockRelationshipInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single StockRelationship record. See Update for
// whitelist behavior description.
func (o *StockRelationship) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single StockRelationship record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *StockRelationship) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the StockRelationship, and panics on error.
// See Update for whitelist behavior description.
func (o *StockRelationship) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the StockRelationship.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *StockRelationship) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	stockRelationshipUpdateCacheMut.RLock()
	cache, cached := stockRelationshipUpdateCache[key]
	stockRelationshipUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(stockRelationshipColumns, stockRelationshipPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update stock_relationship, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stock_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stockRelationshipPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockRelationshipType, stockRelationshipMapping, append(wl, stockRelationshipPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update stock_relationship row")
	}

	if !cached {
		stockRelationshipUpdateCacheMut.Lock()
		stockRelationshipUpdateCache[key] = cache
		stockRelationshipUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q stockRelationshipQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q stockRelationshipQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for stock_relationship")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StockRelationshipSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o StockRelationshipSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o StockRelationshipSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockRelationshipSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockRelationshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"stock_relationship\" SET %s WHERE (\"stock_relationship_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockRelationshipPrimaryKeyColumns), len(colNames)+1, len(stockRelationshipPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in stockRelationship slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *StockRelationship) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *StockRelationship) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *StockRelationship) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *StockRelationship) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stock_relationship provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockRelationshipColumnsWithDefault, o)

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

	stockRelationshipUpsertCacheMut.RLock()
	cache, cached := stockRelationshipUpsertCache[key]
	stockRelationshipUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			stockRelationshipColumns,
			stockRelationshipColumnsWithDefault,
			stockRelationshipColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			stockRelationshipColumns,
			stockRelationshipPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert stock_relationship, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stockRelationshipPrimaryKeyColumns))
			copy(conflict, stockRelationshipPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"stock_relationship\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(stockRelationshipType, stockRelationshipMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockRelationshipType, stockRelationshipMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for stock_relationship")
	}

	if !cached {
		stockRelationshipUpsertCacheMut.Lock()
		stockRelationshipUpsertCache[key] = cache
		stockRelationshipUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single StockRelationship record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockRelationship) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single StockRelationship record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *StockRelationship) DeleteG() error {
	if o == nil {
		return errors.New("models: no StockRelationship provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single StockRelationship record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockRelationship) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single StockRelationship record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *StockRelationship) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no StockRelationship provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockRelationshipPrimaryKeyMapping)
	sql := "DELETE FROM \"stock_relationship\" WHERE \"stock_relationship_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from stock_relationship")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q stockRelationshipQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q stockRelationshipQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no stockRelationshipQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stock_relationship")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o StockRelationshipSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o StockRelationshipSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no StockRelationship slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o StockRelationshipSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockRelationshipSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no StockRelationship slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(stockRelationshipBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockRelationshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"stock_relationship\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockRelationshipPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockRelationshipPrimaryKeyColumns), 1, len(stockRelationshipPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stockRelationship slice")
	}

	if len(stockRelationshipAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *StockRelationship) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *StockRelationship) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *StockRelationship) ReloadG() error {
	if o == nil {
		return errors.New("models: no StockRelationship provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *StockRelationship) Reload(exec boil.Executor) error {
	ret, err := FindStockRelationship(exec, o.StockRelationshipID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockRelationshipSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockRelationshipSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockRelationshipSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty StockRelationshipSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockRelationshipSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	stockRelationships := StockRelationshipSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockRelationshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"stock_relationship\".* FROM \"stock_relationship\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockRelationshipPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(stockRelationshipPrimaryKeyColumns), 1, len(stockRelationshipPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&stockRelationships)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StockRelationshipSlice")
	}

	*o = stockRelationships

	return nil
}

// StockRelationshipExists checks if the StockRelationship row exists.
func StockRelationshipExists(exec boil.Executor, stockRelationshipID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"stock_relationship\" where \"stock_relationship_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, stockRelationshipID)
	}

	row := exec.QueryRow(sql, stockRelationshipID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stock_relationship exists")
	}

	return exists, nil
}

// StockRelationshipExistsG checks if the StockRelationship row exists.
func StockRelationshipExistsG(stockRelationshipID int) (bool, error) {
	return StockRelationshipExists(boil.GetDB(), stockRelationshipID)
}

// StockRelationshipExistsGP checks if the StockRelationship row exists. Panics on error.
func StockRelationshipExistsGP(stockRelationshipID int) bool {
	e, err := StockRelationshipExists(boil.GetDB(), stockRelationshipID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// StockRelationshipExistsP checks if the StockRelationship row exists. Panics on error.
func StockRelationshipExistsP(exec boil.Executor, stockRelationshipID int) bool {
	e, err := StockRelationshipExists(exec, stockRelationshipID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
