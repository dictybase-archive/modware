package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testAuthUserRoles(t *testing.T) {
	t.Parallel()

	query := AuthUserRoles(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testAuthUserRolesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserRole := &AuthUserRole{}
	if err = randomize.Struct(seed, authUserRole, authUserRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = authUserRole.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUserRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthUserRolesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserRole := &AuthUserRole{}
	if err = randomize.Struct(seed, authUserRole, authUserRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AuthUserRoles(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := AuthUserRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthUserRolesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserRole := &AuthUserRole{}
	if err = randomize.Struct(seed, authUserRole, authUserRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserRole.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AuthUserRoleSlice{authUserRole}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUserRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testAuthUserRolesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserRole := &AuthUserRole{}
	if err = randomize.Struct(seed, authUserRole, authUserRoleDBTypes, true, authUserRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserRole.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := AuthUserRoleExists(tx, authUserRole.AuthUserRoleID)
	if err != nil {
		t.Errorf("Unable to check if AuthUserRole exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AuthUserRoleExistsG to return true, but got false.")
	}
}
func testAuthUserRolesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserRole := &AuthUserRole{}
	if err = randomize.Struct(seed, authUserRole, authUserRoleDBTypes, true, authUserRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserRole.Insert(tx); err != nil {
		t.Error(err)
	}

	authUserRoleFound, err := FindAuthUserRole(tx, authUserRole.AuthUserRoleID)
	if err != nil {
		t.Error(err)
	}

	if authUserRoleFound == nil {
		t.Error("want a record, got nil")
	}
}
func testAuthUserRolesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserRole := &AuthUserRole{}
	if err = randomize.Struct(seed, authUserRole, authUserRoleDBTypes, true, authUserRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AuthUserRoles(tx).Bind(authUserRole); err != nil {
		t.Error(err)
	}
}

func testAuthUserRolesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserRole := &AuthUserRole{}
	if err = randomize.Struct(seed, authUserRole, authUserRoleDBTypes, true, authUserRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := AuthUserRoles(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAuthUserRolesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserRoleOne := &AuthUserRole{}
	authUserRoleTwo := &AuthUserRole{}
	if err = randomize.Struct(seed, authUserRoleOne, authUserRoleDBTypes, false, authUserRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}
	if err = randomize.Struct(seed, authUserRoleTwo, authUserRoleDBTypes, false, authUserRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserRoleOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = authUserRoleTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AuthUserRoles(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAuthUserRolesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	authUserRoleOne := &AuthUserRole{}
	authUserRoleTwo := &AuthUserRole{}
	if err = randomize.Struct(seed, authUserRoleOne, authUserRoleDBTypes, false, authUserRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}
	if err = randomize.Struct(seed, authUserRoleTwo, authUserRoleDBTypes, false, authUserRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserRoleOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = authUserRoleTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUserRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func authUserRoleBeforeInsertHook(e boil.Executor, o *AuthUserRole) error {
	*o = AuthUserRole{}
	return nil
}

func authUserRoleAfterInsertHook(e boil.Executor, o *AuthUserRole) error {
	*o = AuthUserRole{}
	return nil
}

func authUserRoleAfterSelectHook(e boil.Executor, o *AuthUserRole) error {
	*o = AuthUserRole{}
	return nil
}

func authUserRoleBeforeUpdateHook(e boil.Executor, o *AuthUserRole) error {
	*o = AuthUserRole{}
	return nil
}

func authUserRoleAfterUpdateHook(e boil.Executor, o *AuthUserRole) error {
	*o = AuthUserRole{}
	return nil
}

func authUserRoleBeforeDeleteHook(e boil.Executor, o *AuthUserRole) error {
	*o = AuthUserRole{}
	return nil
}

func authUserRoleAfterDeleteHook(e boil.Executor, o *AuthUserRole) error {
	*o = AuthUserRole{}
	return nil
}

func authUserRoleBeforeUpsertHook(e boil.Executor, o *AuthUserRole) error {
	*o = AuthUserRole{}
	return nil
}

func authUserRoleAfterUpsertHook(e boil.Executor, o *AuthUserRole) error {
	*o = AuthUserRole{}
	return nil
}

func testAuthUserRolesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &AuthUserRole{}
	o := &AuthUserRole{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, authUserRoleDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AuthUserRole object: %s", err)
	}

	AddAuthUserRoleHook(boil.BeforeInsertHook, authUserRoleBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	authUserRoleBeforeInsertHooks = []AuthUserRoleHook{}

	AddAuthUserRoleHook(boil.AfterInsertHook, authUserRoleAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	authUserRoleAfterInsertHooks = []AuthUserRoleHook{}

	AddAuthUserRoleHook(boil.AfterSelectHook, authUserRoleAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	authUserRoleAfterSelectHooks = []AuthUserRoleHook{}

	AddAuthUserRoleHook(boil.BeforeUpdateHook, authUserRoleBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	authUserRoleBeforeUpdateHooks = []AuthUserRoleHook{}

	AddAuthUserRoleHook(boil.AfterUpdateHook, authUserRoleAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	authUserRoleAfterUpdateHooks = []AuthUserRoleHook{}

	AddAuthUserRoleHook(boil.BeforeDeleteHook, authUserRoleBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	authUserRoleBeforeDeleteHooks = []AuthUserRoleHook{}

	AddAuthUserRoleHook(boil.AfterDeleteHook, authUserRoleAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	authUserRoleAfterDeleteHooks = []AuthUserRoleHook{}

	AddAuthUserRoleHook(boil.BeforeUpsertHook, authUserRoleBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	authUserRoleBeforeUpsertHooks = []AuthUserRoleHook{}

	AddAuthUserRoleHook(boil.AfterUpsertHook, authUserRoleAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	authUserRoleAfterUpsertHooks = []AuthUserRoleHook{}
}
func testAuthUserRolesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserRole := &AuthUserRole{}
	if err = randomize.Struct(seed, authUserRole, authUserRoleDBTypes, true, authUserRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserRole.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUserRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthUserRolesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserRole := &AuthUserRole{}
	if err = randomize.Struct(seed, authUserRole, authUserRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserRole.Insert(tx, authUserRoleColumns...); err != nil {
		t.Error(err)
	}

	count, err := AuthUserRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthUserRoleToOneAuthUserUsingAuthUser(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local AuthUserRole
	var foreign AuthUser

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, authUserRoleDBTypes, true, authUserRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.AuthUserID = foreign.AuthUserID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.AuthUser(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.AuthUserID != foreign.AuthUserID {
		t.Errorf("want: %v, got %v", foreign.AuthUserID, check.AuthUserID)
	}

	slice := AuthUserRoleSlice{&local}
	if err = local.L.LoadAuthUser(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.AuthUser == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.AuthUser = nil
	if err = local.L.LoadAuthUser(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.AuthUser == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAuthUserRoleToOneAuthRoleUsingAuthRole(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local AuthUserRole
	var foreign AuthRole

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, authUserRoleDBTypes, true, authUserRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
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

	slice := AuthUserRoleSlice{&local}
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

func testAuthUserRoleToOneSetOpAuthUserUsingAuthUser(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AuthUserRole
	var b, c AuthUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authUserRoleDBTypes, false, strmangle.SetComplement(authUserRolePrimaryKeyColumns, authUserRoleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, authUserDBTypes, false, strmangle.SetComplement(authUserPrimaryKeyColumns, authUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, authUserDBTypes, false, strmangle.SetComplement(authUserPrimaryKeyColumns, authUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*AuthUser{&b, &c} {
		err = a.SetAuthUser(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.AuthUser != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.AuthUserRoles[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AuthUserID != x.AuthUserID {
			t.Error("foreign key was wrong value", a.AuthUserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AuthUserID))
		reflect.Indirect(reflect.ValueOf(&a.AuthUserID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AuthUserID != x.AuthUserID {
			t.Error("foreign key was wrong value", a.AuthUserID, x.AuthUserID)
		}
	}
}
func testAuthUserRoleToOneSetOpAuthRoleUsingAuthRole(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AuthUserRole
	var b, c AuthRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authUserRoleDBTypes, false, strmangle.SetComplement(authUserRolePrimaryKeyColumns, authUserRoleColumnsWithoutDefault)...); err != nil {
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

		if x.R.AuthUserRoles[0] != &a {
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
func testAuthUserRolesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserRole := &AuthUserRole{}
	if err = randomize.Struct(seed, authUserRole, authUserRoleDBTypes, true, authUserRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = authUserRole.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testAuthUserRolesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserRole := &AuthUserRole{}
	if err = randomize.Struct(seed, authUserRole, authUserRoleDBTypes, true, authUserRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserRole.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AuthUserRoleSlice{authUserRole}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testAuthUserRolesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserRole := &AuthUserRole{}
	if err = randomize.Struct(seed, authUserRole, authUserRoleDBTypes, true, authUserRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserRole.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AuthUserRoles(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	authUserRoleDBTypes = map[string]string{"AuthRoleID": "integer", "AuthUserID": "integer", "AuthUserRoleID": "integer"}
	_                   = bytes.MinRead
)

func testAuthUserRolesUpdate(t *testing.T) {
	t.Parallel()

	if len(authUserRoleColumns) == len(authUserRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	authUserRole := &AuthUserRole{}
	if err = randomize.Struct(seed, authUserRole, authUserRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserRole.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUserRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, authUserRole, authUserRoleDBTypes, true, authUserRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}

	if err = authUserRole.Update(tx); err != nil {
		t.Error(err)
	}
}

func testAuthUserRolesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(authUserRoleColumns) == len(authUserRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	authUserRole := &AuthUserRole{}
	if err = randomize.Struct(seed, authUserRole, authUserRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserRole.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUserRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, authUserRole, authUserRoleDBTypes, true, authUserRolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(authUserRoleColumns, authUserRolePrimaryKeyColumns) {
		fields = authUserRoleColumns
	} else {
		fields = strmangle.SetComplement(
			authUserRoleColumns,
			authUserRolePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(authUserRole))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := AuthUserRoleSlice{authUserRole}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testAuthUserRolesUpsert(t *testing.T) {
	t.Parallel()

	if len(authUserRoleColumns) == len(authUserRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	authUserRole := AuthUserRole{}
	if err = randomize.Struct(seed, &authUserRole, authUserRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserRole.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert AuthUserRole: %s", err)
	}

	count, err := AuthUserRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &authUserRole, authUserRoleDBTypes, false, authUserRolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthUserRole struct: %s", err)
	}

	if err = authUserRole.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert AuthUserRole: %s", err)
	}

	count, err = AuthUserRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
