package accounting_servicev1

import (
	context "context"
	fmt "fmt"

	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm "github.com/jinzhu/gorm"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

type MergeLinkedAccountTokenORM struct {
	AccessToken                      string
	AccountingIntegrationMergeLinkId *uint64
	HrisIntegrationMergeLinkId       *uint64
	Id                               uint64 `gorm:"unique_index:idx_merge_linked_account_tokens_id"`
	ItemId                           string
	KeyId                            string
	MergeEndUserOriginId             string
	MergeIntegrationSlug             string
	Version                          string
}

// TableName overrides the default tablename generated by GORM
func (MergeLinkedAccountTokenORM) TableName() string {
	return "merge_linked_account_tokens"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *MergeLinkedAccountToken) ToORM(ctx context.Context) (MergeLinkedAccountTokenORM, error) {
	to := MergeLinkedAccountTokenORM{}
	var err error
	if prehook, ok := interface{}(m).(MergeLinkedAccountTokenWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.ItemId = m.ItemId
	to.KeyId = m.KeyId
	to.AccessToken = m.AccessToken
	to.Version = m.Version
	to.MergeEndUserOriginId = m.MergeEndUserOriginId
	to.MergeIntegrationSlug = m.MergeIntegrationSlug
	if posthook, ok := interface{}(m).(MergeLinkedAccountTokenWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *MergeLinkedAccountTokenORM) ToPB(ctx context.Context) (MergeLinkedAccountToken, error) {
	to := MergeLinkedAccountToken{}
	var err error
	if prehook, ok := interface{}(m).(MergeLinkedAccountTokenWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.ItemId = m.ItemId
	to.KeyId = m.KeyId
	to.AccessToken = m.AccessToken
	to.Version = m.Version
	to.MergeEndUserOriginId = m.MergeEndUserOriginId
	to.MergeIntegrationSlug = m.MergeIntegrationSlug
	if posthook, ok := interface{}(m).(MergeLinkedAccountTokenWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type MergeLinkedAccountToken the arg will be the target, the caller the one being converted from

// MergeLinkedAccountTokenBeforeToORM called before default ToORM code
type MergeLinkedAccountTokenWithBeforeToORM interface {
	BeforeToORM(context.Context, *MergeLinkedAccountTokenORM) error
}

// MergeLinkedAccountTokenAfterToORM called after default ToORM code
type MergeLinkedAccountTokenWithAfterToORM interface {
	AfterToORM(context.Context, *MergeLinkedAccountTokenORM) error
}

// MergeLinkedAccountTokenBeforeToPB called before default ToPB code
type MergeLinkedAccountTokenWithBeforeToPB interface {
	BeforeToPB(context.Context, *MergeLinkedAccountToken) error
}

// MergeLinkedAccountTokenAfterToPB called after default ToPB code
type MergeLinkedAccountTokenWithAfterToPB interface {
	AfterToPB(context.Context, *MergeLinkedAccountToken) error
}

// DefaultCreateMergeLinkedAccountToken executes a basic gorm create call
func DefaultCreateMergeLinkedAccountToken(ctx context.Context, in *MergeLinkedAccountToken, db *gorm.DB) (*MergeLinkedAccountToken, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(MergeLinkedAccountTokenORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(MergeLinkedAccountTokenORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type MergeLinkedAccountTokenORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type MergeLinkedAccountTokenORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadMergeLinkedAccountToken(ctx context.Context, in *MergeLinkedAccountToken, db *gorm.DB) (*MergeLinkedAccountToken, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(MergeLinkedAccountTokenORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &MergeLinkedAccountTokenORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(MergeLinkedAccountTokenORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := MergeLinkedAccountTokenORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(MergeLinkedAccountTokenORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type MergeLinkedAccountTokenORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type MergeLinkedAccountTokenORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type MergeLinkedAccountTokenORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteMergeLinkedAccountToken(ctx context.Context, in *MergeLinkedAccountToken, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(MergeLinkedAccountTokenORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&MergeLinkedAccountTokenORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(MergeLinkedAccountTokenORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type MergeLinkedAccountTokenORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type MergeLinkedAccountTokenORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteMergeLinkedAccountTokenSet(ctx context.Context, in []*MergeLinkedAccountToken, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []uint64{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == 0 {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&MergeLinkedAccountTokenORM{})).(MergeLinkedAccountTokenORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&MergeLinkedAccountTokenORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&MergeLinkedAccountTokenORM{})).(MergeLinkedAccountTokenORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type MergeLinkedAccountTokenORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*MergeLinkedAccountToken, *gorm.DB) (*gorm.DB, error)
}
type MergeLinkedAccountTokenORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*MergeLinkedAccountToken, *gorm.DB) error
}

// DefaultStrictUpdateMergeLinkedAccountToken clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateMergeLinkedAccountToken(ctx context.Context, in *MergeLinkedAccountToken, db *gorm.DB) (*MergeLinkedAccountToken, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateMergeLinkedAccountToken")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &MergeLinkedAccountTokenORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(MergeLinkedAccountTokenORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(MergeLinkedAccountTokenORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(MergeLinkedAccountTokenORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type MergeLinkedAccountTokenORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type MergeLinkedAccountTokenORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type MergeLinkedAccountTokenORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchMergeLinkedAccountToken executes a basic gorm update call with patch behavior
func DefaultPatchMergeLinkedAccountToken(ctx context.Context, in *MergeLinkedAccountToken, updateMask *field_mask.FieldMask, db *gorm.DB) (*MergeLinkedAccountToken, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj MergeLinkedAccountToken
	var err error
	if hook, ok := interface{}(&pbObj).(MergeLinkedAccountTokenWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadMergeLinkedAccountToken(ctx, &MergeLinkedAccountToken{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(MergeLinkedAccountTokenWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskMergeLinkedAccountToken(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(MergeLinkedAccountTokenWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateMergeLinkedAccountToken(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(MergeLinkedAccountTokenWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type MergeLinkedAccountTokenWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *MergeLinkedAccountToken, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type MergeLinkedAccountTokenWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *MergeLinkedAccountToken, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type MergeLinkedAccountTokenWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *MergeLinkedAccountToken, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type MergeLinkedAccountTokenWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *MergeLinkedAccountToken, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetMergeLinkedAccountToken executes a bulk gorm update call with patch behavior
func DefaultPatchSetMergeLinkedAccountToken(ctx context.Context, objects []*MergeLinkedAccountToken, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*MergeLinkedAccountToken, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*MergeLinkedAccountToken, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchMergeLinkedAccountToken(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskMergeLinkedAccountToken patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskMergeLinkedAccountToken(ctx context.Context, patchee *MergeLinkedAccountToken, patcher *MergeLinkedAccountToken, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*MergeLinkedAccountToken, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"ItemId" {
			patchee.ItemId = patcher.ItemId
			continue
		}
		if f == prefix+"KeyId" {
			patchee.KeyId = patcher.KeyId
			continue
		}
		if f == prefix+"AccessToken" {
			patchee.AccessToken = patcher.AccessToken
			continue
		}
		if f == prefix+"Version" {
			patchee.Version = patcher.Version
			continue
		}
		if f == prefix+"MergeEndUserOriginId" {
			patchee.MergeEndUserOriginId = patcher.MergeEndUserOriginId
			continue
		}
		if f == prefix+"MergeIntegrationSlug" {
			patchee.MergeIntegrationSlug = patcher.MergeIntegrationSlug
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListMergeLinkedAccountToken executes a gorm list call
func DefaultListMergeLinkedAccountToken(ctx context.Context, db *gorm.DB) ([]*MergeLinkedAccountToken, error) {
	in := MergeLinkedAccountToken{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(MergeLinkedAccountTokenORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &MergeLinkedAccountTokenORM{}, &MergeLinkedAccountToken{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(MergeLinkedAccountTokenORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []MergeLinkedAccountTokenORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(MergeLinkedAccountTokenORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*MergeLinkedAccountToken{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type MergeLinkedAccountTokenORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type MergeLinkedAccountTokenORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type MergeLinkedAccountTokenORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]MergeLinkedAccountTokenORM) error
}
