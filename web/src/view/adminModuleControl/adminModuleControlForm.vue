<template>
  <div>
    <div class="gva-form-box">
      <el-form
        ref="elFormRef"
        :model="formData"
        label-position="right"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item
          label="邮箱地址:"
          prop="email"
        >
          <el-input
            v-model="formData.email"
            :clearable="false"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item
          label="账号唯一标识:"
          prop="accountId"
        >
          <el-input
            v-model.number="formData.accountId"
            :clearable="false"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item
          label="玩家标识:"
          prop="playerId"
        >
          <el-input
            v-model="formData.playerId"
            :clearable="false"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item
          label="禁止登录:"
          prop="forbiddenLogin"
        >
          <el-switch
            v-model="formData.forbiddenLogin"
            active-color="#ff4949"
            inactive-color="#13ce66"
            active-text="是"
            inactive-text="否"
            clearable
          />
        </el-form-item>
        <el-form-item
          label="禁止游戏内英雄导出:"
          prop="forbiddenInGameHeroExport"
        >
          <el-switch
            v-model="formData.forbiddenInGameHeroExport"
            active-color="#ff4949"
            inactive-color="#13ce66"
            active-text="是"
            inactive-text="否"
            clearable
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            @click="save"
          >保存</el-button>
          <el-button
            type="primary"
            @click="back"
          >返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createAdminModuleControl,
  updateAdminModuleControl,
  findAdminModuleControl
} from '@/api/adminModuleControl'

defineOptions({
  name: 'AdminModuleControlForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  email: '',
  accountId: 0,
  playerId: '',
  forbiddenLogin: false,
  forbiddenInGameHeroExport: false,
})
// 验证规则
const rule = reactive({
  email: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findAdminModuleControl({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.readminModuleControl
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()
// 保存按钮
const save = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createAdminModuleControl(formData.value)
        break
      case 'update':
        res = await updateAdminModuleControl(formData.value)
        break
      default:
        res = await createAdminModuleControl(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
    }
  })
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style></style>
