// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package pilot

import (
	"os"
	"time"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"openpitrix.io/openpitrix/pkg/logger"
	"openpitrix.io/openpitrix/pkg/manager"
	"openpitrix.io/openpitrix/pkg/pb/metadata/pilot"
	"openpitrix.io/openpitrix/pkg/pb/metadata/types"
	"openpitrix.io/openpitrix/pkg/service/metadata/pilot/pilotutil"
)

type Server struct {
	cfg            *pbtypes.PilotConfig
	pbClientTlsCfg *pbtypes.PilotClientTLSConfig
	fgClientMgr    *FrontgateClientManager
	taskStatusMgr  *TaskStatusManager
}

func Serve(cfg *pbtypes.PilotConfig, pbClientTlsCfg *pbtypes.PilotClientTLSConfig, opts ...Options) {
	if cfg != nil {
		cfg = proto.Clone(cfg).(*pbtypes.PilotConfig)
	} else {
		cfg = NewDefaultConfig()
	}

	for _, fn := range opts {
		fn(cfg)
	}

	p := &Server{
		cfg:            cfg,
		pbClientTlsCfg: proto.Clone(pbClientTlsCfg).(*pbtypes.PilotClientTLSConfig),
		fgClientMgr:    NewFrontgateClientManager(),
		taskStatusMgr:  NewTaskStatusManager(),
	}

	go func() {
		for {
			p.fgClientMgr.CheckAllClient()
			time.Sleep(time.Second * 10)
		}
	}()

	// internal service
	go manager.NewGrpcServer("pilot-service", int(p.cfg.ListenPort)).Serve(
		func(server *grpc.Server) {
			pbpilot.RegisterPilotServiceServer(server, p)
		},
	)

	// tls for public service
	if pbClientTlsCfg != nil {
		tlsCfg, err := pilotutil.NewClientTLSConfigFromPbConfig(pbClientTlsCfg)
		if err != nil {
			logger.Critical("%+v", err)
			os.Exit(1)
		}

		manager.NewGrpcServer("pilot-service-for-frontgate", int(p.cfg.ForFrontgateListenPort)).Serve(
			func(server *grpc.Server) {
				pbpilot.RegisterPilotServiceForFrontgateServer(server, p)
			},
			grpc.Creds(credentials.NewTLS(tlsCfg)),
		)
	} else {
		manager.NewGrpcServer("pilot-service-for-frontgate", int(p.cfg.ForFrontgateListenPort)).Serve(
			func(server *grpc.Server) {
				pbpilot.RegisterPilotServiceForFrontgateServer(server, p)
			},
		)
	}
}
