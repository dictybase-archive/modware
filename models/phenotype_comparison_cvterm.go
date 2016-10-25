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

// PhenotypeComparisonCvterm is an object representing the database table.
type PhenotypeComparisonCvterm struct {
	PhenotypeComparisonCvtermID int `boil:"phenotype_comparison_cvterm_id" json:"phenotype_comparison_cvterm_id" toml:"phenotype_comparison_cvterm_id" yaml:"phenotype_comparison_cvterm_id"`
	PhenotypeComparisonID       int `boil:"phenotype_comparison_id" json:"phenotype_comparison_id" toml:"phenotype_comparison_id" yaml:"phenotype_comparison_id"`
	CvtermID                    int `boil:"cvterm_id" json:"cvterm_id" toml:"cvterm_id" yaml:"cvterm_id"`
	PubID                       int `boil:"pub_id" json:"pub_id" toml:"pub_id" yaml:"pub_id"`
	Rank                        int `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *phenotypeComparisonCvtermR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L phenotypeComparisonCvtermL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// phenotypeComparisonCvtermR is where relationships are stored.
type phenotypeComparisonCvtermR struct {
	PhenotypeComparison *PhenotypeComparison
	Cvterm              *Cvterm
	Pub                 *Pub
}

// phenotypeComparisonCvtermL is where Load methods for each relationship are stored.
type phenotypeComparisonCvtermL struct{}

var (
	phenotypeComparisonCvtermColumns               = []string{"phenotype_comparison_cvterm_id", "phenotype_comparison_id", "cvterm_id", "pub_id", "rank"}
	phenotypeComparisonCvtermColumnsWithoutDefault = []string{"phenotype_comparison_id", "cvterm_id", "pub_id"}
	phenotypeComparisonCvtermColumnsWithDefault    = []string{"phenotype_comparison_cvterm_id", "rank"}
	phenotypeComparisonCvtermPrimaryKeyColumns     = []string{"phenotype_comparison_cvterm_id"}
)

type (
	// PhenotypeComparisonCvtermSlice is an alias for a slice of pointers to PhenotypeComparisonCvterm.
	// This should generally be used opposed to []PhenotypeComparisonCvterm.
	PhenotypeComparisonCvtermSlice []*PhenotypeComparisonCvterm
	// PhenotypeComparisonCvtermHook is the signature for custom PhenotypeComparisonCvterm hook methods
	PhenotypeComparisonCvtermHook func(boil.Executor, *PhenotypeComparisonCvterm) error

	phenotypeComparisonCvtermQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	phenotypeComparisonCvtermType                 = reflect.TypeOf(&PhenotypeComparisonCvterm{})
	phenotypeComparisonCvtermMapping              = queries.MakeStructMapping(phenotypeComparisonCvtermType)
	phenotypeComparisonCvtermPrimaryKeyMapping, _ = queries.BindMapping(phenotypeComparisonCvtermType, phenotypeComparisonCvtermMapping, phenotypeComparisonCvtermPrimaryKeyColumns)
	phenotypeComparisonCvtermInsertCacheMut       sync.RWMutex
	phenotypeComparisonCvtermInsertCache          = make(map[string]insertCache)
	phenotypeComparisonCvtermUpdateCacheMut       sync.RWMutex
	phenotypeComparisonCvtermUpdateCache          = make(map[string]updateCache)
	phenotypeComparisonCvtermUpsertCacheMut       sync.RWMutex
	phenotypeComparisonCvtermUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var phenotypeComparisonCvtermBeforeInsertHooks []PhenotypeComparisonCvtermHook
var phenotypeComparisonCvtermBeforeUpdateHooks []PhenotypeComparisonCvtermHook
var phenotypeComparisonCvtermBeforeDeleteHooks []PhenotypeComparisonCvtermHook
var phenotypeComparisonCvtermBeforeUpsertHooks []PhenotypeComparisonCvtermHook

var phenotypeComparisonCvtermAfterInsertHooks []PhenotypeComparisonCvtermHook
var phenotypeComparisonCvtermAfterSelectHooks []PhenotypeComparisonCvtermHook
var phenotypeComparisonCvtermAfterUpdateHooks []PhenotypeComparisonCvtermHook
var phenotypeComparisonCvtermAfterDeleteHooks []PhenotypeComparisonCvtermHook
var phenotypeComparisonCvtermAfterUpsertHooks []PhenotypeComparisonCvtermHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PhenotypeComparisonCvterm) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeComparisonCvtermBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PhenotypeComparisonCvterm) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeComparisonCvtermBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PhenotypeComparisonCvterm) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeComparisonCvtermBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PhenotypeComparisonCvterm) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeComparisonCvtermBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PhenotypeComparisonCvterm) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeComparisonCvtermAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PhenotypeComparisonCvterm) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeComparisonCvtermAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PhenotypeComparisonCvterm) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeComparisonCvtermAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PhenotypeComparisonCvterm) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeComparisonCvtermAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PhenotypeComparisonCvterm) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeComparisonCvtermAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPhenotypeComparisonCvtermHook registers your hook function for all future operations.
func AddPhenotypeComparisonCvtermHook(hookPoint boil.HookPoint, phenotypeComparisonCvtermHook PhenotypeComparisonCvtermHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		phenotypeComparisonCvtermBeforeInsertHooks = append(phenotypeComparisonCvtermBeforeInsertHooks, phenotypeComparisonCvtermHook)
	case boil.BeforeUpdateHook:
		phenotypeComparisonCvtermBeforeUpdateHooks = append(phenotypeComparisonCvtermBeforeUpdateHooks, phenotypeComparisonCvtermHook)
	case boil.BeforeDeleteHook:
		phenotypeComparisonCvtermBeforeDeleteHooks = append(phenotypeComparisonCvtermBeforeDeleteHooks, phenotypeComparisonCvtermHook)
	case boil.BeforeUpsertHook:
		phenotypeComparisonCvtermBeforeUpsertHooks = append(phenotypeComparisonCvtermBeforeUpsertHooks, phenotypeComparisonCvtermHook)
	case boil.AfterInsertHook:
		phenotypeComparisonCvtermAfterInsertHooks = append(phenotypeComparisonCvtermAfterInsertHooks, phenotypeComparisonCvtermHook)
	case boil.AfterSelectHook:
		phenotypeComparisonCvtermAfterSelectHooks = append(phenotypeComparisonCvtermAfterSelectHooks, phenotypeComparisonCvtermHook)
	case boil.AfterUpdateHook:
		phenotypeComparisonCvtermAfterUpdateHooks = append(phenotypeComparisonCvtermAfterUpdateHooks, phenotypeComparisonCvtermHook)
	case boil.AfterDeleteHook:
		phenotypeComparisonCvtermAfterDeleteHooks = append(phenotypeComparisonCvtermAfterDeleteHooks, phenotypeComparisonCvtermHook)
	case boil.AfterUpsertHook:
		phenotypeComparisonCvtermAfterUpsertHooks = append(phenotypeComparisonCvtermAfterUpsertHooks, phenotypeComparisonCvtermHook)
	}
}

// OneP returns a single phenotypeComparisonCvterm record from the query, and panics on error.
func (q phenotypeComparisonCvtermQuery) OneP() *PhenotypeComparisonCvterm {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single phenotypeComparisonCvterm record from the query.
func (q phenotypeComparisonCvtermQuery) One() (*PhenotypeComparisonCvterm, error) {
	o := &PhenotypeComparisonCvterm{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for phenotype_comparison_cvterm")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all PhenotypeComparisonCvterm records from the query, and panics on error.
func (q phenotypeComparisonCvtermQuery) AllP() PhenotypeComparisonCvtermSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all PhenotypeComparisonCvterm records from the query.
func (q phenotypeComparisonCvtermQuery) All() (PhenotypeComparisonCvtermSlice, error) {
	var o PhenotypeComparisonCvtermSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PhenotypeComparisonCvterm slice")
	}

	if len(phenotypeComparisonCvtermAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all PhenotypeComparisonCvterm records in the query, and panics on error.
func (q phenotypeComparisonCvtermQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all PhenotypeComparisonCvterm records in the query.
func (q phenotypeComparisonCvtermQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count phenotype_comparison_cvterm rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q phenotypeComparisonCvtermQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q phenotypeComparisonCvtermQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if phenotype_comparison_cvterm exists")
	}

	return count > 0, nil
}

// PhenotypeComparisonG pointed to by the foreign key.
func (o *PhenotypeComparisonCvterm) PhenotypeComparisonG(mods ...qm.QueryMod) phenotypeComparisonQuery {
	return o.PhenotypeComparison(boil.GetDB(), mods...)
}

// PhenotypeComparison pointed to by the foreign key.
func (o *PhenotypeComparisonCvterm) PhenotypeComparison(exec boil.Executor, mods ...qm.QueryMod) phenotypeComparisonQuery {
	queryMods := []qm.QueryMod{
		qm.Where("phenotype_comparison_id=$1", o.PhenotypeComparisonID),
	}

	queryMods = append(queryMods, mods...)

	query := PhenotypeComparisons(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype_comparison\"")

	return query
}

// CvtermG pointed to by the foreign key.
func (o *PhenotypeComparisonCvterm) CvtermG(mods ...qm.QueryMod) cvtermQuery {
	return o.Cvterm(boil.GetDB(), mods...)
}

// Cvterm pointed to by the foreign key.
func (o *PhenotypeComparisonCvterm) Cvterm(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// PubG pointed to by the foreign key.
func (o *PhenotypeComparisonCvterm) PubG(mods ...qm.QueryMod) pubQuery {
	return o.Pub(boil.GetDB(), mods...)
}

// Pub pointed to by the foreign key.
func (o *PhenotypeComparisonCvterm) Pub(exec boil.Executor, mods ...qm.QueryMod) pubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub\"")

	return query
}

// LoadPhenotypeComparison allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeComparisonCvtermL) LoadPhenotypeComparison(e boil.Executor, singular bool, maybePhenotypeComparisonCvterm interface{}) error {
	var slice []*PhenotypeComparisonCvterm
	var object *PhenotypeComparisonCvterm

	count := 1
	if singular {
		object = maybePhenotypeComparisonCvterm.(*PhenotypeComparisonCvterm)
	} else {
		slice = *maybePhenotypeComparisonCvterm.(*PhenotypeComparisonCvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeComparisonCvtermR{}
		args[0] = object.PhenotypeComparisonID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeComparisonCvtermR{}
			args[i] = obj.PhenotypeComparisonID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype_comparison\" where \"phenotype_comparison_id\" in (%s)",
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

	if len(phenotypeComparisonCvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.PhenotypeComparison = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PhenotypeComparisonID == foreign.PhenotypeComparisonID {
				local.R.PhenotypeComparison = foreign
				break
			}
		}
	}

	return nil
}

// LoadCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeComparisonCvtermL) LoadCvterm(e boil.Executor, singular bool, maybePhenotypeComparisonCvterm interface{}) error {
	var slice []*PhenotypeComparisonCvterm
	var object *PhenotypeComparisonCvterm

	count := 1
	if singular {
		object = maybePhenotypeComparisonCvterm.(*PhenotypeComparisonCvterm)
	} else {
		slice = *maybePhenotypeComparisonCvterm.(*PhenotypeComparisonCvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeComparisonCvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeComparisonCvtermR{}
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

	if len(phenotypeComparisonCvtermAfterSelectHooks) != 0 {
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
func (phenotypeComparisonCvtermL) LoadPub(e boil.Executor, singular bool, maybePhenotypeComparisonCvterm interface{}) error {
	var slice []*PhenotypeComparisonCvterm
	var object *PhenotypeComparisonCvterm

	count := 1
	if singular {
		object = maybePhenotypeComparisonCvterm.(*PhenotypeComparisonCvterm)
	} else {
		slice = *maybePhenotypeComparisonCvterm.(*PhenotypeComparisonCvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeComparisonCvtermR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeComparisonCvtermR{}
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

	if len(phenotypeComparisonCvtermAfterSelectHooks) != 0 {
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

// SetPhenotypeComparison of the phenotype_comparison_cvterm to the related item.
// Sets o.R.PhenotypeComparison to related.
// Adds o to related.R.PhenotypeComparisonCvterm.
func (o *PhenotypeComparisonCvterm) SetPhenotypeComparison(exec boil.Executor, insert bool, related *PhenotypeComparison) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenotype_comparison_cvterm\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"phenotype_comparison_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenotypeComparisonCvtermPrimaryKeyColumns),
	)
	values := []interface{}{related.PhenotypeComparisonID, o.PhenotypeComparisonCvtermID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PhenotypeComparisonID = related.PhenotypeComparisonID

	if o.R == nil {
		o.R = &phenotypeComparisonCvtermR{
			PhenotypeComparison: related,
		}
	} else {
		o.R.PhenotypeComparison = related
	}

	if related.R == nil {
		related.R = &phenotypeComparisonR{
			PhenotypeComparisonCvterm: o,
		}
	} else {
		related.R.PhenotypeComparisonCvterm = o
	}

	return nil
}

// SetCvterm of the phenotype_comparison_cvterm to the related item.
// Sets o.R.Cvterm to related.
// Adds o to related.R.PhenotypeComparisonCvterm.
func (o *PhenotypeComparisonCvterm) SetCvterm(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenotype_comparison_cvterm\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"cvterm_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenotypeComparisonCvtermPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.PhenotypeComparisonCvtermID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CvtermID = related.CvtermID

	if o.R == nil {
		o.R = &phenotypeComparisonCvtermR{
			Cvterm: related,
		}
	} else {
		o.R.Cvterm = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			PhenotypeComparisonCvterm: o,
		}
	} else {
		related.R.PhenotypeComparisonCvterm = o
	}

	return nil
}

// SetPub of the phenotype_comparison_cvterm to the related item.
// Sets o.R.Pub to related.
// Adds o to related.R.PhenotypeComparisonCvterms.
func (o *PhenotypeComparisonCvterm) SetPub(exec boil.Executor, insert bool, related *Pub) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenotype_comparison_cvterm\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenotypeComparisonCvtermPrimaryKeyColumns),
	)
	values := []interface{}{related.PubID, o.PhenotypeComparisonCvtermID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PubID = related.PubID

	if o.R == nil {
		o.R = &phenotypeComparisonCvtermR{
			Pub: related,
		}
	} else {
		o.R.Pub = related
	}

	if related.R == nil {
		related.R = &pubR{
			PhenotypeComparisonCvterms: PhenotypeComparisonCvtermSlice{o},
		}
	} else {
		related.R.PhenotypeComparisonCvterms = append(related.R.PhenotypeComparisonCvterms, o)
	}

	return nil
}

// PhenotypeComparisonCvtermsG retrieves all records.
func PhenotypeComparisonCvtermsG(mods ...qm.QueryMod) phenotypeComparisonCvtermQuery {
	return PhenotypeComparisonCvterms(boil.GetDB(), mods...)
}

// PhenotypeComparisonCvterms retrieves all the records using an executor.
func PhenotypeComparisonCvterms(exec boil.Executor, mods ...qm.QueryMod) phenotypeComparisonCvtermQuery {
	mods = append(mods, qm.From("\"phenotype_comparison_cvterm\""))
	return phenotypeComparisonCvtermQuery{NewQuery(exec, mods...)}
}

// FindPhenotypeComparisonCvtermG retrieves a single record by ID.
func FindPhenotypeComparisonCvtermG(phenotypeComparisonCvtermID int, selectCols ...string) (*PhenotypeComparisonCvterm, error) {
	return FindPhenotypeComparisonCvterm(boil.GetDB(), phenotypeComparisonCvtermID, selectCols...)
}

// FindPhenotypeComparisonCvtermGP retrieves a single record by ID, and panics on error.
func FindPhenotypeComparisonCvtermGP(phenotypeComparisonCvtermID int, selectCols ...string) *PhenotypeComparisonCvterm {
	retobj, err := FindPhenotypeComparisonCvterm(boil.GetDB(), phenotypeComparisonCvtermID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindPhenotypeComparisonCvterm retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPhenotypeComparisonCvterm(exec boil.Executor, phenotypeComparisonCvtermID int, selectCols ...string) (*PhenotypeComparisonCvterm, error) {
	phenotypeComparisonCvtermObj := &PhenotypeComparisonCvterm{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"phenotype_comparison_cvterm\" where \"phenotype_comparison_cvterm_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, phenotypeComparisonCvtermID)

	err := q.Bind(phenotypeComparisonCvtermObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from phenotype_comparison_cvterm")
	}

	return phenotypeComparisonCvtermObj, nil
}

// FindPhenotypeComparisonCvtermP retrieves a single record by ID with an executor, and panics on error.
func FindPhenotypeComparisonCvtermP(exec boil.Executor, phenotypeComparisonCvtermID int, selectCols ...string) *PhenotypeComparisonCvterm {
	retobj, err := FindPhenotypeComparisonCvterm(exec, phenotypeComparisonCvtermID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *PhenotypeComparisonCvterm) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *PhenotypeComparisonCvterm) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *PhenotypeComparisonCvterm) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *PhenotypeComparisonCvterm) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no phenotype_comparison_cvterm provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(phenotypeComparisonCvtermColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	phenotypeComparisonCvtermInsertCacheMut.RLock()
	cache, cached := phenotypeComparisonCvtermInsertCache[key]
	phenotypeComparisonCvtermInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			phenotypeComparisonCvtermColumns,
			phenotypeComparisonCvtermColumnsWithDefault,
			phenotypeComparisonCvtermColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(phenotypeComparisonCvtermType, phenotypeComparisonCvtermMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(phenotypeComparisonCvtermType, phenotypeComparisonCvtermMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"phenotype_comparison_cvterm\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into phenotype_comparison_cvterm")
	}

	if !cached {
		phenotypeComparisonCvtermInsertCacheMut.Lock()
		phenotypeComparisonCvtermInsertCache[key] = cache
		phenotypeComparisonCvtermInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single PhenotypeComparisonCvterm record. See Update for
// whitelist behavior description.
func (o *PhenotypeComparisonCvterm) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single PhenotypeComparisonCvterm record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *PhenotypeComparisonCvterm) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the PhenotypeComparisonCvterm, and panics on error.
// See Update for whitelist behavior description.
func (o *PhenotypeComparisonCvterm) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the PhenotypeComparisonCvterm.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *PhenotypeComparisonCvterm) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	phenotypeComparisonCvtermUpdateCacheMut.RLock()
	cache, cached := phenotypeComparisonCvtermUpdateCache[key]
	phenotypeComparisonCvtermUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(phenotypeComparisonCvtermColumns, phenotypeComparisonCvtermPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update phenotype_comparison_cvterm, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"phenotype_comparison_cvterm\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, phenotypeComparisonCvtermPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(phenotypeComparisonCvtermType, phenotypeComparisonCvtermMapping, append(wl, phenotypeComparisonCvtermPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update phenotype_comparison_cvterm row")
	}

	if !cached {
		phenotypeComparisonCvtermUpdateCacheMut.Lock()
		phenotypeComparisonCvtermUpdateCache[key] = cache
		phenotypeComparisonCvtermUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q phenotypeComparisonCvtermQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q phenotypeComparisonCvtermQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for phenotype_comparison_cvterm")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o PhenotypeComparisonCvtermSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o PhenotypeComparisonCvtermSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o PhenotypeComparisonCvtermSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PhenotypeComparisonCvtermSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), phenotypeComparisonCvtermPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"phenotype_comparison_cvterm\" SET %s WHERE (\"phenotype_comparison_cvterm_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(phenotypeComparisonCvtermPrimaryKeyColumns), len(colNames)+1, len(phenotypeComparisonCvtermPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in phenotypeComparisonCvterm slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *PhenotypeComparisonCvterm) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *PhenotypeComparisonCvterm) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *PhenotypeComparisonCvterm) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *PhenotypeComparisonCvterm) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no phenotype_comparison_cvterm provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(phenotypeComparisonCvtermColumnsWithDefault, o)

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

	phenotypeComparisonCvtermUpsertCacheMut.RLock()
	cache, cached := phenotypeComparisonCvtermUpsertCache[key]
	phenotypeComparisonCvtermUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			phenotypeComparisonCvtermColumns,
			phenotypeComparisonCvtermColumnsWithDefault,
			phenotypeComparisonCvtermColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			phenotypeComparisonCvtermColumns,
			phenotypeComparisonCvtermPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert phenotype_comparison_cvterm, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(phenotypeComparisonCvtermPrimaryKeyColumns))
			copy(conflict, phenotypeComparisonCvtermPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"phenotype_comparison_cvterm\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(phenotypeComparisonCvtermType, phenotypeComparisonCvtermMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(phenotypeComparisonCvtermType, phenotypeComparisonCvtermMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for phenotype_comparison_cvterm")
	}

	if !cached {
		phenotypeComparisonCvtermUpsertCacheMut.Lock()
		phenotypeComparisonCvtermUpsertCache[key] = cache
		phenotypeComparisonCvtermUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single PhenotypeComparisonCvterm record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *PhenotypeComparisonCvterm) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single PhenotypeComparisonCvterm record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *PhenotypeComparisonCvterm) DeleteG() error {
	if o == nil {
		return errors.New("models: no PhenotypeComparisonCvterm provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single PhenotypeComparisonCvterm record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *PhenotypeComparisonCvterm) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single PhenotypeComparisonCvterm record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PhenotypeComparisonCvterm) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no PhenotypeComparisonCvterm provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), phenotypeComparisonCvtermPrimaryKeyMapping)
	sql := "DELETE FROM \"phenotype_comparison_cvterm\" WHERE \"phenotype_comparison_cvterm_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from phenotype_comparison_cvterm")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q phenotypeComparisonCvtermQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q phenotypeComparisonCvtermQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no phenotypeComparisonCvtermQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from phenotype_comparison_cvterm")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o PhenotypeComparisonCvtermSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o PhenotypeComparisonCvtermSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no PhenotypeComparisonCvterm slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o PhenotypeComparisonCvtermSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PhenotypeComparisonCvtermSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no PhenotypeComparisonCvterm slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(phenotypeComparisonCvtermBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), phenotypeComparisonCvtermPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"phenotype_comparison_cvterm\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, phenotypeComparisonCvtermPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(phenotypeComparisonCvtermPrimaryKeyColumns), 1, len(phenotypeComparisonCvtermPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from phenotypeComparisonCvterm slice")
	}

	if len(phenotypeComparisonCvtermAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *PhenotypeComparisonCvterm) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *PhenotypeComparisonCvterm) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *PhenotypeComparisonCvterm) ReloadG() error {
	if o == nil {
		return errors.New("models: no PhenotypeComparisonCvterm provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PhenotypeComparisonCvterm) Reload(exec boil.Executor) error {
	ret, err := FindPhenotypeComparisonCvterm(exec, o.PhenotypeComparisonCvtermID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PhenotypeComparisonCvtermSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PhenotypeComparisonCvtermSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PhenotypeComparisonCvtermSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty PhenotypeComparisonCvtermSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PhenotypeComparisonCvtermSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	phenotypeComparisonCvterms := PhenotypeComparisonCvtermSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), phenotypeComparisonCvtermPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"phenotype_comparison_cvterm\".* FROM \"phenotype_comparison_cvterm\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, phenotypeComparisonCvtermPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(phenotypeComparisonCvtermPrimaryKeyColumns), 1, len(phenotypeComparisonCvtermPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&phenotypeComparisonCvterms)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PhenotypeComparisonCvtermSlice")
	}

	*o = phenotypeComparisonCvterms

	return nil
}

// PhenotypeComparisonCvtermExists checks if the PhenotypeComparisonCvterm row exists.
func PhenotypeComparisonCvtermExists(exec boil.Executor, phenotypeComparisonCvtermID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"phenotype_comparison_cvterm\" where \"phenotype_comparison_cvterm_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, phenotypeComparisonCvtermID)
	}

	row := exec.QueryRow(sql, phenotypeComparisonCvtermID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if phenotype_comparison_cvterm exists")
	}

	return exists, nil
}

// PhenotypeComparisonCvtermExistsG checks if the PhenotypeComparisonCvterm row exists.
func PhenotypeComparisonCvtermExistsG(phenotypeComparisonCvtermID int) (bool, error) {
	return PhenotypeComparisonCvtermExists(boil.GetDB(), phenotypeComparisonCvtermID)
}

// PhenotypeComparisonCvtermExistsGP checks if the PhenotypeComparisonCvterm row exists. Panics on error.
func PhenotypeComparisonCvtermExistsGP(phenotypeComparisonCvtermID int) bool {
	e, err := PhenotypeComparisonCvtermExists(boil.GetDB(), phenotypeComparisonCvtermID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// PhenotypeComparisonCvtermExistsP checks if the PhenotypeComparisonCvterm row exists. Panics on error.
func PhenotypeComparisonCvtermExistsP(exec boil.Executor, phenotypeComparisonCvtermID int) bool {
	e, err := PhenotypeComparisonCvtermExists(exec, phenotypeComparisonCvtermID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
