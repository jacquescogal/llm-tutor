# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from src.protos import memory_pb2 as protos_dot_memory__pb2

GRPC_GENERATED_VERSION = '1.66.2'
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in protos/memory_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class MemoryServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateMemory = channel.unary_unary(
                '/memory.MemoryService/CreateMemory',
                request_serializer=protos_dot_memory__pb2.CreateMemoryRequest.SerializeToString,
                response_deserializer=protos_dot_memory__pb2.CreateMemoryResponse.FromString,
                _registered_method=True)
        self.GetMemoryById = channel.unary_unary(
                '/memory.MemoryService/GetMemoryById',
                request_serializer=protos_dot_memory__pb2.GetMemoryByIdRequest.SerializeToString,
                response_deserializer=protos_dot_memory__pb2.GetMemoryByIdResponse.FromString,
                _registered_method=True)
        self.GetMemoriesByDocId = channel.unary_unary(
                '/memory.MemoryService/GetMemoriesByDocId',
                request_serializer=protos_dot_memory__pb2.GetMemoriesByDocIdRequest.SerializeToString,
                response_deserializer=protos_dot_memory__pb2.GetMemoriesByDocIdResponse.FromString,
                _registered_method=True)
        self.GetMemoriesByMemoryTitleSearch = channel.unary_unary(
                '/memory.MemoryService/GetMemoriesByMemoryTitleSearch',
                request_serializer=protos_dot_memory__pb2.GetMemoriesByMemoryTitleSearchRequest.SerializeToString,
                response_deserializer=protos_dot_memory__pb2.GetMemoriesByMemoryTitleSearchResponse.FromString,
                _registered_method=True)
        self.UpdateMemory = channel.unary_unary(
                '/memory.MemoryService/UpdateMemory',
                request_serializer=protos_dot_memory__pb2.UpdateMemoryRequest.SerializeToString,
                response_deserializer=protos_dot_memory__pb2.UpdateMemoryResponse.FromString,
                _registered_method=True)
        self.DeleteMemory = channel.unary_unary(
                '/memory.MemoryService/DeleteMemory',
                request_serializer=protos_dot_memory__pb2.DeleteMemoryRequest.SerializeToString,
                response_deserializer=protos_dot_memory__pb2.DeleteMemoryResponse.FromString,
                _registered_method=True)


class MemoryServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateMemory(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetMemoryById(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetMemoriesByDocId(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetMemoriesByMemoryTitleSearch(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateMemory(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteMemory(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MemoryServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateMemory': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateMemory,
                    request_deserializer=protos_dot_memory__pb2.CreateMemoryRequest.FromString,
                    response_serializer=protos_dot_memory__pb2.CreateMemoryResponse.SerializeToString,
            ),
            'GetMemoryById': grpc.unary_unary_rpc_method_handler(
                    servicer.GetMemoryById,
                    request_deserializer=protos_dot_memory__pb2.GetMemoryByIdRequest.FromString,
                    response_serializer=protos_dot_memory__pb2.GetMemoryByIdResponse.SerializeToString,
            ),
            'GetMemoriesByDocId': grpc.unary_unary_rpc_method_handler(
                    servicer.GetMemoriesByDocId,
                    request_deserializer=protos_dot_memory__pb2.GetMemoriesByDocIdRequest.FromString,
                    response_serializer=protos_dot_memory__pb2.GetMemoriesByDocIdResponse.SerializeToString,
            ),
            'GetMemoriesByMemoryTitleSearch': grpc.unary_unary_rpc_method_handler(
                    servicer.GetMemoriesByMemoryTitleSearch,
                    request_deserializer=protos_dot_memory__pb2.GetMemoriesByMemoryTitleSearchRequest.FromString,
                    response_serializer=protos_dot_memory__pb2.GetMemoriesByMemoryTitleSearchResponse.SerializeToString,
            ),
            'UpdateMemory': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateMemory,
                    request_deserializer=protos_dot_memory__pb2.UpdateMemoryRequest.FromString,
                    response_serializer=protos_dot_memory__pb2.UpdateMemoryResponse.SerializeToString,
            ),
            'DeleteMemory': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteMemory,
                    request_deserializer=protos_dot_memory__pb2.DeleteMemoryRequest.FromString,
                    response_serializer=protos_dot_memory__pb2.DeleteMemoryResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'memory.MemoryService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('memory.MemoryService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class MemoryService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateMemory(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/memory.MemoryService/CreateMemory',
            protos_dot_memory__pb2.CreateMemoryRequest.SerializeToString,
            protos_dot_memory__pb2.CreateMemoryResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetMemoryById(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/memory.MemoryService/GetMemoryById',
            protos_dot_memory__pb2.GetMemoryByIdRequest.SerializeToString,
            protos_dot_memory__pb2.GetMemoryByIdResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetMemoriesByDocId(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/memory.MemoryService/GetMemoriesByDocId',
            protos_dot_memory__pb2.GetMemoriesByDocIdRequest.SerializeToString,
            protos_dot_memory__pb2.GetMemoriesByDocIdResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetMemoriesByMemoryTitleSearch(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/memory.MemoryService/GetMemoriesByMemoryTitleSearch',
            protos_dot_memory__pb2.GetMemoriesByMemoryTitleSearchRequest.SerializeToString,
            protos_dot_memory__pb2.GetMemoriesByMemoryTitleSearchResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def UpdateMemory(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/memory.MemoryService/UpdateMemory',
            protos_dot_memory__pb2.UpdateMemoryRequest.SerializeToString,
            protos_dot_memory__pb2.UpdateMemoryResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def DeleteMemory(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/memory.MemoryService/DeleteMemory',
            protos_dot_memory__pb2.DeleteMemoryRequest.SerializeToString,
            protos_dot_memory__pb2.DeleteMemoryResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
