import React, { useEffect, useState } from 'react'
import { Table, Button, Drawer, Form, Input, message, Popconfirm, Pagination } from 'antd'
import { $studio, $del } from '../../api/studioapi';
import { keyboard } from '@testing-library/user-event/dist/keyboard';
import { ExclamationCircleFilled } from '@ant-design/icons'
import AddStudio from './AddStudio';


export default function Studio() {
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
    const del = (Id) => {
        $del({ Id }).then(data => {
            console.log(data.BaseResp.StatusMessage);
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
    }, [Current])
    const columns = [
        {
            title: '编号',
            dataIndex: 'Id',
            width: "100px",
        },
        {
            title: '名称',
            dataIndex: 'Name',
            width: "100px",
        },
        {
            title: '行数',
            dataIndex: 'RowsCount',
            width: "100px",
        },
        {
            title: '列数',
            dataIndex: 'ColsCount',
            width: "100px",
        },
        {
            title: '座位总数',
            dataIndex: 'SeatsCount',
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
                            del(ret.Id)
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
        $studio(Current).then((data, count) => {
            count = data.length
            // let dataList = data.List
            // console.log(dataList)
            data = data.map(r => {
                return {
                    ...r,
                    key: r.Id
                }
            })

            setStudioList(data)
            // const talcut = parseInt(data.length / 10 + 1)
            setCount(count)
            // console.log(count)
        })
    };

    return (
        <> {contextHolder}
            <div className='search'>
                <Button size='small' onClick={() => { setOpen(true) }} >添加</Button>
            </div>
            <Table size='small' dataSource={studioList} columns={columns} pagination={false} />
            <Pagination size='small' pageSize={10} defaultCurrent={Current} total={20} onChange={(page) => {
                setCurrent(page)
            }}
            />
            <AddStudio open={open} setOpen={setOpen} loadList={loadList} Id={Id} setId={setId} />
        </>
    )
}
