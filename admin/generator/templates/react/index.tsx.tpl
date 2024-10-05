import { Button, Form, Input, Modal, Popconfirm } from 'antd';
import { Table } from 'antd';
import type { TableProps } from 'antd';
import { useEffect, useState } from 'react';
import * as Icons from '@ant-design/icons';
import api, {
  {{{.EntityName}}}ListParams,
  {{{.EntityName}}}DataTableItem,
  {{{.EntityName}}}CreateBody,
  {{{.EntityName}}}UpdateBody
} from '@/api{{{.GenPath}}}';

// 类型定义
type TableColumns = TableProps<{{{.EntityName}}}DataTableItem>['columns'];

type TableRowSelection<T extends object = object> =
  TableProps<T>['rowSelection'];

// 主函数
export default function {{{.EntityName}}}Page() {
  // 表单属性定义
  const [searchForm] = Form.useForm();
  const [dataForm] = Form.useForm();
  const [editRecord, setEditRecord] = useState<any>(null);
  const [showDataModal, setShowDataModal] = useState<boolean>(false);
  // 数据表格属性定义
  const [tableData, setTableData] = useState<any>([]);
  const [selectedRowKeys, setSelectedRowKeys] = useState([]);
  // 请求参数定义
  const [params, setParams] = useState<{{{.EntityName}}}ListParams>({
    page: 1,
    limit: 10
  });

  // 事件定义
  // 复选框事件
  const onSelectChange = newSelectedRowKeys => {
    console.log('selectedRowKeys changed: ', newSelectedRowKeys);
    setSelectedRowKeys(newSelectedRowKeys);
  };

  const rowSelection: TableRowSelection<{{{.EntityName}}}DataTableItem> = {
    selectedRowKeys,
    onChange: onSelectChange
  };

  // 查询
  const onSearch = () => {};

  // 详情
  const onDetail = (record: {{{.EntityName}}}DataTableItem) => {
    detail(record.id);
  };

  // 创建
  const onCreate = () => {
    showDataFormModal();
  };

  // 更新
  const onUpdate = (record: {{{.EntityName}}}DataTableItem) => {
    showDataFormModal(record);
  };

  // 删除
  const onDelete = (record: {{{.EntityName}}}DataTableItem) => {
    del([record.id]);
  };

  // 批量删除
  const onBatchDelete = () => {
    const ids = selectedRowKeys;
    del(ids);
  };

  // 弹出层事件
  const showDataFormModal = (record?: {{{.EntityName}}}DataTableItem) => {
    setEditRecord(record);
    dataForm.resetFields();
    if (record) {
      dataForm.setFieldsValue({ ...record });
    }
    setShowDataModal(true);
  };

  const handleDataFormOk = () => {
    if (editRecord) {
      update();
    } else {
      create();
    }
    setShowDataModal(false);
  };

  const handleDataFormCancel = () => {
    setShowDataModal(false);
  };

  // http请求
  // 获取数据列表
  const list = () => {
    api
      .list(params)
      .then(res => {
        setTableData(res.result);
      })
      .catch(err => {
        console.log(err);
      });
  };

  // 详情
  const detail = (id: number) => {
    api
      .detail(id)
      .then(res => {
        console.log(res);
      })
      .catch(err => {
        console.log(err);
      });
  };

  // 新增
  const create = () => {
    api
      .create({ ...dataForm.getFieldsValue() })
      .then(res => {
        console.log(res);
        list();
      })
      .catch(err => {
        console.log(err);
      });
  };

  // 更新
  const update = () => {
    api
      .update({ ...dataForm.getFieldsValue() })
      .then(res => {
        console.log(res);
        list();
      })
      .catch(err => {
        console.log(err);
      });
  };

  // 删除
  const del = (ids: Array<number>) => {
    api
      .delete({ ids })
      .then(res => {
        console.log(res);
      })
      .catch(err => {
        console.log(err);
      });
  };

  useEffect(() => {
    list();
  }, []);

  // 列内容定义
  const columns: TableColumns = [
    {{{- range .Columns }}}
    {{{- if .IsList }}}
    {
      title: '{{{ .ColumnComment }}}',
      dataIndex: '{{{ .GoField }}}',
      key: '{{{ .ColumnComment }}}'
    },
    {{{- end }}}
    {{{- end }}}
    {
      title: '操作',
      key: 'id',
      dataIndex: 'id',
      render: (_, record) => (
        <>
          <div
            className="flex flex-row justify-start gap-3 text-blue-500"
            key={'action_' + record.id}
          >
            <a
              onClick={() => {
                onDetail(record);
              }}
            >
              <Icons.FileTextOutlined />
              详情
            </a>
            <a
              onClick={() => {
                onUpdate(record);
              }}
            >
              <Icons.EditOutlined />
              编辑
            </a>
            <Popconfirm
              title="删除"
              description="确定要删除当前数据吗？"
              onConfirm={() => {
                onDelete(record);
              }}
              okText="确定"
              cancelText="取消"
            >
              <a>
                <Icons.DeleteOutlined className="mr-1" />
                删除
              </a>
            </Popconfirm>
          </div>
        </>
      )
    }
  ];

  return (
    <>
      <div className="un-box w-full h-full bg-white p-4">
        <div className="w-full flex flex-row border-b border-b-gray justify-between pb-4 mb-5">
          <Form layout="inline" form={searchForm}>
            {{{- range .Columns }}}
            <Form.Item label="{{{.ColumnComment}}}" name="{{{.ColumnName}}}">
              <Input
                placeholder="输入{{{.ColumnComment}}}查询"
                allowClear
                onClear={list}
              />
            </Form.Item>
            {{{- end }}}
            <Form.Item>
              <Button type="default" onClick={onSearch}>
                查询
              </Button>
            </Form.Item>
          </Form>
          <div className="flex flex-row gap-5">
            <Button type="primary" onClick={onCreate}>
              <Icons.PlusOutlined />
              新增
            </Button>
            <Button type="default" danger onClick={onBatchDelete}>
              <Icons.DeleteOutlined />
              批量删除
            </Button>
          </div>
        </div>
        <div className="">
          <Table
            size="middle"
            columns={columns}
            rowSelection={rowSelection}
            rowKey={record => record.id}
            dataSource={tableData}
          />
        </div>
      </div>

      <Modal
        title={editRecord ? '编辑数据' : '新增数据'}
        centered={true}
        maskClosable={false}
        open={showDataModal}
        onOk={handleDataFormOk}
        onCancel={handleDataFormCancel}
        width={600}
      >
        <Form
          form={dataForm}
          name="data-form"
          labelCol={{ span: 5 }}
          wrapperCol={{ span: 18 }}
          layout="horizontal"
        >
        {{{- range .Columns }}}
        {{{- if .IsEdit }}}
        {{{- if and (ne $.Table.TreeParent "") (eq .JavaField $.Table.TreeParent) }}}
          tree
        {{{- else if eq .HtmlType "input" }}}
        <Form.Item
            name="cellphone"
            label="{{{.ColumnName}}}"
            rules={[{ required: true, message: '请输入{{{.ColumnComment}}}' }]}
          >
            <Input />
          </Form.Item>
        {{{- else if eq .HtmlType "number" }}}
        number
        {{{- else if eq .HtmlType "textarea" }}}
        textarea
        {{{- else if eq .HtmlType "checkbox" }}}
        checkbox
        {{{- else if eq .HtmlType "select" }}}
        select
        {{{- else if eq .HtmlType "radio" }}}
        radio
        {{{- else if eq .HtmlType "datetime" }}}
        datetime
        {{{- else if eq .HtmlType "imageUpload" }}}
        imageUpload
        {{{- end }}}
        {{{- end }}}
        {{{- end }}}
        </Form>
      </Modal>
    </>
  );
}
