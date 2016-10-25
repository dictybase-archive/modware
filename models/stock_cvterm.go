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

// StockCvterm is an object representing the database table.
type StockCvterm struct {
	StockCvtermID int  `boil:"stock_cvterm_id" json:"stock_cvterm_id" toml:"stock_cvterm_id" yaml:"stock_cvterm_id"`
	StockID       int  `boil:"stock_id" json:"stock_id" toml:"stock_id" yaml:"stock_id"`
	CvtermID      int  `boil:"cvterm_id" json:"cvterm_id" toml:"cvterm_id" yaml:"cvterm_id"`
	PubID         int  `boil:"pub_id" json:"pub_id" toml:"pub_id" yaml:"pub_id"`
	IsNot         bool `boil:"is_not" json:"is_not" toml:"is_not" yaml:"is_not"`
	Rank          int  `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *stockCvtermR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockCvtermL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// stockCvtermR is where relationships are stored.
type stockCvtermR struct {
	Cvterm          *Cvterm
	Pub             *Pub
	Stock           *Stock
	StockCvtermprop *StockCvtermprop
}

// stockCvtermL is where Load methods for each relationship are stored.
type stockCvtermL struct{}

var (
	stockCvtermColumns               = []string{"stock_cvterm_id", "stock_id", "cvterm_id", "pub_id", "is_not", "rank"}
	stockCvtermColumnsWithoutDefault = []string{"stock_id", "cvterm_id", "pub_id"}
	stockCvtermColumnsWithDefault    = []string{"stock_cvterm_id", "is_not", "rank"}
	stockCvtermPrimaryKeyColumns     = []string{"stock_cvterm_id"}
)

type (
	// StockCvtermSlice is an alias for a slice of pointers to StockCvterm.
	// This should generally be used opposed to []StockCvterm.
	StockCvtermSlice []*StockCvterm
	// StockCvtermHook is the signature for custom StockCvterm hook methods
	StockCvtermHook func(boil.Executor, *StockCvterm) error

	stockCvtermQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockCvtermType                 = reflect.TypeOf(&StockCvterm{})
	stockCvtermMapping              = queries.MakeStructMapping(stockCvtermType)
	stockCvtermPrimaryKeyMapping, _ = queries.BindMapping(stockCvtermType, stockCvtermMapping, stockCvtermPrimaryKeyColumns)
	stockCvtermInsertCacheMut       sync.RWMutex
	stockCvtermInsertCache          = make(map[string]insertCache)
	stockCvtermUpdateCacheMut       sync.RWMutex
	stockCvtermUpdateCache          = make(map[string]updateCache)
	stockCvtermUpsertCacheMut       sync.RWMutex
	stockCvtermUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var stockCvtermBeforeInsertHooks []StockCvtermHook
var stockCvtermBeforeUpdateHooks []StockCvtermHook
var stockCvtermBeforeDeleteHooks []StockCvtermHook
var stockCvtermBeforeUpsertHooks []StockCvtermHook

var stockCvtermAfterInsertHooks []StockCvtermHook
var stockCvtermAfterSelectHooks []StockCvtermHook
var stockCvtermAfterUpdateHooks []StockCvtermHook
var stockCvtermAfterDeleteHooks []StockCvtermHook
var stockCvtermAfterUpsertHooks []StockCvtermHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *StockCvterm) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockCvtermBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *StockCvterm) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockCvtermBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *StockCvterm) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockCvtermBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *StockCvterm) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockCvtermBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *StockCvterm) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockCvtermAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *StockCvterm) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range stockCvtermAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *StockCvterm) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockCvtermAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *StockCvterm) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockCvtermAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *StockCvterm) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockCvtermAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStockCvtermHook registers your hook function for all future operations.
func AddStockCvtermHook(hookPoint boil.HookPoint, stockCvtermHook StockCvtermHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stockCvtermBeforeInsertHooks = append(stockCvtermBeforeInsertHooks, stockCvtermHook)
	case boil.BeforeUpdateHook:
		stockCvtermBeforeUpdateHooks = append(stockCvtermBeforeUpdateHooks, stockCvtermHook)
	case boil.BeforeDeleteHook:
		stockCvtermBeforeDeleteHooks = append(stockCvtermBeforeDeleteHooks, stockCvtermHook)
	case boil.BeforeUpsertHook:
		stockCvtermBeforeUpsertHooks = append(stockCvtermBeforeUpsertHooks, stockCvtermHook)
	case boil.AfterInsertHook:
		stockCvtermAfterInsertHooks = append(stockCvtermAfterInsertHooks, stockCvtermHook)
	case boil.AfterSelectHook:
		stockCvtermAfterSelectHooks = append(stockCvtermAfterSelectHooks, stockCvtermHook)
	case boil.AfterUpdateHook:
		stockCvtermAfterUpdateHooks = append(stockCvtermAfterUpdateHooks, stockCvtermHook)
	case boil.AfterDeleteHook:
		stockCvtermAfterDeleteHooks = append(stockCvtermAfterDeleteHooks, stockCvtermHook)
	case boil.AfterUpsertHook:
		stockCvtermAfterUpsertHooks = append(stockCvtermAfterUpsertHooks, stockCvtermHook)
	}
}

// OneP returns a single stockCvterm record from the query, and panics on error.
func (q stockCvtermQuery) OneP() *StockCvterm {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single stockCvterm record from the query.
func (q stockCvtermQuery) One() (*StockCvterm, error) {
	o := &StockCvterm{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stock_cvterm")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all StockCvterm records from the query, and panics on error.
func (q stockCvtermQuery) AllP() StockCvtermSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all StockCvterm records from the query.
func (q stockCvtermQuery) All() (StockCvtermSlice, error) {
	var o StockCvtermSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to StockCvterm slice")
	}

	if len(stockCvtermAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all StockCvterm records in the query, and panics on error.
func (q stockCvtermQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all StockCvterm records in the query.
func (q stockCvtermQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stock_cvterm rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q stockCvtermQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q stockCvtermQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stock_cvterm exists")
	}

	return count > 0, nil
}

// CvtermG pointed to by the foreign key.
func (o *StockCvterm) CvtermG(mods ...qm.QueryMod) cvtermQuery {
	return o.Cvterm(boil.GetDB(), mods...)
}

// Cvterm pointed to by the foreign key.
func (o *StockCvterm) Cvterm(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// PubG pointed to by the foreign key.
func (o *StockCvterm) PubG(mods ...qm.QueryMod) pubQuery {
	return o.Pub(boil.GetDB(), mods...)
}

// Pub pointed to by the foreign key.
func (o *StockCvterm) Pub(exec boil.Executor, mods ...qm.QueryMod) pubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub\"")

	return query
}

// StockG pointed to by the foreign key.
func (o *StockCvterm) StockG(mods ...qm.QueryMod) stockQuery {
	return o.Stock(boil.GetDB(), mods...)
}

// Stock pointed to by the foreign key.
func (o *StockCvterm) Stock(exec boil.Executor, mods ...qm.QueryMod) stockQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_id=$1", o.StockID),
	}

	queryMods = append(queryMods, mods...)

	query := Stocks(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock\"")

	return query
}

// StockCvtermpropG pointed to by the foreign key.
func (o *StockCvterm) StockCvtermpropG(mods ...qm.QueryMod) stockCvtermpropQuery {
	return o.StockCvtermprop(boil.GetDB(), mods...)
}

// StockCvtermprop pointed to by the foreign key.
func (o *StockCvterm) StockCvtermprop(exec boil.Executor, mods ...qm.QueryMod) stockCvtermpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_cvterm_id=$1", o.StockCvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := StockCvtermprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_cvtermprop\"")

	return query
}

// LoadCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockCvtermL) LoadCvterm(e boil.Executor, singular bool, maybeStockCvterm interface{}) error {
	var slice []*StockCvterm
	var object *StockCvterm

	count := 1
	if singular {
		object = maybeStockCvterm.(*StockCvterm)
	} else {
		slice = *maybeStockCvterm.(*StockCvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockCvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &stockCvtermR{}
			args[i] = obj.CvtermID
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

	if len(stockCvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Cvterm = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.CvtermID {
				local.R.Cvterm = foreign
				break
			}
		}
	}

	return nil
}

// LoadPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockCvtermL) LoadPub(e boil.Executor, singular bool, maybeStockCvterm interface{}) error {
	var slice []*StockCvterm
	var object *StockCvterm

	count := 1
	if singular {
		object = maybeStockCvterm.(*StockCvterm)
	} else {
		slice = *maybeStockCvterm.(*StockCvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockCvtermR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &stockCvtermR{}
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

	if len(stockCvtermAfterSelectHooks) != 0 {
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

// LoadStock allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockCvtermL) LoadStock(e boil.Executor, singular bool, maybeStockCvterm interface{}) error {
	var slice []*StockCvterm
	var object *StockCvterm

	count := 1
	if singular {
		object = maybeStockCvterm.(*StockCvterm)
	} else {
		slice = *maybeStockCvterm.(*StockCvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockCvtermR{}
		args[0] = object.StockID
	} else {
		for i, obj := range slice {
			obj.R = &stockCvtermR{}
			args[i] = obj.StockID
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

	if len(stockCvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Stock = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockID == foreign.StockID {
				local.R.Stock = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockCvtermprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockCvtermL) LoadStockCvtermprop(e boil.Executor, singular bool, maybeStockCvterm interface{}) error {
	var slice []*StockCvterm
	var object *StockCvterm

	count := 1
	if singular {
		object = maybeStockCvterm.(*StockCvterm)
	} else {
		slice = *maybeStockCvterm.(*StockCvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockCvtermR{}
		args[0] = object.StockCvtermID
	} else {
		for i, obj := range slice {
			obj.R = &stockCvtermR{}
			args[i] = obj.StockCvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_cvtermprop\" where \"stock_cvterm_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockCvtermprop")
	}
	defer results.Close()

	var resultSlice []*StockCvtermprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockCvtermprop")
	}

	if len(stockCvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.StockCvtermprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockCvtermID == foreign.StockCvtermID {
				local.R.StockCvtermprop = foreign
				break
			}
		}
	}

	return nil
}

// SetCvterm of the stock_cvterm to the related item.
// Sets o.R.Cvterm to related.
// Adds o to related.R.StockCvterm.
func (o *StockCvterm) SetCvterm(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_cvterm\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"cvterm_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockCvtermPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.StockCvtermID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CvtermID = related.CvtermID

	if o.R == nil {
		o.R = &stockCvtermR{
			Cvterm: related,
		}
	} else {
		o.R.Cvterm = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			StockCvterm: o,
		}
	} else {
		related.R.StockCvterm = o
	}

	return nil
}

// SetPub of the stock_cvterm to the related item.
// Sets o.R.Pub to related.
// Adds o to related.R.StockCvterm.
func (o *StockCvterm) SetPub(exec boil.Executor, insert bool, related *Pub) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_cvterm\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockCvtermPrimaryKeyColumns),
	)
	values := []interface{}{related.PubID, o.StockCvtermID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PubID = related.PubID

	if o.R == nil {
		o.R = &stockCvtermR{
			Pub: related,
		}
	} else {
		o.R.Pub = related
	}

	if related.R == nil {
		related.R = &pubR{
			StockCvterm: o,
		}
	} else {
		related.R.StockCvterm = o
	}

	return nil
}

// SetStock of the stock_cvterm to the related item.
// Sets o.R.Stock to related.
// Adds o to related.R.StockCvterm.
func (o *StockCvterm) SetStock(exec boil.Executor, insert bool, related *Stock) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_cvterm\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"stock_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockCvtermPrimaryKeyColumns),
	)
	values := []interface{}{related.StockID, o.StockCvtermID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StockID = related.StockID

	if o.R == nil {
		o.R = &stockCvtermR{
			Stock: related,
		}
	} else {
		o.R.Stock = related
	}

	if related.R == nil {
		related.R = &stockR{
			StockCvterm: o,
		}
	} else {
		related.R.StockCvterm = o
	}

	return nil
}

// SetStockCvtermprop of the stock_cvterm to the related item.
// Sets o.R.StockCvtermprop to related.
// Adds o to related.R.StockCvterm.
func (o *StockCvterm) SetStockCvtermprop(exec boil.Executor, insert bool, related *StockCvtermprop) error {
	var err error

	if insert {
		related.StockCvtermID = o.StockCvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock_cvtermprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"stock_cvterm_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockCvtermpropPrimaryKeyColumns),
		)
		values := []interface{}{o.StockCvtermID, related.StockCvtermpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.StockCvtermID = o.StockCvtermID

	}

	if o.R == nil {
		o.R = &stockCvtermR{
			StockCvtermprop: related,
		}
	} else {
		o.R.StockCvtermprop = related
	}

	if related.R == nil {
		related.R = &stockCvtermpropR{
			StockCvterm: o,
		}
	} else {
		related.R.StockCvterm = o
	}
	return nil
}

// StockCvtermsG retrieves all records.
func StockCvtermsG(mods ...qm.QueryMod) stockCvtermQuery {
	return StockCvterms(boil.GetDB(), mods...)
}

// StockCvterms retrieves all the records using an executor.
func StockCvterms(exec boil.Executor, mods ...qm.QueryMod) stockCvtermQuery {
	mods = append(mods, qm.From("\"stock_cvterm\""))
	return stockCvtermQuery{NewQuery(exec, mods...)}
}

// FindStockCvtermG retrieves a single record by ID.
func FindStockCvtermG(stockCvtermID int, selectCols ...string) (*StockCvterm, error) {
	return FindStockCvterm(boil.GetDB(), stockCvtermID, selectCols...)
}

// FindStockCvtermGP retrieves a single record by ID, and panics on error.
func FindStockCvtermGP(stockCvtermID int, selectCols ...string) *StockCvterm {
	retobj, err := FindStockCvterm(boil.GetDB(), stockCvtermID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindStockCvterm retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStockCvterm(exec boil.Executor, stockCvtermID int, selectCols ...string) (*StockCvterm, error) {
	stockCvtermObj := &StockCvterm{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stock_cvterm\" where \"stock_cvterm_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, stockCvtermID)

	err := q.Bind(stockCvtermObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stock_cvterm")
	}

	return stockCvtermObj, nil
}

// FindStockCvtermP retrieves a single record by ID with an executor, and panics on error.
func FindStockCvtermP(exec boil.Executor, stockCvtermID int, selectCols ...string) *StockCvterm {
	retobj, err := FindStockCvterm(exec, stockCvtermID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *StockCvterm) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *StockCvterm) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *StockCvterm) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *StockCvterm) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stock_cvterm provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockCvtermColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	stockCvtermInsertCacheMut.RLock()
	cache, cached := stockCvtermInsertCache[key]
	stockCvtermInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			stockCvtermColumns,
			stockCvtermColumnsWithDefault,
			stockCvtermColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(stockCvtermType, stockCvtermMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockCvtermType, stockCvtermMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"stock_cvterm\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into stock_cvterm")
	}

	if !cached {
		stockCvtermInsertCacheMut.Lock()
		stockCvtermInsertCache[key] = cache
		stockCvtermInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single StockCvterm record. See Update for
// whitelist behavior description.
func (o *StockCvterm) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single StockCvterm record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *StockCvterm) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the StockCvterm, and panics on error.
// See Update for whitelist behavior description.
func (o *StockCvterm) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the StockCvterm.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *StockCvterm) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	stockCvtermUpdateCacheMut.RLock()
	cache, cached := stockCvtermUpdateCache[key]
	stockCvtermUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(stockCvtermColumns, stockCvtermPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update stock_cvterm, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stock_cvterm\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stockCvtermPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockCvtermType, stockCvtermMapping, append(wl, stockCvtermPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update stock_cvterm row")
	}

	if !cached {
		stockCvtermUpdateCacheMut.Lock()
		stockCvtermUpdateCache[key] = cache
		stockCvtermUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q stockCvtermQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q stockCvtermQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for stock_cvterm")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StockCvtermSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o StockCvtermSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o StockCvtermSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockCvtermSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockCvtermPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"stock_cvterm\" SET %s WHERE (\"stock_cvterm_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockCvtermPrimaryKeyColumns), len(colNames)+1, len(stockCvtermPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in stockCvterm slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *StockCvterm) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *StockCvterm) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *StockCvterm) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *StockCvterm) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stock_cvterm provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockCvtermColumnsWithDefault, o)

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

	stockCvtermUpsertCacheMut.RLock()
	cache, cached := stockCvtermUpsertCache[key]
	stockCvtermUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			stockCvtermColumns,
			stockCvtermColumnsWithDefault,
			stockCvtermColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			stockCvtermColumns,
			stockCvtermPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert stock_cvterm, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stockCvtermPrimaryKeyColumns))
			copy(conflict, stockCvtermPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"stock_cvterm\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(stockCvtermType, stockCvtermMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockCvtermType, stockCvtermMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for stock_cvterm")
	}

	if !cached {
		stockCvtermUpsertCacheMut.Lock()
		stockCvtermUpsertCache[key] = cache
		stockCvtermUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single StockCvterm record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockCvterm) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single StockCvterm record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *StockCvterm) DeleteG() error {
	if o == nil {
		return errors.New("models: no StockCvterm provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single StockCvterm record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockCvterm) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single StockCvterm record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *StockCvterm) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no StockCvterm provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockCvtermPrimaryKeyMapping)
	sql := "DELETE FROM \"stock_cvterm\" WHERE \"stock_cvterm_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from stock_cvterm")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q stockCvtermQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q stockCvtermQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no stockCvtermQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stock_cvterm")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o StockCvtermSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o StockCvtermSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no StockCvterm slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o StockCvtermSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockCvtermSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no StockCvterm slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(stockCvtermBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockCvtermPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"stock_cvterm\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockCvtermPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockCvtermPrimaryKeyColumns), 1, len(stockCvtermPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stockCvterm slice")
	}

	if len(stockCvtermAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *StockCvterm) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *StockCvterm) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *StockCvterm) ReloadG() error {
	if o == nil {
		return errors.New("models: no StockCvterm provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *StockCvterm) Reload(exec boil.Executor) error {
	ret, err := FindStockCvterm(exec, o.StockCvtermID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockCvtermSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockCvtermSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockCvtermSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty StockCvtermSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockCvtermSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	stockCvterms := StockCvtermSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockCvtermPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"stock_cvterm\".* FROM \"stock_cvterm\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockCvtermPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(stockCvtermPrimaryKeyColumns), 1, len(stockCvtermPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&stockCvterms)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StockCvtermSlice")
	}

	*o = stockCvterms

	return nil
}

// StockCvtermExists checks if the StockCvterm row exists.
func StockCvtermExists(exec boil.Executor, stockCvtermID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"stock_cvterm\" where \"stock_cvterm_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, stockCvtermID)
	}

	row := exec.QueryRow(sql, stockCvtermID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stock_cvterm exists")
	}

	return exists, nil
}

// StockCvtermExistsG checks if the StockCvterm row exists.
func StockCvtermExistsG(stockCvtermID int) (bool, error) {
	return StockCvtermExists(boil.GetDB(), stockCvtermID)
}

// StockCvtermExistsGP checks if the StockCvterm row exists. Panics on error.
func StockCvtermExistsGP(stockCvtermID int) bool {
	e, err := StockCvtermExists(boil.GetDB(), stockCvtermID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// StockCvtermExistsP checks if the StockCvterm row exists. Panics on error.
func StockCvtermExistsP(exec boil.Executor, stockCvtermID int) bool {
	e, err := StockCvtermExists(exec, stockCvtermID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
