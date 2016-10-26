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

// StockItemOrder is an object representing the database table.
type StockItemOrder struct {
	StockItemOrderID int `boil:"stock_item_order_id" json:"stock_item_order_id" toml:"stock_item_order_id" yaml:"stock_item_order_id"`
	ItemID           int `boil:"item_id" json:"item_id" toml:"item_id" yaml:"item_id"`
	OrderID          int `boil:"order_id" json:"order_id" toml:"order_id" yaml:"order_id"`
	Quantity         int `boil:"quantity" json:"quantity" toml:"quantity" yaml:"quantity"`

	R *stockItemOrderR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockItemOrderL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// stockItemOrderR is where relationships are stored.
type stockItemOrderR struct {
	Item  *Stock
	Order *StockOrder
}

// stockItemOrderL is where Load methods for each relationship are stored.
type stockItemOrderL struct{}

var (
	stockItemOrderColumns               = []string{"stock_item_order_id", "item_id", "order_id", "quantity"}
	stockItemOrderColumnsWithoutDefault = []string{"item_id", "order_id"}
	stockItemOrderColumnsWithDefault    = []string{"stock_item_order_id", "quantity"}
	stockItemOrderPrimaryKeyColumns     = []string{"stock_item_order_id"}
)

type (
	// StockItemOrderSlice is an alias for a slice of pointers to StockItemOrder.
	// This should generally be used opposed to []StockItemOrder.
	StockItemOrderSlice []*StockItemOrder
	// StockItemOrderHook is the signature for custom StockItemOrder hook methods
	StockItemOrderHook func(boil.Executor, *StockItemOrder) error

	stockItemOrderQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockItemOrderType                 = reflect.TypeOf(&StockItemOrder{})
	stockItemOrderMapping              = queries.MakeStructMapping(stockItemOrderType)
	stockItemOrderPrimaryKeyMapping, _ = queries.BindMapping(stockItemOrderType, stockItemOrderMapping, stockItemOrderPrimaryKeyColumns)
	stockItemOrderInsertCacheMut       sync.RWMutex
	stockItemOrderInsertCache          = make(map[string]insertCache)
	stockItemOrderUpdateCacheMut       sync.RWMutex
	stockItemOrderUpdateCache          = make(map[string]updateCache)
	stockItemOrderUpsertCacheMut       sync.RWMutex
	stockItemOrderUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var stockItemOrderBeforeInsertHooks []StockItemOrderHook
var stockItemOrderBeforeUpdateHooks []StockItemOrderHook
var stockItemOrderBeforeDeleteHooks []StockItemOrderHook
var stockItemOrderBeforeUpsertHooks []StockItemOrderHook

var stockItemOrderAfterInsertHooks []StockItemOrderHook
var stockItemOrderAfterSelectHooks []StockItemOrderHook
var stockItemOrderAfterUpdateHooks []StockItemOrderHook
var stockItemOrderAfterDeleteHooks []StockItemOrderHook
var stockItemOrderAfterUpsertHooks []StockItemOrderHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *StockItemOrder) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockItemOrderBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *StockItemOrder) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockItemOrderBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *StockItemOrder) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockItemOrderBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *StockItemOrder) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockItemOrderBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *StockItemOrder) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockItemOrderAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *StockItemOrder) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range stockItemOrderAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *StockItemOrder) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockItemOrderAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *StockItemOrder) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockItemOrderAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *StockItemOrder) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockItemOrderAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStockItemOrderHook registers your hook function for all future operations.
func AddStockItemOrderHook(hookPoint boil.HookPoint, stockItemOrderHook StockItemOrderHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stockItemOrderBeforeInsertHooks = append(stockItemOrderBeforeInsertHooks, stockItemOrderHook)
	case boil.BeforeUpdateHook:
		stockItemOrderBeforeUpdateHooks = append(stockItemOrderBeforeUpdateHooks, stockItemOrderHook)
	case boil.BeforeDeleteHook:
		stockItemOrderBeforeDeleteHooks = append(stockItemOrderBeforeDeleteHooks, stockItemOrderHook)
	case boil.BeforeUpsertHook:
		stockItemOrderBeforeUpsertHooks = append(stockItemOrderBeforeUpsertHooks, stockItemOrderHook)
	case boil.AfterInsertHook:
		stockItemOrderAfterInsertHooks = append(stockItemOrderAfterInsertHooks, stockItemOrderHook)
	case boil.AfterSelectHook:
		stockItemOrderAfterSelectHooks = append(stockItemOrderAfterSelectHooks, stockItemOrderHook)
	case boil.AfterUpdateHook:
		stockItemOrderAfterUpdateHooks = append(stockItemOrderAfterUpdateHooks, stockItemOrderHook)
	case boil.AfterDeleteHook:
		stockItemOrderAfterDeleteHooks = append(stockItemOrderAfterDeleteHooks, stockItemOrderHook)
	case boil.AfterUpsertHook:
		stockItemOrderAfterUpsertHooks = append(stockItemOrderAfterUpsertHooks, stockItemOrderHook)
	}
}

// OneP returns a single stockItemOrder record from the query, and panics on error.
func (q stockItemOrderQuery) OneP() *StockItemOrder {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single stockItemOrder record from the query.
func (q stockItemOrderQuery) One() (*StockItemOrder, error) {
	o := &StockItemOrder{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for stock_item_order")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all StockItemOrder records from the query, and panics on error.
func (q stockItemOrderQuery) AllP() StockItemOrderSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all StockItemOrder records from the query.
func (q stockItemOrderQuery) All() (StockItemOrderSlice, error) {
	var o StockItemOrderSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to StockItemOrder slice")
	}

	if len(stockItemOrderAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all StockItemOrder records in the query, and panics on error.
func (q stockItemOrderQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all StockItemOrder records in the query.
func (q stockItemOrderQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count stock_item_order rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q stockItemOrderQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q stockItemOrderQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if stock_item_order exists")
	}

	return count > 0, nil
}

// ItemG pointed to by the foreign key.
func (o *StockItemOrder) ItemG(mods ...qm.QueryMod) stockQuery {
	return o.Item(boil.GetDB(), mods...)
}

// Item pointed to by the foreign key.
func (o *StockItemOrder) Item(exec boil.Executor, mods ...qm.QueryMod) stockQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_id=$1", o.ItemID),
	}

	queryMods = append(queryMods, mods...)

	query := Stocks(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock\"")

	return query
}

// OrderG pointed to by the foreign key.
func (o *StockItemOrder) OrderG(mods ...qm.QueryMod) stockOrderQuery {
	return o.Order(boil.GetDB(), mods...)
}

// Order pointed to by the foreign key.
func (o *StockItemOrder) Order(exec boil.Executor, mods ...qm.QueryMod) stockOrderQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_order_id=$1", o.OrderID),
	}

	queryMods = append(queryMods, mods...)

	query := StockOrders(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_order\"")

	return query
}

// LoadItem allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockItemOrderL) LoadItem(e boil.Executor, singular bool, maybeStockItemOrder interface{}) error {
	var slice []*StockItemOrder
	var object *StockItemOrder

	count := 1
	if singular {
		object = maybeStockItemOrder.(*StockItemOrder)
	} else {
		slice = *maybeStockItemOrder.(*StockItemOrderSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockItemOrderR{}
		args[0] = object.ItemID
	} else {
		for i, obj := range slice {
			obj.R = &stockItemOrderR{}
			args[i] = obj.ItemID
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

	if len(stockItemOrderAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Item = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ItemID == foreign.StockID {
				local.R.Item = foreign
				break
			}
		}
	}

	return nil
}

// LoadOrder allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockItemOrderL) LoadOrder(e boil.Executor, singular bool, maybeStockItemOrder interface{}) error {
	var slice []*StockItemOrder
	var object *StockItemOrder

	count := 1
	if singular {
		object = maybeStockItemOrder.(*StockItemOrder)
	} else {
		slice = *maybeStockItemOrder.(*StockItemOrderSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockItemOrderR{}
		args[0] = object.OrderID
	} else {
		for i, obj := range slice {
			obj.R = &stockItemOrderR{}
			args[i] = obj.OrderID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_order\" where \"stock_order_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockOrder")
	}
	defer results.Close()

	var resultSlice []*StockOrder
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockOrder")
	}

	if len(stockItemOrderAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Order = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.OrderID == foreign.StockOrderID {
				local.R.Order = foreign
				break
			}
		}
	}

	return nil
}

// SetItem of the stock_item_order to the related item.
// Sets o.R.Item to related.
// Adds o to related.R.ItemStockItemOrder.
func (o *StockItemOrder) SetItem(exec boil.Executor, insert bool, related *Stock) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_item_order\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"item_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockItemOrderPrimaryKeyColumns),
	)
	values := []interface{}{related.StockID, o.StockItemOrderID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ItemID = related.StockID

	if o.R == nil {
		o.R = &stockItemOrderR{
			Item: related,
		}
	} else {
		o.R.Item = related
	}

	if related.R == nil {
		related.R = &stockR{
			ItemStockItemOrder: o,
		}
	} else {
		related.R.ItemStockItemOrder = o
	}

	return nil
}

// SetOrder of the stock_item_order to the related item.
// Sets o.R.Order to related.
// Adds o to related.R.OrderStockItemOrder.
func (o *StockItemOrder) SetOrder(exec boil.Executor, insert bool, related *StockOrder) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_item_order\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"order_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockItemOrderPrimaryKeyColumns),
	)
	values := []interface{}{related.StockOrderID, o.StockItemOrderID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OrderID = related.StockOrderID

	if o.R == nil {
		o.R = &stockItemOrderR{
			Order: related,
		}
	} else {
		o.R.Order = related
	}

	if related.R == nil {
		related.R = &stockOrderR{
			OrderStockItemOrder: o,
		}
	} else {
		related.R.OrderStockItemOrder = o
	}

	return nil
}

// StockItemOrdersG retrieves all records.
func StockItemOrdersG(mods ...qm.QueryMod) stockItemOrderQuery {
	return StockItemOrders(boil.GetDB(), mods...)
}

// StockItemOrders retrieves all the records using an executor.
func StockItemOrders(exec boil.Executor, mods ...qm.QueryMod) stockItemOrderQuery {
	mods = append(mods, qm.From("\"stock_item_order\""))
	return stockItemOrderQuery{NewQuery(exec, mods...)}
}

// FindStockItemOrderG retrieves a single record by ID.
func FindStockItemOrderG(stockItemOrderID int, selectCols ...string) (*StockItemOrder, error) {
	return FindStockItemOrder(boil.GetDB(), stockItemOrderID, selectCols...)
}

// FindStockItemOrderGP retrieves a single record by ID, and panics on error.
func FindStockItemOrderGP(stockItemOrderID int, selectCols ...string) *StockItemOrder {
	retobj, err := FindStockItemOrder(boil.GetDB(), stockItemOrderID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindStockItemOrder retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStockItemOrder(exec boil.Executor, stockItemOrderID int, selectCols ...string) (*StockItemOrder, error) {
	stockItemOrderObj := &StockItemOrder{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stock_item_order\" where \"stock_item_order_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, stockItemOrderID)

	err := q.Bind(stockItemOrderObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from stock_item_order")
	}

	return stockItemOrderObj, nil
}

// FindStockItemOrderP retrieves a single record by ID with an executor, and panics on error.
func FindStockItemOrderP(exec boil.Executor, stockItemOrderID int, selectCols ...string) *StockItemOrder {
	retobj, err := FindStockItemOrder(exec, stockItemOrderID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *StockItemOrder) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *StockItemOrder) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *StockItemOrder) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *StockItemOrder) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no stock_item_order provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockItemOrderColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	stockItemOrderInsertCacheMut.RLock()
	cache, cached := stockItemOrderInsertCache[key]
	stockItemOrderInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			stockItemOrderColumns,
			stockItemOrderColumnsWithDefault,
			stockItemOrderColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(stockItemOrderType, stockItemOrderMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockItemOrderType, stockItemOrderMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"stock_item_order\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into stock_item_order")
	}

	if !cached {
		stockItemOrderInsertCacheMut.Lock()
		stockItemOrderInsertCache[key] = cache
		stockItemOrderInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single StockItemOrder record. See Update for
// whitelist behavior description.
func (o *StockItemOrder) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single StockItemOrder record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *StockItemOrder) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the StockItemOrder, and panics on error.
// See Update for whitelist behavior description.
func (o *StockItemOrder) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the StockItemOrder.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *StockItemOrder) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	stockItemOrderUpdateCacheMut.RLock()
	cache, cached := stockItemOrderUpdateCache[key]
	stockItemOrderUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(stockItemOrderColumns, stockItemOrderPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update stock_item_order, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stock_item_order\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stockItemOrderPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockItemOrderType, stockItemOrderMapping, append(wl, stockItemOrderPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update stock_item_order row")
	}

	if !cached {
		stockItemOrderUpdateCacheMut.Lock()
		stockItemOrderUpdateCache[key] = cache
		stockItemOrderUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q stockItemOrderQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q stockItemOrderQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for stock_item_order")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StockItemOrderSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o StockItemOrderSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o StockItemOrderSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockItemOrderSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockItemOrderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"stock_item_order\" SET %s WHERE (\"stock_item_order_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockItemOrderPrimaryKeyColumns), len(colNames)+1, len(stockItemOrderPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in stockItemOrder slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *StockItemOrder) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *StockItemOrder) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *StockItemOrder) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *StockItemOrder) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no stock_item_order provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockItemOrderColumnsWithDefault, o)

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

	stockItemOrderUpsertCacheMut.RLock()
	cache, cached := stockItemOrderUpsertCache[key]
	stockItemOrderUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			stockItemOrderColumns,
			stockItemOrderColumnsWithDefault,
			stockItemOrderColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			stockItemOrderColumns,
			stockItemOrderPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert stock_item_order, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stockItemOrderPrimaryKeyColumns))
			copy(conflict, stockItemOrderPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"stock_item_order\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(stockItemOrderType, stockItemOrderMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockItemOrderType, stockItemOrderMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for stock_item_order")
	}

	if !cached {
		stockItemOrderUpsertCacheMut.Lock()
		stockItemOrderUpsertCache[key] = cache
		stockItemOrderUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single StockItemOrder record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockItemOrder) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single StockItemOrder record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *StockItemOrder) DeleteG() error {
	if o == nil {
		return errors.New("chado: no StockItemOrder provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single StockItemOrder record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockItemOrder) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single StockItemOrder record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *StockItemOrder) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no StockItemOrder provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockItemOrderPrimaryKeyMapping)
	sql := "DELETE FROM \"stock_item_order\" WHERE \"stock_item_order_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from stock_item_order")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q stockItemOrderQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q stockItemOrderQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no stockItemOrderQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from stock_item_order")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o StockItemOrderSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o StockItemOrderSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no StockItemOrder slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o StockItemOrderSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockItemOrderSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no StockItemOrder slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(stockItemOrderBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockItemOrderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"stock_item_order\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockItemOrderPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockItemOrderPrimaryKeyColumns), 1, len(stockItemOrderPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from stockItemOrder slice")
	}

	if len(stockItemOrderAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *StockItemOrder) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *StockItemOrder) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *StockItemOrder) ReloadG() error {
	if o == nil {
		return errors.New("chado: no StockItemOrder provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *StockItemOrder) Reload(exec boil.Executor) error {
	ret, err := FindStockItemOrder(exec, o.StockItemOrderID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockItemOrderSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockItemOrderSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockItemOrderSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty StockItemOrderSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockItemOrderSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	stockItemOrders := StockItemOrderSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockItemOrderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"stock_item_order\".* FROM \"stock_item_order\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockItemOrderPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(stockItemOrderPrimaryKeyColumns), 1, len(stockItemOrderPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&stockItemOrders)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in StockItemOrderSlice")
	}

	*o = stockItemOrders

	return nil
}

// StockItemOrderExists checks if the StockItemOrder row exists.
func StockItemOrderExists(exec boil.Executor, stockItemOrderID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"stock_item_order\" where \"stock_item_order_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, stockItemOrderID)
	}

	row := exec.QueryRow(sql, stockItemOrderID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if stock_item_order exists")
	}

	return exists, nil
}

// StockItemOrderExistsG checks if the StockItemOrder row exists.
func StockItemOrderExistsG(stockItemOrderID int) (bool, error) {
	return StockItemOrderExists(boil.GetDB(), stockItemOrderID)
}

// StockItemOrderExistsGP checks if the StockItemOrder row exists. Panics on error.
func StockItemOrderExistsGP(stockItemOrderID int) bool {
	e, err := StockItemOrderExists(boil.GetDB(), stockItemOrderID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// StockItemOrderExistsP checks if the StockItemOrder row exists. Panics on error.
func StockItemOrderExistsP(exec boil.Executor, stockItemOrderID int) bool {
	e, err := StockItemOrderExists(exec, stockItemOrderID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
