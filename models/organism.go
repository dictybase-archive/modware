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

// Organism is an object representing the database table.
type Organism struct {
	OrganismID   int         `boil:"organism_id" json:"organism_id" toml:"organism_id" yaml:"organism_id"`
	Abbreviation null.String `boil:"abbreviation" json:"abbreviation,omitempty" toml:"abbreviation" yaml:"abbreviation,omitempty"`
	Genus        string      `boil:"genus" json:"genus" toml:"genus" yaml:"genus"`
	Species      string      `boil:"species" json:"species" toml:"species" yaml:"species"`
	CommonName   null.String `boil:"common_name" json:"common_name,omitempty" toml:"common_name" yaml:"common_name,omitempty"`
	Comment      null.String `boil:"comment" json:"comment,omitempty" toml:"comment" yaml:"comment,omitempty"`

	R *organismR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L organismL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// organismR is where relationships are stored.
type organismR struct {
	Stock                *Stock
	Feature              *Feature
	JbrowseOrganism      *JbrowseOrganism
	OrganismDbxref       *OrganismDbxref
	Organismprop         *Organismprop
	PhenotypeComparisons PhenotypeComparisonSlice
}

// organismL is where Load methods for each relationship are stored.
type organismL struct{}

var (
	organismColumns               = []string{"organism_id", "abbreviation", "genus", "species", "common_name", "comment"}
	organismColumnsWithoutDefault = []string{"abbreviation", "genus", "species", "common_name", "comment"}
	organismColumnsWithDefault    = []string{"organism_id"}
	organismPrimaryKeyColumns     = []string{"organism_id"}
)

type (
	// OrganismSlice is an alias for a slice of pointers to Organism.
	// This should generally be used opposed to []Organism.
	OrganismSlice []*Organism
	// OrganismHook is the signature for custom Organism hook methods
	OrganismHook func(boil.Executor, *Organism) error

	organismQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	organismType                 = reflect.TypeOf(&Organism{})
	organismMapping              = queries.MakeStructMapping(organismType)
	organismPrimaryKeyMapping, _ = queries.BindMapping(organismType, organismMapping, organismPrimaryKeyColumns)
	organismInsertCacheMut       sync.RWMutex
	organismInsertCache          = make(map[string]insertCache)
	organismUpdateCacheMut       sync.RWMutex
	organismUpdateCache          = make(map[string]updateCache)
	organismUpsertCacheMut       sync.RWMutex
	organismUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var organismBeforeInsertHooks []OrganismHook
var organismBeforeUpdateHooks []OrganismHook
var organismBeforeDeleteHooks []OrganismHook
var organismBeforeUpsertHooks []OrganismHook

var organismAfterInsertHooks []OrganismHook
var organismAfterSelectHooks []OrganismHook
var organismAfterUpdateHooks []OrganismHook
var organismAfterDeleteHooks []OrganismHook
var organismAfterUpsertHooks []OrganismHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Organism) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range organismBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Organism) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range organismBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Organism) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range organismBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Organism) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range organismBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Organism) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range organismAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Organism) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range organismAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Organism) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range organismAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Organism) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range organismAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Organism) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range organismAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOrganismHook registers your hook function for all future operations.
func AddOrganismHook(hookPoint boil.HookPoint, organismHook OrganismHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		organismBeforeInsertHooks = append(organismBeforeInsertHooks, organismHook)
	case boil.BeforeUpdateHook:
		organismBeforeUpdateHooks = append(organismBeforeUpdateHooks, organismHook)
	case boil.BeforeDeleteHook:
		organismBeforeDeleteHooks = append(organismBeforeDeleteHooks, organismHook)
	case boil.BeforeUpsertHook:
		organismBeforeUpsertHooks = append(organismBeforeUpsertHooks, organismHook)
	case boil.AfterInsertHook:
		organismAfterInsertHooks = append(organismAfterInsertHooks, organismHook)
	case boil.AfterSelectHook:
		organismAfterSelectHooks = append(organismAfterSelectHooks, organismHook)
	case boil.AfterUpdateHook:
		organismAfterUpdateHooks = append(organismAfterUpdateHooks, organismHook)
	case boil.AfterDeleteHook:
		organismAfterDeleteHooks = append(organismAfterDeleteHooks, organismHook)
	case boil.AfterUpsertHook:
		organismAfterUpsertHooks = append(organismAfterUpsertHooks, organismHook)
	}
}

// OneP returns a single organism record from the query, and panics on error.
func (q organismQuery) OneP() *Organism {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single organism record from the query.
func (q organismQuery) One() (*Organism, error) {
	o := &Organism{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for organism")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Organism records from the query, and panics on error.
func (q organismQuery) AllP() OrganismSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Organism records from the query.
func (q organismQuery) All() (OrganismSlice, error) {
	var o OrganismSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Organism slice")
	}

	if len(organismAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Organism records in the query, and panics on error.
func (q organismQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Organism records in the query.
func (q organismQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count organism rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q organismQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q organismQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if organism exists")
	}

	return count > 0, nil
}

// StockG pointed to by the foreign key.
func (o *Organism) StockG(mods ...qm.QueryMod) stockQuery {
	return o.Stock(boil.GetDB(), mods...)
}

// Stock pointed to by the foreign key.
func (o *Organism) Stock(exec boil.Executor, mods ...qm.QueryMod) stockQuery {
	queryMods := []qm.QueryMod{
		qm.Where("organism_id=$1", o.OrganismID),
	}

	queryMods = append(queryMods, mods...)

	query := Stocks(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock\"")

	return query
}

// FeatureG pointed to by the foreign key.
func (o *Organism) FeatureG(mods ...qm.QueryMod) featureQuery {
	return o.Feature(boil.GetDB(), mods...)
}

// Feature pointed to by the foreign key.
func (o *Organism) Feature(exec boil.Executor, mods ...qm.QueryMod) featureQuery {
	queryMods := []qm.QueryMod{
		qm.Where("organism_id=$1", o.OrganismID),
	}

	queryMods = append(queryMods, mods...)

	query := Features(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature\"")

	return query
}

// JbrowseOrganismG pointed to by the foreign key.
func (o *Organism) JbrowseOrganismG(mods ...qm.QueryMod) jbrowseOrganismQuery {
	return o.JbrowseOrganism(boil.GetDB(), mods...)
}

// JbrowseOrganism pointed to by the foreign key.
func (o *Organism) JbrowseOrganism(exec boil.Executor, mods ...qm.QueryMod) jbrowseOrganismQuery {
	queryMods := []qm.QueryMod{
		qm.Where("organism_id=$1", o.OrganismID),
	}

	queryMods = append(queryMods, mods...)

	query := JbrowseOrganisms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"jbrowse_organism\"")

	return query
}

// OrganismDbxrefG pointed to by the foreign key.
func (o *Organism) OrganismDbxrefG(mods ...qm.QueryMod) organismDbxrefQuery {
	return o.OrganismDbxref(boil.GetDB(), mods...)
}

// OrganismDbxref pointed to by the foreign key.
func (o *Organism) OrganismDbxref(exec boil.Executor, mods ...qm.QueryMod) organismDbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("organism_id=$1", o.OrganismID),
	}

	queryMods = append(queryMods, mods...)

	query := OrganismDbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"organism_dbxref\"")

	return query
}

// OrganismpropG pointed to by the foreign key.
func (o *Organism) OrganismpropG(mods ...qm.QueryMod) organismpropQuery {
	return o.Organismprop(boil.GetDB(), mods...)
}

// Organismprop pointed to by the foreign key.
func (o *Organism) Organismprop(exec boil.Executor, mods ...qm.QueryMod) organismpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("organism_id=$1", o.OrganismID),
	}

	queryMods = append(queryMods, mods...)

	query := Organismprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"organismprop\"")

	return query
}

// PhenotypeComparisonsG retrieves all the phenotype_comparison's phenotype comparison.
func (o *Organism) PhenotypeComparisonsG(mods ...qm.QueryMod) phenotypeComparisonQuery {
	return o.PhenotypeComparisons(boil.GetDB(), mods...)
}

// PhenotypeComparisons retrieves all the phenotype_comparison's phenotype comparison with an executor.
func (o *Organism) PhenotypeComparisons(exec boil.Executor, mods ...qm.QueryMod) phenotypeComparisonQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"organism_id\"=$1", o.OrganismID),
	)

	query := PhenotypeComparisons(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype_comparison\" as \"a\"")
	return query
}

// LoadStock allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (organismL) LoadStock(e boil.Executor, singular bool, maybeOrganism interface{}) error {
	var slice []*Organism
	var object *Organism

	count := 1
	if singular {
		object = maybeOrganism.(*Organism)
	} else {
		slice = *maybeOrganism.(*OrganismSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &organismR{}
		args[0] = object.OrganismID
	} else {
		for i, obj := range slice {
			obj.R = &organismR{}
			args[i] = obj.OrganismID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock\" where \"organism_id\" in (%s)",
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

	if len(organismAfterSelectHooks) != 0 {
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
			if local.OrganismID == foreign.OrganismID.Int {
				local.R.Stock = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeature allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (organismL) LoadFeature(e boil.Executor, singular bool, maybeOrganism interface{}) error {
	var slice []*Organism
	var object *Organism

	count := 1
	if singular {
		object = maybeOrganism.(*Organism)
	} else {
		slice = *maybeOrganism.(*OrganismSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &organismR{}
		args[0] = object.OrganismID
	} else {
		for i, obj := range slice {
			obj.R = &organismR{}
			args[i] = obj.OrganismID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature\" where \"organism_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Feature")
	}
	defer results.Close()

	var resultSlice []*Feature
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Feature")
	}

	if len(organismAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Feature = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.OrganismID == foreign.OrganismID {
				local.R.Feature = foreign
				break
			}
		}
	}

	return nil
}

// LoadJbrowseOrganism allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (organismL) LoadJbrowseOrganism(e boil.Executor, singular bool, maybeOrganism interface{}) error {
	var slice []*Organism
	var object *Organism

	count := 1
	if singular {
		object = maybeOrganism.(*Organism)
	} else {
		slice = *maybeOrganism.(*OrganismSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &organismR{}
		args[0] = object.OrganismID
	} else {
		for i, obj := range slice {
			obj.R = &organismR{}
			args[i] = obj.OrganismID
		}
	}

	query := fmt.Sprintf(
		"select * from \"jbrowse_organism\" where \"organism_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load JbrowseOrganism")
	}
	defer results.Close()

	var resultSlice []*JbrowseOrganism
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice JbrowseOrganism")
	}

	if len(organismAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.JbrowseOrganism = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.OrganismID == foreign.OrganismID {
				local.R.JbrowseOrganism = foreign
				break
			}
		}
	}

	return nil
}

// LoadOrganismDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (organismL) LoadOrganismDbxref(e boil.Executor, singular bool, maybeOrganism interface{}) error {
	var slice []*Organism
	var object *Organism

	count := 1
	if singular {
		object = maybeOrganism.(*Organism)
	} else {
		slice = *maybeOrganism.(*OrganismSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &organismR{}
		args[0] = object.OrganismID
	} else {
		for i, obj := range slice {
			obj.R = &organismR{}
			args[i] = obj.OrganismID
		}
	}

	query := fmt.Sprintf(
		"select * from \"organism_dbxref\" where \"organism_id\" in (%s)",
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

	if len(organismAfterSelectHooks) != 0 {
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
			if local.OrganismID == foreign.OrganismID {
				local.R.OrganismDbxref = foreign
				break
			}
		}
	}

	return nil
}

// LoadOrganismprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (organismL) LoadOrganismprop(e boil.Executor, singular bool, maybeOrganism interface{}) error {
	var slice []*Organism
	var object *Organism

	count := 1
	if singular {
		object = maybeOrganism.(*Organism)
	} else {
		slice = *maybeOrganism.(*OrganismSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &organismR{}
		args[0] = object.OrganismID
	} else {
		for i, obj := range slice {
			obj.R = &organismR{}
			args[i] = obj.OrganismID
		}
	}

	query := fmt.Sprintf(
		"select * from \"organismprop\" where \"organism_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Organismprop")
	}
	defer results.Close()

	var resultSlice []*Organismprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Organismprop")
	}

	if len(organismAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Organismprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.OrganismID == foreign.OrganismID {
				local.R.Organismprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadPhenotypeComparisons allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (organismL) LoadPhenotypeComparisons(e boil.Executor, singular bool, maybeOrganism interface{}) error {
	var slice []*Organism
	var object *Organism

	count := 1
	if singular {
		object = maybeOrganism.(*Organism)
	} else {
		slice = *maybeOrganism.(*OrganismSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &organismR{}
		args[0] = object.OrganismID
	} else {
		for i, obj := range slice {
			obj.R = &organismR{}
			args[i] = obj.OrganismID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype_comparison\" where \"organism_id\" in (%s)",
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
		object.R.PhenotypeComparisons = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.OrganismID == foreign.OrganismID {
				local.R.PhenotypeComparisons = append(local.R.PhenotypeComparisons, foreign)
				break
			}
		}
	}

	return nil
}

// SetStock of the organism to the related item.
// Sets o.R.Stock to related.
// Adds o to related.R.Organism.
func (o *Organism) SetStock(exec boil.Executor, insert bool, related *Stock) error {
	var err error

	if insert {
		related.OrganismID.Int = o.OrganismID
		related.OrganismID.Valid = true

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"organism_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockPrimaryKeyColumns),
		)
		values := []interface{}{o.OrganismID, related.StockID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.OrganismID.Int = o.OrganismID
		related.OrganismID.Valid = true
	}

	if o.R == nil {
		o.R = &organismR{
			Stock: related,
		}
	} else {
		o.R.Stock = related
	}

	if related.R == nil {
		related.R = &stockR{
			Organism: o,
		}
	} else {
		related.R.Organism = o
	}
	return nil
}

// RemoveStock relationship.
// Sets o.R.Stock to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Organism) RemoveStock(exec boil.Executor, related *Stock) error {
	var err error

	related.OrganismID.Valid = false
	if err = related.Update(exec, "organism_id"); err != nil {
		related.OrganismID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Stock = nil
	if related == nil || related.R == nil {
		return nil
	}

	related.R.Organism = nil
	return nil
}

// SetFeature of the organism to the related item.
// Sets o.R.Feature to related.
// Adds o to related.R.Organism.
func (o *Organism) SetFeature(exec boil.Executor, insert bool, related *Feature) error {
	var err error

	if insert {
		related.OrganismID = o.OrganismID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"organism_id"}),
			strmangle.WhereClause("\"", "\"", 2, featurePrimaryKeyColumns),
		)
		values := []interface{}{o.OrganismID, related.FeatureID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.OrganismID = o.OrganismID

	}

	if o.R == nil {
		o.R = &organismR{
			Feature: related,
		}
	} else {
		o.R.Feature = related
	}

	if related.R == nil {
		related.R = &featureR{
			Organism: o,
		}
	} else {
		related.R.Organism = o
	}
	return nil
}

// SetJbrowseOrganism of the organism to the related item.
// Sets o.R.JbrowseOrganism to related.
// Adds o to related.R.Organism.
func (o *Organism) SetJbrowseOrganism(exec boil.Executor, insert bool, related *JbrowseOrganism) error {
	var err error

	if insert {
		related.OrganismID = o.OrganismID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"jbrowse_organism\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"organism_id"}),
			strmangle.WhereClause("\"", "\"", 2, jbrowseOrganismPrimaryKeyColumns),
		)
		values := []interface{}{o.OrganismID, related.JbrowseOrganismID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.OrganismID = o.OrganismID

	}

	if o.R == nil {
		o.R = &organismR{
			JbrowseOrganism: related,
		}
	} else {
		o.R.JbrowseOrganism = related
	}

	if related.R == nil {
		related.R = &jbrowseOrganismR{
			Organism: o,
		}
	} else {
		related.R.Organism = o
	}
	return nil
}

// SetOrganismDbxref of the organism to the related item.
// Sets o.R.OrganismDbxref to related.
// Adds o to related.R.Organism.
func (o *Organism) SetOrganismDbxref(exec boil.Executor, insert bool, related *OrganismDbxref) error {
	var err error

	if insert {
		related.OrganismID = o.OrganismID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"organism_dbxref\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"organism_id"}),
			strmangle.WhereClause("\"", "\"", 2, organismDbxrefPrimaryKeyColumns),
		)
		values := []interface{}{o.OrganismID, related.OrganismDbxrefID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.OrganismID = o.OrganismID

	}

	if o.R == nil {
		o.R = &organismR{
			OrganismDbxref: related,
		}
	} else {
		o.R.OrganismDbxref = related
	}

	if related.R == nil {
		related.R = &organismDbxrefR{
			Organism: o,
		}
	} else {
		related.R.Organism = o
	}
	return nil
}

// SetOrganismprop of the organism to the related item.
// Sets o.R.Organismprop to related.
// Adds o to related.R.Organism.
func (o *Organism) SetOrganismprop(exec boil.Executor, insert bool, related *Organismprop) error {
	var err error

	if insert {
		related.OrganismID = o.OrganismID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"organismprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"organism_id"}),
			strmangle.WhereClause("\"", "\"", 2, organismpropPrimaryKeyColumns),
		)
		values := []interface{}{o.OrganismID, related.OrganismpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.OrganismID = o.OrganismID

	}

	if o.R == nil {
		o.R = &organismR{
			Organismprop: related,
		}
	} else {
		o.R.Organismprop = related
	}

	if related.R == nil {
		related.R = &organismpropR{
			Organism: o,
		}
	} else {
		related.R.Organism = o
	}
	return nil
}

// AddPhenotypeComparisons adds the given related objects to the existing relationships
// of the organism, optionally inserting them as new records.
// Appends related to o.R.PhenotypeComparisons.
// Sets related.R.Organism appropriately.
func (o *Organism) AddPhenotypeComparisons(exec boil.Executor, insert bool, related ...*PhenotypeComparison) error {
	var err error
	for _, rel := range related {
		rel.OrganismID = o.OrganismID
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "organism_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &organismR{
			PhenotypeComparisons: related,
		}
	} else {
		o.R.PhenotypeComparisons = append(o.R.PhenotypeComparisons, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &phenotypeComparisonR{
				Organism: o,
			}
		} else {
			rel.R.Organism = o
		}
	}
	return nil
}

// OrganismsG retrieves all records.
func OrganismsG(mods ...qm.QueryMod) organismQuery {
	return Organisms(boil.GetDB(), mods...)
}

// Organisms retrieves all the records using an executor.
func Organisms(exec boil.Executor, mods ...qm.QueryMod) organismQuery {
	mods = append(mods, qm.From("\"organism\""))
	return organismQuery{NewQuery(exec, mods...)}
}

// FindOrganismG retrieves a single record by ID.
func FindOrganismG(organismID int, selectCols ...string) (*Organism, error) {
	return FindOrganism(boil.GetDB(), organismID, selectCols...)
}

// FindOrganismGP retrieves a single record by ID, and panics on error.
func FindOrganismGP(organismID int, selectCols ...string) *Organism {
	retobj, err := FindOrganism(boil.GetDB(), organismID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindOrganism retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrganism(exec boil.Executor, organismID int, selectCols ...string) (*Organism, error) {
	organismObj := &Organism{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"organism\" where \"organism_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, organismID)

	err := q.Bind(organismObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from organism")
	}

	return organismObj, nil
}

// FindOrganismP retrieves a single record by ID with an executor, and panics on error.
func FindOrganismP(exec boil.Executor, organismID int, selectCols ...string) *Organism {
	retobj, err := FindOrganism(exec, organismID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Organism) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Organism) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Organism) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Organism) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no organism provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(organismColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	organismInsertCacheMut.RLock()
	cache, cached := organismInsertCache[key]
	organismInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			organismColumns,
			organismColumnsWithDefault,
			organismColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(organismType, organismMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(organismType, organismMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"organism\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into organism")
	}

	if !cached {
		organismInsertCacheMut.Lock()
		organismInsertCache[key] = cache
		organismInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Organism record. See Update for
// whitelist behavior description.
func (o *Organism) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Organism record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Organism) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Organism, and panics on error.
// See Update for whitelist behavior description.
func (o *Organism) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Organism.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Organism) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	organismUpdateCacheMut.RLock()
	cache, cached := organismUpdateCache[key]
	organismUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(organismColumns, organismPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update organism, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"organism\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, organismPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(organismType, organismMapping, append(wl, organismPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update organism row")
	}

	if !cached {
		organismUpdateCacheMut.Lock()
		organismUpdateCache[key] = cache
		organismUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q organismQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q organismQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for organism")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o OrganismSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o OrganismSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o OrganismSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OrganismSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organismPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"organism\" SET %s WHERE (\"organism_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(organismPrimaryKeyColumns), len(colNames)+1, len(organismPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in organism slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Organism) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Organism) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Organism) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Organism) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no organism provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(organismColumnsWithDefault, o)

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

	organismUpsertCacheMut.RLock()
	cache, cached := organismUpsertCache[key]
	organismUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			organismColumns,
			organismColumnsWithDefault,
			organismColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			organismColumns,
			organismPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert organism, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(organismPrimaryKeyColumns))
			copy(conflict, organismPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"organism\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(organismType, organismMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(organismType, organismMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for organism")
	}

	if !cached {
		organismUpsertCacheMut.Lock()
		organismUpsertCache[key] = cache
		organismUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Organism record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Organism) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Organism record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Organism) DeleteG() error {
	if o == nil {
		return errors.New("models: no Organism provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Organism record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Organism) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Organism record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Organism) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Organism provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), organismPrimaryKeyMapping)
	sql := "DELETE FROM \"organism\" WHERE \"organism_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from organism")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q organismQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q organismQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no organismQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from organism")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o OrganismSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o OrganismSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Organism slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o OrganismSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OrganismSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Organism slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(organismBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organismPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"organism\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, organismPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(organismPrimaryKeyColumns), 1, len(organismPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from organism slice")
	}

	if len(organismAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Organism) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Organism) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Organism) ReloadG() error {
	if o == nil {
		return errors.New("models: no Organism provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Organism) Reload(exec boil.Executor) error {
	ret, err := FindOrganism(exec, o.OrganismID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OrganismSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OrganismSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrganismSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty OrganismSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrganismSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	organisms := OrganismSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organismPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"organism\".* FROM \"organism\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, organismPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(organismPrimaryKeyColumns), 1, len(organismPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&organisms)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OrganismSlice")
	}

	*o = organisms

	return nil
}

// OrganismExists checks if the Organism row exists.
func OrganismExists(exec boil.Executor, organismID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"organism\" where \"organism_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, organismID)
	}

	row := exec.QueryRow(sql, organismID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if organism exists")
	}

	return exists, nil
}

// OrganismExistsG checks if the Organism row exists.
func OrganismExistsG(organismID int) (bool, error) {
	return OrganismExists(boil.GetDB(), organismID)
}

// OrganismExistsGP checks if the Organism row exists. Panics on error.
func OrganismExistsGP(organismID int) bool {
	e, err := OrganismExists(boil.GetDB(), organismID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// OrganismExistsP checks if the Organism row exists. Panics on error.
func OrganismExistsP(exec boil.Executor, organismID int) bool {
	e, err := OrganismExists(exec, organismID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
