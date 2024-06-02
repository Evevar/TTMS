import { json } from "react-router-dom";
import axios from "../utils/request";

//演出厅列表
export const $schedule = async (params) => {
    let { data } = await axios.get('ttms/schedule/all', {
        params: {
            Current: params,
            PageSize: 10
        }
    })
    // sessionStorage.setItem('token', success.Token);
    // console.log(data)
    return data
    // return data
}

//添加演出厅
export const $addschedule = async (params) => {
    let mon = (params.ShowTime.$M + 1) < 10 ? '0' + (params.ShowTime.$M + 1) : params.ShowTime.$M + 1
    let day = (params.ShowTime.$D + 1) < 10 ? '0' + params.ShowTime.$D : params.ShowTime.$D
    let hour = params.ShowTime.$d
    let newhour = hour.toString()
    let newriqi = params.ShowTime.$y + '-' + mon + '-' + day + newhour.substring(15, 24)
    console.log(newriqi)
    let { data } = await axios.post('ttms/schedule/add', {
        "PlayId": parseInt(params.PlayId),
        "StudioId": parseInt(params.StudioId),
        "ShowTime": newriqi,
    })
    return data
}


//删除演出厅
export const $delschedule = async (params) => {
    let { data } = await axios.post('ttms/schedule/delete', { "Id": parseInt(params.Id) })
    return data
}

//修改演出厅
export const $reviseschedule = async (params) => {
    let mon = (params.ShowTime.$M + 1) < 10 ? '0' + (params.ShowTime.$M + 1) : params.ShowTime.$M + 1
    let day = (params.ShowTime.$D + 1) < 10 ? '0' + params.ShowTime.$D : params.ShowTime.$D
    let hour = params.ShowTime.$d
    let newhour = hour.toString()
    let newriqi = params.ShowTime.$y + '-' + mon + '-' + day + newhour.substring(15, 24)
    let { data } = await axios.post('ttms/schedule/update', {


        "Id": parseInt(params.Id),
        "PlayId": parseInt(params.PlayId),
        "StudioId": parseInt(params.StudioId),
        "ShowTime": newriqi,
    })
    return data
}