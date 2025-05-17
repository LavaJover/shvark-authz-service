package grpcapi

import (
	"context"

	authzpb "github.com/LavaJover/shvark-authz-service/proto/gen"
	"github.com/casbin/casbin/v2"
)

type AuthzService struct {
	authzpb.UnimplementedAuthzServiceServer
	Enforcer *casbin.Enforcer
}

func (s *AuthzService) AssignRole(ctx context.Context, req *authzpb.AssignRoleRequest) (*authzpb.AssignRoleResponse, error) {
	ok, err := s.Enforcer.AddGroupingPolicy(req.UserId, req.Role)
	return &authzpb.AssignRoleResponse{Success: ok}, err
}

func (s *AuthzService) RevokeRole(ctx context.Context, req *authzpb.RevokeRoleRequest) (*authzpb.RevokeRoleResponse, error) {
	ok, err := s.Enforcer.RemoveGroupingPolicy(req.UserId, req.Role)
	return &authzpb.RevokeRoleResponse{Success: ok}, err
}

func (s *AuthzService) AddPolicy(ctx context.Context, req *authzpb.AddPolicyRequest) (*authzpb.AddPolicyResponse, error) {
	ok, err := s.Enforcer.AddPolicy(req.Subject, req.Object, req.Action)
	return &authzpb.AddPolicyResponse{Success: ok}, err
}

func (s *AuthzService) DeletePolicy(ctx context.Context, req *authzpb.DeletePolicyRequest) (*authzpb.DeletePolicyResponse, error) {
	ok, err := s.Enforcer.RemovePolicy(req.Subject, req.Object, req.Action)
	return &authzpb.DeletePolicyResponse{Success: ok}, err
}

func (s *AuthzService) CheckPermission(ctx context.Context, req *authzpb.CheckPermissionRequest) (*authzpb.CheckPermissionResponse, error) {
	ok, err := s.Enforcer.Enforce(req.UserId, req.Object, req.Action)
	return &authzpb.CheckPermissionResponse{Allowed: ok}, err
}