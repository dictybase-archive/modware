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

// StockRelationshipCvterm is an object representing the database table.
type StockRelationshipCvterm struct {
	StockRelationshipCvtermID int      `boil:"stock_relationship_cvterm_id" json:"stock_relationship_cvterm_id" toml:"stock_relationship_cvterm_id" yaml:"stock_relationship_cvterm_id"`
	StockRelationshipID       int      `boil:"stock_relationship_id" json:"stock_relationship_id" toml:"stock_relationship_id" yaml:"stock_relationship_id"`
	CvtermID                  int      `boil:"cvterm_id" json:"cvterm_id" toml:"cvterm_id" yaml:"cvterm_id"`
	PubID                     null.Int `boil:"pub_id" json:"pub_id,omitempty" toml:"pub_id" yaml:"pub_id,omitempty"`

	R *stockRelationshipCvtermR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockRelationshipCvtermL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// stockRelationshipCvtermR is where relationships are stored.
type stockRelationshipCvtermR struct {
	Cvterm            *Cvterm
	Pub               *Pub
	StockRelationship *StockRelationship
}

// stockRelationshipCvtermL is where Load methods for each relationship are stored.
type stockRelationshipCvtermL struct{}

var (
	stockRelationshipCvtermColumns               = []string{"stock_relationship_cvterm_id", "stock_relationship_id", "cvterm_id", "pub_id"}
	stockRelationshipCvtermColumnsWithoutDefault = []string{"stock_relationship_id", "cvterm_id", "pub_id"}
	stockRelationshipCvtermColumnsWithDefault    = []string{"stock_relationship_cvterm_id"}
	stockRelationshipCvtermPrimaryKeyColumns     = []string{"stock_relationship_cvterm_id"}
)

type (
	// StockRelationshipCvtermSlice is an alias for a slice of pointers to StockRelationshipCvterm.
	// This should generally be used opposed to []StockRelationshipCvterm.
	StockRelationshipCvtermSlice []*StockRelationshipCvterm
	// StockRelationshipCvtermHook is the signature for custom StockRelationshipCvterm hook methods
	StockRelationshipCvtermHook func(boil.Executor, *StockRelationshipCvterm) error

	stockRelationshipCvtermQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockRelationshipCvtermType                 = reflect.TypeOf(&StockRelationshipCvterm{})
	stockRelationshipCvtermMapping              = queries.MakeStructMapping(stockRelationshipCvtermType)
	stockRelationshipCvtermPrimaryKeyMapping, _ = queries.BindMapping(stockRelationshipCvtermType, stockRelationshipCvtermMapping, stockRelationshipCvtermPrimaryKeyColumns)
	stockRelationshipCvtermInsertCacheMut       sync.RWMutex
	stockRelationshipCvtermInsertCache          = make(map[string]insertCache)
	stockRelationshipCvtermUpdateCacheMut       sync.RWMutex
	stockRelationshipCvtermUpdateCache          = make(map[string]updateCache)
	stockRelationshipCvtermUpsertCacheMut       sync.RWMutex
	stockRelationshipCvtermUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var stockRelationshipCvtermBeforeInsertHooks []StockRelationshipCvtermHook
var stockRelationshipCvtermBeforeUpdateHooks []StockRelationshipCvtermHook
var stockRelationshipCvtermBeforeDeleteHooks []StockRelationshipCvtermHook
var stockRelationshipCvtermBeforeUpsertHooks []StockRelationshipCvtermHook

var stockRelationshipCvtermAfterInsertHooks []StockRelationshipCvtermHook
var stockRelationshipCvtermAfterSelectHooks []StockRelationshipCvtermHook
var stockRelationshipCvtermAfterUpdateHooks []StockRelationshipCvtermHook
var stockRelationshipCvtermAfterDeleteHooks []StockRelationshipCvtermHook
var stockRelationshipCvtermAfterUpsertHooks []StockRelationshipCvtermHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *StockRelationshipCvterm) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipCvtermBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *StockRelationshipCvterm) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipCvtermBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *StockRelationshipCvterm) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipCvtermBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *StockRelationshipCvterm) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipCvtermBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *StockRelationshipCvterm) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipCvtermAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *StockRelationshipCvterm) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipCvtermAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *StockRelationshipCvterm) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipCvtermAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *StockRelationshipCvterm) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipCvtermAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *StockRelationshipCvterm) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipCvtermAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStockRelationshipCvtermHook registers your hook function for all future operations.
func AddStockRelationshipCvtermHook(hookPoint boil.HookPoint, stockRelationshipCvtermHook StockRelationshipCvtermHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stockRelationshipCvtermBeforeInsertHooks = append(stockRelationshipCvtermBeforeInsertHooks, stockRelationshipCvtermHook)
	case boil.BeforeUpdateHook:
		stockRelationshipCvtermBeforeUpdateHooks = append(stockRelationshipCvtermBeforeUpdateHooks, stockRelationshipCvtermHook)
	case boil.BeforeDeleteHook:
		stockRelationshipCvtermBeforeDeleteHooks = append(stockRelationshipCvtermBeforeDeleteHooks, stockRelationshipCvtermHook)
	case boil.BeforeUpsertHook:
		stockRelationshipCvtermBeforeUpsertHooks = append(stockRelationshipCvtermBeforeUpsertHooks, stockRelationshipCvtermHook)
	case boil.AfterInsertHook:
		stockRelationshipCvtermAfterInsertHooks = append(stockRelationshipCvtermAfterInsertHooks, stockRelationshipCvtermHook)
	case boil.AfterSelectHook:
		stockRelationshipCvtermAfterSelectHooks = append(stockRelationshipCvtermAfterSelectHooks, stockRelationshipCvtermHook)
	case boil.AfterUpdateHook:
		stockRelationshipCvtermAfterUpdateHooks = append(stockRelationshipCvtermAfterUpdateHooks, stockRelationshipCvtermHook)
	case boil.AfterDeleteHook:
		stockRelationshipCvtermAfterDeleteHooks = append(stockRelationshipCvtermAfterDeleteHooks, stockRelationshipCvtermHook)
	case boil.AfterUpsertHook:
		stockRelationshipCvtermAfterUpsertHooks = append(stockRelationshipCvtermAfterUpsertHooks, stockRelationshipCvtermHook)
	}
}

// OneP returns a single stockRelationshipCvterm record from the query, and panics on error.
func (q stockRelationshipCvtermQuery) OneP() *StockRelationshipCvterm {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single stockRelationshipCvterm record from the query.
func (q stockRelationshipCvtermQuery) One() (*StockRelationshipCvterm, error) {
	o := &StockRelationshipCvterm{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for stock_relationship_cvterm")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all StockRelationshipCvterm records from the query, and panics on error.
func (q stockRelationshipCvtermQuery) AllP() StockRelationshipCvtermSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all StockRelationshipCvterm records from the query.
func (q stockRelationshipCvtermQuery) All() (StockRelationshipCvtermSlice, error) {
	var o StockRelationshipCvtermSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to StockRelationshipCvterm slice")
	}

	if len(stockRelationshipCvtermAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all StockRelationshipCvterm records in the query, and panics on error.
func (q stockRelationshipCvtermQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all StockRelationshipCvterm records in the query.
func (q stockRelationshipCvtermQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count stock_relationship_cvterm rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q stockRelationshipCvtermQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q stockRelationshipCvtermQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if stock_relationship_cvterm exists")
	}

	return count > 0, nil
}

// CvtermG pointed to by the foreign key.
func (o *StockRelationshipCvterm) CvtermG(mods ...qm.QueryMod) cvtermQuery {
	return o.Cvterm(boil.GetDB(), mods...)
}

// Cvterm pointed to by the foreign key.
func (o *StockRelationshipCvterm) Cvterm(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// PubG pointed to by the foreign key.
func (o *StockRelationshipCvterm) PubG(mods ...qm.QueryMod) pubQuery {
	return o.Pub(boil.GetDB(), mods...)
}

// Pub pointed to by the foreign key.
func (o *StockRelationshipCvterm) Pub(exec boil.Executor, mods ...qm.QueryMod) pubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub\"")

	return query
}

// StockRelationshipG pointed to by the foreign key.
func (o *StockRelationshipCvterm) StockRelationshipG(mods ...qm.QueryMod) stockRelationshipQuery {
	return o.StockRelationship(boil.GetDB(), mods...)
}

// StockRelationship pointed to by the foreign key.
func (o *StockRelationshipCvterm) StockRelationship(exec boil.Executor, mods ...qm.QueryMod) stockRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_relationship_id=$1", o.StockRelationshipID),
	}

	queryMods = append(queryMods, mods...)

	query := StockRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_relationship\"")

	return query
}

// LoadCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockRelationshipCvtermL) LoadCvterm(e boil.Executor, singular bool, maybeStockRelationshipCvterm interface{}) error {
	var slice []*StockRelationshipCvterm
	var object *StockRelationshipCvterm

	count := 1
	if singular {
		object = maybeStockRelationshipCvterm.(*StockRelationshipCvterm)
	} else {
		slice = *maybeStockRelationshipCvterm.(*StockRelationshipCvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockRelationshipCvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &stockRelationshipCvtermR{}
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

	if len(stockRelationshipCvtermAfterSelectHooks) != 0 {
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
func (stockRelationshipCvtermL) LoadPub(e boil.Executor, singular bool, maybeStockRelationshipCvterm interface{}) error {
	var slice []*StockRelationshipCvterm
	var object *StockRelationshipCvterm

	count := 1
	if singular {
		object = maybeStockRelationshipCvterm.(*StockRelationshipCvterm)
	} else {
		slice = *maybeStockRelationshipCvterm.(*StockRelationshipCvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockRelationshipCvtermR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &stockRelationshipCvtermR{}
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

	if len(stockRelationshipCvtermAfterSelectHooks) != 0 {
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
			if local.PubID.Int == foreign.PubID {
				local.R.Pub = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockRelationshipCvtermL) LoadStockRelationship(e boil.Executor, singular bool, maybeStockRelationshipCvterm interface{}) error {
	var slice []*StockRelationshipCvterm
	var object *StockRelationshipCvterm

	count := 1
	if singular {
		object = maybeStockRelationshipCvterm.(*StockRelationshipCvterm)
	} else {
		slice = *maybeStockRelationshipCvterm.(*StockRelationshipCvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockRelationshipCvtermR{}
		args[0] = object.StockRelationshipID
	} else {
		for i, obj := range slice {
			obj.R = &stockRelationshipCvtermR{}
			args[i] = obj.StockRelationshipID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_relationship\" where \"stock_relationship_id\" in (%s)",
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

	if len(stockRelationshipCvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.StockRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockRelationshipID == foreign.StockRelationshipID {
				local.R.StockRelationship = foreign
				break
			}
		}
	}

	return nil
}

// SetCvterm of the stock_relationship_cvterm to the related item.
// Sets o.R.Cvterm to related.
// Adds o to related.R.StockRelationshipCvterms.
func (o *StockRelationshipCvterm) SetCvterm(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_relationship_cvterm\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"cvterm_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockRelationshipCvtermPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.StockRelationshipCvtermID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CvtermID = related.CvtermID

	if o.R == nil {
		o.R = &stockRelationshipCvtermR{
			Cvterm: related,
		}
	} else {
		o.R.Cvterm = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			StockRelationshipCvterms: StockRelationshipCvtermSlice{o},
		}
	} else {
		related.R.StockRelationshipCvterms = append(related.R.StockRelationshipCvterms, o)
	}

	return nil
}

// SetPub of the stock_relationship_cvterm to the related item.
// Sets o.R.Pub to related.
// Adds o to related.R.StockRelationshipCvterms.
func (o *StockRelationshipCvterm) SetPub(exec boil.Executor, insert bool, related *Pub) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_relationship_cvterm\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockRelationshipCvtermPrimaryKeyColumns),
	)
	values := []interface{}{related.PubID, o.StockRelationshipCvtermID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PubID.Int = related.PubID
	o.PubID.Valid = true

	if o.R == nil {
		o.R = &stockRelationshipCvtermR{
			Pub: related,
		}
	} else {
		o.R.Pub = related
	}

	if related.R == nil {
		related.R = &pubR{
			StockRelationshipCvterms: StockRelationshipCvtermSlice{o},
		}
	} else {
		related.R.StockRelationshipCvterms = append(related.R.StockRelationshipCvterms, o)
	}

	return nil
}

// RemovePub relationship.
// Sets o.R.Pub to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *StockRelationshipCvterm) RemovePub(exec boil.Executor, related *Pub) error {
	var err error

	o.PubID.Valid = false
	if err = o.Update(exec, "pub_id"); err != nil {
		o.PubID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Pub = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.StockRelationshipCvterms {
		if o.PubID.Int != ri.PubID.Int {
			continue
		}

		ln := len(related.R.StockRelationshipCvterms)
		if ln > 1 && i < ln-1 {
			related.R.StockRelationshipCvterms[i] = related.R.StockRelationshipCvterms[ln-1]
		}
		related.R.StockRelationshipCvterms = related.R.StockRelationshipCvterms[:ln-1]
		break
	}
	return nil
}

// SetStockRelationship of the stock_relationship_cvterm to the related item.
// Sets o.R.StockRelationship to related.
// Adds o to related.R.StockRelationshipCvterms.
func (o *StockRelationshipCvterm) SetStockRelationship(exec boil.Executor, insert bool, related *StockRelationship) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_relationship_cvterm\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"stock_relationship_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockRelationshipCvtermPrimaryKeyColumns),
	)
	values := []interface{}{related.StockRelationshipID, o.StockRelationshipCvtermID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StockRelationshipID = related.StockRelationshipID

	if o.R == nil {
		o.R = &stockRelationshipCvtermR{
			StockRelationship: related,
		}
	} else {
		o.R.StockRelationship = related
	}

	if related.R == nil {
		related.R = &stockRelationshipR{
			StockRelationshipCvterms: StockRelationshipCvtermSlice{o},
		}
	} else {
		related.R.StockRelationshipCvterms = append(related.R.StockRelationshipCvterms, o)
	}

	return nil
}

// StockRelationshipCvtermsG retrieves all records.
func StockRelationshipCvtermsG(mods ...qm.QueryMod) stockRelationshipCvtermQuery {
	return StockRelationshipCvterms(boil.GetDB(), mods...)
}

// StockRelationshipCvterms retrieves all the records using an executor.
func StockRelationshipCvterms(exec boil.Executor, mods ...qm.QueryMod) stockRelationshipCvtermQuery {
	mods = append(mods, qm.From("\"stock_relationship_cvterm\""))
	return stockRelationshipCvtermQuery{NewQuery(exec, mods...)}
}

// FindStockRelationshipCvtermG retrieves a single record by ID.
func FindStockRelationshipCvtermG(stockRelationshipCvtermID int, selectCols ...string) (*StockRelationshipCvterm, error) {
	return FindStockRelationshipCvterm(boil.GetDB(), stockRelationshipCvtermID, selectCols...)
}

// FindStockRelationshipCvtermGP retrieves a single record by ID, and panics on error.
func FindStockRelationshipCvtermGP(stockRelationshipCvtermID int, selectCols ...string) *StockRelationshipCvterm {
	retobj, err := FindStockRelationshipCvterm(boil.GetDB(), stockRelationshipCvtermID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindStockRelationshipCvterm retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStockRelationshipCvterm(exec boil.Executor, stockRelationshipCvtermID int, selectCols ...string) (*StockRelationshipCvterm, error) {
	stockRelationshipCvtermObj := &StockRelationshipCvterm{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stock_relationship_cvterm\" where \"stock_relationship_cvterm_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, stockRelationshipCvtermID)

	err := q.Bind(stockRelationshipCvtermObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from stock_relationship_cvterm")
	}

	return stockRelationshipCvtermObj, nil
}

// FindStockRelationshipCvtermP retrieves a single record by ID with an executor, and panics on error.
func FindStockRelationshipCvtermP(exec boil.Executor, stockRelationshipCvtermID int, selectCols ...string) *StockRelationshipCvterm {
	retobj, err := FindStockRelationshipCvterm(exec, stockRelationshipCvtermID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *StockRelationshipCvterm) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *StockRelationshipCvterm) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *StockRelationshipCvterm) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *StockRelationshipCvterm) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no stock_relationship_cvterm provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockRelationshipCvtermColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	stockRelationshipCvtermInsertCacheMut.RLock()
	cache, cached := stockRelationshipCvtermInsertCache[key]
	stockRelationshipCvtermInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			stockRelationshipCvtermColumns,
			stockRelationshipCvtermColumnsWithDefault,
			stockRelationshipCvtermColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(stockRelationshipCvtermType, stockRelationshipCvtermMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockRelationshipCvtermType, stockRelationshipCvtermMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"stock_relationship_cvterm\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into stock_relationship_cvterm")
	}

	if !cached {
		stockRelationshipCvtermInsertCacheMut.Lock()
		stockRelationshipCvtermInsertCache[key] = cache
		stockRelationshipCvtermInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single StockRelationshipCvterm record. See Update for
// whitelist behavior description.
func (o *StockRelationshipCvterm) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single StockRelationshipCvterm record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *StockRelationshipCvterm) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the StockRelationshipCvterm, and panics on error.
// See Update for whitelist behavior description.
func (o *StockRelationshipCvterm) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the StockRelationshipCvterm.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *StockRelationshipCvterm) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	stockRelationshipCvtermUpdateCacheMut.RLock()
	cache, cached := stockRelationshipCvtermUpdateCache[key]
	stockRelationshipCvtermUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(stockRelationshipCvtermColumns, stockRelationshipCvtermPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update stock_relationship_cvterm, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stock_relationship_cvterm\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stockRelationshipCvtermPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockRelationshipCvtermType, stockRelationshipCvtermMapping, append(wl, stockRelationshipCvtermPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update stock_relationship_cvterm row")
	}

	if !cached {
		stockRelationshipCvtermUpdateCacheMut.Lock()
		stockRelationshipCvtermUpdateCache[key] = cache
		stockRelationshipCvtermUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q stockRelationshipCvtermQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q stockRelationshipCvtermQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for stock_relationship_cvterm")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StockRelationshipCvtermSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o StockRelationshipCvtermSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o StockRelationshipCvtermSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockRelationshipCvtermSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockRelationshipCvtermPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"stock_relationship_cvterm\" SET %s WHERE (\"stock_relationship_cvterm_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockRelationshipCvtermPrimaryKeyColumns), len(colNames)+1, len(stockRelationshipCvtermPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in stockRelationshipCvterm slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *StockRelationshipCvterm) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *StockRelationshipCvterm) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *StockRelationshipCvterm) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *StockRelationshipCvterm) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no stock_relationship_cvterm provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockRelationshipCvtermColumnsWithDefault, o)

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

	stockRelationshipCvtermUpsertCacheMut.RLock()
	cache, cached := stockRelationshipCvtermUpsertCache[key]
	stockRelationshipCvtermUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			stockRelationshipCvtermColumns,
			stockRelationshipCvtermColumnsWithDefault,
			stockRelationshipCvtermColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			stockRelationshipCvtermColumns,
			stockRelationshipCvtermPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert stock_relationship_cvterm, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stockRelationshipCvtermPrimaryKeyColumns))
			copy(conflict, stockRelationshipCvtermPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"stock_relationship_cvterm\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(stockRelationshipCvtermType, stockRelationshipCvtermMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockRelationshipCvtermType, stockRelationshipCvtermMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for stock_relationship_cvterm")
	}

	if !cached {
		stockRelationshipCvtermUpsertCacheMut.Lock()
		stockRelationshipCvtermUpsertCache[key] = cache
		stockRelationshipCvtermUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single StockRelationshipCvterm record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockRelationshipCvterm) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single StockRelationshipCvterm record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *StockRelationshipCvterm) DeleteG() error {
	if o == nil {
		return errors.New("chado: no StockRelationshipCvterm provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single StockRelationshipCvterm record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockRelationshipCvterm) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single StockRelationshipCvterm record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *StockRelationshipCvterm) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no StockRelationshipCvterm provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockRelationshipCvtermPrimaryKeyMapping)
	sql := "DELETE FROM \"stock_relationship_cvterm\" WHERE \"stock_relationship_cvterm_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from stock_relationship_cvterm")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q stockRelationshipCvtermQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q stockRelationshipCvtermQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no stockRelationshipCvtermQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from stock_relationship_cvterm")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o StockRelationshipCvtermSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o StockRelationshipCvtermSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no StockRelationshipCvterm slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o StockRelationshipCvtermSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockRelationshipCvtermSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no StockRelationshipCvterm slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(stockRelationshipCvtermBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockRelationshipCvtermPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"stock_relationship_cvterm\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockRelationshipCvtermPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockRelationshipCvtermPrimaryKeyColumns), 1, len(stockRelationshipCvtermPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from stockRelationshipCvterm slice")
	}

	if len(stockRelationshipCvtermAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *StockRelationshipCvterm) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *StockRelationshipCvterm) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *StockRelationshipCvterm) ReloadG() error {
	if o == nil {
		return errors.New("chado: no StockRelationshipCvterm provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *StockRelationshipCvterm) Reload(exec boil.Executor) error {
	ret, err := FindStockRelationshipCvterm(exec, o.StockRelationshipCvtermID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockRelationshipCvtermSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockRelationshipCvtermSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockRelationshipCvtermSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty StockRelationshipCvtermSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockRelationshipCvtermSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	stockRelationshipCvterms := StockRelationshipCvtermSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockRelationshipCvtermPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"stock_relationship_cvterm\".* FROM \"stock_relationship_cvterm\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockRelationshipCvtermPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(stockRelationshipCvtermPrimaryKeyColumns), 1, len(stockRelationshipCvtermPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&stockRelationshipCvterms)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in StockRelationshipCvtermSlice")
	}

	*o = stockRelationshipCvterms

	return nil
}

// StockRelationshipCvtermExists checks if the StockRelationshipCvterm row exists.
func StockRelationshipCvtermExists(exec boil.Executor, stockRelationshipCvtermID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"stock_relationship_cvterm\" where \"stock_relationship_cvterm_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, stockRelationshipCvtermID)
	}

	row := exec.QueryRow(sql, stockRelationshipCvtermID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if stock_relationship_cvterm exists")
	}

	return exists, nil
}

// StockRelationshipCvtermExistsG checks if the StockRelationshipCvterm row exists.
func StockRelationshipCvtermExistsG(stockRelationshipCvtermID int) (bool, error) {
	return StockRelationshipCvtermExists(boil.GetDB(), stockRelationshipCvtermID)
}

// StockRelationshipCvtermExistsGP checks if the StockRelationshipCvterm row exists. Panics on error.
func StockRelationshipCvtermExistsGP(stockRelationshipCvtermID int) bool {
	e, err := StockRelationshipCvtermExists(boil.GetDB(), stockRelationshipCvtermID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// StockRelationshipCvtermExistsP checks if the StockRelationshipCvterm row exists. Panics on error.
func StockRelationshipCvtermExistsP(exec boil.Executor, stockRelationshipCvtermID int) bool {
	e, err := StockRelationshipCvtermExists(exec, stockRelationshipCvtermID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
