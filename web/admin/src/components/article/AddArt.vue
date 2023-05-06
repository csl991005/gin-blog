<template>
    <div>
        <a-card>
            <h3>{{ id ? '编辑文章' : '新增文章' }}</h3>
            <a-form-model
                :model="artInfo"
                ref="artInfoRef"
                :rules="artInfoRules"
                :hideRequiredMark="true"
            >
                <a-form-model-item label="文章标题" prop="title">
                    <a-input
                        style="width: 300px"
                        v-model="artInfo.title"
                    ></a-input>
                </a-form-model-item>
                <a-form-model-item label="文章分类" prop="cid">
                    <a-select
                        style="width: 200px"
                        v-model="artInfo.cid"
                        placeholder="请选择分类"
                        @change="cateChange"
                    >
                        <a-select-option
                            v-for="item in Catelist"
                            :key="item.id"
                            :value="item.id"
                            >{{ item.name }}</a-select-option
                        >
                    </a-select>
                </a-form-model-item>
                <a-form-model-item label="文章描述" prop="desc">
                    <a-input type="textarea" v-model="artInfo.desc"></a-input>
                </a-form-model-item>
                <a-form-model-item label="文章缩略图" prop="img">
                    <a-upload
                        listType="picture"
                        name="file"
                        :action="upUrl"
                        :headers="headers"
                        @change="upChange"
                    >
                        <a-button> <a-icon type="upload" /> 点击上传 </a-button>
                        <br />
                        <template v-if="id">
                            <img
                                :src="artInfo.img"
                                style="width: 120px; height: 100px"
                            />
                        </template>
                    </a-upload>
                </a-form-model-item>
                <a-form-model-item label="文章内容" prop="content">
                    <div style="border: 1px solid #ccc">
                        <Toolbar
                            style="border-bottom: 1px solid #ccc"
                            :editor="editor"
                            :defaultConfig="toolbarConfig"
                            :mode="mode"
                        />
                        <Editor
                            style="height: 300px; line-height: 100%"
                            v-model="artInfo.content"
                            :defaultConfig="editorConfig"
                            :mode="mode"
                            @onCreated="onCreated"
                        />
                    </div>
                </a-form-model-item>
                <a-form-model-item>
                    <a-button
                        type="danger"
                        style="margin-right: 15px"
                        @click="artOk(artInfo.id)"
                        >{{ artInfo.id ? '更新' : '提交' }}</a-button
                    >
                    <a-button type="primary" @click="addCancel">取消</a-button>
                </a-form-model-item>
            </a-form-model>
        </a-card>
    </div>
</template>
  
<script>
import { Url } from '../../plugin/http'

import { Editor, Toolbar } from '@wangeditor/editor-for-vue'

export default {
    components: { Editor, Toolbar },
    props: ['id'],
    data() {
        return {
            editor: null,
            toolbarConfig: {},
            editorConfig: { MENU_CONF: {}, placeholder: '请输入文章内容' },
            mode: 'default', // or 'simple'
            artInfo: {
                id: 0,
                title: '',
                cid: undefined,
                desc: '',
                content: '',
                img: ''
            },
            Catelist: [],
            upUrl: Url + '/upload',
            headers: {},
            fileList: [],
            artInfoRules: {
                title: [
                    {
                        required: true,
                        message: '请输入文章标题',
                        trigger: 'blur'
                    }
                ],
                cid: [
                    {
                        required: true,
                        message: '请选择文章分类',
                        trigger: 'change'
                    }
                ],
                desc: [
                    {
                        required: true,
                        message: '请输入文章描述',
                        trigger: 'blur'
                    },
                    {
                        max: 120,
                        message: '描述最多可写120个字符',
                        trigger: 'change'
                    }
                ],
                img: [
                    {
                        required: true,
                        message: '请选择文章缩略图',
                        trigger: 'blur'
                    }
                ],
                content: [
                    {
                        required: true,
                        message: '请输入文章内容',
                        trigger: 'blur'
                    }
                ]
            }
        }
    },
    created() {
        this.getCateList()
        this.headers = {
            Authorization: `Bearer ${window.sessionStorage.getItem('token')}`
        }
        if (this.id) {
            this.getArtInfo(this.id)
        }
        this.editorConfig.MENU_CONF['uploadImage'] = {
            server: 'http://localhost:8080/api/v1/upload',
            fieldName: 'custom-field-name',
            headers: {
                Authorization: `Bearer ${window.sessionStorage.getItem(
                    'token'
                )}`
            },
            customUpload: this.customUpload
        }
    },
    beforeDestroy() {
        const editor = this.editor
        if (editor == null) return
        editor.destroy() // 组件销毁时，及时销毁编辑器
    },
    methods: {
        async customUpload(file, insertFn) {
            let formData = new FormData()
            formData.append('file', file)
            const { data: res } = await this.$http.post('/upload', formData)
            insertFn(res.url)
        },
        onCreated(editor) {
            this.editor = Object.seal(editor) // 一定要用 Object.seal() ，否则会报错
        },
        // 查询文章信息
        async getArtInfo(id) {
            const { data: res } = await this.$http.get(`article/info/${id}`)
            if (res.status != 200) {
                return this.$message.error(res.message)
            }
            this.artInfo = res.data
            this.artInfo.id = res.data.ID
        },
        // 获取分类列表
        async getCateList() {
            const { data: res } = await this.$http.get('category')
            if (res.status != 200) {
                return this.$message.error(res.message)
            }
            this.Catelist = res.data
        },
        // 选择分类
        cateChange(value) {
            this.artInfo.cid = value
        },
        // 上传图片
        upChange(info) {
            if (info.file.status !== 'uploading') {
            }
            if (info.file.status === 'done') {
                this.$message.success(`图片上传成功`)
                const imgUrl = info.file.response.url
                this.artInfo.img = imgUrl
            } else if (info.file.status === 'error') {
                this.$message.error(`图片上传失败`)
            }
        },
        // 提交或更新文章
        artOk(id) {
            this.$refs.artInfoRef.validate(async (valid) => {
                if (!valid) {
                    return this.$message.error(
                        '参数验证未通过，请按要求录入文章内容'
                    )
                }
                if (id === 0) {
                    const { data: res } = await this.$http.post(
                        'article/add',
                        this.artInfo
                    )
                    if (res.status != 200) {
                        return this.$message.error(res.message)
                    }
                    this.$message.success('添加文章成功')
                    this.$router.push('/artlist')
                } else {
                    const { data: res } = await this.$http.put(
                        `article/${id}`,
                        this.artInfo
                    )
                    if (res.status != 200) {
                        return this.$message.error(res.message)
                    }
                    this.$message.success('更新文章成功')
                    this.$router.push('/artlist')
                }
            })
        },
        // 取消按钮
        addCancel() {
            this.$refs.artInfoRef.resetFields()
        }
    }
}
</script>
  
<style src="@wangeditor/editor/dist/css/style.css"></style>