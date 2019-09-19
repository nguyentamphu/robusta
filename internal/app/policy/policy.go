package policy

import (
	"context"

	"github.com/pthethanh/robusta/internal/pkg/db/mongodb"

	"github.com/casbin/casbin"
	mongodbadapter "github.com/casbin/mongodb-adapter"
)

type (
	CasbinConfig struct {
		MongoDB    mongodb.Config
		ConfigPath string `envconfig:"CONFIG_PATH" default:"configs/casbin.conf"`
	}
	Service struct {
		enforcer *casbin.Enforcer
	}
)

const (
	ObjectAny = "*"

	EffectAllow = "allow"
	EffectDeny  = "deny"

	ActionAny    = "*"
	ActionCreate = "create"
	ActionUpdate = "update"
	ActionRead   = "read"
	ActionDelete = "delete"
)

// New return a new instance of policy service
func New(enforcer *casbin.Enforcer) (*Service, error) {
	enforcer.EnableAutoSave(true) // auto save...
	if err := enforcer.LoadPolicy(); err != nil {
		return nil, err
	}
	return &Service{
		enforcer: enforcer,
	}, nil
}

// NewMongoDBCasbinEnforcer return new instance of cashout.Enforcer
// which use MongoDB as adapter.
func NewMongoDBCasbinEnforcer(conf CasbinConfig) *casbin.Enforcer {
	dialInfo := conf.MongoDB.DialInfo()
	adapter := mongodbadapter.NewAdapterWithDialInfo(dialInfo)
	enforcer := casbin.NewEnforcer(conf.ConfigPath, adapter)
	// TODO add watchers in the future to sync the policies between multiple nodes
	// see https://casbin.org/docs/en/watchers
	return enforcer
}

// AddPolicy add new policy
// For specific user:
// - s.AddPolicy("alice", "article_1", "read", allow)
// For group:
// - s.AddPolicy("group_admin", "article_1", "*", deny)
// - s.AddPolicy("group_admin", "article_2", "read", allow)
// - s.AddPolicy("group_admin", "article_3", "write", allow)
func (s *Service) AddPolicy(ctx context.Context, sub string, obj string, act string, eft string) error {
	_, err := s.enforcer.AddPolicySafe(sub, obj, act, eft)
	return err
}

// AddGroupingPolicy add grouping policy...
// Example adding user alice into the group_admin
// that would make alice inherits all permissions from the group
// - e.AddGroupingPolicy("alice", "group_admin")
func (s *Service) AddGroupingPolicy(ctx context.Context, sub string, group string) error {
	_, err := s.enforcer.AddGroupingPolicySafe(sub, group)
	return err
}

// IsAllowed check if the sub is allowed to do the act on the obj
func (s *Service) IsAllowed(ctx context.Context, sub string, obj string, act string) bool {
	ok, err := s.enforcer.EnforceSafe(sub, obj, act)
	return err == nil && ok
}

// MakeOwner make the sub to be owner of the obj
func (s *Service) MakeOwner(ctx context.Context, sub string, obj string) error {
	return s.AddPolicy(ctx, sub, obj, ActionAny, EffectAllow)
}
