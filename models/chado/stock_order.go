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

// StockOrder is an object representing the database table.
type StockOrder struct {
	StockOrderID int       `boil:"stock_order_id" json:"stock_order_id" toml:"stock_order_id" yaml:"stock_order_id"`
	UserID       int       `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	CreatedAt    null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt    null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *stockOrderR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockOrderL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// stockOrderR is where relationships are stored.
type stockOrderR struct {
	User                *AuthUser
	OrderStockItemOrder *StockItemOrder
}

// stockOrderL is where Load methods for each relationship are stored.
type stockOrderL struct{}

var (
	stockOrderColumns               = []string{"stock_order_id", "user_id", "created_at", "updated_at"}
	stockOrderColumnsWithoutDefault = []string{"user_id"}
	stockOrderColumnsWithDefault    = []string{"stock_order_id", "created_at", "updated_at"}
	stockOrderPrimaryKeyColumns     = []string{"stock_order_id"}
)

type (
	// StockOrderSlice is an alias for a slice of pointers to StockOrder.
	// This should generally be used opposed to []StockOrder.
	StockOrderSlice []*StockOrder
	// StockOrderHook is the signature for custom StockOrder hook methods
	StockOrderHook func(boil.Executor, *StockOrder) error

	stockOrderQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockOrderType                 = reflect.TypeOf(&StockOrder{})
	stockOrderMapping              = queries.MakeStructMapping(stockOrderType)
	stockOrderPrimaryKeyMapping, _ = queries.BindMapping(stockOrderType, stockOrderMapping, stockOrderPrimaryKeyColumns)
	stockOrderInsertCacheMut       sync.RWMutex
	stockOrderInsertCache          = make(map[string]insertCache)
	stockOrderUpdateCacheMut       sync.RWMutex
	stockOrderUpdateCache          = make(map[string]updateCache)
	stockOrderUpsertCacheMut       sync.RWMutex
	stockOrderUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var stockOrderBeforeInsertHooks []StockOrderHook
var stockOrderBeforeUpdateHooks []StockOrderHook
var stockOrderBeforeDeleteHooks []StockOrderHook
var stockOrderBeforeUpsertHooks []StockOrderHook

var stockOrderAfterInsertHooks []StockOrderHook
var stockOrderAfterSelectHooks []StockOrderHook
var stockOrderAfterUpdateHooks []StockOrderHook
var stockOrderAfterDeleteHooks []StockOrderHook
var stockOrderAfterUpsertHooks []StockOrderHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *StockOrder) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockOrderBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *StockOrder) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockOrderBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *StockOrder) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockOrderBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *StockOrder) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockOrderBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *StockOrder) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockOrderAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *StockOrder) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range stockOrderAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *StockOrder) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockOrderAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *StockOrder) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockOrderAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *StockOrder) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockOrderAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStockOrderHook registers your hook function for all future operations.
func AddStockOrderHook(hookPoint boil.HookPoint, stockOrderHook StockOrderHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stockOrderBeforeInsertHooks = append(stockOrderBeforeInsertHooks, stockOrderHook)
	case boil.BeforeUpdateHook:
		stockOrderBeforeUpdateHooks = append(stockOrderBeforeUpdateHooks, stockOrderHook)
	case boil.BeforeDeleteHook:
		stockOrderBeforeDeleteHooks = append(stockOrderBeforeDeleteHooks, stockOrderHook)
	case boil.BeforeUpsertHook:
		stockOrderBeforeUpsertHooks = append(stockOrderBeforeUpsertHooks, stockOrderHook)
	case boil.AfterInsertHook:
		stockOrderAfterInsertHooks = append(stockOrderAfterInsertHooks, stockOrderHook)
	case boil.AfterSelectHook:
		stockOrderAfterSelectHooks = append(stockOrderAfterSelectHooks, stockOrderHook)
	case boil.AfterUpdateHook:
		stockOrderAfterUpdateHooks = append(stockOrderAfterUpdateHooks, stockOrderHook)
	case boil.AfterDeleteHook:
		stockOrderAfterDeleteHooks = append(stockOrderAfterDeleteHooks, stockOrderHook)
	case boil.AfterUpsertHook:
		stockOrderAfterUpsertHooks = append(stockOrderAfterUpsertHooks, stockOrderHook)
	}
}

// OneP returns a single stockOrder record from the query, and panics on error.
func (q stockOrderQuery) OneP() *StockOrder {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single stockOrder record from the query.
func (q stockOrderQuery) One() (*StockOrder, error) {
	o := &StockOrder{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for stock_order")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all StockOrder records from the query, and panics on error.
func (q stockOrderQuery) AllP() StockOrderSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all StockOrder records from the query.
func (q stockOrderQuery) All() (StockOrderSlice, error) {
	var o StockOrderSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to StockOrder slice")
	}

	if len(stockOrderAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all StockOrder records in the query, and panics on error.
func (q stockOrderQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all StockOrder records in the query.
func (q stockOrderQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count stock_order rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q stockOrderQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q stockOrderQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if stock_order exists")
	}

	return count > 0, nil
}

// UserG pointed to by the foreign key.
func (o *StockOrder) UserG(mods ...qm.QueryMod) authUserQuery {
	return o.User(boil.GetDB(), mods...)
}

// User pointed to by the foreign key.
func (o *StockOrder) User(exec boil.Executor, mods ...qm.QueryMod) authUserQuery {
	queryMods := []qm.QueryMod{
		qm.Where("auth_user_id=$1", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := AuthUsers(exec, queryMods...)
	queries.SetFrom(query.Query, "\"auth_user\"")

	return query
}

// OrderStockItemOrderG pointed to by the foreign key.
func (o *StockOrder) OrderStockItemOrderG(mods ...qm.QueryMod) stockItemOrderQuery {
	return o.OrderStockItemOrder(boil.GetDB(), mods...)
}

// OrderStockItemOrder pointed to by the foreign key.
func (o *StockOrder) OrderStockItemOrder(exec boil.Executor, mods ...qm.QueryMod) stockItemOrderQuery {
	queryMods := []qm.QueryMod{
		qm.Where("order_id=$1", o.StockOrderID),
	}

	queryMods = append(queryMods, mods...)

	query := StockItemOrders(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_item_order\"")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockOrderL) LoadUser(e boil.Executor, singular bool, maybeStockOrder interface{}) error {
	var slice []*StockOrder
	var object *StockOrder

	count := 1
	if singular {
		object = maybeStockOrder.(*StockOrder)
	} else {
		slice = *maybeStockOrder.(*StockOrderSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockOrderR{}
		args[0] = object.UserID
	} else {
		for i, obj := range slice {
			obj.R = &stockOrderR{}
			args[i] = obj.UserID
		}
	}

	query := fmt.Sprintf(
		"select * from \"auth_user\" where \"auth_user_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AuthUser")
	}
	defer results.Close()

	var resultSlice []*AuthUser
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AuthUser")
	}

	if len(stockOrderAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.User = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.UserID == foreign.AuthUserID {
				local.R.User = foreign
				break
			}
		}
	}

	return nil
}

// LoadOrderStockItemOrder allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockOrderL) LoadOrderStockItemOrder(e boil.Executor, singular bool, maybeStockOrder interface{}) error {
	var slice []*StockOrder
	var object *StockOrder

	count := 1
	if singular {
		object = maybeStockOrder.(*StockOrder)
	} else {
		slice = *maybeStockOrder.(*StockOrderSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockOrderR{}
		args[0] = object.StockOrderID
	} else {
		for i, obj := range slice {
			obj.R = &stockOrderR{}
			args[i] = obj.StockOrderID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_item_order\" where \"order_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockItemOrder")
	}
	defer results.Close()

	var resultSlice []*StockItemOrder
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockItemOrder")
	}

	if len(stockOrderAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.OrderStockItemOrder = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockOrderID == foreign.OrderID {
				local.R.OrderStockItemOrder = foreign
				break
			}
		}
	}

	return nil
}

// SetUser of the stock_order to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserStockOrders.
func (o *StockOrder) SetUser(exec boil.Executor, insert bool, related *AuthUser) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_order\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockOrderPrimaryKeyColumns),
	)
	values := []interface{}{related.AuthUserID, o.StockOrderID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.AuthUserID

	if o.R == nil {
		o.R = &stockOrderR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &authUserR{
			UserStockOrders: StockOrderSlice{o},
		}
	} else {
		related.R.UserStockOrders = append(related.R.UserStockOrders, o)
	}

	return nil
}

// SetOrderStockItemOrder of the stock_order to the related item.
// Sets o.R.OrderStockItemOrder to related.
// Adds o to related.R.Order.
func (o *StockOrder) SetOrderStockItemOrder(exec boil.Executor, insert bool, related *StockItemOrder) error {
	var err error

	if insert {
		related.OrderID = o.StockOrderID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock_item_order\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"order_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockItemOrderPrimaryKeyColumns),
		)
		values := []interface{}{o.StockOrderID, related.StockItemOrderID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.OrderID = o.StockOrderID

	}

	if o.R == nil {
		o.R = &stockOrderR{
			OrderStockItemOrder: related,
		}
	} else {
		o.R.OrderStockItemOrder = related
	}

	if related.R == nil {
		related.R = &stockItemOrderR{
			Order: o,
		}
	} else {
		related.R.Order = o
	}
	return nil
}

// StockOrdersG retrieves all records.
func StockOrdersG(mods ...qm.QueryMod) stockOrderQuery {
	return StockOrders(boil.GetDB(), mods...)
}

// StockOrders retrieves all the records using an executor.
func StockOrders(exec boil.Executor, mods ...qm.QueryMod) stockOrderQuery {
	mods = append(mods, qm.From("\"stock_order\""))
	return stockOrderQuery{NewQuery(exec, mods...)}
}

// FindStockOrderG retrieves a single record by ID.
func FindStockOrderG(stockOrderID int, selectCols ...string) (*StockOrder, error) {
	return FindStockOrder(boil.GetDB(), stockOrderID, selectCols...)
}

// FindStockOrderGP retrieves a single record by ID, and panics on error.
func FindStockOrderGP(stockOrderID int, selectCols ...string) *StockOrder {
	retobj, err := FindStockOrder(boil.GetDB(), stockOrderID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindStockOrder retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStockOrder(exec boil.Executor, stockOrderID int, selectCols ...string) (*StockOrder, error) {
	stockOrderObj := &StockOrder{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stock_order\" where \"stock_order_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, stockOrderID)

	err := q.Bind(stockOrderObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from stock_order")
	}

	return stockOrderObj, nil
}

// FindStockOrderP retrieves a single record by ID with an executor, and panics on error.
func FindStockOrderP(exec boil.Executor, stockOrderID int, selectCols ...string) *StockOrder {
	retobj, err := FindStockOrder(exec, stockOrderID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *StockOrder) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *StockOrder) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *StockOrder) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *StockOrder) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no stock_order provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.Time.IsZero() {
		o.CreatedAt.Time = currTime
		o.CreatedAt.Valid = true
	}
	if o.UpdatedAt.Time.IsZero() {
		o.UpdatedAt.Time = currTime
		o.UpdatedAt.Valid = true
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockOrderColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	stockOrderInsertCacheMut.RLock()
	cache, cached := stockOrderInsertCache[key]
	stockOrderInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			stockOrderColumns,
			stockOrderColumnsWithDefault,
			stockOrderColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(stockOrderType, stockOrderMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockOrderType, stockOrderMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"stock_order\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into stock_order")
	}

	if !cached {
		stockOrderInsertCacheMut.Lock()
		stockOrderInsertCache[key] = cache
		stockOrderInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single StockOrder record. See Update for
// whitelist behavior description.
func (o *StockOrder) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single StockOrder record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *StockOrder) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the StockOrder, and panics on error.
// See Update for whitelist behavior description.
func (o *StockOrder) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the StockOrder.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *StockOrder) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	stockOrderUpdateCacheMut.RLock()
	cache, cached := stockOrderUpdateCache[key]
	stockOrderUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(stockOrderColumns, stockOrderPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update stock_order, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stock_order\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stockOrderPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockOrderType, stockOrderMapping, append(wl, stockOrderPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update stock_order row")
	}

	if !cached {
		stockOrderUpdateCacheMut.Lock()
		stockOrderUpdateCache[key] = cache
		stockOrderUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q stockOrderQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q stockOrderQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for stock_order")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StockOrderSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o StockOrderSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o StockOrderSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockOrderSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockOrderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"stock_order\" SET %s WHERE (\"stock_order_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockOrderPrimaryKeyColumns), len(colNames)+1, len(stockOrderPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in stockOrder slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *StockOrder) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *StockOrder) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *StockOrder) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *StockOrder) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no stock_order provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.Time.IsZero() {
		o.CreatedAt.Time = currTime
		o.CreatedAt.Valid = true
	}
	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockOrderColumnsWithDefault, o)

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

	stockOrderUpsertCacheMut.RLock()
	cache, cached := stockOrderUpsertCache[key]
	stockOrderUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			stockOrderColumns,
			stockOrderColumnsWithDefault,
			stockOrderColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			stockOrderColumns,
			stockOrderPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert stock_order, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stockOrderPrimaryKeyColumns))
			copy(conflict, stockOrderPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"stock_order\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(stockOrderType, stockOrderMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockOrderType, stockOrderMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for stock_order")
	}

	if !cached {
		stockOrderUpsertCacheMut.Lock()
		stockOrderUpsertCache[key] = cache
		stockOrderUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single StockOrder record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockOrder) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single StockOrder record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *StockOrder) DeleteG() error {
	if o == nil {
		return errors.New("chado: no StockOrder provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single StockOrder record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockOrder) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single StockOrder record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *StockOrder) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no StockOrder provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockOrderPrimaryKeyMapping)
	sql := "DELETE FROM \"stock_order\" WHERE \"stock_order_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from stock_order")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q stockOrderQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q stockOrderQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no stockOrderQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from stock_order")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o StockOrderSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o StockOrderSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no StockOrder slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o StockOrderSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockOrderSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no StockOrder slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(stockOrderBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockOrderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"stock_order\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockOrderPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockOrderPrimaryKeyColumns), 1, len(stockOrderPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from stockOrder slice")
	}

	if len(stockOrderAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *StockOrder) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *StockOrder) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *StockOrder) ReloadG() error {
	if o == nil {
		return errors.New("chado: no StockOrder provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *StockOrder) Reload(exec boil.Executor) error {
	ret, err := FindStockOrder(exec, o.StockOrderID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockOrderSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockOrderSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockOrderSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty StockOrderSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockOrderSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	stockOrders := StockOrderSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockOrderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"stock_order\".* FROM \"stock_order\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockOrderPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(stockOrderPrimaryKeyColumns), 1, len(stockOrderPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&stockOrders)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in StockOrderSlice")
	}

	*o = stockOrders

	return nil
}

// StockOrderExists checks if the StockOrder row exists.
func StockOrderExists(exec boil.Executor, stockOrderID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"stock_order\" where \"stock_order_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, stockOrderID)
	}

	row := exec.QueryRow(sql, stockOrderID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if stock_order exists")
	}

	return exists, nil
}

// StockOrderExistsG checks if the StockOrder row exists.
func StockOrderExistsG(stockOrderID int) (bool, error) {
	return StockOrderExists(boil.GetDB(), stockOrderID)
}

// StockOrderExistsGP checks if the StockOrder row exists. Panics on error.
func StockOrderExistsGP(stockOrderID int) bool {
	e, err := StockOrderExists(boil.GetDB(), stockOrderID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// StockOrderExistsP checks if the StockOrder row exists. Panics on error.
func StockOrderExistsP(exec boil.Executor, stockOrderID int) bool {
	e, err := StockOrderExists(exec, stockOrderID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
