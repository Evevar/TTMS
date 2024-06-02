import React, { useEffect, useState } from 'react'
import { Table, Button, Drawer, Form, Input, message, Popconfirm, Pagination } from 'antd'
import { $play, $delplay } from '../../api/playapi';
import { keyboard } from '@testing-library/user-event/dist/keyboard';
import { ExclamationCircleFilled } from '@ant-design/icons'
import AddPlay from './AddPlay';


export default function Play() {
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
        $delplay({ Id }).then(data => {
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
            title: '类型',
            dataIndex: 'Type',
            width: "100px",
        },
        {
            title: '地域',
            dataIndex: 'Area',
            width: "100px",
        },
        {
            title: '剧目等级',
            dataIndex: 'Rating',
            width: "100px",
        },
        {
            title: '时长',
            dataIndex: 'Duration',
            width: "100px",
        },
        {
            title: '上映日期',
            dataIndex: 'StartDate',
            width: "100px",
        },
        {
            title: '结束日期',
            dataIndex: 'EndDate',
            width: "100px",
        },
        {
            title: '价格',
            dataIndex: 'Price',
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
                        <Button style={{ marginLeft: '1px' }} danger size='small' >删除</Button>
                    </Popconfirm>
                </>
            ),
        },
    ];
    const loadList = () => {
        $play(Current).then((data, count) => {
            count = data.Data.Total
            let dataList = data.Data.List
            // console.log(dataList)
            dataList = dataList.map(r => {
                return {
                    ...r,
                    key: r.Id,
                }
            })

            setStudioList(dataList)
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
            <Pagination size='small' pageSize={10} defaultCurrent={Current} total={count} onChange={(page) => {
                setCurrent(page)
            }}
            />
            <AddPlay open={open} setOpen={setOpen} loadList={loadList} Id={Id} setId={setId} />
        </>
    )
}
