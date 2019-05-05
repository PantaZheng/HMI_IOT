<template>
  <div >
    <div v-for="item of list" :key="item.pId">
    <router-link :to="'/frame/'+item.pName+'/'+item.pId" class="weui-cell cell" >
      <div class="weui-cell__hd textColor">
          <span class="iconMargin">
            <svg class="icon" aria-hidden="true">
              <use xlink:href="#icon-"></use>
            </svg>
          </span>
        {{item.pName}}
      </div>
      <div class="weui-cell__bd textColor">
        <label class="weui-label right">
          <svg class="icon iconMargin iconColor" aria-hidden="true">
            <use xlink:href="#iconfuzeren3"></use>
          </svg>
            {{item.head}}
        </label>
      </div>
    </router-link>
    </div>
  </div>
</template>

<script>

import CellBox from 'vux/src/components/cell-box/index'
export default {
  name: 'list',
  props: {
    list: Array
  },
  components: {CellBox}
}
</script>

<style lang="stylus" scoped>
  .right
    float: right
  .iconMargin
    margin-right .5rem
  .textColor
    color #0D0D0D
  .iconColor
    color #88d4ff
</style>
