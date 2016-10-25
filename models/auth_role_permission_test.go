package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testAuthRolePermissions(t *testing.T) {
	t.Parallel()

	query := AuthRolePermissions(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testAuthRolePermissionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRolePermission := &AuthRolePermission{}
	if err = randomize.Struct(seed, authRolePermission, authRolePermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRolePermission.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = authRolePermission.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthRolePermissions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthRolePermissionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRolePermission := &AuthRolePermission{}
	if err = randomize.Struct(seed, authRolePermission, authRolePermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRolePermission.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AuthRolePermissions(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := AuthRolePermissions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthRolePermissionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRolePermission := &AuthRolePermission{}
	if err = randomize.Struct(seed, authRolePermission, authRolePermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRolePermission.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AuthRolePermissionSlice{authRolePermission}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthRolePermissions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testAuthRolePermissionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRolePermission := &AuthRolePermission{}
	if err = randomize.Struct(seed, authRolePermission, authRolePermissionDBTypes, true, authRolePermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRolePermission.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := AuthRolePermissionExists(tx, authRolePermission.AuthRolePermissionID)
	if err != nil {
		t.Errorf("Unable to check if AuthRolePermission exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AuthRolePermissionExistsG to return true, but got false.")
	}
}
func testAuthRolePermissionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRolePermission := &AuthRolePermission{}
	if err = randomize.Struct(seed, authRolePermission, authRolePermissionDBTypes, true, authRolePermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRolePermission.Insert(tx); err != nil {
		t.Error(err)
	}

	authRolePermissionFound, err := FindAuthRolePermission(tx, authRolePermission.AuthRolePermissionID)
	if err != nil {
		t.Error(err)
	}

	if authRolePermissionFound == nil {
		t.Error("want a record, got nil")
	}
}
func testAuthRolePermissionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRolePermission := &AuthRolePermission{}
	if err = randomize.Struct(seed, authRolePermission, authRolePermissionDBTypes, true, authRolePermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRolePermission.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AuthRolePermissions(tx).Bind(authRolePermission); err != nil {
		t.Error(err)
	}
}

func testAuthRolePermissionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRolePermission := &AuthRolePermission{}
	if err = randomize.Struct(seed, authRolePermission, authRolePermissionDBTypes, true, authRolePermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRolePermission.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := AuthRolePermissions(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAuthRolePermissionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRolePermissionOne := &AuthRolePermission{}
	authRolePermissionTwo := &AuthRolePermission{}
	if err = randomize.Struct(seed, authRolePermissionOne, authRolePermissionDBTypes, false, authRolePermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}
	if err = randomize.Struct(seed, authRolePermissionTwo, authRolePermissionDBTypes, false, authRolePermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRolePermissionOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = authRolePermissionTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AuthRolePermissions(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAuthRolePermissionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	authRolePermissionOne := &AuthRolePermission{}
	authRolePermissionTwo := &AuthRolePermission{}
	if err = randomize.Struct(seed, authRolePermissionOne, authRolePermissionDBTypes, false, authRolePermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}
	if err = randomize.Struct(seed, authRolePermissionTwo, authRolePermissionDBTypes, false, authRolePermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRolePermissionOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = authRolePermissionTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthRolePermissions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func authRolePermissionBeforeInsertHook(e boil.Executor, o *AuthRolePermission) error {
	*o = AuthRolePermission{}
	return nil
}

func authRolePermissionAfterInsertHook(e boil.Executor, o *AuthRolePermission) error {
	*o = AuthRolePermission{}
	return nil
}

func authRolePermissionAfterSelectHook(e boil.Executor, o *AuthRolePermission) error {
	*o = AuthRolePermission{}
	return nil
}

func authRolePermissionBeforeUpdateHook(e boil.Executor, o *AuthRolePermission) error {
	*o = AuthRolePermission{}
	return nil
}

func authRolePermissionAfterUpdateHook(e boil.Executor, o *AuthRolePermission) error {
	*o = AuthRolePermission{}
	return nil
}

func authRolePermissionBeforeDeleteHook(e boil.Executor, o *AuthRolePermission) error {
	*o = AuthRolePermission{}
	return nil
}

func authRolePermissionAfterDeleteHook(e boil.Executor, o *AuthRolePermission) error {
	*o = AuthRolePermission{}
	return nil
}

func authRolePermissionBeforeUpsertHook(e boil.Executor, o *AuthRolePermission) error {
	*o = AuthRolePermission{}
	return nil
}

func authRolePermissionAfterUpsertHook(e boil.Executor, o *AuthRolePermission) error {
	*o = AuthRolePermission{}
	return nil
}

func testAuthRolePermissionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &AuthRolePermission{}
	o := &AuthRolePermission{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, authRolePermissionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission object: %s", err)
	}

	AddAuthRolePermissionHook(boil.BeforeInsertHook, authRolePermissionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	authRolePermissionBeforeInsertHooks = []AuthRolePermissionHook{}

	AddAuthRolePermissionHook(boil.AfterInsertHook, authRolePermissionAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	authRolePermissionAfterInsertHooks = []AuthRolePermissionHook{}

	AddAuthRolePermissionHook(boil.AfterSelectHook, authRolePermissionAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	authRolePermissionAfterSelectHooks = []AuthRolePermissionHook{}

	AddAuthRolePermissionHook(boil.BeforeUpdateHook, authRolePermissionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	authRolePermissionBeforeUpdateHooks = []AuthRolePermissionHook{}

	AddAuthRolePermissionHook(boil.AfterUpdateHook, authRolePermissionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	authRolePermissionAfterUpdateHooks = []AuthRolePermissionHook{}

	AddAuthRolePermissionHook(boil.BeforeDeleteHook, authRolePermissionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	authRolePermissionBeforeDeleteHooks = []AuthRolePermissionHook{}

	AddAuthRolePermissionHook(boil.AfterDeleteHook, authRolePermissionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	authRolePermissionAfterDeleteHooks = []AuthRolePermissionHook{}

	AddAuthRolePermissionHook(boil.BeforeUpsertHook, authRolePermissionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	authRolePermissionBeforeUpsertHooks = []AuthRolePermissionHook{}

	AddAuthRolePermissionHook(boil.AfterUpsertHook, authRolePermissionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	authRolePermissionAfterUpsertHooks = []AuthRolePermissionHook{}
}
func testAuthRolePermissionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRolePermission := &AuthRolePermission{}
	if err = randomize.Struct(seed, authRolePermission, authRolePermissionDBTypes, true, authRolePermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRolePermission.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthRolePermissions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthRolePermissionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRolePermission := &AuthRolePermission{}
	if err = randomize.Struct(seed, authRolePermission, authRolePermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRolePermission.Insert(tx, authRolePermissionColumns...); err != nil {
		t.Error(err)
	}

	count, err := AuthRolePermissions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthRolePermissionToOneAuthRoleUsingAuthRole(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local AuthRolePermission
	var foreign AuthRole

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, authRolePermissionDBTypes, true, authRolePermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, authRoleDBTypes, true, authRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRole struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.AuthRoleID = foreign.AuthRoleID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.AuthRole(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.AuthRoleID != foreign.AuthRoleID {
		t.Errorf("want: %v, got %v", foreign.AuthRoleID, check.AuthRoleID)
	}

	slice := AuthRolePermissionSlice{&local}
	if err = local.L.LoadAuthRole(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.AuthRole == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.AuthRole = nil
	if err = local.L.LoadAuthRole(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.AuthRole == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAuthRolePermissionToOneAuthPermissionUsingAuthPermission(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local AuthRolePermission
	var foreign AuthPermission

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, authRolePermissionDBTypes, true, authRolePermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.AuthPermissionID = foreign.AuthPermissionID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.AuthPermission(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.AuthPermissionID != foreign.AuthPermissionID {
		t.Errorf("want: %v, got %v", foreign.AuthPermissionID, check.AuthPermissionID)
	}

	slice := AuthRolePermissionSlice{&local}
	if err = local.L.LoadAuthPermission(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.AuthPermission == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.AuthPermission = nil
	if err = local.L.LoadAuthPermission(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.AuthPermission == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAuthRolePermissionToOneSetOpAuthRoleUsingAuthRole(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AuthRolePermission
	var b, c AuthRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authRolePermissionDBTypes, false, strmangle.SetComplement(authRolePermissionPrimaryKeyColumns, authRolePermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, authRoleDBTypes, false, strmangle.SetComplement(authRolePrimaryKeyColumns, authRoleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, authRoleDBTypes, false, strmangle.SetComplement(authRolePrimaryKeyColumns, authRoleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*AuthRole{&b, &c} {
		err = a.SetAuthRole(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.AuthRole != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.AuthRolePermission != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AuthRoleID != x.AuthRoleID {
			t.Error("foreign key was wrong value", a.AuthRoleID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AuthRoleID))
		reflect.Indirect(reflect.ValueOf(&a.AuthRoleID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AuthRoleID != x.AuthRoleID {
			t.Error("foreign key was wrong value", a.AuthRoleID, x.AuthRoleID)
		}
	}
}
func testAuthRolePermissionToOneSetOpAuthPermissionUsingAuthPermission(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AuthRolePermission
	var b, c AuthPermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authRolePermissionDBTypes, false, strmangle.SetComplement(authRolePermissionPrimaryKeyColumns, authRolePermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, authPermissionDBTypes, false, strmangle.SetComplement(authPermissionPrimaryKeyColumns, authPermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, authPermissionDBTypes, false, strmangle.SetComplement(authPermissionPrimaryKeyColumns, authPermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*AuthPermission{&b, &c} {
		err = a.SetAuthPermission(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.AuthPermission != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.AuthRolePermission != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AuthPermissionID != x.AuthPermissionID {
			t.Error("foreign key was wrong value", a.AuthPermissionID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AuthPermissionID))
		reflect.Indirect(reflect.ValueOf(&a.AuthPermissionID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AuthPermissionID != x.AuthPermissionID {
			t.Error("foreign key was wrong value", a.AuthPermissionID, x.AuthPermissionID)
		}
	}
}
func testAuthRolePermissionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRolePermission := &AuthRolePermission{}
	if err = randomize.Struct(seed, authRolePermission, authRolePermissionDBTypes, true, authRolePermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRolePermission.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = authRolePermission.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testAuthRolePermissionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRolePermission := &AuthRolePermission{}
	if err = randomize.Struct(seed, authRolePermission, authRolePermissionDBTypes, true, authRolePermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRolePermission.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AuthRolePermissionSlice{authRolePermission}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testAuthRolePermissionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authRolePermission := &AuthRolePermission{}
	if err = randomize.Struct(seed, authRolePermission, authRolePermissionDBTypes, true, authRolePermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRolePermission.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AuthRolePermissions(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	authRolePermissionDBTypes = map[string]string{"AuthPermissionID": "integer", "AuthRoleID": "integer", "AuthRolePermissionID": "integer", "CreatedAt": "timestamp with time zone", "UpdatedAt": "timestamp with time zone"}
	_                         = bytes.MinRead
)

func testAuthRolePermissionsUpdate(t *testing.T) {
	t.Parallel()

	if len(authRolePermissionColumns) == len(authRolePermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	authRolePermission := &AuthRolePermission{}
	if err = randomize.Struct(seed, authRolePermission, authRolePermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRolePermission.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthRolePermissions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, authRolePermission, authRolePermissionDBTypes, true, authRolePermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}

	if err = authRolePermission.Update(tx); err != nil {
		t.Error(err)
	}
}

func testAuthRolePermissionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(authRolePermissionColumns) == len(authRolePermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	authRolePermission := &AuthRolePermission{}
	if err = randomize.Struct(seed, authRolePermission, authRolePermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRolePermission.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthRolePermissions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, authRolePermission, authRolePermissionDBTypes, true, authRolePermissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(authRolePermissionColumns, authRolePermissionPrimaryKeyColumns) {
		fields = authRolePermissionColumns
	} else {
		fields = strmangle.SetComplement(
			authRolePermissionColumns,
			authRolePermissionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(authRolePermission))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := AuthRolePermissionSlice{authRolePermission}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testAuthRolePermissionsUpsert(t *testing.T) {
	t.Parallel()

	if len(authRolePermissionColumns) == len(authRolePermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	authRolePermission := AuthRolePermission{}
	if err = randomize.Struct(seed, &authRolePermission, authRolePermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authRolePermission.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert AuthRolePermission: %s", err)
	}

	count, err := AuthRolePermissions(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &authRolePermission, authRolePermissionDBTypes, false, authRolePermissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthRolePermission struct: %s", err)
	}

	if err = authRolePermission.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert AuthRolePermission: %s", err)
	}

	count, err = AuthRolePermissions(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
