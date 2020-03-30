package rpc

import (
	"time"

	sliverpb "github.com/bishopfox/sliver/protobuf/sliver"
	"github.com/bishopfox/sliver/server/core"
	"github.com/golang/protobuf/proto"
)

/*
	Sliver Implant Framework
	Copyright (C) 2019  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

func rpcNamedPipesListener(req []byte, timeout time.Duration, resp RPCResponse) {
	namedpipeReq := &sliverpb.NamedPipesReq{}
	err := proto.Unmarshal(req, namedpipeReq)
	if err != nil {
		resp([]byte{}, err)
		return
	}
	sliver := (*core.Hive.Slivers)[namedpipeReq.SliverID]
	if sliver == nil {
		resp([]byte{}, err)
		return
	}
	data, _ := proto.Marshal(&sliverpb.NamedPipesReq{
		PipeName: namedpipeReq.PipeName,
	})
	data, err = sliver.Request(sliverpb.MsgNamedPipesReq, timeout, data)
	resp(data, err)
}
