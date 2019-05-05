<template>
  <div>
    <div class="weui_tab_bd_item ">
      <div class="weui-news">
        <ul class="weui-news-list">
          <li v-for="data in projectList" :key="data.pId">
            <router-link :to="'/project/'+data.pId" class="weui-news-item">
            <div class="weui-news-inners">
              <div class="weui-news-title">{{data.pName}}<span style="float: right;">{{data.sTime}}</span></div>
              <div class="weui-news-text">
                <p class="weui-news-p">{{data.content}}</p>
              </div>
              <div class="weui-news-info">
                <div class="weui-news-infoitem">
                  <span>
                    <svg class="icon iconMargin" aria-hidden="true">
                      <use xlink:href="#iconicon_signal"></use>
                    </svg>
                    {{data.head}}
                  </span>
                </div>
              </div>
            </div>
            </router-link>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'list',
  data () {
    return {
      projectList: Array
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
  .iconMargin
    margin-right .25rem
</style>
