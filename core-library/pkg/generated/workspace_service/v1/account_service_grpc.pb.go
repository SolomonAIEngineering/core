// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: workspace_service/v1/account_service.proto

package workspace_servicev1

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
	WorkspaceService_CreateAccount_FullMethodName   = "/workspace_service.v1.WorkspaceService/CreateAccount"
	WorkspaceService_GetAccount_FullMethodName      = "/workspace_service.v1.WorkspaceService/GetAccount"
	WorkspaceService_DeleteAccount_FullMethodName   = "/workspace_service.v1.WorkspaceService/DeleteAccount"
	WorkspaceService_DeleteFile_FullMethodName      = "/workspace_service.v1.WorkspaceService/DeleteFile"
	WorkspaceService_UpdateFile_FullMethodName      = "/workspace_service.v1.WorkspaceService/UpdateFile"
	WorkspaceService_CreateFolder_FullMethodName    = "/workspace_service.v1.WorkspaceService/CreateFolder"
	WorkspaceService_ListFolder_FullMethodName      = "/workspace_service.v1.WorkspaceService/ListFolder"
	WorkspaceService_UpdateFolder_FullMethodName    = "/workspace_service.v1.WorkspaceService/UpdateFolder"
	WorkspaceService_DeleteFolder_FullMethodName    = "/workspace_service.v1.WorkspaceService/DeleteFolder"
	WorkspaceService_CreateWorkspace_FullMethodName = "/workspace_service.v1.WorkspaceService/CreateWorkspace"
	WorkspaceService_ListWorkspace_FullMethodName   = "/workspace_service.v1.WorkspaceService/ListWorkspace"
	WorkspaceService_DeleteWorkspace_FullMethodName = "/workspace_service.v1.WorkspaceService/DeleteWorkspace"
	WorkspaceService_UpdateWorkspace_FullMethodName = "/workspace_service.v1.WorkspaceService/UpdateWorkspace"
)

// WorkspaceServiceClient is the client API for WorkspaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkspaceServiceClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error)
	// Delete an Account
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error)
	// Delete a File
	DeleteFile(ctx context.Context, in *DeleteFileRequest, opts ...grpc.CallOption) (*DeleteFileResponse, error)
	// Update a File
	UpdateFile(ctx context.Context, in *UpdateFileRequest, opts ...grpc.CallOption) (*UpdateFileResponse, error)
	// Create a Folder Metadata
	CreateFolder(ctx context.Context, in *CreateFolderRequest, opts ...grpc.CallOption) (*CreateFolderResponse, error)
	// Read a Folder Metadata
	ListFolder(ctx context.Context, in *ListFolderRequest, opts ...grpc.CallOption) (*ListFolderResponse, error)
	// Update a Folder Metadata
	UpdateFolder(ctx context.Context, in *UpdateFolderRequest, opts ...grpc.CallOption) (*UpdateFolderResponse, error)
	// Delete a Folder Metadata
	DeleteFolder(ctx context.Context, in *DeleteFolderRequest, opts ...grpc.CallOption) (*DeleteFolderResponse, error)
	// Create a Workspace
	CreateWorkspace(ctx context.Context, in *CreateWorkspaceRequest, opts ...grpc.CallOption) (*CreateWorkspaceResponse, error)
	// Read a Workspace
	ListWorkspace(ctx context.Context, in *ListWorkspaceRequest, opts ...grpc.CallOption) (*ListWorkspaceResponse, error)
	// Delete a Workspace
	DeleteWorkspace(ctx context.Context, in *DeleteWorkspaceRequest, opts ...grpc.CallOption) (*DeleteWorkspaceResponse, error)
	// Update a Workspace
	UpdateWorkspace(ctx context.Context, in *UpdateWorkspaceRequest, opts ...grpc.CallOption) (*UpdateWorkspaceResponse, error)
}

type workspaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkspaceServiceClient(cc grpc.ClientConnInterface) WorkspaceServiceClient {
	return &workspaceServiceClient{cc}
}

func (c *workspaceServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_CreateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	out := new(GetAccountResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_GetAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error) {
	out := new(DeleteAccountResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_DeleteAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) DeleteFile(ctx context.Context, in *DeleteFileRequest, opts ...grpc.CallOption) (*DeleteFileResponse, error) {
	out := new(DeleteFileResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_DeleteFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) UpdateFile(ctx context.Context, in *UpdateFileRequest, opts ...grpc.CallOption) (*UpdateFileResponse, error) {
	out := new(UpdateFileResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_UpdateFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) CreateFolder(ctx context.Context, in *CreateFolderRequest, opts ...grpc.CallOption) (*CreateFolderResponse, error) {
	out := new(CreateFolderResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_CreateFolder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) ListFolder(ctx context.Context, in *ListFolderRequest, opts ...grpc.CallOption) (*ListFolderResponse, error) {
	out := new(ListFolderResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_ListFolder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) UpdateFolder(ctx context.Context, in *UpdateFolderRequest, opts ...grpc.CallOption) (*UpdateFolderResponse, error) {
	out := new(UpdateFolderResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_UpdateFolder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) DeleteFolder(ctx context.Context, in *DeleteFolderRequest, opts ...grpc.CallOption) (*DeleteFolderResponse, error) {
	out := new(DeleteFolderResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_DeleteFolder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) CreateWorkspace(ctx context.Context, in *CreateWorkspaceRequest, opts ...grpc.CallOption) (*CreateWorkspaceResponse, error) {
	out := new(CreateWorkspaceResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_CreateWorkspace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) ListWorkspace(ctx context.Context, in *ListWorkspaceRequest, opts ...grpc.CallOption) (*ListWorkspaceResponse, error) {
	out := new(ListWorkspaceResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_ListWorkspace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) DeleteWorkspace(ctx context.Context, in *DeleteWorkspaceRequest, opts ...grpc.CallOption) (*DeleteWorkspaceResponse, error) {
	out := new(DeleteWorkspaceResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_DeleteWorkspace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) UpdateWorkspace(ctx context.Context, in *UpdateWorkspaceRequest, opts ...grpc.CallOption) (*UpdateWorkspaceResponse, error) {
	out := new(UpdateWorkspaceResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_UpdateWorkspace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspaceServiceServer is the server API for WorkspaceService service.
// All implementations must embed UnimplementedWorkspaceServiceServer
// for forward compatibility
type WorkspaceServiceServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error)
	// Delete an Account
	DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error)
	// Delete a File
	DeleteFile(context.Context, *DeleteFileRequest) (*DeleteFileResponse, error)
	// Update a File
	UpdateFile(context.Context, *UpdateFileRequest) (*UpdateFileResponse, error)
	// Create a Folder Metadata
	CreateFolder(context.Context, *CreateFolderRequest) (*CreateFolderResponse, error)
	// Read a Folder Metadata
	ListFolder(context.Context, *ListFolderRequest) (*ListFolderResponse, error)
	// Update a Folder Metadata
	UpdateFolder(context.Context, *UpdateFolderRequest) (*UpdateFolderResponse, error)
	// Delete a Folder Metadata
	DeleteFolder(context.Context, *DeleteFolderRequest) (*DeleteFolderResponse, error)
	// Create a Workspace
	CreateWorkspace(context.Context, *CreateWorkspaceRequest) (*CreateWorkspaceResponse, error)
	// Read a Workspace
	ListWorkspace(context.Context, *ListWorkspaceRequest) (*ListWorkspaceResponse, error)
	// Delete a Workspace
	DeleteWorkspace(context.Context, *DeleteWorkspaceRequest) (*DeleteWorkspaceResponse, error)
	// Update a Workspace
	UpdateWorkspace(context.Context, *UpdateWorkspaceRequest) (*UpdateWorkspaceResponse, error)
	mustEmbedUnimplementedWorkspaceServiceServer()
}

// UnimplementedWorkspaceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkspaceServiceServer struct {
}

func (UnimplementedWorkspaceServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedWorkspaceServiceServer) GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedWorkspaceServiceServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedWorkspaceServiceServer) DeleteFile(context.Context, *DeleteFileRequest) (*DeleteFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (UnimplementedWorkspaceServiceServer) UpdateFile(context.Context, *UpdateFileRequest) (*UpdateFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFile not implemented")
}
func (UnimplementedWorkspaceServiceServer) CreateFolder(context.Context, *CreateFolderRequest) (*CreateFolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFolder not implemented")
}
func (UnimplementedWorkspaceServiceServer) ListFolder(context.Context, *ListFolderRequest) (*ListFolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFolder not implemented")
}
func (UnimplementedWorkspaceServiceServer) UpdateFolder(context.Context, *UpdateFolderRequest) (*UpdateFolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFolder not implemented")
}
func (UnimplementedWorkspaceServiceServer) DeleteFolder(context.Context, *DeleteFolderRequest) (*DeleteFolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFolder not implemented")
}
func (UnimplementedWorkspaceServiceServer) CreateWorkspace(context.Context, *CreateWorkspaceRequest) (*CreateWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) ListWorkspace(context.Context, *ListWorkspaceRequest) (*ListWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) DeleteWorkspace(context.Context, *DeleteWorkspaceRequest) (*DeleteWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) UpdateWorkspace(context.Context, *UpdateWorkspaceRequest) (*UpdateWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) mustEmbedUnimplementedWorkspaceServiceServer() {}

// UnsafeWorkspaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkspaceServiceServer will
// result in compilation errors.
type UnsafeWorkspaceServiceServer interface {
	mustEmbedUnimplementedWorkspaceServiceServer()
}

func RegisterWorkspaceServiceServer(s grpc.ServiceRegistrar, srv WorkspaceServiceServer) {
	s.RegisterService(&WorkspaceService_ServiceDesc, srv)
}

func _WorkspaceService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_GetAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_DeleteAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_DeleteFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).DeleteFile(ctx, req.(*DeleteFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_UpdateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).UpdateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_UpdateFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).UpdateFile(ctx, req.(*UpdateFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_CreateFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).CreateFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_CreateFolder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).CreateFolder(ctx, req.(*CreateFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_ListFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).ListFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_ListFolder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).ListFolder(ctx, req.(*ListFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_UpdateFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).UpdateFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_UpdateFolder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).UpdateFolder(ctx, req.(*UpdateFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_DeleteFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).DeleteFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_DeleteFolder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).DeleteFolder(ctx, req.(*DeleteFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_CreateWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).CreateWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_CreateWorkspace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).CreateWorkspace(ctx, req.(*CreateWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_ListWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).ListWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_ListWorkspace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).ListWorkspace(ctx, req.(*ListWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_DeleteWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).DeleteWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_DeleteWorkspace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).DeleteWorkspace(ctx, req.(*DeleteWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_UpdateWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).UpdateWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_UpdateWorkspace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).UpdateWorkspace(ctx, req.(*UpdateWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkspaceService_ServiceDesc is the grpc.ServiceDesc for WorkspaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkspaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "workspace_service.v1.WorkspaceService",
	HandlerType: (*WorkspaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _WorkspaceService_CreateAccount_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _WorkspaceService_GetAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _WorkspaceService_DeleteAccount_Handler,
		},
		{
			MethodName: "DeleteFile",
			Handler:    _WorkspaceService_DeleteFile_Handler,
		},
		{
			MethodName: "UpdateFile",
			Handler:    _WorkspaceService_UpdateFile_Handler,
		},
		{
			MethodName: "CreateFolder",
			Handler:    _WorkspaceService_CreateFolder_Handler,
		},
		{
			MethodName: "ListFolder",
			Handler:    _WorkspaceService_ListFolder_Handler,
		},
		{
			MethodName: "UpdateFolder",
			Handler:    _WorkspaceService_UpdateFolder_Handler,
		},
		{
			MethodName: "DeleteFolder",
			Handler:    _WorkspaceService_DeleteFolder_Handler,
		},
		{
			MethodName: "CreateWorkspace",
			Handler:    _WorkspaceService_CreateWorkspace_Handler,
		},
		{
			MethodName: "ListWorkspace",
			Handler:    _WorkspaceService_ListWorkspace_Handler,
		},
		{
			MethodName: "DeleteWorkspace",
			Handler:    _WorkspaceService_DeleteWorkspace_Handler,
		},
		{
			MethodName: "UpdateWorkspace",
			Handler:    _WorkspaceService_UpdateWorkspace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workspace_service/v1/account_service.proto",
}
