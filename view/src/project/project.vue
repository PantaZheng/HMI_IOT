<template>
  <div>
    <project-header :pName="pName"></project-header>
    <project-detail :project="project"></project-detail>
    <project-module :module="module"></project-module>
  </div>
</template>

<script>
import ProjectHeader from './projectComponents/header'
import ProjectDetail from './projectComponents/detail/header'
import ProjectModule from './projectComponents/module'
import axios from 'axios'

export default {
  name: 'project',
  components: {
    ProjectHeader,
    ProjectDetail,
    ProjectModule
  },
  data () {
    return {
      project: this.project,
      module: [],
      pName: this.pName
    }
  },
  methods: {
    getProject () {
      axios.get('/api/project.json')
        .then(this.getProjectSucc)
    },
    getProjectSucc (res) {
      if (res) {
        this.project = res
        const data = res.data
        this.module = data.module
        this.pName = data.pName
      }
    }
  },
  mounted () {
    this.getProject()
  }
}
</script>

<style scoped>

</style>
