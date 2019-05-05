<template>
  <div>
    <project-header></project-header>
    <project-list></project-list>
  </div>
</template>

<script>
import ProjectHeader from './components/header'
import ProjectList from './components/list'

export default {
  name: 'index',
  components: {
    ProjectHeader,
    ProjectList
  }
}
</script>

<style scoped>

</style>
