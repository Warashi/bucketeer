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

package domain

import (
	"errors"
	"time"

	pb "github.com/golang/protobuf/proto" // nolint:staticcheck
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"

	"github.com/bucketeer-io/bucketeer/pkg/uuid"
	proto "github.com/bucketeer-io/bucketeer/proto/autoops"
)

var (
	errClauseNotFound = errors.New("autoOpsRule: clause not found")
	errClauseEmpty    = errors.New("autoOpsRule: clause cannot be empty")

	OpsEventRateClause = &proto.OpsEventRateClause{}
	DatetimeClause     = &proto.DatetimeClause{}
)

type AutoOpsRule struct {
	*proto.AutoOpsRule
}

func NewAutoOpsRule(
	featureID string,
	opsType proto.OpsType,
	opsEventRateClauses []*proto.OpsEventRateClause,
	datetimeClauses []*proto.DatetimeClause,
) (*AutoOpsRule, error) {
	now := time.Now().Unix()
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	autoOpsRule := &AutoOpsRule{&proto.AutoOpsRule{
		Id:            id.String(),
		FeatureId:     featureID,
		OpsType:       opsType,
		Clauses:       []*proto.Clause{},
		CreatedAt:     now,
		UpdatedAt:     now,
		AutoOpsStatus: proto.AutoOpsStatus_WAITING,
		StoppedAt:     0,
	}}
	if opsType == proto.OpsType_EVENT_RATE {
		for _, c := range opsEventRateClauses {
			if _, err := autoOpsRule.AddOpsEventRateClause(c); err != nil {
				return nil, err
			}
		}
	}
	if opsType == proto.OpsType_SCHEDULE {
		for _, c := range datetimeClauses {
			if _, err := autoOpsRule.AddDatetimeClause(c); err != nil {
				return nil, err
			}
		}
	}
	if len(autoOpsRule.Clauses) == 0 {
		return nil, errClauseEmpty
	}
	return autoOpsRule, nil

}

func (a *AutoOpsRule) SetStopped() {
	now := time.Now().Unix()
	a.AutoOpsRule.AutoOpsStatus = proto.AutoOpsStatus_STOPPED
	a.AutoOpsRule.StoppedAt = now
	a.AutoOpsRule.UpdatedAt = now
}

func (a *AutoOpsRule) SetDeleted() {
	a.AutoOpsRule.AutoOpsStatus = proto.AutoOpsStatus_DELETED
	a.AutoOpsRule.UpdatedAt = time.Now().Unix()
}

func (a *AutoOpsRule) SetCompleted() {
	a.AutoOpsRule.AutoOpsStatus = proto.AutoOpsStatus_COMPLETED
	a.AutoOpsRule.UpdatedAt = time.Now().Unix()
}

func (a *AutoOpsRule) HasExecuteClause() bool {
	return a.AutoOpsStatus == proto.AutoOpsStatus_WAITING || a.AutoOpsStatus == proto.AutoOpsStatus_RUNNING
}

func (a *AutoOpsRule) SetOpsType(opsType proto.OpsType) {
	a.AutoOpsRule.OpsType = opsType
	a.AutoOpsRule.UpdatedAt = time.Now().Unix()
	if a.AutoOpsRule.AutoOpsStatus != proto.AutoOpsStatus_RUNNING {
		a.AutoOpsRule.AutoOpsStatus = proto.AutoOpsStatus_WAITING
		a.AutoOpsRule.StoppedAt = 0
	}
}

func (a *AutoOpsRule) SetAutoOpsStatus(status proto.AutoOpsStatus) {
	a.AutoOpsRule.AutoOpsStatus = status
	a.AutoOpsRule.UpdatedAt = time.Now().Unix()
}

func (a *AutoOpsRule) AddOpsEventRateClause(oerc *proto.OpsEventRateClause) (*proto.Clause, error) {
	ac, err := ptypes.MarshalAny(oerc)
	if err != nil {
		return nil, err
	}
	clause, err := a.addClause(ac, oerc.ActionType)
	if err != nil {
		return nil, err
	}
	a.AutoOpsRule.UpdatedAt = time.Now().Unix()
	if a.AutoOpsRule.AutoOpsStatus != proto.AutoOpsStatus_RUNNING {
		a.AutoOpsRule.AutoOpsStatus = proto.AutoOpsStatus_WAITING
		a.AutoOpsRule.StoppedAt = 0
	}
	return clause, nil
}

func (a *AutoOpsRule) AddDatetimeClause(dc *proto.DatetimeClause) (*proto.Clause, error) {
	ac, err := ptypes.MarshalAny(dc)
	if err != nil {
		return nil, err
	}

	var insertIndex = int64(-1)
	for i, c := range a.Clauses {
		datetimeClause, err := a.UnmarshalDatetimeClause(c)
		if err != nil {
			return nil, err
		}
		if datetimeClause == nil {
			continue
		}
		if dc.Time <= datetimeClause.Time {
			insertIndex = int64(i)
			break
		}
	}
	if insertIndex == -1 {
		insertIndex = int64(len(a.Clauses))
	}
	clause, err := a.insertClause(ac, dc.ActionType, insertIndex)
	if err != nil {
		return nil, err
	}
	a.AutoOpsRule.UpdatedAt = time.Now().Unix()
	if a.AutoOpsRule.AutoOpsStatus != proto.AutoOpsStatus_RUNNING {
		a.AutoOpsRule.AutoOpsStatus = proto.AutoOpsStatus_WAITING
		a.AutoOpsRule.StoppedAt = 0
	}
	return clause, nil
}

func (a *AutoOpsRule) addClause(ac *any.Any, actionType proto.ActionType) (*proto.Clause, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	clause := &proto.Clause{
		Id:         id.String(),
		Clause:     ac,
		ActionType: actionType,
	}
	a.AutoOpsRule.Clauses = append(a.AutoOpsRule.Clauses, clause)
	return clause, nil
}

func (a *AutoOpsRule) insertClause(ac *any.Any, actionType proto.ActionType, index int64) (*proto.Clause, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	clause := &proto.Clause{
		Id:         id.String(),
		Clause:     ac,
		ActionType: actionType,
	}
	a.AutoOpsRule.Clauses = append(a.AutoOpsRule.Clauses[:index], append([]*proto.Clause{clause}, a.AutoOpsRule.Clauses[index:]...)...)
	return clause, nil
}

func (a *AutoOpsRule) ChangeOpsEventRateClause(id string, oerc *proto.OpsEventRateClause) error {
	err := a.changeClause(id, oerc, oerc.ActionType)
	if err != nil {
		return err
	}
	a.AutoOpsRule.UpdatedAt = time.Now().Unix()
	if a.AutoOpsRule.AutoOpsStatus != proto.AutoOpsStatus_RUNNING {
		a.AutoOpsRule.AutoOpsStatus = proto.AutoOpsStatus_WAITING
		a.AutoOpsRule.StoppedAt = 0
	}
	return nil
}

func (a *AutoOpsRule) ChangeDatetimeClause(id string, dc *proto.DatetimeClause) error {
	var index = int64(-1)
	var swapIndex = int64(-1)
	for i, c := range a.Clauses {
		if c.Id == id {
			index = int64(i)
		}
		datetimeClause, err := a.UnmarshalDatetimeClause(c)
		if err != nil {
			return err
		}
		if datetimeClause == nil {
			continue
		}
		if swapIndex == -1 && dc.Time <= datetimeClause.Time {
			swapIndex = int64(i)
		}
	}
	if index == -1 {
		return errClauseNotFound
	}
	if swapIndex == -1 {
		swapIndex = int64(len(a.Clauses) - 1)
	}

	err := a.swapClause(id, dc, dc.ActionType, swapIndex)
	if err != nil {
		return err
	}
	a.AutoOpsRule.UpdatedAt = time.Now().Unix()
	if a.AutoOpsRule.AutoOpsStatus != proto.AutoOpsStatus_RUNNING {
		a.AutoOpsRule.AutoOpsStatus = proto.AutoOpsStatus_WAITING
		a.AutoOpsRule.StoppedAt = 0
	}
	return nil
}

func (a *AutoOpsRule) changeClause(id string, mc pb.Message, actionType proto.ActionType) error {
	for _, c := range a.Clauses {
		if c.Id == id {
			clause, err := ptypes.MarshalAny(mc)
			if err != nil {
				return err
			}
			c.Clause = clause
			c.ActionType = actionType
			return nil
		}
	}
	return errClauseNotFound
}

func (a *AutoOpsRule) swapClause(id string, mc pb.Message, actionType proto.ActionType, swapIndex int64) error {
	clause, err := ptypes.MarshalAny(mc)
	if err != nil {
		return err
	}

	index := -1
	for i, c := range a.Clauses {
		if c.Id == id {
			index = i
		}
	}
	if index == -1 || swapIndex > int64(len(a.Clauses)) {
		return errClauseNotFound
	}

	a.Clauses[index].Clause = clause
	a.Clauses[index].ActionType = actionType

	tmp := a.Clauses[swapIndex]
	copy(a.Clauses[swapIndex:], a.Clauses[swapIndex+1:])
	a.Clauses[index] = tmp
	return nil
}

func (a *AutoOpsRule) DeleteClause(id string) error {
	if len(a.Clauses) <= 1 {
		return errClauseEmpty
	}
	a.AutoOpsRule.UpdatedAt = time.Now().Unix()
	var clauses []*proto.Clause
	for i, c := range a.Clauses {
		if c.Id == id {
			clauses = append(a.Clauses[:i], a.Clauses[i+1:]...)
			continue
		}
	}
	if len(clauses) > 0 {
		a.Clauses = clauses
		return nil
	}
	return errClauseNotFound
}

func (a *AutoOpsRule) HasEventRateOps() (bool, error) {
	clauses, err := a.ExtractOpsEventRateClauses()
	if err != nil {
		return false, err
	}
	return len(clauses) > 0, nil
}

func (a *AutoOpsRule) ExtractOpsEventRateClauses() (map[string]*proto.OpsEventRateClause, error) {
	opsEventRateClauses := map[string]*proto.OpsEventRateClause{}
	for _, c := range a.Clauses {
		opsEventRateClause, err := a.UnmarshalOpsEventRateClause(c)
		if err != nil {
			return nil, err
		}
		if opsEventRateClause == nil {
			continue
		}
		opsEventRateClauses[c.Id] = opsEventRateClause
	}
	return opsEventRateClauses, nil
}

func (a *AutoOpsRule) ExtractOpsClausesForEventRateClauses() (map[string]*proto.Clause, error) {
	opsEventRateClauses := map[string]*proto.Clause{}
	for _, c := range a.Clauses {
		opsEventRateClause, err := a.UnmarshalOpsEventRateClause(c)
		if err != nil {
			return nil, err
		}
		if opsEventRateClause == nil {
			continue
		}
		opsEventRateClauses[c.Id] = c
	}
	return opsEventRateClauses, nil
}

func (a *AutoOpsRule) UnmarshalOpsEventRateClause(clause *proto.Clause) (*proto.OpsEventRateClause, error) {
	if ptypes.Is(clause.Clause, OpsEventRateClause) {
		c := &proto.OpsEventRateClause{}
		if err := ptypes.UnmarshalAny(clause.Clause, c); err != nil {
			return nil, err
		}
		return c, nil
	}
	return nil, nil
}

func (a *AutoOpsRule) HasScheduleOps() (bool, error) {
	clauses, err := a.ExtractDatetimeClauses()
	if err != nil {
		return false, err
	}
	return len(clauses) > 0, nil
}

func (a *AutoOpsRule) ExtractDatetimeClauses() ([]*proto.Clause, error) {
	datetimeClauses := []*proto.Clause{}
	for _, c := range a.Clauses {
		datetimeClause, err := a.UnmarshalDatetimeClause(c)
		if err != nil {
			return nil, err
		}
		if datetimeClause == nil {
			continue
		}
		datetimeClauses = append(datetimeClauses, c)
	}
	return datetimeClauses, nil
}

func (a *AutoOpsRule) UnmarshalDatetimeClause(clause *proto.Clause) (*proto.DatetimeClause, error) {
	if ptypes.Is(clause.Clause, DatetimeClause) {
		c := &proto.DatetimeClause{}
		if err := ptypes.UnmarshalAny(clause.Clause, c); err != nil {
			return nil, err
		}
		return c, nil
	}
	return nil, nil
}
