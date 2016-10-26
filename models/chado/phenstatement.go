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

// Phenstatement is an object representing the database table.
type Phenstatement struct {
	PhenstatementID int `boil:"phenstatement_id" json:"phenstatement_id" toml:"phenstatement_id" yaml:"phenstatement_id"`
	GenotypeID      int `boil:"genotype_id" json:"genotype_id" toml:"genotype_id" yaml:"genotype_id"`
	EnvironmentID   int `boil:"environment_id" json:"environment_id" toml:"environment_id" yaml:"environment_id"`
	PhenotypeID     int `boil:"phenotype_id" json:"phenotype_id" toml:"phenotype_id" yaml:"phenotype_id"`
	TypeID          int `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	PubID           int `boil:"pub_id" json:"pub_id" toml:"pub_id" yaml:"pub_id"`

	R *phenstatementR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L phenstatementL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// phenstatementR is where relationships are stored.
type phenstatementR struct {
	Environment *Environment
	Genotype    *Genotype
	Phenotype   *Phenotype
	Type        *Cvterm
	Pub         *Pub
}

// phenstatementL is where Load methods for each relationship are stored.
type phenstatementL struct{}

var (
	phenstatementColumns               = []string{"phenstatement_id", "genotype_id", "environment_id", "phenotype_id", "type_id", "pub_id"}
	phenstatementColumnsWithoutDefault = []string{"genotype_id", "environment_id", "phenotype_id", "type_id", "pub_id"}
	phenstatementColumnsWithDefault    = []string{"phenstatement_id"}
	phenstatementPrimaryKeyColumns     = []string{"phenstatement_id"}
)

type (
	// PhenstatementSlice is an alias for a slice of pointers to Phenstatement.
	// This should generally be used opposed to []Phenstatement.
	PhenstatementSlice []*Phenstatement
	// PhenstatementHook is the signature for custom Phenstatement hook methods
	PhenstatementHook func(boil.Executor, *Phenstatement) error

	phenstatementQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	phenstatementType                 = reflect.TypeOf(&Phenstatement{})
	phenstatementMapping              = queries.MakeStructMapping(phenstatementType)
	phenstatementPrimaryKeyMapping, _ = queries.BindMapping(phenstatementType, phenstatementMapping, phenstatementPrimaryKeyColumns)
	phenstatementInsertCacheMut       sync.RWMutex
	phenstatementInsertCache          = make(map[string]insertCache)
	phenstatementUpdateCacheMut       sync.RWMutex
	phenstatementUpdateCache          = make(map[string]updateCache)
	phenstatementUpsertCacheMut       sync.RWMutex
	phenstatementUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var phenstatementBeforeInsertHooks []PhenstatementHook
var phenstatementBeforeUpdateHooks []PhenstatementHook
var phenstatementBeforeDeleteHooks []PhenstatementHook
var phenstatementBeforeUpsertHooks []PhenstatementHook

var phenstatementAfterInsertHooks []PhenstatementHook
var phenstatementAfterSelectHooks []PhenstatementHook
var phenstatementAfterUpdateHooks []PhenstatementHook
var phenstatementAfterDeleteHooks []PhenstatementHook
var phenstatementAfterUpsertHooks []PhenstatementHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Phenstatement) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phenstatementBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Phenstatement) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range phenstatementBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Phenstatement) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range phenstatementBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Phenstatement) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phenstatementBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Phenstatement) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phenstatementAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Phenstatement) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range phenstatementAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Phenstatement) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range phenstatementAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Phenstatement) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range phenstatementAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Phenstatement) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phenstatementAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPhenstatementHook registers your hook function for all future operations.
func AddPhenstatementHook(hookPoint boil.HookPoint, phenstatementHook PhenstatementHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		phenstatementBeforeInsertHooks = append(phenstatementBeforeInsertHooks, phenstatementHook)
	case boil.BeforeUpdateHook:
		phenstatementBeforeUpdateHooks = append(phenstatementBeforeUpdateHooks, phenstatementHook)
	case boil.BeforeDeleteHook:
		phenstatementBeforeDeleteHooks = append(phenstatementBeforeDeleteHooks, phenstatementHook)
	case boil.BeforeUpsertHook:
		phenstatementBeforeUpsertHooks = append(phenstatementBeforeUpsertHooks, phenstatementHook)
	case boil.AfterInsertHook:
		phenstatementAfterInsertHooks = append(phenstatementAfterInsertHooks, phenstatementHook)
	case boil.AfterSelectHook:
		phenstatementAfterSelectHooks = append(phenstatementAfterSelectHooks, phenstatementHook)
	case boil.AfterUpdateHook:
		phenstatementAfterUpdateHooks = append(phenstatementAfterUpdateHooks, phenstatementHook)
	case boil.AfterDeleteHook:
		phenstatementAfterDeleteHooks = append(phenstatementAfterDeleteHooks, phenstatementHook)
	case boil.AfterUpsertHook:
		phenstatementAfterUpsertHooks = append(phenstatementAfterUpsertHooks, phenstatementHook)
	}
}

// OneP returns a single phenstatement record from the query, and panics on error.
func (q phenstatementQuery) OneP() *Phenstatement {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single phenstatement record from the query.
func (q phenstatementQuery) One() (*Phenstatement, error) {
	o := &Phenstatement{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for phenstatement")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Phenstatement records from the query, and panics on error.
func (q phenstatementQuery) AllP() PhenstatementSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Phenstatement records from the query.
func (q phenstatementQuery) All() (PhenstatementSlice, error) {
	var o PhenstatementSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to Phenstatement slice")
	}

	if len(phenstatementAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Phenstatement records in the query, and panics on error.
func (q phenstatementQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Phenstatement records in the query.
func (q phenstatementQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count phenstatement rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q phenstatementQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q phenstatementQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if phenstatement exists")
	}

	return count > 0, nil
}

// EnvironmentG pointed to by the foreign key.
func (o *Phenstatement) EnvironmentG(mods ...qm.QueryMod) environmentQuery {
	return o.Environment(boil.GetDB(), mods...)
}

// Environment pointed to by the foreign key.
func (o *Phenstatement) Environment(exec boil.Executor, mods ...qm.QueryMod) environmentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("environment_id=$1", o.EnvironmentID),
	}

	queryMods = append(queryMods, mods...)

	query := Environments(exec, queryMods...)
	queries.SetFrom(query.Query, "\"environment\"")

	return query
}

// GenotypeG pointed to by the foreign key.
func (o *Phenstatement) GenotypeG(mods ...qm.QueryMod) genotypeQuery {
	return o.Genotype(boil.GetDB(), mods...)
}

// Genotype pointed to by the foreign key.
func (o *Phenstatement) Genotype(exec boil.Executor, mods ...qm.QueryMod) genotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("genotype_id=$1", o.GenotypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Genotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"genotype\"")

	return query
}

// PhenotypeG pointed to by the foreign key.
func (o *Phenstatement) PhenotypeG(mods ...qm.QueryMod) phenotypeQuery {
	return o.Phenotype(boil.GetDB(), mods...)
}

// Phenotype pointed to by the foreign key.
func (o *Phenstatement) Phenotype(exec boil.Executor, mods ...qm.QueryMod) phenotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("phenotype_id=$1", o.PhenotypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Phenotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype\"")

	return query
}

// TypeG pointed to by the foreign key.
func (o *Phenstatement) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Phenstatement) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// PubG pointed to by the foreign key.
func (o *Phenstatement) PubG(mods ...qm.QueryMod) pubQuery {
	return o.Pub(boil.GetDB(), mods...)
}

// Pub pointed to by the foreign key.
func (o *Phenstatement) Pub(exec boil.Executor, mods ...qm.QueryMod) pubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub\"")

	return query
}

// LoadEnvironment allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenstatementL) LoadEnvironment(e boil.Executor, singular bool, maybePhenstatement interface{}) error {
	var slice []*Phenstatement
	var object *Phenstatement

	count := 1
	if singular {
		object = maybePhenstatement.(*Phenstatement)
	} else {
		slice = *maybePhenstatement.(*PhenstatementSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenstatementR{}
		args[0] = object.EnvironmentID
	} else {
		for i, obj := range slice {
			obj.R = &phenstatementR{}
			args[i] = obj.EnvironmentID
		}
	}

	query := fmt.Sprintf(
		"select * from \"environment\" where \"environment_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Environment")
	}
	defer results.Close()

	var resultSlice []*Environment
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Environment")
	}

	if len(phenstatementAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Environment = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.EnvironmentID == foreign.EnvironmentID {
				local.R.Environment = foreign
				break
			}
		}
	}

	return nil
}

// LoadGenotype allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenstatementL) LoadGenotype(e boil.Executor, singular bool, maybePhenstatement interface{}) error {
	var slice []*Phenstatement
	var object *Phenstatement

	count := 1
	if singular {
		object = maybePhenstatement.(*Phenstatement)
	} else {
		slice = *maybePhenstatement.(*PhenstatementSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenstatementR{}
		args[0] = object.GenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &phenstatementR{}
			args[i] = obj.GenotypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"genotype\" where \"genotype_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Genotype")
	}
	defer results.Close()

	var resultSlice []*Genotype
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Genotype")
	}

	if len(phenstatementAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Genotype = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.GenotypeID == foreign.GenotypeID {
				local.R.Genotype = foreign
				break
			}
		}
	}

	return nil
}

// LoadPhenotype allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenstatementL) LoadPhenotype(e boil.Executor, singular bool, maybePhenstatement interface{}) error {
	var slice []*Phenstatement
	var object *Phenstatement

	count := 1
	if singular {
		object = maybePhenstatement.(*Phenstatement)
	} else {
		slice = *maybePhenstatement.(*PhenstatementSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenstatementR{}
		args[0] = object.PhenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &phenstatementR{}
			args[i] = obj.PhenotypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype\" where \"phenotype_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Phenotype")
	}
	defer results.Close()

	var resultSlice []*Phenotype
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Phenotype")
	}

	if len(phenstatementAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Phenotype = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PhenotypeID == foreign.PhenotypeID {
				local.R.Phenotype = foreign
				break
			}
		}
	}

	return nil
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenstatementL) LoadType(e boil.Executor, singular bool, maybePhenstatement interface{}) error {
	var slice []*Phenstatement
	var object *Phenstatement

	count := 1
	if singular {
		object = maybePhenstatement.(*Phenstatement)
	} else {
		slice = *maybePhenstatement.(*PhenstatementSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenstatementR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &phenstatementR{}
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

	if len(phenstatementAfterSelectHooks) != 0 {
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

// LoadPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenstatementL) LoadPub(e boil.Executor, singular bool, maybePhenstatement interface{}) error {
	var slice []*Phenstatement
	var object *Phenstatement

	count := 1
	if singular {
		object = maybePhenstatement.(*Phenstatement)
	} else {
		slice = *maybePhenstatement.(*PhenstatementSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenstatementR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &phenstatementR{}
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

	if len(phenstatementAfterSelectHooks) != 0 {
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

// SetEnvironment of the phenstatement to the related item.
// Sets o.R.Environment to related.
// Adds o to related.R.Phenstatement.
func (o *Phenstatement) SetEnvironment(exec boil.Executor, insert bool, related *Environment) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenstatement\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"environment_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenstatementPrimaryKeyColumns),
	)
	values := []interface{}{related.EnvironmentID, o.PhenstatementID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.EnvironmentID = related.EnvironmentID

	if o.R == nil {
		o.R = &phenstatementR{
			Environment: related,
		}
	} else {
		o.R.Environment = related
	}

	if related.R == nil {
		related.R = &environmentR{
			Phenstatement: o,
		}
	} else {
		related.R.Phenstatement = o
	}

	return nil
}

// SetGenotype of the phenstatement to the related item.
// Sets o.R.Genotype to related.
// Adds o to related.R.Phenstatement.
func (o *Phenstatement) SetGenotype(exec boil.Executor, insert bool, related *Genotype) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenstatement\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"genotype_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenstatementPrimaryKeyColumns),
	)
	values := []interface{}{related.GenotypeID, o.PhenstatementID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.GenotypeID = related.GenotypeID

	if o.R == nil {
		o.R = &phenstatementR{
			Genotype: related,
		}
	} else {
		o.R.Genotype = related
	}

	if related.R == nil {
		related.R = &genotypeR{
			Phenstatement: o,
		}
	} else {
		related.R.Phenstatement = o
	}

	return nil
}

// SetPhenotype of the phenstatement to the related item.
// Sets o.R.Phenotype to related.
// Adds o to related.R.Phenstatement.
func (o *Phenstatement) SetPhenotype(exec boil.Executor, insert bool, related *Phenotype) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenstatement\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"phenotype_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenstatementPrimaryKeyColumns),
	)
	values := []interface{}{related.PhenotypeID, o.PhenstatementID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PhenotypeID = related.PhenotypeID

	if o.R == nil {
		o.R = &phenstatementR{
			Phenotype: related,
		}
	} else {
		o.R.Phenotype = related
	}

	if related.R == nil {
		related.R = &phenotypeR{
			Phenstatement: o,
		}
	} else {
		related.R.Phenstatement = o
	}

	return nil
}

// SetType of the phenstatement to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypePhenstatement.
func (o *Phenstatement) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenstatement\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenstatementPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.PhenstatementID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &phenstatementR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypePhenstatement: o,
		}
	} else {
		related.R.TypePhenstatement = o
	}

	return nil
}

// SetPub of the phenstatement to the related item.
// Sets o.R.Pub to related.
// Adds o to related.R.Phenstatement.
func (o *Phenstatement) SetPub(exec boil.Executor, insert bool, related *Pub) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenstatement\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenstatementPrimaryKeyColumns),
	)
	values := []interface{}{related.PubID, o.PhenstatementID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PubID = related.PubID

	if o.R == nil {
		o.R = &phenstatementR{
			Pub: related,
		}
	} else {
		o.R.Pub = related
	}

	if related.R == nil {
		related.R = &pubR{
			Phenstatement: o,
		}
	} else {
		related.R.Phenstatement = o
	}

	return nil
}

// PhenstatementsG retrieves all records.
func PhenstatementsG(mods ...qm.QueryMod) phenstatementQuery {
	return Phenstatements(boil.GetDB(), mods...)
}

// Phenstatements retrieves all the records using an executor.
func Phenstatements(exec boil.Executor, mods ...qm.QueryMod) phenstatementQuery {
	mods = append(mods, qm.From("\"phenstatement\""))
	return phenstatementQuery{NewQuery(exec, mods...)}
}

// FindPhenstatementG retrieves a single record by ID.
func FindPhenstatementG(phenstatementID int, selectCols ...string) (*Phenstatement, error) {
	return FindPhenstatement(boil.GetDB(), phenstatementID, selectCols...)
}

// FindPhenstatementGP retrieves a single record by ID, and panics on error.
func FindPhenstatementGP(phenstatementID int, selectCols ...string) *Phenstatement {
	retobj, err := FindPhenstatement(boil.GetDB(), phenstatementID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindPhenstatement retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPhenstatement(exec boil.Executor, phenstatementID int, selectCols ...string) (*Phenstatement, error) {
	phenstatementObj := &Phenstatement{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"phenstatement\" where \"phenstatement_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, phenstatementID)

	err := q.Bind(phenstatementObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from phenstatement")
	}

	return phenstatementObj, nil
}

// FindPhenstatementP retrieves a single record by ID with an executor, and panics on error.
func FindPhenstatementP(exec boil.Executor, phenstatementID int, selectCols ...string) *Phenstatement {
	retobj, err := FindPhenstatement(exec, phenstatementID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Phenstatement) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Phenstatement) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Phenstatement) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Phenstatement) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no phenstatement provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(phenstatementColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	phenstatementInsertCacheMut.RLock()
	cache, cached := phenstatementInsertCache[key]
	phenstatementInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			phenstatementColumns,
			phenstatementColumnsWithDefault,
			phenstatementColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(phenstatementType, phenstatementMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(phenstatementType, phenstatementMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"phenstatement\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into phenstatement")
	}

	if !cached {
		phenstatementInsertCacheMut.Lock()
		phenstatementInsertCache[key] = cache
		phenstatementInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Phenstatement record. See Update for
// whitelist behavior description.
func (o *Phenstatement) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Phenstatement record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Phenstatement) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Phenstatement, and panics on error.
// See Update for whitelist behavior description.
func (o *Phenstatement) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Phenstatement.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Phenstatement) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	phenstatementUpdateCacheMut.RLock()
	cache, cached := phenstatementUpdateCache[key]
	phenstatementUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(phenstatementColumns, phenstatementPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update phenstatement, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"phenstatement\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, phenstatementPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(phenstatementType, phenstatementMapping, append(wl, phenstatementPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update phenstatement row")
	}

	if !cached {
		phenstatementUpdateCacheMut.Lock()
		phenstatementUpdateCache[key] = cache
		phenstatementUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q phenstatementQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q phenstatementQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for phenstatement")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o PhenstatementSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o PhenstatementSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o PhenstatementSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PhenstatementSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), phenstatementPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"phenstatement\" SET %s WHERE (\"phenstatement_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(phenstatementPrimaryKeyColumns), len(colNames)+1, len(phenstatementPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in phenstatement slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Phenstatement) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Phenstatement) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Phenstatement) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Phenstatement) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no phenstatement provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(phenstatementColumnsWithDefault, o)

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

	phenstatementUpsertCacheMut.RLock()
	cache, cached := phenstatementUpsertCache[key]
	phenstatementUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			phenstatementColumns,
			phenstatementColumnsWithDefault,
			phenstatementColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			phenstatementColumns,
			phenstatementPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert phenstatement, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(phenstatementPrimaryKeyColumns))
			copy(conflict, phenstatementPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"phenstatement\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(phenstatementType, phenstatementMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(phenstatementType, phenstatementMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for phenstatement")
	}

	if !cached {
		phenstatementUpsertCacheMut.Lock()
		phenstatementUpsertCache[key] = cache
		phenstatementUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Phenstatement record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Phenstatement) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Phenstatement record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Phenstatement) DeleteG() error {
	if o == nil {
		return errors.New("chado: no Phenstatement provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Phenstatement record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Phenstatement) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Phenstatement record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Phenstatement) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Phenstatement provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), phenstatementPrimaryKeyMapping)
	sql := "DELETE FROM \"phenstatement\" WHERE \"phenstatement_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from phenstatement")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q phenstatementQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q phenstatementQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no phenstatementQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from phenstatement")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o PhenstatementSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o PhenstatementSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no Phenstatement slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o PhenstatementSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PhenstatementSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Phenstatement slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(phenstatementBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), phenstatementPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"phenstatement\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, phenstatementPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(phenstatementPrimaryKeyColumns), 1, len(phenstatementPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from phenstatement slice")
	}

	if len(phenstatementAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Phenstatement) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Phenstatement) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Phenstatement) ReloadG() error {
	if o == nil {
		return errors.New("chado: no Phenstatement provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Phenstatement) Reload(exec boil.Executor) error {
	ret, err := FindPhenstatement(exec, o.PhenstatementID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PhenstatementSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PhenstatementSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PhenstatementSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty PhenstatementSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PhenstatementSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	phenstatements := PhenstatementSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), phenstatementPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"phenstatement\".* FROM \"phenstatement\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, phenstatementPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(phenstatementPrimaryKeyColumns), 1, len(phenstatementPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&phenstatements)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in PhenstatementSlice")
	}

	*o = phenstatements

	return nil
}

// PhenstatementExists checks if the Phenstatement row exists.
func PhenstatementExists(exec boil.Executor, phenstatementID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"phenstatement\" where \"phenstatement_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, phenstatementID)
	}

	row := exec.QueryRow(sql, phenstatementID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if phenstatement exists")
	}

	return exists, nil
}

// PhenstatementExistsG checks if the Phenstatement row exists.
func PhenstatementExistsG(phenstatementID int) (bool, error) {
	return PhenstatementExists(boil.GetDB(), phenstatementID)
}

// PhenstatementExistsGP checks if the Phenstatement row exists. Panics on error.
func PhenstatementExistsGP(phenstatementID int) bool {
	e, err := PhenstatementExists(boil.GetDB(), phenstatementID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// PhenstatementExistsP checks if the Phenstatement row exists. Panics on error.
func PhenstatementExistsP(exec boil.Executor, phenstatementID int) bool {
	e, err := PhenstatementExists(exec, phenstatementID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
