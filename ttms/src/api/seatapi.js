import axios from "../utils/request";


export const $seat = async (value) => {

    // console.log(value)
    let { data } = await axios.get('ttms/seat/all', {
        params: {
            StudioId: value.selectchage,
            Current: value.Current,
            PageSize: 10
        }
    })
    // sessionStorage.setItem('token', success.Token);
    // console.log(data)
    return data
    // return data
}


export const $delseat = async (params) => {
    // console.log(params.deldata)
    let { data } = await axios.post('ttms/seat/delete', {
        "StudioId": parseInt(params.deldata.StudioId),
        "Row": parseInt(params.deldata.Row),
        "Col": parseInt(params.deldata.Col)
    })
    return data
}


export const $addseat = async (params) => {
    let { data } = await axios.post('ttms/seat/add', {
        "StudioId": parseInt(params.StudioId),
        "Row": parseInt(params.Row),
        "Col": parseInt(params.Col),
        "Status": parseInt(params.Status),
    })
    return data
}


export const $reviseseat = async (params) => {
    let { data } = await axios.post('ttms/seat/update', {
        "StudioId": parseInt(params.StudioId),
        "Row": parseInt(params.Row),
        "Col": parseInt(params.Col),
        "Status": parseInt(params.Status),
    })
    return data
}
