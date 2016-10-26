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

// Genotype is an object representing the database table.
type Genotype struct {
	GenotypeID  int         `boil:"genotype_id" json:"genotype_id" toml:"genotype_id" yaml:"genotype_id"`
	Name        null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Uniquename  string      `boil:"uniquename" json:"uniquename" toml:"uniquename" yaml:"uniquename"`
	Description null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	TypeID      int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`

	R *genotypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L genotypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// genotypeR is where relationships are stored.
type genotypeR struct {
	Type                         *Cvterm
	FeatureGenotype              *FeatureGenotype
	Genotypeprop                 *Genotypeprop
	Phendesc                     *Phendesc
	Phenstatement                *Phenstatement
	StockGenotype                *StockGenotype
	Genotype1PhenotypeComparison *PhenotypeComparison
	Genotype2PhenotypeComparison *PhenotypeComparison
}

// genotypeL is where Load methods for each relationship are stored.
type genotypeL struct{}

var (
	genotypeColumns               = []string{"genotype_id", "name", "uniquename", "description", "type_id"}
	genotypeColumnsWithoutDefault = []string{"name", "uniquename", "description", "type_id"}
	genotypeColumnsWithDefault    = []string{"genotype_id"}
	genotypePrimaryKeyColumns     = []string{"genotype_id"}
)

type (
	// GenotypeSlice is an alias for a slice of pointers to Genotype.
	// This should generally be used opposed to []Genotype.
	GenotypeSlice []*Genotype
	// GenotypeHook is the signature for custom Genotype hook methods
	GenotypeHook func(boil.Executor, *Genotype) error

	genotypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	genotypeType                 = reflect.TypeOf(&Genotype{})
	genotypeMapping              = queries.MakeStructMapping(genotypeType)
	genotypePrimaryKeyMapping, _ = queries.BindMapping(genotypeType, genotypeMapping, genotypePrimaryKeyColumns)
	genotypeInsertCacheMut       sync.RWMutex
	genotypeInsertCache          = make(map[string]insertCache)
	genotypeUpdateCacheMut       sync.RWMutex
	genotypeUpdateCache          = make(map[string]updateCache)
	genotypeUpsertCacheMut       sync.RWMutex
	genotypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var genotypeBeforeInsertHooks []GenotypeHook
var genotypeBeforeUpdateHooks []GenotypeHook
var genotypeBeforeDeleteHooks []GenotypeHook
var genotypeBeforeUpsertHooks []GenotypeHook

var genotypeAfterInsertHooks []GenotypeHook
var genotypeAfterSelectHooks []GenotypeHook
var genotypeAfterUpdateHooks []GenotypeHook
var genotypeAfterDeleteHooks []GenotypeHook
var genotypeAfterUpsertHooks []GenotypeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Genotype) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range genotypeBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Genotype) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range genotypeBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Genotype) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range genotypeBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Genotype) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range genotypeBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Genotype) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range genotypeAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Genotype) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range genotypeAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Genotype) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range genotypeAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Genotype) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range genotypeAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Genotype) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range genotypeAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGenotypeHook registers your hook function for all future operations.
func AddGenotypeHook(hookPoint boil.HookPoint, genotypeHook GenotypeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		genotypeBeforeInsertHooks = append(genotypeBeforeInsertHooks, genotypeHook)
	case boil.BeforeUpdateHook:
		genotypeBeforeUpdateHooks = append(genotypeBeforeUpdateHooks, genotypeHook)
	case boil.BeforeDeleteHook:
		genotypeBeforeDeleteHooks = append(genotypeBeforeDeleteHooks, genotypeHook)
	case boil.BeforeUpsertHook:
		genotypeBeforeUpsertHooks = append(genotypeBeforeUpsertHooks, genotypeHook)
	case boil.AfterInsertHook:
		genotypeAfterInsertHooks = append(genotypeAfterInsertHooks, genotypeHook)
	case boil.AfterSelectHook:
		genotypeAfterSelectHooks = append(genotypeAfterSelectHooks, genotypeHook)
	case boil.AfterUpdateHook:
		genotypeAfterUpdateHooks = append(genotypeAfterUpdateHooks, genotypeHook)
	case boil.AfterDeleteHook:
		genotypeAfterDeleteHooks = append(genotypeAfterDeleteHooks, genotypeHook)
	case boil.AfterUpsertHook:
		genotypeAfterUpsertHooks = append(genotypeAfterUpsertHooks, genotypeHook)
	}
}

// OneP returns a single genotype record from the query, and panics on error.
func (q genotypeQuery) OneP() *Genotype {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single genotype record from the query.
func (q genotypeQuery) One() (*Genotype, error) {
	o := &Genotype{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for genotype")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Genotype records from the query, and panics on error.
func (q genotypeQuery) AllP() GenotypeSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Genotype records from the query.
func (q genotypeQuery) All() (GenotypeSlice, error) {
	var o GenotypeSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to Genotype slice")
	}

	if len(genotypeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Genotype records in the query, and panics on error.
func (q genotypeQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Genotype records in the query.
func (q genotypeQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count genotype rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q genotypeQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q genotypeQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if genotype exists")
	}

	return count > 0, nil
}

// TypeG pointed to by the foreign key.
func (o *Genotype) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Genotype) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// FeatureGenotypeG pointed to by the foreign key.
func (o *Genotype) FeatureGenotypeG(mods ...qm.QueryMod) featureGenotypeQuery {
	return o.FeatureGenotype(boil.GetDB(), mods...)
}

// FeatureGenotype pointed to by the foreign key.
func (o *Genotype) FeatureGenotype(exec boil.Executor, mods ...qm.QueryMod) featureGenotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("genotype_id=$1", o.GenotypeID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureGenotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_genotype\"")

	return query
}

// GenotypepropG pointed to by the foreign key.
func (o *Genotype) GenotypepropG(mods ...qm.QueryMod) genotypepropQuery {
	return o.Genotypeprop(boil.GetDB(), mods...)
}

// Genotypeprop pointed to by the foreign key.
func (o *Genotype) Genotypeprop(exec boil.Executor, mods ...qm.QueryMod) genotypepropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("genotype_id=$1", o.GenotypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Genotypeprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"genotypeprop\"")

	return query
}

// PhendescG pointed to by the foreign key.
func (o *Genotype) PhendescG(mods ...qm.QueryMod) phendescQuery {
	return o.Phendesc(boil.GetDB(), mods...)
}

// Phendesc pointed to by the foreign key.
func (o *Genotype) Phendesc(exec boil.Executor, mods ...qm.QueryMod) phendescQuery {
	queryMods := []qm.QueryMod{
		qm.Where("genotype_id=$1", o.GenotypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Phendescs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phendesc\"")

	return query
}

// PhenstatementG pointed to by the foreign key.
func (o *Genotype) PhenstatementG(mods ...qm.QueryMod) phenstatementQuery {
	return o.Phenstatement(boil.GetDB(), mods...)
}

// Phenstatement pointed to by the foreign key.
func (o *Genotype) Phenstatement(exec boil.Executor, mods ...qm.QueryMod) phenstatementQuery {
	queryMods := []qm.QueryMod{
		qm.Where("genotype_id=$1", o.GenotypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Phenstatements(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenstatement\"")

	return query
}

// StockGenotypeG pointed to by the foreign key.
func (o *Genotype) StockGenotypeG(mods ...qm.QueryMod) stockGenotypeQuery {
	return o.StockGenotype(boil.GetDB(), mods...)
}

// StockGenotype pointed to by the foreign key.
func (o *Genotype) StockGenotype(exec boil.Executor, mods ...qm.QueryMod) stockGenotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("genotype_id=$1", o.GenotypeID),
	}

	queryMods = append(queryMods, mods...)

	query := StockGenotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_genotype\"")

	return query
}

// Genotype1PhenotypeComparisonG pointed to by the foreign key.
func (o *Genotype) Genotype1PhenotypeComparisonG(mods ...qm.QueryMod) phenotypeComparisonQuery {
	return o.Genotype1PhenotypeComparison(boil.GetDB(), mods...)
}

// Genotype1PhenotypeComparison pointed to by the foreign key.
func (o *Genotype) Genotype1PhenotypeComparison(exec boil.Executor, mods ...qm.QueryMod) phenotypeComparisonQuery {
	queryMods := []qm.QueryMod{
		qm.Where("genotype1_id=$1", o.GenotypeID),
	}

	queryMods = append(queryMods, mods...)

	query := PhenotypeComparisons(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype_comparison\"")

	return query
}

// Genotype2PhenotypeComparisonG pointed to by the foreign key.
func (o *Genotype) Genotype2PhenotypeComparisonG(mods ...qm.QueryMod) phenotypeComparisonQuery {
	return o.Genotype2PhenotypeComparison(boil.GetDB(), mods...)
}

// Genotype2PhenotypeComparison pointed to by the foreign key.
func (o *Genotype) Genotype2PhenotypeComparison(exec boil.Executor, mods ...qm.QueryMod) phenotypeComparisonQuery {
	queryMods := []qm.QueryMod{
		qm.Where("genotype2_id=$1", o.GenotypeID),
	}

	queryMods = append(queryMods, mods...)

	query := PhenotypeComparisons(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype_comparison\"")

	return query
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (genotypeL) LoadType(e boil.Executor, singular bool, maybeGenotype interface{}) error {
	var slice []*Genotype
	var object *Genotype

	count := 1
	if singular {
		object = maybeGenotype.(*Genotype)
	} else {
		slice = *maybeGenotype.(*GenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &genotypeR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &genotypeR{}
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

	if len(genotypeAfterSelectHooks) != 0 {
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

// LoadFeatureGenotype allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (genotypeL) LoadFeatureGenotype(e boil.Executor, singular bool, maybeGenotype interface{}) error {
	var slice []*Genotype
	var object *Genotype

	count := 1
	if singular {
		object = maybeGenotype.(*Genotype)
	} else {
		slice = *maybeGenotype.(*GenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &genotypeR{}
		args[0] = object.GenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &genotypeR{}
			args[i] = obj.GenotypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_genotype\" where \"genotype_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureGenotype")
	}
	defer results.Close()

	var resultSlice []*FeatureGenotype
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureGenotype")
	}

	if len(genotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeatureGenotype = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.GenotypeID == foreign.GenotypeID {
				local.R.FeatureGenotype = foreign
				break
			}
		}
	}

	return nil
}

// LoadGenotypeprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (genotypeL) LoadGenotypeprop(e boil.Executor, singular bool, maybeGenotype interface{}) error {
	var slice []*Genotype
	var object *Genotype

	count := 1
	if singular {
		object = maybeGenotype.(*Genotype)
	} else {
		slice = *maybeGenotype.(*GenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &genotypeR{}
		args[0] = object.GenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &genotypeR{}
			args[i] = obj.GenotypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"genotypeprop\" where \"genotype_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Genotypeprop")
	}
	defer results.Close()

	var resultSlice []*Genotypeprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Genotypeprop")
	}

	if len(genotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Genotypeprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.GenotypeID == foreign.GenotypeID {
				local.R.Genotypeprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadPhendesc allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (genotypeL) LoadPhendesc(e boil.Executor, singular bool, maybeGenotype interface{}) error {
	var slice []*Genotype
	var object *Genotype

	count := 1
	if singular {
		object = maybeGenotype.(*Genotype)
	} else {
		slice = *maybeGenotype.(*GenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &genotypeR{}
		args[0] = object.GenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &genotypeR{}
			args[i] = obj.GenotypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phendesc\" where \"genotype_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Phendesc")
	}
	defer results.Close()

	var resultSlice []*Phendesc
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Phendesc")
	}

	if len(genotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Phendesc = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.GenotypeID == foreign.GenotypeID {
				local.R.Phendesc = foreign
				break
			}
		}
	}

	return nil
}

// LoadPhenstatement allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (genotypeL) LoadPhenstatement(e boil.Executor, singular bool, maybeGenotype interface{}) error {
	var slice []*Genotype
	var object *Genotype

	count := 1
	if singular {
		object = maybeGenotype.(*Genotype)
	} else {
		slice = *maybeGenotype.(*GenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &genotypeR{}
		args[0] = object.GenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &genotypeR{}
			args[i] = obj.GenotypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenstatement\" where \"genotype_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Phenstatement")
	}
	defer results.Close()

	var resultSlice []*Phenstatement
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Phenstatement")
	}

	if len(genotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Phenstatement = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.GenotypeID == foreign.GenotypeID {
				local.R.Phenstatement = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockGenotype allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (genotypeL) LoadStockGenotype(e boil.Executor, singular bool, maybeGenotype interface{}) error {
	var slice []*Genotype
	var object *Genotype

	count := 1
	if singular {
		object = maybeGenotype.(*Genotype)
	} else {
		slice = *maybeGenotype.(*GenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &genotypeR{}
		args[0] = object.GenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &genotypeR{}
			args[i] = obj.GenotypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_genotype\" where \"genotype_id\" in (%s)",
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

	if len(genotypeAfterSelectHooks) != 0 {
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
			if local.GenotypeID == foreign.GenotypeID {
				local.R.StockGenotype = foreign
				break
			}
		}
	}

	return nil
}

// LoadGenotype1PhenotypeComparison allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (genotypeL) LoadGenotype1PhenotypeComparison(e boil.Executor, singular bool, maybeGenotype interface{}) error {
	var slice []*Genotype
	var object *Genotype

	count := 1
	if singular {
		object = maybeGenotype.(*Genotype)
	} else {
		slice = *maybeGenotype.(*GenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &genotypeR{}
		args[0] = object.GenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &genotypeR{}
			args[i] = obj.GenotypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype_comparison\" where \"genotype1_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load PhenotypeComparison")
	}
	defer results.Close()

	var resultSlice []*PhenotypeComparison
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice PhenotypeComparison")
	}

	if len(genotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Genotype1PhenotypeComparison = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.GenotypeID == foreign.Genotype1ID {
				local.R.Genotype1PhenotypeComparison = foreign
				break
			}
		}
	}

	return nil
}

// LoadGenotype2PhenotypeComparison allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (genotypeL) LoadGenotype2PhenotypeComparison(e boil.Executor, singular bool, maybeGenotype interface{}) error {
	var slice []*Genotype
	var object *Genotype

	count := 1
	if singular {
		object = maybeGenotype.(*Genotype)
	} else {
		slice = *maybeGenotype.(*GenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &genotypeR{}
		args[0] = object.GenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &genotypeR{}
			args[i] = obj.GenotypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype_comparison\" where \"genotype2_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load PhenotypeComparison")
	}
	defer results.Close()

	var resultSlice []*PhenotypeComparison
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice PhenotypeComparison")
	}

	if len(genotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Genotype2PhenotypeComparison = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.GenotypeID == foreign.Genotype2ID {
				local.R.Genotype2PhenotypeComparison = foreign
				break
			}
		}
	}

	return nil
}

// SetType of the genotype to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeGenotypes.
func (o *Genotype) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"genotype\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, genotypePrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.GenotypeID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &genotypeR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeGenotypes: GenotypeSlice{o},
		}
	} else {
		related.R.TypeGenotypes = append(related.R.TypeGenotypes, o)
	}

	return nil
}

// SetFeatureGenotype of the genotype to the related item.
// Sets o.R.FeatureGenotype to related.
// Adds o to related.R.Genotype.
func (o *Genotype) SetFeatureGenotype(exec boil.Executor, insert bool, related *FeatureGenotype) error {
	var err error

	if insert {
		related.GenotypeID = o.GenotypeID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_genotype\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"genotype_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureGenotypePrimaryKeyColumns),
		)
		values := []interface{}{o.GenotypeID, related.FeatureGenotypeID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.GenotypeID = o.GenotypeID

	}

	if o.R == nil {
		o.R = &genotypeR{
			FeatureGenotype: related,
		}
	} else {
		o.R.FeatureGenotype = related
	}

	if related.R == nil {
		related.R = &featureGenotypeR{
			Genotype: o,
		}
	} else {
		related.R.Genotype = o
	}
	return nil
}

// SetGenotypeprop of the genotype to the related item.
// Sets o.R.Genotypeprop to related.
// Adds o to related.R.Genotype.
func (o *Genotype) SetGenotypeprop(exec boil.Executor, insert bool, related *Genotypeprop) error {
	var err error

	if insert {
		related.GenotypeID = o.GenotypeID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"genotypeprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"genotype_id"}),
			strmangle.WhereClause("\"", "\"", 2, genotypepropPrimaryKeyColumns),
		)
		values := []interface{}{o.GenotypeID, related.GenotypepropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.GenotypeID = o.GenotypeID

	}

	if o.R == nil {
		o.R = &genotypeR{
			Genotypeprop: related,
		}
	} else {
		o.R.Genotypeprop = related
	}

	if related.R == nil {
		related.R = &genotypepropR{
			Genotype: o,
		}
	} else {
		related.R.Genotype = o
	}
	return nil
}

// SetPhendesc of the genotype to the related item.
// Sets o.R.Phendesc to related.
// Adds o to related.R.Genotype.
func (o *Genotype) SetPhendesc(exec boil.Executor, insert bool, related *Phendesc) error {
	var err error

	if insert {
		related.GenotypeID = o.GenotypeID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phendesc\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"genotype_id"}),
			strmangle.WhereClause("\"", "\"", 2, phendescPrimaryKeyColumns),
		)
		values := []interface{}{o.GenotypeID, related.PhendescID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.GenotypeID = o.GenotypeID

	}

	if o.R == nil {
		o.R = &genotypeR{
			Phendesc: related,
		}
	} else {
		o.R.Phendesc = related
	}

	if related.R == nil {
		related.R = &phendescR{
			Genotype: o,
		}
	} else {
		related.R.Genotype = o
	}
	return nil
}

// SetPhenstatement of the genotype to the related item.
// Sets o.R.Phenstatement to related.
// Adds o to related.R.Genotype.
func (o *Genotype) SetPhenstatement(exec boil.Executor, insert bool, related *Phenstatement) error {
	var err error

	if insert {
		related.GenotypeID = o.GenotypeID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phenstatement\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"genotype_id"}),
			strmangle.WhereClause("\"", "\"", 2, phenstatementPrimaryKeyColumns),
		)
		values := []interface{}{o.GenotypeID, related.PhenstatementID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.GenotypeID = o.GenotypeID

	}

	if o.R == nil {
		o.R = &genotypeR{
			Phenstatement: related,
		}
	} else {
		o.R.Phenstatement = related
	}

	if related.R == nil {
		related.R = &phenstatementR{
			Genotype: o,
		}
	} else {
		related.R.Genotype = o
	}
	return nil
}

// SetStockGenotype of the genotype to the related item.
// Sets o.R.StockGenotype to related.
// Adds o to related.R.Genotype.
func (o *Genotype) SetStockGenotype(exec boil.Executor, insert bool, related *StockGenotype) error {
	var err error

	if insert {
		related.GenotypeID = o.GenotypeID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock_genotype\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"genotype_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockGenotypePrimaryKeyColumns),
		)
		values := []interface{}{o.GenotypeID, related.StockGenotypeID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.GenotypeID = o.GenotypeID

	}

	if o.R == nil {
		o.R = &genotypeR{
			StockGenotype: related,
		}
	} else {
		o.R.StockGenotype = related
	}

	if related.R == nil {
		related.R = &stockGenotypeR{
			Genotype: o,
		}
	} else {
		related.R.Genotype = o
	}
	return nil
}

// SetGenotype1PhenotypeComparison of the genotype to the related item.
// Sets o.R.Genotype1PhenotypeComparison to related.
// Adds o to related.R.Genotype1.
func (o *Genotype) SetGenotype1PhenotypeComparison(exec boil.Executor, insert bool, related *PhenotypeComparison) error {
	var err error

	if insert {
		related.Genotype1ID = o.GenotypeID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phenotype_comparison\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"genotype1_id"}),
			strmangle.WhereClause("\"", "\"", 2, phenotypeComparisonPrimaryKeyColumns),
		)
		values := []interface{}{o.GenotypeID, related.PhenotypeComparisonID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.Genotype1ID = o.GenotypeID

	}

	if o.R == nil {
		o.R = &genotypeR{
			Genotype1PhenotypeComparison: related,
		}
	} else {
		o.R.Genotype1PhenotypeComparison = related
	}

	if related.R == nil {
		related.R = &phenotypeComparisonR{
			Genotype1: o,
		}
	} else {
		related.R.Genotype1 = o
	}
	return nil
}

// SetGenotype2PhenotypeComparison of the genotype to the related item.
// Sets o.R.Genotype2PhenotypeComparison to related.
// Adds o to related.R.Genotype2.
func (o *Genotype) SetGenotype2PhenotypeComparison(exec boil.Executor, insert bool, related *PhenotypeComparison) error {
	var err error

	if insert {
		related.Genotype2ID = o.GenotypeID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phenotype_comparison\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"genotype2_id"}),
			strmangle.WhereClause("\"", "\"", 2, phenotypeComparisonPrimaryKeyColumns),
		)
		values := []interface{}{o.GenotypeID, related.PhenotypeComparisonID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.Genotype2ID = o.GenotypeID

	}

	if o.R == nil {
		o.R = &genotypeR{
			Genotype2PhenotypeComparison: related,
		}
	} else {
		o.R.Genotype2PhenotypeComparison = related
	}

	if related.R == nil {
		related.R = &phenotypeComparisonR{
			Genotype2: o,
		}
	} else {
		related.R.Genotype2 = o
	}
	return nil
}

// GenotypesG retrieves all records.
func GenotypesG(mods ...qm.QueryMod) genotypeQuery {
	return Genotypes(boil.GetDB(), mods...)
}

// Genotypes retrieves all the records using an executor.
func Genotypes(exec boil.Executor, mods ...qm.QueryMod) genotypeQuery {
	mods = append(mods, qm.From("\"genotype\""))
	return genotypeQuery{NewQuery(exec, mods...)}
}

// FindGenotypeG retrieves a single record by ID.
func FindGenotypeG(genotypeID int, selectCols ...string) (*Genotype, error) {
	return FindGenotype(boil.GetDB(), genotypeID, selectCols...)
}

// FindGenotypeGP retrieves a single record by ID, and panics on error.
func FindGenotypeGP(genotypeID int, selectCols ...string) *Genotype {
	retobj, err := FindGenotype(boil.GetDB(), genotypeID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindGenotype retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGenotype(exec boil.Executor, genotypeID int, selectCols ...string) (*Genotype, error) {
	genotypeObj := &Genotype{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"genotype\" where \"genotype_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, genotypeID)

	err := q.Bind(genotypeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from genotype")
	}

	return genotypeObj, nil
}

// FindGenotypeP retrieves a single record by ID with an executor, and panics on error.
func FindGenotypeP(exec boil.Executor, genotypeID int, selectCols ...string) *Genotype {
	retobj, err := FindGenotype(exec, genotypeID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Genotype) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Genotype) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Genotype) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Genotype) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no genotype provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(genotypeColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	genotypeInsertCacheMut.RLock()
	cache, cached := genotypeInsertCache[key]
	genotypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			genotypeColumns,
			genotypeColumnsWithDefault,
			genotypeColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(genotypeType, genotypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(genotypeType, genotypeMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"genotype\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into genotype")
	}

	if !cached {
		genotypeInsertCacheMut.Lock()
		genotypeInsertCache[key] = cache
		genotypeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Genotype record. See Update for
// whitelist behavior description.
func (o *Genotype) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Genotype record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Genotype) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Genotype, and panics on error.
// See Update for whitelist behavior description.
func (o *Genotype) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Genotype.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Genotype) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	genotypeUpdateCacheMut.RLock()
	cache, cached := genotypeUpdateCache[key]
	genotypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(genotypeColumns, genotypePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update genotype, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"genotype\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, genotypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(genotypeType, genotypeMapping, append(wl, genotypePrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update genotype row")
	}

	if !cached {
		genotypeUpdateCacheMut.Lock()
		genotypeUpdateCache[key] = cache
		genotypeUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q genotypeQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q genotypeQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for genotype")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o GenotypeSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o GenotypeSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o GenotypeSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GenotypeSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), genotypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"genotype\" SET %s WHERE (\"genotype_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(genotypePrimaryKeyColumns), len(colNames)+1, len(genotypePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in genotype slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Genotype) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Genotype) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Genotype) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Genotype) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no genotype provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(genotypeColumnsWithDefault, o)

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

	genotypeUpsertCacheMut.RLock()
	cache, cached := genotypeUpsertCache[key]
	genotypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			genotypeColumns,
			genotypeColumnsWithDefault,
			genotypeColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			genotypeColumns,
			genotypePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert genotype, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(genotypePrimaryKeyColumns))
			copy(conflict, genotypePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"genotype\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(genotypeType, genotypeMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(genotypeType, genotypeMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for genotype")
	}

	if !cached {
		genotypeUpsertCacheMut.Lock()
		genotypeUpsertCache[key] = cache
		genotypeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Genotype record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Genotype) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Genotype record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Genotype) DeleteG() error {
	if o == nil {
		return errors.New("chado: no Genotype provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Genotype record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Genotype) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Genotype record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Genotype) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Genotype provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), genotypePrimaryKeyMapping)
	sql := "DELETE FROM \"genotype\" WHERE \"genotype_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from genotype")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q genotypeQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q genotypeQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no genotypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from genotype")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o GenotypeSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o GenotypeSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no Genotype slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o GenotypeSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GenotypeSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Genotype slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(genotypeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), genotypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"genotype\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, genotypePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(genotypePrimaryKeyColumns), 1, len(genotypePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from genotype slice")
	}

	if len(genotypeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Genotype) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Genotype) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Genotype) ReloadG() error {
	if o == nil {
		return errors.New("chado: no Genotype provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Genotype) Reload(exec boil.Executor) error {
	ret, err := FindGenotype(exec, o.GenotypeID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *GenotypeSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *GenotypeSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GenotypeSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty GenotypeSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GenotypeSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	genotypes := GenotypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), genotypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"genotype\".* FROM \"genotype\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, genotypePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(genotypePrimaryKeyColumns), 1, len(genotypePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&genotypes)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in GenotypeSlice")
	}

	*o = genotypes

	return nil
}

// GenotypeExists checks if the Genotype row exists.
func GenotypeExists(exec boil.Executor, genotypeID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"genotype\" where \"genotype_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, genotypeID)
	}

	row := exec.QueryRow(sql, genotypeID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if genotype exists")
	}

	return exists, nil
}

// GenotypeExistsG checks if the Genotype row exists.
func GenotypeExistsG(genotypeID int) (bool, error) {
	return GenotypeExists(boil.GetDB(), genotypeID)
}

// GenotypeExistsGP checks if the Genotype row exists. Panics on error.
func GenotypeExistsGP(genotypeID int) bool {
	e, err := GenotypeExists(boil.GetDB(), genotypeID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// GenotypeExistsP checks if the Genotype row exists. Panics on error.
func GenotypeExistsP(exec boil.Executor, genotypeID int) bool {
	e, err := GenotypeExists(exec, genotypeID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
