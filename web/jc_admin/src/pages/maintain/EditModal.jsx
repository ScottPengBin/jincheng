import config from "src/commons/config-hoc";
import {useCallback, useMemo, useState} from "react";
import {ModalContent, FormItem, Content, useDebounceValidator} from '@ra-lib/admin';
import {Button, Card, Col, Form, Row} from "antd";


export default config({
    modal: {
        title: (props) => {
            if (props?.record?.isDetail) return '查看记录';

            return props.isEdit ? '编辑会员' : '创建会员';
        },
        width: '70%',
        top: 50,
    },
})(function Edit(props) {
    const {record, isEdit, onOk, onCancel} = props;
    const [loading, setLoading] = useState(false);
    const [form] = Form.useForm();
    const isDetail = record?.isDetail;

    const params = useMemo(() => {
        return {id: record?.id};
    }, [record]);
    // 编辑时，查询详情数据
    props.ajax.useGet('/member/getOne', params, [params], {
        mountFire: isEdit,
        setLoading,
        formatResult: (res) => {
            if (!res) return;
            form.setFieldsValue(res);
        },
    });

    const {run: save} = props.ajax.usePost('/member/add', null, {setLoading, successTip: '创建成功！'});
    const {run: update} = props.ajax.usePost('/member/updateMemberById', null, {setLoading, successTip: '修改成功！'});
    const {run: fetchMemberByName} = props.ajax.useGet('/member/getOneByMobile');
    const handleSubmit = useCallback(
        async (values) => {
            const roleIds = values.roleIds?.filter((id) => !`${id}`.startsWith('systemId'));
            const params = {
                ...values,
                roleIds,
            };

            if (isEdit) {
                await update(params);
            } else {
                await save(params);
            }

            onOk();
        },
        [isEdit, update, save, onOk],
    );

    const disabled = isDetail;
    const layout = {
        labelCol: {flex: '100px'},
        disabled,
    };
    const colLayout = {
        xs: {span: 24},
        sm: {span: 12},
    };


    const checkMemberMobile = useDebounceValidator(async (rule, value) => {
        if (!value) return;

        const systemId = form.getFieldValue('systemId');
        const role = await fetchMemberByName({mobile: value, systemId});
        if (!role) return;

        const id = form.getFieldValue('member')?.id;
        if (isEdit && role.id !== id && role.mobile === value) throw Error('电话号码不能重复！');
        if (!isEdit && role.mobile === value) throw Error('电话号码不能重复！');
    });
    return (
        <Form form={form} name="member" onFinish={handleSubmit} initialValues={{enabled: true}}>
            <ModalContent
                loading={loading}
                okText="保存"
                okHtmlType="submit"
                cancelText="重置"
                onCancel={() => form.resetFields()}
                footer={disabled ? <Button onClick={onCancel}>关闭</Button> : undefined}
            >
                {isEdit ? <FormItem hidden name={["member", "id"]}/> : null}
                {isEdit ? <FormItem hidden name={["car_info", "id"]}/> : null}
                <Row gutter={8}>
                    <Col {...colLayout}>
                        <Card title="基础信息">
                            <Content fitHeight otherHeight={160}>
                                <FormItem {...layout} label="姓名" name={["member", "member_name"]} required noSpace/>
                                <FormItem {...layout} label="电话" name={["member", "mobile"]}
                                          rules={[{validator: checkMemberMobile}]}
                                          required noSpace/>
                                <FormItem {...layout} label="生日" type={"date"} dateFormat={"YYYY-MM-DD"} name={["member", "brith_day"]}
                                          required noSpace/>
                                <FormItem {...layout} label="性别" name={["member", "gender"]} type="select" options={[
                                    {value: '男', label: '男'},
                                    {value: '女', label: '女'},
                                ]} required noSpace/>
                                <FormItem {...layout} label="备注" type={"textarea"} name={["member", "member_note"]}
                                          noSpace/>
                            </Content>
                        </Card>
                    </Col>
                    <Col {...colLayout}>
                        <Card title="车辆信息" >
                            <Content fitHeight otherHeight={160}>
                                <FormItem {...layout} label="车牌号" name={["car_info", "car_no"]} required noSpace/>
                                <FormItem {...layout} label="车辆名称" name={["car_info", "car_name"]} required
                                          noSpace/>
                                <FormItem {...layout} label="车辆颜色" name={["car_info", "car_color"]} required
                                          noSpace/>
                                <FormItem {...layout} label="车辆备注" type={"textarea"} name={["car_info", "car_note"]}
                                          noSpace/>
                            </Content>
                        </Card>
                    </Col>
                </Row>
            </ModalContent>
        </Form>
    );
})
