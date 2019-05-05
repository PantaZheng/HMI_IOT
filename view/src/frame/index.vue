<template>
  <div>
   <frame-header></frame-header>
    <frame-list :list="projectList"></frame-list>
  </div>
</template>

<script>
import FrameHeader from './components/header'
import FrameList from './components/list'
import axios from 'axios'

export default {
  name: 'BindHeader',
  components: {
    FrameHeader,
    FrameList
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

<style lang="stylus" scoped>
  @import "~styles/varibles.styl"
  .header
    position: relative
    line-height: $headerHeight
</style>
