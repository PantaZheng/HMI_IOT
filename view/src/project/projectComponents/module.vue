<template>
  <div>
    <div v-for="data in module" :key="data.moId">
      <div>{{data.cTime}}</div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'projectModule',
  props: {
    module: Array
  }
}
</script>

<style scoped>

</style>
