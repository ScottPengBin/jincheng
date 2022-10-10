import {useState, useCallback, useMemo} from 'react';
import config from 'src/commons/config-hoc';
import {
    PageContent,
    QueryBar,
    FormItem,
    ToolBar,
    Table,
    Pagination, Operator,
} from '@ra-lib/admin';
import {Form, Space, Button} from 'antd';

export default config({
    path: '/member',
})(function MemberManager(props) {
    const [loading, setLoading] = useState(false);
    const [pageNum, setPageNum] = useState(1);
    const [pageSize, setPageSize] = useState(20);
    const [conditions, setConditions] = useState({});
    const [form] = Form.useForm();

    const params = useMemo(() => {
        return {
            ...conditions,
            pageNum,
            pageSize,
        };
    }, [conditions, pageNum, pageSize]);


    // 使用现有查询条件，重新发起请求
    const refreshSearch = useCallback(() => {
        setConditions(form.getFieldsValue());
    }, [form]);

    //列表
    const {data: {dataSource, total} = {}} = props.ajax.useGet('/member/getList', params, [params], {
        setLoading,
        formatResult: (res) => {

            return {
                dataSource: res?.records || [],
                total: res?.total || 0,
            };
        },
    });
    const columns = [
        {title: '姓名', dataIndex: 'member_name'},
        {title: '电话', dataIndex: 'mobile'},
        {title: '生日', dataIndex: 'brith_day'},
        {title: '年龄', dataIndex: 'age'},
        {title: '性别', dataIndex: 'gender'},
        {title: '创建时间', dataIndex: 'created_at'},
        {title: '更新时间', dataIndex: 'update_at'},
        {title: '备注', dataIndex: 'member_note'},
        {
            title: '操作',
            dataIndex: 'operator',
            render: (value, record) => {
                const {id, name} = record;
                const items = [
                    {
                        label: '查看',
                    },
                    {
                        label: '修改',
                    },
                    {
                        label: '删除',
                        color: 'red',
                        confirm: {
                            title: `您确定删除「${name}」吗？`,
                            onConfirm: () => handleDelete(id),
                        },
                    },
                ];

                return <Operator items={items}/>;
            },
        },
    ];

    // 删除
    const {run: deleteRecord} = props.ajax.useDel('/member/delete', {id: ':id'}, {setLoading, successTip: '删除成功！'});
    const handleDelete = useCallback(
        async (id) => {
            await deleteRecord(id);
            // 触发列表更新
            refreshSearch();
        },
        [deleteRecord, refreshSearch],
    );


    const handleSubmit = useCallback((values) => {
        console.log(values)

    }, []);


    const handlePageNumChange = useCallback((pageNum) => {
        setPageNum(pageNum);
    }, []);

    const handlePageSizeChange = useCallback((pageSize) => {
        setPageNum(1);
        setPageSize(pageSize);
    }, []);

    return (
        <PageContent loading={loading}>
            <QueryBar>
                <Form layout="inline"  name="member" form={form} onFinish={(values) => setPageNum(1) || setConditions(values)}>
                    <FormItem label="姓名" name="name"/>
                    <FormItem label="电话" name="mobile"/>
                    <FormItem label="添加日期" name="created_at" type="date"/>
                    <FormItem>
                        <Space>
                            <Button type="primary" htmlType="submit">
                                提交
                            </Button>
                            <Button>重置</Button>
                        </Space>
                    </FormItem>
                </Form>
            </QueryBar>
            <ToolBar>
                <Button type="primary">添加</Button>
                <Button type="primary" danger>
                    批量删除
                </Button>
                <Button>导出</Button>
            </ToolBar>
            <Table
                pagination={false}
                dataSource={dataSource}
                rowKey="id"
                columns={columns}
            />
            <Pagination
                total={total || 0}
                pageNum={pageNum}
                pageSize={pageSize}
                onPageNumChange={handlePageNumChange}
                onPageSizeChange={handlePageSizeChange}
            />
        </PageContent>
    );
});
