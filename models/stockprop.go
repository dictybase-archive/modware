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

// Stockprop is an object representing the database table.
type Stockprop struct {
	StockpropID int         `boil:"stockprop_id" json:"stockprop_id" toml:"stockprop_id" yaml:"stockprop_id"`
	StockID     int         `boil:"stock_id" json:"stock_id" toml:"stock_id" yaml:"stock_id"`
	TypeID      int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Value       null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	Rank        int         `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *stockpropR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockpropL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// stockpropR is where relationships are stored.
type stockpropR struct {
	Type         *Cvterm
	Stock        *Stock
	StockpropPub *StockpropPub
}

// stockpropL is where Load methods for each relationship are stored.
type stockpropL struct{}

var (
	stockpropColumns               = []string{"stockprop_id", "stock_id", "type_id", "value", "rank"}
	stockpropColumnsWithoutDefault = []string{"stock_id", "type_id", "value"}
	stockpropColumnsWithDefault    = []string{"stockprop_id", "rank"}
	stockpropPrimaryKeyColumns     = []string{"stockprop_id"}
)

type (
	// StockpropSlice is an alias for a slice of pointers to Stockprop.
	// This should generally be used opposed to []Stockprop.
	StockpropSlice []*Stockprop
	// StockpropHook is the signature for custom Stockprop hook methods
	StockpropHook func(boil.Executor, *Stockprop) error

	stockpropQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockpropType                 = reflect.TypeOf(&Stockprop{})
	stockpropMapping              = queries.MakeStructMapping(stockpropType)
	stockpropPrimaryKeyMapping, _ = queries.BindMapping(stockpropType, stockpropMapping, stockpropPrimaryKeyColumns)
	stockpropInsertCacheMut       sync.RWMutex
	stockpropInsertCache          = make(map[string]insertCache)
	stockpropUpdateCacheMut       sync.RWMutex
	stockpropUpdateCache          = make(map[string]updateCache)
	stockpropUpsertCacheMut       sync.RWMutex
	stockpropUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var stockpropBeforeInsertHooks []StockpropHook
var stockpropBeforeUpdateHooks []StockpropHook
var stockpropBeforeDeleteHooks []StockpropHook
var stockpropBeforeUpsertHooks []StockpropHook

var stockpropAfterInsertHooks []StockpropHook
var stockpropAfterSelectHooks []StockpropHook
var stockpropAfterUpdateHooks []StockpropHook
var stockpropAfterDeleteHooks []StockpropHook
var stockpropAfterUpsertHooks []StockpropHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Stockprop) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockpropBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Stockprop) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockpropBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Stockprop) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockpropBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Stockprop) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockpropBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Stockprop) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockpropAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Stockprop) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range stockpropAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Stockprop) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockpropAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Stockprop) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockpropAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Stockprop) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockpropAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStockpropHook registers your hook function for all future operations.
func AddStockpropHook(hookPoint boil.HookPoint, stockpropHook StockpropHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stockpropBeforeInsertHooks = append(stockpropBeforeInsertHooks, stockpropHook)
	case boil.BeforeUpdateHook:
		stockpropBeforeUpdateHooks = append(stockpropBeforeUpdateHooks, stockpropHook)
	case boil.BeforeDeleteHook:
		stockpropBeforeDeleteHooks = append(stockpropBeforeDeleteHooks, stockpropHook)
	case boil.BeforeUpsertHook:
		stockpropBeforeUpsertHooks = append(stockpropBeforeUpsertHooks, stockpropHook)
	case boil.AfterInsertHook:
		stockpropAfterInsertHooks = append(stockpropAfterInsertHooks, stockpropHook)
	case boil.AfterSelectHook:
		stockpropAfterSelectHooks = append(stockpropAfterSelectHooks, stockpropHook)
	case boil.AfterUpdateHook:
		stockpropAfterUpdateHooks = append(stockpropAfterUpdateHooks, stockpropHook)
	case boil.AfterDeleteHook:
		stockpropAfterDeleteHooks = append(stockpropAfterDeleteHooks, stockpropHook)
	case boil.AfterUpsertHook:
		stockpropAfterUpsertHooks = append(stockpropAfterUpsertHooks, stockpropHook)
	}
}

// OneP returns a single stockprop record from the query, and panics on error.
func (q stockpropQuery) OneP() *Stockprop {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single stockprop record from the query.
func (q stockpropQuery) One() (*Stockprop, error) {
	o := &Stockprop{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stockprop")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Stockprop records from the query, and panics on error.
func (q stockpropQuery) AllP() StockpropSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Stockprop records from the query.
func (q stockpropQuery) All() (StockpropSlice, error) {
	var o StockpropSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Stockprop slice")
	}

	if len(stockpropAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Stockprop records in the query, and panics on error.
func (q stockpropQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Stockprop records in the query.
func (q stockpropQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stockprop rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q stockpropQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q stockpropQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stockprop exists")
	}

	return count > 0, nil
}

// TypeG pointed to by the foreign key.
func (o *Stockprop) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Stockprop) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// StockG pointed to by the foreign key.
func (o *Stockprop) StockG(mods ...qm.QueryMod) stockQuery {
	return o.Stock(boil.GetDB(), mods...)
}

// Stock pointed to by the foreign key.
func (o *Stockprop) Stock(exec boil.Executor, mods ...qm.QueryMod) stockQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_id=$1", o.StockID),
	}

	queryMods = append(queryMods, mods...)

	query := Stocks(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock\"")

	return query
}

// StockpropPubG pointed to by the foreign key.
func (o *Stockprop) StockpropPubG(mods ...qm.QueryMod) stockpropPubQuery {
	return o.StockpropPub(boil.GetDB(), mods...)
}

// StockpropPub pointed to by the foreign key.
func (o *Stockprop) StockpropPub(exec boil.Executor, mods ...qm.QueryMod) stockpropPubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stockprop_id=$1", o.StockpropID),
	}

	queryMods = append(queryMods, mods...)

	query := StockpropPubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stockprop_pub\"")

	return query
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockpropL) LoadType(e boil.Executor, singular bool, maybeStockprop interface{}) error {
	var slice []*Stockprop
	var object *Stockprop

	count := 1
	if singular {
		object = maybeStockprop.(*Stockprop)
	} else {
		slice = *maybeStockprop.(*StockpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockpropR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &stockpropR{}
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

	if len(stockpropAfterSelectHooks) != 0 {
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

// LoadStock allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockpropL) LoadStock(e boil.Executor, singular bool, maybeStockprop interface{}) error {
	var slice []*Stockprop
	var object *Stockprop

	count := 1
	if singular {
		object = maybeStockprop.(*Stockprop)
	} else {
		slice = *maybeStockprop.(*StockpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockpropR{}
		args[0] = object.StockID
	} else {
		for i, obj := range slice {
			obj.R = &stockpropR{}
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

	if len(stockpropAfterSelectHooks) != 0 {
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

// LoadStockpropPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockpropL) LoadStockpropPub(e boil.Executor, singular bool, maybeStockprop interface{}) error {
	var slice []*Stockprop
	var object *Stockprop

	count := 1
	if singular {
		object = maybeStockprop.(*Stockprop)
	} else {
		slice = *maybeStockprop.(*StockpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockpropR{}
		args[0] = object.StockpropID
	} else {
		for i, obj := range slice {
			obj.R = &stockpropR{}
			args[i] = obj.StockpropID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stockprop_pub\" where \"stockprop_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockpropPub")
	}
	defer results.Close()

	var resultSlice []*StockpropPub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockpropPub")
	}

	if len(stockpropAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.StockpropPub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockpropID == foreign.StockpropID {
				local.R.StockpropPub = foreign
				break
			}
		}
	}

	return nil
}

// SetType of the stockprop to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeStockprop.
func (o *Stockprop) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stockprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockpropPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.StockpropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &stockpropR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeStockprop: o,
		}
	} else {
		related.R.TypeStockprop = o
	}

	return nil
}

// SetStock of the stockprop to the related item.
// Sets o.R.Stock to related.
// Adds o to related.R.Stockprop.
func (o *Stockprop) SetStock(exec boil.Executor, insert bool, related *Stock) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stockprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"stock_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockpropPrimaryKeyColumns),
	)
	values := []interface{}{related.StockID, o.StockpropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StockID = related.StockID

	if o.R == nil {
		o.R = &stockpropR{
			Stock: related,
		}
	} else {
		o.R.Stock = related
	}

	if related.R == nil {
		related.R = &stockR{
			Stockprop: o,
		}
	} else {
		related.R.Stockprop = o
	}

	return nil
}

// SetStockpropPub of the stockprop to the related item.
// Sets o.R.StockpropPub to related.
// Adds o to related.R.Stockprop.
func (o *Stockprop) SetStockpropPub(exec boil.Executor, insert bool, related *StockpropPub) error {
	var err error

	if insert {
		related.StockpropID = o.StockpropID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stockprop_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"stockprop_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockpropPubPrimaryKeyColumns),
		)
		values := []interface{}{o.StockpropID, related.StockpropPubID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.StockpropID = o.StockpropID

	}

	if o.R == nil {
		o.R = &stockpropR{
			StockpropPub: related,
		}
	} else {
		o.R.StockpropPub = related
	}

	if related.R == nil {
		related.R = &stockpropPubR{
			Stockprop: o,
		}
	} else {
		related.R.Stockprop = o
	}
	return nil
}

// StockpropsG retrieves all records.
func StockpropsG(mods ...qm.QueryMod) stockpropQuery {
	return Stockprops(boil.GetDB(), mods...)
}

// Stockprops retrieves all the records using an executor.
func Stockprops(exec boil.Executor, mods ...qm.QueryMod) stockpropQuery {
	mods = append(mods, qm.From("\"stockprop\""))
	return stockpropQuery{NewQuery(exec, mods...)}
}

// FindStockpropG retrieves a single record by ID.
func FindStockpropG(stockpropID int, selectCols ...string) (*Stockprop, error) {
	return FindStockprop(boil.GetDB(), stockpropID, selectCols...)
}

// FindStockpropGP retrieves a single record by ID, and panics on error.
func FindStockpropGP(stockpropID int, selectCols ...string) *Stockprop {
	retobj, err := FindStockprop(boil.GetDB(), stockpropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindStockprop retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStockprop(exec boil.Executor, stockpropID int, selectCols ...string) (*Stockprop, error) {
	stockpropObj := &Stockprop{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stockprop\" where \"stockprop_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, stockpropID)

	err := q.Bind(stockpropObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stockprop")
	}

	return stockpropObj, nil
}

// FindStockpropP retrieves a single record by ID with an executor, and panics on error.
func FindStockpropP(exec boil.Executor, stockpropID int, selectCols ...string) *Stockprop {
	retobj, err := FindStockprop(exec, stockpropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Stockprop) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Stockprop) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Stockprop) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Stockprop) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stockprop provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockpropColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	stockpropInsertCacheMut.RLock()
	cache, cached := stockpropInsertCache[key]
	stockpropInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			stockpropColumns,
			stockpropColumnsWithDefault,
			stockpropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(stockpropType, stockpropMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockpropType, stockpropMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"stockprop\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into stockprop")
	}

	if !cached {
		stockpropInsertCacheMut.Lock()
		stockpropInsertCache[key] = cache
		stockpropInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Stockprop record. See Update for
// whitelist behavior description.
func (o *Stockprop) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Stockprop record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Stockprop) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Stockprop, and panics on error.
// See Update for whitelist behavior description.
func (o *Stockprop) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Stockprop.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Stockprop) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	stockpropUpdateCacheMut.RLock()
	cache, cached := stockpropUpdateCache[key]
	stockpropUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(stockpropColumns, stockpropPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update stockprop, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stockprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stockpropPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockpropType, stockpropMapping, append(wl, stockpropPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update stockprop row")
	}

	if !cached {
		stockpropUpdateCacheMut.Lock()
		stockpropUpdateCache[key] = cache
		stockpropUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q stockpropQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q stockpropQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for stockprop")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StockpropSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o StockpropSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o StockpropSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockpropSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"stockprop\" SET %s WHERE (\"stockprop_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockpropPrimaryKeyColumns), len(colNames)+1, len(stockpropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in stockprop slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Stockprop) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Stockprop) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Stockprop) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Stockprop) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stockprop provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockpropColumnsWithDefault, o)

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

	stockpropUpsertCacheMut.RLock()
	cache, cached := stockpropUpsertCache[key]
	stockpropUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			stockpropColumns,
			stockpropColumnsWithDefault,
			stockpropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			stockpropColumns,
			stockpropPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert stockprop, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stockpropPrimaryKeyColumns))
			copy(conflict, stockpropPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"stockprop\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(stockpropType, stockpropMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockpropType, stockpropMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for stockprop")
	}

	if !cached {
		stockpropUpsertCacheMut.Lock()
		stockpropUpsertCache[key] = cache
		stockpropUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Stockprop record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Stockprop) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Stockprop record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Stockprop) DeleteG() error {
	if o == nil {
		return errors.New("models: no Stockprop provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Stockprop record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Stockprop) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Stockprop record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Stockprop) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Stockprop provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockpropPrimaryKeyMapping)
	sql := "DELETE FROM \"stockprop\" WHERE \"stockprop_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from stockprop")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q stockpropQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q stockpropQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no stockpropQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stockprop")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o StockpropSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o StockpropSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Stockprop slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o StockpropSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockpropSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Stockprop slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(stockpropBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"stockprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockpropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockpropPrimaryKeyColumns), 1, len(stockpropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stockprop slice")
	}

	if len(stockpropAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Stockprop) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Stockprop) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Stockprop) ReloadG() error {
	if o == nil {
		return errors.New("models: no Stockprop provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Stockprop) Reload(exec boil.Executor) error {
	ret, err := FindStockprop(exec, o.StockpropID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockpropSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockpropSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockpropSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty StockpropSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockpropSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	stockprops := StockpropSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"stockprop\".* FROM \"stockprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockpropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(stockpropPrimaryKeyColumns), 1, len(stockpropPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&stockprops)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StockpropSlice")
	}

	*o = stockprops

	return nil
}

// StockpropExists checks if the Stockprop row exists.
func StockpropExists(exec boil.Executor, stockpropID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"stockprop\" where \"stockprop_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, stockpropID)
	}

	row := exec.QueryRow(sql, stockpropID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stockprop exists")
	}

	return exists, nil
}

// StockpropExistsG checks if the Stockprop row exists.
func StockpropExistsG(stockpropID int) (bool, error) {
	return StockpropExists(boil.GetDB(), stockpropID)
}

// StockpropExistsGP checks if the Stockprop row exists. Panics on error.
func StockpropExistsGP(stockpropID int) bool {
	e, err := StockpropExists(boil.GetDB(), stockpropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// StockpropExistsP checks if the Stockprop row exists. Panics on error.
func StockpropExistsP(exec boil.Executor, stockpropID int) bool {
	e, err := StockpropExists(exec, stockpropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
