<template>
    <div>
      <div class="v-center header">
        {{head}}
      </div>
      <div v-for="data in PR" :key="data.PRID">
        <div class="v-center header">
          <div>{{data.name}}</div>
          <div>{{data.tel}}</div>
        </div>
        <div v-for="module in data.module" :key="module.moId">
          <div class="v-center header">
            {{module.moName}}
          </div>
          <div v-for="stu in data.stu" :key="stu">
            {{stu}}
          </div>
        </div>
      </div>
    </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'projectFrame',
  data () {
    return {
      PR: String,
      head: Array
    }
  },
  methods: {
    getFramework () {
      axios.get('/api/framework.json')
        .then(this.getFrameworkSucc)
    },
    getFrameworkSucc (res) {
      res = res.data[0].frame
      if (res) {
        const data = res
        this.PR = data.PR
        this.head = data.head
        console.log(this.PR)
      }
    }
  },
  mounted () {
    this.getFramework()
  }
}
</script>

<style lang="stylus" scoped>
  .header
    color $headerFontColor
    background $headerBgColor
    position: relative
    line-height: $headerHeight
</style>
