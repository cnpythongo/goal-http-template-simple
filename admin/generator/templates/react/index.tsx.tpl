import { Button, Form, Input, Modal, Popconfirm, InputNumber, Table, Radio, Checkbox, Upload, DatePicker, GetProp, UploadProps  } from 'antd';
import type { TableProps } from 'antd';
import { useEffect, useState } from 'react';
import * as Icons from '@ant-design/icons';
import api, {
  {{{.EntityName}}}ListParams,
  {{{.EntityName}}}Item,
  {{{.EntityName}}}CreateBody,
  {{{.EntityName}}}UpdateBody
} from '@/api{{{.GenPath}}}';
import dayjs from 'dayjs';

const { TextArea } = Input;
const { RangePicker } = DatePicker;

type FileType = Parameters<GetProp<UploadProps, 'beforeUpload'>>[0];

const getBase64 = (img: FileType, callback: (url: string) => void) => {
  const reader = new FileReader();
  reader.addEventListener('load', () => callback(reader.result as string));
  reader.readAsDataURL(img);
};

const beforeUpload = (file: FileType) => {
  const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png';
  if (!isJpgOrPng) {
    message.error('You can only upload JPG/PNG file!');
  }
  const isLt2M = file.size / 1024 / 1024 < 2;
  if (!isLt2M) {
    message.error('Image must smaller than 2MB!');
  }
  return isJpgOrPng && isLt2M;
};


// 类型定义
type TableColumns = TableProps<{{{.EntityName}}}Item>['columns'];

type TableRowSelection<T extends object = object> =
  TableProps<T>['rowSelection'];

// 主函数
export default function {{{.EntityName}}}Page() {
  // 表单属性定义
  const [searchForm] = Form.useForm();
  const [editForm] = Form.useForm();
  const [editRecord, setEditRecord] = useState<any>(null);
  const [showEditModal, setShowEditModal] = useState<boolean>(false);
  const [showDataModal, setShowDataModal] = useState<boolean>(false);
  const getDateInterval = (days: number) => {
      return [dayjs().subtract(days, 'days').unix(), dayjs().unix()];
    };
    const [initial_start, initial_end] = getDateInterval(30);
    const [total, setTotal] = useState(0);
    const [page, setPage] = useState(1);
    const [pageLimit, setPageLimit] = useState(10);
    const [dateFilter, setDateFilter] = useState({
      create_time_start: initial_start,
      create_time_end: initial_end
    });

  // 数据表格属性定义
  const [tableData, setTableData] = useState<any>([]);
  const [selectedRowKeys, setSelectedRowKeys] = useState([]);

  // 上传组件
  const [loading, setLoading] = useState(false);
  const [imageUrl, setImageUrl] = useState<string>();
  const handleChange: UploadProps['onChange'] = (info) => {
      if (info.file.status === 'uploading') {
        setLoading(true);
        return;
      }
      if (info.file.status === 'done') {
        // Get this url from response in real world.
        getBase64(info.file.originFileObj as FileType, (url) => {
          setLoading(false);
          setImageUrl(url);
        });
      }
    };

  const uploadButton = (
      <button style={{ border: 0, background: 'none' }} type="button">
        {loading ? <Icons.LoadingOutlined /> : <Icons.PlusOutlined />}
        <div style={{ marginTop: 8 }}>Upload</div>
      </button>
    );
    
  // 请求参数定义
  const [params, setParams] = useState<{{{.EntityName}}}ListParams>({
      ...dateFilter,
      page,
      limit: pageLimit
    });

  // 事件定义
  // 复选框事件
  const onSelectChange = newSelectedRowKeys => {
    console.log('selectedRowKeys changed: ', newSelectedRowKeys);
    setSelectedRowKeys(newSelectedRowKeys);
  };

  const rowSelection: TableRowSelection<{{{.EntityName}}}Item> = {
    selectedRowKeys,
    onChange: onSelectChange
  };

  const onSearchDateChange = (_, dateStrings) => {
      setDateFilter({
        create_time_start: dayjs(dateStrings[0]).unix(),
        create_time_end: dayjs(dateStrings[1]).unix()
      });
    };

  // 查询
  const onSearch = () => {
      const formValues = searchForm.getFieldsValues();
      setParams({
        ...params,
        ...formValues,
        ...dateFilter,
        page: 1,
        limit: pageLimit
      });
    };

  // 详情
  const onDetail = (record: {{{.EntityName}}}Item) => {
    detail(record.id);
  };

  // 创建
  const onCreate = () => {
    showEditFormModal();
  };

  // 更新
  const onUpdate = (record: {{{.EntityName}}}Item) => {
    showEditFormModal(record);
  };

  // 删除
  const onDelete = (record: {{{.EntityName}}}Item) => {
    del([record.id]);
  };

  // 批量删除
  const onBatchDelete = () => {
    const ids = selectedRowKeys;
    del(ids);
  };

  // 弹出层事件
  const showEditFormModal = (record?: {{{.EntityName}}}Item) => {
    setEditRecord(record);
    editForm.resetFields();
    if (record) {
      editForm.setFieldsValue({ ...record });
    }
    setShowEditModal(true);
  };

  const handleDataFormOk = () => {
    if (editRecord) {
      update();
    } else {
      create();
    }
    setShowEditModal(false);
  };

  const handleDataFormCancel = () => {
    setShowEditModal(false);
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
      .create({ ...editForm.getFieldsValue() })
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
      .update({ ...editRecord, ...editForm.getFieldsValue() })
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
  const paginationParams = {
      total,
      current: page,
      pageLimit,
      showTotal: (v: any) => `共${v}条记录`,
      onChange: (pNumber: any, pSize: any) => {
        setPage(pNumber);
        setPageLimit(pSize);
        setParams({ ...params, page: pNumber, limit: pSize });
      }
    };

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
            pagination={paginationParams}
          />
        </div>
      </div>

      <Modal
        title={editRecord ? '编辑数据' : '新增数据'}
        centered={true}
        maskClosable={false}
        open={showEditModal}
        onOk={handleDataFormOk}
        onCancel={handleDataFormCancel}
        width={600}
      >
        <Form
          form={editForm}
          name="data-form"
          labelCol={{ span: 5 }}
          wrapperCol={{ span: 18 }}
          layout="horizontal"
        >
        {{{- range .Columns }}}
        {{{- if .IsEdit }}}
        <Form.Item
          name="{{{.ColumnName}}}"
          label="{{{.ColumnComment}}}"
          rules={[{ required: true, message: '请输入{{{.ColumnComment}}}' }]}
        >
        {{{- if and (ne $.Table.TreeParent "") (eq .GoField $.Table.TreeParent) }}}
            <TreeSelect
              treeDefaultExpandAll={true}
              treeLine={true}
              treeData={[]}
              disabled={disableTreeSelect}
              fieldNames={{
                label: 'name',
                value: 'id',
                children: 'children'
              }}
            />
        {{{- else if eq .HtmlType "input" }}}
          <Input />
        {{{- else if eq .HtmlType "number" }}}
          <InputNumber min={1} max={10} defaultValue={3} onChange={onChange} />
        {{{- else if eq .HtmlType "textarea" }}}
          <TextArea />
        {{{- else if eq .HtmlType "checkbox" }}}
            <Checkbox />
        {{{- else if eq .HtmlType "select" }}}
        <Select options={[]} />
        {{{- else if eq .HtmlType "radio" }}}
        <Radio.Group onChange={onUpdateGenTpl}>
          <Radio value={'crud'}>单表</Radio>
          <Radio value={'tree'}>树表</Radio>
        </Radio.Group>
        {{{- else if eq .HtmlType "date" }}}
        <DatePicker onChange={onChange} />
        <RangePicker showTime />
        {{{- else if eq .HtmlType "datetime" }}}
        <Form.Item label="选择日期">
          <RangePicker
            onChange={onSearchDateChange}
            allowClear
            value={[
              dayjs.unix(dateFilter.create_time_start),
              dayjs.unix(dateFilter.create_time_end)
            ]}
            format="YYYY-MM-DD"
            separator="~"
          />
        </Form.Item>
        {{{- else if eq .HtmlType "imageUpload" }}}
        <Upload
            name="avatar"
            listType="picture-card"
            className="avatar-uploader"
            showUploadList={false}
            action="https://660d2bd96ddfa2943b33731c.mockapi.io/api/upload"
            beforeUpload={beforeUpload}
            onChange={handleChange}
          >
            {imageUrl ? <img src={imageUrl} alt="avatar" style={{ width: '100%' }} /> : uploadButton}
          </Upload>
        {{{- end }}}
        </Form.Item>
        {{{- end }}}
        {{{- end }}}
        </Form>
      </Modal>
      
      <Modal
          title={'数据详情'}
          centered={true}
          maskClosable={false}
          open={showDataModal}
          onCancel={() => {
            setShowDataModal(false);
          }}
          width={600}
        >
        {editRecord && (
        <>
        {{{- range .Columns }}}
          <div className="flex flex-row mb-[1px]">
            <div className="basis-1/4 text-right bg-slate-200">{{{ .ColumnComment }}}：</div>
            <div className="basis-3/4">{editRecord.{{{ .ColumnName }}}}</div>
          </div>
        {{{- end }}}
        </>
        )}
        </Modal>
    </>
  );
}
