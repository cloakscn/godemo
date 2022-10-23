package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"example.com/ms_api/user_web/forms"
	"example.com/ms_api/user_web/global"
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

// GetUserList
func GetUserList(ctx *gin.Context) {
	cc, err := grpc.Dial(fmt.Sprintf("%s:%s", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("GetUserList connect user server failed", "msg", err.Error())
	}

	uc := proto.NewUserClient(cc)
	
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "0"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	zap.S().Debug(pageNum)
	zap.S().Debug(pageSize)

	rsp, err := uc.GetUserList(context.Background(), &proto.PageInfo{
		PageNum: uint32(pageNum),
		PageSize: uint32(pageSize),
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

// PasswordLogin
func PasswordLogin(ctx *gin.Context) {
	passwordLoginForm := forms.PasswordLoginForm{}
	if err := ctx.ShouldBindJSON(&passwordLoginForm); err != nil {
		zap.S().Errorf(err.Error())
	}
}