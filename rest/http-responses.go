package rest

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

// Maps some common http codes for more descriptive errors. Not my original code
func MapHttpToGrpcErrorCode(resp *http.Response) error {
	if resp == nil {
		return status.Error(codes.Unknown, " The error is unknown or unspecified.")
	}
	code_ := resp.StatusCode
	switch code_ {
	case 400:
		return status.Error(codes.InvalidArgument, " The client provided an invalid argument.")
	case 403:
		return status.Error(codes.PermissionDenied, " The caller does not have permission to perform the operation. ")
	case 404:
		return status.Error(codes.NotFound, " The requested entity or resource was not found. ")
	case 409:
		return status.Error(codes.Aborted, "  The operation was aborted, typically due to a concurrency issue. ")
	case 416:
		return status.Error(codes.OutOfRange, "  The operation was attempted past the valid range. ")
	case 429:
		return status.Error(codes.ResourceExhausted, "  The resource limits have been exceeded. ")
	case 499:
		return status.Error(codes.Canceled, "  The operation was cancelled (usually by the client). ")
	case 504:
		return status.Error(codes.DeadlineExceeded, "  The deadline for the operation has expired. ")
	case 501:
		return status.Error(codes.Unimplemented, "  : The requested operation is not implemented or not supported. ")
	case 503:
		return status.Error(codes.Unavailable, "  The service is currently unavailable. ")
	case 401:
		return status.Error(codes.Unauthenticated, "  The request does not have valid authentication credentials. ")
	case 422:
		return status.Error(codes.InvalidArgument, "  The request received unprocessable entity check for invalid parameters. ")
	default:
		{
			if code_ >= 400 && code_ < 500 {
				return status.Error(codes.FailedPrecondition, " The operation was rejected due to the current state of the resource.  ")
			}
			if code_ >= 500 && code_ < 600 {
				return status.Error(codes.Internal, "  : An external server error occurred. ")
			}
		}
	}
	return status.Error(codes.Unknown, "  The error is unknown or unspecified. ")
}
