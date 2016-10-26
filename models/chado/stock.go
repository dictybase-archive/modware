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

// Stock is an object representing the database table.
type Stock struct {
	StockID     int         `boil:"stock_id" json:"stock_id" toml:"stock_id" yaml:"stock_id"`
	DbxrefID    null.Int    `boil:"dbxref_id" json:"dbxref_id,omitempty" toml:"dbxref_id" yaml:"dbxref_id,omitempty"`
	OrganismID  null.Int    `boil:"organism_id" json:"organism_id,omitempty" toml:"organism_id" yaml:"organism_id,omitempty"`
	Name        null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Uniquename  string      `boil:"uniquename" json:"uniquename" toml:"uniquename" yaml:"uniquename"`
	Description null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	TypeID      int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	IsObsolete  bool        `boil:"is_obsolete" json:"is_obsolete" toml:"is_obsolete" yaml:"is_obsolete"`

	R *stockR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// stockR is where relationships are stored.
type stockR struct {
	Organism                 *Organism
	Type                     *Cvterm
	Dbxref                   *Dbxref
	StockPub                 *StockPub
	ItemStockItemOrder       *StockItemOrder
	StockcollectionStock     *StockcollectionStock
	StockCvterm              *StockCvterm
	ObjectStockRelationship  *StockRelationship
	SubjectStockRelationship *StockRelationship
	StockDbxref              *StockDbxref
	StockGenotype            *StockGenotype
	Stockprop                *Stockprop
}

// stockL is where Load methods for each relationship are stored.
type stockL struct{}

var (
	stockColumns               = []string{"stock_id", "dbxref_id", "organism_id", "name", "uniquename", "description", "type_id", "is_obsolete"}
	stockColumnsWithoutDefault = []string{"dbxref_id", "organism_id", "name", "uniquename", "description", "type_id"}
	stockColumnsWithDefault    = []string{"stock_id", "is_obsolete"}
	stockPrimaryKeyColumns     = []string{"stock_id"}
)

type (
	// StockSlice is an alias for a slice of pointers to Stock.
	// This should generally be used opposed to []Stock.
	StockSlice []*Stock
	// StockHook is the signature for custom Stock hook methods
	StockHook func(boil.Executor, *Stock) error

	stockQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockType                 = reflect.TypeOf(&Stock{})
	stockMapping              = queries.MakeStructMapping(stockType)
	stockPrimaryKeyMapping, _ = queries.BindMapping(stockType, stockMapping, stockPrimaryKeyColumns)
	stockInsertCacheMut       sync.RWMutex
	stockInsertCache          = make(map[string]insertCache)
	stockUpdateCacheMut       sync.RWMutex
	stockUpdateCache          = make(map[string]updateCache)
	stockUpsertCacheMut       sync.RWMutex
	stockUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var stockBeforeInsertHooks []StockHook
var stockBeforeUpdateHooks []StockHook
var stockBeforeDeleteHooks []StockHook
var stockBeforeUpsertHooks []StockHook

var stockAfterInsertHooks []StockHook
var stockAfterSelectHooks []StockHook
var stockAfterUpdateHooks []StockHook
var stockAfterDeleteHooks []StockHook
var stockAfterUpsertHooks []StockHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Stock) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Stock) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Stock) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Stock) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Stock) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Stock) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range stockAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Stock) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Stock) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Stock) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStockHook registers your hook function for all future operations.
func AddStockHook(hookPoint boil.HookPoint, stockHook StockHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stockBeforeInsertHooks = append(stockBeforeInsertHooks, stockHook)
	case boil.BeforeUpdateHook:
		stockBeforeUpdateHooks = append(stockBeforeUpdateHooks, stockHook)
	case boil.BeforeDeleteHook:
		stockBeforeDeleteHooks = append(stockBeforeDeleteHooks, stockHook)
	case boil.BeforeUpsertHook:
		stockBeforeUpsertHooks = append(stockBeforeUpsertHooks, stockHook)
	case boil.AfterInsertHook:
		stockAfterInsertHooks = append(stockAfterInsertHooks, stockHook)
	case boil.AfterSelectHook:
		stockAfterSelectHooks = append(stockAfterSelectHooks, stockHook)
	case boil.AfterUpdateHook:
		stockAfterUpdateHooks = append(stockAfterUpdateHooks, stockHook)
	case boil.AfterDeleteHook:
		stockAfterDeleteHooks = append(stockAfterDeleteHooks, stockHook)
	case boil.AfterUpsertHook:
		stockAfterUpsertHooks = append(stockAfterUpsertHooks, stockHook)
	}
}

// OneP returns a single stock record from the query, and panics on error.
func (q stockQuery) OneP() *Stock {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single stock record from the query.
func (q stockQuery) One() (*Stock, error) {
	o := &Stock{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for stock")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Stock records from the query, and panics on error.
func (q stockQuery) AllP() StockSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Stock records from the query.
func (q stockQuery) All() (StockSlice, error) {
	var o StockSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to Stock slice")
	}

	if len(stockAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Stock records in the query, and panics on error.
func (q stockQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Stock records in the query.
func (q stockQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count stock rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q stockQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q stockQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if stock exists")
	}

	return count > 0, nil
}

// OrganismG pointed to by the foreign key.
func (o *Stock) OrganismG(mods ...qm.QueryMod) organismQuery {
	return o.Organism(boil.GetDB(), mods...)
}

// Organism pointed to by the foreign key.
func (o *Stock) Organism(exec boil.Executor, mods ...qm.QueryMod) organismQuery {
	queryMods := []qm.QueryMod{
		qm.Where("organism_id=$1", o.OrganismID),
	}

	queryMods = append(queryMods, mods...)

	query := Organisms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"organism\"")

	return query
}

// TypeG pointed to by the foreign key.
func (o *Stock) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Stock) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// DbxrefG pointed to by the foreign key.
func (o *Stock) DbxrefG(mods ...qm.QueryMod) dbxrefQuery {
	return o.Dbxref(boil.GetDB(), mods...)
}

// Dbxref pointed to by the foreign key.
func (o *Stock) Dbxref(exec boil.Executor, mods ...qm.QueryMod) dbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("dbxref_id=$1", o.DbxrefID),
	}

	queryMods = append(queryMods, mods...)

	query := Dbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"dbxref\"")

	return query
}

// StockPubG pointed to by the foreign key.
func (o *Stock) StockPubG(mods ...qm.QueryMod) stockPubQuery {
	return o.StockPub(boil.GetDB(), mods...)
}

// StockPub pointed to by the foreign key.
func (o *Stock) StockPub(exec boil.Executor, mods ...qm.QueryMod) stockPubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_id=$1", o.StockID),
	}

	queryMods = append(queryMods, mods...)

	query := StockPubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_pub\"")

	return query
}

// ItemStockItemOrderG pointed to by the foreign key.
func (o *Stock) ItemStockItemOrderG(mods ...qm.QueryMod) stockItemOrderQuery {
	return o.ItemStockItemOrder(boil.GetDB(), mods...)
}

// ItemStockItemOrder pointed to by the foreign key.
func (o *Stock) ItemStockItemOrder(exec boil.Executor, mods ...qm.QueryMod) stockItemOrderQuery {
	queryMods := []qm.QueryMod{
		qm.Where("item_id=$1", o.StockID),
	}

	queryMods = append(queryMods, mods...)

	query := StockItemOrders(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_item_order\"")

	return query
}

// StockcollectionStockG pointed to by the foreign key.
func (o *Stock) StockcollectionStockG(mods ...qm.QueryMod) stockcollectionStockQuery {
	return o.StockcollectionStock(boil.GetDB(), mods...)
}

// StockcollectionStock pointed to by the foreign key.
func (o *Stock) StockcollectionStock(exec boil.Executor, mods ...qm.QueryMod) stockcollectionStockQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_id=$1", o.StockID),
	}

	queryMods = append(queryMods, mods...)

	query := StockcollectionStocks(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stockcollection_stock\"")

	return query
}

// StockCvtermG pointed to by the foreign key.
func (o *Stock) StockCvtermG(mods ...qm.QueryMod) stockCvtermQuery {
	return o.StockCvterm(boil.GetDB(), mods...)
}

// StockCvterm pointed to by the foreign key.
func (o *Stock) StockCvterm(exec boil.Executor, mods ...qm.QueryMod) stockCvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_id=$1", o.StockID),
	}

	queryMods = append(queryMods, mods...)

	query := StockCvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_cvterm\"")

	return query
}

// ObjectStockRelationshipG pointed to by the foreign key.
func (o *Stock) ObjectStockRelationshipG(mods ...qm.QueryMod) stockRelationshipQuery {
	return o.ObjectStockRelationship(boil.GetDB(), mods...)
}

// ObjectStockRelationship pointed to by the foreign key.
func (o *Stock) ObjectStockRelationship(exec boil.Executor, mods ...qm.QueryMod) stockRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("object_id=$1", o.StockID),
	}

	queryMods = append(queryMods, mods...)

	query := StockRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_relationship\"")

	return query
}

// SubjectStockRelationshipG pointed to by the foreign key.
func (o *Stock) SubjectStockRelationshipG(mods ...qm.QueryMod) stockRelationshipQuery {
	return o.SubjectStockRelationship(boil.GetDB(), mods...)
}

// SubjectStockRelationship pointed to by the foreign key.
func (o *Stock) SubjectStockRelationship(exec boil.Executor, mods ...qm.QueryMod) stockRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("subject_id=$1", o.StockID),
	}

	queryMods = append(queryMods, mods...)

	query := StockRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_relationship\"")

	return query
}

// StockDbxrefG pointed to by the foreign key.
func (o *Stock) StockDbxrefG(mods ...qm.QueryMod) stockDbxrefQuery {
	return o.StockDbxref(boil.GetDB(), mods...)
}

// StockDbxref pointed to by the foreign key.
func (o *Stock) StockDbxref(exec boil.Executor, mods ...qm.QueryMod) stockDbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_id=$1", o.StockID),
	}

	queryMods = append(queryMods, mods...)

	query := StockDbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_dbxref\"")

	return query
}

// StockGenotypeG pointed to by the foreign key.
func (o *Stock) StockGenotypeG(mods ...qm.QueryMod) stockGenotypeQuery {
	return o.StockGenotype(boil.GetDB(), mods...)
}

// StockGenotype pointed to by the foreign key.
func (o *Stock) StockGenotype(exec boil.Executor, mods ...qm.QueryMod) stockGenotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_id=$1", o.StockID),
	}

	queryMods = append(queryMods, mods...)

	query := StockGenotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_genotype\"")

	return query
}

// StockpropG pointed to by the foreign key.
func (o *Stock) StockpropG(mods ...qm.QueryMod) stockpropQuery {
	return o.Stockprop(boil.GetDB(), mods...)
}

// Stockprop pointed to by the foreign key.
func (o *Stock) Stockprop(exec boil.Executor, mods ...qm.QueryMod) stockpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_id=$1", o.StockID),
	}

	queryMods = append(queryMods, mods...)

	query := Stockprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stockprop\"")

	return query
}

// LoadOrganism allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockL) LoadOrganism(e boil.Executor, singular bool, maybeStock interface{}) error {
	var slice []*Stock
	var object *Stock

	count := 1
	if singular {
		object = maybeStock.(*Stock)
	} else {
		slice = *maybeStock.(*StockSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockR{}
		args[0] = object.OrganismID
	} else {
		for i, obj := range slice {
			obj.R = &stockR{}
			args[i] = obj.OrganismID
		}
	}

	query := fmt.Sprintf(
		"select * from \"organism\" where \"organism_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Organism")
	}
	defer results.Close()

	var resultSlice []*Organism
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Organism")
	}

	if len(stockAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Organism = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.OrganismID.Int == foreign.OrganismID {
				local.R.Organism = foreign
				break
			}
		}
	}

	return nil
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockL) LoadType(e boil.Executor, singular bool, maybeStock interface{}) error {
	var slice []*Stock
	var object *Stock

	count := 1
	if singular {
		object = maybeStock.(*Stock)
	} else {
		slice = *maybeStock.(*StockSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &stockR{}
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

	if len(stockAfterSelectHooks) != 0 {
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

// LoadDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockL) LoadDbxref(e boil.Executor, singular bool, maybeStock interface{}) error {
	var slice []*Stock
	var object *Stock

	count := 1
	if singular {
		object = maybeStock.(*Stock)
	} else {
		slice = *maybeStock.(*StockSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockR{}
		args[0] = object.DbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &stockR{}
			args[i] = obj.DbxrefID
		}
	}

	query := fmt.Sprintf(
		"select * from \"dbxref\" where \"dbxref_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Dbxref")
	}
	defer results.Close()

	var resultSlice []*Dbxref
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Dbxref")
	}

	if len(stockAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Dbxref = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.DbxrefID.Int == foreign.DbxrefID {
				local.R.Dbxref = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockL) LoadStockPub(e boil.Executor, singular bool, maybeStock interface{}) error {
	var slice []*Stock
	var object *Stock

	count := 1
	if singular {
		object = maybeStock.(*Stock)
	} else {
		slice = *maybeStock.(*StockSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockR{}
		args[0] = object.StockID
	} else {
		for i, obj := range slice {
			obj.R = &stockR{}
			args[i] = obj.StockID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_pub\" where \"stock_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockPub")
	}
	defer results.Close()

	var resultSlice []*StockPub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockPub")
	}

	if len(stockAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.StockPub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockID == foreign.StockID {
				local.R.StockPub = foreign
				break
			}
		}
	}

	return nil
}

// LoadItemStockItemOrder allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockL) LoadItemStockItemOrder(e boil.Executor, singular bool, maybeStock interface{}) error {
	var slice []*Stock
	var object *Stock

	count := 1
	if singular {
		object = maybeStock.(*Stock)
	} else {
		slice = *maybeStock.(*StockSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockR{}
		args[0] = object.StockID
	} else {
		for i, obj := range slice {
			obj.R = &stockR{}
			args[i] = obj.StockID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_item_order\" where \"item_id\" in (%s)",
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

	if len(stockAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.ItemStockItemOrder = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockID == foreign.ItemID {
				local.R.ItemStockItemOrder = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockcollectionStock allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockL) LoadStockcollectionStock(e boil.Executor, singular bool, maybeStock interface{}) error {
	var slice []*Stock
	var object *Stock

	count := 1
	if singular {
		object = maybeStock.(*Stock)
	} else {
		slice = *maybeStock.(*StockSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockR{}
		args[0] = object.StockID
	} else {
		for i, obj := range slice {
			obj.R = &stockR{}
			args[i] = obj.StockID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stockcollection_stock\" where \"stock_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockcollectionStock")
	}
	defer results.Close()

	var resultSlice []*StockcollectionStock
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockcollectionStock")
	}

	if len(stockAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.StockcollectionStock = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockID == foreign.StockID {
				local.R.StockcollectionStock = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockL) LoadStockCvterm(e boil.Executor, singular bool, maybeStock interface{}) error {
	var slice []*Stock
	var object *Stock

	count := 1
	if singular {
		object = maybeStock.(*Stock)
	} else {
		slice = *maybeStock.(*StockSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockR{}
		args[0] = object.StockID
	} else {
		for i, obj := range slice {
			obj.R = &stockR{}
			args[i] = obj.StockID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_cvterm\" where \"stock_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockCvterm")
	}
	defer results.Close()

	var resultSlice []*StockCvterm
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockCvterm")
	}

	if len(stockAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.StockCvterm = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockID == foreign.StockID {
				local.R.StockCvterm = foreign
				break
			}
		}
	}

	return nil
}

// LoadObjectStockRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockL) LoadObjectStockRelationship(e boil.Executor, singular bool, maybeStock interface{}) error {
	var slice []*Stock
	var object *Stock

	count := 1
	if singular {
		object = maybeStock.(*Stock)
	} else {
		slice = *maybeStock.(*StockSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockR{}
		args[0] = object.StockID
	} else {
		for i, obj := range slice {
			obj.R = &stockR{}
			args[i] = obj.StockID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_relationship\" where \"object_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockRelationship")
	}
	defer results.Close()

	var resultSlice []*StockRelationship
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockRelationship")
	}

	if len(stockAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.ObjectStockRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockID == foreign.ObjectID {
				local.R.ObjectStockRelationship = foreign
				break
			}
		}
	}

	return nil
}

// LoadSubjectStockRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockL) LoadSubjectStockRelationship(e boil.Executor, singular bool, maybeStock interface{}) error {
	var slice []*Stock
	var object *Stock

	count := 1
	if singular {
		object = maybeStock.(*Stock)
	} else {
		slice = *maybeStock.(*StockSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockR{}
		args[0] = object.StockID
	} else {
		for i, obj := range slice {
			obj.R = &stockR{}
			args[i] = obj.StockID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_relationship\" where \"subject_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockRelationship")
	}
	defer results.Close()

	var resultSlice []*StockRelationship
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockRelationship")
	}

	if len(stockAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.SubjectStockRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockID == foreign.SubjectID {
				local.R.SubjectStockRelationship = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockL) LoadStockDbxref(e boil.Executor, singular bool, maybeStock interface{}) error {
	var slice []*Stock
	var object *Stock

	count := 1
	if singular {
		object = maybeStock.(*Stock)
	} else {
		slice = *maybeStock.(*StockSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockR{}
		args[0] = object.StockID
	} else {
		for i, obj := range slice {
			obj.R = &stockR{}
			args[i] = obj.StockID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_dbxref\" where \"stock_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockDbxref")
	}
	defer results.Close()

	var resultSlice []*StockDbxref
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockDbxref")
	}

	if len(stockAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.StockDbxref = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockID == foreign.StockID {
				local.R.StockDbxref = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockGenotype allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockL) LoadStockGenotype(e boil.Executor, singular bool, maybeStock interface{}) error {
	var slice []*Stock
	var object *Stock

	count := 1
	if singular {
		object = maybeStock.(*Stock)
	} else {
		slice = *maybeStock.(*StockSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockR{}
		args[0] = object.StockID
	} else {
		for i, obj := range slice {
			obj.R = &stockR{}
			args[i] = obj.StockID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_genotype\" where \"stock_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockGenotype")
	}
	defer results.Close()

	var resultSlice []*StockGenotype
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockGenotype")
	}

	if len(stockAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.StockGenotype = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockID == foreign.StockID {
				local.R.StockGenotype = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockL) LoadStockprop(e boil.Executor, singular bool, maybeStock interface{}) error {
	var slice []*Stock
	var object *Stock

	count := 1
	if singular {
		object = maybeStock.(*Stock)
	} else {
		slice = *maybeStock.(*StockSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockR{}
		args[0] = object.StockID
	} else {
		for i, obj := range slice {
			obj.R = &stockR{}
			args[i] = obj.StockID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stockprop\" where \"stock_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Stockprop")
	}
	defer results.Close()

	var resultSlice []*Stockprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Stockprop")
	}

	if len(stockAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Stockprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockID == foreign.StockID {
				local.R.Stockprop = foreign
				break
			}
		}
	}

	return nil
}

// SetOrganism of the stock to the related item.
// Sets o.R.Organism to related.
// Adds o to related.R.Stock.
func (o *Stock) SetOrganism(exec boil.Executor, insert bool, related *Organism) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"organism_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockPrimaryKeyColumns),
	)
	values := []interface{}{related.OrganismID, o.StockID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OrganismID.Int = related.OrganismID
	o.OrganismID.Valid = true

	if o.R == nil {
		o.R = &stockR{
			Organism: related,
		}
	} else {
		o.R.Organism = related
	}

	if related.R == nil {
		related.R = &organismR{
			Stock: o,
		}
	} else {
		related.R.Stock = o
	}

	return nil
}

// RemoveOrganism relationship.
// Sets o.R.Organism to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Stock) RemoveOrganism(exec boil.Executor, related *Organism) error {
	var err error

	o.OrganismID.Valid = false
	if err = o.Update(exec, "organism_id"); err != nil {
		o.OrganismID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Organism = nil
	if related == nil || related.R == nil {
		return nil
	}

	related.R.Stock = nil
	return nil
}

// SetType of the stock to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeStock.
func (o *Stock) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.StockID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &stockR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeStock: o,
		}
	} else {
		related.R.TypeStock = o
	}

	return nil
}

// SetDbxref of the stock to the related item.
// Sets o.R.Dbxref to related.
// Adds o to related.R.Stocks.
func (o *Stock) SetDbxref(exec boil.Executor, insert bool, related *Dbxref) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"dbxref_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockPrimaryKeyColumns),
	)
	values := []interface{}{related.DbxrefID, o.StockID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.DbxrefID.Int = related.DbxrefID
	o.DbxrefID.Valid = true

	if o.R == nil {
		o.R = &stockR{
			Dbxref: related,
		}
	} else {
		o.R.Dbxref = related
	}

	if related.R == nil {
		related.R = &dbxrefR{
			Stocks: StockSlice{o},
		}
	} else {
		related.R.Stocks = append(related.R.Stocks, o)
	}

	return nil
}

// RemoveDbxref relationship.
// Sets o.R.Dbxref to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Stock) RemoveDbxref(exec boil.Executor, related *Dbxref) error {
	var err error

	o.DbxrefID.Valid = false
	if err = o.Update(exec, "dbxref_id"); err != nil {
		o.DbxrefID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Dbxref = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Stocks {
		if o.DbxrefID.Int != ri.DbxrefID.Int {
			continue
		}

		ln := len(related.R.Stocks)
		if ln > 1 && i < ln-1 {
			related.R.Stocks[i] = related.R.Stocks[ln-1]
		}
		related.R.Stocks = related.R.Stocks[:ln-1]
		break
	}
	return nil
}

// SetStockPub of the stock to the related item.
// Sets o.R.StockPub to related.
// Adds o to related.R.Stock.
func (o *Stock) SetStockPub(exec boil.Executor, insert bool, related *StockPub) error {
	var err error

	if insert {
		related.StockID = o.StockID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"stock_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockPubPrimaryKeyColumns),
		)
		values := []interface{}{o.StockID, related.StockPubID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.StockID = o.StockID

	}

	if o.R == nil {
		o.R = &stockR{
			StockPub: related,
		}
	} else {
		o.R.StockPub = related
	}

	if related.R == nil {
		related.R = &stockPubR{
			Stock: o,
		}
	} else {
		related.R.Stock = o
	}
	return nil
}

// SetItemStockItemOrder of the stock to the related item.
// Sets o.R.ItemStockItemOrder to related.
// Adds o to related.R.Item.
func (o *Stock) SetItemStockItemOrder(exec boil.Executor, insert bool, related *StockItemOrder) error {
	var err error

	if insert {
		related.ItemID = o.StockID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock_item_order\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"item_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockItemOrderPrimaryKeyColumns),
		)
		values := []interface{}{o.StockID, related.StockItemOrderID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.ItemID = o.StockID

	}

	if o.R == nil {
		o.R = &stockR{
			ItemStockItemOrder: related,
		}
	} else {
		o.R.ItemStockItemOrder = related
	}

	if related.R == nil {
		related.R = &stockItemOrderR{
			Item: o,
		}
	} else {
		related.R.Item = o
	}
	return nil
}

// SetStockcollectionStock of the stock to the related item.
// Sets o.R.StockcollectionStock to related.
// Adds o to related.R.Stock.
func (o *Stock) SetStockcollectionStock(exec boil.Executor, insert bool, related *StockcollectionStock) error {
	var err error

	if insert {
		related.StockID = o.StockID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stockcollection_stock\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"stock_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockcollectionStockPrimaryKeyColumns),
		)
		values := []interface{}{o.StockID, related.StockcollectionStockID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.StockID = o.StockID

	}

	if o.R == nil {
		o.R = &stockR{
			StockcollectionStock: related,
		}
	} else {
		o.R.StockcollectionStock = related
	}

	if related.R == nil {
		related.R = &stockcollectionStockR{
			Stock: o,
		}
	} else {
		related.R.Stock = o
	}
	return nil
}

// SetStockCvterm of the stock to the related item.
// Sets o.R.StockCvterm to related.
// Adds o to related.R.Stock.
func (o *Stock) SetStockCvterm(exec boil.Executor, insert bool, related *StockCvterm) error {
	var err error

	if insert {
		related.StockID = o.StockID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock_cvterm\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"stock_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockCvtermPrimaryKeyColumns),
		)
		values := []interface{}{o.StockID, related.StockCvtermID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.StockID = o.StockID

	}

	if o.R == nil {
		o.R = &stockR{
			StockCvterm: related,
		}
	} else {
		o.R.StockCvterm = related
	}

	if related.R == nil {
		related.R = &stockCvtermR{
			Stock: o,
		}
	} else {
		related.R.Stock = o
	}
	return nil
}

// SetObjectStockRelationship of the stock to the related item.
// Sets o.R.ObjectStockRelationship to related.
// Adds o to related.R.Object.
func (o *Stock) SetObjectStockRelationship(exec boil.Executor, insert bool, related *StockRelationship) error {
	var err error

	if insert {
		related.ObjectID = o.StockID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"object_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockRelationshipPrimaryKeyColumns),
		)
		values := []interface{}{o.StockID, related.StockRelationshipID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.ObjectID = o.StockID

	}

	if o.R == nil {
		o.R = &stockR{
			ObjectStockRelationship: related,
		}
	} else {
		o.R.ObjectStockRelationship = related
	}

	if related.R == nil {
		related.R = &stockRelationshipR{
			Object: o,
		}
	} else {
		related.R.Object = o
	}
	return nil
}

// SetSubjectStockRelationship of the stock to the related item.
// Sets o.R.SubjectStockRelationship to related.
// Adds o to related.R.Subject.
func (o *Stock) SetSubjectStockRelationship(exec boil.Executor, insert bool, related *StockRelationship) error {
	var err error

	if insert {
		related.SubjectID = o.StockID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"subject_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockRelationshipPrimaryKeyColumns),
		)
		values := []interface{}{o.StockID, related.StockRelationshipID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.SubjectID = o.StockID

	}

	if o.R == nil {
		o.R = &stockR{
			SubjectStockRelationship: related,
		}
	} else {
		o.R.SubjectStockRelationship = related
	}

	if related.R == nil {
		related.R = &stockRelationshipR{
			Subject: o,
		}
	} else {
		related.R.Subject = o
	}
	return nil
}

// SetStockDbxref of the stock to the related item.
// Sets o.R.StockDbxref to related.
// Adds o to related.R.Stock.
func (o *Stock) SetStockDbxref(exec boil.Executor, insert bool, related *StockDbxref) error {
	var err error

	if insert {
		related.StockID = o.StockID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock_dbxref\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"stock_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockDbxrefPrimaryKeyColumns),
		)
		values := []interface{}{o.StockID, related.StockDbxrefID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.StockID = o.StockID

	}

	if o.R == nil {
		o.R = &stockR{
			StockDbxref: related,
		}
	} else {
		o.R.StockDbxref = related
	}

	if related.R == nil {
		related.R = &stockDbxrefR{
			Stock: o,
		}
	} else {
		related.R.Stock = o
	}
	return nil
}

// SetStockGenotype of the stock to the related item.
// Sets o.R.StockGenotype to related.
// Adds o to related.R.Stock.
func (o *Stock) SetStockGenotype(exec boil.Executor, insert bool, related *StockGenotype) error {
	var err error

	if insert {
		related.StockID = o.StockID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock_genotype\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"stock_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockGenotypePrimaryKeyColumns),
		)
		values := []interface{}{o.StockID, related.StockGenotypeID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.StockID = o.StockID

	}

	if o.R == nil {
		o.R = &stockR{
			StockGenotype: related,
		}
	} else {
		o.R.StockGenotype = related
	}

	if related.R == nil {
		related.R = &stockGenotypeR{
			Stock: o,
		}
	} else {
		related.R.Stock = o
	}
	return nil
}

// SetStockprop of the stock to the related item.
// Sets o.R.Stockprop to related.
// Adds o to related.R.Stock.
func (o *Stock) SetStockprop(exec boil.Executor, insert bool, related *Stockprop) error {
	var err error

	if insert {
		related.StockID = o.StockID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stockprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"stock_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockpropPrimaryKeyColumns),
		)
		values := []interface{}{o.StockID, related.StockpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.StockID = o.StockID

	}

	if o.R == nil {
		o.R = &stockR{
			Stockprop: related,
		}
	} else {
		o.R.Stockprop = related
	}

	if related.R == nil {
		related.R = &stockpropR{
			Stock: o,
		}
	} else {
		related.R.Stock = o
	}
	return nil
}

// StocksG retrieves all records.
func StocksG(mods ...qm.QueryMod) stockQuery {
	return Stocks(boil.GetDB(), mods...)
}

// Stocks retrieves all the records using an executor.
func Stocks(exec boil.Executor, mods ...qm.QueryMod) stockQuery {
	mods = append(mods, qm.From("\"stock\""))
	return stockQuery{NewQuery(exec, mods...)}
}

// FindStockG retrieves a single record by ID.
func FindStockG(stockID int, selectCols ...string) (*Stock, error) {
	return FindStock(boil.GetDB(), stockID, selectCols...)
}

// FindStockGP retrieves a single record by ID, and panics on error.
func FindStockGP(stockID int, selectCols ...string) *Stock {
	retobj, err := FindStock(boil.GetDB(), stockID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindStock retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStock(exec boil.Executor, stockID int, selectCols ...string) (*Stock, error) {
	stockObj := &Stock{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stock\" where \"stock_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, stockID)

	err := q.Bind(stockObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from stock")
	}

	return stockObj, nil
}

// FindStockP retrieves a single record by ID with an executor, and panics on error.
func FindStockP(exec boil.Executor, stockID int, selectCols ...string) *Stock {
	retobj, err := FindStock(exec, stockID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Stock) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Stock) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Stock) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Stock) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no stock provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	stockInsertCacheMut.RLock()
	cache, cached := stockInsertCache[key]
	stockInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			stockColumns,
			stockColumnsWithDefault,
			stockColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(stockType, stockMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockType, stockMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"stock\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into stock")
	}

	if !cached {
		stockInsertCacheMut.Lock()
		stockInsertCache[key] = cache
		stockInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Stock record. See Update for
// whitelist behavior description.
func (o *Stock) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Stock record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Stock) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Stock, and panics on error.
// See Update for whitelist behavior description.
func (o *Stock) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Stock.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Stock) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	stockUpdateCacheMut.RLock()
	cache, cached := stockUpdateCache[key]
	stockUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(stockColumns, stockPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update stock, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stock\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stockPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockType, stockMapping, append(wl, stockPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update stock row")
	}

	if !cached {
		stockUpdateCacheMut.Lock()
		stockUpdateCache[key] = cache
		stockUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q stockQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q stockQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for stock")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StockSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o StockSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o StockSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"stock\" SET %s WHERE (\"stock_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockPrimaryKeyColumns), len(colNames)+1, len(stockPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in stock slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Stock) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Stock) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Stock) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Stock) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no stock provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockColumnsWithDefault, o)

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

	stockUpsertCacheMut.RLock()
	cache, cached := stockUpsertCache[key]
	stockUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			stockColumns,
			stockColumnsWithDefault,
			stockColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			stockColumns,
			stockPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert stock, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stockPrimaryKeyColumns))
			copy(conflict, stockPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"stock\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(stockType, stockMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockType, stockMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for stock")
	}

	if !cached {
		stockUpsertCacheMut.Lock()
		stockUpsertCache[key] = cache
		stockUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Stock record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Stock) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Stock record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Stock) DeleteG() error {
	if o == nil {
		return errors.New("chado: no Stock provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Stock record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Stock) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Stock record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Stock) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Stock provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockPrimaryKeyMapping)
	sql := "DELETE FROM \"stock\" WHERE \"stock_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from stock")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q stockQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q stockQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no stockQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from stock")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o StockSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o StockSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no Stock slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o StockSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Stock slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(stockBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"stock\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockPrimaryKeyColumns), 1, len(stockPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from stock slice")
	}

	if len(stockAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Stock) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Stock) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Stock) ReloadG() error {
	if o == nil {
		return errors.New("chado: no Stock provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Stock) Reload(exec boil.Executor) error {
	ret, err := FindStock(exec, o.StockID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty StockSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	stocks := StockSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"stock\".* FROM \"stock\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(stockPrimaryKeyColumns), 1, len(stockPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&stocks)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in StockSlice")
	}

	*o = stocks

	return nil
}

// StockExists checks if the Stock row exists.
func StockExists(exec boil.Executor, stockID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"stock\" where \"stock_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, stockID)
	}

	row := exec.QueryRow(sql, stockID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if stock exists")
	}

	return exists, nil
}

// StockExistsG checks if the Stock row exists.
func StockExistsG(stockID int) (bool, error) {
	return StockExists(boil.GetDB(), stockID)
}

// StockExistsGP checks if the Stock row exists. Panics on error.
func StockExistsGP(stockID int) bool {
	e, err := StockExists(boil.GetDB(), stockID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// StockExistsP checks if the Stock row exists. Panics on error.
func StockExistsP(exec boil.Executor, stockID int) bool {
	e, err := StockExists(exec, stockID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
