<template>
    <div class="page-bg">
      <div class="weui-cells">
        <div class="weui-cell">
          <div class="weui-cell__hd"><label class="weui-label">项目名称</label></div>
          <div class="weui-cell__bd ">
            <input class="weui-input" pattern="[0-9]*" placeholder="请输入项目名称" type="name">
          </div>
        </div>
        <div class="weui-cell">
          <div class="weui-cell__hd">
            <label class="weui-label">开始时间</label>
          </div>
          <div class="weui-cell__bd">
            <div class="weui_cell_bd weui_cell_primary">
              <input class="weui_input" type="date" value="">
            </div>
          </div>
        </div>
        <div class="weui-cell">
          <div class="weui-cell__hd">
            <label class="weui-label">结束时间</label>
          </div>
          <div class="weui-cell__bd">
            <div class="weui_cell_bd weui_cell_primary">
              <input class="weui_input" type="date" value="">
            </div>
          </div>
        </div>
      </div>
      <div class="weui_cells margin">
        <div class="weui_cell weui_cell_select">
          <div class="weui_cell_bd weui_cell_primary">
            <select class="weui_select">
              <option value="">请选择项目总负责人</option>
            </select>
          </div>
        </div>
        <div class="weui_cell weui_cell_select">
          <div class="weui_cell_bd weui_cell_primary">
            <select class="weui_select">
              <option value="">请选择项目类型</option>
              <option value="h">横向项目</option>
              <option value="z">纵向项目</option>
            </select>
          </div>
        </div>
        <div class="weui_cell weui_cell_select">
          <div class="weui_cell_bd weui_cell_primary">
            <select class="weui_select">
              <option value="">请选择参与项目的老师</option>
            </select>
          </div>
        </div>
      </div>
      <div class="weui-cells weui-cells_form margin">
        <div class="weui-cell">
          <div class="weui-cell__bd">
            <textarea class="weui-textarea" placeholder="请输入验收目标" rows="3"></textarea>
            <div class="weui-textarea-counter"><span>0</span>/<i>100</i></div>
          </div>
          <i class="weui-icon-clear"></i>
        </div>
      </div>
      <div class="weui-cells weui-cells_form margin">
        <div class="weui-cell">
          <div class="weui-cell__bd">
            <textarea class="weui-textarea" placeholder="请输入项目主页主要内容" rows="3"></textarea>
            <div class="weui-textarea-counter"><span>0</span>/<i>100</i></div>
          </div>
          <i class="weui-icon-clear"></i>
        </div>
      </div>
      <div class="weui-btn-area">
        <a href="javascript:;" class="weui-btn weui-btn_primary">提交</a>
      </div>
    </div>
</template>

<script>
export default {
  name: 'assistantForm'
}
</script>

<style lang="stylus" scoped>
  .margin
    margin-top 15px
</style>
