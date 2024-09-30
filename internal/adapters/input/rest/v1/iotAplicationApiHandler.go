package restv1

import (
	"errors"
	"github.com/gin-gonic/gin"
	inputports "go-hexagonal-practice/internal/application/ports/input"
	"net/http"
)

type IotApplicationHandler struct {
	authHeader            string
	iotApplicationService inputports.IoTApplicationService
}

func NewIotApplicationHandler(authHeader string, iotApplicationService inputports.IoTApplicationService) *IotApplicationHandler {
	return &IotApplicationHandler{
		authHeader:            authHeader,
		iotApplicationService: iotApplicationService,
	}
}

// @tags Server
// @produce json
// @success 200 {object} domain.IoTApplication
// @failure 500 {object} ErrorResponse
// @param x-amzn-oidc-identity header string true "user Id" example(6c6ca355-8a69-475a-b8b1-16648ea4fb0a)
// @param request body ApplicationRequest true "{ "applicationName": "new-app-name" }"
// @router /v1/applications [post]
func (h *IotApplicationHandler) CreateApplication(c *gin.Context) {
	auth, err := h.authenticate(c)
	if err != nil {
		handleError(c, http.StatusUnauthorized, err.Error())
	}
	var request *ApplicationRequest
	if err := c.BindJSON(&request); err != nil {
		handleError(c, http.StatusBadRequest, "Error decoding JSON request body")
		return
	}
	if len(request.ApplicationName) == 0 {
		handleError(c, http.StatusBadRequest, "Error decoding JSON request body")
		return
	}
	cmd := &inputports.CreateIoTApplicationCommand{
		ApplicationName: request.ApplicationName,
	}
	app, err := h.iotApplicationService.CreateApplication(auth, cmd)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Service unavailable")
		return
	}
	c.JSON(http.StatusOK, app)
}

// @tags Server
// @produce json
// @success 200 {array} domain.IoTApplication
// @failure 500 {object} ErrorResponse
// @param x-amzn-oidc-identity header string true "user Id" example(6c6ca355-8a69-475a-b8b1-16648ea4fb0a)
// @router /v1/applications [get]
func (h *IotApplicationHandler) ListApplications(c *gin.Context) {
	auth, err := h.authenticate(c)
	if err != nil {
		handleError(c, http.StatusUnauthorized, err.Error())
	}
	app, err := h.iotApplicationService.ListUserApplications(auth)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Service unavailable")
		return
	}
	c.JSON(http.StatusOK, app)
}

// @tags Server
// @produce json
// @success 200 {object} domain.IoTApplication
// @failure 400 {object} ErrorResponse
// @failure 404 {object} ErrorResponse
// @failure 500 {object} ErrorResponse
// @param x-amzn-oidc-identity header string true "user Id" example(6c6ca355-8a69-475a-b8b1-16648ea4fb0a)
// @param id path string true "Application Id" example(6c6ca355-8a69-475a-b8b1-16648ea4fb0a)
// @router /v1/applications/{id} [get]
func (h *IotApplicationHandler) GetApplication(c *gin.Context) {
	auth, err := h.authenticate(c)
	if err != nil {
		handleError(c, http.StatusUnauthorized, err.Error())
	}
	appId := c.Param("id")
	app, err := h.iotApplicationService.GetApplication(auth, appId)
	if errors.Is(err, inputports.ErrNotFound) {
		handleError(c, http.StatusNotFound, "Application not found")
		return
	}
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Service unavailable")
		return
	}
	c.JSON(http.StatusOK, app)
}

// @tags Server
// @produce json
// @success 200 {object} domain.IoTApplication
// @failure 400 {object} ErrorResponse
// @failure 404 {object} ErrorResponse
// @failure 500 {object} ErrorResponse
// @param x-amzn-oidc-identity header string true "user Id" example(6c6ca355-8a69-475a-b8b1-16648ea4fb0a)
// @param request body ApplicationRequest true "{ "applicationName": "new-app-name" }"
// @param id path string true "Application Id" example(6c6ca355-8a69-475a-b8b1-16648ea4fb0a)
// @router /v1/applications/{id} [post]
func (h *IotApplicationHandler) UpdateApplication(c *gin.Context) {
	auth, err := h.authenticate(c)
	if err != nil {
		handleError(c, http.StatusUnauthorized, err.Error())
	}
	appId := c.Param("id")
	var request *ApplicationRequest
	if err := c.BindJSON(&request); err != nil {
		handleError(c, http.StatusBadRequest, "Error decoding JSON request body")
		return
	}
	if len(request.ApplicationName) == 0 {
		handleError(c, http.StatusBadRequest, "Error decoding JSON request body")
		return
	}
	cmd := &inputports.UpdateIoTApplicationCommand{
		ApplicationId:   appId,
		ApplicationName: request.ApplicationName,
	}
	app, err := h.iotApplicationService.UpdateApplication(auth, cmd)
	if errors.Is(err, inputports.ErrNotFound) {
		handleError(c, http.StatusNotFound, "Application not found")
		return
	}
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Service unavailable")
		return
	}
	c.JSON(http.StatusOK, app)
}

// @tags Server
// @produce json
// @success 200
// @failure 400 {object} ErrorResponse
// @failure 404 {object} ErrorResponse
// @failure 500 {object} ErrorResponse
// @param x-amzn-oidc-identity header string true "user Id" example(6c6ca355-8a69-475a-b8b1-16648ea4fb0a)
// @param id path string true "Application Id" example(6c6ca355-8a69-475a-b8b1-16648ea4fb0a)
// @router /v1/applications/{id} [delete]
func (h *IotApplicationHandler) DeleteApplication(c *gin.Context) {
	auth, err := h.authenticate(c)
	if err != nil {
		handleError(c, http.StatusUnauthorized, err.Error())
	}
	appId := c.Param("id")
	err = h.iotApplicationService.DeleteApplication(auth, appId)
	if errors.Is(err, inputports.ErrNotFound) {
		handleError(c, http.StatusNotFound, "Application not found")
		return
	}
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Service unavailable")
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}

func (h *IotApplicationHandler) InitRoutes(router *gin.RouterGroup) {
	router.POST("/", h.CreateApplication)
	router.POST("/:id", h.UpdateApplication)
	router.GET("/", h.ListApplications)
	router.GET("/:id", h.GetApplication)
	router.DELETE("/:id", h.DeleteApplication)
}

// Service is supposed to be running behind API Gateway which will pass userId as request header so no token verification required
func (h *IotApplicationHandler) authenticate(c *gin.Context) (*inputports.Authentication, error) {
	userId := c.GetHeader(h.authHeader)
	if len(userId) == 0 {
		return nil, errors.New("unauthenticated request attempt")
	}
	return inputports.UserAuthentication(userId), nil
}

func handleError(c *gin.Context, statusCode int, errorMessage string) {
	response := &ErrorResponse{
		Error: errorMessage,
	}
	c.JSON(statusCode, response)
}

type ApplicationRequest struct {
	ApplicationName string
}

type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}
