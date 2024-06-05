import React, { useEffect, useState } from 'react'
import { Table, Button, Drawer, Form, Input, message, Popconfirm, Pagination, Select } from 'antd'
import { $seat, $delseat } from '../../api/seatapi';

import { keyboard } from '@testing-library/user-event/dist/keyboard';
import { ExclamationCircleFilled } from '@ant-design/icons'
import AddSeat from './AddSeat';
import { $studio } from '../../api/studioapi';


export default function Seat() {
    // $studio().then((data) => {
    //     // const selectelement = document.getElementsByClassName('myselect')
    //     let getId = []
    //     for (var a = 0; a <= 9; a++) {
    //         getId.push({ 'value': data[a].Name, 'label': data[a].Id })
    //     }
    //     // selectelement.options
    //     console.log(getId)
    // })
    // console.log(getId)
    let [selectchage, setselectchage] = useState(1)
    //座位演出厅选择
    const handleChange = async (value) => {
        // console.log(`selected ${value}`);
        loadList()
        setselectchage(parseInt(`${value}`))
    };

    //总数量
    let [count, setCount] = useState(1)
    //分页
    let [Current, setCurrent] = useState(1)
    //编辑状态Id
    let [Id, setId] = useState(0)
    //提示框
    const [messageApi, contextHolder] = message.useMessage();
    // 是否打开抽屉
    const [open, setOpen] = useState(false);
    //编辑方法
    const edit = (Id) => {
        setOpen(true)//打开抽屉
        setId(Id)//设置为编辑状态
    }
    //删除方法
    const del = (StudioId, Row, Col) => {
        const deldata = { StudioId, Row, Col }

        $delseat({ deldata }).then(data => {
            // console.log(deldata)
            // console.log(data.BaseResp.StatusMessage);
            if (data.BaseResp.StatusMessage === 'success') {
                messageApi.open({
                    type: 'success',
                    content: '删除成功',
                });
                loadList()//重新加载数据
            } else {
                messageApi.open({
                    type: 'error',
                    content: '删除失败',
                });
            }
        })
    }
    //角色列表数据
    let [studioList, setStudioList] = useState([])
    // let studiodata = $studio()
    // console.log(studiodata)
    useEffect(() => {
        loadList()
    }, [Current, selectchage])
    const columns = [
        {
            title: '编号',
            dataIndex: 'Id',
            width: "100px",
        },
        {
            title: '演出厅ID',
            dataIndex: 'StudioId',
            width: "100px",
        },
        {
            title: '某行',
            dataIndex: 'Row',
            width: "100px",
        },
        {
            title: '某列',
            dataIndex: 'Col',
            width: "100px",
        },
        {
            title: '座位状态',
            dataIndex: 'Status',
            width: "100px",
        },
        {
            title: '操作',
            key: 'action',
            width: '100px',

            render: (ret) => (
                <>
                    <Button size='small' style={{ borderColor: 'green', color: 'green' }} onClick={() => {
                        edit(ret.Id)
                    }}>编辑</Button>
                    <Popconfirm
                        title="提示"
                        description="确定删除吗?"
                        onConfirm={() => {
                            del(ret.StudioId, ret.Row, ret.Col)
                        }}
                        okText="确定"
                        cancelText="取消"
                    >
                        <Button style={{ marginLeft: '15px' }} danger size='small' >删除</Button>
                    </Popconfirm>
                </>
            ),
        },
    ];
    const loadList = () => {

        // const getselect = document.getElementsByClassName('myselect')
        const selectdata = { selectchage, Current }
        $seat(selectdata).then((data, count) => {
            count = data.Data.Total
            let dataList = data.Data.List
            dataList = dataList.map(r => {
                return {
                    ...r,
                    key: r.Id
                }
            })

            setStudioList(dataList)
            // const talcut = parseInt(data.length / 10 + 1)
            setCount(count)

        })
    };

    return (
        <> {contextHolder}
            <Select
                className='myselect'
                defaultValue="1"
                size='small'
                style={{
                    float: 'left',
                    width: 150,
                }}
                onChange={handleChange}
                options={[
                    // {
                    //     value: 'jack',
                    //     label: 12,
                    // },
                    { value: 4, label: 4 },

                    { value: 5, label: 5 },

                    { value: 6, label: 6 },

                    { value: 25, label: 25 },

                    { value: 26, label: 26 },

                    { value: 27, label: 27 },

                    { value: 28, label: 28 },

                    { value: 29, label: 29 },

                    { value: 30, label: 30 },

                    { value: 31, label: 31 }
                ]}
            />


            <div className='search'>
                <Button size='small' onClick={() => { setOpen(true) }} >添加</Button>
            </div>
            <Table size='small' dataSource={studioList} columns={columns} pagination={false} />
            <Pagination size='small' defaultCurrent={Current} total={count} onChange={(page) => {
                setCurrent(page)
            }}
            />
            <AddSeat open={open} setOpen={setOpen} loadList={loadList} Id={Id} setId={setId} />
        </>
    )
}
