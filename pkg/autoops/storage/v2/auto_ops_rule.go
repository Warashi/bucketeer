// Copyright 2024 The Bucketeer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate mockgen -source=$GOFILE -package=mock -destination=./mock/$GOFILE
package v2

import (
	"context"
	_ "embed"
	"errors"
	"fmt"

	"github.com/bucketeer-io/bucketeer/pkg/autoops/domain"
	"github.com/bucketeer-io/bucketeer/pkg/storage/v2/mysql"
	proto "github.com/bucketeer-io/bucketeer/proto/autoops"
)

var (
	//go:embed sql/auto_ops_rule/insert_auto_ops_rule.sql
	insertAutoOpsRuleSQL string
	//go:embed sql/auto_ops_rule/update_auto_ops_rule.sql
	updateAutoOpsRuleSQL string
	//go:embed sql/auto_ops_rule/select_auto_ops_rule.sql
	selectAutoOpsRuleSQL string
	//go:embed sql/auto_ops_rule/select_auto_ops_rules.sql
	selectAutoOpsRulesSQL string
)

var (
	ErrAutoOpsRuleAlreadyExists          = errors.New("autoOpsRule: already exists")
	ErrAutoOpsRuleNotFound               = errors.New("autoOpsRule: not found")
	ErrAutoOpsRuleUnexpectedAffectedRows = errors.New("autoOpsRule: unexpected affected rows")
)

type AutoOpsRuleStorage interface {
	CreateAutoOpsRule(ctx context.Context, e *domain.AutoOpsRule, environmentId string) error
	UpdateAutoOpsRule(ctx context.Context, e *domain.AutoOpsRule, environmentId string) error
	GetAutoOpsRule(ctx context.Context, id, environmentId string) (*domain.AutoOpsRule, error)
	ListAutoOpsRules(
		ctx context.Context,
		options *mysql.ListOptions,
	) ([]*proto.AutoOpsRule, int, error)
}

type autoOpsRuleStorage struct {
	qe mysql.QueryExecer
}

func NewAutoOpsRuleStorage(qe mysql.QueryExecer) AutoOpsRuleStorage {
	return &autoOpsRuleStorage{qe: qe}
}

func (s *autoOpsRuleStorage) CreateAutoOpsRule(
	ctx context.Context,
	e *domain.AutoOpsRule,
	environmentId string,
) error {
	_, err := s.qe.ExecContext(
		ctx,
		insertAutoOpsRuleSQL,
		e.Id,
		e.FeatureId,
		int32(e.OpsType),
		mysql.JSONObject{Val: e.Clauses},
		e.CreatedAt,
		e.UpdatedAt,
		e.Deleted,
		int32(e.AutoOpsStatus),
		environmentId,
	)
	if err != nil {
		if err == mysql.ErrDuplicateEntry {
			return ErrAutoOpsRuleAlreadyExists
		}
		return err
	}
	return nil
}

func (s *autoOpsRuleStorage) UpdateAutoOpsRule(
	ctx context.Context,
	e *domain.AutoOpsRule,
	environmentId string,
) error {
	result, err := s.qe.ExecContext(
		ctx,
		updateAutoOpsRuleSQL,
		e.FeatureId,
		int32(e.OpsType),
		mysql.JSONObject{Val: e.Clauses},
		e.CreatedAt,
		e.UpdatedAt,
		e.Deleted,
		int32(e.AutoOpsStatus),
		e.Id,
		environmentId,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return ErrAutoOpsRuleUnexpectedAffectedRows
	}
	return nil
}

func (s *autoOpsRuleStorage) GetAutoOpsRule(
	ctx context.Context,
	id, environmentId string,
) (*domain.AutoOpsRule, error) {
	autoOpsRule := proto.AutoOpsRule{}
	var opsType int32
	err := s.qe.QueryRowContext(
		ctx,
		selectAutoOpsRuleSQL,
		id,
		environmentId,
	).Scan(
		&autoOpsRule.Id,
		&autoOpsRule.FeatureId,
		&opsType,
		&mysql.JSONObject{Val: &autoOpsRule.Clauses},
		&autoOpsRule.CreatedAt,
		&autoOpsRule.UpdatedAt,
		&autoOpsRule.Deleted,
		&autoOpsRule.AutoOpsStatus,
	)
	if err != nil {
		if err == mysql.ErrNoRows {
			return nil, ErrAutoOpsRuleNotFound
		}
		return nil, err
	}
	autoOpsRule.OpsType = proto.OpsType(opsType)
	return &domain.AutoOpsRule{AutoOpsRule: &autoOpsRule}, nil
}

func (s *autoOpsRuleStorage) ListAutoOpsRules(
	ctx context.Context,
	options *mysql.ListOptions,
) ([]*proto.AutoOpsRule, int, error) {
	var whereParts []mysql.WherePart = []mysql.WherePart{}
	var orderBySQL string = ""
	var limitOffsetSQL string = ""
	var limit int = 0
	var offset int = 0
	if options != nil {
		whereParts = options.CreateWhereParts()
		orderBySQL = mysql.ConstructOrderBySQLString(options.Orders)
		limitOffsetSQL = mysql.ConstructLimitOffsetSQLString(options.Limit, options.Offset)
		limit = options.Limit
		offset = options.Offset
	}

	whereSQL, whereArgs := mysql.ConstructWhereSQLString(whereParts)
	query := fmt.Sprintf(selectAutoOpsRulesSQL, whereSQL, orderBySQL, limitOffsetSQL)
	rows, err := s.qe.QueryContext(ctx, query, whereArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	autoOpsRules := make([]*proto.AutoOpsRule, 0, limit)
	for rows.Next() {
		autoOpsRule := proto.AutoOpsRule{}
		var opsType int32
		err := rows.Scan(
			&autoOpsRule.Id,
			&autoOpsRule.FeatureId,
			&opsType,
			&mysql.JSONObject{Val: &autoOpsRule.Clauses},
			&autoOpsRule.CreatedAt,
			&autoOpsRule.UpdatedAt,
			&autoOpsRule.Deleted,
			&autoOpsRule.AutoOpsStatus,
		)
		if err != nil {
			return nil, 0, err
		}
		autoOpsRule.OpsType = proto.OpsType(opsType)
		autoOpsRules = append(autoOpsRules, &autoOpsRule)
	}
	if rows.Err() != nil {
		return nil, 0, err
	}
	nextOffset := offset + len(autoOpsRules)
	return autoOpsRules, nextOffset, nil
}
