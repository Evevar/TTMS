//图标导入
import {
    MenuFoldOutlined,
    UploadOutlined,
    UserOutlined,
    VideoCameraOutlined,
    AppstoreOutlined, MailOutlined, SettingOutlined, ExclamationCircleFilled
} from '@ant-design/icons';
import { MenuUnfoldOutlined } from '@ant-design/icons';
import { Layout, Menu, Modal } from 'antd';
import React, { useState, useEffect } from 'react';
import { useNavigate, Outlet, Navigate } from 'react-router-dom'
const { confirm } = Modal;
const { Header, Sider, Content } = Layout;
import './Layout.scss'


export default function () {

    const navigate = useNavigate()
    useEffect(() => {
        if (!sessionStorage.getItem('token')) {
            navigate('/')
        }
    }, [])
    const [current, setCurrent] = useState('mail');
    //顶菜单横向
    const items = [
        // {
        //     label: '首页',
        //     key: 'mail',
        //     icon: <MailOutlined />,
        // },
        {
            label: '个人中心',
            key: 'mine',
            icon: <SettingOutlined />,
            children: [
                {
                    key: 'my',
                    label: '个人信息',
                },
                {
                    key: 'exit',
                    label: '退出登陆',
                },
            ],
        },
        {
            label: (
                <a href='http://116.204.90.27:82/' target="_blank" rel="noopener noreferrer">
                    进入购票页面
                </a>
            ),
            key: 'alipay',
        },
    ];
    //左侧菜单项
    const items2 = [
        {
            key: 'studio',
            icon: <UploadOutlined />,
            label: '演出厅管理',
        },
        {
            key: 'seat',
            icon: <UserOutlined />,
            label: '座位管理',
        },
        {
            key: 'play',
            icon: <VideoCameraOutlined />,
            label: '剧目管理',
        },
        {
            key: 'schedule',
            icon: <UploadOutlined />,
            label: '演出计划管理',
        },
        {
            key: 'order',
            icon: <UploadOutlined />,
            label: '统计票房',
        },

    ];
    //点击菜单方法
    const OnclickMeu = (e) => {
        setCurrent(e.key)
        switch (e.key) {
            case 'order':
                navigate('/layout/order')
                break
            case 'schedule':
                navigate('/layout/schedule')
                break
            case 'play':
                navigate('/layout/play')
                break
            case 'seat':
                navigate('/layout/seat')
                alert("请先选择所需查询演出厅id")
                break
            case 'studio':
                navigate('/layout/studio')
                break
            //退出登陆
            case 'exit':
                confirm({
                    title: '系统提示',
                    icon: <ExclamationCircleFilled />,
                    content: '确定退出系统吗？',
                    okText: '确定',
                    cancelText: '取消',
                    onOk() {
                        sessionStorage.clear()
                        localStorage.clear()
                        navigate('/')
                    },
                });
                break
        }
    }
    //侧边栏折叠状态
    const [collapsed, setCollapsed] = useState(false);

    return (
        <Layout className='layout'>
            <Sider trigger={null} collapsible collapsed={collapsed}>
                <div className="logo" > {collapsed ? '长' : '长乐剧院票务管理系统'}</div>
                <Menu
                    onClick={OnclickMeu}
                    theme="dark"
                    mode="inline"
                    defaultSelectedKeys={['1']}
                    items={items2}
                />
            </Sider>
            <Layout >
                <Header className='header'>
                    {React.createElement(collapsed ? MenuUnfoldOutlined : MenuFoldOutlined, {
                        className: 'trigger',
                        onClick: () => setCollapsed(!collapsed),
                    })}
                    <Menu onClick={OnclickMeu} theme='dark' className='menu' selectedKeys={[current]} mode="horizontal" items={items} />
                </Header>
                <Content className='content'>
                    <Outlet></Outlet>
                </Content>
            </Layout>
        </Layout>
    );
};

