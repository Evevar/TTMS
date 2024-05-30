import React, { useState, useEffect } from 'react'
import './Login.scss'
import { Button, Form, Input, notification } from 'antd';
import { $login } from '../../api/userapi'
import { Navigate, useNavigate } from 'react-router-dom'
// import MyNotification from '../../components/MyNotification/MyNotification';

export default function Login() {
    //判断是否已登陆成功
    let navigate = useNavigate()

    // // //通知框状态
    // let [notiMsg, setNotiMsg] = useState({ type: '', description: '' })
    // //提示框
    // const [api, contextHolder] = notification.useNotification();
    // //打开提示框
    // const openNotification = (type, description) => {
    //     api[type]({
    //         message: '系统提示',
    //         description
    //     });
    // };
    //表单
    let [form] = Form.useForm()
    //表单成功提交   
    const onFinish = async (values) => {
        // values = JSON.stringify({ "UserId": 1, "Password": "z3" });
        let success = await $login(values)


        // console.log()
        if (success.BaseResp.StatusMessage === 'success') {
            // openNotification('success', message)
            // setNotiMsg({ type: 'success', description: message })
            // alert("登录成功");
            sessionStorage.setItem('token', success.Token);
            // sessionStorage.setItem('Authorization', success.Token);
            navigate('/layout/studio')
        }
        // if (success.BaseResp.StatusMessage != 'success')
        else {
            // openNotification('error', message)
            // 
            alert(success.BaseResp.StatusMessage);
        }
        // // console.log('Success:', values);
        // console.log(values.UserId)
    };
    useEffect(() => {
        if (sessionStorage.getItem('token')) {
            navigate('/layout')
        }
    }, [])

    return (

        <div className='login'>
            <div className='content'>
                <h2>长乐剧院票务管理系统</h2>
                <Form
                    name="basic"
                    form={form}
                    labelCol={{
                        span: 4,
                    }}
                    wrapperCol={{
                        span: 18,
                    }}
                    initialValues={{
                        UserId: 7,
                        Password: 'z3'
                    }}
                    onFinish={onFinish}
                    autoComplete="off"
                >
                    <Form.Item
                        label="账号"
                        name="UserId"
                        rules={[
                            {

                                required: true,
                                message: '请输入账号',
                            },
                        ]}
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        label="密码"
                        name="Password"
                        rules={[
                            {
                                required: true,
                                message: '请输入密码',
                            },
                        ]}
                    >
                        <Input.Password />
                    </Form.Item>



                    <Form.Item
                        wrapperCol={{
                            offset: 6,
                            span: 16,
                        }}
                    >
                        <Button type="primary" htmlType="submit">
                            登录
                        </Button>
                        <Button onClick={() => {
                            form.resetFields()
                        }} style={{ marginLeft: '125px' }} >
                            取消
                        </Button>
                    </Form.Item>
                </Form>
            </div>
            {/* <MyNotification notiMsg={notiMsg} /> */}
        </div>
    )
}
