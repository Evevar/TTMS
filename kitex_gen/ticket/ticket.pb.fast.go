// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package ticket

import (
	studio "TTMS/kitex_gen/studio"
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *BaseResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BaseResp[number], err)
}

func (x *BaseResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BaseResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMessage, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Ticket) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Ticket[number], err)
}

func (x *Ticket) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Ticket) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ScheduleId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Ticket) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.SeatRow, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Ticket) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.SeatCol, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Ticket) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Price, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Ticket) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.PlayName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Ticket) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.StudioId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Ticket) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.Status, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *BatchAddTicketRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BatchAddTicketRequest[number], err)
}

func (x *BatchAddTicketRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ScheduleId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BatchAddTicketRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StudioId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BatchAddTicketRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Price, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *BatchAddTicketRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.PlayName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BatchAddTicketRequest) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	var v studio.Seat
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.List = append(x.List, &v)
	return offset, nil
}

func (x *BatchAddTicketResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BatchAddTicketResponse[number], err)
}

func (x *BatchAddTicketResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *UpdateTicketRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpdateTicketRequest[number], err)
}

func (x *UpdateTicketRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ScheduleId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UpdateTicketRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.SeatRow, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *UpdateTicketRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.SeatCol, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *UpdateTicketRequest) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Price, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *UpdateTicketRequest) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Status, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *UpdateTicketResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpdateTicketResponse[number], err)
}

func (x *UpdateTicketResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *GetAllTicketRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetAllTicketRequest[number], err)
}

func (x *GetAllTicketRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ScheduleId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetAllTicketResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetAllTicketResponse[number], err)
}

func (x *GetAllTicketResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *GetAllTicketResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v Ticket
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.List = append(x.List, &v)
	return offset, nil
}

func (x *BuyTicketRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BuyTicketRequest[number], err)
}

func (x *BuyTicketRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ScheduleId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BuyTicketRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SeatRow, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *BuyTicketRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.SeatCol, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *BuyTicketRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BuyTicketResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BuyTicketResponse[number], err)
}

func (x *BuyTicketResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *ReturnTicketRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ReturnTicketRequest[number], err)
}

func (x *ReturnTicketRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *ReturnTicketRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ScheduleId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *ReturnTicketRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.SeatRow, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *ReturnTicketRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.SeatCol, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *ReturnTicketResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ReturnTicketResponse[number], err)
}

func (x *ReturnTicketResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *BaseResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *BaseResp) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *BaseResp) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMessage == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.StatusMessage)
	return offset
}

func (x *Ticket) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	return offset
}

func (x *Ticket) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.Id)
	return offset
}

func (x *Ticket) fastWriteField2(buf []byte) (offset int) {
	if x.ScheduleId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.ScheduleId)
	return offset
}

func (x *Ticket) fastWriteField3(buf []byte) (offset int) {
	if x.SeatRow == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.SeatRow)
	return offset
}

func (x *Ticket) fastWriteField4(buf []byte) (offset int) {
	if x.SeatCol == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 4, x.SeatCol)
	return offset
}

func (x *Ticket) fastWriteField5(buf []byte) (offset int) {
	if x.Price == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 5, x.Price)
	return offset
}

func (x *Ticket) fastWriteField6(buf []byte) (offset int) {
	if x.PlayName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.PlayName)
	return offset
}

func (x *Ticket) fastWriteField7(buf []byte) (offset int) {
	if x.StudioId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 7, x.StudioId)
	return offset
}

func (x *Ticket) fastWriteField8(buf []byte) (offset int) {
	if x.Status == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 8, x.Status)
	return offset
}

func (x *BatchAddTicketRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *BatchAddTicketRequest) fastWriteField1(buf []byte) (offset int) {
	if x.ScheduleId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.ScheduleId)
	return offset
}

func (x *BatchAddTicketRequest) fastWriteField2(buf []byte) (offset int) {
	if x.StudioId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.StudioId)
	return offset
}

func (x *BatchAddTicketRequest) fastWriteField3(buf []byte) (offset int) {
	if x.Price == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.Price)
	return offset
}

func (x *BatchAddTicketRequest) fastWriteField4(buf []byte) (offset int) {
	if x.PlayName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.PlayName)
	return offset
}

func (x *BatchAddTicketRequest) fastWriteField5(buf []byte) (offset int) {
	if x.List == nil {
		return offset
	}
	for i := range x.List {
		offset += fastpb.WriteMessage(buf[offset:], 5, x.List[i])
	}
	return offset
}

func (x *BatchAddTicketResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *BatchAddTicketResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.BaseResp)
	return offset
}

func (x *UpdateTicketRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	return offset
}

func (x *UpdateTicketRequest) fastWriteField1(buf []byte) (offset int) {
	if x.ScheduleId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.ScheduleId)
	return offset
}

func (x *UpdateTicketRequest) fastWriteField3(buf []byte) (offset int) {
	if x.SeatRow == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.SeatRow)
	return offset
}

func (x *UpdateTicketRequest) fastWriteField4(buf []byte) (offset int) {
	if x.SeatCol == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 4, x.SeatCol)
	return offset
}

func (x *UpdateTicketRequest) fastWriteField5(buf []byte) (offset int) {
	if x.Price == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 5, x.Price)
	return offset
}

func (x *UpdateTicketRequest) fastWriteField6(buf []byte) (offset int) {
	if x.Status == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 6, x.Status)
	return offset
}

func (x *UpdateTicketResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *UpdateTicketResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.BaseResp)
	return offset
}

func (x *GetAllTicketRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetAllTicketRequest) fastWriteField1(buf []byte) (offset int) {
	if x.ScheduleId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.ScheduleId)
	return offset
}

func (x *GetAllTicketResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetAllTicketResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.BaseResp)
	return offset
}

func (x *GetAllTicketResponse) fastWriteField2(buf []byte) (offset int) {
	if x.List == nil {
		return offset
	}
	for i := range x.List {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.List[i])
	}
	return offset
}

func (x *BuyTicketRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *BuyTicketRequest) fastWriteField1(buf []byte) (offset int) {
	if x.ScheduleId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.ScheduleId)
	return offset
}

func (x *BuyTicketRequest) fastWriteField2(buf []byte) (offset int) {
	if x.SeatRow == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.SeatRow)
	return offset
}

func (x *BuyTicketRequest) fastWriteField3(buf []byte) (offset int) {
	if x.SeatCol == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.SeatCol)
	return offset
}

func (x *BuyTicketRequest) fastWriteField4(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.UserId)
	return offset
}

func (x *BuyTicketResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *BuyTicketResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.BaseResp)
	return offset
}

func (x *ReturnTicketRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *ReturnTicketRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.UserId)
	return offset
}

func (x *ReturnTicketRequest) fastWriteField2(buf []byte) (offset int) {
	if x.ScheduleId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.ScheduleId)
	return offset
}

func (x *ReturnTicketRequest) fastWriteField3(buf []byte) (offset int) {
	if x.SeatRow == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.SeatRow)
	return offset
}

func (x *ReturnTicketRequest) fastWriteField4(buf []byte) (offset int) {
	if x.SeatCol == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 4, x.SeatCol)
	return offset
}

func (x *ReturnTicketResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *ReturnTicketResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.BaseResp)
	return offset
}

func (x *BaseResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *BaseResp) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.StatusCode)
	return n
}

func (x *BaseResp) sizeField2() (n int) {
	if x.StatusMessage == "" {
		return n
	}
	n += fastpb.SizeString(2, x.StatusMessage)
	return n
}

func (x *Ticket) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	return n
}

func (x *Ticket) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.Id)
	return n
}

func (x *Ticket) sizeField2() (n int) {
	if x.ScheduleId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.ScheduleId)
	return n
}

func (x *Ticket) sizeField3() (n int) {
	if x.SeatRow == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.SeatRow)
	return n
}

func (x *Ticket) sizeField4() (n int) {
	if x.SeatCol == 0 {
		return n
	}
	n += fastpb.SizeInt32(4, x.SeatCol)
	return n
}

func (x *Ticket) sizeField5() (n int) {
	if x.Price == 0 {
		return n
	}
	n += fastpb.SizeInt32(5, x.Price)
	return n
}

func (x *Ticket) sizeField6() (n int) {
	if x.PlayName == "" {
		return n
	}
	n += fastpb.SizeString(6, x.PlayName)
	return n
}

func (x *Ticket) sizeField7() (n int) {
	if x.StudioId == 0 {
		return n
	}
	n += fastpb.SizeInt64(7, x.StudioId)
	return n
}

func (x *Ticket) sizeField8() (n int) {
	if x.Status == 0 {
		return n
	}
	n += fastpb.SizeInt32(8, x.Status)
	return n
}

func (x *BatchAddTicketRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *BatchAddTicketRequest) sizeField1() (n int) {
	if x.ScheduleId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.ScheduleId)
	return n
}

func (x *BatchAddTicketRequest) sizeField2() (n int) {
	if x.StudioId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.StudioId)
	return n
}

func (x *BatchAddTicketRequest) sizeField3() (n int) {
	if x.Price == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.Price)
	return n
}

func (x *BatchAddTicketRequest) sizeField4() (n int) {
	if x.PlayName == "" {
		return n
	}
	n += fastpb.SizeString(4, x.PlayName)
	return n
}

func (x *BatchAddTicketRequest) sizeField5() (n int) {
	if x.List == nil {
		return n
	}
	for i := range x.List {
		n += fastpb.SizeMessage(5, x.List[i])
	}
	return n
}

func (x *BatchAddTicketResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *BatchAddTicketResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.BaseResp)
	return n
}

func (x *UpdateTicketRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	return n
}

func (x *UpdateTicketRequest) sizeField1() (n int) {
	if x.ScheduleId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.ScheduleId)
	return n
}

func (x *UpdateTicketRequest) sizeField3() (n int) {
	if x.SeatRow == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.SeatRow)
	return n
}

func (x *UpdateTicketRequest) sizeField4() (n int) {
	if x.SeatCol == 0 {
		return n
	}
	n += fastpb.SizeInt32(4, x.SeatCol)
	return n
}

func (x *UpdateTicketRequest) sizeField5() (n int) {
	if x.Price == 0 {
		return n
	}
	n += fastpb.SizeInt32(5, x.Price)
	return n
}

func (x *UpdateTicketRequest) sizeField6() (n int) {
	if x.Status == 0 {
		return n
	}
	n += fastpb.SizeInt32(6, x.Status)
	return n
}

func (x *UpdateTicketResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *UpdateTicketResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.BaseResp)
	return n
}

func (x *GetAllTicketRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetAllTicketRequest) sizeField1() (n int) {
	if x.ScheduleId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.ScheduleId)
	return n
}

func (x *GetAllTicketResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetAllTicketResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.BaseResp)
	return n
}

func (x *GetAllTicketResponse) sizeField2() (n int) {
	if x.List == nil {
		return n
	}
	for i := range x.List {
		n += fastpb.SizeMessage(2, x.List[i])
	}
	return n
}

func (x *BuyTicketRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *BuyTicketRequest) sizeField1() (n int) {
	if x.ScheduleId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.ScheduleId)
	return n
}

func (x *BuyTicketRequest) sizeField2() (n int) {
	if x.SeatRow == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.SeatRow)
	return n
}

func (x *BuyTicketRequest) sizeField3() (n int) {
	if x.SeatCol == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.SeatCol)
	return n
}

func (x *BuyTicketRequest) sizeField4() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.UserId)
	return n
}

func (x *BuyTicketResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *BuyTicketResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.BaseResp)
	return n
}

func (x *ReturnTicketRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *ReturnTicketRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.UserId)
	return n
}

func (x *ReturnTicketRequest) sizeField2() (n int) {
	if x.ScheduleId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.ScheduleId)
	return n
}

func (x *ReturnTicketRequest) sizeField3() (n int) {
	if x.SeatRow == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.SeatRow)
	return n
}

func (x *ReturnTicketRequest) sizeField4() (n int) {
	if x.SeatCol == 0 {
		return n
	}
	n += fastpb.SizeInt32(4, x.SeatCol)
	return n
}

func (x *ReturnTicketResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *ReturnTicketResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.BaseResp)
	return n
}

var fieldIDToName_BaseResp = map[int32]string{
	1: "StatusCode",
	2: "StatusMessage",
}

var fieldIDToName_Ticket = map[int32]string{
	1: "Id",
	2: "ScheduleId",
	3: "SeatRow",
	4: "SeatCol",
	5: "Price",
	6: "PlayName",
	7: "StudioId",
	8: "Status",
}

var fieldIDToName_BatchAddTicketRequest = map[int32]string{
	1: "ScheduleId",
	2: "StudioId",
	3: "Price",
	4: "PlayName",
	5: "List",
}

var fieldIDToName_BatchAddTicketResponse = map[int32]string{
	1: "BaseResp",
}

var fieldIDToName_UpdateTicketRequest = map[int32]string{
	1: "ScheduleId",
	3: "SeatRow",
	4: "SeatCol",
	5: "Price",
	6: "Status",
}

var fieldIDToName_UpdateTicketResponse = map[int32]string{
	1: "BaseResp",
}

var fieldIDToName_GetAllTicketRequest = map[int32]string{
	1: "ScheduleId",
}

var fieldIDToName_GetAllTicketResponse = map[int32]string{
	1: "BaseResp",
	2: "List",
}

var fieldIDToName_BuyTicketRequest = map[int32]string{
	1: "ScheduleId",
	2: "SeatRow",
	3: "SeatCol",
	4: "UserId",
}

var fieldIDToName_BuyTicketResponse = map[int32]string{
	1: "BaseResp",
}

var fieldIDToName_ReturnTicketRequest = map[int32]string{
	1: "UserId",
	2: "ScheduleId",
	3: "SeatRow",
	4: "SeatCol",
}

var fieldIDToName_ReturnTicketResponse = map[int32]string{
	1: "BaseResp",
}

var _ = studio.File_studio_proto