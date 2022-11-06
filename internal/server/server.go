package server

import (
	"context"
	"github.com/go-grpc-course/config/configpb"
	"github.com/go-grpc-course/internal/vault"
	"log"
)

type Server struct {
	configpb.UnimplementedConfigServiceServer
}

func (s *Server) Put(ctx context.Context, req *configpb.ConfigRequestWithData) (*configpb.ConfigResponse, error) {
	log.Printf("invoke put request %v", req)
	reqData := make(map[string]interface{}, 0)
	for k, v := range req.GetData() {
		reqData[k] = v
	}
	reqData["process"] = false

	client, err := vault.InitClient()
	if err != nil {
		return nil, err
	}

	KVSecret, err := client.KVv2(vault.MountPath).Put(ctx, req.GetService(), reqData)
	if err != nil {
		return nil, err
	}

	var respData map[string]string
	for key, value := range KVSecret.Data {
		if v, ok := value.(string); ok {
			respData[key] = v
		}
	}

	return &configpb.ConfigResponse{Status: "successful", Data: respData}, nil
}

func (s *Server) Get(ctx context.Context, req *configpb.ConfigRequestWithoutData) (*configpb.ConfigResponse, error) {
	log.Printf("invoke get request %v", req)
	client, err := vault.InitClient()
	if err != nil {
		return nil, err
	}

	processPatch := make(map[string]interface{}, 0)
	processPatch["process"] = true
	_, err = client.KVv2(vault.MountPath).Patch(ctx, req.GetService(), processPatch)
	if err != nil {
		return nil, err
	}

	KVSecret, err := client.KVv2(vault.MountPath).Get(ctx, req.GetService())
	if err != nil {
		return nil, err
	}

	respData := make(map[string]string, 0)
	for key, value := range KVSecret.Data {
		if v, ok := value.(string); ok {
			respData[key] = v
		}
	}

	return &configpb.ConfigResponse{Status: "successful", Data: respData}, nil
}

func (s *Server) Update(ctx context.Context, req *configpb.ConfigRequestWithData) (*configpb.ConfigResponse, error) {
	log.Printf("invoke update request %v", req)
	reqData := make(map[string]interface{}, 0)
	for k, v := range req.GetData() {
		reqData[k] = v
	}

	client, err := vault.InitClient()
	if err != nil {
		return nil, err
	}

	reqData["process"] = false

	KVSecret, err := client.KVv2(vault.MountPath).Patch(ctx, req.GetService(), reqData)
	if err != nil {
		return nil, err
	}

	var respData map[string]string
	for key, value := range KVSecret.Data {
		if v, ok := value.(string); ok {
			respData[key] = v
		}
	}

	return &configpb.ConfigResponse{Status: "successful", Data: respData}, nil
}

func (s *Server) Delete(ctx context.Context, req *configpb.ConfigRequestWithoutData) (*configpb.ConfigResponse, error) {
	log.Printf("invoke delete request %v", req)
	client, err := vault.InitClient()
	if err != nil {
		return nil, err
	}

	KVSecret, err := client.KVv2(vault.MountPath).Get(ctx, req.GetService())
	if err != nil {
		return nil, err
	}

	if val, ok := KVSecret.Data["process"]; ok {
		if val == false {
			err = client.KVv2(vault.MountPath).Delete(ctx, req.GetService())
			if err != nil {
				return nil, err
			}

			return &configpb.ConfigResponse{Status: "config deleted"}, nil
		} else {
			return &configpb.ConfigResponse{Status: "config is already use"}, nil
		}
	}
	return &configpb.ConfigResponse{Status: "error: no 'process' data", Data: nil}, nil
}
