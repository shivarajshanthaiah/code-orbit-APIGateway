package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/admin/handler"
)

func (a *Admin) AdminLogin(ctx *gin.Context) {
	handler.AdminLoginHandler(ctx, a.Client)
}

func (a *Admin) BlockUser(ctx *gin.Context) {
	handler.BlockUserHandler(ctx, a.Client)
}

func (a *Admin) UnblockUser(ctx *gin.Context) {
	handler.UnBlockUserHandler(ctx, a.Client)
}

func (a *Admin) GetAllUsers(ctx *gin.Context) {
	handler.FindAllUsersHandler(ctx, a.Client)
}

func (a *Admin) GetUserByID(ctx *gin.Context) {
	handler.FindUserByIDHandler(ctx, a.Client)
}

func (a *Admin) InsertProblem(ctx *gin.Context) {
	handler.InsertProblemHanlder(ctx, a.Client)
}

func (a *Admin) GetAllProblems(ctx *gin.Context) {
	handler.AdminGetAllProblemsHandler(ctx, a.Client)
}

func (a *Admin) EditProblem(ctx *gin.Context) {
	handler.EditProblemHandler(ctx, a.Client)
}

func (a *Admin) InsertTestCases(ctx *gin.Context) {
	handler.InsertTestCaseHandler(ctx, a.Client)
}

func (a *Admin) UpdateTestCases(ctx *gin.Context) {
	handler.UpdateTestCaseHandler(ctx, a.Client)
}
