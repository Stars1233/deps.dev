// Copyright 2024 Google LLC
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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.3
// source: api.proto

package v3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Insights_GetPackage_FullMethodName                = "/deps_dev.v3.Insights/GetPackage"
	Insights_GetVersion_FullMethodName                = "/deps_dev.v3.Insights/GetVersion"
	Insights_GetRequirements_FullMethodName           = "/deps_dev.v3.Insights/GetRequirements"
	Insights_GetDependencies_FullMethodName           = "/deps_dev.v3.Insights/GetDependencies"
	Insights_GetProject_FullMethodName                = "/deps_dev.v3.Insights/GetProject"
	Insights_GetProjectPackageVersions_FullMethodName = "/deps_dev.v3.Insights/GetProjectPackageVersions"
	Insights_GetAdvisory_FullMethodName               = "/deps_dev.v3.Insights/GetAdvisory"
	Insights_Query_FullMethodName                     = "/deps_dev.v3.Insights/Query"
)

// InsightsClient is the client API for Insights service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InsightsClient interface {
	// GetPackage returns information about a package, including a list of its
	// available versions, with the default version marked if known.
	GetPackage(ctx context.Context, in *GetPackageRequest, opts ...grpc.CallOption) (*Package, error)
	// GetVersion returns information about a specific package version, including
	// its licenses and any security advisories known to affect it.
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*Version, error)
	// GetRequirements returns the requirements for a given version in a
	// system-specific format. Requirements are currently available for
	// Maven, npm, NuGet and RubyGems.
	//
	// Requirements are the dependency constraints specified by the version.
	GetRequirements(ctx context.Context, in *GetRequirementsRequest, opts ...grpc.CallOption) (*Requirements, error)
	// GetDependencies returns a resolved dependency graph for the given package
	// version. Dependencies are currently available for Go, npm, Cargo, Maven
	// and PyPI.
	//
	// Dependencies are the resolution of the requirements (dependency
	// constraints) specified by a version.
	//
	// The dependency graph should be similar to one produced by installing the
	// package version on a generic 64-bit Linux system, with no other
	// dependencies present. The precise meaning of this varies from system to
	// system.
	GetDependencies(ctx context.Context, in *GetDependenciesRequest, opts ...grpc.CallOption) (*Dependencies, error)
	// GetProject returns information about projects hosted by GitHub, GitLab, or
	// BitBucket, when known to us.
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error)
	// GetProjectPackageVersions returns known mappings between the requested
	// project and package versions.
	// At most 1500 package versions are returned. Mappings which were derived
	// from attestations are served first.
	GetProjectPackageVersions(ctx context.Context, in *GetProjectPackageVersionsRequest, opts ...grpc.CallOption) (*ProjectPackageVersions, error)
	// GetAdvisory returns information about security advisories hosted by OSV.
	GetAdvisory(ctx context.Context, in *GetAdvisoryRequest, opts ...grpc.CallOption) (*Advisory, error)
	// Query returns information about multiple package versions, which can be
	// specified by name, content hash, or both. If a hash was specified in the
	// request, it returns the artifacts that matched the hash.
	//
	// Querying by content hash is currently supported for npm, Cargo, Maven,
	// NuGet, PyPI and RubyGems. It is typical for hash queries to return many
	// results; hashes are matched against multiple release artifacts (such as
	// JAR files) that comprise package versions, and any given artifact may
	// appear in several package versions.
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResult, error)
}

type insightsClient struct {
	cc grpc.ClientConnInterface
}

func NewInsightsClient(cc grpc.ClientConnInterface) InsightsClient {
	return &insightsClient{cc}
}

func (c *insightsClient) GetPackage(ctx context.Context, in *GetPackageRequest, opts ...grpc.CallOption) (*Package, error) {
	out := new(Package)
	err := c.cc.Invoke(ctx, Insights_GetPackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *insightsClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := c.cc.Invoke(ctx, Insights_GetVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *insightsClient) GetRequirements(ctx context.Context, in *GetRequirementsRequest, opts ...grpc.CallOption) (*Requirements, error) {
	out := new(Requirements)
	err := c.cc.Invoke(ctx, Insights_GetRequirements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *insightsClient) GetDependencies(ctx context.Context, in *GetDependenciesRequest, opts ...grpc.CallOption) (*Dependencies, error) {
	out := new(Dependencies)
	err := c.cc.Invoke(ctx, Insights_GetDependencies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *insightsClient) GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, Insights_GetProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *insightsClient) GetProjectPackageVersions(ctx context.Context, in *GetProjectPackageVersionsRequest, opts ...grpc.CallOption) (*ProjectPackageVersions, error) {
	out := new(ProjectPackageVersions)
	err := c.cc.Invoke(ctx, Insights_GetProjectPackageVersions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *insightsClient) GetAdvisory(ctx context.Context, in *GetAdvisoryRequest, opts ...grpc.CallOption) (*Advisory, error) {
	out := new(Advisory)
	err := c.cc.Invoke(ctx, Insights_GetAdvisory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *insightsClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResult, error) {
	out := new(QueryResult)
	err := c.cc.Invoke(ctx, Insights_Query_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InsightsServer is the server API for Insights service.
// All implementations must embed UnimplementedInsightsServer
// for forward compatibility
type InsightsServer interface {
	// GetPackage returns information about a package, including a list of its
	// available versions, with the default version marked if known.
	GetPackage(context.Context, *GetPackageRequest) (*Package, error)
	// GetVersion returns information about a specific package version, including
	// its licenses and any security advisories known to affect it.
	GetVersion(context.Context, *GetVersionRequest) (*Version, error)
	// GetRequirements returns the requirements for a given version in a
	// system-specific format. Requirements are currently available for
	// Maven, npm, NuGet and RubyGems.
	//
	// Requirements are the dependency constraints specified by the version.
	GetRequirements(context.Context, *GetRequirementsRequest) (*Requirements, error)
	// GetDependencies returns a resolved dependency graph for the given package
	// version. Dependencies are currently available for Go, npm, Cargo, Maven
	// and PyPI.
	//
	// Dependencies are the resolution of the requirements (dependency
	// constraints) specified by a version.
	//
	// The dependency graph should be similar to one produced by installing the
	// package version on a generic 64-bit Linux system, with no other
	// dependencies present. The precise meaning of this varies from system to
	// system.
	GetDependencies(context.Context, *GetDependenciesRequest) (*Dependencies, error)
	// GetProject returns information about projects hosted by GitHub, GitLab, or
	// BitBucket, when known to us.
	GetProject(context.Context, *GetProjectRequest) (*Project, error)
	// GetProjectPackageVersions returns known mappings between the requested
	// project and package versions.
	// At most 1500 package versions are returned. Mappings which were derived
	// from attestations are served first.
	GetProjectPackageVersions(context.Context, *GetProjectPackageVersionsRequest) (*ProjectPackageVersions, error)
	// GetAdvisory returns information about security advisories hosted by OSV.
	GetAdvisory(context.Context, *GetAdvisoryRequest) (*Advisory, error)
	// Query returns information about multiple package versions, which can be
	// specified by name, content hash, or both. If a hash was specified in the
	// request, it returns the artifacts that matched the hash.
	//
	// Querying by content hash is currently supported for npm, Cargo, Maven,
	// NuGet, PyPI and RubyGems. It is typical for hash queries to return many
	// results; hashes are matched against multiple release artifacts (such as
	// JAR files) that comprise package versions, and any given artifact may
	// appear in several package versions.
	Query(context.Context, *QueryRequest) (*QueryResult, error)
	mustEmbedUnimplementedInsightsServer()
}

// UnimplementedInsightsServer must be embedded to have forward compatible implementations.
type UnimplementedInsightsServer struct {
}

func (UnimplementedInsightsServer) GetPackage(context.Context, *GetPackageRequest) (*Package, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPackage not implemented")
}
func (UnimplementedInsightsServer) GetVersion(context.Context, *GetVersionRequest) (*Version, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedInsightsServer) GetRequirements(context.Context, *GetRequirementsRequest) (*Requirements, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequirements not implemented")
}
func (UnimplementedInsightsServer) GetDependencies(context.Context, *GetDependenciesRequest) (*Dependencies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDependencies not implemented")
}
func (UnimplementedInsightsServer) GetProject(context.Context, *GetProjectRequest) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedInsightsServer) GetProjectPackageVersions(context.Context, *GetProjectPackageVersionsRequest) (*ProjectPackageVersions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectPackageVersions not implemented")
}
func (UnimplementedInsightsServer) GetAdvisory(context.Context, *GetAdvisoryRequest) (*Advisory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdvisory not implemented")
}
func (UnimplementedInsightsServer) Query(context.Context, *QueryRequest) (*QueryResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedInsightsServer) mustEmbedUnimplementedInsightsServer() {}

// UnsafeInsightsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InsightsServer will
// result in compilation errors.
type UnsafeInsightsServer interface {
	mustEmbedUnimplementedInsightsServer()
}

func RegisterInsightsServer(s grpc.ServiceRegistrar, srv InsightsServer) {
	s.RegisterService(&Insights_ServiceDesc, srv)
}

func _Insights_GetPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).GetPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_GetPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).GetPackage(ctx, req.(*GetPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Insights_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_GetVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Insights_GetRequirements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequirementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).GetRequirements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_GetRequirements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).GetRequirements(ctx, req.(*GetRequirementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Insights_GetDependencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDependenciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).GetDependencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_GetDependencies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).GetDependencies(ctx, req.(*GetDependenciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Insights_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_GetProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).GetProject(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Insights_GetProjectPackageVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectPackageVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).GetProjectPackageVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_GetProjectPackageVersions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).GetProjectPackageVersions(ctx, req.(*GetProjectPackageVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Insights_GetAdvisory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdvisoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).GetAdvisory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_GetAdvisory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).GetAdvisory(ctx, req.(*GetAdvisoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Insights_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_Query_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Insights_ServiceDesc is the grpc.ServiceDesc for Insights service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Insights_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "deps_dev.v3.Insights",
	HandlerType: (*InsightsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPackage",
			Handler:    _Insights_GetPackage_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _Insights_GetVersion_Handler,
		},
		{
			MethodName: "GetRequirements",
			Handler:    _Insights_GetRequirements_Handler,
		},
		{
			MethodName: "GetDependencies",
			Handler:    _Insights_GetDependencies_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _Insights_GetProject_Handler,
		},
		{
			MethodName: "GetProjectPackageVersions",
			Handler:    _Insights_GetProjectPackageVersions_Handler,
		},
		{
			MethodName: "GetAdvisory",
			Handler:    _Insights_GetAdvisory_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Insights_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
