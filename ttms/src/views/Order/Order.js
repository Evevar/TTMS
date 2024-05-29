import React from 'react'
import { $order } from '../../api/playapi'
import { Button, Checkbox, Form, Input, Card, message } from 'antd';


export default function Order() {
    const [messageApi, contextHolder] = message.useMessage();
    const onFinish = async (values) => {
        let success = await $order(values)
        if (success) {
            // console.log(success.PlayId)
            if (success.PlayId) {
                document.getElementById('top').innerHTML = "剧目ID:" + success.PlayId
            }
            if (success.PlayName) {
                document.getElementById('op').innerHTML = "剧目名称：" + success.PlayName
            }
            if (success.PlayArea) {
                document.getElementById('to').innerHTML = "演出地址:" + success.PlayArea
            }
            if (success.PlayDuration) {
                document.getElementById('so').innerHTML = "演出时长:" + success.PlayDuration
            }
            if (success.StartData) {
                document.getElementById('fo').innerHTML = "开始放映日期:" + success.StartData
            }
            if (success.EndData) {
                document.getElementById('fio').innerHTML = "结束放映日期:" + success.EndData
            }
            if (success.TotalTicket) {
                document.getElementById('so').innerHTML = "总票数:" + success.TotalTicket
            }
            if (success.Sales) {
                document.getElementById('seo').innerHTML = "总售出:" + success.Sales
            }
            if (success.Price) {
                document.getElementById('eo').innerHTML = "价格:" + success.Price
            }


        } else {
            messageApi.open({
                type: 'error',
                content: '该剧目不存在',
            });
        }

    };

    return (
        <>
            {contextHolder}
            <Form
                name="basic"
                labelCol={{
                    span: 8,
                }}
                wrapperCol={{
                    span: 16,
                }}
                style={{

                    maxWidth: 600,
                }}
                initialValues={{
                    remember: true,
                }}
                onFinish={onFinish}
                // onFinishFailed={onFinishFailed}
                autoComplete="off"
            >
                <Form.Item
                    label="剧目ID"
                    name="PlayId"
                    rules={[
                        {
                            required: true,
                            message: '请输入剧目ID!',
                        },
                    ]}
                    style={{ textAlign: 'center', marginLeft: 230 }}
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    wrapperCol={{
                        offset: 8,
                        span: 16,
                    }}
                >
                    <Button style={{ marginLeft: 100 }} type="primary" htmlType="submit">
                        提交
                    </Button>
                </Form.Item>
                <Card
                    title="统计票房"
                    // extra={<a href="#">More</a>}
                    style={{
                        width: 500,
                        height: 500,
                        marginLeft: 300,
                    }}
                >
                    <p id='top'>暂无内容</p>
                    <p id='op'></p>
                    <p id='to'></p>
                    <p id='so'></p>
                    <p id='fo'></p>
                    <p id='fio'></p>
                    <p id='so'></p>
                    <p id='seo'></p>
                    <p id='eo'></p>

                </Card>
            </Form>


        </>
    )
}
