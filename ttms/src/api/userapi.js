import axios from "../utils/request";
// import md5 from 'md5'

//登录
// 
// export const $login = async (params) => {

export const $login = async (values) => {

    //导航
    // params.Password = md5(md5(params.Password).split('').reverse().join(''))
    let { data } = await axios.post('ttms/user/login/', { UserId: parseInt(values.UserId), Password: values.Password })
    // let { data } = await axios.post('ttms/user/login', { params })
    // if (data.BaseResp.StatusMessage) {
    //     // console.log(data.Token)

    // }
    return data
}
