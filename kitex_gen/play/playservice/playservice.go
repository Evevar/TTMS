// Code generated by Kitex v0.4.4. DO NOT EDIT.

package playservice

import (
	play "TTMS/kitex_gen/play"
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return playServiceServiceInfo
}

var playServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "PlayService"
	handlerType := (*play.PlayService)(nil)
	methods := map[string]kitex.MethodInfo{
		"AddPlay":        kitex.NewMethodInfo(addPlayHandler, newAddPlayArgs, newAddPlayResult, false),
		"UpdatePlay":     kitex.NewMethodInfo(updatePlayHandler, newUpdatePlayArgs, newUpdatePlayResult, false),
		"DeletePlay":     kitex.NewMethodInfo(deletePlayHandler, newDeletePlayArgs, newDeletePlayResult, false),
		"GetAllPlay":     kitex.NewMethodInfo(getAllPlayHandler, newGetAllPlayArgs, newGetAllPlayResult, false),
		"AddSchedule":    kitex.NewMethodInfo(addScheduleHandler, newAddScheduleArgs, newAddScheduleResult, false),
		"UpdateSchedule": kitex.NewMethodInfo(updateScheduleHandler, newUpdateScheduleArgs, newUpdateScheduleResult, false),
		"DeleteSchedule": kitex.NewMethodInfo(deleteScheduleHandler, newDeleteScheduleArgs, newDeleteScheduleResult, false),
		"GetAllSchedule": kitex.NewMethodInfo(getAllScheduleHandler, newGetAllScheduleArgs, newGetAllScheduleResult, false),
		"PlayToSchedule": kitex.NewMethodInfo(playToScheduleHandler, newPlayToScheduleArgs, newPlayToScheduleResult, false),
		"GetSchedule":    kitex.NewMethodInfo(getScheduleHandler, newGetScheduleArgs, newGetScheduleResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "play",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func addPlayHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(play.AddPlayRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(play.PlayService).AddPlay(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *AddPlayArgs:
		success, err := handler.(play.PlayService).AddPlay(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AddPlayResult)
		realResult.Success = success
	}
	return nil
}
func newAddPlayArgs() interface{} {
	return &AddPlayArgs{}
}

func newAddPlayResult() interface{} {
	return &AddPlayResult{}
}

type AddPlayArgs struct {
	Req *play.AddPlayRequest
}

func (p *AddPlayArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(play.AddPlayRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AddPlayArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AddPlayArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AddPlayArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AddPlayArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *AddPlayArgs) Unmarshal(in []byte) error {
	msg := new(play.AddPlayRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AddPlayArgs_Req_DEFAULT *play.AddPlayRequest

func (p *AddPlayArgs) GetReq() *play.AddPlayRequest {
	if !p.IsSetReq() {
		return AddPlayArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AddPlayArgs) IsSetReq() bool {
	return p.Req != nil
}

type AddPlayResult struct {
	Success *play.AddPlayResponse
}

var AddPlayResult_Success_DEFAULT *play.AddPlayResponse

func (p *AddPlayResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(play.AddPlayResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AddPlayResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AddPlayResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AddPlayResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AddPlayResult")
	}
	return proto.Marshal(p.Success)
}

func (p *AddPlayResult) Unmarshal(in []byte) error {
	msg := new(play.AddPlayResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AddPlayResult) GetSuccess() *play.AddPlayResponse {
	if !p.IsSetSuccess() {
		return AddPlayResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AddPlayResult) SetSuccess(x interface{}) {
	p.Success = x.(*play.AddPlayResponse)
}

func (p *AddPlayResult) IsSetSuccess() bool {
	return p.Success != nil
}

func updatePlayHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(play.UpdatePlayRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(play.PlayService).UpdatePlay(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UpdatePlayArgs:
		success, err := handler.(play.PlayService).UpdatePlay(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdatePlayResult)
		realResult.Success = success
	}
	return nil
}
func newUpdatePlayArgs() interface{} {
	return &UpdatePlayArgs{}
}

func newUpdatePlayResult() interface{} {
	return &UpdatePlayResult{}
}

type UpdatePlayArgs struct {
	Req *play.UpdatePlayRequest
}

func (p *UpdatePlayArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(play.UpdatePlayRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdatePlayArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdatePlayArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdatePlayArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdatePlayArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UpdatePlayArgs) Unmarshal(in []byte) error {
	msg := new(play.UpdatePlayRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdatePlayArgs_Req_DEFAULT *play.UpdatePlayRequest

func (p *UpdatePlayArgs) GetReq() *play.UpdatePlayRequest {
	if !p.IsSetReq() {
		return UpdatePlayArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdatePlayArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdatePlayResult struct {
	Success *play.UpdatePlayResponse
}

var UpdatePlayResult_Success_DEFAULT *play.UpdatePlayResponse

func (p *UpdatePlayResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(play.UpdatePlayResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdatePlayResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdatePlayResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdatePlayResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdatePlayResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UpdatePlayResult) Unmarshal(in []byte) error {
	msg := new(play.UpdatePlayResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatePlayResult) GetSuccess() *play.UpdatePlayResponse {
	if !p.IsSetSuccess() {
		return UpdatePlayResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdatePlayResult) SetSuccess(x interface{}) {
	p.Success = x.(*play.UpdatePlayResponse)
}

func (p *UpdatePlayResult) IsSetSuccess() bool {
	return p.Success != nil
}

func deletePlayHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(play.DeletePlayRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(play.PlayService).DeletePlay(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DeletePlayArgs:
		success, err := handler.(play.PlayService).DeletePlay(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeletePlayResult)
		realResult.Success = success
	}
	return nil
}
func newDeletePlayArgs() interface{} {
	return &DeletePlayArgs{}
}

func newDeletePlayResult() interface{} {
	return &DeletePlayResult{}
}

type DeletePlayArgs struct {
	Req *play.DeletePlayRequest
}

func (p *DeletePlayArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(play.DeletePlayRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeletePlayArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeletePlayArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeletePlayArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeletePlayArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *DeletePlayArgs) Unmarshal(in []byte) error {
	msg := new(play.DeletePlayRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeletePlayArgs_Req_DEFAULT *play.DeletePlayRequest

func (p *DeletePlayArgs) GetReq() *play.DeletePlayRequest {
	if !p.IsSetReq() {
		return DeletePlayArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeletePlayArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeletePlayResult struct {
	Success *play.DeletePlayResponse
}

var DeletePlayResult_Success_DEFAULT *play.DeletePlayResponse

func (p *DeletePlayResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(play.DeletePlayResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeletePlayResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeletePlayResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeletePlayResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeletePlayResult")
	}
	return proto.Marshal(p.Success)
}

func (p *DeletePlayResult) Unmarshal(in []byte) error {
	msg := new(play.DeletePlayResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeletePlayResult) GetSuccess() *play.DeletePlayResponse {
	if !p.IsSetSuccess() {
		return DeletePlayResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeletePlayResult) SetSuccess(x interface{}) {
	p.Success = x.(*play.DeletePlayResponse)
}

func (p *DeletePlayResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getAllPlayHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(play.GetAllPlayRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(play.PlayService).GetAllPlay(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetAllPlayArgs:
		success, err := handler.(play.PlayService).GetAllPlay(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetAllPlayResult)
		realResult.Success = success
	}
	return nil
}
func newGetAllPlayArgs() interface{} {
	return &GetAllPlayArgs{}
}

func newGetAllPlayResult() interface{} {
	return &GetAllPlayResult{}
}

type GetAllPlayArgs struct {
	Req *play.GetAllPlayRequest
}

func (p *GetAllPlayArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(play.GetAllPlayRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetAllPlayArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetAllPlayArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetAllPlayArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetAllPlayArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetAllPlayArgs) Unmarshal(in []byte) error {
	msg := new(play.GetAllPlayRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetAllPlayArgs_Req_DEFAULT *play.GetAllPlayRequest

func (p *GetAllPlayArgs) GetReq() *play.GetAllPlayRequest {
	if !p.IsSetReq() {
		return GetAllPlayArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetAllPlayArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetAllPlayResult struct {
	Success *play.GetAllPlayResponse
}

var GetAllPlayResult_Success_DEFAULT *play.GetAllPlayResponse

func (p *GetAllPlayResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(play.GetAllPlayResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetAllPlayResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetAllPlayResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetAllPlayResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetAllPlayResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetAllPlayResult) Unmarshal(in []byte) error {
	msg := new(play.GetAllPlayResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAllPlayResult) GetSuccess() *play.GetAllPlayResponse {
	if !p.IsSetSuccess() {
		return GetAllPlayResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetAllPlayResult) SetSuccess(x interface{}) {
	p.Success = x.(*play.GetAllPlayResponse)
}

func (p *GetAllPlayResult) IsSetSuccess() bool {
	return p.Success != nil
}

func addScheduleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(play.AddScheduleRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(play.PlayService).AddSchedule(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *AddScheduleArgs:
		success, err := handler.(play.PlayService).AddSchedule(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AddScheduleResult)
		realResult.Success = success
	}
	return nil
}
func newAddScheduleArgs() interface{} {
	return &AddScheduleArgs{}
}

func newAddScheduleResult() interface{} {
	return &AddScheduleResult{}
}

type AddScheduleArgs struct {
	Req *play.AddScheduleRequest
}

func (p *AddScheduleArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(play.AddScheduleRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AddScheduleArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AddScheduleArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AddScheduleArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AddScheduleArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *AddScheduleArgs) Unmarshal(in []byte) error {
	msg := new(play.AddScheduleRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AddScheduleArgs_Req_DEFAULT *play.AddScheduleRequest

func (p *AddScheduleArgs) GetReq() *play.AddScheduleRequest {
	if !p.IsSetReq() {
		return AddScheduleArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AddScheduleArgs) IsSetReq() bool {
	return p.Req != nil
}

type AddScheduleResult struct {
	Success *play.AddScheduleResponse
}

var AddScheduleResult_Success_DEFAULT *play.AddScheduleResponse

func (p *AddScheduleResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(play.AddScheduleResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AddScheduleResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AddScheduleResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AddScheduleResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AddScheduleResult")
	}
	return proto.Marshal(p.Success)
}

func (p *AddScheduleResult) Unmarshal(in []byte) error {
	msg := new(play.AddScheduleResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AddScheduleResult) GetSuccess() *play.AddScheduleResponse {
	if !p.IsSetSuccess() {
		return AddScheduleResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AddScheduleResult) SetSuccess(x interface{}) {
	p.Success = x.(*play.AddScheduleResponse)
}

func (p *AddScheduleResult) IsSetSuccess() bool {
	return p.Success != nil
}

func updateScheduleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(play.UpdateScheduleRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(play.PlayService).UpdateSchedule(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UpdateScheduleArgs:
		success, err := handler.(play.PlayService).UpdateSchedule(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateScheduleResult)
		realResult.Success = success
	}
	return nil
}
func newUpdateScheduleArgs() interface{} {
	return &UpdateScheduleArgs{}
}

func newUpdateScheduleResult() interface{} {
	return &UpdateScheduleResult{}
}

type UpdateScheduleArgs struct {
	Req *play.UpdateScheduleRequest
}

func (p *UpdateScheduleArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(play.UpdateScheduleRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateScheduleArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateScheduleArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateScheduleArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateScheduleArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateScheduleArgs) Unmarshal(in []byte) error {
	msg := new(play.UpdateScheduleRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateScheduleArgs_Req_DEFAULT *play.UpdateScheduleRequest

func (p *UpdateScheduleArgs) GetReq() *play.UpdateScheduleRequest {
	if !p.IsSetReq() {
		return UpdateScheduleArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateScheduleArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateScheduleResult struct {
	Success *play.UpdateScheduleResponse
}

var UpdateScheduleResult_Success_DEFAULT *play.UpdateScheduleResponse

func (p *UpdateScheduleResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(play.UpdateScheduleResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateScheduleResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateScheduleResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateScheduleResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateScheduleResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateScheduleResult) Unmarshal(in []byte) error {
	msg := new(play.UpdateScheduleResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateScheduleResult) GetSuccess() *play.UpdateScheduleResponse {
	if !p.IsSetSuccess() {
		return UpdateScheduleResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateScheduleResult) SetSuccess(x interface{}) {
	p.Success = x.(*play.UpdateScheduleResponse)
}

func (p *UpdateScheduleResult) IsSetSuccess() bool {
	return p.Success != nil
}

func deleteScheduleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(play.DeleteScheduleRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(play.PlayService).DeleteSchedule(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DeleteScheduleArgs:
		success, err := handler.(play.PlayService).DeleteSchedule(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteScheduleResult)
		realResult.Success = success
	}
	return nil
}
func newDeleteScheduleArgs() interface{} {
	return &DeleteScheduleArgs{}
}

func newDeleteScheduleResult() interface{} {
	return &DeleteScheduleResult{}
}

type DeleteScheduleArgs struct {
	Req *play.DeleteScheduleRequest
}

func (p *DeleteScheduleArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(play.DeleteScheduleRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteScheduleArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteScheduleArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteScheduleArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteScheduleArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteScheduleArgs) Unmarshal(in []byte) error {
	msg := new(play.DeleteScheduleRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteScheduleArgs_Req_DEFAULT *play.DeleteScheduleRequest

func (p *DeleteScheduleArgs) GetReq() *play.DeleteScheduleRequest {
	if !p.IsSetReq() {
		return DeleteScheduleArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteScheduleArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteScheduleResult struct {
	Success *play.DeleteScheduleResponse
}

var DeleteScheduleResult_Success_DEFAULT *play.DeleteScheduleResponse

func (p *DeleteScheduleResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(play.DeleteScheduleResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteScheduleResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteScheduleResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteScheduleResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteScheduleResult")
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteScheduleResult) Unmarshal(in []byte) error {
	msg := new(play.DeleteScheduleResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteScheduleResult) GetSuccess() *play.DeleteScheduleResponse {
	if !p.IsSetSuccess() {
		return DeleteScheduleResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteScheduleResult) SetSuccess(x interface{}) {
	p.Success = x.(*play.DeleteScheduleResponse)
}

func (p *DeleteScheduleResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getAllScheduleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(play.GetAllScheduleRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(play.PlayService).GetAllSchedule(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetAllScheduleArgs:
		success, err := handler.(play.PlayService).GetAllSchedule(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetAllScheduleResult)
		realResult.Success = success
	}
	return nil
}
func newGetAllScheduleArgs() interface{} {
	return &GetAllScheduleArgs{}
}

func newGetAllScheduleResult() interface{} {
	return &GetAllScheduleResult{}
}

type GetAllScheduleArgs struct {
	Req *play.GetAllScheduleRequest
}

func (p *GetAllScheduleArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(play.GetAllScheduleRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetAllScheduleArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetAllScheduleArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetAllScheduleArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetAllScheduleArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetAllScheduleArgs) Unmarshal(in []byte) error {
	msg := new(play.GetAllScheduleRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetAllScheduleArgs_Req_DEFAULT *play.GetAllScheduleRequest

func (p *GetAllScheduleArgs) GetReq() *play.GetAllScheduleRequest {
	if !p.IsSetReq() {
		return GetAllScheduleArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetAllScheduleArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetAllScheduleResult struct {
	Success *play.GetAllScheduleResponse
}

var GetAllScheduleResult_Success_DEFAULT *play.GetAllScheduleResponse

func (p *GetAllScheduleResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(play.GetAllScheduleResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetAllScheduleResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetAllScheduleResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetAllScheduleResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetAllScheduleResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetAllScheduleResult) Unmarshal(in []byte) error {
	msg := new(play.GetAllScheduleResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAllScheduleResult) GetSuccess() *play.GetAllScheduleResponse {
	if !p.IsSetSuccess() {
		return GetAllScheduleResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetAllScheduleResult) SetSuccess(x interface{}) {
	p.Success = x.(*play.GetAllScheduleResponse)
}

func (p *GetAllScheduleResult) IsSetSuccess() bool {
	return p.Success != nil
}

func playToScheduleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(play.PlayToScheduleRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(play.PlayService).PlayToSchedule(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PlayToScheduleArgs:
		success, err := handler.(play.PlayService).PlayToSchedule(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PlayToScheduleResult)
		realResult.Success = success
	}
	return nil
}
func newPlayToScheduleArgs() interface{} {
	return &PlayToScheduleArgs{}
}

func newPlayToScheduleResult() interface{} {
	return &PlayToScheduleResult{}
}

type PlayToScheduleArgs struct {
	Req *play.PlayToScheduleRequest
}

func (p *PlayToScheduleArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(play.PlayToScheduleRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PlayToScheduleArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PlayToScheduleArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PlayToScheduleArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PlayToScheduleArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PlayToScheduleArgs) Unmarshal(in []byte) error {
	msg := new(play.PlayToScheduleRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PlayToScheduleArgs_Req_DEFAULT *play.PlayToScheduleRequest

func (p *PlayToScheduleArgs) GetReq() *play.PlayToScheduleRequest {
	if !p.IsSetReq() {
		return PlayToScheduleArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PlayToScheduleArgs) IsSetReq() bool {
	return p.Req != nil
}

type PlayToScheduleResult struct {
	Success *play.PlayToScheduleResponse
}

var PlayToScheduleResult_Success_DEFAULT *play.PlayToScheduleResponse

func (p *PlayToScheduleResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(play.PlayToScheduleResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PlayToScheduleResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PlayToScheduleResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PlayToScheduleResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PlayToScheduleResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PlayToScheduleResult) Unmarshal(in []byte) error {
	msg := new(play.PlayToScheduleResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PlayToScheduleResult) GetSuccess() *play.PlayToScheduleResponse {
	if !p.IsSetSuccess() {
		return PlayToScheduleResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PlayToScheduleResult) SetSuccess(x interface{}) {
	p.Success = x.(*play.PlayToScheduleResponse)
}

func (p *PlayToScheduleResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getScheduleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(play.GetScheduleRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(play.PlayService).GetSchedule(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetScheduleArgs:
		success, err := handler.(play.PlayService).GetSchedule(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetScheduleResult)
		realResult.Success = success
	}
	return nil
}
func newGetScheduleArgs() interface{} {
	return &GetScheduleArgs{}
}

func newGetScheduleResult() interface{} {
	return &GetScheduleResult{}
}

type GetScheduleArgs struct {
	Req *play.GetScheduleRequest
}

func (p *GetScheduleArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(play.GetScheduleRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetScheduleArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetScheduleArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetScheduleArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetScheduleArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetScheduleArgs) Unmarshal(in []byte) error {
	msg := new(play.GetScheduleRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetScheduleArgs_Req_DEFAULT *play.GetScheduleRequest

func (p *GetScheduleArgs) GetReq() *play.GetScheduleRequest {
	if !p.IsSetReq() {
		return GetScheduleArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetScheduleArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetScheduleResult struct {
	Success *play.GetScheduleResponse
}

var GetScheduleResult_Success_DEFAULT *play.GetScheduleResponse

func (p *GetScheduleResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(play.GetScheduleResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetScheduleResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetScheduleResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetScheduleResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetScheduleResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetScheduleResult) Unmarshal(in []byte) error {
	msg := new(play.GetScheduleResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetScheduleResult) GetSuccess() *play.GetScheduleResponse {
	if !p.IsSetSuccess() {
		return GetScheduleResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetScheduleResult) SetSuccess(x interface{}) {
	p.Success = x.(*play.GetScheduleResponse)
}

func (p *GetScheduleResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) AddPlay(ctx context.Context, Req *play.AddPlayRequest) (r *play.AddPlayResponse, err error) {
	var _args AddPlayArgs
	_args.Req = Req
	var _result AddPlayResult
	if err = p.c.Call(ctx, "AddPlay", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdatePlay(ctx context.Context, Req *play.UpdatePlayRequest) (r *play.UpdatePlayResponse, err error) {
	var _args UpdatePlayArgs
	_args.Req = Req
	var _result UpdatePlayResult
	if err = p.c.Call(ctx, "UpdatePlay", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeletePlay(ctx context.Context, Req *play.DeletePlayRequest) (r *play.DeletePlayResponse, err error) {
	var _args DeletePlayArgs
	_args.Req = Req
	var _result DeletePlayResult
	if err = p.c.Call(ctx, "DeletePlay", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAllPlay(ctx context.Context, Req *play.GetAllPlayRequest) (r *play.GetAllPlayResponse, err error) {
	var _args GetAllPlayArgs
	_args.Req = Req
	var _result GetAllPlayResult
	if err = p.c.Call(ctx, "GetAllPlay", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddSchedule(ctx context.Context, Req *play.AddScheduleRequest) (r *play.AddScheduleResponse, err error) {
	var _args AddScheduleArgs
	_args.Req = Req
	var _result AddScheduleResult
	if err = p.c.Call(ctx, "AddSchedule", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateSchedule(ctx context.Context, Req *play.UpdateScheduleRequest) (r *play.UpdateScheduleResponse, err error) {
	var _args UpdateScheduleArgs
	_args.Req = Req
	var _result UpdateScheduleResult
	if err = p.c.Call(ctx, "UpdateSchedule", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteSchedule(ctx context.Context, Req *play.DeleteScheduleRequest) (r *play.DeleteScheduleResponse, err error) {
	var _args DeleteScheduleArgs
	_args.Req = Req
	var _result DeleteScheduleResult
	if err = p.c.Call(ctx, "DeleteSchedule", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAllSchedule(ctx context.Context, Req *play.GetAllScheduleRequest) (r *play.GetAllScheduleResponse, err error) {
	var _args GetAllScheduleArgs
	_args.Req = Req
	var _result GetAllScheduleResult
	if err = p.c.Call(ctx, "GetAllSchedule", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PlayToSchedule(ctx context.Context, Req *play.PlayToScheduleRequest) (r *play.PlayToScheduleResponse, err error) {
	var _args PlayToScheduleArgs
	_args.Req = Req
	var _result PlayToScheduleResult
	if err = p.c.Call(ctx, "PlayToSchedule", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetSchedule(ctx context.Context, Req *play.GetScheduleRequest) (r *play.GetScheduleResponse, err error) {
	var _args GetScheduleArgs
	_args.Req = Req
	var _result GetScheduleResult
	if err = p.c.Call(ctx, "GetSchedule", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
