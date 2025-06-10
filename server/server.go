package server

import (
	pb "NotificationSender/proto"
	"NotificationSender/service"
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type notificationServer struct {
	pb.UnimplementedNotificationServiceServer
}

func (s *notificationServer) SendUpdateNotification(ctx context.Context, req *pb.UpdateNotificationRequest) (*pb.UpdateNotificationResponse, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Ошибка загрузки .env файла: %v", err)
	}
	log.Printf("Отправка уведомления пользователю %d:", req.UserId)

	oldStatus, newStatus, err := service.GetLastStatusChange(req.RequestId)
	if err != nil {
		log.Printf("Ошибка получения статусов: %v", err)
		return &pb.UpdateNotificationResponse{Success: false}, nil
	}

	err = service.SendUpdateEmail("yar1k3lfg@gmail.com", os.Getenv("UPDATE_EMAIL_SUBJECT"), "service/updateTemplate.html", req.RequestId, oldStatus, newStatus)
	if err != nil {
		log.Printf("Ошибка отправки email: %v", err)
		return &pb.UpdateNotificationResponse{Success: false}, nil
	}
	// Здесь будет логика отправки email/SMS (пока просто лог)
	return &pb.UpdateNotificationResponse{Success: true}, nil
}

func (s *notificationServer) SendCreateNotification(ctx context.Context, req *pb.CreateNotificationRequest) (*pb.CreateNotificationResponse, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Ошибка загрузки .env файла: %v", err)
	}
	log.Printf("Отправка уведомления пользователю %d:", req.UserId)

	err = service.SendCreateEmail("yar1k3lfg@gmail.com", os.Getenv("CREATE_EMAIL_SUBJECT"), "service/createTemplate.html", req.RequestId)
	if err != nil {
		log.Printf("Ошибка отправки email: %v", err)
		return &pb.CreateNotificationResponse{Success: false}, nil
	}
	//Здесь будет логика отправки email/SMS (пока просто лог)
	return &pb.CreateNotificationResponse{Success: true}, nil
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Ошибка запуска gRPC-сервера: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNotificationServiceServer(grpcServer, &notificationServer{})

	log.Println("NotificationService запущен на порту 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Ошибка запуска: %v", err)
	}
}
