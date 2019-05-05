<template>
    <div>
      <bind-header></bind-header>
      <bind-form></bind-form>
    </div>
</template>

<script>
import BindHeader from '@/bind/components/header'
import BindForm from '@/bind/components/form'
// import axios from 'axios'
export default {
  name: 'index',
  components: {
    BindHeader,
    BindForm
  }
}
</script>

<style lang="stylus" scoped>
</style>
