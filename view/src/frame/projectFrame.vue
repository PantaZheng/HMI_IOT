<template>
  <div>
    <project-header :pName="pName"></project-header>
    <framework :pId="pId"></framework>
  </div>
</template>

<script>
import ProjectHeader from './projectFrameComponents/header'
import Framework from './projectFrameComponents/framework'

export default {
  name: 'projectFrame',
  data () {
    return {
      pName: this.$route.params.pName,
      pId: this.$route.params.pId
    }
  },
  components: {
    ProjectHeader,
    Framework
  }
}
</script>

<style scoped>

</style>
