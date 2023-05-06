package helper

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func PrettyGRPCError(ctx *gin.Context, err error) {
	e, ok := status.FromError(err)
	if ok {
		switch e.Code() {
		case codes.Aborted:
			ctx.AbortWithStatusJSON(409, gin.H{
				"error": "Aborted",
			})
			return
		case codes.AlreadyExists:
			ctx.AbortWithStatusJSON(409, gin.H{
				"error": e.Message(),
			})
			return
		case codes.Canceled:
			ctx.AbortWithStatusJSON(499, gin.H{
				"error": "Canceled",
			})
			return
		case codes.DataLoss:
			ctx.AbortWithStatusJSON(500, gin.H{
				"error": "Data loss",
			})
			return
		case codes.DeadlineExceeded:
			ctx.AbortWithStatusJSON(504, gin.H{
				"error": "Deadline exceeded",
			})
			return
		case codes.FailedPrecondition:
			ctx.AbortWithStatusJSON(400, gin.H{
				"error": "Failed precondition",
			})
			return
		case codes.Internal:
			ctx.AbortWithStatusJSON(500, gin.H{
				"error": e.Message(),
			})
			return
		case codes.InvalidArgument:
			ctx.AbortWithStatusJSON(400, gin.H{
				"error": e.Message(),
			})
			return
		case codes.NotFound:
			ctx.AbortWithStatusJSON(404, gin.H{
				"error": e.Message(),
			})
			return
		case codes.OK:
			ctx.AbortWithStatusJSON(200, gin.H{
				"error": "OK",
			})
			return
		case codes.OutOfRange:
			ctx.AbortWithStatusJSON(400, gin.H{
				"error": "Out of range",
			})
			return
		case codes.PermissionDenied:
			ctx.AbortWithStatusJSON(403, gin.H{
				"error": "Permission denied",
			})
			return
		case codes.ResourceExhausted:
			ctx.AbortWithStatusJSON(429, gin.H{
				"error": "Resource exhausted",
			})
			return
		case codes.Unauthenticated:
			ctx.AbortWithStatusJSON(401, gin.H{
				"error": "Unauthenticated",
			})
			return

		case codes.Unavailable:
			ctx.AbortWithStatusJSON(503, gin.H{
				"error": "Unavailable",
			})
			return
		case codes.Unimplemented:
			ctx.AbortWithStatusJSON(501, gin.H{
				"error": "Unimplemented",
			})
			return
		case codes.Unknown:
			ctx.AbortWithStatusJSON(500, gin.H{
				"error": "Unknown",
			})
			return
		default:
			ctx.AbortWithStatusJSON(500, gin.H{
				"error": "Internal server error",
			})
			return

		}
	}
	ctx.AbortWithStatusJSON(500, gin.H{
		"error": "Internal server error",
	})
}
