import axios from "../utils/request";

//演出厅列表
export const $studio = async (params) => {
    let { data } = await axios.get('ttms/studio/all', {
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
export const $addstudio = async (params) => {
    let { data } = await axios.post('ttms/studio/add', {
        "Name": params.Name,
        "RowsCount": parseInt(params.RowsCount),
        "ColsCount": parseInt(params.ColsCount),
    })
    return data
}


//删除演出厅
export const $del = async (params) => {
    let { data } = await axios.post('ttms/studio/delete', { "Id": parseInt(params.Id) })
    return data
}

//修改演出厅
export const $revise = async (params) => {
    let { data } = await axios.post('ttms/studio/update', {
        "Id": parseInt(params.Id),
        "Name": params.Name,
        "RowsCount": parseInt(params.RowsCount),
        "ColsCount": parseInt(params.ColsCount),
    })
    return data
}

//获取单个演出厅
export const $getOne = async (params) => {
    let { data } = await axios.get('ttms/studio/info', {
        params: {
            Id: parseInt(params.Id)
        }
    })
    return data.result
}