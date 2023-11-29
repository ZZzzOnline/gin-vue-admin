<template>
  <div>
    <warning-bar
      :title="t('view.dictionary.sysDictionary.dictNote')"
    />
    <div class="dict-box flex gap-4">
      <div class="w-64 bg-white p-4">
        <div class="flex justify-between items-center">
          <span class="text font-bold">字典列表</span>
          <el-button
            type="primary"
            @click="openDialog"
          >
            {{ t('general.add') }}
          </el-button>
        </div>
        <el-scrollbar
          class="mt-4"
          style="height: calc(100vh - 300px)"
        >
          <div
            v-for="dictionary in dictionaryData"
            :key="dictionary.ID"
            class="rounded flex justify-between items-center px-2 py-4 cursor-pointer mt-2 hover:bg-blue-50 hover:text-gray-800 group bg-gray-50"
            :class="selectID === dictionary.ID && 'active'"
            @click="toDetail(dictionary)"
          >
            <span class="max-w-[160px] truncate">{{ dictionary.name }}</span>
            <div>
              <el-icon
                class="group-hover:text-blue-500"
                :class="selectID === dictionary.ID ? 'text-white-800':'text-blue-500'"
                @click.stop="updateSysDictionaryFunc(dictionary)"
              >
                <Edit />
              </el-icon>
              <el-popover
                placement="top"
                width="160"
              >
                <p>{{ t('general.deleteConfirm') }}</p>
                <div style="text-align: right; margin-top: 8px;">
                  <el-button
                    type="primary"
                    link
                    @click="dictionary.visible = false"
                  >{{ t('general.cancel') }}</el-button>
                  <el-button
                    type="primary"
                    @click="deleteSysDictionaryFunc(dictionary)"
                  >{{ t('general.confirm') }}</el-button>
                </div>
                <template #reference>
                  <el-icon
                    class="ml-2 group-hover:text-red-500"
                    :class="selectID === dictionary.ID ? 'text-white-800':'text-red-500'"
                  >
                    <Delete />
                  </el-icon>
                </template>
              </el-popover>
            </div>
          </div>
        </el-scrollbar>
      </div>
      <div class="flex-1 bg-white">
        <sysDictionaryDetail :sys-dictionary-i-d="selectID" />
      </div>
    </div>
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      :title="type==='create'?t('view.dictionary.sysDictionary.addDictionary'):t('view.dictionary.sysDictionary.editDictionary')"
    >
      <el-form
        ref="dialogForm"
        :model="formData"
        :rules="rules"
        label-width="110px"
      >
        <el-form-item
        :label="t('view.dictionary.sysDictionary.dictName')"
          prop="name"
        >
          <el-input
            v-model="formData.name"
            :placeholder="t('view.dictionary.sysDictionary.enterDictName')"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
        <el-form-item
        :label="t('view.dictionary.sysDictionary.dictNameEn')"
          prop="type"
        >
          <el-input
            v-model="formData.type"
            :placeholder="t('view.dictionary.sysDictionary.enterDictNameEn')"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
        <el-form-item
        :label="t('view.dictionary.sysDictionary.status')"
          prop="status"
          required
        >
          <el-switch
            v-model="formData.status"
            :active-text="t('general.enable')"
            :inactive-text="t('general.disable')"
          />
        </el-form-item>
        <el-form-item
        :label="t('general.description')"
          prop="desc"
        >
          <el-input
            v-model="formData.desc"
            :placeholder="t('view.dictionary.sysDictionary.enterDescription')"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">{{ t('general.close') }}</el-button>
          <el-button

            type="primary"
            @click="enterDialog"
          >{{ t('general.confirm') }}</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createSysDictionary,
  deleteSysDictionary,
  updateSysDictionary,
  findSysDictionary,
  getSysDictionaryList,
} from '@/api/sysDictionary' // 此处请自行替换地址
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'

import sysDictionaryDetail from './sysDictionaryDetail.vue'
import { Edit, Plus } from '@element-plus/icons-vue'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage
defineOptions({
  name: 'SysDictionary',
})

const selectID = ref(1)

const formData = ref({
  name: null,
  type: null,
  status: true,
  desc: null,
})
const rules = ref({
  name: [
    {
      required: true,
      message: t('view.dictionary.sysDictionary.enterDictName'),
      trigger: 'blur',
    },
  ],
  type: [
    {
      required: true,
      message: t('view.dictionary.sysDictionary.enterDictNameEn'),
      trigger: 'blur',
    },
  ],
  desc: [
    {
      required: true,
      message: t('view.dictionary.sysDictionary.enterDescription'),
      trigger: 'blur',
    },
  ],
})

const dictionaryData = ref([])

// 查询
const getTableData = async() => {
  const res = await getSysDictionaryList()
  if (res.code === 0) {
    dictionaryData.value = res.data
  }
}

getTableData()

const toDetail = (row) => {
  selectID.value = row.ID
}

const dialogFormVisible = ref(false)
const type = ref('')
const updateSysDictionaryFunc = async(row) => {
  const res = await findSysDictionary({ ID: row.ID, status: row.status })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.resysDictionary
    dialogFormVisible.value = true
  }
}
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    name: null,
    type: null,
    status: true,
    desc: null,
  }
}
const deleteSysDictionaryFunc = async(row) => {
  row.visible = false
  const res = await deleteSysDictionary({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: t('general.deleteSuccess'),
    })
    getTableData()
  }
}

const dialogForm = ref(null)
const enterDialog = async() => {
  dialogForm.value.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createSysDictionary(formData.value)
        break
      case 'update':
        res = await updateSysDictionary(formData.value)
        break
      default:
        res = await createSysDictionary(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage.success('操作成功')
      closeDialog()
      getTableData()
    }
  })
}
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}
</script>

<style>
.dict-box{
  height: calc(100vh - 240px);
}
.active {
  background-color: var(--el-color-primary) !important;
  color: #fff;
}
</style>
