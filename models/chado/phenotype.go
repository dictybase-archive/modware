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

// Phenotype is an object representing the database table.
type Phenotype struct {
	PhenotypeID  int         `boil:"phenotype_id" json:"phenotype_id" toml:"phenotype_id" yaml:"phenotype_id"`
	Uniquename   string      `boil:"uniquename" json:"uniquename" toml:"uniquename" yaml:"uniquename"`
	Name         null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	ObservableID null.Int    `boil:"observable_id" json:"observable_id,omitempty" toml:"observable_id" yaml:"observable_id,omitempty"`
	AttrID       null.Int    `boil:"attr_id" json:"attr_id,omitempty" toml:"attr_id" yaml:"attr_id,omitempty"`
	Value        null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	CvalueID     null.Int    `boil:"cvalue_id" json:"cvalue_id,omitempty" toml:"cvalue_id" yaml:"cvalue_id,omitempty"`
	AssayID      null.Int    `boil:"assay_id" json:"assay_id,omitempty" toml:"assay_id" yaml:"assay_id,omitempty"`

	R *phenotypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L phenotypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// phenotypeR is where relationships are stored.
type phenotypeR struct {
	Assay                          *Cvterm
	Attr                           *Cvterm
	Cvalue                         *Cvterm
	Observable                     *Cvterm
	Phenotypeprop                  *Phenotypeprop
	FeaturePhenotype               *FeaturePhenotype
	Phenstatement                  *Phenstatement
	PhenotypeCvterm                *PhenotypeCvterm
	Phenotype1PhenotypeComparison  *PhenotypeComparison
	Phenotype2PhenotypeComparisons PhenotypeComparisonSlice
}

// phenotypeL is where Load methods for each relationship are stored.
type phenotypeL struct{}

var (
	phenotypeColumns               = []string{"phenotype_id", "uniquename", "name", "observable_id", "attr_id", "value", "cvalue_id", "assay_id"}
	phenotypeColumnsWithoutDefault = []string{"uniquename", "name", "observable_id", "attr_id", "value", "cvalue_id", "assay_id"}
	phenotypeColumnsWithDefault    = []string{"phenotype_id"}
	phenotypePrimaryKeyColumns     = []string{"phenotype_id"}
)

type (
	// PhenotypeSlice is an alias for a slice of pointers to Phenotype.
	// This should generally be used opposed to []Phenotype.
	PhenotypeSlice []*Phenotype
	// PhenotypeHook is the signature for custom Phenotype hook methods
	PhenotypeHook func(boil.Executor, *Phenotype) error

	phenotypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	phenotypeType                 = reflect.TypeOf(&Phenotype{})
	phenotypeMapping              = queries.MakeStructMapping(phenotypeType)
	phenotypePrimaryKeyMapping, _ = queries.BindMapping(phenotypeType, phenotypeMapping, phenotypePrimaryKeyColumns)
	phenotypeInsertCacheMut       sync.RWMutex
	phenotypeInsertCache          = make(map[string]insertCache)
	phenotypeUpdateCacheMut       sync.RWMutex
	phenotypeUpdateCache          = make(map[string]updateCache)
	phenotypeUpsertCacheMut       sync.RWMutex
	phenotypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var phenotypeBeforeInsertHooks []PhenotypeHook
var phenotypeBeforeUpdateHooks []PhenotypeHook
var phenotypeBeforeDeleteHooks []PhenotypeHook
var phenotypeBeforeUpsertHooks []PhenotypeHook

var phenotypeAfterInsertHooks []PhenotypeHook
var phenotypeAfterSelectHooks []PhenotypeHook
var phenotypeAfterUpdateHooks []PhenotypeHook
var phenotypeAfterDeleteHooks []PhenotypeHook
var phenotypeAfterUpsertHooks []PhenotypeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Phenotype) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Phenotype) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Phenotype) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Phenotype) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Phenotype) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Phenotype) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Phenotype) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Phenotype) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Phenotype) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPhenotypeHook registers your hook function for all future operations.
func AddPhenotypeHook(hookPoint boil.HookPoint, phenotypeHook PhenotypeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		phenotypeBeforeInsertHooks = append(phenotypeBeforeInsertHooks, phenotypeHook)
	case boil.BeforeUpdateHook:
		phenotypeBeforeUpdateHooks = append(phenotypeBeforeUpdateHooks, phenotypeHook)
	case boil.BeforeDeleteHook:
		phenotypeBeforeDeleteHooks = append(phenotypeBeforeDeleteHooks, phenotypeHook)
	case boil.BeforeUpsertHook:
		phenotypeBeforeUpsertHooks = append(phenotypeBeforeUpsertHooks, phenotypeHook)
	case boil.AfterInsertHook:
		phenotypeAfterInsertHooks = append(phenotypeAfterInsertHooks, phenotypeHook)
	case boil.AfterSelectHook:
		phenotypeAfterSelectHooks = append(phenotypeAfterSelectHooks, phenotypeHook)
	case boil.AfterUpdateHook:
		phenotypeAfterUpdateHooks = append(phenotypeAfterUpdateHooks, phenotypeHook)
	case boil.AfterDeleteHook:
		phenotypeAfterDeleteHooks = append(phenotypeAfterDeleteHooks, phenotypeHook)
	case boil.AfterUpsertHook:
		phenotypeAfterUpsertHooks = append(phenotypeAfterUpsertHooks, phenotypeHook)
	}
}

// OneP returns a single phenotype record from the query, and panics on error.
func (q phenotypeQuery) OneP() *Phenotype {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single phenotype record from the query.
func (q phenotypeQuery) One() (*Phenotype, error) {
	o := &Phenotype{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for phenotype")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Phenotype records from the query, and panics on error.
func (q phenotypeQuery) AllP() PhenotypeSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Phenotype records from the query.
func (q phenotypeQuery) All() (PhenotypeSlice, error) {
	var o PhenotypeSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to Phenotype slice")
	}

	if len(phenotypeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Phenotype records in the query, and panics on error.
func (q phenotypeQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Phenotype records in the query.
func (q phenotypeQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count phenotype rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q phenotypeQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q phenotypeQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if phenotype exists")
	}

	return count > 0, nil
}

// AssayG pointed to by the foreign key.
func (o *Phenotype) AssayG(mods ...qm.QueryMod) cvtermQuery {
	return o.Assay(boil.GetDB(), mods...)
}

// Assay pointed to by the foreign key.
func (o *Phenotype) Assay(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.AssayID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// AttrG pointed to by the foreign key.
func (o *Phenotype) AttrG(mods ...qm.QueryMod) cvtermQuery {
	return o.Attr(boil.GetDB(), mods...)
}

// Attr pointed to by the foreign key.
func (o *Phenotype) Attr(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.AttrID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// CvalueG pointed to by the foreign key.
func (o *Phenotype) CvalueG(mods ...qm.QueryMod) cvtermQuery {
	return o.Cvalue(boil.GetDB(), mods...)
}

// Cvalue pointed to by the foreign key.
func (o *Phenotype) Cvalue(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.CvalueID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// ObservableG pointed to by the foreign key.
func (o *Phenotype) ObservableG(mods ...qm.QueryMod) cvtermQuery {
	return o.Observable(boil.GetDB(), mods...)
}

// Observable pointed to by the foreign key.
func (o *Phenotype) Observable(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.ObservableID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// PhenotypepropG pointed to by the foreign key.
func (o *Phenotype) PhenotypepropG(mods ...qm.QueryMod) phenotypepropQuery {
	return o.Phenotypeprop(boil.GetDB(), mods...)
}

// Phenotypeprop pointed to by the foreign key.
func (o *Phenotype) Phenotypeprop(exec boil.Executor, mods ...qm.QueryMod) phenotypepropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("phenotype_id=$1", o.PhenotypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Phenotypeprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotypeprop\"")

	return query
}

// FeaturePhenotypeG pointed to by the foreign key.
func (o *Phenotype) FeaturePhenotypeG(mods ...qm.QueryMod) featurePhenotypeQuery {
	return o.FeaturePhenotype(boil.GetDB(), mods...)
}

// FeaturePhenotype pointed to by the foreign key.
func (o *Phenotype) FeaturePhenotype(exec boil.Executor, mods ...qm.QueryMod) featurePhenotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("phenotype_id=$1", o.PhenotypeID),
	}

	queryMods = append(queryMods, mods...)

	query := FeaturePhenotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_phenotype\"")

	return query
}

// PhenstatementG pointed to by the foreign key.
func (o *Phenotype) PhenstatementG(mods ...qm.QueryMod) phenstatementQuery {
	return o.Phenstatement(boil.GetDB(), mods...)
}

// Phenstatement pointed to by the foreign key.
func (o *Phenotype) Phenstatement(exec boil.Executor, mods ...qm.QueryMod) phenstatementQuery {
	queryMods := []qm.QueryMod{
		qm.Where("phenotype_id=$1", o.PhenotypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Phenstatements(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenstatement\"")

	return query
}

// PhenotypeCvtermG pointed to by the foreign key.
func (o *Phenotype) PhenotypeCvtermG(mods ...qm.QueryMod) phenotypeCvtermQuery {
	return o.PhenotypeCvterm(boil.GetDB(), mods...)
}

// PhenotypeCvterm pointed to by the foreign key.
func (o *Phenotype) PhenotypeCvterm(exec boil.Executor, mods ...qm.QueryMod) phenotypeCvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("phenotype_id=$1", o.PhenotypeID),
	}

	queryMods = append(queryMods, mods...)

	query := PhenotypeCvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype_cvterm\"")

	return query
}

// Phenotype1PhenotypeComparisonG pointed to by the foreign key.
func (o *Phenotype) Phenotype1PhenotypeComparisonG(mods ...qm.QueryMod) phenotypeComparisonQuery {
	return o.Phenotype1PhenotypeComparison(boil.GetDB(), mods...)
}

// Phenotype1PhenotypeComparison pointed to by the foreign key.
func (o *Phenotype) Phenotype1PhenotypeComparison(exec boil.Executor, mods ...qm.QueryMod) phenotypeComparisonQuery {
	queryMods := []qm.QueryMod{
		qm.Where("phenotype1_id=$1", o.PhenotypeID),
	}

	queryMods = append(queryMods, mods...)

	query := PhenotypeComparisons(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype_comparison\"")

	return query
}

// Phenotype2PhenotypeComparisonsG retrieves all the phenotype_comparison's phenotype comparison via phenotype2_id column.
func (o *Phenotype) Phenotype2PhenotypeComparisonsG(mods ...qm.QueryMod) phenotypeComparisonQuery {
	return o.Phenotype2PhenotypeComparisons(boil.GetDB(), mods...)
}

// Phenotype2PhenotypeComparisons retrieves all the phenotype_comparison's phenotype comparison with an executor via phenotype2_id column.
func (o *Phenotype) Phenotype2PhenotypeComparisons(exec boil.Executor, mods ...qm.QueryMod) phenotypeComparisonQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"phenotype2_id\"=$1", o.PhenotypeID),
	)

	query := PhenotypeComparisons(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype_comparison\" as \"a\"")
	return query
}

// LoadAssay allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeL) LoadAssay(e boil.Executor, singular bool, maybePhenotype interface{}) error {
	var slice []*Phenotype
	var object *Phenotype

	count := 1
	if singular {
		object = maybePhenotype.(*Phenotype)
	} else {
		slice = *maybePhenotype.(*PhenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeR{}
		args[0] = object.AssayID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeR{}
			args[i] = obj.AssayID
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

	if len(phenotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Assay = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AssayID.Int == foreign.CvtermID {
				local.R.Assay = foreign
				break
			}
		}
	}

	return nil
}

// LoadAttr allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeL) LoadAttr(e boil.Executor, singular bool, maybePhenotype interface{}) error {
	var slice []*Phenotype
	var object *Phenotype

	count := 1
	if singular {
		object = maybePhenotype.(*Phenotype)
	} else {
		slice = *maybePhenotype.(*PhenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeR{}
		args[0] = object.AttrID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeR{}
			args[i] = obj.AttrID
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

	if len(phenotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Attr = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AttrID.Int == foreign.CvtermID {
				local.R.Attr = foreign
				break
			}
		}
	}

	return nil
}

// LoadCvalue allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeL) LoadCvalue(e boil.Executor, singular bool, maybePhenotype interface{}) error {
	var slice []*Phenotype
	var object *Phenotype

	count := 1
	if singular {
		object = maybePhenotype.(*Phenotype)
	} else {
		slice = *maybePhenotype.(*PhenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeR{}
		args[0] = object.CvalueID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeR{}
			args[i] = obj.CvalueID
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

	if len(phenotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Cvalue = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvalueID.Int == foreign.CvtermID {
				local.R.Cvalue = foreign
				break
			}
		}
	}

	return nil
}

// LoadObservable allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeL) LoadObservable(e boil.Executor, singular bool, maybePhenotype interface{}) error {
	var slice []*Phenotype
	var object *Phenotype

	count := 1
	if singular {
		object = maybePhenotype.(*Phenotype)
	} else {
		slice = *maybePhenotype.(*PhenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeR{}
		args[0] = object.ObservableID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeR{}
			args[i] = obj.ObservableID
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

	if len(phenotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Observable = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ObservableID.Int == foreign.CvtermID {
				local.R.Observable = foreign
				break
			}
		}
	}

	return nil
}

// LoadPhenotypeprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeL) LoadPhenotypeprop(e boil.Executor, singular bool, maybePhenotype interface{}) error {
	var slice []*Phenotype
	var object *Phenotype

	count := 1
	if singular {
		object = maybePhenotype.(*Phenotype)
	} else {
		slice = *maybePhenotype.(*PhenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeR{}
		args[0] = object.PhenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeR{}
			args[i] = obj.PhenotypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotypeprop\" where \"phenotype_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Phenotypeprop")
	}
	defer results.Close()

	var resultSlice []*Phenotypeprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Phenotypeprop")
	}

	if len(phenotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Phenotypeprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PhenotypeID == foreign.PhenotypeID {
				local.R.Phenotypeprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeaturePhenotype allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeL) LoadFeaturePhenotype(e boil.Executor, singular bool, maybePhenotype interface{}) error {
	var slice []*Phenotype
	var object *Phenotype

	count := 1
	if singular {
		object = maybePhenotype.(*Phenotype)
	} else {
		slice = *maybePhenotype.(*PhenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeR{}
		args[0] = object.PhenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeR{}
			args[i] = obj.PhenotypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_phenotype\" where \"phenotype_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeaturePhenotype")
	}
	defer results.Close()

	var resultSlice []*FeaturePhenotype
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeaturePhenotype")
	}

	if len(phenotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeaturePhenotype = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PhenotypeID == foreign.PhenotypeID {
				local.R.FeaturePhenotype = foreign
				break
			}
		}
	}

	return nil
}

// LoadPhenstatement allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeL) LoadPhenstatement(e boil.Executor, singular bool, maybePhenotype interface{}) error {
	var slice []*Phenotype
	var object *Phenotype

	count := 1
	if singular {
		object = maybePhenotype.(*Phenotype)
	} else {
		slice = *maybePhenotype.(*PhenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeR{}
		args[0] = object.PhenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeR{}
			args[i] = obj.PhenotypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenstatement\" where \"phenotype_id\" in (%s)",
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

	if len(phenotypeAfterSelectHooks) != 0 {
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
			if local.PhenotypeID == foreign.PhenotypeID {
				local.R.Phenstatement = foreign
				break
			}
		}
	}

	return nil
}

// LoadPhenotypeCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeL) LoadPhenotypeCvterm(e boil.Executor, singular bool, maybePhenotype interface{}) error {
	var slice []*Phenotype
	var object *Phenotype

	count := 1
	if singular {
		object = maybePhenotype.(*Phenotype)
	} else {
		slice = *maybePhenotype.(*PhenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeR{}
		args[0] = object.PhenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeR{}
			args[i] = obj.PhenotypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype_cvterm\" where \"phenotype_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load PhenotypeCvterm")
	}
	defer results.Close()

	var resultSlice []*PhenotypeCvterm
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice PhenotypeCvterm")
	}

	if len(phenotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.PhenotypeCvterm = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PhenotypeID == foreign.PhenotypeID {
				local.R.PhenotypeCvterm = foreign
				break
			}
		}
	}

	return nil
}

// LoadPhenotype1PhenotypeComparison allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeL) LoadPhenotype1PhenotypeComparison(e boil.Executor, singular bool, maybePhenotype interface{}) error {
	var slice []*Phenotype
	var object *Phenotype

	count := 1
	if singular {
		object = maybePhenotype.(*Phenotype)
	} else {
		slice = *maybePhenotype.(*PhenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeR{}
		args[0] = object.PhenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeR{}
			args[i] = obj.PhenotypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype_comparison\" where \"phenotype1_id\" in (%s)",
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

	if len(phenotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Phenotype1PhenotypeComparison = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PhenotypeID == foreign.Phenotype1ID {
				local.R.Phenotype1PhenotypeComparison = foreign
				break
			}
		}
	}

	return nil
}

// LoadPhenotype2PhenotypeComparisons allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeL) LoadPhenotype2PhenotypeComparisons(e boil.Executor, singular bool, maybePhenotype interface{}) error {
	var slice []*Phenotype
	var object *Phenotype

	count := 1
	if singular {
		object = maybePhenotype.(*Phenotype)
	} else {
		slice = *maybePhenotype.(*PhenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeR{}
		args[0] = object.PhenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeR{}
			args[i] = obj.PhenotypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype_comparison\" where \"phenotype2_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load phenotype_comparison")
	}
	defer results.Close()

	var resultSlice []*PhenotypeComparison
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice phenotype_comparison")
	}

	if len(phenotypeComparisonAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Phenotype2PhenotypeComparisons = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PhenotypeID == foreign.Phenotype2ID.Int {
				local.R.Phenotype2PhenotypeComparisons = append(local.R.Phenotype2PhenotypeComparisons, foreign)
				break
			}
		}
	}

	return nil
}

// SetAssay of the phenotype to the related item.
// Sets o.R.Assay to related.
// Adds o to related.R.AssayPhenotypes.
func (o *Phenotype) SetAssay(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenotype\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"assay_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenotypePrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.PhenotypeID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AssayID.Int = related.CvtermID
	o.AssayID.Valid = true

	if o.R == nil {
		o.R = &phenotypeR{
			Assay: related,
		}
	} else {
		o.R.Assay = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			AssayPhenotypes: PhenotypeSlice{o},
		}
	} else {
		related.R.AssayPhenotypes = append(related.R.AssayPhenotypes, o)
	}

	return nil
}

// RemoveAssay relationship.
// Sets o.R.Assay to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Phenotype) RemoveAssay(exec boil.Executor, related *Cvterm) error {
	var err error

	o.AssayID.Valid = false
	if err = o.Update(exec, "assay_id"); err != nil {
		o.AssayID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Assay = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.AssayPhenotypes {
		if o.AssayID.Int != ri.AssayID.Int {
			continue
		}

		ln := len(related.R.AssayPhenotypes)
		if ln > 1 && i < ln-1 {
			related.R.AssayPhenotypes[i] = related.R.AssayPhenotypes[ln-1]
		}
		related.R.AssayPhenotypes = related.R.AssayPhenotypes[:ln-1]
		break
	}
	return nil
}

// SetAttr of the phenotype to the related item.
// Sets o.R.Attr to related.
// Adds o to related.R.AttrPhenotypes.
func (o *Phenotype) SetAttr(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenotype\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"attr_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenotypePrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.PhenotypeID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AttrID.Int = related.CvtermID
	o.AttrID.Valid = true

	if o.R == nil {
		o.R = &phenotypeR{
			Attr: related,
		}
	} else {
		o.R.Attr = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			AttrPhenotypes: PhenotypeSlice{o},
		}
	} else {
		related.R.AttrPhenotypes = append(related.R.AttrPhenotypes, o)
	}

	return nil
}

// RemoveAttr relationship.
// Sets o.R.Attr to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Phenotype) RemoveAttr(exec boil.Executor, related *Cvterm) error {
	var err error

	o.AttrID.Valid = false
	if err = o.Update(exec, "attr_id"); err != nil {
		o.AttrID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Attr = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.AttrPhenotypes {
		if o.AttrID.Int != ri.AttrID.Int {
			continue
		}

		ln := len(related.R.AttrPhenotypes)
		if ln > 1 && i < ln-1 {
			related.R.AttrPhenotypes[i] = related.R.AttrPhenotypes[ln-1]
		}
		related.R.AttrPhenotypes = related.R.AttrPhenotypes[:ln-1]
		break
	}
	return nil
}

// SetCvalue of the phenotype to the related item.
// Sets o.R.Cvalue to related.
// Adds o to related.R.CvaluePhenotypes.
func (o *Phenotype) SetCvalue(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenotype\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"cvalue_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenotypePrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.PhenotypeID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CvalueID.Int = related.CvtermID
	o.CvalueID.Valid = true

	if o.R == nil {
		o.R = &phenotypeR{
			Cvalue: related,
		}
	} else {
		o.R.Cvalue = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			CvaluePhenotypes: PhenotypeSlice{o},
		}
	} else {
		related.R.CvaluePhenotypes = append(related.R.CvaluePhenotypes, o)
	}

	return nil
}

// RemoveCvalue relationship.
// Sets o.R.Cvalue to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Phenotype) RemoveCvalue(exec boil.Executor, related *Cvterm) error {
	var err error

	o.CvalueID.Valid = false
	if err = o.Update(exec, "cvalue_id"); err != nil {
		o.CvalueID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Cvalue = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.CvaluePhenotypes {
		if o.CvalueID.Int != ri.CvalueID.Int {
			continue
		}

		ln := len(related.R.CvaluePhenotypes)
		if ln > 1 && i < ln-1 {
			related.R.CvaluePhenotypes[i] = related.R.CvaluePhenotypes[ln-1]
		}
		related.R.CvaluePhenotypes = related.R.CvaluePhenotypes[:ln-1]
		break
	}
	return nil
}

// SetObservable of the phenotype to the related item.
// Sets o.R.Observable to related.
// Adds o to related.R.ObservablePhenotypes.
func (o *Phenotype) SetObservable(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenotype\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"observable_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenotypePrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.PhenotypeID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ObservableID.Int = related.CvtermID
	o.ObservableID.Valid = true

	if o.R == nil {
		o.R = &phenotypeR{
			Observable: related,
		}
	} else {
		o.R.Observable = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			ObservablePhenotypes: PhenotypeSlice{o},
		}
	} else {
		related.R.ObservablePhenotypes = append(related.R.ObservablePhenotypes, o)
	}

	return nil
}

// RemoveObservable relationship.
// Sets o.R.Observable to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Phenotype) RemoveObservable(exec boil.Executor, related *Cvterm) error {
	var err error

	o.ObservableID.Valid = false
	if err = o.Update(exec, "observable_id"); err != nil {
		o.ObservableID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Observable = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.ObservablePhenotypes {
		if o.ObservableID.Int != ri.ObservableID.Int {
			continue
		}

		ln := len(related.R.ObservablePhenotypes)
		if ln > 1 && i < ln-1 {
			related.R.ObservablePhenotypes[i] = related.R.ObservablePhenotypes[ln-1]
		}
		related.R.ObservablePhenotypes = related.R.ObservablePhenotypes[:ln-1]
		break
	}
	return nil
}

// SetPhenotypeprop of the phenotype to the related item.
// Sets o.R.Phenotypeprop to related.
// Adds o to related.R.Phenotype.
func (o *Phenotype) SetPhenotypeprop(exec boil.Executor, insert bool, related *Phenotypeprop) error {
	var err error

	if insert {
		related.PhenotypeID = o.PhenotypeID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phenotypeprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"phenotype_id"}),
			strmangle.WhereClause("\"", "\"", 2, phenotypepropPrimaryKeyColumns),
		)
		values := []interface{}{o.PhenotypeID, related.PhenotypepropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PhenotypeID = o.PhenotypeID

	}

	if o.R == nil {
		o.R = &phenotypeR{
			Phenotypeprop: related,
		}
	} else {
		o.R.Phenotypeprop = related
	}

	if related.R == nil {
		related.R = &phenotypepropR{
			Phenotype: o,
		}
	} else {
		related.R.Phenotype = o
	}
	return nil
}

// SetFeaturePhenotype of the phenotype to the related item.
// Sets o.R.FeaturePhenotype to related.
// Adds o to related.R.Phenotype.
func (o *Phenotype) SetFeaturePhenotype(exec boil.Executor, insert bool, related *FeaturePhenotype) error {
	var err error

	if insert {
		related.PhenotypeID = o.PhenotypeID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_phenotype\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"phenotype_id"}),
			strmangle.WhereClause("\"", "\"", 2, featurePhenotypePrimaryKeyColumns),
		)
		values := []interface{}{o.PhenotypeID, related.FeaturePhenotypeID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PhenotypeID = o.PhenotypeID

	}

	if o.R == nil {
		o.R = &phenotypeR{
			FeaturePhenotype: related,
		}
	} else {
		o.R.FeaturePhenotype = related
	}

	if related.R == nil {
		related.R = &featurePhenotypeR{
			Phenotype: o,
		}
	} else {
		related.R.Phenotype = o
	}
	return nil
}

// SetPhenstatement of the phenotype to the related item.
// Sets o.R.Phenstatement to related.
// Adds o to related.R.Phenotype.
func (o *Phenotype) SetPhenstatement(exec boil.Executor, insert bool, related *Phenstatement) error {
	var err error

	if insert {
		related.PhenotypeID = o.PhenotypeID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phenstatement\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"phenotype_id"}),
			strmangle.WhereClause("\"", "\"", 2, phenstatementPrimaryKeyColumns),
		)
		values := []interface{}{o.PhenotypeID, related.PhenstatementID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PhenotypeID = o.PhenotypeID

	}

	if o.R == nil {
		o.R = &phenotypeR{
			Phenstatement: related,
		}
	} else {
		o.R.Phenstatement = related
	}

	if related.R == nil {
		related.R = &phenstatementR{
			Phenotype: o,
		}
	} else {
		related.R.Phenotype = o
	}
	return nil
}

// SetPhenotypeCvterm of the phenotype to the related item.
// Sets o.R.PhenotypeCvterm to related.
// Adds o to related.R.Phenotype.
func (o *Phenotype) SetPhenotypeCvterm(exec boil.Executor, insert bool, related *PhenotypeCvterm) error {
	var err error

	if insert {
		related.PhenotypeID = o.PhenotypeID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phenotype_cvterm\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"phenotype_id"}),
			strmangle.WhereClause("\"", "\"", 2, phenotypeCvtermPrimaryKeyColumns),
		)
		values := []interface{}{o.PhenotypeID, related.PhenotypeCvtermID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PhenotypeID = o.PhenotypeID

	}

	if o.R == nil {
		o.R = &phenotypeR{
			PhenotypeCvterm: related,
		}
	} else {
		o.R.PhenotypeCvterm = related
	}

	if related.R == nil {
		related.R = &phenotypeCvtermR{
			Phenotype: o,
		}
	} else {
		related.R.Phenotype = o
	}
	return nil
}

// SetPhenotype1PhenotypeComparison of the phenotype to the related item.
// Sets o.R.Phenotype1PhenotypeComparison to related.
// Adds o to related.R.Phenotype1.
func (o *Phenotype) SetPhenotype1PhenotypeComparison(exec boil.Executor, insert bool, related *PhenotypeComparison) error {
	var err error

	if insert {
		related.Phenotype1ID = o.PhenotypeID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phenotype_comparison\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"phenotype1_id"}),
			strmangle.WhereClause("\"", "\"", 2, phenotypeComparisonPrimaryKeyColumns),
		)
		values := []interface{}{o.PhenotypeID, related.PhenotypeComparisonID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.Phenotype1ID = o.PhenotypeID

	}

	if o.R == nil {
		o.R = &phenotypeR{
			Phenotype1PhenotypeComparison: related,
		}
	} else {
		o.R.Phenotype1PhenotypeComparison = related
	}

	if related.R == nil {
		related.R = &phenotypeComparisonR{
			Phenotype1: o,
		}
	} else {
		related.R.Phenotype1 = o
	}
	return nil
}

// AddPhenotype2PhenotypeComparisons adds the given related objects to the existing relationships
// of the phenotype, optionally inserting them as new records.
// Appends related to o.R.Phenotype2PhenotypeComparisons.
// Sets related.R.Phenotype2 appropriately.
func (o *Phenotype) AddPhenotype2PhenotypeComparisons(exec boil.Executor, insert bool, related ...*PhenotypeComparison) error {
	var err error
	for _, rel := range related {
		rel.Phenotype2ID.Int = o.PhenotypeID
		rel.Phenotype2ID.Valid = true
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "phenotype2_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &phenotypeR{
			Phenotype2PhenotypeComparisons: related,
		}
	} else {
		o.R.Phenotype2PhenotypeComparisons = append(o.R.Phenotype2PhenotypeComparisons, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &phenotypeComparisonR{
				Phenotype2: o,
			}
		} else {
			rel.R.Phenotype2 = o
		}
	}
	return nil
}

// SetPhenotype2PhenotypeComparisons removes all previously related items of the
// phenotype replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Phenotype2's Phenotype2PhenotypeComparisons accordingly.
// Replaces o.R.Phenotype2PhenotypeComparisons with related.
// Sets related.R.Phenotype2's Phenotype2PhenotypeComparisons accordingly.
func (o *Phenotype) SetPhenotype2PhenotypeComparisons(exec boil.Executor, insert bool, related ...*PhenotypeComparison) error {
	query := "update \"phenotype_comparison\" set \"phenotype2_id\" = null where \"phenotype2_id\" = $1"
	values := []interface{}{o.PhenotypeID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.Phenotype2PhenotypeComparisons {
			rel.Phenotype2ID.Valid = false
			if rel.R == nil {
				continue
			}

			rel.R.Phenotype2 = nil
		}

		o.R.Phenotype2PhenotypeComparisons = nil
	}
	return o.AddPhenotype2PhenotypeComparisons(exec, insert, related...)
}

// RemovePhenotype2PhenotypeComparisons relationships from objects passed in.
// Removes related items from R.Phenotype2PhenotypeComparisons (uses pointer comparison, removal does not keep order)
// Sets related.R.Phenotype2.
func (o *Phenotype) RemovePhenotype2PhenotypeComparisons(exec boil.Executor, related ...*PhenotypeComparison) error {
	var err error
	for _, rel := range related {
		rel.Phenotype2ID.Valid = false
		if rel.R != nil {
			rel.R.Phenotype2 = nil
		}
		if err = rel.Update(exec, "phenotype2_id"); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Phenotype2PhenotypeComparisons {
			if rel != ri {
				continue
			}

			ln := len(o.R.Phenotype2PhenotypeComparisons)
			if ln > 1 && i < ln-1 {
				o.R.Phenotype2PhenotypeComparisons[i] = o.R.Phenotype2PhenotypeComparisons[ln-1]
			}
			o.R.Phenotype2PhenotypeComparisons = o.R.Phenotype2PhenotypeComparisons[:ln-1]
			break
		}
	}

	return nil
}

// PhenotypesG retrieves all records.
func PhenotypesG(mods ...qm.QueryMod) phenotypeQuery {
	return Phenotypes(boil.GetDB(), mods...)
}

// Phenotypes retrieves all the records using an executor.
func Phenotypes(exec boil.Executor, mods ...qm.QueryMod) phenotypeQuery {
	mods = append(mods, qm.From("\"phenotype\""))
	return phenotypeQuery{NewQuery(exec, mods...)}
}

// FindPhenotypeG retrieves a single record by ID.
func FindPhenotypeG(phenotypeID int, selectCols ...string) (*Phenotype, error) {
	return FindPhenotype(boil.GetDB(), phenotypeID, selectCols...)
}

// FindPhenotypeGP retrieves a single record by ID, and panics on error.
func FindPhenotypeGP(phenotypeID int, selectCols ...string) *Phenotype {
	retobj, err := FindPhenotype(boil.GetDB(), phenotypeID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindPhenotype retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPhenotype(exec boil.Executor, phenotypeID int, selectCols ...string) (*Phenotype, error) {
	phenotypeObj := &Phenotype{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"phenotype\" where \"phenotype_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, phenotypeID)

	err := q.Bind(phenotypeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from phenotype")
	}

	return phenotypeObj, nil
}

// FindPhenotypeP retrieves a single record by ID with an executor, and panics on error.
func FindPhenotypeP(exec boil.Executor, phenotypeID int, selectCols ...string) *Phenotype {
	retobj, err := FindPhenotype(exec, phenotypeID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Phenotype) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Phenotype) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Phenotype) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Phenotype) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no phenotype provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(phenotypeColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	phenotypeInsertCacheMut.RLock()
	cache, cached := phenotypeInsertCache[key]
	phenotypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			phenotypeColumns,
			phenotypeColumnsWithDefault,
			phenotypeColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(phenotypeType, phenotypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(phenotypeType, phenotypeMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"phenotype\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into phenotype")
	}

	if !cached {
		phenotypeInsertCacheMut.Lock()
		phenotypeInsertCache[key] = cache
		phenotypeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Phenotype record. See Update for
// whitelist behavior description.
func (o *Phenotype) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Phenotype record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Phenotype) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Phenotype, and panics on error.
// See Update for whitelist behavior description.
func (o *Phenotype) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Phenotype.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Phenotype) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	phenotypeUpdateCacheMut.RLock()
	cache, cached := phenotypeUpdateCache[key]
	phenotypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(phenotypeColumns, phenotypePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update phenotype, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"phenotype\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, phenotypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(phenotypeType, phenotypeMapping, append(wl, phenotypePrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update phenotype row")
	}

	if !cached {
		phenotypeUpdateCacheMut.Lock()
		phenotypeUpdateCache[key] = cache
		phenotypeUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q phenotypeQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q phenotypeQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for phenotype")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o PhenotypeSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o PhenotypeSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o PhenotypeSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PhenotypeSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), phenotypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"phenotype\" SET %s WHERE (\"phenotype_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(phenotypePrimaryKeyColumns), len(colNames)+1, len(phenotypePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in phenotype slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Phenotype) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Phenotype) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Phenotype) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Phenotype) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no phenotype provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(phenotypeColumnsWithDefault, o)

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

	phenotypeUpsertCacheMut.RLock()
	cache, cached := phenotypeUpsertCache[key]
	phenotypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			phenotypeColumns,
			phenotypeColumnsWithDefault,
			phenotypeColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			phenotypeColumns,
			phenotypePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert phenotype, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(phenotypePrimaryKeyColumns))
			copy(conflict, phenotypePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"phenotype\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(phenotypeType, phenotypeMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(phenotypeType, phenotypeMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for phenotype")
	}

	if !cached {
		phenotypeUpsertCacheMut.Lock()
		phenotypeUpsertCache[key] = cache
		phenotypeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Phenotype record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Phenotype) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Phenotype record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Phenotype) DeleteG() error {
	if o == nil {
		return errors.New("chado: no Phenotype provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Phenotype record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Phenotype) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Phenotype record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Phenotype) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Phenotype provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), phenotypePrimaryKeyMapping)
	sql := "DELETE FROM \"phenotype\" WHERE \"phenotype_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from phenotype")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q phenotypeQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q phenotypeQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no phenotypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from phenotype")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o PhenotypeSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o PhenotypeSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no Phenotype slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o PhenotypeSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PhenotypeSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Phenotype slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(phenotypeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), phenotypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"phenotype\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, phenotypePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(phenotypePrimaryKeyColumns), 1, len(phenotypePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from phenotype slice")
	}

	if len(phenotypeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Phenotype) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Phenotype) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Phenotype) ReloadG() error {
	if o == nil {
		return errors.New("chado: no Phenotype provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Phenotype) Reload(exec boil.Executor) error {
	ret, err := FindPhenotype(exec, o.PhenotypeID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PhenotypeSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PhenotypeSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PhenotypeSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty PhenotypeSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PhenotypeSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	phenotypes := PhenotypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), phenotypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"phenotype\".* FROM \"phenotype\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, phenotypePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(phenotypePrimaryKeyColumns), 1, len(phenotypePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&phenotypes)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in PhenotypeSlice")
	}

	*o = phenotypes

	return nil
}

// PhenotypeExists checks if the Phenotype row exists.
func PhenotypeExists(exec boil.Executor, phenotypeID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"phenotype\" where \"phenotype_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, phenotypeID)
	}

	row := exec.QueryRow(sql, phenotypeID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if phenotype exists")
	}

	return exists, nil
}

// PhenotypeExistsG checks if the Phenotype row exists.
func PhenotypeExistsG(phenotypeID int) (bool, error) {
	return PhenotypeExists(boil.GetDB(), phenotypeID)
}

// PhenotypeExistsGP checks if the Phenotype row exists. Panics on error.
func PhenotypeExistsGP(phenotypeID int) bool {
	e, err := PhenotypeExists(boil.GetDB(), phenotypeID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// PhenotypeExistsP checks if the Phenotype row exists. Panics on error.
func PhenotypeExistsP(exec boil.Executor, phenotypeID int) bool {
	e, err := PhenotypeExists(exec, phenotypeID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
