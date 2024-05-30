import axios from "../utils/request";

//演出厅列表
export const $play = async (params) => {
    let { data } = await axios.get('ttms/play/all/', {
        params: {
            Current: params,
            PageSize: 10
        }
    })
    // sessionStorage.setItem('token', success.Token);
    // console.log(data)
    return data.List
    // return data
}

//添加演出厅
export const $addplay = async (params) => {
    let mon = (params.StartDate.$M + 1) < 10 ? '0' + (params.StartDate.$M + 1) : params.StartDate.$M + 1
    let day = (params.StartDate.$D + 1) < 10 ? '0' + params.StartDate.$D : params.StartDate.$D
    let newriqi = params.StartDate.$y + '-' + mon + '-' + day

    let endmon = (params.EndDate.$M + 1) < 10 ? '0' + (params.EndDate.$M + 1) : params.EndDate.$M + 1
    let endday = (params.EndDate.$D + 1) < 10 ? '0' + params.EndDate.$D : params.EndDate.$D
    let endnewriqi = params.EndDate.$y + '-' + endmon + '-' + endday
    let { data } = await axios.post('ttms/play/add/', {
        "Name": params.Name,
        "Type": parseInt(params.Type),
        "Area": params.Area,
        "Rating": parseInt(params.Rating),
        "Duration": params.Duration,
        "StartDate": newriqi,
        "EndDate": endnewriqi,
        "Price": parseInt(params.Price),
    })
    return data
}


//删除演出厅
export const $delplay = async (params) => {
    let { data } = await axios.post('ttms/play/delete/', { "Id": parseInt(params.Id) })
    return data
}

//修改演出厅
export const $reviseplay = async (params) => {
    let mon = (params.StartDate.$M + 1) < 10 ? '0' + (params.StartDate.$M + 1) : params.StartDate.$M + 1
    let day = (params.StartDate.$D + 1) < 10 ? '0' + params.StartDate.$D : params.StartDate.$D
    let newriqi = params.StartDate.$y + '-' + mon + '-' + day

    let endmon = (params.EndDate.$M + 1) < 10 ? '0' + (params.EndDate.$M + 1) : params.EndDate.$M + 1
    let endday = (params.EndDate.$D + 1) < 10 ? '0' + params.EndDate.$D : params.EndDate.$D
    let endnewriqi = params.EndDate.$y + '-' + endmon + '-' + endday
    // console.log(endnewriqi)
    let { data } = await axios.post('ttms/play/update/', {
        "Id": parseInt(params.Id),
        "Name": params.Name,
        "Type": parseInt(params.Type),
        "Area": params.Area,
        "Rating": parseInt(params.Rating),
        "Duration": params.Duration,
        "StartDate": newriqi,
        "EndDate": endnewriqi,
        "Price": parseInt(params.Price),
    })
    return data
}


export const $order = async (params) => {
    let { data } = await axios.get('ttms/order/analysis/', {
        params: {
            PlayId: parseInt(params.PlayId)
        }
    })
    // let newdata = Array.from(data.OrderAnalysis);
    // sessionStorage.setItem('token', success.Token);
    // console.log(newdata)
    return data.OrderAnalysis
    // return data
}