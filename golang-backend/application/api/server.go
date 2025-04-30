package api

import (
	"context"
	"database/sql"
	"demo/application/internal/repository/postgredb"
	"demo/application/middleware"
	"demo/application/service/patientservice"
	"demo/application/service/userservice"
	"demo/db"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	router         *gin.Engine
	logger         *zap.Logger
	patientService patientservice.PatientService
	userservice    *userservice.UserService
}

func NewServer(ctx context.Context, logger *zap.Logger, conn *sql.DB) *Server {
	queries := db.New(conn)
	database := postgredb.NewRepo(queries)
	patientService := patientservice.NewPatientService(ctx, logger, database)
	userService := userservice.NewService(ctx, logger, database)
	server := &Server{
		logger:         logger,
		patientService: *patientService,
		userservice:    userService,
	}

	router := gin.Default()

	// Public Routes
	router.POST("/signup", server.SignUp)
	router.POST("/login", server.Login)

	// Protected Routes - require JWT authentication
	auth := router.Group("/")
	auth.Use(middleware.JWTAuth())

	// Receptionist Routes - require receptionist role
	receptionist := auth.Group("/receptionist")
	receptionist.Use(middleware.RoleAuth("receptionist"))
	receptionist.POST("/addPatient", server.addPatient)
	receptionist.GET("/getPatient/:id", server.getPatient)
	receptionist.GET("/listPatients", server.listPatients)

	// Doctor Routes - require doctor role
	doctor := auth.Group("/doctor")
	doctor.Use(middleware.RoleAuth("doctor"))
	doctor.POST("/createAppointment/:patient_id", server.createAppointment)
	doctor.GET("/getAppointment/:appointment_id", server.getAppointment)
	doctor.GET("/listAppointmentsByPatient/:patient_id", server.listAppointmentsByPatient)
	doctor.PUT("/updatePatientStatus/:id", server.updatePatientStatus)

	// Future endpoints
	// router.PUT("/updateAppointment/:appointment_id", server.updateAppointment)
	// router.DELETE("/deleteAppointment/:appointment_id", server.deleteAppointment)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
