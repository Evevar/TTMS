import React, { useState, useEffect } from 'react'
import { Button, Drawer, Form, Input, message, Modal, DatePicker, Space, TreeSelect } from 'antd'
import { $getOne } from '../../api/studioapi';
import { $addplay, $reviseplay } from '../../api/playapi';
import { ExclamationCircleFilled } from '@ant-design/icons'
import 'dayjs/locale/zh-cn';
import locale from 'antd/es/date-picker/locale/zh_CN';
const { confirm } = Modal;

export default function AddPlay({ open, setOpen, loadList, Id, setId }) {
    const [value, setValue] = useState();
    const onChange = (newValue) => {
        setValue(newValue);
    };
    const treeData = [

        {
            value: '1',
            title: '1',
        },
        {
            value: '2',
            title: '2',
        },
        {
            value: '3',
            title: '3',
        },
    ];

    // const [selectedDate, setselectedDate] = useState('')
    //时间修改

    // const onChange = (date, dateString) => {
    //     // setselectedDate(date, dateString)
    // };
    //提示框
    const [messageApi, contextHolder] = message.useMessage();
    //是否打开抽屉
    // const [open, setOpen] = useState(false);
    //定义表单事例
    let [form] = Form.useForm()
    useEffect(() => {
        if (Id !== 0) {
            $getOne({ Id }).then(data => {
                form.setFieldsValue({ "Id": Id })
            })
        }
    }, [Id])

    //关闭抽屉的方法
    const onClose = () => {
        clear();
        setId(0)
        setOpen(false);
    };
    //表单提交的方法
    const onFinish = async (values) => {

        // values.StartDate.spli console.log(values.StartDate)t(0, 10)
        if (Id) {
            $reviseplay(values).then(data => {
                console.log(values)
                if (data.BaseResp.StatusMessage === 'success') {
                    messageApi.open({
                        type: 'success',
                        content: '修改成功',
                    });
                    loadList()//重新加载数据
                } else {
                    messageApi.open({
                        type: 'error',
                        content: '修改失败',
                    });
                }
            })
        } else {
            let success = await $addplay(values)
            // console.log(success.BaseResp.StatusMessage)
            // .then(success, message => {
            if (success.BaseResp.StatusMessage === 'success') {
                confirm({
                    title: '系统提示',
                    icon: <ExclamationCircleFilled />,
                    content: '添加成功',
                    okText: '确定',
                    cancelText: '取消',
                });
                clear()//清空表单
                loadList()//加载表单
            } else {
                confirm({
                    title: '系统提示',
                    icon: <ExclamationCircleFilled />,
                    content: '添加失败',
                    okText: '确定',
                    cancelText: '取消',
                });
            }
            // })
            // console.log('Success:', values);
        }

    };
    //清空表单方法
    const clear = () => {
        form.setFieldsValue({
            Name: '',
            Type: '',
            Area: '',
            Rating: '',
            Duration: '',
            StartDate: '',
            EndDate: '',
            Price: '',
        })
    }
    return (
        <>
            {contextHolder}
            <Drawer title={Id ? '修改剧目' : '添加剧目'} width={500} placement="right" onClose={onClose} open={open}>
                <Form
                    name="basic"
                    form={form}
                    labelCol={{
                        span: 6,
                    }}
                    wrapperCol={{
                        span: 16,
                    }}
                    style={{
                        maxWidth: 600,
                    }}

                    onFinish={onFinish}
                    autoComplete="off"
                >
                    <Form.Item
                        label="编号"
                        name="Id"
                        hidden
                    >
                        <Input />
                    </Form.Item>
                    <Form.Item
                        label="剧目名称"
                        name="Name"
                        rules={[
                            {
                                required: true,
                                message: '请输入剧目名称!',
                            },
                        ]}
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        label="类型"
                        name="Type"
                        rules={[
                            {
                                required: true,
                                message: '请输入类型!',
                            },
                        ]}
                    >
                        <TreeSelect
                            showSearch
                            style={{
                                width: '100%',
                            }}
                            value={value}
                            dropdownStyle={{
                                maxHeight: 400,
                                overflow: 'auto',
                            }}
                            placeholder="请选择类型"
                            allowClear
                            treeDefaultExpandAll
                            onChange={onChange}
                            treeData={treeData}
                        />
                    </Form.Item>
                    <Form.Item
                        label="地域"
                        name="Area"
                        rules={[
                            {
                                required: true,
                                message: '请输入地域!',
                            },
                        ]}
                    >
                        <Input />
                    </Form.Item>
                    <Form.Item
                        label="剧目等级"
                        name="Rating"
                        rules={[
                            {
                                required: true,
                                message: '请输入剧目等级!',
                            },
                        ]}
                    >
                        <TreeSelect
                            showSearch
                            style={{
                                width: '100%',
                            }}
                            value={value}
                            dropdownStyle={{
                                maxHeight: 400,
                                overflow: 'auto',
                            }}
                            placeholder="请选择剧目等级"
                            allowClear
                            treeDefaultExpandAll
                            onChange={onChange}
                            treeData={treeData}
                        />
                    </Form.Item>
                    <Form.Item
                        label="时长"
                        name="Duration"
                        rules={[
                            {
                                required: true,
                                message: '请输入时长!',
                            },
                        ]}
                    >
                        <Input />
                    </Form.Item>
                    <Form.Item
                        label="上映时间"
                        name="StartDate"
                        rules={[
                            {
                                required: true,
                                message: '请输入上映时间!',
                            },
                        ]}
                    >
                        <DatePicker locale={locale} format="YYYY-MM-DD" onChange={onChange} />
                    </Form.Item>
                    <Form.Item
                        label="结束时间"
                        name="EndDate"
                        rules={[
                            {
                                required: true,
                                message: '请输入结束时间!',
                            },
                        ]}
                    >
                        <DatePicker locale={locale} format="YYYY-MM-DD" onChange={onChange} />
                    </Form.Item>
                    <Form.Item
                        label="票价"
                        name="Price"
                        rules={[
                            {
                                required: true,
                                message: '请输入票价!',
                            },
                        ]}
                    >
                        <Input />
                    </Form.Item>


                    <Form.Item
                        wrapperCol={{
                            offset: 9,
                            span: 16,
                        }}
                    >
                        <Button type="primary" htmlType="submit">
                            {Id ? '修改' : '添加'}
                        </Button>
                        <Button onClick={clear} style={{ marginLeft: '10px' }} >
                            取消
                        </Button>
                    </Form.Item>
                </Form>
            </Drawer>
        </>
    )
}