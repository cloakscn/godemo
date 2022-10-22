package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"example.com/ms_api/user_web/global/response"
	"example.com/ms_api/user_web/proto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleGrpcErrorToHttp
func HandleGrpcErrorToHttp(err error, ctx *gin.Context) {
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, gin.H{
					"msg": s.Message(),
				})
			case codes.Internal:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				ctx.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "其他错误 ",
				})
			}
		}
	}
}


func GetUserList(ctx *gin.Context) {
	ip := "127.0.0.1"
	port := 8080

	cc, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("GetUserList connect user server failed", "msg", err.Error())
	}

	uc := proto.NewUserClient(cc)
	rsp, err := uc.GetUserList(context.Background(), &proto.PageInfo{
		PageNum: 1,
		PageSize: 10,
	})
	if err != nil {
		zap.S().Errorw("GetUserList visited user server failed", "msg", err.Error())
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	result := make([]interface{}, 0)
	for _, v := range rsp.Data {
		data := response.UserResponse{
			Id:       v.Id,
			NickName: v.NickName,
			Birthday: time.Time(time.Unix(int64(v.Birthday), 0)).Format("2006-01-02"),
			Gender:   v.Gender,
			Mobile:   v.Mobile,
		}

		result = append(result, data)
	}
	ctx.JSON(http.StatusOK, rsp)
}