<template class="page-bg">
  <div >
    <div class="weui-cells form weui-cells_form">
      <div class="weui-cell">
        <div class="weui-cell__hd"><label class="weui-label req">姓名</label></div>
        <div class="weui-cell__bd ">
          <input class="weui-input" pattern="[0-9]*" placeholder="请准确输入您的姓名" type="name">
        </div>
      </div>
      <div class="weui-cell">
        <div class="weui-cell__hd">
          <label class="weui-label">手机号</label>
        </div>
        <div class="weui-cell__bd">
          <input class="weui-input" placeholder="请输入手机号" type="tel">
        </div>
      </div>
      <div class="weui-cell">
        <div class="weui-cell__hd">
          <label class="weui-label req">身份证号</label>
        </div>
        <div class="weui-cell__bd">
          <input class="weui-input " placeholder="请输入您的身份证号" type="tel">
        </div>
      </div>
    </div>

    <div class="weui-btn-area">
      <a href="javascript:;" class="weui-btn weui-btn_primary" id="sd2">确定</a>
    </div>
  </div>
</template>

<script>
export default {
  name: 'BindForm'
}
</script>

<style lang="stylus" scoped>
  .form
   margin-top 0
  .btn
    background  #C0BFC4
</style>
