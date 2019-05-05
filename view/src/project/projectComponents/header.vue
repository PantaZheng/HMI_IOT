<template>
  <div class="header">
    <div class="header-left back-icon">
      <svg class="icon " aria-hidden="true">
        <use xlink:href="#iconleft"></use>
      </svg>
    </div>
    <div class="header-title v-center">
      {{pName}}
    </div>
    <div class="header-right">
      <svg class="icon " aria-hidden="true">
        <use xlink:href="#icondian2"></use>
      </svg>
    </div>
  </div>
</template>

<script>
export default {
  name: 'projectHeader',
  props: {
    pName: String
  }
}
</script>

<style lang="stylus" scoped>
  @import "~styles/varibles.styl"
  .header
    display flex
    color $headerFontColor
    background $headerBgColor
    position: relative
    line-height: $headerHeight
  .arrow
    float left
  .header-left
    margin-left .4rem
    width: .64rem
    float left
    .back-icon
      text-align: center
      font-size: .4rem
  .header-title
    flex: 1
    font-size $headerFontSize
  .header-right
    min-width: 2rem
</style>
