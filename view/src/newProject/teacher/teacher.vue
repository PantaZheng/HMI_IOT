<template>
  <div>
    <teacher-header></teacher-header>
    <teacher-list :list="projectList"></teacher-list>
  </div>
</template>

<script>
import TeacherHeader from './components/header'
import TeacherList from './components/list'
import axios from 'axios'

export default {
  name: 'assistant',
  components: {
    TeacherHeader,
    TeacherList
  },
  data () {
    return {
      projectList: []
    }
  },
  methods: {
    getProjectList () {
      axios.get('/api/projectList.json')
        .then(this.getProjectListSucc)
    },
    getProjectListSucc (res) {
      res = res.data
      if (res) {
        this.projectList = res
      }
    }
  },
  mounted () {
    this.getProjectList()
  }
}
</script>

<style scoped>

</style>
