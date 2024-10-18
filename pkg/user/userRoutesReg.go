package user

import (
	"github.com/gin-gonic/gin"
	"github.com/shivaraj-shanthaiah/code_orbit_apigateway/pkg/user/handler"
)

func (u *User) UserSignup(ctx *gin.Context) {
	handler.UserSignupHandler(ctx, u.Client)
}

func (u *User) UserVerify(ctx *gin.Context) {
	handler.VerificationHandler(ctx, u.Client)
}

func (u *User) UserLogin(ctx *gin.Context) {
	handler.UserLoginHandler(ctx, u.Client)
}

func (u *User) ViewProfile(ctx *gin.Context) {
	handler.ViewProfileHandler(ctx, u.Client)
}

func (u *User) EditProfile(ctx *gin.Context) {
	handler.EditProfileHandler(ctx, u.Client)
}

func (u *User) ChangePassword(ctx *gin.Context) {
	handler.ChangePasswordHandler(ctx, u.Client)
}

func (u *User) UGetAllProblems(ctx *gin.Context) {
	handler.UserGetAllProblemsHandler(ctx, u.Client)
}

func (u *User) UGetProblemByID(ctx *gin.Context) {
	handler.GetProblemWithTestCasesHandler(ctx, u.Client)
}

func (u *User) SubmitCode(ctx *gin.Context) {
	handler.SubmitCodeHandler(ctx, u.Client)
}

func (u *User) GetUserStats(ctx *gin.Context) {
	handler.GetUserStatsHandler(ctx, u.Client)
}

func (u *User) GetAllPlans(ctx *gin.Context) {
	handler.UserGetAllPlans(ctx, u.Client)
}

func (u *User) GenerateInvoice(ctx *gin.Context) {
	handler.GenerateInvoiceHandler(ctx, u.Client)
}

func (u *User) MakePayment(ctx *gin.Context) {
	handler.MakePaymentHandler(ctx, u.Client)
}

func (u *User) ConfirmPayment(ctx *gin.Context) {
	handler.PaymentSuccessHandler(ctx, u.Client)
}

func (u *User) SuccessPage(ctx *gin.Context) {
	handler.PaymentSuccessPage(ctx, u.Client)
}

func (u *User) UserGetProblemStats(ctx *gin.Context){
	handler.UserGetProblemStatsHandler(ctx, u.Client)
}

func (u *User) UserGetLeaderBoard(ctx *gin.Context){
	handler.UserGetLeaderboardHandler(ctx, u.Client)
}
