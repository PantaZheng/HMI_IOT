<template>
  <div class="v-center header page-bg page-hd-title">
    绑定
  </div>
</template>

<script>
export default {
  name: 'BindHeader'
}
</script>

<style lang="stylus" scoped>
  @import "~styles/varibles.styl"
  .header
    color $headerFontColor
    background $headerBgColor
    position: relative
    line-height: $headerHeight
    font-size $headerFontSize
</style>
