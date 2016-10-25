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

// Dbxref is an object representing the database table.
type Dbxref struct {
	DbxrefID    int         `boil:"dbxref_id" json:"dbxref_id" toml:"dbxref_id" yaml:"dbxref_id"`
	DBID        int         `boil:"db_id" json:"db_id" toml:"db_id" yaml:"db_id"`
	Accession   string      `boil:"accession" json:"accession" toml:"accession" yaml:"accession"`
	Version     string      `boil:"version" json:"version" toml:"version" yaml:"version"`
	Description null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`

	R *dbxrefR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dbxrefL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// dbxrefR is where relationships are stored.
type dbxrefR struct {
	DB                  *DB
	PubDbxref           *PubDbxref
	Cvterm              *Cvterm
	FeatureDbxref       *FeatureDbxref
	Dbxrefprop          *Dbxrefprop
	FeatureCvtermDbxref *FeatureCvtermDbxref
	CvtermDbxref        *CvtermDbxref
	StockDbxref         *StockDbxref
	OrganismDbxref      *OrganismDbxref
	Stocks              StockSlice
	Features            FeatureSlice
}

// dbxrefL is where Load methods for each relationship are stored.
type dbxrefL struct{}

var (
	dbxrefColumns               = []string{"dbxref_id", "db_id", "accession", "version", "description"}
	dbxrefColumnsWithoutDefault = []string{"db_id", "accession", "description"}
	dbxrefColumnsWithDefault    = []string{"dbxref_id", "version"}
	dbxrefPrimaryKeyColumns     = []string{"dbxref_id"}
)

type (
	// DbxrefSlice is an alias for a slice of pointers to Dbxref.
	// This should generally be used opposed to []Dbxref.
	DbxrefSlice []*Dbxref
	// DbxrefHook is the signature for custom Dbxref hook methods
	DbxrefHook func(boil.Executor, *Dbxref) error

	dbxrefQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dbxrefType                 = reflect.TypeOf(&Dbxref{})
	dbxrefMapping              = queries.MakeStructMapping(dbxrefType)
	dbxrefPrimaryKeyMapping, _ = queries.BindMapping(dbxrefType, dbxrefMapping, dbxrefPrimaryKeyColumns)
	dbxrefInsertCacheMut       sync.RWMutex
	dbxrefInsertCache          = make(map[string]insertCache)
	dbxrefUpdateCacheMut       sync.RWMutex
	dbxrefUpdateCache          = make(map[string]updateCache)
	dbxrefUpsertCacheMut       sync.RWMutex
	dbxrefUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var dbxrefBeforeInsertHooks []DbxrefHook
var dbxrefBeforeUpdateHooks []DbxrefHook
var dbxrefBeforeDeleteHooks []DbxrefHook
var dbxrefBeforeUpsertHooks []DbxrefHook

var dbxrefAfterInsertHooks []DbxrefHook
var dbxrefAfterSelectHooks []DbxrefHook
var dbxrefAfterUpdateHooks []DbxrefHook
var dbxrefAfterDeleteHooks []DbxrefHook
var dbxrefAfterUpsertHooks []DbxrefHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Dbxref) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range dbxrefBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Dbxref) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range dbxrefBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Dbxref) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range dbxrefBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Dbxref) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range dbxrefBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Dbxref) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range dbxrefAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Dbxref) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range dbxrefAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Dbxref) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range dbxrefAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Dbxref) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range dbxrefAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Dbxref) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range dbxrefAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDbxrefHook registers your hook function for all future operations.
func AddDbxrefHook(hookPoint boil.HookPoint, dbxrefHook DbxrefHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		dbxrefBeforeInsertHooks = append(dbxrefBeforeInsertHooks, dbxrefHook)
	case boil.BeforeUpdateHook:
		dbxrefBeforeUpdateHooks = append(dbxrefBeforeUpdateHooks, dbxrefHook)
	case boil.BeforeDeleteHook:
		dbxrefBeforeDeleteHooks = append(dbxrefBeforeDeleteHooks, dbxrefHook)
	case boil.BeforeUpsertHook:
		dbxrefBeforeUpsertHooks = append(dbxrefBeforeUpsertHooks, dbxrefHook)
	case boil.AfterInsertHook:
		dbxrefAfterInsertHooks = append(dbxrefAfterInsertHooks, dbxrefHook)
	case boil.AfterSelectHook:
		dbxrefAfterSelectHooks = append(dbxrefAfterSelectHooks, dbxrefHook)
	case boil.AfterUpdateHook:
		dbxrefAfterUpdateHooks = append(dbxrefAfterUpdateHooks, dbxrefHook)
	case boil.AfterDeleteHook:
		dbxrefAfterDeleteHooks = append(dbxrefAfterDeleteHooks, dbxrefHook)
	case boil.AfterUpsertHook:
		dbxrefAfterUpsertHooks = append(dbxrefAfterUpsertHooks, dbxrefHook)
	}
}

// OneP returns a single dbxref record from the query, and panics on error.
func (q dbxrefQuery) OneP() *Dbxref {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single dbxref record from the query.
func (q dbxrefQuery) One() (*Dbxref, error) {
	o := &Dbxref{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for dbxref")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Dbxref records from the query, and panics on error.
func (q dbxrefQuery) AllP() DbxrefSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Dbxref records from the query.
func (q dbxrefQuery) All() (DbxrefSlice, error) {
	var o DbxrefSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Dbxref slice")
	}

	if len(dbxrefAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Dbxref records in the query, and panics on error.
func (q dbxrefQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Dbxref records in the query.
func (q dbxrefQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count dbxref rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q dbxrefQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q dbxrefQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if dbxref exists")
	}

	return count > 0, nil
}

// DBG pointed to by the foreign key.
func (o *Dbxref) DBG(mods ...qm.QueryMod) dbQuery {
	return o.DB(boil.GetDB(), mods...)
}

// DB pointed to by the foreign key.
func (o *Dbxref) DB(exec boil.Executor, mods ...qm.QueryMod) dbQuery {
	queryMods := []qm.QueryMod{
		qm.Where("db_id=$1", o.DBID),
	}

	queryMods = append(queryMods, mods...)

	query := DBS(exec, queryMods...)
	queries.SetFrom(query.Query, "\"db\"")

	return query
}

// PubDbxrefG pointed to by the foreign key.
func (o *Dbxref) PubDbxrefG(mods ...qm.QueryMod) pubDbxrefQuery {
	return o.PubDbxref(boil.GetDB(), mods...)
}

// PubDbxref pointed to by the foreign key.
func (o *Dbxref) PubDbxref(exec boil.Executor, mods ...qm.QueryMod) pubDbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("dbxref_id=$1", o.DbxrefID),
	}

	queryMods = append(queryMods, mods...)

	query := PubDbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub_dbxref\"")

	return query
}

// CvtermG pointed to by the foreign key.
func (o *Dbxref) CvtermG(mods ...qm.QueryMod) cvtermQuery {
	return o.Cvterm(boil.GetDB(), mods...)
}

// Cvterm pointed to by the foreign key.
func (o *Dbxref) Cvterm(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("dbxref_id=$1", o.DbxrefID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// FeatureDbxrefG pointed to by the foreign key.
func (o *Dbxref) FeatureDbxrefG(mods ...qm.QueryMod) featureDbxrefQuery {
	return o.FeatureDbxref(boil.GetDB(), mods...)
}

// FeatureDbxref pointed to by the foreign key.
func (o *Dbxref) FeatureDbxref(exec boil.Executor, mods ...qm.QueryMod) featureDbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("dbxref_id=$1", o.DbxrefID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureDbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_dbxref\"")

	return query
}

// DbxrefpropG pointed to by the foreign key.
func (o *Dbxref) DbxrefpropG(mods ...qm.QueryMod) dbxrefpropQuery {
	return o.Dbxrefprop(boil.GetDB(), mods...)
}

// Dbxrefprop pointed to by the foreign key.
func (o *Dbxref) Dbxrefprop(exec boil.Executor, mods ...qm.QueryMod) dbxrefpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("dbxref_id=$1", o.DbxrefID),
	}

	queryMods = append(queryMods, mods...)

	query := Dbxrefprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"dbxrefprop\"")

	return query
}

// FeatureCvtermDbxrefG pointed to by the foreign key.
func (o *Dbxref) FeatureCvtermDbxrefG(mods ...qm.QueryMod) featureCvtermDbxrefQuery {
	return o.FeatureCvtermDbxref(boil.GetDB(), mods...)
}

// FeatureCvtermDbxref pointed to by the foreign key.
func (o *Dbxref) FeatureCvtermDbxref(exec boil.Executor, mods ...qm.QueryMod) featureCvtermDbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("dbxref_id=$1", o.DbxrefID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureCvtermDbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_cvterm_dbxref\"")

	return query
}

// CvtermDbxrefG pointed to by the foreign key.
func (o *Dbxref) CvtermDbxrefG(mods ...qm.QueryMod) cvtermDbxrefQuery {
	return o.CvtermDbxref(boil.GetDB(), mods...)
}

// CvtermDbxref pointed to by the foreign key.
func (o *Dbxref) CvtermDbxref(exec boil.Executor, mods ...qm.QueryMod) cvtermDbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("dbxref_id=$1", o.DbxrefID),
	}

	queryMods = append(queryMods, mods...)

	query := CvtermDbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm_dbxref\"")

	return query
}

// StockDbxrefG pointed to by the foreign key.
func (o *Dbxref) StockDbxrefG(mods ...qm.QueryMod) stockDbxrefQuery {
	return o.StockDbxref(boil.GetDB(), mods...)
}

// StockDbxref pointed to by the foreign key.
func (o *Dbxref) StockDbxref(exec boil.Executor, mods ...qm.QueryMod) stockDbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("dbxref_id=$1", o.DbxrefID),
	}

	queryMods = append(queryMods, mods...)

	query := StockDbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_dbxref\"")

	return query
}

// OrganismDbxrefG pointed to by the foreign key.
func (o *Dbxref) OrganismDbxrefG(mods ...qm.QueryMod) organismDbxrefQuery {
	return o.OrganismDbxref(boil.GetDB(), mods...)
}

// OrganismDbxref pointed to by the foreign key.
func (o *Dbxref) OrganismDbxref(exec boil.Executor, mods ...qm.QueryMod) organismDbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("dbxref_id=$1", o.DbxrefID),
	}

	queryMods = append(queryMods, mods...)

	query := OrganismDbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"organism_dbxref\"")

	return query
}

// StocksG retrieves all the stock's stock.
func (o *Dbxref) StocksG(mods ...qm.QueryMod) stockQuery {
	return o.Stocks(boil.GetDB(), mods...)
}

// Stocks retrieves all the stock's stock with an executor.
func (o *Dbxref) Stocks(exec boil.Executor, mods ...qm.QueryMod) stockQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"dbxref_id\"=$1", o.DbxrefID),
	)

	query := Stocks(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock\" as \"a\"")
	return query
}

// FeaturesG retrieves all the feature's feature.
func (o *Dbxref) FeaturesG(mods ...qm.QueryMod) featureQuery {
	return o.Features(boil.GetDB(), mods...)
}

// Features retrieves all the feature's feature with an executor.
func (o *Dbxref) Features(exec boil.Executor, mods ...qm.QueryMod) featureQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"dbxref_id\"=$1", o.DbxrefID),
	)

	query := Features(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature\" as \"a\"")
	return query
}

// LoadDB allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dbxrefL) LoadDB(e boil.Executor, singular bool, maybeDbxref interface{}) error {
	var slice []*Dbxref
	var object *Dbxref

	count := 1
	if singular {
		object = maybeDbxref.(*Dbxref)
	} else {
		slice = *maybeDbxref.(*DbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &dbxrefR{}
		args[0] = object.DBID
	} else {
		for i, obj := range slice {
			obj.R = &dbxrefR{}
			args[i] = obj.DBID
		}
	}

	query := fmt.Sprintf(
		"select * from \"db\" where \"db_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DB")
	}
	defer results.Close()

	var resultSlice []*DB
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DB")
	}

	if len(dbxrefAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.DB = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.DBID == foreign.DBID {
				local.R.DB = foreign
				break
			}
		}
	}

	return nil
}

// LoadPubDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dbxrefL) LoadPubDbxref(e boil.Executor, singular bool, maybeDbxref interface{}) error {
	var slice []*Dbxref
	var object *Dbxref

	count := 1
	if singular {
		object = maybeDbxref.(*Dbxref)
	} else {
		slice = *maybeDbxref.(*DbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &dbxrefR{}
		args[0] = object.DbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &dbxrefR{}
			args[i] = obj.DbxrefID
		}
	}

	query := fmt.Sprintf(
		"select * from \"pub_dbxref\" where \"dbxref_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load PubDbxref")
	}
	defer results.Close()

	var resultSlice []*PubDbxref
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice PubDbxref")
	}

	if len(dbxrefAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.PubDbxref = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.DbxrefID == foreign.DbxrefID {
				local.R.PubDbxref = foreign
				break
			}
		}
	}

	return nil
}

// LoadCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dbxrefL) LoadCvterm(e boil.Executor, singular bool, maybeDbxref interface{}) error {
	var slice []*Dbxref
	var object *Dbxref

	count := 1
	if singular {
		object = maybeDbxref.(*Dbxref)
	} else {
		slice = *maybeDbxref.(*DbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &dbxrefR{}
		args[0] = object.DbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &dbxrefR{}
			args[i] = obj.DbxrefID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvterm\" where \"dbxref_id\" in (%s)",
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

	if len(dbxrefAfterSelectHooks) != 0 {
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
			if local.DbxrefID == foreign.DbxrefID {
				local.R.Cvterm = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeatureDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dbxrefL) LoadFeatureDbxref(e boil.Executor, singular bool, maybeDbxref interface{}) error {
	var slice []*Dbxref
	var object *Dbxref

	count := 1
	if singular {
		object = maybeDbxref.(*Dbxref)
	} else {
		slice = *maybeDbxref.(*DbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &dbxrefR{}
		args[0] = object.DbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &dbxrefR{}
			args[i] = obj.DbxrefID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_dbxref\" where \"dbxref_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureDbxref")
	}
	defer results.Close()

	var resultSlice []*FeatureDbxref
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureDbxref")
	}

	if len(dbxrefAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeatureDbxref = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.DbxrefID == foreign.DbxrefID {
				local.R.FeatureDbxref = foreign
				break
			}
		}
	}

	return nil
}

// LoadDbxrefprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dbxrefL) LoadDbxrefprop(e boil.Executor, singular bool, maybeDbxref interface{}) error {
	var slice []*Dbxref
	var object *Dbxref

	count := 1
	if singular {
		object = maybeDbxref.(*Dbxref)
	} else {
		slice = *maybeDbxref.(*DbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &dbxrefR{}
		args[0] = object.DbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &dbxrefR{}
			args[i] = obj.DbxrefID
		}
	}

	query := fmt.Sprintf(
		"select * from \"dbxrefprop\" where \"dbxref_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Dbxrefprop")
	}
	defer results.Close()

	var resultSlice []*Dbxrefprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Dbxrefprop")
	}

	if len(dbxrefAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Dbxrefprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.DbxrefID == foreign.DbxrefID {
				local.R.Dbxrefprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeatureCvtermDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dbxrefL) LoadFeatureCvtermDbxref(e boil.Executor, singular bool, maybeDbxref interface{}) error {
	var slice []*Dbxref
	var object *Dbxref

	count := 1
	if singular {
		object = maybeDbxref.(*Dbxref)
	} else {
		slice = *maybeDbxref.(*DbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &dbxrefR{}
		args[0] = object.DbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &dbxrefR{}
			args[i] = obj.DbxrefID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_cvterm_dbxref\" where \"dbxref_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureCvtermDbxref")
	}
	defer results.Close()

	var resultSlice []*FeatureCvtermDbxref
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureCvtermDbxref")
	}

	if len(dbxrefAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeatureCvtermDbxref = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.DbxrefID == foreign.DbxrefID {
				local.R.FeatureCvtermDbxref = foreign
				break
			}
		}
	}

	return nil
}

// LoadCvtermDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dbxrefL) LoadCvtermDbxref(e boil.Executor, singular bool, maybeDbxref interface{}) error {
	var slice []*Dbxref
	var object *Dbxref

	count := 1
	if singular {
		object = maybeDbxref.(*Dbxref)
	} else {
		slice = *maybeDbxref.(*DbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &dbxrefR{}
		args[0] = object.DbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &dbxrefR{}
			args[i] = obj.DbxrefID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvterm_dbxref\" where \"dbxref_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load CvtermDbxref")
	}
	defer results.Close()

	var resultSlice []*CvtermDbxref
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice CvtermDbxref")
	}

	if len(dbxrefAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.CvtermDbxref = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.DbxrefID == foreign.DbxrefID {
				local.R.CvtermDbxref = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dbxrefL) LoadStockDbxref(e boil.Executor, singular bool, maybeDbxref interface{}) error {
	var slice []*Dbxref
	var object *Dbxref

	count := 1
	if singular {
		object = maybeDbxref.(*Dbxref)
	} else {
		slice = *maybeDbxref.(*DbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &dbxrefR{}
		args[0] = object.DbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &dbxrefR{}
			args[i] = obj.DbxrefID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_dbxref\" where \"dbxref_id\" in (%s)",
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

	if len(dbxrefAfterSelectHooks) != 0 {
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
			if local.DbxrefID == foreign.DbxrefID {
				local.R.StockDbxref = foreign
				break
			}
		}
	}

	return nil
}

// LoadOrganismDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dbxrefL) LoadOrganismDbxref(e boil.Executor, singular bool, maybeDbxref interface{}) error {
	var slice []*Dbxref
	var object *Dbxref

	count := 1
	if singular {
		object = maybeDbxref.(*Dbxref)
	} else {
		slice = *maybeDbxref.(*DbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &dbxrefR{}
		args[0] = object.DbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &dbxrefR{}
			args[i] = obj.DbxrefID
		}
	}

	query := fmt.Sprintf(
		"select * from \"organism_dbxref\" where \"dbxref_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load OrganismDbxref")
	}
	defer results.Close()

	var resultSlice []*OrganismDbxref
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice OrganismDbxref")
	}

	if len(dbxrefAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.OrganismDbxref = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.DbxrefID == foreign.DbxrefID {
				local.R.OrganismDbxref = foreign
				break
			}
		}
	}

	return nil
}

// LoadStocks allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dbxrefL) LoadStocks(e boil.Executor, singular bool, maybeDbxref interface{}) error {
	var slice []*Dbxref
	var object *Dbxref

	count := 1
	if singular {
		object = maybeDbxref.(*Dbxref)
	} else {
		slice = *maybeDbxref.(*DbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &dbxrefR{}
		args[0] = object.DbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &dbxrefR{}
			args[i] = obj.DbxrefID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock\" where \"dbxref_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load stock")
	}
	defer results.Close()

	var resultSlice []*Stock
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice stock")
	}

	if len(stockAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Stocks = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.DbxrefID == foreign.DbxrefID.Int {
				local.R.Stocks = append(local.R.Stocks, foreign)
				break
			}
		}
	}

	return nil
}

// LoadFeatures allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dbxrefL) LoadFeatures(e boil.Executor, singular bool, maybeDbxref interface{}) error {
	var slice []*Dbxref
	var object *Dbxref

	count := 1
	if singular {
		object = maybeDbxref.(*Dbxref)
	} else {
		slice = *maybeDbxref.(*DbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &dbxrefR{}
		args[0] = object.DbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &dbxrefR{}
			args[i] = obj.DbxrefID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature\" where \"dbxref_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load feature")
	}
	defer results.Close()

	var resultSlice []*Feature
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice feature")
	}

	if len(featureAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Features = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.DbxrefID == foreign.DbxrefID.Int {
				local.R.Features = append(local.R.Features, foreign)
				break
			}
		}
	}

	return nil
}

// SetDB of the dbxref to the related item.
// Sets o.R.DB to related.
// Adds o to related.R.Dbxref.
func (o *Dbxref) SetDB(exec boil.Executor, insert bool, related *DB) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"dbxref\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"db_id"}),
		strmangle.WhereClause("\"", "\"", 2, dbxrefPrimaryKeyColumns),
	)
	values := []interface{}{related.DBID, o.DbxrefID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.DBID = related.DBID

	if o.R == nil {
		o.R = &dbxrefR{
			DB: related,
		}
	} else {
		o.R.DB = related
	}

	if related.R == nil {
		related.R = &dbR{
			Dbxref: o,
		}
	} else {
		related.R.Dbxref = o
	}

	return nil
}

// SetPubDbxref of the dbxref to the related item.
// Sets o.R.PubDbxref to related.
// Adds o to related.R.Dbxref.
func (o *Dbxref) SetPubDbxref(exec boil.Executor, insert bool, related *PubDbxref) error {
	var err error

	if insert {
		related.DbxrefID = o.DbxrefID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"pub_dbxref\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"dbxref_id"}),
			strmangle.WhereClause("\"", "\"", 2, pubDbxrefPrimaryKeyColumns),
		)
		values := []interface{}{o.DbxrefID, related.PubDbxrefID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.DbxrefID = o.DbxrefID

	}

	if o.R == nil {
		o.R = &dbxrefR{
			PubDbxref: related,
		}
	} else {
		o.R.PubDbxref = related
	}

	if related.R == nil {
		related.R = &pubDbxrefR{
			Dbxref: o,
		}
	} else {
		related.R.Dbxref = o
	}
	return nil
}

// SetCvterm of the dbxref to the related item.
// Sets o.R.Cvterm to related.
// Adds o to related.R.Dbxref.
func (o *Dbxref) SetCvterm(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error

	if insert {
		related.DbxrefID = o.DbxrefID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"cvterm\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"dbxref_id"}),
			strmangle.WhereClause("\"", "\"", 2, cvtermPrimaryKeyColumns),
		)
		values := []interface{}{o.DbxrefID, related.CvtermID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.DbxrefID = o.DbxrefID

	}

	if o.R == nil {
		o.R = &dbxrefR{
			Cvterm: related,
		}
	} else {
		o.R.Cvterm = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			Dbxref: o,
		}
	} else {
		related.R.Dbxref = o
	}
	return nil
}

// SetFeatureDbxref of the dbxref to the related item.
// Sets o.R.FeatureDbxref to related.
// Adds o to related.R.Dbxref.
func (o *Dbxref) SetFeatureDbxref(exec boil.Executor, insert bool, related *FeatureDbxref) error {
	var err error

	if insert {
		related.DbxrefID = o.DbxrefID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_dbxref\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"dbxref_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureDbxrefPrimaryKeyColumns),
		)
		values := []interface{}{o.DbxrefID, related.FeatureDbxrefID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.DbxrefID = o.DbxrefID

	}

	if o.R == nil {
		o.R = &dbxrefR{
			FeatureDbxref: related,
		}
	} else {
		o.R.FeatureDbxref = related
	}

	if related.R == nil {
		related.R = &featureDbxrefR{
			Dbxref: o,
		}
	} else {
		related.R.Dbxref = o
	}
	return nil
}

// SetDbxrefprop of the dbxref to the related item.
// Sets o.R.Dbxrefprop to related.
// Adds o to related.R.Dbxref.
func (o *Dbxref) SetDbxrefprop(exec boil.Executor, insert bool, related *Dbxrefprop) error {
	var err error

	if insert {
		related.DbxrefID = o.DbxrefID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"dbxrefprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"dbxref_id"}),
			strmangle.WhereClause("\"", "\"", 2, dbxrefpropPrimaryKeyColumns),
		)
		values := []interface{}{o.DbxrefID, related.DbxrefpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.DbxrefID = o.DbxrefID

	}

	if o.R == nil {
		o.R = &dbxrefR{
			Dbxrefprop: related,
		}
	} else {
		o.R.Dbxrefprop = related
	}

	if related.R == nil {
		related.R = &dbxrefpropR{
			Dbxref: o,
		}
	} else {
		related.R.Dbxref = o
	}
	return nil
}

// SetFeatureCvtermDbxref of the dbxref to the related item.
// Sets o.R.FeatureCvtermDbxref to related.
// Adds o to related.R.Dbxref.
func (o *Dbxref) SetFeatureCvtermDbxref(exec boil.Executor, insert bool, related *FeatureCvtermDbxref) error {
	var err error

	if insert {
		related.DbxrefID = o.DbxrefID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_cvterm_dbxref\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"dbxref_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureCvtermDbxrefPrimaryKeyColumns),
		)
		values := []interface{}{o.DbxrefID, related.FeatureCvtermDbxrefID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.DbxrefID = o.DbxrefID

	}

	if o.R == nil {
		o.R = &dbxrefR{
			FeatureCvtermDbxref: related,
		}
	} else {
		o.R.FeatureCvtermDbxref = related
	}

	if related.R == nil {
		related.R = &featureCvtermDbxrefR{
			Dbxref: o,
		}
	} else {
		related.R.Dbxref = o
	}
	return nil
}

// SetCvtermDbxref of the dbxref to the related item.
// Sets o.R.CvtermDbxref to related.
// Adds o to related.R.Dbxref.
func (o *Dbxref) SetCvtermDbxref(exec boil.Executor, insert bool, related *CvtermDbxref) error {
	var err error

	if insert {
		related.DbxrefID = o.DbxrefID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"cvterm_dbxref\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"dbxref_id"}),
			strmangle.WhereClause("\"", "\"", 2, cvtermDbxrefPrimaryKeyColumns),
		)
		values := []interface{}{o.DbxrefID, related.CvtermDbxrefID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.DbxrefID = o.DbxrefID

	}

	if o.R == nil {
		o.R = &dbxrefR{
			CvtermDbxref: related,
		}
	} else {
		o.R.CvtermDbxref = related
	}

	if related.R == nil {
		related.R = &cvtermDbxrefR{
			Dbxref: o,
		}
	} else {
		related.R.Dbxref = o
	}
	return nil
}

// SetStockDbxref of the dbxref to the related item.
// Sets o.R.StockDbxref to related.
// Adds o to related.R.Dbxref.
func (o *Dbxref) SetStockDbxref(exec boil.Executor, insert bool, related *StockDbxref) error {
	var err error

	if insert {
		related.DbxrefID = o.DbxrefID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock_dbxref\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"dbxref_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockDbxrefPrimaryKeyColumns),
		)
		values := []interface{}{o.DbxrefID, related.StockDbxrefID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.DbxrefID = o.DbxrefID

	}

	if o.R == nil {
		o.R = &dbxrefR{
			StockDbxref: related,
		}
	} else {
		o.R.StockDbxref = related
	}

	if related.R == nil {
		related.R = &stockDbxrefR{
			Dbxref: o,
		}
	} else {
		related.R.Dbxref = o
	}
	return nil
}

// SetOrganismDbxref of the dbxref to the related item.
// Sets o.R.OrganismDbxref to related.
// Adds o to related.R.Dbxref.
func (o *Dbxref) SetOrganismDbxref(exec boil.Executor, insert bool, related *OrganismDbxref) error {
	var err error

	if insert {
		related.DbxrefID = o.DbxrefID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"organism_dbxref\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"dbxref_id"}),
			strmangle.WhereClause("\"", "\"", 2, organismDbxrefPrimaryKeyColumns),
		)
		values := []interface{}{o.DbxrefID, related.OrganismDbxrefID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.DbxrefID = o.DbxrefID

	}

	if o.R == nil {
		o.R = &dbxrefR{
			OrganismDbxref: related,
		}
	} else {
		o.R.OrganismDbxref = related
	}

	if related.R == nil {
		related.R = &organismDbxrefR{
			Dbxref: o,
		}
	} else {
		related.R.Dbxref = o
	}
	return nil
}

// AddStocks adds the given related objects to the existing relationships
// of the dbxref, optionally inserting them as new records.
// Appends related to o.R.Stocks.
// Sets related.R.Dbxref appropriately.
func (o *Dbxref) AddStocks(exec boil.Executor, insert bool, related ...*Stock) error {
	var err error
	for _, rel := range related {
		rel.DbxrefID.Int = o.DbxrefID
		rel.DbxrefID.Valid = true
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "dbxref_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &dbxrefR{
			Stocks: related,
		}
	} else {
		o.R.Stocks = append(o.R.Stocks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &stockR{
				Dbxref: o,
			}
		} else {
			rel.R.Dbxref = o
		}
	}
	return nil
}

// SetStocks removes all previously related items of the
// dbxref replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Dbxref's Stocks accordingly.
// Replaces o.R.Stocks with related.
// Sets related.R.Dbxref's Stocks accordingly.
func (o *Dbxref) SetStocks(exec boil.Executor, insert bool, related ...*Stock) error {
	query := "update \"stock\" set \"dbxref_id\" = null where \"dbxref_id\" = $1"
	values := []interface{}{o.DbxrefID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.Stocks {
			rel.DbxrefID.Valid = false
			if rel.R == nil {
				continue
			}

			rel.R.Dbxref = nil
		}

		o.R.Stocks = nil
	}
	return o.AddStocks(exec, insert, related...)
}

// RemoveStocks relationships from objects passed in.
// Removes related items from R.Stocks (uses pointer comparison, removal does not keep order)
// Sets related.R.Dbxref.
func (o *Dbxref) RemoveStocks(exec boil.Executor, related ...*Stock) error {
	var err error
	for _, rel := range related {
		rel.DbxrefID.Valid = false
		if rel.R != nil {
			rel.R.Dbxref = nil
		}
		if err = rel.Update(exec, "dbxref_id"); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Stocks {
			if rel != ri {
				continue
			}

			ln := len(o.R.Stocks)
			if ln > 1 && i < ln-1 {
				o.R.Stocks[i] = o.R.Stocks[ln-1]
			}
			o.R.Stocks = o.R.Stocks[:ln-1]
			break
		}
	}

	return nil
}

// AddFeatures adds the given related objects to the existing relationships
// of the dbxref, optionally inserting them as new records.
// Appends related to o.R.Features.
// Sets related.R.Dbxref appropriately.
func (o *Dbxref) AddFeatures(exec boil.Executor, insert bool, related ...*Feature) error {
	var err error
	for _, rel := range related {
		rel.DbxrefID.Int = o.DbxrefID
		rel.DbxrefID.Valid = true
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "dbxref_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &dbxrefR{
			Features: related,
		}
	} else {
		o.R.Features = append(o.R.Features, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &featureR{
				Dbxref: o,
			}
		} else {
			rel.R.Dbxref = o
		}
	}
	return nil
}

// SetFeatures removes all previously related items of the
// dbxref replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Dbxref's Features accordingly.
// Replaces o.R.Features with related.
// Sets related.R.Dbxref's Features accordingly.
func (o *Dbxref) SetFeatures(exec boil.Executor, insert bool, related ...*Feature) error {
	query := "update \"feature\" set \"dbxref_id\" = null where \"dbxref_id\" = $1"
	values := []interface{}{o.DbxrefID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.Features {
			rel.DbxrefID.Valid = false
			if rel.R == nil {
				continue
			}

			rel.R.Dbxref = nil
		}

		o.R.Features = nil
	}
	return o.AddFeatures(exec, insert, related...)
}

// RemoveFeatures relationships from objects passed in.
// Removes related items from R.Features (uses pointer comparison, removal does not keep order)
// Sets related.R.Dbxref.
func (o *Dbxref) RemoveFeatures(exec boil.Executor, related ...*Feature) error {
	var err error
	for _, rel := range related {
		rel.DbxrefID.Valid = false
		if rel.R != nil {
			rel.R.Dbxref = nil
		}
		if err = rel.Update(exec, "dbxref_id"); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Features {
			if rel != ri {
				continue
			}

			ln := len(o.R.Features)
			if ln > 1 && i < ln-1 {
				o.R.Features[i] = o.R.Features[ln-1]
			}
			o.R.Features = o.R.Features[:ln-1]
			break
		}
	}

	return nil
}

// DbxrefsG retrieves all records.
func DbxrefsG(mods ...qm.QueryMod) dbxrefQuery {
	return Dbxrefs(boil.GetDB(), mods...)
}

// Dbxrefs retrieves all the records using an executor.
func Dbxrefs(exec boil.Executor, mods ...qm.QueryMod) dbxrefQuery {
	mods = append(mods, qm.From("\"dbxref\""))
	return dbxrefQuery{NewQuery(exec, mods...)}
}

// FindDbxrefG retrieves a single record by ID.
func FindDbxrefG(dbxrefID int, selectCols ...string) (*Dbxref, error) {
	return FindDbxref(boil.GetDB(), dbxrefID, selectCols...)
}

// FindDbxrefGP retrieves a single record by ID, and panics on error.
func FindDbxrefGP(dbxrefID int, selectCols ...string) *Dbxref {
	retobj, err := FindDbxref(boil.GetDB(), dbxrefID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDbxref retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDbxref(exec boil.Executor, dbxrefID int, selectCols ...string) (*Dbxref, error) {
	dbxrefObj := &Dbxref{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"dbxref\" where \"dbxref_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, dbxrefID)

	err := q.Bind(dbxrefObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from dbxref")
	}

	return dbxrefObj, nil
}

// FindDbxrefP retrieves a single record by ID with an executor, and panics on error.
func FindDbxrefP(exec boil.Executor, dbxrefID int, selectCols ...string) *Dbxref {
	retobj, err := FindDbxref(exec, dbxrefID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Dbxref) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Dbxref) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Dbxref) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Dbxref) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no dbxref provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dbxrefColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	dbxrefInsertCacheMut.RLock()
	cache, cached := dbxrefInsertCache[key]
	dbxrefInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			dbxrefColumns,
			dbxrefColumnsWithDefault,
			dbxrefColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(dbxrefType, dbxrefMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dbxrefType, dbxrefMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"dbxref\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into dbxref")
	}

	if !cached {
		dbxrefInsertCacheMut.Lock()
		dbxrefInsertCache[key] = cache
		dbxrefInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Dbxref record. See Update for
// whitelist behavior description.
func (o *Dbxref) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Dbxref record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Dbxref) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Dbxref, and panics on error.
// See Update for whitelist behavior description.
func (o *Dbxref) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Dbxref.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Dbxref) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	dbxrefUpdateCacheMut.RLock()
	cache, cached := dbxrefUpdateCache[key]
	dbxrefUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(dbxrefColumns, dbxrefPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update dbxref, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"dbxref\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dbxrefPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dbxrefType, dbxrefMapping, append(wl, dbxrefPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update dbxref row")
	}

	if !cached {
		dbxrefUpdateCacheMut.Lock()
		dbxrefUpdateCache[key] = cache
		dbxrefUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q dbxrefQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q dbxrefQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for dbxref")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DbxrefSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DbxrefSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DbxrefSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DbxrefSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"dbxref\" SET %s WHERE (\"dbxref_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dbxrefPrimaryKeyColumns), len(colNames)+1, len(dbxrefPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dbxref slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Dbxref) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Dbxref) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Dbxref) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Dbxref) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no dbxref provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dbxrefColumnsWithDefault, o)

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

	dbxrefUpsertCacheMut.RLock()
	cache, cached := dbxrefUpsertCache[key]
	dbxrefUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			dbxrefColumns,
			dbxrefColumnsWithDefault,
			dbxrefColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			dbxrefColumns,
			dbxrefPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert dbxref, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(dbxrefPrimaryKeyColumns))
			copy(conflict, dbxrefPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"dbxref\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(dbxrefType, dbxrefMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dbxrefType, dbxrefMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for dbxref")
	}

	if !cached {
		dbxrefUpsertCacheMut.Lock()
		dbxrefUpsertCache[key] = cache
		dbxrefUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Dbxref record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Dbxref) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Dbxref record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Dbxref) DeleteG() error {
	if o == nil {
		return errors.New("models: no Dbxref provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Dbxref record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Dbxref) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Dbxref record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Dbxref) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Dbxref provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dbxrefPrimaryKeyMapping)
	sql := "DELETE FROM \"dbxref\" WHERE \"dbxref_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from dbxref")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q dbxrefQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q dbxrefQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no dbxrefQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dbxref")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DbxrefSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DbxrefSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Dbxref slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DbxrefSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DbxrefSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Dbxref slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(dbxrefBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"dbxref\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dbxrefPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dbxrefPrimaryKeyColumns), 1, len(dbxrefPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dbxref slice")
	}

	if len(dbxrefAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Dbxref) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Dbxref) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Dbxref) ReloadG() error {
	if o == nil {
		return errors.New("models: no Dbxref provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Dbxref) Reload(exec boil.Executor) error {
	ret, err := FindDbxref(exec, o.DbxrefID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DbxrefSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DbxrefSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DbxrefSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DbxrefSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DbxrefSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	dbxrefs := DbxrefSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"dbxref\".* FROM \"dbxref\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dbxrefPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(dbxrefPrimaryKeyColumns), 1, len(dbxrefPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&dbxrefs)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DbxrefSlice")
	}

	*o = dbxrefs

	return nil
}

// DbxrefExists checks if the Dbxref row exists.
func DbxrefExists(exec boil.Executor, dbxrefID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"dbxref\" where \"dbxref_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, dbxrefID)
	}

	row := exec.QueryRow(sql, dbxrefID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if dbxref exists")
	}

	return exists, nil
}

// DbxrefExistsG checks if the Dbxref row exists.
func DbxrefExistsG(dbxrefID int) (bool, error) {
	return DbxrefExists(boil.GetDB(), dbxrefID)
}

// DbxrefExistsGP checks if the Dbxref row exists. Panics on error.
func DbxrefExistsGP(dbxrefID int) bool {
	e, err := DbxrefExists(boil.GetDB(), dbxrefID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DbxrefExistsP checks if the Dbxref row exists. Panics on error.
func DbxrefExistsP(exec boil.Executor, dbxrefID int) bool {
	e, err := DbxrefExists(exec, dbxrefID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
