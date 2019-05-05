<template>
  <div>
    <div v-for="data in project" :key="data.pId" v-if="data.pId == 1">
      <div class="weui-cell">{{data.content}}</div>
      <div class="weui-cells">
        <div class="weui-cell">
          <div class="weui-cell__hd">
            <label class="weui-label">开始时间</label>
          </div>
          <div class="weui-cell__bd cellMargin">
            <div class="weui_cell_bd weui_cell_primary">
              {{data.sTime}}
            </div>
          </div>
        </div>
        <div class="weui-cell">
          <div class="weui-cell__hd">
            <label class="weui-label">结束时间</label>
          </div>
          <div class="weui-cell__bd cellMargin">
            <div class="weui_cell_bd weui_cell_primary">
              {{data.eTime}}
            </div>
          </div>
        </div>
        <div class="weui-cell">
          <div class="weui-cell__hd">
            <label class="weui-label">项目总负责人</label>
          </div>
          <div class="weui-cell__bd cellMargin">
            <div class="weui_cell_bd weui_cell_primary">
              {{data.charge}}
            </div>
          </div>
        </div>
        <div class="weui-cell">
          <div class="weui-cell__hd">
            <label class="weui-label">项目参与老师</label>
          </div>
          <div class="weui-cell__bd cellMargin">
            <div class="weui_cell_bd weui_cell_primary">
              <span  v-for="teacher in data.teacher" :key="teacher.name" class="nameMargin">{{teacher.name}}</span>
            </div>
          </div>
        </div>
        <div class="weui-cell">
          <div class="weui-cell__hd">
            <label class="weui-label">项目验收目标</label>
          </div>
          <div class="weui-cell__bd cellMargin">
            <div class="weui_cell_bd weui_cell_primary" v-for="target in data.target" :key="target" >
              {{target}}
            </div>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'detail',
  props: {
    project: Object
  }
}
</script>

<style lang="stylus" scoped>
  .cellMargin
    margin-left .7rem
  .nameMargin
    margin-right .5rem
</style>
