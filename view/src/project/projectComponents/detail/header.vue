<template>
  <div>
    <el-collapse @change="handleChange">
      <el-collapse-item title="点击查看项目详情">
        <detail :project="project"></detail>
      </el-collapse-item>
    </el-collapse>
  </div>
</template>

<script>
import Detail from './detail'

export default {
  name: 'detailHeader',
  components: {
    Detail
  },
  props: {
    project: Object
  },
  methods: {
    handleChange (val) {
      console.log(val)
    }
  }
}
</script>

<style scoped>

</style>
